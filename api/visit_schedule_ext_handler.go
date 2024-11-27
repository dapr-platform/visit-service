package api

import (
	"context"
	"net/http"
	"strconv"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitVisit_scheduleExtRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/visit-schedule/set-all-status", SetVisitScheduleStatusHandler)

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
