-- Table: public.metadata_history

-- DROP TABLE public.metadata_history;

CREATE TABLE public.metadata_history
(
  id bigserial NOT NULL, -- รหัสประวัติการแก้ไขบัญชีข้อมูล
  metadata_id bigint NOT NULL, -- รหัสบัญชีข้อมูล metadata serial number
  metadata_datetime time with time zone, -- วันที่แก้ไขบัญชีข้อมูล
  history_description text, -- รายละเอียดของการแก้ไขบัญชีข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_metadata_history PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata_history
  IS 'ประวัติการแก้ไขบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_history.id IS 'รหัสประวัติการแก้ไขบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_history.metadata_id IS 'รหัสบัญชีข้อมูล metadata serial number';
COMMENT ON COLUMN public.metadata_history.metadata_datetime IS 'วันที่แก้ไขบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_history.history_description IS 'รายละเอียดของการแก้ไขบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_history.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata_history.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata_history.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata_history.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata_history.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata_history.deleted_at IS 'วันที่ลบข้อมูล deleted date';

