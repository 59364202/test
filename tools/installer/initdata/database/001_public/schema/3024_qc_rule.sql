-- Table: public.qc_rule

-- DROP TABLE public.qc_rule;

CREATE TABLE public.qc_rule
(
  id bigserial NOT NULL, -- รหัสการตรวจสอบคุณภาพข้อมูล quality data 's serial number
  dataimport_dataset_id bigint, -- รหัส dataset
  group_name text NOT NULL, -- ชื่อกลุ่มของการตรวจสอบ
  expression text, -- กฎของการตรวจสอบ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_qc_rule PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.qc_rule
  IS 'เกณฑ์การวัดคุณภาพของข้อมูล';
COMMENT ON COLUMN public.qc_rule.id IS 'รหัสการตรวจสอบคุณภาพข้อมูล quality data ''s serial number';
COMMENT ON COLUMN public.qc_rule.dataimport_dataset_id IS 'รหัส dataset';
COMMENT ON COLUMN public.qc_rule.group_name IS 'ชื่อกลุ่มของการตรวจสอบ';
COMMENT ON COLUMN public.qc_rule.expression IS 'กฎของการตรวจสอบ';
COMMENT ON COLUMN public.qc_rule.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.qc_rule.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.qc_rule.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.qc_rule.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.qc_rule.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.qc_rule.deleted_at IS 'วันที่ลบข้อมูล deleted date';

