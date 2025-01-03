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


Table: v_bed_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] ward_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] bed_no                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] camera_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] vr_camera_id                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] camera_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] vr_camera_name                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "oZmlbnHqsNcCoNghAfEuZpGhc",    "ward_id": "eaexmLVXpfkESZtltdOaCMlYD",    "bed_no": "snvlhqlDvylIUVXDrBWcXkqSa",    "camera_id": "nGcPnUXBIsuRQDYEBkRLIsNUY",    "vr_camera_id": "xprhFwcCRSuxyqaAgkoklLqao",    "type": 31,    "status": 23,    "ward_name": "vHkZZnOJmjpnRgaeXxJVOxOUD",    "camera_name": "xduhujTehfSoAgoUWrrADPmpa",    "vr_camera_name": "hgwedMapsCcsZFWiSBAbFVAmq"}


Comments
-------------------------------------
[ 0] Warning table: v_bed_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_bed_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Bed_info_FIELD_NAME_id = "id"

	Bed_info_FIELD_NAME_ward_id = "ward_id"

	Bed_info_FIELD_NAME_bed_no = "bed_no"

	Bed_info_FIELD_NAME_camera_id = "camera_id"

	Bed_info_FIELD_NAME_vr_camera_id = "vr_camera_id"

	Bed_info_FIELD_NAME_type = "type"

	Bed_info_FIELD_NAME_status = "status"

	Bed_info_FIELD_NAME_ward_name = "ward_name"

	Bed_info_FIELD_NAME_camera_name = "camera_name"

	Bed_info_FIELD_NAME_vr_camera_name = "vr_camera_name"
)

// Bed_info struct is a row record of the v_bed_info table in the  database
type Bed_info struct {
	ID string `json:"id"` //病床ID

	WardID string `json:"ward_id"` //病房ID

	BedNo string `json:"bed_no"` //床位号

	CameraID string `json:"camera_id"` //camera_id

	VrCameraID string `json:"vr_camera_id"` //vr_camera_id

	Type int32 `json:"type"` //床位类型(0:普通,1:ICU)

	Status int32 `json:"status"` //床位状态(0:空闲,1:占用)

	WardName string `json:"ward_name"` //病房名称

	CameraName string `json:"camera_name"` //床头摄像头名称

	VrCameraName string `json:"vr_camera_name"` //VR摄像头名称

}

var Bed_infoTableInfo = &TableInfo{
	Name: "v_bed_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `病床ID`,
			Notes: `Warning table: v_bed_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_bed_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "ward_id",
			Comment:            `病房ID`,
			Notes:              ``,
			Nullable:           true,
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
			Nullable:           true,
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
			Comment:            `camera_id`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "vr_camera_id",
			Comment:            `vr_camera_id`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "type",
			Comment:            `床位类型(0:普通,1:ICU)`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "ward_name",
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
			GoFieldName:        "WardName",
			GoFieldType:        "string",
			JSONFieldName:      "ward_name",
			ProtobufFieldName:  "ward_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "camera_name",
			Comment:            `床头摄像头名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "CameraName",
			GoFieldType:        "string",
			JSONFieldName:      "camera_name",
			ProtobufFieldName:  "camera_name",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "vr_camera_name",
			Comment:            `VR摄像头名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "VrCameraName",
			GoFieldType:        "string",
			JSONFieldName:      "vr_camera_name",
			ProtobufFieldName:  "vr_camera_name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Bed_info) TableName() string {
	return "v_bed_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Bed_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Bed_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Bed_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Bed_info) TableInfo() *TableInfo {
	return Bed_infoTableInfo
}
