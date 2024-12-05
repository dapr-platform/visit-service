package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"time"
)

var _ = time.Now()

func InitVisit_record_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/visit-record-info/page", Visit_record_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/visit-record-info", Visit_record_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视记录信息视图
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
// @Param send_sms_status query string false "send_sms_status"
// @Param send_prompt_sms_status query string false "send_prompt_sms_status"
// @Param patient_name query string false "patient_name"
// @Param patient_ward_name query string false "patient_ward_name"
// @Param patient_bed_no query string false "patient_bed_no"
// @Param stream_id query string false "stream_id"
// @Param stream_url_suffix query string false "stream_url_suffix"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Visit_record_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-info/page [get]
func Visit_record_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Visit_record_info](w, r, common.GetDaprClient(), "v_visit_record_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视记录信息视图
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
// @Param send_sms_status query string false "send_sms_status"
// @Param send_prompt_sms_status query string false "send_prompt_sms_status"
// @Param patient_name query string false "patient_name"
// @Param patient_ward_name query string false "patient_ward_name"
// @Param patient_bed_no query string false "patient_bed_no"
// @Param stream_id query string false "stream_id"
// @Param stream_url_suffix query string false "stream_url_suffix"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Visit_record_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-info [get]
func Visit_record_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Visit_record_info](w, r, common.GetDaprClient(), "v_visit_record_info", "id")
}
