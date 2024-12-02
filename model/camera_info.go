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


Table: v_camera_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] device_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] device_no                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] location_type                                  INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] device_type                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] manufacturer                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] main_stream_url                                VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 7] sub_stream_url                                 VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 8] rel_vr_camera_id                               VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "NcdNGTQqTatiSgPWpiRCeljHo",    "device_name": "cigFhDXmfVCZRumLjclIuEnEH",    "device_no": "foXevdtGmObndTtxsWJDLwqMe",    "location_type": 49,    "device_type": 81,    "manufacturer": "XhSItkBnWUTJsEQdVKBENCXUH",    "main_stream_url": "jIjKuHKUumSndcSaiIchQdkQy",    "sub_stream_url": "rGNAIiTdlGiBdMVOTjjWioLmc",    "rel_vr_camera_id": "GwGhBDrfHlPYWeqmLnnDghuXW",    "status": 44}


Comments
-------------------------------------
[ 0] Warning table: v_camera_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_camera_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Camera_info_FIELD_NAME_id = "id"

	Camera_info_FIELD_NAME_device_name = "device_name"

	Camera_info_FIELD_NAME_device_no = "device_no"

	Camera_info_FIELD_NAME_location_type = "location_type"

	Camera_info_FIELD_NAME_device_type = "device_type"

	Camera_info_FIELD_NAME_manufacturer = "manufacturer"

	Camera_info_FIELD_NAME_main_stream_url = "main_stream_url"

	Camera_info_FIELD_NAME_sub_stream_url = "sub_stream_url"

	Camera_info_FIELD_NAME_rel_vr_camera_id = "rel_vr_camera_id"

	Camera_info_FIELD_NAME_status = "status"
)

// Camera_info struct is a row record of the v_camera_info table in the  database
type Camera_info struct {
	ID string `json:"id"` //摄像头ID

	DeviceName string `json:"device_name"` //设备名称

	DeviceNo string `json:"device_no"` //设备编号

	LocationType int32 `json:"location_type"` //位置类型(0:固定,1:可移动)

	DeviceType int32 `json:"device_type"` //设备类型(0:普通,1:普通云台,2:VR)

	Manufacturer string `json:"manufacturer"` //设备厂商

	MainStreamURL string `json:"main_stream_url"` //主码流URL

	SubStreamURL string `json:"sub_stream_url"` //辅码流URL

	RelVrCameraID string `json:"rel_vr_camera_id"` //关联VR摄像头ID(当位置为可移动时，需关联VR摄像头)

	Status int32 `json:"status"` //状态(0:正常,1:禁用)

}

var Camera_infoTableInfo = &TableInfo{
	Name: "v_camera_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `摄像头ID`,
			Notes: `Warning table: v_camera_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_camera_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "location_type",
			Comment:            `位置类型(0:固定,1:可移动)`,
			Notes:              ``,
			Nullable:           true,
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
			Comment:            `设备类型(0:普通,1:普通云台,2:VR)`,
			Notes:              ``,
			Nullable:           true,
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
			Nullable:           true,
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
			Nullable:           true,
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
			Nullable:           true,
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
			Nullable:           true,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Camera_info) TableName() string {
	return "v_camera_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Camera_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Camera_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Camera_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Camera_info) TableInfo() *TableInfo {
	return Camera_infoTableInfo
}
