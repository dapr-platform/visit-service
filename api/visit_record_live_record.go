package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"visit-service/model"

	"strings"

	"time"
)

var _ = time.Now()

func InitVisit_record_live_recordRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/visit-record-live-record/page", Visit_record_live_recordPageListHandler)
	r.Get(common.BASE_CONTEXT+"/visit-record-live-record", Visit_record_live_recordListHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record-live-record", UpsertVisit_record_live_recordHandler)

	r.Delete(common.BASE_CONTEXT+"/visit-record-live-record/{id}", DeleteVisit_record_live_recordHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record-live-record/batch-delete", batchDeleteVisit_record_live_recordHandler)

	r.Post(common.BASE_CONTEXT+"/visit-record-live-record/batch-upsert", batchUpsertVisit_record_live_recordHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 探视记录-直播记录关联
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record/batch-upsert [post]
func batchUpsertVisit_record_live_recordHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Visit_record_live_record
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Visit_record_live_record")
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

	err = common.DbBatchUpsert[model.Visit_record_live_record](r.Context(), common.GetDaprClient(), entities, model.Visit_record_live_recordTableInfo.Name, model.Visit_record_live_record_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 探视记录-直播记录关联
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param visit_record_id query string false "visit_record_id"
// @Param live_record_id query string false "live_record_id"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Visit_record_live_record}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record/page [get]
func Visit_record_live_recordPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Visit_record_live_record](w, r, common.GetDaprClient(), "o_visit_record_live_record", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 探视记录-直播记录关联
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param visit_record_id query string false "visit_record_id"
// @Param live_record_id query string false "live_record_id"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Visit_record_live_record} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record [get]
func Visit_record_live_recordListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Visit_record_live_record](w, r, common.GetDaprClient(), "o_visit_record_live_record", "id")
}

// @Summary save
// @Description save
// @Tags 探视记录-直播记录关联
// @Accept       json
// @Param item body model.Visit_record_live_record true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_record_live_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record [post]
func UpsertVisit_record_live_recordHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Visit_record_live_record
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Visit_record_live_record")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Visit_record_live_record)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}

	err = common.DbUpsert[model.Visit_record_live_record](r.Context(), common.GetDaprClient(), val, model.Visit_record_live_recordTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 探视记录-直播记录关联
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Visit_record_live_record} "object"
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record/{id} [delete]
func DeleteVisit_record_live_recordHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Visit_record_live_record")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_visit_record_live_record", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 探视记录-直播记录关联
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /visit-record-live-record/batch-delete [post]
func batchDeleteVisit_record_live_recordHandler(w http.ResponseWriter, r *http.Request) {

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
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Visit_record_live_record")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_visit_record_live_record", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
