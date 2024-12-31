package service

import (
	"context"
	"time"
	"visit-service/config"
	"visit-service/model"

	"github.com/dapr-platform/common"
)

var RECORD_STATUS_STARTING = 1
var RECORD_STATUS_ENDING = 2

func AddLiveRecord(ctx context.Context, visitRecord *model.Visit_record, userID, cameraID, streamID string) error {
	record := &model.Live_record{
		ID:              common.NanoId(),
		ScheduleID:      visitRecord.ID,
		PatientID:       visitRecord.PatientID,
		RelativeID:      visitRecord.RelativeID,
		CameraID:        cameraID,
		VrCameraID:      visitRecord.VrCameraID,
		StreamID:        streamID,
		StreamURLSuffix: config.ZLMEDIAKIT_STREAM_URL_PREFIX + "live/" + streamID + ".live.flv",
		StartTime:       common.LocalTime(time.Now()),
		Status:          int32(RECORD_STATUS_STARTING),
		VisitRecordID:   visitRecord.ID,
	}
	_, err := common.DbInsert(ctx, common.GetDaprClient(), record, model.Live_recordTableInfo.Name)
	if err != nil {
		common.Logger.Error("add live record", "visitRecord", visitRecord, "err", err)
		return err
	}
	return nil
}
func StopLiveRecordOnlyStatus(ctx context.Context, streamID string) error {
	qstr := model.Live_record_FIELD_NAME_stream_id + "=" + streamID
	record, err := common.DbGetOne[model.Live_record](ctx, common.GetDaprClient(), model.Live_recordTableInfo.Name, qstr)
	if err != nil {
		return err
	}
	if record == nil {
		return nil
	}
	record.EndTime = common.LocalTime(time.Now())
	record.Status = int32(RECORD_STATUS_ENDING)
	err = common.DbUpsert(ctx, common.GetDaprClient(), *record, model.Live_recordTableInfo.Name, model.Live_record_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error("stop live record", "streamID", streamID, "record", record, "err", err)
		return err
	}
	return nil
}
func StopLiveRecord(ctx context.Context, streamID string, fileSize int64, url string) error {
	qstr := model.Live_record_FIELD_NAME_stream_id + "=" + streamID
	record, err := common.DbGetOne[model.Live_record](ctx, common.GetDaprClient(), model.Live_recordTableInfo.Name, qstr)
	if err != nil {
		return err
	}
	if record == nil {
		common.Logger.Error("stop live record", "streamID", streamID, "record", record)
		return nil
	}
	record.EndTime = common.LocalTime(time.Now())
	record.Status = int32(RECORD_STATUS_ENDING)
	record.FileSize = int32(fileSize)
	record.StreamURLSuffix = config.ZLMEDIAKIT_STREAM_URL_PREFIX + url
	err = common.DbUpsert(ctx, common.GetDaprClient(), *record, model.Live_recordTableInfo.Name, model.Live_record_FIELD_NAME_id)
	if err != nil {
		common.Logger.Error("stop live record", "streamID", streamID, "record", record, "err", err)
		return err
	}
	return nil
}
