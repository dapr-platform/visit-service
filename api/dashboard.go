package api

import (
	"net/http"
	"sort"
	"visit-service/entity"
	"visit-service/model"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitDashboardRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/dashboard/stats", DashboardStatsHandler)
	r.Get(common.BASE_CONTEXT+"/dashboard/monthly-stats", DashboardMonthlyStatsHandler)
}

// @Summary 获取探视预约总量，当前系统普通用户人数，当前系统病患人数，当前家属人数
// @Description 获取探视预约总量，当前系统普通用户人数，当前系统病患人数，当前家属人数
// @Tags 仪表盘
// @Produce json
// @Success 200 {object} common.Response{data=map[string]int64} "统计数据"
// @Failure 500 {object} common.Response ""
// @Router /dashboard/stats [get]
func DashboardStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := entity.DashboardStatsResponse{}

	// 获取探视预约总量
	visitCount, err := common.DbGetCount(r.Context(), common.GetDaprClient(), model.Visit_recordTableInfo.Name, model.Visit_record_FIELD_NAME_id, "")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	stats.VisitCount = visitCount

	// 获取普通用户数量
	userCount, err := common.DbGetCount(r.Context(), common.GetDaprClient(), model.UserTableInfo.Name, model.User_FIELD_NAME_id, "type=2")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	stats.UserCount = userCount

	// 获取病患数量
	patientCount, err := common.DbGetCount(r.Context(), common.GetDaprClient(), model.PatientTableInfo.Name, model.Patient_FIELD_NAME_id, "")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	stats.PatientCount = patientCount

	// 获取家属数量
	relativeCount, err := common.DbGetCount(r.Context(), common.GetDaprClient(), model.UserTableInfo.Name, model.User_FIELD_NAME_id, "type=3")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	stats.FamilyCount = relativeCount

	common.HttpResult(w, common.OK.WithData(stats))
}

// @Summary 获取探视预约按月数量统计
// @Description 获取探视预约按月数量统计
// @Tags 仪表盘
// @Produce json
// @Success 200 {object} common.Response{data=[]entity.VisitRecordStats} "月度统计数据"
// @Failure 500 {object} common.Response ""
// @Router /dashboard/monthly-stats [get]
func DashboardMonthlyStatsHandler(w http.ResponseWriter, r *http.Request) {

	selectSql := "to_char(visit_end_time, 'YYYY-MM') as month, COUNT(*) as count"
	fromSql := "o_visit_record"
	whereSql := ` 1==1 GROUP BY to_char(visit_end_time, 'YYYY-MM')
			ORDER BY month DESC 
			LIMIT 12`
	type Stats struct {
		Month string `json:"month"`
		Count int64  `json:"count"`
	}
	result, err := common.CustomSql[entity.VisitRecordStats](r.Context(), common.GetDaprClient(), selectSql, fromSql, whereSql)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	// Sort result by month ascending
	sort.Slice(result, func(i, j int) bool {
		return result[i].Month < result[j].Month
	})

	common.HttpResult(w, common.OK.WithData(result))
}
