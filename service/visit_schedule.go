package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"visit-service/model"

	"github.com/dapr-platform/common"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
)

var scheduleCron *cron.Cron

func init() {
	// 创建一个新的定时任务
	scheduleCron = cron.New(cron.WithSeconds())
	// 每天凌晨1点执行
	_, err := scheduleCron.AddFunc("0 0 1 * * ?", func() {
		if err := initVisitScheduleDaily(false); err != nil {
			common.Logger.Errorf("Failed to init visit schedule: %v", err)
		}
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to add cron job: %v", err))
	}
	scheduleCron.Start()
	common.RegisterUpsertBeforeHook("Visit_schedule", UpsertVisitSchedule)
}

func UpsertVisitSchedule(r *http.Request, in any) (out any, err error) {
	schedule := in.(model.Visit_schedule)
	if time.Time(schedule.StartTime).IsZero() {
		return nil, fmt.Errorf("start time is zero")
	}
	intervalConfig, err := GetConfig(CONFIG_SCHEDULE_INTERVAL)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_interval config: %v", err)
		return
	}
	schedule.EndTime = common.LocalTime(time.Time(schedule.StartTime).Add(time.Duration(cast.ToInt(intervalConfig.ConfigValue)) * time.Minute))
	if schedule.ID == "" {
		schedule.ID = time.Time(schedule.StartTime).Format("20060102150405")
	}
	return schedule, nil
}

func FindVisitScheduleByStartTime(ctx context.Context, startTime common.LocalTime) (*model.Visit_schedule, error) {
	qstr := model.Visit_schedule_FIELD_NAME_start_time + "=" + startTime.DbString()
	schedule, err := common.DbGetOne[model.Visit_schedule](ctx, common.GetDaprClient(), model.Visit_scheduleTableInfo.Name, qstr)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}
func IncreaseVisitScheduleVisitors(ctx context.Context, schedule *model.Visit_schedule) error {
	schedule.ScheduleVisitors++
	err := common.DbUpsert[model.Visit_schedule](ctx, common.GetDaprClient(), *schedule, model.Visit_scheduleTableInfo.Name, model.Visit_schedule_FIELD_NAME_id)
	if err != nil {
		return err
	}
	return nil
}
func DecreaseVisitScheduleVisitors(ctx context.Context, schedule *model.Visit_schedule) error {
	schedule.ScheduleVisitors--
	err := common.DbUpsert[model.Visit_schedule](ctx, common.GetDaprClient(), *schedule, model.Visit_scheduleTableInfo.Name, model.Visit_schedule_FIELD_NAME_id)
	if err != nil {
		return err
	}
	return nil
}

// ManuAddVisitSchedule 手动添加排班
func ManuAddVisitSchedule(startTime time.Time, endTime time.Time, totalVisitors int) error {
	// 检查时间段是否已存在
	exists, err := checkTimeSlotExists(startTime)
	if err != nil {
		return fmt.Errorf("failed to check time slot: %v", err)
	}
	if exists {
		return fmt.Errorf("time slot already exists")
	}

	// 创建排班记录
	schedule := model.Visit_schedule{
		ID:               startTime.Format("20060102150405"),
		StartTime:        common.LocalTime(startTime),
		EndTime:          common.LocalTime(endTime),
		TotalVisitors:    int32(totalVisitors),
		ScheduleVisitors: int32(0),
		Status:           1,
	}

	// 插入数据库
	_, err = common.DbInsert(
		context.Background(),
		common.GetDaprClient(),
		&schedule,
		model.Visit_scheduleTableInfo.Name,
	)
	if err != nil {
		return fmt.Errorf("failed to insert schedule: %v", err)
	}

	return nil
}

// ManualInitVisitSchedule 用于调试的初始化函数
func ManualInitVisitSchedule(forceUpdate bool) error {
	return initVisitScheduleDaily(forceUpdate)
}

func DeleteVisitSchedule(startDay time.Time) error {
	nextDay := startDay.AddDate(0, 0, 1)
	return common.DbDeleteByOps(
		context.Background(),
		common.GetDaprClient(),
		model.Visit_scheduleTableInfo.Name,
		[]string{model.Visit_schedule_FIELD_NAME_start_time, model.Visit_schedule_FIELD_NAME_end_time},
		[]string{">=", "<"},
		[]any{startDay.Format("2006-01-02T15:04:05"), nextDay.Format("2006-01-02T15:04:05")},
	)
}
func SetAllVisitScheduleStatus(ctx context.Context, status int) error {
	params := map[string]any{
		model.Visit_schedule_FIELD_NAME_status: status,
	}
	return common.DbUpdateByOps(
		ctx,
		common.GetDaprClient(),
		model.Visit_scheduleTableInfo.Name,
		params,
		[]string{},
		[]string{},
		[]any{},
	)
}

// checkTimeSlotExists 检查时间段是否已存在
func checkTimeSlotExists(startTime time.Time) (bool, error) {
	// 查询与给定时间段有重叠的排班
	qstr := model.Visit_schedule_FIELD_NAME_id + "=" + startTime.Format("20060102150405")

	schedules, err := common.DbQuery[model.Visit_schedule](
		context.Background(),
		common.GetDaprClient(),
		model.Visit_scheduleTableInfo.Name,
		qstr,
	)
	if err != nil {
		return false, err
	}

	return len(schedules) > 0, nil
}

func initVisitScheduleDaily(forceUpdate bool) error {
	// 获取配置，如果出错则记录日志并返回
	beginHourConfig, err := GetConfig(CONFIG_SCHEDULE_BEGIN_HOUR)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_begin_hour config: %v", err)
		return err
	}
	endHourConfig, err := GetConfig(CONFIG_SCHEDULE_END_HOUR)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_end_hour config: %v", err)
		return err
	}

	autoAvailableBeginHourConfig, err := GetConfig(CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_BEGIN_HOUR)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_state_auto_available_begin_hour config: %v", err)
		return err
	}
	autoAvailableEndHourConfig, err := GetConfig(CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_END_HOUR)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_state_auto_available_end_hour config: %v", err)
		return err
	}
	intervalConfig, err := GetConfig(CONFIG_SCHEDULE_INTERVAL)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_interval config: %v", err)
		return err
	}
	timeSpanConfig, err := GetConfig(CONFIG_SCHEDULE_TIME_SPAN)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_time_span config: %v", err)
		return err
	}
	generateDaysConfig, err := GetConfig(CONFIG_SCHEDULE_GENERATE_DAYS)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_generate_days config: %v", err)
		return err
	}
	visitorsConfig, err := GetConfig(CONFIG_SCHEDULE_MAX_VISITORS)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_max_visitors config: %v", err)
		return err
	}
	autoAvailableVisitorsConfig, err := GetConfig(CONFIG_SCHEDULE_AUTO_AVAILABLE_MAX_VISITORS)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_auto_available_max_visitors config: %v", err)
		return err
	}

	startHour := cast.ToInt(beginHourConfig.ConfigValue)
	endHour := cast.ToInt(endHourConfig.ConfigValue)
	interval := cast.ToInt(intervalConfig.ConfigValue)
	timeSpan := cast.ToInt(timeSpanConfig.ConfigValue)
	generateDays := cast.ToInt(generateDaysConfig.ConfigValue)
	maxVisitors := cast.ToInt(visitorsConfig.ConfigValue)
	common.Logger.Infof("startHour: %v, endHour: %v, interval: %v, timeSpan: %v, generateDays: %v, maxVisitors: %v", startHour, endHour, interval, timeSpan, generateDays, maxVisitors)
	// 从明天开始生成
	startDate := time.Now().AddDate(0, 0, 1)
	common.Logger.Infof("startDate: %v", startDate)
	// 生成未来 generateDays 天的排班
	for day := 0; day < generateDays; day++ {
		currentDate := startDate.AddDate(0, 0, day)
		common.Logger.Infof("currentDate: %v", currentDate)

		if forceUpdate {
			common.Logger.Infof("forceUpdate: %v, delete visit schedule: %v", forceUpdate, currentDate)
			DeleteVisitSchedule(currentDate)
		}

		startTime := time.Date(
			currentDate.Year(),
			currentDate.Month(),
			currentDate.Day(),
			startHour,
			0,
			0,
			0,
			time.Local,
		)

		// 生成当天的时间段
		for minute := 0; minute < 1440; minute += (interval + timeSpan) {
			startTime = startTime.Add(time.Duration(minute) * time.Minute)
			endTime := startTime.Add(time.Duration(timeSpan) * time.Minute)

			// 检查时间段是否已存在
			exists, err := checkTimeSlotExists(startTime)
			if err != nil {
				common.Logger.Errorf("Error checking time slot: %v", err)
				continue
			}
			if exists && !forceUpdate {
				continue
			}
			status := 0
			totalVisitors := maxVisitors
			autoAvailableVisitors := cast.ToInt(autoAvailableVisitorsConfig.ConfigValue)
			if startTime.Hour() >= cast.ToInt(autoAvailableBeginHourConfig.ConfigValue) && startTime.Hour() <= cast.ToInt(autoAvailableEndHourConfig.ConfigValue) {
				status = 1
				totalVisitors = autoAvailableVisitors
			}

			// 创建排班记录
			schedule := model.Visit_schedule{
				ID:               startTime.Format("20060102150405"),
				StartTime:        common.LocalTime(startTime),
				EndTime:          common.LocalTime(endTime),
				TotalVisitors:    int32(totalVisitors),
				ScheduleVisitors: int32(0),
				Status:           int32(status),
			}

			// 插入数据库
			_, err = common.DbInsert(
				context.Background(),
				common.GetDaprClient(),
				&schedule,
				model.Visit_scheduleTableInfo.Name,
			)
			if err != nil {
				common.Logger.Errorf("Error inserting schedule: %v", err)
				continue
			}
			common.Logger.Infof("Inserted schedule: %v", schedule)
		}
	}

	return nil
}
