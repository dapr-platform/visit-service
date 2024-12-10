package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"time"
)

var _ = time.Now()

func InitPatient_relative_infoRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/patient-relative-info/page", Patient_relative_infoPageListHandler)
	r.Get(common.BASE_CONTEXT+"/patient-relative-info", Patient_relative_infoListHandler)

}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 病患-家属关联信息视图
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param create_time query string false "create_time"
// @Param check_status query string false "check_status"
// @Param patient_name query string false "patient_name"
// @Param hospital_no query string false "hospital_no"
// @Param patient_status query string false "patient_status"
// @Param ward_name query string false "ward_name"
// @Param bed_no query string false "bed_no"
// @Param relative_name query string false "relative_name"
// @Param relative_mobile query string false "relative_mobile"
// @Param relative_id_card query string false "relative_id_card"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Patient_relative_info}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative-info/page [get]
func Patient_relative_infoPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Patient_relative_info](w, r, common.GetDaprClient(), "v_patient_relative_info", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 病患-家属关联信息视图
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param patient_id query string false "patient_id"
// @Param relative_id query string false "relative_id"
// @Param relationship query string false "relationship"
// @Param status query string false "status"
// @Param create_time query string false "create_time"
// @Param check_status query string false "check_status"
// @Param patient_name query string false "patient_name"
// @Param hospital_no query string false "hospital_no"
// @Param patient_status query string false "patient_status"
// @Param ward_name query string false "ward_name"
// @Param bed_no query string false "bed_no"
// @Param relative_name query string false "relative_name"
// @Param relative_mobile query string false "relative_mobile"
// @Param relative_id_card query string false "relative_id_card"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Patient_relative_info} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /patient-relative-info [get]
func Patient_relative_infoListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Patient_relative_info](w, r, common.GetDaprClient(), "v_patient_relative_info", "id")
}
