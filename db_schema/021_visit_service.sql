-- +goose Up
-- +goose StatementBegin

-- 病房表
CREATE TABLE o_ward (
    id VARCHAR(32) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_ward IS '病房表';
COMMENT ON COLUMN o_ward.id IS '病房ID';
COMMENT ON COLUMN o_ward.name IS '病房名称';
COMMENT ON COLUMN o_ward.type IS '病房类型';
COMMENT ON COLUMN o_ward.status IS '病房状态';

-- 病床表
CREATE TABLE o_bed (
    id VARCHAR(32) NOT NULL,
    ward_id VARCHAR(32) NOT NULL,
    bed_no VARCHAR(32) NOT NULL,
    type INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_bed IS '病床表';
COMMENT ON COLUMN o_bed.id IS '病床ID';
COMMENT ON COLUMN o_bed.ward_id IS '病房ID';
COMMENT ON COLUMN o_bed.bed_no IS '床位号';
COMMENT ON COLUMN o_bed.type IS '床位类型';
COMMENT ON COLUMN o_bed.status IS '床位状态';

CREATE OR REPLACE VIEW v_bed_info AS
SELECT o.*,w.name AS ward_name FROM o_bed o,o_ward w WHERE o.ward_id = w.id;

COMMENT ON VIEW v_bed_info IS '病床信息视图';
COMMENT ON COLUMN v_bed_info.id IS '病床ID';
COMMENT ON COLUMN v_bed_info.ward_id IS '病房ID';
COMMENT ON COLUMN v_bed_info.bed_no IS '床位号';
COMMENT ON COLUMN v_bed_info.type IS '床位类型';
COMMENT ON COLUMN v_bed_info.status IS '床位状态';
COMMENT ON COLUMN v_bed_info.ward_name IS '病房名称';

-- 病床摄像头绑定关系表
CREATE TABLE o_bed_camera (
    id VARCHAR(32) NOT NULL,
    bed_id VARCHAR(32) NOT NULL,
    camera_id VARCHAR(32) NOT NULL,
    bind_time TIMESTAMP NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_bed_camera IS '病床摄像头绑定关系表';
COMMENT ON COLUMN o_bed_camera.id IS '绑定关系ID';
COMMENT ON COLUMN o_bed_camera.bed_id IS '病床ID';
COMMENT ON COLUMN o_bed_camera.camera_id IS '摄像头ID';
COMMENT ON COLUMN o_bed_camera.bind_time IS '绑定时间';
COMMENT ON COLUMN o_bed_camera.status IS '状态';

-- 病患表
CREATE TABLE o_patient (
    id VARCHAR(32) NOT NULL,
    ward_id VARCHAR(32) NOT NULL,
    bed_id VARCHAR(32) NOT NULL,
    name VARCHAR(255) NOT NULL,
    hospital_no VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    remark VARCHAR(1024) DEFAULT '',
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_patient IS '病患表';
COMMENT ON COLUMN o_patient.id IS '病患ID';
COMMENT ON COLUMN o_patient.ward_id IS '病房ID';
COMMENT ON COLUMN o_patient.bed_id IS '床位ID';
COMMENT ON COLUMN o_patient.name IS '病患姓名';
COMMENT ON COLUMN o_patient.hospital_no IS '住院号';
COMMENT ON COLUMN o_patient.status IS '状态';
COMMENT ON COLUMN o_patient.remark IS '备注';

CREATE OR REPLACE VIEW v_patient_info AS
SELECT p.*,b.bed_no,w.name as ward_name FROM o_patient p,o_bed b,o_ward w WHERE p.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_patient_info IS '病患信息视图';
COMMENT ON COLUMN v_patient_info.id IS '病患ID';
COMMENT ON COLUMN v_patient_info.ward_id IS '病房ID';
COMMENT ON COLUMN v_patient_info.bed_id IS '床位ID';
COMMENT ON COLUMN v_patient_info.name IS '病患姓名';
COMMENT ON COLUMN v_patient_info.hospital_no IS '住院号';
COMMENT ON COLUMN v_patient_info.status IS '状态';
COMMENT ON COLUMN v_patient_info.remark IS '备注';
COMMENT ON COLUMN v_patient_info.bed_no IS '床位号';
COMMENT ON COLUMN v_patient_info.ward_name IS '病房名称';

-- 探视排班表
CREATE TABLE o_visit_schedule (
    id VARCHAR(32) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    total_visitors INTEGER NOT NULL,
    remaining_visitors INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_visit_schedule IS '探视排班表';
COMMENT ON COLUMN o_visit_schedule.id IS '排班ID';
COMMENT ON COLUMN o_visit_schedule.start_time IS '开始时间';
COMMENT ON COLUMN o_visit_schedule.end_time IS '结束时间';
COMMENT ON COLUMN o_visit_schedule.total_visitors IS '探视总人数';
COMMENT ON COLUMN o_visit_schedule.remaining_visitors IS '探视剩余人数';
COMMENT ON COLUMN o_visit_schedule.status IS '状态';

-- 探视排班摄像头表
CREATE TABLE o_schedule_camera (
    id VARCHAR(32) NOT NULL,
    schedule_id VARCHAR(32) NOT NULL,
    camera_id VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_schedule_camera IS '探视排班摄像头表';
COMMENT ON COLUMN o_schedule_camera.id IS '排班摄像头ID';
COMMENT ON COLUMN o_schedule_camera.schedule_id IS '探视排班ID';
COMMENT ON COLUMN o_schedule_camera.camera_id IS '摄像头ID';
COMMENT ON COLUMN o_schedule_camera.status IS '状态';

-- 探视登记表
CREATE TABLE o_visit_record (
    id VARCHAR(32) NOT NULL,
    patient_id VARCHAR(32) NOT NULL,
    relative_id VARCHAR(32) NOT NULL,
    visit_start_time TIMESTAMP NOT NULL,
    visit_end_time TIMESTAMP NOT NULL,
    visitor_name VARCHAR(255) NOT NULL,
    visitor_phone VARCHAR(32) NOT NULL,
    visitor_id_card VARCHAR(32) NOT NULL,
    relationship VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    remark VARCHAR(1024) DEFAULT '',
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_visit_record IS '探视登记表';
COMMENT ON COLUMN o_visit_record.id IS '探视登记ID';
COMMENT ON COLUMN o_visit_record.patient_id IS '病患ID';
COMMENT ON COLUMN o_visit_record.relative_id IS '家属ID';
COMMENT ON COLUMN o_visit_record.visit_start_time IS '探视开始时间';
COMMENT ON COLUMN o_visit_record.visit_end_time IS '探视结束时间';
COMMENT ON COLUMN o_visit_record.visitor_name IS '探视人';
COMMENT ON COLUMN o_visit_record.visitor_phone IS '探视人电话';
COMMENT ON COLUMN o_visit_record.visitor_id_card IS '探视人身份证号';
COMMENT ON COLUMN o_visit_record.relationship IS '探视人关系';
COMMENT ON COLUMN o_visit_record.status IS '状态';
COMMENT ON COLUMN o_visit_record.remark IS '备注';

CREATE OR REPLACE VIEW v_visit_record_info AS
SELECT r.*,p.name AS patient_name,p.name AS patient_ward_name,b.bed_no AS patient_bed_no FROM o_visit_record r,o_patient p,o_bed b,o_ward w WHERE r.patient_id = p.id AND p.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_visit_record_info IS '探视记录信息视图';
COMMENT ON COLUMN v_visit_record_info.id IS '探视记录ID';
COMMENT ON COLUMN v_visit_record_info.patient_id IS '病患ID';
COMMENT ON COLUMN v_visit_record_info.relative_id IS '家属ID';
COMMENT ON COLUMN v_visit_record_info.visit_start_time IS '探视开始时间';
COMMENT ON COLUMN v_visit_record_info.visit_end_time IS '探视结束时间';
COMMENT ON COLUMN v_visit_record_info.visitor_name IS '探视人姓名';
COMMENT ON COLUMN v_visit_record_info.visitor_phone IS '探视人电话';
COMMENT ON COLUMN v_visit_record_info.visitor_id_card IS '探视人身份证号';
COMMENT ON COLUMN v_visit_record_info.relationship IS '探视人关系';
COMMENT ON COLUMN v_visit_record_info.status IS '状态';
COMMENT ON COLUMN v_visit_record_info.remark IS '备注';
COMMENT ON COLUMN v_visit_record_info.patient_name IS '病患姓名';
COMMENT ON COLUMN v_visit_record_info.patient_ward_name IS '病房名称';
COMMENT ON COLUMN v_visit_record_info.patient_bed_no IS '床位号';

-- 摄像头表
CREATE TABLE o_camera (
    id VARCHAR(32) NOT NULL,
    device_name VARCHAR(255) NOT NULL,
    device_no VARCHAR(32) NOT NULL,
    location_type INTEGER NOT NULL DEFAULT 0,
    ward_id VARCHAR(32),
    bed_id VARCHAR(32),
    device_type INTEGER NOT NULL DEFAULT 0  ,
    manufacturer VARCHAR(255) NOT NULL,
    main_stream_url VARCHAR(1024) NOT NULL,
    sub_stream_url VARCHAR(1024) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_camera IS '摄像头表';
COMMENT ON COLUMN o_camera.id IS '摄像头ID';
COMMENT ON COLUMN o_camera.device_name IS '设备名称';
COMMENT ON COLUMN o_camera.device_no IS '设备编号';
COMMENT ON COLUMN o_camera.location_type IS '位置类型';
COMMENT ON COLUMN o_camera.ward_id IS '所属病房';
COMMENT ON COLUMN o_camera.bed_id IS '所属床位';
COMMENT ON COLUMN o_camera.device_type IS '设备类型';
COMMENT ON COLUMN o_camera.manufacturer IS '设备厂商';
COMMENT ON COLUMN o_camera.main_stream_url IS '主码流URL';
COMMENT ON COLUMN o_camera.sub_stream_url IS '辅码流URL';
COMMENT ON COLUMN o_camera.status IS '状态';

CREATE OR REPLACE VIEW v_camera_info AS
SELECT c.*,b.bed_no,w.name as ward_name FROM o_camera c,o_bed b,o_ward w WHERE c.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_camera_info IS '摄像头信息视图';
COMMENT ON COLUMN v_camera_info.id IS '摄像头ID';
COMMENT ON COLUMN v_camera_info.device_name IS '设备名称';
COMMENT ON COLUMN v_camera_info.device_no IS '设备编号';
COMMENT ON COLUMN v_camera_info.location_type IS '位置类型';
COMMENT ON COLUMN v_camera_info.ward_id IS '所属病房ID';
COMMENT ON COLUMN v_camera_info.bed_id IS '所属床位ID';
COMMENT ON COLUMN v_camera_info.device_type IS '设备类型';
COMMENT ON COLUMN v_camera_info.manufacturer IS '设备厂商';
COMMENT ON COLUMN v_camera_info.main_stream_url IS '主码流URL';
COMMENT ON COLUMN v_camera_info.sub_stream_url IS '辅码流URL';
COMMENT ON COLUMN v_camera_info.status IS '状态';
COMMENT ON COLUMN v_camera_info.bed_no IS '床位号';
COMMENT ON COLUMN v_camera_info.ward_name IS '病房名称';

-- 移动摄像头绑定关系表
CREATE TABLE o_mobile_camera_binding (
    id VARCHAR(32) NOT NULL,
    camera_id VARCHAR(32) NOT NULL,
    panoramic_camera_id VARCHAR(32) NOT NULL,
    bind_time TIMESTAMP NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_mobile_camera_binding IS '移动摄像头绑定关系表';
COMMENT ON COLUMN o_mobile_camera_binding.id IS '绑定关系ID';
COMMENT ON COLUMN o_mobile_camera_binding.camera_id IS '摄像头ID';
COMMENT ON COLUMN o_mobile_camera_binding.panoramic_camera_id IS '全景摄像头ID';
COMMENT ON COLUMN o_mobile_camera_binding.bind_time IS '绑定时间';
COMMENT ON COLUMN o_mobile_camera_binding.status IS '状态';

-- 头显表
CREATE TABLE o_head_display (
    id VARCHAR(32) NOT NULL,
    device_name VARCHAR(255) NOT NULL,
    device_no VARCHAR(32) NOT NULL,
    model VARCHAR(255) NOT NULL,
    ward_id VARCHAR(32) NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_head_display IS '头显表';
COMMENT ON COLUMN o_head_display.id IS '头显ID';
COMMENT ON COLUMN o_head_display.device_name IS '设备名称';
COMMENT ON COLUMN o_head_display.device_no IS '设备编号';
COMMENT ON COLUMN o_head_display.model IS '型号';
COMMENT ON COLUMN o_head_display.ward_id IS '病房';

CREATE OR REPLACE VIEW v_head_display_info AS
SELECT h.*,w.name AS ward_name FROM o_head_display h,o_ward w WHERE h.ward_id = w.id;

COMMENT ON VIEW v_head_display_info IS '头显信息视图';
COMMENT ON COLUMN v_head_display_info.id IS '头显ID';
COMMENT ON COLUMN v_head_display_info.device_name IS '设备名称';
COMMENT ON COLUMN v_head_display_info.device_no IS '设备编号';
COMMENT ON COLUMN v_head_display_info.model IS '型号';
COMMENT ON COLUMN v_head_display_info.ward_id IS '病房ID';
COMMENT ON COLUMN v_head_display_info.ward_name IS '病房名称';

-- 直播记录表
CREATE TABLE o_live_record (
    id VARCHAR(32) NOT NULL,
    schedule_id VARCHAR(32) NOT NULL,
    patient_id VARCHAR(32) NOT NULL,
    relative_id VARCHAR(32) NOT NULL,
    device_id VARCHAR(32) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    file_size BIGINT,
    stream_id VARCHAR(32),
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_live_record IS '直播记录表';
COMMENT ON COLUMN o_live_record.id IS '直播记录ID';
COMMENT ON COLUMN o_live_record.schedule_id IS '探视排班ID';
COMMENT ON COLUMN o_live_record.patient_id IS '病患ID';
COMMENT ON COLUMN o_live_record.relative_id IS '家属ID';
COMMENT ON COLUMN o_live_record.device_id IS '设备ID';
COMMENT ON COLUMN o_live_record.start_time IS '直播开始时间';
COMMENT ON COLUMN o_live_record.end_time IS '直播结束时间';
COMMENT ON COLUMN o_live_record.file_size IS '文件大小';
COMMENT ON COLUMN o_live_record.stream_id IS '流ID';
COMMENT ON COLUMN o_live_record.status IS '状态(0:未开始,1:直播中,2:已结束)';

CREATE OR REPLACE VIEW v_live_record_info AS
SELECT l.*,p.name AS patient_name,p.name AS patient_ward_name,b.bed_no AS patient_bed_no FROM o_live_record l,o_patient p,o_bed b,o_ward w WHERE l.patient_id = p.id AND p.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_live_record_info IS '直播记录信息视图';
COMMENT ON COLUMN v_live_record_info.id IS '直播记录ID';
COMMENT ON COLUMN v_live_record_info.schedule_id IS '探视排班ID';
COMMENT ON COLUMN v_live_record_info.patient_id IS '病患ID';
COMMENT ON COLUMN v_live_record_info.relative_id IS '家属ID';
COMMENT ON COLUMN v_live_record_info.device_id IS '设备ID';
COMMENT ON COLUMN v_live_record_info.start_time IS '直播开始时间';
COMMENT ON COLUMN v_live_record_info.end_time IS '直播结束时间';
COMMENT ON COLUMN v_live_record_info.file_size IS '文件大小';
COMMENT ON COLUMN v_live_record_info.stream_id IS '流ID';
COMMENT ON COLUMN v_live_record_info.status IS '状态(0:未开始,1:直播中,2:已结束)';
COMMENT ON COLUMN v_live_record_info.patient_name IS '病患姓名';
COMMENT ON COLUMN v_live_record_info.patient_ward_name IS '病房名称';
COMMENT ON COLUMN v_live_record_info.patient_bed_no IS '床位号';

-- 系统配置表
CREATE TABLE o_system_config (
    id VARCHAR(32) NOT NULL,
    config_name VARCHAR(255) NOT NULL,
    config_value VARCHAR(1024) NOT NULL,
    config_type VARCHAR(32) NOT NULL DEFAULT 'string',
    config_unit VARCHAR(32) DEFAULT '',
    config_desc VARCHAR(1024) DEFAULT '',
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_system_config IS '系统配置表';
COMMENT ON COLUMN o_system_config.id IS '配置ID';
COMMENT ON COLUMN o_system_config.config_name IS '配置名称';
COMMENT ON COLUMN o_system_config.config_value IS '配置值';
COMMENT ON COLUMN o_system_config.config_type IS '配置类型';
COMMENT ON COLUMN o_system_config.config_unit IS '配置单位';
COMMENT ON COLUMN o_system_config.config_desc IS '配置描述';
COMMENT ON COLUMN o_system_config.status IS '状态';

CREATE OR REPLACE VIEW v_family_member_info AS
SELECT * FROM o_user WHERE type = 3;

COMMENT ON VIEW v_family_member_info IS '家属信息视图';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS v_family_member_info CASCADE;
DROP VIEW IF EXISTS v_bed_info CASCADE;
DROP VIEW IF EXISTS v_patient_info CASCADE;
DROP VIEW IF EXISTS v_visit_record CASCADE;
DROP VIEW IF EXISTS v_camera_info CASCADE;
DROP VIEW IF EXISTS v_live_record_info CASCADE;
DROP VIEW IF EXISTS v_head_display_info CASCADE;
DROP TABLE IF EXISTS o_system_config CASCADE;
DROP TABLE IF EXISTS o_live_record CASCADE;
DROP TABLE IF EXISTS o_head_display CASCADE;
DROP TABLE IF EXISTS o_mobile_camera_binding CASCADE;
DROP TABLE IF EXISTS o_camera CASCADE;
DROP TABLE IF EXISTS o_visit_record CASCADE;
DROP TABLE IF EXISTS o_schedule_camera CASCADE;
DROP TABLE IF EXISTS o_visit_schedule CASCADE;
DROP TABLE IF EXISTS o_patient CASCADE;
DROP TABLE IF EXISTS o_bed_camera CASCADE;
DROP TABLE IF EXISTS o_bed CASCADE;
DROP TABLE IF EXISTS o_ward CASCADE;

-- +goose StatementEnd
