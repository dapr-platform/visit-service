package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitWardRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ward/page", WardPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ward", WardListHandler)

	r.Post(common.BASE_CONTEXT+"/ward", UpsertWardHandler)

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

	err = common.DbUpsert[model.Ward](r.Context(), common.GetDaprClient(), val, model.WardTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
