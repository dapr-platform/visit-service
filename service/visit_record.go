package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
	"visit-service/config"
	"visit-service/model"
	"visit-service/sms"

	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
)

var CHECK_STATUS_UNCHECKED int32 = 0     //未审核
var CHECK_STATUS_CHECKED int32 = 1       //已审核
var CHECK_STATUS_CHECK_FAILED int32 = 2  //审核不通过
var VISIT_RECORD_STATUS_NORMAL int32 = 0 //正常
var VISIT_RECORD_STATUS_CANCEL int32 = 1 //取消
var SEND_SMS_STATUS_UNSENT int32 = 0     //未发送
var SEND_SMS_STATUS_SENT int32 = 1       //已发送

var visitRecordNewLock = sync.Mutex{}

func init() {
	// 创建一个新的定时任务
	scheduleCron = cron.New(cron.WithSeconds())
	// 每分钟执行
	_, err := scheduleCron.AddFunc("0 * * * * ?", func() {
		if err := LoopCheckVisitRecordWillAfter5Minutes(); err != nil {
			common.Logger.Errorf("Failed to loop check visit record: %v", err)
		}
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to add cron job: %v", err))
	}
	scheduleCron.Start()
	common.RegisterUpsertBeforeHook("Visit_record", UpsertVisit_record)
}

func UpsertVisit_record(r *http.Request, in any) (out any, err error) {
	record := in.(model.Visit_record)
	if record.CheckStatus == CHECK_STATUS_UNCHECKED { //未审核
		if record.Status == VISIT_RECORD_STATUS_NORMAL && record.ID == "" { //正常状态可以新建
			err = newAddVisitRecord(r.Context(), &record)
			if err != nil {
				return nil, errors.Wrap(err, "新建预约记录失败")
			}
			config, err := GetConfig(CONFIG_VISIT_REGISTER_AUTO_AUDIT)
			if err != nil {
				return nil, errors.Wrap(err, "查询系统配置失败")
			}
			if config.ConfigValue == "1" {
				record.CheckStatus = CHECK_STATUS_CHECKED
				SendVisitRecordSms(&record)
			}
		} else if record.Status == VISIT_RECORD_STATUS_CANCEL && record.CameraID != "" { //取消状态，需要释放摄像头,并减少排班已预约人数
			err = cancelVisitRecord(r.Context(), &record)
			if err != nil {
				return nil, errors.Wrap(err, "取消预约记录失败")
			}
		}
	} else if record.CheckStatus == CHECK_STATUS_CHECK_FAILED && record.CameraID != "" { //审核不通过，删除摄像头信息，并减少排班已预约人数
		err = cancelVisitRecord(r.Context(), &record)
	} else if record.CheckStatus == CHECK_STATUS_CHECKED && record.Status == VISIT_RECORD_STATUS_NORMAL { //审核通过，发送短信
		SendVisitRecordSms(&record)
	}

	return record, err
}
func cancelVisitRecord(ctx context.Context, record *model.Visit_record) error {
	schedule, err := FindVisitScheduleByStartTime(ctx, record.VisitStartTime)
	if err != nil {
		return errors.Wrap(err, "查询排班信息失败")
	}
	if schedule != nil {
		err = DecreaseVisitScheduleVisitors(ctx, schedule)
		if err != nil {
			return errors.Wrap(err, "减少排班已预约人数失败")
		}
	}
	record.CameraID = ""
	record.VrCameraID = ""
	return nil
}

func newAddVisitRecord(ctx context.Context, record *model.Visit_record) error {
	configPerDay, err := GetConfig(CONFIG_MAX_VISIT_PER_DAY)
	if err != nil {
		return errors.Wrap(err, "查询系统配置失败")
	}
	maxVisitPerDay := cast.ToInt(configPerDay.ConfigValue)
	if maxVisitPerDay > 0 {
		count, err := findVisitRecordCountByRelativeIDAndStartDay(ctx, record.RelativeID, record.VisitStartTime)
		if err != nil {
			return errors.Wrap(err, "查询已存在预约记录失败")
		}
		if count >= maxVisitPerDay {
			return errors.New("已达到预约次数上限")
		}
	}
	existRecord, err := findVisitRecordByRelativeIDAndStartTime(ctx, record.RelativeID, record.VisitStartTime)
	if err != nil {
		return errors.Wrap(err, "查询已存在预约记录失败")
	}
	if existRecord != nil {
		return errors.New("当前时段已存在预约记录")
	}
	schedule, err := FindVisitScheduleByStartTime(ctx, record.VisitStartTime)
	if err != nil {
		return errors.Wrap(err, "查询排班信息失败")
	}
	if schedule == nil {
		return errors.New("排班信息不存在")
	}
	if schedule.ScheduleVisitors >= schedule.TotalVisitors {
		return errors.New("排班已满")
	}

	patient, err := FindPatientInfo(ctx, record.PatientID)
	if err != nil {
		return errors.Wrap(err, "查询患者信息失败")
	}
	if patient == nil {
		return errors.New("患者信息不存在")
	}
	if patient.CameraID != "" { //固定摄像头
		record.CameraID = patient.CameraID
		record.VrCameraID = patient.VrCameraID
	} else { //需要分配可移动摄像头。需要检查是否与其他预约时间冲突
		visitRecordNewLock.Lock()
		defer visitRecordNewLock.Unlock()
		cameras, err := FindMovableCameras(ctx)
		if err != nil {
			return errors.Wrap(err, "查询可用摄像头失败")
		}
		if len(cameras) == 0 {
			return errors.New("没有可用摄像头")
		} else {
			existRecords, err := FindValidExistVisitRecordsByStartTime(ctx, record.VisitStartTime)
			if err != nil {
				return errors.Wrap(err, "查询已存在预约记录失败")
			}
			usedCameraIDMap := make(map[string]bool)
			for _, existRecord := range existRecords {
				usedCameraIDMap[existRecord.CameraID] = true
			}
			for _, camera := range cameras {
				if !usedCameraIDMap[camera.ID] {
					record.CameraID = camera.ID
					record.VrCameraID = camera.RelVrCameraID
					break
				}
			}
			if record.CameraID == "" {
				return errors.New("没有可用摄像头")
			}
		}
	}
	err = IncreaseVisitScheduleVisitors(ctx, schedule)
	if err != nil {
		return errors.Wrap(err, "更新排班已预约人数失败")
	}
	return nil
}

func findVisitRecordCountByRelativeIDAndStartDay(ctx context.Context, relativeID string, startDay common.LocalTime) (int, error) {
	startDayStr := time.Time(startDay).Format("2006-01-02")
	endDayStr := time.Time(startDay).AddDate(0, 0, 1).Format("2006-01-02")
	qstr := model.Visit_record_FIELD_NAME_relative_id + "=" + relativeID + "&" + model.Visit_record_FIELD_NAME_visit_start_time + "=$gt." + startDayStr + "&" + model.Visit_record_FIELD_NAME_visit_start_time + "=$lt." + endDayStr
	count, err := common.DbGetCount(ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, model.Visit_record_FIELD_NAME_id, qstr)
	if err != nil {
		return 0, errors.Wrap(err, "查询已存在预约记录失败")
	}
	return int(count), nil
}
func findVisitRecordByRelativeIDAndStartTime(ctx context.Context, relativeID string, startTime common.LocalTime) (*model.Visit_record, error) {
	qstr := model.Visit_record_FIELD_NAME_relative_id + "=" + relativeID + "&" + model.Visit_record_FIELD_NAME_visit_start_time + "=" + startTime.DbString()
	record, err := common.DbGetOne[model.Visit_record](ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func FindValidExistVisitRecordsByStartTime(ctx context.Context, startTime common.LocalTime) ([]model.Visit_record, error) {
	qstr := model.Visit_record_FIELD_NAME_visit_start_time + "=" + startTime.DbString() + "&" + model.Visit_record_FIELD_NAME_check_status + "=$ne." + cast.ToString(CHECK_STATUS_CHECK_FAILED)
	records, err := common.DbQuery[model.Visit_record](ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return nil, errors.Wrap(err, "查询已存在预约记录失败")
	}
	return records, nil
}
func FindNearestVisitRecord(ctx context.Context, familyMemberID string) (*model.Visit_record, error) {
	qstr := model.Visit_record_FIELD_NAME_relative_id + "=" + familyMemberID + "&_order=-" + model.Visit_record_FIELD_NAME_visit_start_time
	record, err := common.DbGetOne[model.Visit_record](ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return nil, err
	}
	return record, nil
}
func FindVisitRecord(ctx context.Context, visitRecordID string) (*model.Visit_record, error) {
	qstr := model.Visit_record_FIELD_NAME_id + "=" + visitRecordID
	record, err := common.DbGetOne[model.Visit_record](ctx, common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func LoopCheckVisitRecordWillAfter5Minutes() error {
	qstr := model.Visit_record_FIELD_NAME_visit_start_time + "=" + common.LocalTime(time.Now().Add(time.Minute*5)).DbString()
	qstr += "&" + model.Visit_record_FIELD_NAME_check_status + "=" + cast.ToString(CHECK_STATUS_CHECKED)
	qstr += "&" + model.Visit_record_FIELD_NAME_status + "=" + cast.ToString(VISIT_RECORD_STATUS_NORMAL)
	records, err := common.DbQuery[model.Visit_record](context.Background(), common.GetDaprClient(), model.Visit_recordTableInfo.Name, qstr)
	if err != nil {
		return errors.Wrap(err, "查询即将开始的预约记录失败")
	}
	for _, record := range records {
		if record.SendSmsStatus == SEND_SMS_STATUS_UNSENT {
			SendVisitRecordPromptSms(&record)
		}
	}
	return nil
}
func SendVisitRecordPromptSms(record *model.Visit_record) error {
	common.Logger.Infof("send visit record prompt sms: %v", record)
	//TODO: 发送短信
	return nil
}
func SendVisitRecordSms(record *model.Visit_record) error {
	common.Logger.Infof("send visit record sms: %v", record)
	visitRecordInfo, err := common.DbGetOne[model.Visit_record_info](context.Background(), common.GetDaprClient(), model.Visit_record_infoTableInfo.Name, model.Visit_record_info_FIELD_NAME_id+"="+record.ID)
	if err != nil {
		return errors.Wrap(err, "查询探视记录信息失败")
	}
	if visitRecordInfo == nil {
		return errors.New("探视记录信息不存在")
	}
	phone := visitRecordInfo.VisitorPhone
	if phone == "" {
		return errors.New("家属手机号不存在")
	}
	statusStr := "审核通过"
	if record.Status == VISIT_RECORD_STATUS_CANCEL {
		statusStr = "审核不通过"
	}
	templateParam := map[string]string{
		"time":   time.Time(record.VisitStartTime).Format("2006-01-02年 15时04分"),
		"name":   visitRecordInfo.PatientName,
		"status": statusStr,
	}
	err = sms.SendSms(config.ALI_SMS_REGION, config.ALI_SMS_ACCESS_ID, config.ALI_SMS_ACCESS_SECRET, config.ALI_SMS_SIGN_NAME, config.ALI_SMS_TEMPLATE_VISIT_CHECK_CODE, phone, templateParam)
	if err != nil {
		return errors.Wrap(err, "发送短信失败")
	}
	return nil
}
