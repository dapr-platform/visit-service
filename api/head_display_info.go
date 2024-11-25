package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitHead_display_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/head-display-info/page", Head_display_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/head-display-info", Head_display_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Head_display_info
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param model query string false "model"
// @Param ward_id query string false "ward_id"
// @Param ward_name query string false "ward_name"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Head_display_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /head-display-info/page [get]
func Head_display_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Head_display_info](w, r, common.GetDaprClient(), "v_head_display_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Head_display_info
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param model query string false "model"
// @Param ward_id query string false "ward_id"
// @Param ward_name query string false "ward_name"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Head_display_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /head-display-info [get]
func Head_display_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Head_display_info](w, r, common.GetDaprClient(), "v_head_display_info", "id")
}
