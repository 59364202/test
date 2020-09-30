-- Table: public.lt_metadata_status

-- DROP TABLE public.lt_metadata_status;

CREATE TABLE public.lt_metadata_status
(
  id bigserial NOT NULL, -- รหัสสถานะของบัญชีข้อมูล status of metadata serial number
  metadatastatus_name json, -- ชื่อสถานะของบัญชีข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_metadata_status PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_metadata_status
  IS 'สถานะของบัญชีข้อมูล';
COMMENT ON COLUMN public.lt_metadata_status.id IS 'รหัสสถานะของบัญชีข้อมูล status of metadata serial number';
COMMENT ON COLUMN public.lt_metadata_status.metadatastatus_name IS 'ชื่อสถานะของบัญชีข้อมูล';
COMMENT ON COLUMN public.lt_metadata_status.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_metadata_status.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_metadata_status.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_metadata_status.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_metadata_status.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_metadata_status.deleted_at IS 'วันที่ลบข้อมูล deleted date';
