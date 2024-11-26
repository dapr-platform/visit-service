package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitSchedule_cameraRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/schedule-camera/page", Schedule_cameraPageListHandler)
	r.Get(common.BASE_CONTEXT+"/schedule-camera", Schedule_cameraListHandler)

	r.Post(common.BASE_CONTEXT+"/schedule-camera", UpsertSchedule_cameraHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视排班摄像头
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param schedule_id query string false "schedule_id"
// @Param camera_id query string false "camera_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Schedule_camera}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /schedule-camera/page [get]
func Schedule_cameraPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Schedule_camera](w, r, common.GetDaprClient(), "o_schedule_camera", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视排班摄像头
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param schedule_id query string false "schedule_id"
// @Param camera_id query string false "camera_id"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Schedule_camera} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /schedule-camera [get]
func Schedule_cameraListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Schedule_camera](w, r, common.GetDaprClient(), "o_schedule_camera", "id")
}

// @Summary save
// @Description save
// @Tags 探视排班摄像头
// @Accept       json
// @Param item body model.Schedule_camera true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Schedule_camera} "object"
// @Failure 500 {object} common.Response ""
// @Router /schedule-camera [post]
func UpsertSchedule_cameraHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Schedule_camera
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Schedule_camera")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Schedule_camera)
	}

	err = common.DbUpsert[model.Schedule_camera](r.Context(), common.GetDaprClient(), val, model.Schedule_cameraTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
