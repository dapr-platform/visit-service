package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"time"
)

var _ = time.Now()

func InitCamera_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/camera-info/page", Camera_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/camera-info", Camera_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 摄像头信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param location_type query string false "location_type"
// @Param device_type query string false "device_type"
// @Param manufacturer query string false "manufacturer"
// @Param main_stream_url query string false "main_stream_url"
// @Param sub_stream_url query string false "sub_stream_url"
// @Param rel_vr_camera_id query string false "rel_vr_camera_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Camera_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /camera-info/page [get]
func Camera_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Camera_info](w, r, common.GetDaprClient(), "v_camera_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 摄像头信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param location_type query string false "location_type"
// @Param device_type query string false "device_type"
// @Param manufacturer query string false "manufacturer"
// @Param main_stream_url query string false "main_stream_url"
// @Param sub_stream_url query string false "sub_stream_url"
// @Param rel_vr_camera_id query string false "rel_vr_camera_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Camera_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /camera-info [get]
func Camera_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Camera_info](w, r, common.GetDaprClient(), "v_camera_info", "id")
}
