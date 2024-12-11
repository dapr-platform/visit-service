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


Table: o_camera
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] device_name                                    VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] device_no                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] location_type                                  INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 4] device_type                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 5] manufacturer                                   VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] main_stream_url                                VARCHAR(1024)        null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 7] sub_stream_url                                 VARCHAR(1024)        null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 8] rel_vr_camera_id                               VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "hGrhOsyHqdTXViuVhpSlZtBho",    "device_name": "KsHDDaGhARtcHCoBoIdhfLNcW",    "device_no": "SQKFFkHgYfgtUmvRgUqhbKcaF",    "location_type": 97,    "device_type": 63,    "manufacturer": "dtbQqKKWiNOoqquTTkcoWqZfj",    "main_stream_url": "OfKgpFfitVWeKTxeIsGtchTuG",    "sub_stream_url": "RfdXHywTmYiQpfbWJAbiHDMPL",    "rel_vr_camera_id": "aRdCQpNluwGnhGpAomDANIDkd",    "status": 29}



*/

var (
	Camera_FIELD_NAME_id = "id"

	Camera_FIELD_NAME_device_name = "device_name"

	Camera_FIELD_NAME_device_no = "device_no"

	Camera_FIELD_NAME_location_type = "location_type"

	Camera_FIELD_NAME_device_type = "device_type"

	Camera_FIELD_NAME_manufacturer = "manufacturer"

	Camera_FIELD_NAME_main_stream_url = "main_stream_url"

	Camera_FIELD_NAME_sub_stream_url = "sub_stream_url"

	Camera_FIELD_NAME_rel_vr_camera_id = "rel_vr_camera_id"

	Camera_FIELD_NAME_status = "status"
)

// Camera struct is a row record of the o_camera table in the  database
type Camera struct {
	ID string `json:"id"` //摄像头ID

	DeviceName string `json:"device_name"` //设备名称

	DeviceNo string `json:"device_no"` //设备编号

	LocationType int32 `json:"location_type"` //位置类型

	DeviceType int32 `json:"device_type"` //设备类型

	Manufacturer string `json:"manufacturer"` //设备厂商

	MainStreamURL string `json:"main_stream_url"` //主码流URL

	SubStreamURL string `json:"sub_stream_url"` //辅码流URL

	RelVrCameraID string `json:"rel_vr_camera_id"` //关联VR摄像头ID(当位置为可移动时，需关联VR摄像头)

	Status int32 `json:"status"` //状态(0:正常,1:禁用)

}

var CameraTableInfo = &TableInfo{
	Name: "o_camera",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `摄像头ID`,
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
			Name:               "device_name",
			Comment:            `设备名称`,
			Notes:              ``,
			Nullable:           false,
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
			Nullable:           false,
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
			Name:               "location_type",
			Comment:            `位置类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "LocationType",
			GoFieldType:        "int32",
			JSONFieldName:      "location_type",
			ProtobufFieldName:  "location_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "device_type",
			Comment:            `设备类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "DeviceType",
			GoFieldType:        "int32",
			JSONFieldName:      "device_type",
			ProtobufFieldName:  "device_type",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "manufacturer",
			Comment:            `设备厂商`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Manufacturer",
			GoFieldType:        "string",
			JSONFieldName:      "manufacturer",
			ProtobufFieldName:  "manufacturer",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "main_stream_url",
			Comment:            `主码流URL`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "MainStreamURL",
			GoFieldType:        "string",
			JSONFieldName:      "main_stream_url",
			ProtobufFieldName:  "main_stream_url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sub_stream_url",
			Comment:            `辅码流URL`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "SubStreamURL",
			GoFieldType:        "string",
			JSONFieldName:      "sub_stream_url",
			ProtobufFieldName:  "sub_stream_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "rel_vr_camera_id",
			Comment:            `关联VR摄像头ID(当位置为可移动时，需关联VR摄像头)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RelVrCameraID",
			GoFieldType:        "string",
			JSONFieldName:      "rel_vr_camera_id",
			ProtobufFieldName:  "rel_vr_camera_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
			Comment:            `状态(0:正常,1:禁用)`,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Camera) TableName() string {
	return "o_camera"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Camera) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Camera) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Camera) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Camera) TableInfo() *TableInfo {
	return CameraTableInfo
}
