package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitBedRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/bed/page", BedPageListHandler)
	r.Get(common.BASE_CONTEXT+"/bed", BedListHandler)

	r.Post(common.BASE_CONTEXT+"/bed", UpsertBedHandler)

	r.Delete(common.BASE_CONTEXT+"/bed/{id}", DeleteBedHandler)

	r.Post(common.BASE_CONTEXT+"/bed/batch-delete", batchDeleteBedHandler)

	r.Post(common.BASE_CONTEXT+"/bed/batch-upsert", batchUpsertBedHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 病床
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /bed/batch-upsert [post]
func batchUpsertBedHandler(w http.ResponseWriter, r *http.Request) {

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

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.BedTableInfo.Name, model.Bed_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病床
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
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Bed}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /bed/page [get]
func BedPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Bed](w, r, common.GetDaprClient(), "o_bed", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病床
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param ward_id query string false "ward_id"
// @Param bed_no query string false "bed_no"
// @Param camera_id query string false "camera_id"
// @Param vr_camera_id query string false "vr_camera_id"
// @Param type query string false "type"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Bed} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /bed [get]
func BedListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Bed](w, r, common.GetDaprClient(), "o_bed", "id")
}

// @Summary save
// @Description save
// @Tags 病床
// @Accept       json
// @Param item body model.Bed true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Bed} "object"
// @Failure 500 {object} common.Response ""
// @Router /bed [post]
func UpsertBedHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Bed
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Bed")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Bed)
	}

	err = common.DbUpsert[model.Bed](r.Context(), common.GetDaprClient(), val, model.BedTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 病床
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Bed} "object"
// @Failure 500 {object} common.Response ""
// @Router /bed/{id} [delete]
func DeleteBedHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Bed")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_bed", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 病床
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /bed/batch-delete [post]
func batchDeleteBedHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Bed")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_bed", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
