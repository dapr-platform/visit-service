package service

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

// 系统配置项名称常量
const (
	CONFIG_VIDEO_RETENTION_DAYS                     = "video_retention_days"
	CONFIG_RELATIVE_PATIENT_AUTO_AUDIT              = "relative_patient_auto_audit"
	CONFIG_LATEST_VISIT_PERIOD                      = "latest_visit_period"
	CONFIG_EARLIEST_VISIT_PERIOD                    = "earliest_visit_period"
	CONFIG_SYSTEM_DEPLOY_TIME                       = "system_deploy_time"
	CONFIG_NURSE_VISIT_NOTIFY_TIME                  = "nurse_visit_notify_time"
	CONFIG_RELATIVE_VISIT_NOTIFY_TIME               = "relative_visit_notify_time"
	CONFIG_VISIT_REGISTER_AUTO_AUDIT                = "visit_register_auto_audit"
	CONFIG_SCHEDULE_INTERVAL                        = "schedule_interval"
	CONFIG_SCHEDULE_BEGIN_HOUR                      = "schedule_begin_hour"
	CONFIG_SCHEDULE_END_HOUR                        = "schedule_end_hour"
	CONFIG_SCHEDULE_TIME_SPAN                       = "schedule_time_span"
	CONFIG_SCHEDULE_GENERATE_DAYS                   = "schedule_generate_days"
	CONFIG_SCHEDULE_AUTO_AVAILABLE_MAX_VISITORS     = "schedule_auto_available_max_visitors"
	CONFIG_SCHEDULE_MAX_VISITORS                    = "schedule_max_visitors"
	CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_BEGIN_HOUR = "schedule_state_auto_available_begin_hour"
	CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_END_HOUR   = "schedule_state_auto_available_end_hour"
)

// ConfigManager 系统配置管理器
type ConfigManager struct {
	configs map[string]*model.System_config
	loaded  bool
	mu      sync.RWMutex
}

var (
	configManager = &ConfigManager{
		configs: make(map[string]*model.System_config),
	}

	// 系统配置定义
	videoRetentionDays = model.System_config{
		ConfigName:  CONFIG_VIDEO_RETENTION_DAYS,
		ConfigValue: "30",
		ConfigType:  "number",
		ConfigUnit:  "天",
		ConfigDesc:  "视频保留天数",
		Status:      1,
	}
	relativePatientAutoAudit = model.System_config{
		ConfigName:  CONFIG_RELATIVE_PATIENT_AUTO_AUDIT,
		ConfigValue: "1",
		ConfigType:  "number",
		ConfigDesc:  "家属病人关系绑定自动审核开关(1:自动审核，0:手动审核)",
		Status:      1,
	}
	latestVisitPeriod = model.System_config{
		ConfigName:  CONFIG_LATEST_VISIT_PERIOD,
		ConfigValue: "7",
		ConfigType:  "number",
		ConfigDesc:  "最晚预约探视周期",
		Status:      1,
	}
	earliestVisitPeriod = model.System_config{
		ConfigName:  CONFIG_EARLIEST_VISIT_PERIOD,
		ConfigValue: "1",
		ConfigType:  "number",
		ConfigDesc:  "最早预约探视周期（天）。0为今天，1为明天，2为后天，以此类推",
		Status:      1,
	}
	systemDeployTime = model.System_config{
		ConfigName:  CONFIG_SYSTEM_DEPLOY_TIME,
		ConfigValue: "2024-01-01 00:00:00",
		ConfigType:  "datetime",
		ConfigDesc:  "系统部署时间（影响：数据统计初始时间）",
		Status:      1,
	}
	nurseVisitNotifyTime = model.System_config{
		ConfigName:  CONFIG_NURSE_VISIT_NOTIFY_TIME,
		ConfigValue: "300",
		ConfigType:  "number",
		ConfigDesc:  "护士探视通知时间",
		Status:      1,
	}
	relativeVisitNotifyTime = model.System_config{
		ConfigName:  CONFIG_RELATIVE_VISIT_NOTIFY_TIME,
		ConfigValue: "300",
		ConfigType:  "number",
		ConfigDesc:  "家属探视通知时间",
		Status:      1,
	}
	visitRegisterAutoAudit = model.System_config{
		ConfigName:  CONFIG_VISIT_REGISTER_AUTO_AUDIT,
		ConfigValue: "1",
		ConfigType:  "number",
		ConfigDesc:  "探视登记预约是否自动审核(1:自动审核，0:手动审核)",
		Status:      1,
	}
	scheduleInterval = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_INTERVAL,
		ConfigValue: "30",
		ConfigType:  "number",
		ConfigDesc:  "排班间隔时间(分钟)",
		Status:      1,
	}
	scheduleBeginHour = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_BEGIN_HOUR,
		ConfigValue: "0",
		ConfigType:  "number",
		ConfigDesc:  "排班开始时间（小时）0-23",
		Status:      1,
	}
	scheduleEndHour = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_END_HOUR,
		ConfigValue: "23",
		ConfigType:  "number",
		ConfigDesc:  "排班结束时间（小时）0-23",
		Status:      1,
	}
	scheduleTimeSpan = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_TIME_SPAN,
		ConfigValue: "15",
		ConfigType:  "number",
		ConfigDesc:  "排班时间跨度(分钟)",
		Status:      1,
	}
	scheduleGenerateDays = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_GENERATE_DAYS,
		ConfigValue: "8",
		ConfigType:  "number",
		ConfigDesc:  "排班生成天数",
		Status:      1,
	}
	scheduleAutoAvailableMaxVisitors = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_AUTO_AVAILABLE_MAX_VISITORS,
		ConfigValue: "10",
		ConfigType:  "number",
		ConfigDesc:  "自动可预约时段最大访客数",
		Status:      1,
	}
	scheduleMaxVisitors = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_MAX_VISITORS,
		ConfigValue: "1",
		ConfigType:  "number",
		ConfigDesc:  "每个时段最大访客数",
		Status:      1,
	}
	scheduleStateAutoAvailableBeginHour = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_BEGIN_HOUR,
		ConfigValue: "15",
		ConfigType:  "number",
		ConfigDesc:  "排班状态自动可预约开始时间(小时)",
		Status:      1,
	}
	scheduleStateAutoAvailableEndHour = model.System_config{
		ConfigName:  CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_END_HOUR,
		ConfigValue: "16",
		ConfigType:  "number",
		ConfigDesc:  "排班状态自动可预约结束时间(小时)",
		Status:      1,
	}

	// 默认配置映射
	defaultSystemConfigListMap = map[string]*model.System_config{
		CONFIG_VIDEO_RETENTION_DAYS:                     &videoRetentionDays,
		CONFIG_RELATIVE_PATIENT_AUTO_AUDIT:              &relativePatientAutoAudit,
		CONFIG_LATEST_VISIT_PERIOD:                      &latestVisitPeriod,
		CONFIG_EARLIEST_VISIT_PERIOD:                    &earliestVisitPeriod,
		CONFIG_SYSTEM_DEPLOY_TIME:                       &systemDeployTime,
		CONFIG_NURSE_VISIT_NOTIFY_TIME:                  &nurseVisitNotifyTime,
		CONFIG_RELATIVE_VISIT_NOTIFY_TIME:               &relativeVisitNotifyTime,
		CONFIG_VISIT_REGISTER_AUTO_AUDIT:                &visitRegisterAutoAudit,
		CONFIG_SCHEDULE_INTERVAL:                        &scheduleInterval,
		CONFIG_SCHEDULE_BEGIN_HOUR:                      &scheduleBeginHour,
		CONFIG_SCHEDULE_END_HOUR:                        &scheduleEndHour,
		CONFIG_SCHEDULE_TIME_SPAN:                       &scheduleTimeSpan,
		CONFIG_SCHEDULE_GENERATE_DAYS:                   &scheduleGenerateDays,
		CONFIG_SCHEDULE_AUTO_AVAILABLE_MAX_VISITORS:     &scheduleAutoAvailableMaxVisitors,
		CONFIG_SCHEDULE_MAX_VISITORS:                     &scheduleMaxVisitors,
		CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_BEGIN_HOUR: &scheduleStateAutoAvailableBeginHour,
		CONFIG_SCHEDULE_STATE_AUTO_AVAILABLE_END_HOUR:   &scheduleStateAutoAvailableEndHour,
	}
)

func init() {
	// 初始化默认配置到内存
	for name, config := range defaultSystemConfigListMap {
		configManager.configs[name] = config
	}

	common.RegisterUpsertBeforeHook("System_config", UpsertSystem_config)
}

// GetConfig 获取配置，如果配置未加载会尝试加载
func GetConfig(name string) (*model.System_config, error) {
	// 先尝试读取
	configManager.mu.RLock()
	if config, exists := configManager.configs[name]; exists {
		if configManager.loaded {
			configManager.mu.RUnlock()
			return config, nil
		}
	}
	configManager.mu.RUnlock()

	// 如果未加载，尝试加载
	configManager.mu.Lock()
	defer configManager.mu.Unlock()

	// 双重检查
	if configManager.loaded {
		return configManager.configs[name], nil
	}

	// 尝试从数据库加载
	if err := loadConfigs(); err != nil {
		return nil, err
	}

	config, exists := configManager.configs[name]
	if !exists {
		return nil, errors.New("config not found")
	}
	return config, nil
}

// loadConfigs 从数据库加载配置
func loadConfigs() error {
	configs, err := common.DbQuery[model.System_config](context.Background(), common.GetDaprClient(), model.System_configTableInfo.Name, "")
	if err != nil {
		return err
	}

	// 更新内存中的配置
	for _, config := range configs {
		if defaultConfig, exists := configManager.configs[config.ConfigName]; exists {
			defaultConfig.ConfigValue = config.ConfigValue
		}
	}

	// 将默认配置写入数据库（如果不存在）
	for name, config := range configManager.configs {
		exists := false
		for _, dbConfig := range configs {
			if dbConfig.ConfigName == name {
				exists = true
				break
			}
		}
		if !exists {
			configToInsert := *config
			configToInsert.ID = common.NanoId()
			_, err = common.DbInsert(context.Background(), common.GetDaprClient(), &configToInsert, model.System_configTableInfo.Name)
			if err != nil {
				return err
			}
		}
	}

	configManager.loaded = true
	return nil
}

// InitSystemConfig 初始化系统配置，应在服务启动后调用
func InitSystemConfig() error {
	_, err := GetConfig(CONFIG_VIDEO_RETENTION_DAYS) // 尝试获取任意配置来触发加载
	return err
}

func UpsertSystem_config(r *http.Request, in any) (out any, err error) {
	v, ok := in.(model.System_config)
	if !ok {
		return nil, errors.New("invalid input")
	}

	configManager.mu.Lock()
	defer configManager.mu.Unlock()

	config, exists := configManager.configs[v.ConfigName]
	if !exists {
		return nil, errors.New("config name not found")
	}

	config.ConfigValue = v.ConfigValue
	return v, nil
}
