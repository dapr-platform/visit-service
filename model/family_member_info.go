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


Table: v_family_member_info
[ 0] id                                             VARCHAR(25)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 1] tenant_id                                      VARCHAR(25)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 25      default: []
[ 2] mobile                                         VARCHAR(15)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[ 3] email                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] identity                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] gender                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] address                                        VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 8] password                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[10] org_id                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] id_card                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] work_number                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] avatar_url                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] create_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[15] update_at                                      TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[16] remark                                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[17] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "AHqwhETrfhAWMKIybKLxXCHPa",    "tenant_id": "mOPvRTCOoLNxaWPUWkGPeMoJC",    "mobile": "hvmurCbDeOCcmMRQmsTannlno",    "email": "sOLqZIxAVVqXKIWcgwNXYJfqB",    "identity": "kHEMYEFahMipKXbSWtWxmIjRF",    "name": "qFMhiHimRwtPyYChttTktWIsP",    "gender": 74,    "address": "qZnDIjhZPwCqKWTISKRfDaAkn",    "password": "IqgwnFSRmUOiUokJrOLthuypA",    "type": 28,    "org_id": "WuvHsMpmgnAdPUqZZKkKGGAmL",    "id_card": "hdplpxkqtBuKqEOmfOwKQavwR",    "work_number": "CGBBDBZJUACqPcJLViPfcBiwq",    "avatar_url": "pNEoWjYKOTEpcINQVgVAMSVax",    "create_at": 94,    "update_at": 91,    "remark": "nrsElPeyJcTreWtrVOaekcnRS",    "status": 19}


Comments
-------------------------------------
[ 0] Warning table: v_family_member_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_family_member_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Family_member_info_FIELD_NAME_id = "id"

	Family_member_info_FIELD_NAME_tenant_id = "tenant_id"

	Family_member_info_FIELD_NAME_mobile = "mobile"

	Family_member_info_FIELD_NAME_email = "email"

	Family_member_info_FIELD_NAME_identity = "identity"

	Family_member_info_FIELD_NAME_name = "name"

	Family_member_info_FIELD_NAME_gender = "gender"

	Family_member_info_FIELD_NAME_address = "address"

	Family_member_info_FIELD_NAME_password = "password"

	Family_member_info_FIELD_NAME_type = "type"

	Family_member_info_FIELD_NAME_org_id = "org_id"

	Family_member_info_FIELD_NAME_id_card = "id_card"

	Family_member_info_FIELD_NAME_work_number = "work_number"

	Family_member_info_FIELD_NAME_avatar_url = "avatar_url"

	Family_member_info_FIELD_NAME_create_at = "create_at"

	Family_member_info_FIELD_NAME_update_at = "update_at"

	Family_member_info_FIELD_NAME_remark = "remark"

	Family_member_info_FIELD_NAME_status = "status"
)

// Family_member_info struct is a row record of the v_family_member_info table in the  database
type Family_member_info struct {
	ID string `json:"id"` //id

	TenantID string `json:"tenant_id"` //tenant_id

	Mobile string `json:"mobile"` //mobile

	Email string `json:"email"` //email

	Identity string `json:"identity"` //identity

	Name string `json:"name"` //name

	Gender int32 `json:"gender"` //gender

	Address string `json:"address"` //address

	Password string `json:"password"` //password

	Type int32 `json:"type"` //type

	OrgID string `json:"org_id"` //org_id

	IDCard string `json:"id_card"` //id_card

	WorkNumber string `json:"work_number"` //work_number

	AvatarURL string `json:"avatar_url"` //avatar_url

	CreateAt common.LocalTime `json:"create_at"` //create_at

	UpdateAt common.LocalTime `json:"update_at"` //update_at

	Remark string `json:"remark"` //remark

	Status int32 `json:"status"` //status

}

var Family_member_infoTableInfo = &TableInfo{
	Name: "v_family_member_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_family_member_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_family_member_info primary key column id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "tenant_id",
			Comment:            `tenant_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(25)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       25,
			GoFieldName:        "TenantID",
			GoFieldType:        "string",
			JSONFieldName:      "tenant_id",
			ProtobufFieldName:  "tenant_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mobile",
			Comment:            `mobile`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       15,
			GoFieldName:        "Mobile",
			GoFieldType:        "string",
			JSONFieldName:      "mobile",
			ProtobufFieldName:  "mobile",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "email",
			Comment:            `email`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "identity",
			Comment:            `identity`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Identity",
			GoFieldType:        "string",
			JSONFieldName:      "identity",
			ProtobufFieldName:  "identity",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "name",
			Comment:            `name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "gender",
			Comment:            `gender`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "address",
			Comment:            `address`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "password",
			Comment:            `password`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "type",
			Comment:            `type`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "org_id",
			Comment:            `org_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "OrgID",
			GoFieldType:        "string",
			JSONFieldName:      "org_id",
			ProtobufFieldName:  "org_id",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "id_card",
			Comment:            `id_card`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "IDCard",
			GoFieldType:        "string",
			JSONFieldName:      "id_card",
			ProtobufFieldName:  "id_card",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "work_number",
			Comment:            `work_number`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "WorkNumber",
			GoFieldType:        "string",
			JSONFieldName:      "work_number",
			ProtobufFieldName:  "work_number",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "avatar_url",
			Comment:            `avatar_url`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "AvatarURL",
			GoFieldType:        "string",
			JSONFieldName:      "avatar_url",
			ProtobufFieldName:  "avatar_url",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "create_at",
			Comment:            `create_at`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "create_at",
			ProtobufFieldName:  "create_at",
			ProtobufType:       "uint64",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "update_at",
			Comment:            `update_at`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdateAt",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "update_at",
			ProtobufFieldName:  "update_at",
			ProtobufType:       "uint64",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "remark",
			Comment:            `remark`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *Family_member_info) TableName() string {
	return "v_family_member_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *Family_member_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *Family_member_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *Family_member_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *Family_member_info) TableInfo() *TableInfo {
	return Family_member_infoTableInfo
}
