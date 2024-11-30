package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitFamily_member_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/family-member-info/page", Family_member_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/family-member-info", Family_member_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 家属信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param zh_name query string false "zh_name"
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
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Family_member_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /family-member-info/page [get]
func Family_member_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Family_member_info](w, r, common.GetDaprClient(), "v_family_member_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 家属信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param tenant_id query string false "tenant_id"
// @Param mobile query string false "mobile"
// @Param email query string false "email"
// @Param identity query string false "identity"
// @Param name query string false "name"
// @Param zh_name query string false "zh_name"
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
// @Success 200 {object} common.Response{data=[]model.Family_member_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /family-member-info [get]
func Family_member_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Family_member_info](w, r, common.GetDaprClient(), "v_family_member_info", "id")
}
