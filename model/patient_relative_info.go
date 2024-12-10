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


Table: v_patient_relative_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] patient_id                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] relative_id                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] relationship                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] create_time                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] check_status                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] patient_name                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] hospital_no                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] patient_status                                 INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[10] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] bed_no                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] relative_name                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] relative_mobile                                VARCHAR(15)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[14] relative_id_card                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "YFIhgxFuoccaagZWHGiuZkCje",    "patient_id": "SZGMlpZqGgjeSYKHdGFbTfTTo",    "relative_id": "SgklwXfSXuNcqZltFWgYXDEpO",    "relationship": "MWrIZyvasxLpYZXgEXBucLNHu",    "status": 12,    "create_time": 28,    "check_status": 48,    "patient_name": "FcqaExnRshUiZKRyGJYqbMMbj",    "hospital_no": "UcMHmGoUZZesIRuRtAZkbEKQE",    "patient_status": 31,    "ward_name": "GTCeqTNSXdQxbXFAAdCVDRXQe",    "bed_no": "RvnlSjIBouPPWFLZtqVEsMxkk",    "relative_name": "UNqUdDVSJBPgSoNxWLJGGJLnf",    "relative_mobile": "WjkxxmsabwcTueGcZoKIpIHaV",    "relative_id_card": "yjWFUOvyPobnXaPVFdiBJiLrh"}


Comments
-------------------------------------
[ 0] Warning table: v_patient_relative_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_patient_relative_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Patient_relative_info_FIELD_NAME_id = "id"

	Patient_relative_info_FIELD_NAME_patient_id = "patient_id"

	Patient_relative_info_FIELD_NAME_relative_id = "relative_id"

	Patient_relative_info_FIELD_NAME_relationship = "relationship"

	Patient_relative_info_FIELD_NAME_status = "status"

	Patient_relative_info_FIELD_NAME_create_time = "create_time"

	Patient_relative_info_FIELD_NAME_check_status = "check_status"

	Patient_relative_info_FIELD_NAME_patient_name = "patient_name"

	Patient_relative_info_FIELD_NAME_hospital_no = "hospital_no"

	Patient_relative_info_FIELD_NAME_patient_status = "patient_status"

	Patient_relative_info_FIELD_NAME_ward_name = "ward_name"

	Patient_relative_info_FIELD_NAME_bed_no = "bed_no"

	Patient_relative_info_FIELD_NAME_relative_name = "relative_name"

	Patient_relative_info_FIELD_NAME_relative_mobile = "relative_mobile"

	Patient_relative_info_FIELD_NAME_relative_id_card = "relative_id_card"
)

// Patient_relative_info struct is a row record of the v_patient_relative_info table in the  database
type Patient_relative_info struct {
	ID string `json:"id"` //关联ID

	PatientID string `json:"patient_id"` //病患ID

	RelativeID string `json:"relative_id"` //家属ID

	Relationship string `json:"relationship"` //与患者关系

	Status int32 `json:"status"` //状态

	CreateTime common.LocalTime `json:"create_time"` //创建时间

	CheckStatus int32 `json:"check_status"` //审核状态(0:未审核,1:已审核,2:审核不通过)

	PatientName string `json:"patient_name"` //病患姓名

	HospitalNo string `json:"hospital_no"` //住院号

	PatientStatus int32 `json:"patient_status"` //病患状态

	WardName string `json:"ward_name"` //病房名称

	BedNo string `json:"bed_no"` //床位号

	RelativeName string `json:"relative_name"` //家属姓名

	RelativeMobile string `json:"relative_mobile"` //家属手机号

	RelativeIDCard string `json:"relative_id_card"` //家属身份证号

}

var Patient_relative_infoTableInfo = &TableInfo{
	Name: "v_patient_relative_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `关联ID`,
			Notes: `Warning table: v_patient_relative_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_patient_relative_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "relationship",
			Comment:            `与患者关系`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "status",
			Comment:            `状态`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "create_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "hospital_no",
			Comment:            `住院号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "HospitalNo",
			GoFieldType:        "string",
			JSONFieldName:      "hospital_no",
			ProtobufFieldName:  "hospital_no",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "patient_status",
			Comment:            `病患状态`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "PatientStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "patient_status",
			ProtobufFieldName:  "patient_status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "relative_name",
			Comment:            `家属姓名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RelativeName",
			GoFieldType:        "string",
			JSONFieldName:      "relative_name",
			ProtobufFieldName:  "relative_name",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "relative_mobile",
			Comment:            `家属手机号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       15,
			GoFieldName:        "RelativeMobile",
			GoFieldType:        "string",
			JSONFieldName:      "relative_mobile",
			ProtobufFieldName:  "relative_mobile",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "relative_id_card",
			Comment:            `家属身份证号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RelativeIDCard",
			GoFieldType:        "string",
			JSONFieldName:      "relative_id_card",
			ProtobufFieldName:  "relative_id_card",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Patient_relative_info) TableName() string {
	return "v_patient_relative_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Patient_relative_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Patient_relative_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Patient_relative_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Patient_relative_info) TableInfo() *TableInfo {
	return Patient_relative_infoTableInfo
}
