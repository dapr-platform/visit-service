package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitUserRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/user/page", UserPageListHandler)
	r.Get(common.BASE_CONTEXT+"/user", UserListHandler)

	r.Post(common.BASE_CONTEXT+"/user", UpsertUserHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags User
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param gender query string false "gender"
// @Param address query string false "address"
// @Param password query string false "password"
// @Param type query string false "type"
// @Param org_id query string false "org_id"
// @Param id_card query string false "id_card"
// @Param work_number query string false "work_number"
// @Param avatar_url query string false "avatar_url"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param remark query string false "remark"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.User}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user/page [get]
func UserPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.User](w, r, common.GetDaprClient(), "o_user", "id")

}

// @Summary query objects
// @Description query objects
// @Tags User
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param gender query string false "gender"
// @Param address query string false "address"
// @Param password query string false "password"
// @Param type query string false "type"
// @Param org_id query string false "org_id"
// @Param id_card query string false "id_card"
// @Param work_number query string false "work_number"
// @Param avatar_url query string false "avatar_url"
// @Param create_at query string false "create_at"
// @Param update_at query string false "update_at"
// @Param remark query string false "remark"
// @Param status query string false "status"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.User} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /user [get]
func UserListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.User](w, r, common.GetDaprClient(), "o_user", "id")
}

// @Summary save
// @Description save
// @Tags User
// @Accept       json
// @Param item body model.User true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User} "object"
// @Failure 500 {object} common.Response ""
// @Router /user [post]
func UpsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var val model.User
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("User")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.User)
	}

	err = common.DbUpsert[model.User](r.Context(), common.GetDaprClient(), val, model.UserTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
