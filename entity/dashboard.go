package entity

type DashboardStatsResponse struct {
	VisitCount   int64 `json:"visit_count"` // 探视预约总量
	UserCount    int64 `json:"user_count"`
	PatientCount int64 `json:"patient_count"`
	FamilyCount  int64 `json:"family_count"`
}

type DashboardMonthlyStatsResponse struct {
	Month string `json:"month"` // 月份
	Count int64  `json:"count"` // 数量
}
type VisitRecordStats struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}
