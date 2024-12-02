package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitLive_recordRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/live-record/page", Live_recordPageListHandler)
	r.Get(common.BASE_CONTEXT+"/live-record", Live_recordListHandler)

	r.Post(common.BASE_CONTEXT+"/live-record", UpsertLive_recordHandler)

	r.Delete(common.BASE_CONTEXT+"/live-record/{id}", DeleteLive_recordHandler)

	r.Post(common.BASE_CONTEXT+"/live-record/batch-delete", batchDeleteLive_recordHandler)

	r.Post(common.BASE_CONTEXT+"/live-record/batch-upsert", batchUpsertLive_recordHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 直播记录
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /live-record/batch-upsert [post]
func batchUpsertLive_recordHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Live_record
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Live_record")
	if exists {
		for _, v := range entities {
			_, err1 := beforeHook(r, v)
			if err1 != nil {
				common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
				return
			}
		}

	}
	for _, v := range entities {
		if v.ID == "" {
			v.ID = common.NanoId()
		}
	}

	err = common.DbBatchUpsert[model.Live_record](r.Context(), common.GetDaprClient(), entities, model.Live_recordTableInfo.Name, model.Live_record_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
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
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
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
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
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
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	err = common.DbUpsert[model.Live_record](r.Context(), common.GetDaprClient(), val, model.Live_recordTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 直播记录
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Live_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /live-record/{id} [delete]
func DeleteLive_recordHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Live_record")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_live_record", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 直播记录
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /live-record/batch-delete [post]
func batchDeleteLive_recordHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Live_record")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_live_record", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
