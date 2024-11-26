package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitCameraRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/camera/page", CameraPageListHandler)
	r.Get(common.BASE_CONTEXT+"/camera", CameraListHandler)

	r.Post(common.BASE_CONTEXT+"/camera", UpsertCameraHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags Camera
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param location_type query string false "location_type"
// @Param ward_id query string false "ward_id"
// @Param bed_id query string false "bed_id"
// @Param device_type query string false "device_type"
// @Param manufacturer query string false "manufacturer"
// @Param main_stream_url query string false "main_stream_url"
// @Param sub_stream_url query string false "sub_stream_url"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Camera}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /camera/page [get]
func CameraPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Camera](w, r, common.GetDaprClient(), "o_camera", "id")

}

// @Summary query objects
// @Description query objects
// @Tags Camera
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param location_type query string false "location_type"
// @Param ward_id query string false "ward_id"
// @Param bed_id query string false "bed_id"
// @Param device_type query string false "device_type"
// @Param manufacturer query string false "manufacturer"
// @Param main_stream_url query string false "main_stream_url"
// @Param sub_stream_url query string false "sub_stream_url"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Camera} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /camera [get]
func CameraListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Camera](w, r, common.GetDaprClient(), "o_camera", "id")
}

// @Summary save
// @Description save
// @Tags Camera
// @Accept       json
// @Param item body model.Camera true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Camera} "object"
// @Failure 500 {object} common.Response ""
// @Router /camera [post]
func UpsertCameraHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Camera
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Camera")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Camera)
	}

	err = common.DbUpsert[model.Camera](r.Context(), common.GetDaprClient(), val, model.CameraTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
