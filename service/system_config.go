package service

import (
	"context"
	"errors"
	"net/http"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

/*
默认系统配置, 如果数据库中没有该配置, 则插入默认配置,包括
.视频保留天数
.家属病人关系绑定自动审核开关(1:自动审核，0:手动审核)
.最晚预约探视周期
.最早预约探视周期（天）。0为今天，1为明天，2为后天，以此类推
.系统部署时间（影响：数据统计初始时间）
.护士探视通知时间,单位秒
.家属探视通知时间,单位秒
.探视登记预约是否自动审核(1:自动审核，0:手动审核)
.排班间隔时间,单位秒
*/
var videoRetentionDays = model.System_config{
	ConfigName:  "video_retention_days",
	ConfigValue: "30",
	ConfigType:  "number",
	ConfigUnit:  "天",
	ConfigDesc:  "视频保留天数",
	Status:      1,
}
var relativePatientAutoAudit = model.System_config{
	ConfigName:  "relative_patient_auto_audit",
	ConfigValue: "1",
	ConfigType:  "number",
	ConfigDesc:  "家属病人关系绑定自动审核开关(1:自动审核，0:手动审核)",
	Status:      1,
}
var latestVisitPeriod = model.System_config{
	ConfigName:  "latest_visit_period",
	ConfigValue: "7",
	ConfigType:  "number",
	ConfigDesc:  "最晚预约探视周期",
	Status:      1,
}
var earliestVisitPeriod = model.System_config{
	ConfigName:  "earliest_visit_period",
	ConfigValue: "1",
	ConfigType:  "number",
	ConfigDesc:  "最早预约探视周期（天）。0为今天，1为明天，2为后天，以此类推",
	Status:      1,
}
var systemDeployTime = model.System_config{
	ConfigName:  "system_deploy_time",
	ConfigValue: "2024-01-01 00:00:00",
	ConfigType:  "datetime",
	ConfigDesc:  "系统部署时间（影响：数据统计初始时间）",
	Status:      1,
}
var nurseVisitNotifyTime = model.System_config{
	ConfigName:  "nurse_visit_notify_time",
	ConfigValue: "300",
	ConfigType:  "number",
	ConfigDesc:  "护士探视通知时间",
	Status:      1,
}
var relativeVisitNotifyTime = model.System_config{
	ConfigName:  "relative_visit_notify_time",
	ConfigValue: "300",
	ConfigType:  "number",
	ConfigDesc:  "家属探视通知时间",
	Status:      1,
}
var visitRegisterAutoAudit = model.System_config{
	ConfigName:  "visit_register_auto_audit",
	ConfigValue: "1",
	ConfigType:  "number",
	ConfigDesc:  "探视登记预约是否自动审核(1:自动审核，0:手动审核)",
	Status:      1,
}
var scheduleInterval = model.System_config{
	ConfigName:  "schedule_interval",
	ConfigValue: "1800",
	ConfigType:  "number",
	ConfigDesc:  "排班间隔时间",
	Status:      1,
}
var defaultSystemConfigListMap = map[string]*model.System_config{
	videoRetentionDays.ConfigName:       &videoRetentionDays,
	relativePatientAutoAudit.ConfigName: &relativePatientAutoAudit,
	latestVisitPeriod.ConfigName:        &latestVisitPeriod,
	earliestVisitPeriod.ConfigName:      &earliestVisitPeriod,
	systemDeployTime.ConfigName:         &systemDeployTime,
	nurseVisitNotifyTime.ConfigName:     &nurseVisitNotifyTime,
	relativeVisitNotifyTime.ConfigName:  &relativeVisitNotifyTime,
	visitRegisterAutoAudit.ConfigName:   &visitRegisterAutoAudit,
	scheduleInterval.ConfigName:         &scheduleInterval,
}

func init() {
	common.RegisterUpsertBeforeHook("System_config", func(r *http.Request, in any) (out any, err error) {
		return in, nil
	})
	initDefaultSystemConfig()
}

func initDefaultSystemConfig() {
	configs, err := common.DbQuery[model.System_config](context.Background(), common.GetDaprClient(), model.System_configTableInfo.Name, "")
	if err != nil {
		panic(err)
	}

	// 创建一个map来加速查找
	existingConfigs := make(map[string]model.System_config)
	for _, config := range configs {
		existingConfigs[config.ConfigName] = config
	}

	// 遍历默认配置
	for name, defaultConfig := range defaultSystemConfigListMap {
		if existingConfig, exists := existingConfigs[name]; exists {
			// 更新内存中的默认配置值
			defaultConfig.ConfigValue = existingConfig.ConfigValue
		} else {
			// 创建新配置的副本
			configToInsert := *defaultConfig // 解引用创建副本
			_, err = common.DbInsert(context.Background(), common.GetDaprClient(), &configToInsert, model.System_configTableInfo.Name)
			if err != nil {
				panic(err)
			}
		}
	}
}

func UpsertSystem_config(r *http.Request, in any) (out any, err error) {
	v, ok := in.(model.System_config)
	if !ok {
		return nil, errors.New("invalid input")
	}
	
	defaultConfig, exists := defaultSystemConfigListMap[v.ConfigName]
	if !exists {
		return nil, errors.New("config name not found")
	}
	
	defaultConfig.ConfigValue = v.ConfigValue
	return v, nil
}
