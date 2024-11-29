package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitVisit_recordRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/visit-record/page", Visit_recordPageListHandler)
	r.Get(common.BASE_CONTEXT+"/visit-record", Visit_recordListHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record", UpsertVisit_recordHandler)

	r.Delete(common.BASE_CONTEXT+"/visit-record/{id}", DeleteVisit_recordHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record/batch-delete", batchDeleteVisit_recordHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record/batch-upsert", batchUpsertVisit_recordHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 探视登记
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-record/batch-upsert [post]
func batchUpsertVisit_recordHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Visit_recordTableInfo.Name, model.Visit_record_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视登记
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param visit_start_time query string false "visit_start_time"
// @Param visit_end_time query string false "visit_end_time"
// @Param visitor_name query string false "visitor_name"
// @Param visitor_phone query string false "visitor_phone"
// @Param visitor_id_card query string false "visitor_id_card"
// @Param relationship query string false "relationship"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param check_status query string false "check_status"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Visit_record}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record/page [get]
func Visit_recordPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Visit_record](w, r, common.GetDaprClient(), "o_visit_record", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视登记
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param visit_start_time query string false "visit_start_time"
// @Param visit_end_time query string false "visit_end_time"
// @Param visitor_name query string false "visitor_name"
// @Param visitor_phone query string false "visitor_phone"
// @Param visitor_id_card query string false "visitor_id_card"
// @Param relationship query string false "relationship"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param check_status query string false "check_status"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Visit_record} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record [get]
func Visit_recordListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Visit_record](w, r, common.GetDaprClient(), "o_visit_record", "id")
}

// @Summary save
// @Description save
// @Tags 探视登记
// @Accept       json
// @Param item body model.Visit_record true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-record [post]
func UpsertVisit_recordHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Visit_record
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Visit_record")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Visit_record)
	}

	err = common.DbUpsert[model.Visit_record](r.Context(), common.GetDaprClient(), val, model.Visit_recordTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 探视登记
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-record/{id} [delete]
func DeleteVisit_recordHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Visit_record")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_visit_record", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 探视登记
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-record/batch-delete [post]
func batchDeleteVisit_recordHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Visit_record")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_visit_record", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
