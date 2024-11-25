package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitSystem_configRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/system-config/page", System_configPageListHandler)
	r.Get(common.BASE_CONTEXT+"/system-config", System_configListHandler)

	r.Post(common.BASE_CONTEXT+"/system-config", UpsertSystem_configHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags System_config
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param config_name query string false "config_name"
// @Param config_value query string false "config_value"
// @Param config_type query string false "config_type"
// @Param config_unit query string false "config_unit"
// @Param config_desc query string false "config_desc"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.System_config}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /system-config/page [get]
func System_configPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.System_config](w, r, common.GetDaprClient(), "o_system_config", "id")

}

// @Summary query objects
// @Description query objects
// @Tags System_config
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param config_name query string false "config_name"
// @Param config_value query string false "config_value"
// @Param config_type query string false "config_type"
// @Param config_unit query string false "config_unit"
// @Param config_desc query string false "config_desc"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.System_config} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /system-config [get]
func System_configListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.System_config](w, r, common.GetDaprClient(), "o_system_config", "id")
}

// @Summary save
// @Description save
// @Tags System_config
// @Accept       json
// @Param item body model.System_config true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.System_config} "object"
// @Failure 500 {object} common.Response ""
// @Router /system-config [post]
func UpsertSystem_configHandler(w http.ResponseWriter, r *http.Request) {
	var val model.System_config
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("System_config")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.System_config)
	}

	err = common.DbUpsert[model.System_config](r.Context(), common.GetDaprClient(), val, model.System_configTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
