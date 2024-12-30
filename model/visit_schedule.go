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


Table: o_visit_schedule
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] start_time                                     TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 2] end_time                                       TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] total_visitors                                 INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] schedule_visitors                              INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 5] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "WhVuhxXDWGkPUCOtDuqOqJNcH",    "start_time": 20,    "end_time": 23,    "total_visitors": 10,    "schedule_visitors": 16,    "status": 87}



*/

var (
	Visit_schedule_FIELD_NAME_id = "id"

	Visit_schedule_FIELD_NAME_start_time = "start_time"

	Visit_schedule_FIELD_NAME_end_time = "end_time"

	Visit_schedule_FIELD_NAME_total_visitors = "total_visitors"

	Visit_schedule_FIELD_NAME_schedule_visitors = "schedule_visitors"

	Visit_schedule_FIELD_NAME_status = "status"
)

// Visit_schedule struct is a row record of the o_visit_schedule table in the  database
type Visit_schedule struct {
	ID string `json:"id"` //排班ID

	StartTime common.LocalTime `json:"start_time"` //开始时间

	EndTime common.LocalTime `json:"end_time"` //结束时间

	TotalVisitors int32 `json:"total_visitors"` //探视总人数

	ScheduleVisitors int32 `json:"schedule_visitors"` //已预约人数

	Status int32 `json:"status"` //状态(0:不可预约,1:可预约)

}

var Visit_scheduleTableInfo = &TableInfo{
	Name: "o_visit_schedule",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `排班ID`,
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
			Name:               "start_time",
			Comment:            `开始时间`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "end_time",
			Comment:            `结束时间`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "total_visitors",
			Comment:            `探视总人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "TotalVisitors",
			GoFieldType:        "int32",
			JSONFieldName:      "total_visitors",
			ProtobufFieldName:  "total_visitors",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "schedule_visitors",
			Comment:            `已预约人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ScheduleVisitors",
			GoFieldType:        "int32",
			JSONFieldName:      "schedule_visitors",
			ProtobufFieldName:  "schedule_visitors",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `状态(0:不可预约,1:可预约)`,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *Visit_schedule) TableName() string {
	return "o_visit_schedule"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *Visit_schedule) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *Visit_schedule) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *Visit_schedule) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *Visit_schedule) TableInfo() *TableInfo {
	return Visit_scheduleTableInfo
}
