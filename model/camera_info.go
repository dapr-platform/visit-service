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
[ 4] ward_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] bed_id                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] device_type                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] manufacturer                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] main_stream_url                                VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 9] sub_stream_url                                 VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[10] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[11] bed_no                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "gOdlqeWLpqRslXbwchJXclMZK",    "device_name": "QiNOnkhLnhEBUXqWoIPXnLrDD",    "device_no": "bhWMOmHcPNJFtgwdDQwrcJqFf",    "location_type": 18,    "ward_id": "jtmqQGltdcAYYVJQrWyUmGDqY",    "bed_id": "udmTWmarHdVTTePnIZwfLpprF",    "device_type": 66,    "manufacturer": "diTmvXFZOBCDSfRwUlhCLKHjw",    "main_stream_url": "foIdlXkYfQDQeFlMeiHnshZhM",    "sub_stream_url": "siYsImLhjxoSkKtSikbMslioo",    "status": 22,    "bed_no": "muXUjZRtkSMASNYcqiOBJvCgl",    "ward_name": "UEVGvFJxKIWHmTgLwEKtPnPuI"}


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

	Camera_info_FIELD_NAME_ward_id = "ward_id"

	Camera_info_FIELD_NAME_bed_id = "bed_id"

	Camera_info_FIELD_NAME_device_type = "device_type"

	Camera_info_FIELD_NAME_manufacturer = "manufacturer"

	Camera_info_FIELD_NAME_main_stream_url = "main_stream_url"

	Camera_info_FIELD_NAME_sub_stream_url = "sub_stream_url"

	Camera_info_FIELD_NAME_status = "status"

	Camera_info_FIELD_NAME_bed_no = "bed_no"

	Camera_info_FIELD_NAME_ward_name = "ward_name"
)

// Camera_info struct is a row record of the v_camera_info table in the  database
type Camera_info struct {
	ID            string `json:"id"`              //id
	DeviceName    string `json:"device_name"`     //device_name
	DeviceNo      string `json:"device_no"`       //device_no
	LocationType  int32  `json:"location_type"`   //location_type
	WardID        string `json:"ward_id"`         //ward_id
	BedID         string `json:"bed_id"`          //bed_id
	DeviceType    int32  `json:"device_type"`     //device_type
	Manufacturer  string `json:"manufacturer"`    //manufacturer
	MainStreamURL string `json:"main_stream_url"` //main_stream_url
	SubStreamURL  string `json:"sub_stream_url"`  //sub_stream_url
	Status        int32  `json:"status"`          //status
	BedNo         string `json:"bed_no"`          //bed_no
	WardName      string `json:"ward_name"`       //ward_name

}

var Camera_infoTableInfo = &TableInfo{
	Name: "v_camera_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
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
			Comment:            `device_name`,
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
			Comment:            `device_no`,
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
			Comment:            `location_type`,
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
			Name:               "ward_id",
			Comment:            `ward_id`,
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
			Name:               "bed_id",
			Comment:            `bed_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "BedID",
			GoFieldType:        "string",
			JSONFieldName:      "bed_id",
			ProtobufFieldName:  "bed_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "device_type",
			Comment:            `device_type`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "manufacturer",
			Comment:            `manufacturer`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "main_stream_url",
			Comment:            `main_stream_url`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "sub_stream_url",
			Comment:            `sub_stream_url`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "status",
			Comment:            `status`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "bed_no",
			Comment:            `bed_no`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "ward_name",
			Comment:            `ward_name`,
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
			ProtobufPos:        13,
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