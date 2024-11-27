package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitCameraRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/camera/page", CameraPageListHandler)
	r.Get(common.BASE_CONTEXT+"/camera", CameraListHandler)

	r.Post(common.BASE_CONTEXT+"/camera", UpsertCameraHandler)

	r.Delete(common.BASE_CONTEXT+"/camera/{id}", DeleteCameraHandler)

	r.Post(common.BASE_CONTEXT+"/camera/batch-delete", batchDeleteCameraHandler)

	r.Post(common.BASE_CONTEXT+"/camera/batch-upsert", batchUpsertCameraHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 摄像头
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /camera/batch-upsert [post]
func batchUpsertCameraHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.CameraTableInfo.Name, model.Camera_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 摄像头
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
// @Tags 摄像头
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
// @Success 200 {object} common.Response{data=[]model.Camera} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /camera [get]
func CameraListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Camera](w, r, common.GetDaprClient(), "o_camera", "id")
}

// @Summary save
// @Description save
// @Tags 摄像头
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

// @Summary delete
// @Description delete
// @Tags 摄像头
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Camera} "object"
// @Failure 500 {object} common.Response ""
// @Router /camera/{id} [delete]
func DeleteCameraHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Camera")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_camera", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 摄像头
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /camera/batch-delete [post]
func batchDeleteCameraHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Camera")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_camera", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
