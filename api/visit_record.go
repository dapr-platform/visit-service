package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"
)

func InitVisit_recordRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/visit-record/page", Visit_recordPageListHandler)
	r.Get(common.BASE_CONTEXT+"/visit-record", Visit_recordListHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record", UpsertVisit_recordHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视登记
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param visit_start_time query string false "visit_start_time"
// @Param visit_end_time query string false "visit_end_time"
// @Param visitor_name query string false "visitor_name"
// @Param visitor_phone query string false "visitor_phone"
// @Param visitor_id_card query string false "visitor_id_card"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Visit_record}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record/page [get]
func Visit_recordPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Visit_record](w, r, common.GetDaprClient(), "o_visit_record", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视登记
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param visit_start_time query string false "visit_start_time"
// @Param visit_end_time query string false "visit_end_time"
// @Param visitor_name query string false "visitor_name"
// @Param visitor_phone query string false "visitor_phone"
// @Param visitor_id_card query string false "visitor_id_card"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param remark query string false "remark"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Visit_record} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record [get]
func Visit_recordListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Visit_record](w, r, common.GetDaprClient(), "o_visit_record", "id")
}

// @Summary save
// @Description save
// @Tags 探视登记
// @Accept       json
// @Param item body model.Visit_record true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-record [post]
func UpsertVisit_recordHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Visit_record
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	beforeHook, exists := common.GetUpsertBeforeHook("Visit_record")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Visit_record)
	}

	err = common.DbUpsert[model.Visit_record](r.Context(), common.GetDaprClient(), val, model.Visit_recordTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}
