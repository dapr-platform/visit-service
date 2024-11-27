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


Table: v_head_display_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] device_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] device_no                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] model                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] ward_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "REadfTFxvAPMdhSlFZhoOYDtw",    "device_name": "BkiKvPiIBSiSKCBHwyuoYtNop",    "device_no": "fGiryHLGseqhfRmeZhyZUxePb",    "model": "FmXxObsKMwAXLWTuEgpoKVcKL",    "ward_id": "sHNHpmbsoFxuRbXiBwCysjIIR",    "ward_name": "CBhCayWCnGUZhqpJVtSnALonE"}


Comments
-------------------------------------
[ 0] Warning table: v_head_display_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_head_display_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Head_display_info_FIELD_NAME_id = "id"

	Head_display_info_FIELD_NAME_device_name = "device_name"

	Head_display_info_FIELD_NAME_device_no = "device_no"

	Head_display_info_FIELD_NAME_model = "model"

	Head_display_info_FIELD_NAME_ward_id = "ward_id"

	Head_display_info_FIELD_NAME_ward_name = "ward_name"
)

// Head_display_info struct is a row record of the v_head_display_info table in the  database
type Head_display_info struct {
	ID string `json:"id"` // [ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	DeviceName string `json:"device_name"` // [ 1] device_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []

	DeviceNo string `json:"device_no"` // [ 2] device_no                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	Model string `json:"model"` // [ 3] model                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []

	WardID string `json:"ward_id"` // [ 4] ward_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []

	WardName string `json:"ward_name"` // [ 5] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []

}

var Head_display_infoTableInfo = &TableInfo{
	Name: "v_head_display_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `头显ID`,
			Notes: `Warning table: v_head_display_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_head_display_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "device_name",
			Comment:            `设备名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DeviceName",
			GoFieldType:        "string",
			JSONFieldName:      "device_name",
			ProtobufFieldName:  "device_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "device_no",
			Comment:            `设备编号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "DeviceNo",
			GoFieldType:        "string",
			JSONFieldName:      "device_no",
			ProtobufFieldName:  "device_no",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "model",
			Comment:            `型号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Model",
			GoFieldType:        "string",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (h *Head_display_info) TableName() string {
	return "v_head_display_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *Head_display_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *Head_display_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *Head_display_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *Head_display_info) TableInfo() *TableInfo {
	return Head_display_infoTableInfo
}
