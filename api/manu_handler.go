package api

import (
	"net/http"
	"time"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitManualHandler(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/manual/init-visit-schedule", ManualInitVisitScheduleHandler)
	r.Post(common.BASE_CONTEXT+"/manual/delete-visit-schedule", DeleteVisitScheduleHandler)
}

// @Summary 手动初始化排班
// @Description 手动初始化排班
// @Tags Manual
// @Accept json
// @Produce json
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /manual/init-visit-schedule [post]
func ManualInitVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {
	err := service.ManualInitVisitSchedule()
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK)
}

// @Summary 手动删除排班
// @Description 手动删除排班
// @Tags Manual
// @Accept json
// @Produce json
// @Param start_day query string true "删除开始时间，2006-01-02"
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /manual/delete-visit-schedule [post]
func DeleteVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {
	startDayStr := r.URL.Query().Get("start_day")
	startDay, err := time.Parse(startDayStr, "2006-01-02")
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	err = service.DeleteVisitSchedule(startDay)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK)
}
