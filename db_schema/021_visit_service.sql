-- +goose Up
-- +goose StatementBegin


-- 摄像头表
CREATE TABLE o_camera (
    id VARCHAR(32) NOT NULL,
    device_name VARCHAR(255) NOT NULL,
    device_no VARCHAR(32) NOT NULL,
    location_type INTEGER NOT NULL DEFAULT 0,
    device_type INTEGER NOT NULL DEFAULT 0  ,
    manufacturer VARCHAR(255) NOT NULL DEFAULT '',
    main_stream_url VARCHAR(1024) NOT NULL DEFAULT '',
    sub_stream_url VARCHAR(1024) NOT NULL DEFAULT '',
    rel_vr_camera_id VARCHAR(32) NOT NULL DEFAULT '',
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_camera IS '摄像头表';
COMMENT ON COLUMN o_camera.id IS '摄像头ID';
COMMENT ON COLUMN o_camera.device_name IS '设备名称';
COMMENT ON COLUMN o_camera.device_no IS '设备编号';
COMMENT ON COLUMN o_camera.location_type IS '位置类型';
COMMENT ON COLUMN o_camera.device_type IS '设备类型';
COMMENT ON COLUMN o_camera.manufacturer IS '设备厂商';
COMMENT ON COLUMN o_camera.main_stream_url IS '主码流URL';
COMMENT ON COLUMN o_camera.sub_stream_url IS '辅码流URL';
COMMENT ON COLUMN o_camera.rel_vr_camera_id IS '关联VR摄像头ID(当位置为可移动时，需关联VR摄像头)';
COMMENT ON COLUMN o_camera.status IS '状态(0:正常,1:禁用)';

CREATE OR REPLACE VIEW v_camera_info AS
SELECT * FROM o_camera;

COMMENT ON VIEW v_camera_info IS '摄像头信息视图';
COMMENT ON COLUMN v_camera_info.id IS '摄像头ID';
COMMENT ON COLUMN v_camera_info.device_name IS '设备名称';
COMMENT ON COLUMN v_camera_info.device_no IS '设备编号';
COMMENT ON COLUMN v_camera_info.location_type IS '位置类型(0:固定,1:可移动)';
COMMENT ON COLUMN v_camera_info.device_type IS '设备类型(0:普通,1:普通云台,2:VR)';
COMMENT ON COLUMN v_camera_info.manufacturer IS '设备厂商';
COMMENT ON COLUMN v_camera_info.main_stream_url IS '主码流URL';
COMMENT ON COLUMN v_camera_info.sub_stream_url IS '辅码流URL';
COMMENT ON COLUMN v_camera_info.rel_vr_camera_id IS '关联VR摄像头ID(当位置为可移动时，需关联VR摄像头)';
COMMENT ON COLUMN v_camera_info.status IS '状态(0:正常,1:禁用)';

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
COMMENT ON COLUMN o_ward.type IS '病房类型(0:普通,1:ICU)';
COMMENT ON COLUMN o_ward.status IS '病房状态(0:空闲,1:占用)';

-- 病床表
CREATE TABLE o_bed (
    id VARCHAR(32) NOT NULL,
    ward_id VARCHAR(32) NOT NULL,
    bed_no VARCHAR(32) NOT NULL,
    camera_id VARCHAR(32) NOT NULL DEFAULT '',
    vr_camera_id VARCHAR(32) NOT NULL DEFAULT '',
    type INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
COMMENT ON TABLE o_bed IS '病床表';
COMMENT ON COLUMN o_bed.id IS '病床ID';
COMMENT ON COLUMN o_bed.ward_id IS '病房ID';
COMMENT ON COLUMN o_bed.bed_no IS '床位号';
COMMENT ON COLUMN o_bed.camera_id IS '摄像头ID(可关联固定床头摄像头)';
COMMENT ON COLUMN o_bed.vr_camera_id IS 'VR摄像头ID(可关联固定VR摄像头)';
COMMENT ON COLUMN o_bed.type IS '床位类型(0:普通,1:ICU)';
COMMENT ON COLUMN o_bed.status IS '床位状态(0:空闲,1:占用)';

CREATE OR REPLACE VIEW v_bed_info AS
SELECT o.*,w.name AS ward_name,(select name from o_camera where id = o.camera_id ) as camera_name,(select name from o_camera where id = o.vr_camera_id) as vr_camera_name FROM o_bed o,o_ward w WHERE o.ward_id = w.id;

COMMENT ON VIEW v_bed_info IS '病床信息视图';
COMMENT ON COLUMN v_bed_info.id IS '病床ID';
COMMENT ON COLUMN v_bed_info.ward_id IS '病房ID';
COMMENT ON COLUMN v_bed_info.bed_no IS '床位号';
COMMENT ON COLUMN v_bed_info.type IS '床位类型(0:普通,1:ICU)';
COMMENT ON COLUMN v_bed_info.status IS '床位状态(0:空闲,1:占用)';
COMMENT ON COLUMN v_bed_info.ward_name IS '病房名称';
COMMENT ON COLUMN v_bed_info.camera_name IS '床头摄像头名称';
COMMENT ON COLUMN v_bed_info.vr_camera_name IS 'VR摄像头名称';


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
COMMENT ON COLUMN o_patient.status IS '状态(0:正常,1:出院)';
COMMENT ON COLUMN o_patient.remark IS '备注';

CREATE OR REPLACE VIEW v_patient_info AS
SELECT p.*,b.bed_no,b.camera_id,b.vr_camera_id,w.name as ward_name FROM o_patient p,o_bed b,o_ward w WHERE p.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_patient_info IS '病患信息视图';
COMMENT ON COLUMN v_patient_info.id IS '病患ID';
COMMENT ON COLUMN v_patient_info.ward_id IS '病房ID';
COMMENT ON COLUMN v_patient_info.bed_id IS '床位ID';
COMMENT ON COLUMN v_patient_info.name IS '病患姓名';
COMMENT ON COLUMN v_patient_info.hospital_no IS '住院号';
COMMENT ON COLUMN v_patient_info.status IS '状态(0:正常,1:出院)';
COMMENT ON COLUMN v_patient_info.remark IS '备注';
COMMENT ON COLUMN v_patient_info.bed_no IS '床位号';
COMMENT ON COLUMN v_patient_info.camera_id IS '床头摄像头ID';
COMMENT ON COLUMN v_patient_info.vr_camera_id IS 'VR摄像头ID';
COMMENT ON COLUMN v_patient_info.ward_name IS '病房名称';

-- 探视排班表
CREATE TABLE o_visit_schedule (
    id VARCHAR(32) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    total_visitors INTEGER NOT NULL,
    schedule_visitors INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);
CREATE INDEX idx_visit_schedule_start_time ON o_visit_schedule (start_time);
COMMENT ON TABLE o_visit_schedule IS '探视排班表';
COMMENT ON COLUMN o_visit_schedule.id IS '排班ID';
COMMENT ON COLUMN o_visit_schedule.start_time IS '开始时间';
COMMENT ON COLUMN o_visit_schedule.end_time IS '结束时间';
COMMENT ON COLUMN o_visit_schedule.total_visitors IS '探视总人数';
COMMENT ON COLUMN o_visit_schedule.schedule_visitors IS '已预约人数';
COMMENT ON COLUMN o_visit_schedule.status IS '状态(0:可预约,1:不可预约)';


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
    stream_url_suffix VARCHAR(1024),
    camera_id VARCHAR(32),
    vr_camera_id VARCHAR(32),
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
COMMENT ON COLUMN o_live_record.stream_url_suffix IS '流URL后缀';
COMMENT ON COLUMN o_live_record.camera_id IS '床头摄像头ID';
COMMENT ON COLUMN o_live_record.vr_camera_id IS 'VR摄像头ID';
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
    camera_id VARCHAR(32) NOT NULL,
    vr_camera_id VARCHAR(32) NOT NULL,
    check_status INTEGER NOT NULL DEFAULT 0,
    send_sms_status INTEGER NOT NULL DEFAULT 0,
    send_prompt_sms_status INTEGER NOT NULL DEFAULT 0,
    status INTEGER NOT NULL DEFAULT 0,
    remark VARCHAR(1024) DEFAULT '',
    PRIMARY KEY (id)
);
CREATE INDEX idx_visit_record_visit_start_time ON o_visit_record (visit_start_time);
COMMENT ON TABLE o_visit_record IS '探视登记表';
COMMENT ON COLUMN o_visit_record.id IS '探视登记ID';
COMMENT ON COLUMN o_visit_record.patient_id IS '病患ID';
COMMENT ON COLUMN o_visit_record.relative_id IS '家属ID';
COMMENT ON COLUMN o_visit_record.visit_start_time IS '探视开始时间';
COMMENT ON COLUMN o_visit_record.visit_end_time IS '探视结束时间';
COMMENT ON COLUMN o_visit_record.visitor_name IS '探视人';
COMMENT ON COLUMN o_visit_record.visitor_phone IS '探视人电话';
COMMENT ON COLUMN o_visit_record.visitor_id_card IS '探视人身份证号';
COMMENT ON COLUMN o_visit_record.relationship IS '探视人与患者关系(父母，配偶，子女，其他)';
COMMENT ON COLUMN o_visit_record.camera_id IS '床头摄像头ID';
COMMENT ON COLUMN o_visit_record.vr_camera_id IS 'VR摄像头ID';
COMMENT ON COLUMN o_visit_record.check_status IS '审核状态(0:未审核,1:已审核,2:审核不通过)';
COMMENT ON COLUMN o_visit_record.send_sms_status IS '发送审核短信状态(0:未发送,1:已发送)';
COMMENT ON COLUMN o_visit_record.send_prompt_sms_status IS '发送提醒短信状态(0:未发送,1:已发送)';
COMMENT ON COLUMN o_visit_record.status IS '状态(0:正常,1:取消)';
COMMENT ON COLUMN o_visit_record.remark IS '备注';

CREATE OR REPLACE VIEW v_visit_record_info AS
SELECT r.*,p.name AS patient_name,w.name AS patient_ward_name,b.bed_no AS patient_bed_no,
(SELECT l.stream_id FROM o_live_record l WHERE l.schedule_id = r.id LIMIT 1) as stream_id,
(SELECT l.stream_url_suffix FROM o_live_record l WHERE l.schedule_id = r.id LIMIT 1) as stream_url_suffix
FROM o_visit_record r,o_patient p,o_bed b,o_ward w 
WHERE r.patient_id = p.id AND p.bed_id = b.id AND b.ward_id = w.id;

COMMENT ON VIEW v_visit_record_info IS '探视记录信息视图';
COMMENT ON COLUMN v_visit_record_info.id IS '探视记录ID';
COMMENT ON COLUMN v_visit_record_info.patient_id IS '病患ID';
COMMENT ON COLUMN v_visit_record_info.relative_id IS '家属ID';
COMMENT ON COLUMN v_visit_record_info.visit_start_time IS '探视开始时间';
COMMENT ON COLUMN v_visit_record_info.visit_end_time IS '探视结束时间';
COMMENT ON COLUMN v_visit_record_info.visitor_name IS '探视人姓名';
COMMENT ON COLUMN v_visit_record_info.visitor_phone IS '探视人电话';
COMMENT ON COLUMN v_visit_record_info.visitor_id_card IS '探视人身份证号';
COMMENT ON COLUMN v_visit_record_info.relationship IS '探视人与患者关系(父母，配偶，子女，其他)';
COMMENT ON COLUMN v_visit_record_info.camera_id IS '床头摄像头ID';
COMMENT ON COLUMN v_visit_record_info.vr_camera_id IS 'VR摄像头ID';
COMMENT ON COLUMN v_visit_record_info.check_status IS '审核状态(0:未审核,1:已审核,2:审核不通过)';
COMMENT ON COLUMN v_visit_record_info.status IS '状态(0:正常,1:取消)';
COMMENT ON COLUMN v_visit_record_info.remark IS '备注';
COMMENT ON COLUMN v_visit_record_info.patient_name IS '病患姓名';
COMMENT ON COLUMN v_visit_record_info.patient_ward_name IS '病房名称';
COMMENT ON COLUMN v_visit_record_info.patient_bed_no IS '床位号';
COMMENT ON COLUMN v_visit_record_info.stream_id IS '直播流ID';


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


-- 病患-家属关联表
CREATE TABLE o_patient_relative (
    id VARCHAR(32) NOT NULL,
    patient_id VARCHAR(32) NOT NULL,
    relative_id VARCHAR(32) NOT NULL,
    relationship VARCHAR(32) NOT NULL,
    status INTEGER NOT NULL DEFAULT 0,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_patient_relative_patient_id ON o_patient_relative (patient_id);
CREATE INDEX idx_patient_relative_relative_id ON o_patient_relative (relative_id);

COMMENT ON TABLE o_patient_relative IS '病患-家属关联表';
COMMENT ON COLUMN o_patient_relative.id IS '关联ID';
COMMENT ON COLUMN o_patient_relative.patient_id IS '病患ID';
COMMENT ON COLUMN o_patient_relative.relative_id IS '家属ID';
COMMENT ON COLUMN o_patient_relative.relationship IS '与患者关系(父母,配偶,子女,其他)';
COMMENT ON COLUMN o_patient_relative.status IS '状态(0:正常,1:解除关联)';
COMMENT ON COLUMN o_patient_relative.create_time IS '创建时间';

CREATE OR REPLACE VIEW v_patient_relative_info AS
SELECT 
    pr.*,
    p.name AS patient_name,
    p.hospital_no,
    p.status AS patient_status,
    w.name AS ward_name,
    b.bed_no,
    u.zh_name AS relative_name,
    u.mobile AS relative_mobile,
    u.id_card AS relative_id_card
FROM o_patient_relative pr
LEFT JOIN o_patient p ON pr.patient_id = p.id
LEFT JOIN o_bed b ON p.bed_id = b.id 
LEFT JOIN o_ward w ON p.ward_id = w.id
LEFT JOIN o_user u ON pr.relative_id = u.id;

COMMENT ON VIEW v_patient_relative_info IS '病患-家属关联信息视图';
COMMENT ON COLUMN v_patient_relative_info.id IS '关联ID';
COMMENT ON COLUMN v_patient_relative_info.patient_id IS '病患ID';
COMMENT ON COLUMN v_patient_relative_info.patient_status IS '病患状态';
COMMENT ON COLUMN v_patient_relative_info.relative_id IS '家属ID';
COMMENT ON COLUMN v_patient_relative_info.relationship IS '与患者关系';
COMMENT ON COLUMN v_patient_relative_info.status IS '状态';
COMMENT ON COLUMN v_patient_relative_info.create_time IS '创建时间';
COMMENT ON COLUMN v_patient_relative_info.patient_name IS '病患姓名';
COMMENT ON COLUMN v_patient_relative_info.hospital_no IS '住院号';
COMMENT ON COLUMN v_patient_relative_info.ward_name IS '病房名称';
COMMENT ON COLUMN v_patient_relative_info.bed_no IS '床位号';
COMMENT ON COLUMN v_patient_relative_info.relative_name IS '家属姓名';
COMMENT ON COLUMN v_patient_relative_info.relative_mobile IS '家属手机号';
COMMENT ON COLUMN v_patient_relative_info.relative_id_card IS '家属身份证号';


CREATE OR REPLACE VIEW v_family_member_info AS
SELECT 
    u.id,
    u.tenant_id,
    u.mobile,
    u.email,
    u.identity,
    u.name,
    u.zh_name,
    u.gender,
    u.address,
    u.type,
    u.org_id,
    u.id_card,
    u.work_number,
    u.avatar_url,
    u.create_at,
    u.update_at,
    u.remark,
    u.status,
    json_agg(
        json_build_object(
            'patient_id', p.id,
            'patient_name', p.name,
            'hospital_no', p.hospital_no,
            'ward_name', w.name,
            'bed_no', b.bed_no,
            'relationship', pr.relationship,
            'relative_status', pr.status
        )
    ) FILTER (WHERE p.id IS NOT NULL) as patients
FROM o_user u 
LEFT JOIN o_patient_relative pr ON u.id = pr.relative_id AND pr.status = 0
LEFT JOIN o_patient p ON pr.patient_id = p.id AND p.status = 0
LEFT JOIN o_bed b ON p.bed_id = b.id
LEFT JOIN o_ward w ON p.ward_id = w.id
WHERE u.type = 3
GROUP BY u.id, u.tenant_id, u.mobile, u.email, u.identity, u.name, u.zh_name, 
         u.gender, u.address, u.type, u.org_id, u.id_card, u.work_number,
         u.avatar_url, u.create_at, u.update_at, u.remark, u.status;

COMMENT ON VIEW v_family_member_info IS '家属信息视图';
COMMENT ON COLUMN v_family_member_info.id IS '家属ID';
COMMENT ON COLUMN v_family_member_info.tenant_id IS '租户ID';
COMMENT ON COLUMN v_family_member_info.mobile IS '手机号';
COMMENT ON COLUMN v_family_member_info.email IS '邮箱';
COMMENT ON COLUMN v_family_member_info.identity IS '用户标识';
COMMENT ON COLUMN v_family_member_info.name IS '姓名';
COMMENT ON COLUMN v_family_member_info.zh_name IS '中文名';
COMMENT ON COLUMN v_family_member_info.gender IS '性别(0:未知,1:男,2:女)';
COMMENT ON COLUMN v_family_member_info.address IS '地址';
COMMENT ON COLUMN v_family_member_info.type IS '用户类型(3:访客)';
COMMENT ON COLUMN v_family_member_info.org_id IS '组织ID';
COMMENT ON COLUMN v_family_member_info.id_card IS '身份证';
COMMENT ON COLUMN v_family_member_info.work_number IS '工号';
COMMENT ON COLUMN v_family_member_info.avatar_url IS '头像';
COMMENT ON COLUMN v_family_member_info.create_at IS '创建时间';
COMMENT ON COLUMN v_family_member_info.update_at IS '更新时间';
COMMENT ON COLUMN v_family_member_info.remark IS '备注';
COMMENT ON COLUMN v_family_member_info.status IS '状态(1:正常,2:禁止登陆,3:删除)';
COMMENT ON COLUMN v_family_member_info.patients IS '关联的病患信息列表';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS v_family_member_info CASCADE;
DROP VIEW IF EXISTS v_patient_relative_info CASCADE;
DROP TABLE IF EXISTS o_patient_relative CASCADE;
DROP VIEW IF EXISTS v_bed_info CASCADE;
DROP VIEW IF EXISTS v_patient_info CASCADE;
DROP VIEW IF EXISTS v_visit_record_info CASCADE;
DROP VIEW IF EXISTS v_live_record_info CASCADE;
DROP VIEW IF EXISTS v_head_display_info CASCADE;
DROP TABLE IF EXISTS o_system_config CASCADE;
DROP TABLE IF EXISTS o_live_record CASCADE;
DROP TABLE IF EXISTS o_head_display CASCADE;
DROP TABLE IF EXISTS o_visit_record CASCADE;
DROP TABLE IF EXISTS o_visit_schedule CASCADE;
DROP TABLE IF EXISTS o_patient CASCADE;
DROP TABLE IF EXISTS o_bed CASCADE;
DROP TABLE IF EXISTS o_ward CASCADE;
DROP VIEW IF EXISTS v_camera_info CASCADE;
DROP TABLE IF EXISTS o_camera CASCADE;

-- +goose StatementEnd
