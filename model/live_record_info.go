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


Table: v_live_record_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] schedule_id                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] patient_id                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] relative_id                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] device_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] start_time                                     TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] end_time                                       TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 7] file_size                                      INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] stream_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] stream_url_suffix                              VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[10] camera_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] vr_camera_id                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[13] patient_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] patient_ward_name                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[15] patient_bed_no                                 VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []


JSON Sample
-------------------------------------
{    "id": "TEQWQDqymWqpOVAdwJCBqjWIE",    "schedule_id": "ejKVBuDEHooZgntQZWIAtlYlT",    "patient_id": "kacXRWGigNLQTmrKRHeFFIXBR",    "relative_id": "nUAMxqruGkBaSJELdQOonrHkl",    "device_id": "nRiNCFAcoinFsOospeeTbgOsx",    "start_time": 28,    "end_time": 87,    "file_size": 90,    "stream_id": "hMyDGampMcVhffhsPpKyebdgl",    "stream_url_suffix": "uPxqKYXWRybsWnotcBOTfHlRD",    "camera_id": "uHNKNDwMAIeHOlBYrbULixJgN",    "vr_camera_id": "OpbyCBfmyTkYxBgAypjUhuUTE",    "status": 29,    "patient_name": "sYHMyqaTmvMgNaOdIyoKplaeu",    "patient_ward_name": "giKIAkMavkAjOEJNYPgCMdoyp",    "patient_bed_no": "uiKnasIEZnaiPXDUKNdHDdoTA"}


Comments
-------------------------------------
[ 0] Warning table: v_live_record_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_live_record_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Live_record_info_FIELD_NAME_id = "id"

	Live_record_info_FIELD_NAME_schedule_id = "schedule_id"

	Live_record_info_FIELD_NAME_patient_id = "patient_id"

	Live_record_info_FIELD_NAME_relative_id = "relative_id"

	Live_record_info_FIELD_NAME_device_id = "device_id"

	Live_record_info_FIELD_NAME_start_time = "start_time"

	Live_record_info_FIELD_NAME_end_time = "end_time"

	Live_record_info_FIELD_NAME_file_size = "file_size"

	Live_record_info_FIELD_NAME_stream_id = "stream_id"

	Live_record_info_FIELD_NAME_stream_url_suffix = "stream_url_suffix"

	Live_record_info_FIELD_NAME_camera_id = "camera_id"

	Live_record_info_FIELD_NAME_vr_camera_id = "vr_camera_id"

	Live_record_info_FIELD_NAME_status = "status"

	Live_record_info_FIELD_NAME_patient_name = "patient_name"

	Live_record_info_FIELD_NAME_patient_ward_name = "patient_ward_name"

	Live_record_info_FIELD_NAME_patient_bed_no = "patient_bed_no"
)

// Live_record_info struct is a row record of the v_live_record_info table in the  database
type Live_record_info struct {
	ID string `json:"id"` //直播记录ID

	ScheduleID string `json:"schedule_id"` //探视排班ID

	PatientID string `json:"patient_id"` //病患ID

	RelativeID string `json:"relative_id"` //家属ID

	DeviceID string `json:"device_id"` //设备ID

	StartTime common.LocalTime `json:"start_time"` //直播开始时间

	EndTime common.LocalTime `json:"end_time"` //直播结束时间

	FileSize int32 `json:"file_size"` //文件大小

	StreamID string `json:"stream_id"` //流ID

	StreamURLSuffix string `json:"stream_url_suffix"` //流URL后缀

	CameraID string `json:"camera_id"` //床头摄像头ID

	VrCameraID string `json:"vr_camera_id"` //VR摄像头ID

	Status int32 `json:"status"` //状态(0:未开始,1:直播中,2:已结束)

	PatientName string `json:"patient_name"` //病患姓名

	PatientWardName string `json:"patient_ward_name"` //病房名称

	PatientBedNo string `json:"patient_bed_no"` //床位号

}

var Live_record_infoTableInfo = &TableInfo{
	Name: "v_live_record_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `直播记录ID`,
			Notes: `Warning table: v_live_record_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_live_record_info primary key column id is nullable column, setting it as NOT NULL
`,
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
			Nullable:           true,
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
			Name:               "patient_id",
			Comment:            `病患ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "PatientID",
			GoFieldType:        "string",
			JSONFieldName:      "patient_id",
			ProtobufFieldName:  "patient_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "relative_id",
			Comment:            `家属ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RelativeID",
			GoFieldType:        "string",
			JSONFieldName:      "relative_id",
			ProtobufFieldName:  "relative_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "device_id",
			Comment:            `设备ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "DeviceID",
			GoFieldType:        "string",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "start_time",
			Comment:            `直播开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "end_time",
			Comment:            `直播结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "file_size",
			Comment:            `文件大小`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "FileSize",
			GoFieldType:        "int32",
			JSONFieldName:      "file_size",
			ProtobufFieldName:  "file_size",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "stream_id",
			Comment:            `流ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "StreamID",
			GoFieldType:        "string",
			JSONFieldName:      "stream_id",
			ProtobufFieldName:  "stream_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "stream_url_suffix",
			Comment:            `流URL后缀`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "StreamURLSuffix",
			GoFieldType:        "string",
			JSONFieldName:      "stream_url_suffix",
			ProtobufFieldName:  "stream_url_suffix",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "camera_id",
			Comment:            `床头摄像头ID`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "vr_camera_id",
			Comment:            `VR摄像头ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "VrCameraID",
			GoFieldType:        "string",
			JSONFieldName:      "vr_camera_id",
			ProtobufFieldName:  "vr_camera_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
			Comment:            `状态(0:未开始,1:直播中,2:已结束)`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "patient_name",
			Comment:            `病患姓名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "PatientName",
			GoFieldType:        "string",
			JSONFieldName:      "patient_name",
			ProtobufFieldName:  "patient_name",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "patient_ward_name",
			Comment:            `病房名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "PatientWardName",
			GoFieldType:        "string",
			JSONFieldName:      "patient_ward_name",
			ProtobufFieldName:  "patient_ward_name",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "patient_bed_no",
			Comment:            `床位号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "PatientBedNo",
			GoFieldType:        "string",
			JSONFieldName:      "patient_bed_no",
			ProtobufFieldName:  "patient_bed_no",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *Live_record_info) TableName() string {
	return "v_live_record_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *Live_record_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *Live_record_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *Live_record_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *Live_record_info) TableInfo() *TableInfo {
	return Live_record_infoTableInfo
}
