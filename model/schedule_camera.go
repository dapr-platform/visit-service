package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_schedule_camera
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] schedule_id                                    VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] camera_id                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "IsAcXvDXuJxRIKUrLQfSqXUYu",    "schedule_id": "hAGkELKGZssNdKdVWvmfqtlcV",    "camera_id": "jgDuXhpoFYENFwGmlvuPfRguH",    "status": 44}



*/

var (
	Schedule_camera_FIELD_NAME_id = "id"

	Schedule_camera_FIELD_NAME_schedule_id = "schedule_id"

	Schedule_camera_FIELD_NAME_camera_id = "camera_id"

	Schedule_camera_FIELD_NAME_status = "status"
)

// Schedule_camera struct is a row record of the o_schedule_camera table in the  database
type Schedule_camera struct {
	ID string `json:"id"` // [ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	ScheduleID string `json:"schedule_id"` // [ 1] schedule_id                                    VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	CameraID string `json:"camera_id"` // [ 2] camera_id                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	Status int32 `json:"status"` // [ 3] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]

}

var Schedule_cameraTableInfo = &TableInfo{
	Name: "o_schedule_camera",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `排班摄像头ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "schedule_id",
			Comment:            `探视排班ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ScheduleID",
			GoFieldType:        "string",
			JSONFieldName:      "schedule_id",
			ProtobufFieldName:  "schedule_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "camera_id",
			Comment:            `摄像头ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CameraID",
			GoFieldType:        "string",
			JSONFieldName:      "camera_id",
			ProtobufFieldName:  "camera_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Schedule_camera) TableName() string {
	return "o_schedule_camera"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Schedule_camera) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Schedule_camera) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Schedule_camera) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Schedule_camera) TableInfo() *TableInfo {
	return Schedule_cameraTableInfo
}
