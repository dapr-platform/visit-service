package service

import (
	"context"
	"time"
	"visit-service/model"

	"github.com/dapr-platform/common"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
)

func init() {
	// 创建定时任务，每天凌晨1点执行清理历史数据
	historyCron := cron.New(cron.WithSeconds())
	_, err := historyCron.AddFunc("0 0 1 * * ?", cleanHistoryData)
	if err != nil {
		common.Logger.Errorf("Failed to add history clean cron job: %v", err)
		return
	}
	historyCron.Start()
}

func cleanHistoryData() {
	// 获取历史数据保留天数配置
	config, err := GetConfig(CONFIG_HISTORY_DATA_RETENTION_DAYS)
	if err != nil {
		common.Logger.Errorf("Failed to get history data retention days config: %v", err)
		return
	}

	// 计算截止时间
	retentionDays := cast.ToInt(config.ConfigValue)
	if retentionDays <= 0 {
		return
	}
	cutoffTime := time.Now().AddDate(0, 0, -retentionDays)
	cutoffTimeStr := cutoffTime.Format("2006-01-02T15:04:05")

	ctx := context.Background()

	// 删除历史访视记录
	err = common.DbDeleteByOps(ctx, common.GetDaprClient(),
		model.Visit_recordTableInfo.Name,
		[]string{model.Visit_record_FIELD_NAME_visit_end_time},
		[]string{"<"},
		[]any{cutoffTimeStr})
	if err != nil {
		common.Logger.Errorf("Failed to clean visit records: %v", err)
	}

	// 删除历史直播记录
	err = common.DbDeleteByOps(ctx, common.GetDaprClient(),
		model.Live_recordTableInfo.Name,
		[]string{model.Live_record_FIELD_NAME_end_time},
		[]string{"<"},
		[]any{cutoffTimeStr})
	if err != nil {
		common.Logger.Errorf("Failed to clean live records: %v", err)
	}

	// 删除历史排班记录
	err = common.DbDeleteByOps(ctx, common.GetDaprClient(),
		model.Visit_scheduleTableInfo.Name,
		[]string{model.Visit_schedule_FIELD_NAME_start_time},
		[]string{"<"},
		[]any{cutoffTimeStr})
	if err != nil {
		common.Logger.Errorf("Failed to clean visit schedules: %v", err)
	}
}
