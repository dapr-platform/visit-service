package api

import (
	"net/http"
	"visit-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitManualHandler(mux *chi.Mux) {
	mux.HandleFunc("/manual/init-visit-schedule", DebugInitVisitScheduleHandler)
	mux.HandleFunc("/manual/delete-visit-schedule", DeleteVisitScheduleHandler)
}

// @Summary 手动初始化排班
// @Description 手动初始化排班
// @Tags Manual
// @Accept json
// @Produce json
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /debug/init-visit-schedule [post]
func DebugInitVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {
	err := service.DebugInitVisitSchedule()
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
// @Success 200 {object} common.Response "Success response"
// @Failure 400 {object} common.Response "Invalid request parameters"
// @Failure 500 {object} common.Response "Internal server error"
// @Router /manual/delete-visit-schedule [post]
func DeleteVisitScheduleHandler(w http.ResponseWriter, r *http.Request) {

}
