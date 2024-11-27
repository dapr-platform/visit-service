package service

import (
	"context"
	"net/http"
	"sync"
	"visit-service/model"

	"github.com/dapr-platform/common"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

var CHECK_STATUS_UNCHECKED int32 = 0     //未审核
var CHECK_STATUS_CHECKED int32 = 1       //已审核
var CHECK_STATUS_CHECK_FAILED int32 = 2  //审核不通过
var VISIT_RECORD_STATUS_NORMAL int32 = 0 //正常
var VISIT_RECORD_STATUS_CANCEL int32 = 1 //取消

var visitRecordNewLock = sync.Mutex{}

func init() {

	common.RegisterUpsertBeforeHook("Visit_record", UpsertVisit_record)
}

func UpsertVisit_record(r *http.Request, in any) (out any, err error) {
	record := in.(model.Visit_record)
	if record.CheckStatus == CHECK_STATUS_UNCHECKED { //未审核
		if record.Status == VISIT_RECORD_STATUS_NORMAL { //正常状态可以新建
			err = newAddVisitRecord(r.Context(), &record)
		} else if record.Status == VISIT_RECORD_STATUS_CANCEL { //取消状态，需要释放摄像头,并增加排班剩余人数
			err = cancelVisitRecord(r.Context(), &record)
		}
	} else if record.CheckStatus == CHECK_STATUS_CHECK_FAILED { //审核不通过，删除摄像头信息
		record.CameraID = ""
		record.VrCameraID = ""
	}
	return record, err
}
func cancelVisitRecord(ctx context.Context, record *model.Visit_record) error {
	schedule, err := FindVisitScheduleByStartTime(ctx, record.VisitStartTime)
	if err != nil {
		return errors.Wrap(err, "查询排班信息失败")
	}
	if schedule != nil {
		err = IncreaseVisitScheduleRemainingVisitors(ctx, schedule)
		if err != nil {
			return errors.Wrap(err, "增加排班剩余人数失败")
		}
	}
	record.CameraID = ""
	record.VrCameraID = ""
	return nil
}

func newAddVisitRecord(ctx context.Context, record *model.Visit_record) error {
	schedule, err := FindVisitScheduleByStartTime(ctx, record.VisitStartTime)
	if err != nil {
		return errors.Wrap(err, "查询排班信息失败")
	}
	if schedule == nil {
		return errors.New("排班信息不存在")
	}
	if schedule.RemainingVisitors <= 0 {
		return errors.New("排班剩余人数不足")
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
	err = DecreaseVisitScheduleRemainingVisitors(ctx, schedule)
	if err != nil {
		return errors.Wrap(err, "更新排班剩余人数失败")
	}
	return nil
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
