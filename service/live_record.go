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
		ID:         common.NanoId(),
		ScheduleID: visitRecord.ID,
		PatientID:  visitRecord.PatientID,
		RelativeID: visitRecord.RelativeID,
		DeviceID:   cameraID,
		StreamID:   streamID,
		StartTime:  common.LocalTime(time.Now()),
		Status:     int32(RECORD_STATUS_STARTING),
	}
	_, err := common.DbInsert(ctx, common.GetDaprClient(), record, model.Live_recordTableInfo.Name)
	return err
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
	return common.DbUpsert(ctx, common.GetDaprClient(), *record, model.Live_recordTableInfo.Name, model.Live_record_FIELD_NAME_id)
}
