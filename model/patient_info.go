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


Table: v_patient_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] ward_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] bed_id                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] hospital_no                                    VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] remark                                         VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 7] bed_no                                         VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] camera_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] vr_camera_id                                   VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] ward_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "lOhBFTbcATKUSGROseeICmZSw",    "ward_id": "KZkgbuvvDvPNSYMrHYfHCckWv",    "bed_id": "juwWNpkYhRBUBySkVwLhGsAwM",    "name": "hZXlaqrnHtBHwMchCPBoXEUlh",    "hospital_no": "jWthbfNRDTDpTHYYtCdfUiwgG",    "status": 68,    "remark": "LCuLkBJpAlyDIJHVlsoWssKFE",    "bed_no": "GGasabyDumiknAviZYSHicKpi",    "camera_id": "KFICPBwPLfokNoOCrcBpCjmyI",    "vr_camera_id": "JdRrZwAaPKCFhwvHmmFCMXwNS",    "ward_name": "qleAXblFnoVUdisfIXvJQgIMg"}


Comments
-------------------------------------
[ 0] Warning table: v_patient_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_patient_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Patient_info_FIELD_NAME_id = "id"

	Patient_info_FIELD_NAME_ward_id = "ward_id"

	Patient_info_FIELD_NAME_bed_id = "bed_id"

	Patient_info_FIELD_NAME_name = "name"

	Patient_info_FIELD_NAME_hospital_no = "hospital_no"

	Patient_info_FIELD_NAME_status = "status"

	Patient_info_FIELD_NAME_remark = "remark"

	Patient_info_FIELD_NAME_bed_no = "bed_no"

	Patient_info_FIELD_NAME_camera_id = "camera_id"

	Patient_info_FIELD_NAME_vr_camera_id = "vr_camera_id"

	Patient_info_FIELD_NAME_ward_name = "ward_name"
)

// Patient_info struct is a row record of the v_patient_info table in the  database
type Patient_info struct {
	ID string `json:"id"` //病患ID

	WardID string `json:"ward_id"` //病房ID

	BedID string `json:"bed_id"` //床位ID

	Name string `json:"name"` //病患姓名

	HospitalNo string `json:"hospital_no"` //住院号

	Status int32 `json:"status"` //状态(0:正常,1:出院)

	Remark string `json:"remark"` //备注

	BedNo string `json:"bed_no"` //床位号

	CameraID string `json:"camera_id"` //床头摄像头ID

	VrCameraID string `json:"vr_camera_id"` //VR摄像头ID

	WardName string `json:"ward_name"` //病房名称

}

var Patient_infoTableInfo = &TableInfo{
	Name: "v_patient_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `病患ID`,
			Notes: `Warning table: v_patient_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_patient_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "bed_id",
			Comment:            `床位ID`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `状态(0:正常,1:出院)`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "remark",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "camera_id",
			Comment:            `床头摄像头ID`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "vr_camera_id",
			Comment:            `VR摄像头ID`,
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
	},
}

// TableName sets the insert table name for this struct type
func (p *Patient_info) TableName() string {
	return "v_patient_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Patient_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Patient_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Patient_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Patient_info) TableInfo() *TableInfo {
	return Patient_infoTableInfo
}
