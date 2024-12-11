package api

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitVisit_scheduleExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/visit-schedule/set-all-status", SetVisitScheduleStatusHandler)
	r.Post(common.BASE_CONTEXT+"/visit-schedule/init-visit-schedule", ManualInitVisitScheduleHandler)
	r.Post(common.BASE_CONTEXT+"/visit-schedule/delete-visit-schedule", DeleteVisitScheduleHandler)
}

// @Summary 手动初始化排班
// @Description 手动初始化排班
// @Tags 探视排班
// @Accept json
// @Produce json
// @Param force_update query bool false "是否强制更新, true: 强制更新, false: 不强制更新"
// @Param start_date query string false "开始日期, 2024-01-01"
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /visit-schedule/init-visit-schedule [post]
func ManualInitVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {
	forceUpdate := r.FormValue("force_update") == "true"
	startDateStr := r.FormValue("start_date")
	startDate := time.Time{}
	var err error
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
			return
		}
	}
	go func() {
		err := service.ManualInitVisitSchedule(forceUpdate, startDate)
		if err != nil {
			common.Logger.Error("init visit schedule error", err)
		}
	}()
	common.HttpResult(w, common.OK.AppendMsg("后台运行，请查看日志"))
}

// @Summary 手动删除排班
// @Description 手动删除排班, 删除开始时间到第二天开始时间之间的排班
// @Tags 探视排班
// @Accept json
// @Produce json
// @Param start_day query string true "删除开始时间，2006-01-02"
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /visit-schedule/delete-visit-schedule [post]
func DeleteVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {
	startDayStr := r.URL.Query().Get("start_day")
	startDay, err := time.Parse(startDayStr, "2006-01-02")
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	go func() {
		err = service.DeleteVisitSchedule(startDay)
		if err != nil {
			common.Logger.Error("delete visit schedule error", err)
		}
	}()
	common.HttpResult(w, common.OK.AppendMsg("后台运行，请查看日志"))
}

// 设置全部排班状态
// @Summary 设置全部排班状态
// @Description 设置全部排班状态
// @Tags 探视排班
// @Param status query int true "状态"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Router /visit-schedule/set-all-status [post]
func SetVisitScheduleStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := r.FormValue("status")
	if status == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("status is required"))
		return
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("status is invalid"))
		return
	}
	err = service.SetAllVisitScheduleStatus(context.Background(), statusInt)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, nil)
}
