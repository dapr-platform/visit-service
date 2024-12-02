package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitPatient_relativeRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/patient-relative/page", Patient_relativePageListHandler)
	r.Get(common.BASE_CONTEXT+"/patient-relative", Patient_relativeListHandler)

	r.Post(common.BASE_CONTEXT+"/patient-relative", UpsertPatient_relativeHandler)

	r.Delete(common.BASE_CONTEXT+"/patient-relative/{id}", DeletePatient_relativeHandler)

	r.Post(common.BASE_CONTEXT+"/patient-relative/batch-delete", batchDeletePatient_relativeHandler)

	r.Post(common.BASE_CONTEXT+"/patient-relative/batch-upsert", batchUpsertPatient_relativeHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 病患-家属关联
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /patient-relative/batch-upsert [post]
func batchUpsertPatient_relativeHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Patient_relative
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Patient_relative")
	if exists {
		for _, v := range entities {
			_, err1 := beforeHook(r, v)
			if err1 != nil {
				common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
				return
			}
		}

	}
	for _, v := range entities {
		if v.ID == "" {
			v.ID = common.NanoId()
		}
	}

	err = common.DbBatchUpsert[model.Patient_relative](r.Context(), common.GetDaprClient(), entities, model.Patient_relativeTableInfo.Name, model.Patient_relative_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病患-家属关联
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param create_time query string false "create_time"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Patient_relative}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative/page [get]
func Patient_relativePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Patient_relative](w, r, common.GetDaprClient(), "o_patient_relative", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病患-家属关联
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param create_time query string false "create_time"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Patient_relative} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative [get]
func Patient_relativeListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Patient_relative](w, r, common.GetDaprClient(), "o_patient_relative", "id")
}

// @Summary save
// @Description save
// @Tags 病患-家属关联
// @Accept       json
// @Param item body model.Patient_relative true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Patient_relative} "object"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative [post]
func UpsertPatient_relativeHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Patient_relative
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Patient_relative")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Patient_relative)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	err = common.DbUpsert[model.Patient_relative](r.Context(), common.GetDaprClient(), val, model.Patient_relativeTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 病患-家属关联
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Patient_relative} "object"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative/{id} [delete]
func DeletePatient_relativeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Patient_relative")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_patient_relative", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 病患-家属关联
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /patient-relative/batch-delete [post]
func batchDeletePatient_relativeHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Patient_relative")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_patient_relative", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
