package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitPatientRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/patient/page", PatientPageListHandler)
	r.Get(common.BASE_CONTEXT+"/patient", PatientListHandler)

	r.Post(common.BASE_CONTEXT+"/patient", UpsertPatientHandler)

	r.Delete(common.BASE_CONTEXT+"/patient/{id}", DeletePatientHandler)

	r.Post(common.BASE_CONTEXT+"/patient/batch-delete", batchDeletePatientHandler)

	r.Post(common.BASE_CONTEXT+"/patient/batch-upsert", batchUpsertPatientHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 病患
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /patient/batch-upsert [post]
func batchUpsertPatientHandler(w http.ResponseWriter, r *http.Request) {

	var entities []map[string]any
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}
	for _, v := range entities {
		if v["id"] == "" {
			v["id"] = common.NanoId()
		}
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.PatientTableInfo.Name, model.Patient_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病患
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param ward_id query string false "ward_id"
// @Param bed_id query string false "bed_id"
// @Param name query string false "name"
// @Param hospital_no query string false "hospital_no"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Patient}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient/page [get]
func PatientPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Patient](w, r, common.GetDaprClient(), "o_patient", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病患
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param ward_id query string false "ward_id"
// @Param bed_id query string false "bed_id"
// @Param name query string false "name"
// @Param hospital_no query string false "hospital_no"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Patient} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient [get]
func PatientListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Patient](w, r, common.GetDaprClient(), "o_patient", "id")
}

// @Summary save
// @Description save
// @Tags 病患
// @Accept       json
// @Param item body model.Patient true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Patient} "object"
// @Failure 500 {object} common.Response ""
// @Router /patient [post]
func UpsertPatientHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Patient
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Patient")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Patient)
	}

	err = common.DbUpsert[model.Patient](r.Context(), common.GetDaprClient(), val, model.PatientTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 病患
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Patient} "object"
// @Failure 500 {object} common.Response ""
// @Router /patient/{id} [delete]
func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Patient")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_patient", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 病患
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /patient/batch-delete [post]
func batchDeletePatientHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of ids is 0"))
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Patient")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_patient", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
