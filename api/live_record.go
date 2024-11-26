package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitLive_recordRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/live-record/page", Live_recordPageListHandler)
	r.Get(common.BASE_CONTEXT+"/live-record", Live_recordListHandler)

	r.Post(common.BASE_CONTEXT+"/live-record", UpsertLive_recordHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 直播记录
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param schedule_id query string false "schedule_id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param device_id query string false "device_id"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param file_size query string false "file_size"
// @Param stream_id query string false "stream_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Live_record}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /live-record/page [get]
func Live_recordPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Live_record](w, r, common.GetDaprClient(), "o_live_record", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 直播记录
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param schedule_id query string false "schedule_id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param device_id query string false "device_id"
// @Param start_time query string false "start_time"
// @Param end_time query string false "end_time"
// @Param file_size query string false "file_size"
// @Param stream_id query string false "stream_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Live_record} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /live-record [get]
func Live_recordListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Live_record](w, r, common.GetDaprClient(), "o_live_record", "id")
}

// @Summary save
// @Description save
// @Tags 直播记录
// @Accept       json
// @Param item body model.Live_record true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Live_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /live-record [post]
func UpsertLive_recordHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Live_record
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Live_record")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Live_record)
	}

	err = common.DbUpsert[model.Live_record](r.Context(), common.GetDaprClient(), val, model.Live_recordTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
