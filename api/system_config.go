package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"
)

func InitSystem_configRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/system-config/page", System_configPageListHandler)
	r.Get(common.BASE_CONTEXT+"/system-config", System_configListHandler)

	r.Post(common.BASE_CONTEXT+"/system-config", UpsertSystem_configHandler)

	r.Delete(common.BASE_CONTEXT+"/system-config/{id}", DeleteSystem_configHandler)

	r.Post(common.BASE_CONTEXT+"/system-config/batch-delete", batchDeleteSystem_configHandler)

	r.Post(common.BASE_CONTEXT+"/system-config/batch-upsert", batchUpsertSystem_configHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 系统配置
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /system-config/batch-upsert [post]
func batchUpsertSystem_configHandler(w http.ResponseWriter, r *http.Request) {

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
	for _, v := range entities {
		if v["id"] == "" {
			v["id"] = common.NanoId()
		}
	}

	err = common.DbBatchUpsert[map[string]any](r.Context(), common.GetDaprClient(), entities, model.System_configTableInfo.Name, model.System_config_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 系统配置
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
// @Tags 系统配置
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
// @Tags 系统配置
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
	if val.ID == "" {
		val.ID = common.NanoId()
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

// @Summary delete
// @Description delete
// @Tags 系统配置
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.System_config} "object"
// @Failure 500 {object} common.Response ""
// @Router /system-config/{id} [delete]
func DeleteSystem_configHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("System_config")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_system_config", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 系统配置
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /system-config/batch-delete [post]
func batchDeleteSystem_configHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("System_config")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_system_config", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
