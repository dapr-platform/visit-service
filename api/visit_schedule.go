package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitVisit_scheduleRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/visit-schedule/page", Visit_schedulePageListHandler)
	r.Get(common.BASE_CONTEXT+"/visit-schedule", Visit_scheduleListHandler)

	r.Post(common.BASE_CONTEXT+"/visit-schedule", UpsertVisit_scheduleHandler)

	r.Delete(common.BASE_CONTEXT+"/visit-schedule/{id}", DeleteVisit_scheduleHandler)

	r.Post(common.BASE_CONTEXT+"/visit-schedule/batch-delete", batchDeleteVisit_scheduleHandler)

	r.Post(common.BASE_CONTEXT+"/visit-schedule/batch-upsert", batchUpsertVisit_scheduleHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 探视排班
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule/batch-upsert [post]
func batchUpsertVisit_scheduleHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Visit_scheduleTableInfo.Name, model.Visit_schedule_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视排班
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param total_visitors query string false "total_visitors"
// @Param remaining_visitors query string false "remaining_visitors"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Visit_schedule}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule/page [get]
func Visit_schedulePageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Visit_schedule](w, r, common.GetDaprClient(), "o_visit_schedule", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视排班
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param total_visitors query string false "total_visitors"
// @Param remaining_visitors query string false "remaining_visitors"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Visit_schedule} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule [get]
func Visit_scheduleListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Visit_schedule](w, r, common.GetDaprClient(), "o_visit_schedule", "id")
}

// @Summary save
// @Description save
// @Tags 探视排班
// @Accept       json
// @Param item body model.Visit_schedule true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_schedule} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule [post]
func UpsertVisit_scheduleHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Visit_schedule
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Visit_schedule")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Visit_schedule)
	}

	err = common.DbUpsert[model.Visit_schedule](r.Context(), common.GetDaprClient(), val, model.Visit_scheduleTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 探视排班
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_schedule} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule/{id} [delete]
func DeleteVisit_scheduleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Visit_schedule")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_visit_schedule", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 探视排班
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-schedule/batch-delete [post]
func batchDeleteVisit_scheduleHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Visit_schedule")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_visit_schedule", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
