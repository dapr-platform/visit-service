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


Table: o_visit_record_live_record
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] visit_record_id                                VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] live_record_id                                 VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []


JSON Sample
-------------------------------------
{    "id": "soeVHaFBtmlVnZaHNvdGHDVJf",    "visit_record_id": "nqbdhhiTlJJpjuugsanOBLcmg",    "live_record_id": "SRKPUWxXvsAVvacgOdDjjomuN"}



*/

var (
	Visit_record_live_record_FIELD_NAME_id = "id"

	Visit_record_live_record_FIELD_NAME_visit_record_id = "visit_record_id"

	Visit_record_live_record_FIELD_NAME_live_record_id = "live_record_id"
)

// Visit_record_live_record struct is a row record of the o_visit_record_live_record table in the  database
type Visit_record_live_record struct {
	ID string `json:"id"` //关联ID

	VisitRecordID string `json:"visit_record_id"` //探视记录ID

	LiveRecordID string `json:"live_record_id"` //直播记录ID

}

var Visit_record_live_recordTableInfo = &TableInfo{
	Name: "o_visit_record_live_record",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `关联ID`,
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
			Name:               "visit_record_id",
			Comment:            `探视记录ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "VisitRecordID",
			GoFieldType:        "string",
			JSONFieldName:      "visit_record_id",
			ProtobufFieldName:  "visit_record_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "live_record_id",
			Comment:            `直播记录ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "LiveRecordID",
			GoFieldType:        "string",
			JSONFieldName:      "live_record_id",
			ProtobufFieldName:  "live_record_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (v *Visit_record_live_record) TableName() string {
	return "o_visit_record_live_record"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (v *Visit_record_live_record) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (v *Visit_record_live_record) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (v *Visit_record_live_record) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (v *Visit_record_live_record) TableInfo() *TableInfo {
	return Visit_record_live_recordTableInfo
}
