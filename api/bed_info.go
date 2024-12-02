package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"time"
)

var _ = time.Now()

func InitBed_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/bed-info/page", Bed_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/bed-info", Bed_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病床信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param ward_id query string false "ward_id"
// @Param bed_no query string false "bed_no"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param ward_name query string false "ward_name"
// @Param camera_name query string false "camera_name"
// @Param vr_camera_name query string false "vr_camera_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Bed_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /bed-info/page [get]
func Bed_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Bed_info](w, r, common.GetDaprClient(), "v_bed_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病床信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param ward_id query string false "ward_id"
// @Param bed_no query string false "bed_no"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param type query string false "type"
// @Param status query string false "status"
// @Param ward_name query string false "ward_name"
// @Param camera_name query string false "camera_name"
// @Param vr_camera_name query string false "vr_camera_name"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Bed_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /bed-info [get]
func Bed_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Bed_info](w, r, common.GetDaprClient(), "v_bed_info", "id")
}
