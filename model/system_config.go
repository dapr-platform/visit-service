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


Table: o_system_config
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] config_name                                    VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] config_value                                   VARCHAR(1024)        null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 3] config_type                                    VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: [string]
[ 4] config_unit                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] config_desc                                    VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 6] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": "TpIxCCbkHnuueUIdFvXOsgiUw",    "config_name": "SQDQawymiLUgtakbZSlsAEAXK",    "config_value": "gIwYEHuISqeKdmcuSEsYCELtw",    "config_type": "jYLweDBNaMMcpudSbPLWVITEm",    "config_unit": "vyCeveEOpiucYwwZIuyRNlrWD",    "config_desc": "iBaRRXydsQYtZlCsKokoxvCcy",    "status": 20}



*/

var (
	System_config_FIELD_NAME_id = "id"

	System_config_FIELD_NAME_config_name = "config_name"

	System_config_FIELD_NAME_config_value = "config_value"

	System_config_FIELD_NAME_config_type = "config_type"

	System_config_FIELD_NAME_config_unit = "config_unit"

	System_config_FIELD_NAME_config_desc = "config_desc"

	System_config_FIELD_NAME_status = "status"
)

// System_config struct is a row record of the o_system_config table in the  database
type System_config struct {
	ID string `json:"id"` //配置ID

	ConfigName string `json:"config_name"` //配置名称

	ConfigValue string `json:"config_value"` //配置值

	ConfigType string `json:"config_type"` //配置类型

	ConfigUnit string `json:"config_unit"` //配置单位

	ConfigDesc string `json:"config_desc"` //配置描述

	Status int32 `json:"status"` //状态

}

var System_configTableInfo = &TableInfo{
	Name: "o_system_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `配置ID`,
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
			Name:               "config_name",
			Comment:            `配置名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ConfigName",
			GoFieldType:        "string",
			JSONFieldName:      "config_name",
			ProtobufFieldName:  "config_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "config_value",
			Comment:            `配置值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "ConfigValue",
			GoFieldType:        "string",
			JSONFieldName:      "config_value",
			ProtobufFieldName:  "config_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "config_type",
			Comment:            `配置类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ConfigType",
			GoFieldType:        "string",
			JSONFieldName:      "config_type",
			ProtobufFieldName:  "config_type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "config_unit",
			Comment:            `配置单位`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ConfigUnit",
			GoFieldType:        "string",
			JSONFieldName:      "config_unit",
			ProtobufFieldName:  "config_unit",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "config_desc",
			Comment:            `配置描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "ConfigDesc",
			GoFieldType:        "string",
			JSONFieldName:      "config_desc",
			ProtobufFieldName:  "config_desc",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `状态`,
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
func (s *System_config) TableName() string {
	return "o_system_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *System_config) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *System_config) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *System_config) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *System_config) TableInfo() *TableInfo {
	return System_configTableInfo
}
