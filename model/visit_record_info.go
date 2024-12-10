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


Table: v_visit_record_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] patient_id                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] relative_id                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] visit_start_time                               TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] visit_end_time                                 TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] visitor_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] visitor_phone                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] visitor_id_card                                VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] relationship                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] camera_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] vr_camera_id                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] check_status                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[12] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[13] remark                                         VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[14] send_sms_status                                INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] send_prompt_sms_status                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[16] schedule_id                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[17] patient_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[18] patient_ward_name                              VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[19] patient_bed_no                                 VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[20] stream_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[21] stream_url_suffix                              VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []


JSON Sample
-------------------------------------
{    "id": "AQwQGjnPAkWgRFSTPbhVQjPtx",    "patient_id": "QBBvwJkeTcyDAlWZuXXGngUgA",    "relative_id": "bjcLZdWkhEtEfpboXuSYTJFGs",    "visit_start_time": 13,    "visit_end_time": 70,    "visitor_name": "xpMMMOhXdYsZqUfnWsnODZXyF",    "visitor_phone": "okfjbfugYgioKtgTcgRsaARiE",    "visitor_id_card": "ewWQQbSbZIdPMHrOYaRVBNmBU",    "relationship": "UVxgVDaQkByouWYmlVsRoqhmB",    "camera_id": "sOBXoelrFNZbcwZpRVyOnmhCA",    "vr_camera_id": "tbZPskKewvrKxMrsZcOmHFsUt",    "check_status": 95,    "status": 21,    "remark": "qNDrLFmLbRcotwtbAmRynwQpV",    "send_sms_status": 66,    "send_prompt_sms_status": 59,    "schedule_id": "cFxXVOeuMImiVkMSkGXVhBRfu",    "patient_name": "dkdsFiomjNCCjjEQGZHilvrQk",    "patient_ward_name": "aDxDElNAisjjLlAJYqjIoJmOM",    "patient_bed_no": "hsOkrHxfuHDjwIhoVppvhluXT",    "stream_id": "oegRoFmkjEfDCKWvPZaZwdKVJ",    "stream_url_suffix": "QSsVtUxTrjxaUZYGrpDtyRBCJ"}


Comments
-------------------------------------
[ 0] Warning table: v_visit_record_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_visit_record_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Visit_record_info_FIELD_NAME_id = "id"

	Visit_record_info_FIELD_NAME_patient_id = "patient_id"

	Visit_record_info_FIELD_NAME_relative_id = "relative_id"

	Visit_record_info_FIELD_NAME_visit_start_time = "visit_start_time"

	Visit_record_info_FIELD_NAME_visit_end_time = "visit_end_time"

	Visit_record_info_FIELD_NAME_visitor_name = "visitor_name"

	Visit_record_info_FIELD_NAME_visitor_phone = "visitor_phone"

	Visit_record_info_FIELD_NAME_visitor_id_card = "visitor_id_card"

	Visit_record_info_FIELD_NAME_relationship = "relationship"

	Visit_record_info_FIELD_NAME_camera_id = "camera_id"

	Visit_record_info_FIELD_NAME_vr_camera_id = "vr_camera_id"

	Visit_record_info_FIELD_NAME_check_status = "check_status"

	Visit_record_info_FIELD_NAME_status = "status"

	Visit_record_info_FIELD_NAME_remark = "remark"

	Visit_record_info_FIELD_NAME_send_sms_status = "send_sms_status"

	Visit_record_info_FIELD_NAME_send_prompt_sms_status = "send_prompt_sms_status"

	Visit_record_info_FIELD_NAME_schedule_id = "schedule_id"

	Visit_record_info_FIELD_NAME_patient_name = "patient_name"

	Visit_record_info_FIELD_NAME_patient_ward_name = "patient_ward_name"

	Visit_record_info_FIELD_NAME_patient_bed_no = "patient_bed_no"

	Visit_record_info_FIELD_NAME_stream_id = "stream_id"

	Visit_record_info_FIELD_NAME_stream_url_suffix = "stream_url_suffix"
)

// Visit_record_info struct is a row record of the v_visit_record_info table in the  database
type Visit_record_info struct {
	ID string `json:"id"` //探视记录ID

	PatientID string `json:"patient_id"` //病患ID

	RelativeID string `json:"relative_id"` //家属ID

	VisitStartTime common.LocalTime `json:"visit_start_time"` //探视开始时间

	VisitEndTime common.LocalTime `json:"visit_end_time"` //探视结束时间

	VisitorName string `json:"visitor_name"` //探视人姓名

	VisitorPhone string `json:"visitor_phone"` //探视人电话

	VisitorIDCard string `json:"visitor_id_card"` //探视人身份证号

	Relationship string `json:"relationship"` //探视人与患者关系(父母，配偶，子女，其他)

	CameraID string `json:"camera_id"` //床头摄像头ID

	VrCameraID string `json:"vr_camera_id"` //VR摄像头ID

	CheckStatus int32 `json:"check_status"` //审核状态(0:未审核,1:已审核,2:审核不通过)

	Status int32 `json:"status"` //状态(0:正常,1:取消)

	Remark string `json:"remark"` //备注

	SendSmsStatus int32 `json:"send_sms_status"` //send_sms_status

	SendPromptSmsStatus int32 `json:"send_prompt_sms_status"` //send_prompt_sms_status

	ScheduleID string `json:"schedule_id"` //排班ID

	PatientName string `json:"patient_name"` //病患姓名

	PatientWardName string `json:"patient_ward_name"` //病房名称

	PatientBedNo string `json:"patient_bed_no"` //床位号

	StreamID string `json:"stream_id"` //直播流ID

	StreamURLSuffix string `json:"stream_url_suffix"` //stream_url_suffix

}

var Visit_record_infoTableInfo = &TableInfo{
	Name: "v_visit_record_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `探视记录ID`,
			Notes: `Warning table: v_visit_record_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_visit_record_info primary key column id is nullable column, setting it as NOT NULL
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "visit_start_time",
			Comment:            `探视开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "VisitStartTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "visit_start_time",
			ProtobufFieldName:  "visit_start_time",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "visit_end_time",
			Comment:            `探视结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "VisitEndTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "visit_end_time",
			ProtobufFieldName:  "visit_end_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "visitor_name",
			Comment:            `探视人姓名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "VisitorName",
			GoFieldType:        "string",
			JSONFieldName:      "visitor_name",
			ProtobufFieldName:  "visitor_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "visitor_phone",
			Comment:            `探视人电话`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "VisitorPhone",
			GoFieldType:        "string",
			JSONFieldName:      "visitor_phone",
			ProtobufFieldName:  "visitor_phone",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "visitor_id_card",
			Comment:            `探视人身份证号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "VisitorIDCard",
			GoFieldType:        "string",
			JSONFieldName:      "visitor_id_card",
			ProtobufFieldName:  "visitor_id_card",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "relationship",
			Comment:            `探视人与患者关系(父母，配偶，子女，其他)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "Relationship",
			GoFieldType:        "string",
			JSONFieldName:      "relationship",
			ProtobufFieldName:  "relationship",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "check_status",
			Comment:            `审核状态(0:未审核,1:已审核,2:审核不通过)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "CheckStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "check_status",
			ProtobufFieldName:  "check_status",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
			Comment:            `状态(0:正常,1:取消)`,
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
			Name:               "remark",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "send_sms_status",
			Comment:            `send_sms_status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SendSmsStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "send_sms_status",
			ProtobufFieldName:  "send_sms_status",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "send_prompt_sms_status",
			Comment:            `send_prompt_sms_status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SendPromptSmsStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "send_prompt_sms_status",
			ProtobufFieldName:  "send_prompt_sms_status",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "schedule_id",
			Comment:            `排班ID`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "stream_id",
			Comment:            `直播流ID`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "stream_url_suffix",
			Comment:            `stream_url_suffix`,
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
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *Visit_record_info) TableName() string {
	return "v_visit_record_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *Visit_record_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *Visit_record_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *Visit_record_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *Visit_record_info) TableInfo() *TableInfo {
	return Visit_record_infoTableInfo
}
