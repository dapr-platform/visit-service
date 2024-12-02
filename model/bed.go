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


Table: o_bed
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] ward_id                                        VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] bed_no                                         VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] camera_id                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] vr_camera_id                                   VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 6] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "TLydaJMJrEqcLmnnKJCTYTaKl",    "ward_id": "loxFXVPYPyCwjqxXUwtbPWJWK",    "bed_no": "gSRaKcwvSCgjeWDWNQCClIaMV",    "camera_id": "LfUlUvoIUYYplNRiIqddxcMHd",    "vr_camera_id": "rlroirGloguiFOAjalNfhiDGH",    "type": 42,    "status": 9}



*/

var (
	Bed_FIELD_NAME_id = "id"

	Bed_FIELD_NAME_ward_id = "ward_id"

	Bed_FIELD_NAME_bed_no = "bed_no"

	Bed_FIELD_NAME_camera_id = "camera_id"

	Bed_FIELD_NAME_vr_camera_id = "vr_camera_id"

	Bed_FIELD_NAME_type = "type"

	Bed_FIELD_NAME_status = "status"
)

// Bed struct is a row record of the o_bed table in the  database
type Bed struct {
	ID string `json:"id"` //病床ID

	WardID string `json:"ward_id"` //病房ID

	BedNo string `json:"bed_no"` //床位号

	CameraID string `json:"camera_id"` //摄像头ID(可关联固定床头摄像头)

	VrCameraID string `json:"vr_camera_id"` //VR摄像头ID(可关联固定VR摄像头)

	Type int32 `json:"type"` //床位类型(0:普通,1:ICU)

	Status int32 `json:"status"` //床位状态(0:空闲,1:占用)

}

var BedTableInfo = &TableInfo{
	Name: "o_bed",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `病床ID`,
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
			Name:               "ward_id",
			Comment:            `病房ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "WardID",
			GoFieldType:        "string",
			JSONFieldName:      "ward_id",
			ProtobufFieldName:  "ward_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "bed_no",
			Comment:            `床位号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "BedNo",
			GoFieldType:        "string",
			JSONFieldName:      "bed_no",
			ProtobufFieldName:  "bed_no",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "camera_id",
			Comment:            `摄像头ID(可关联固定床头摄像头)`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "vr_camera_id",
			Comment:            `VR摄像头ID(可关联固定VR摄像头)`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "type",
			Comment:            `床位类型(0:普通,1:ICU)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `床位状态(0:空闲,1:占用)`,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Bed) TableName() string {
	return "o_bed"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Bed) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Bed) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Bed) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Bed) TableInfo() *TableInfo {
	return BedTableInfo
}
