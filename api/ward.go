package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitWardRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ward/page", WardPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ward", WardListHandler)

	r.Post(common.BASE_CONTEXT+"/ward", UpsertWardHandler)

	r.Delete(common.BASE_CONTEXT+"/ward/{id}", DeleteWardHandler)

	r.Post(common.BASE_CONTEXT+"/ward/batch-delete", batchDeleteWardHandler)

	r.Post(common.BASE_CONTEXT+"/ward/batch-upsert", batchUpsertWardHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 病房
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ward/batch-upsert [post]
func batchUpsertWardHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Ward
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ward")
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

	err = common.DbBatchUpsert[model.Ward](r.Context(), common.GetDaprClient(), entities, model.WardTableInfo.Name, model.Ward_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病房
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ward}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ward/page [get]
func WardPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ward](w, r, common.GetDaprClient(), "o_ward", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病房
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param name query string false "name"
// @Param type query string false "type"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ward} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ward [get]
func WardListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ward](w, r, common.GetDaprClient(), "o_ward", "id")
}

// @Summary save
// @Description save
// @Tags 病房
// @Accept       json
// @Param item body model.Ward true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ward} "object"
// @Failure 500 {object} common.Response ""
// @Router /ward [post]
func UpsertWardHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ward
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ward")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ward)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}
	err = common.DbUpsert[model.Ward](r.Context(), common.GetDaprClient(), val, model.WardTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 病房
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ward} "object"
// @Failure 500 {object} common.Response ""
// @Router /ward/{id} [delete]
func DeleteWardHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ward")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ward", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 病房
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ward/batch-delete [post]
func batchDeleteWardHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ward")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ward", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
