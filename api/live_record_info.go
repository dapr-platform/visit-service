package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"time"
)

var _ = time.Now()

func InitLive_record_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/live-record-info/page", Live_record_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/live-record-info", Live_record_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 直播记录信息视图
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
// @Param stream_url_suffix query string false "stream_url_suffix"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param status query string false "status"
// @Param patient_name query string false "patient_name"
// @Param patient_ward_name query string false "patient_ward_name"
// @Param patient_bed_no query string false "patient_bed_no"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Live_record_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /live-record-info/page [get]
func Live_record_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Live_record_info](w, r, common.GetDaprClient(), "v_live_record_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 直播记录信息视图
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
// @Param stream_url_suffix query string false "stream_url_suffix"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param status query string false "status"
// @Param patient_name query string false "patient_name"
// @Param patient_ward_name query string false "patient_ward_name"
// @Param patient_bed_no query string false "patient_bed_no"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Live_record_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /live-record-info [get]
func Live_record_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Live_record_info](w, r, common.GetDaprClient(), "v_live_record_info", "id")
}
