package service

import (
	"context"
	"fmt"
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
		if err := initVisitScheduleDaily(); err != nil {
			common.Logger.Errorf("Failed to init visit schedule: %v", err)
		}
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to add cron job: %v", err))
	}
	scheduleCron.Start()
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
		ID:                startTime.Format("20060102150405"),
		StartTime:         common.LocalTime(startTime),
		EndTime:           common.LocalTime(endTime),
		TotalVisitors:     int32(totalVisitors),
		RemainingVisitors: int32(totalVisitors),
		Status:            1,
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

// DebugInitVisitSchedule 用于调试的初始化函数
func DebugInitVisitSchedule() error {
	return initVisitScheduleDaily()
}

func DeleteVisitSchedule(startDay time.Time) error {
	return common.DbDeleteByOps(
		context.Background(),
		common.GetDaprClient(),
		model.Visit_scheduleTableInfo.Name,
		[]string{model.Visit_schedule_FIELD_NAME_start_time},
		[]string{">="},
		[]any{startDay.Format("2006-01-02T00:00:00")},
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

func initVisitScheduleDaily() error {
	// 获取配置，如果出错则记录日志并返回
	hourBeginConfig, err := GetConfig(CONFIG_SCHEDULE_HOUR_BEGIN)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_hour_begin config: %v", err)
		return err
	}
	hourEndConfig, err := GetConfig(CONFIG_SCHEDULE_HOUR_END)
	if err != nil {
		common.Logger.Errorf("Failed to get schedule_hour_end config: %v", err)
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

	startHour := cast.ToInt(hourBeginConfig.ConfigValue)
	endHour := cast.ToInt(hourEndConfig.ConfigValue)
	interval := cast.ToInt(intervalConfig.ConfigValue)
	timeSpan := cast.ToInt(timeSpanConfig.ConfigValue)
	generateDays := cast.ToInt(generateDaysConfig.ConfigValue)
	maxVisitors := cast.ToInt(visitorsConfig.ConfigValue)

	// 从明天开始生成
	startDate := time.Now().AddDate(0, 0, 1)

	// 生成未来 generateDays 天的排班
	for day := 0; day < generateDays; day++ {
		currentDate := startDate.AddDate(0, 0, day)

		// 生成当天的时间段
		for hour := startHour; hour < endHour; hour++ {
			for minute := 0; minute < 60; minute += (interval + timeSpan) {
				startTime := time.Date(
					currentDate.Year(),
					currentDate.Month(),
					currentDate.Day(),
					hour,
					minute,
					0,
					0,
					time.Local,
				)
				endTime := startTime.Add(time.Duration(timeSpan) * time.Minute)

				// 检查时间段是否已存在
				exists, err := checkTimeSlotExists(startTime)
				if err != nil {
					common.Logger.Errorf("Error checking time slot: %v", err)
					continue
				}
				if exists {
					continue
				}

				// 创建排班记录
				schedule := model.Visit_schedule{
					ID:                startTime.Format("20060102150405"),
					StartTime:         common.LocalTime(startTime),
					EndTime:           common.LocalTime(endTime),
					TotalVisitors:     int32(maxVisitors),
					RemainingVisitors: int32(maxVisitors),
					Status:            1,
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
			}
		}
	}

	return nil
}
