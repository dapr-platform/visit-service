package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitHead_displayRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/head-display/page", Head_displayPageListHandler)
	r.Get(common.BASE_CONTEXT+"/head-display", Head_displayListHandler)

	r.Post(common.BASE_CONTEXT+"/head-display", UpsertHead_displayHandler)

	r.Delete(common.BASE_CONTEXT+"/head-display/{id}", DeleteHead_displayHandler)

	r.Post(common.BASE_CONTEXT+"/head-display/batch-delete", batchDeleteHead_displayHandler)

	r.Post(common.BASE_CONTEXT+"/head-display/batch-upsert", batchUpsertHead_displayHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 头显
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /head-display/batch-upsert [post]
func batchUpsertHead_displayHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.Head_displayTableInfo.Name, model.Head_display_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 头显
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param model query string false "model"
// @Param ward_id query string false "ward_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Head_display}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /head-display/page [get]
func Head_displayPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Head_display](w, r, common.GetDaprClient(), "o_head_display", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 头显
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param device_name query string false "device_name"
// @Param device_no query string false "device_no"
// @Param model query string false "model"
// @Param ward_id query string false "ward_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Head_display} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /head-display [get]
func Head_displayListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Head_display](w, r, common.GetDaprClient(), "o_head_display", "id")
}

// @Summary save
// @Description save
// @Tags 头显
// @Accept       json
// @Param item body model.Head_display true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Head_display} "object"
// @Failure 500 {object} common.Response ""
// @Router /head-display [post]
func UpsertHead_displayHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Head_display
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Head_display")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Head_display)
	}

	err = common.DbUpsert[model.Head_display](r.Context(), common.GetDaprClient(), val, model.Head_displayTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 头显
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Head_display} "object"
// @Failure 500 {object} common.Response ""
// @Router /head-display/{id} [delete]
func DeleteHead_displayHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Head_display")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_head_display", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 头显
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /head-display/batch-delete [post]
func batchDeleteHead_displayHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Head_display")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_head_display", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
