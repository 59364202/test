-- Table: public.lt_dataunit

-- DROP TABLE public.lt_dataunit;

CREATE TABLE public.lt_dataunit
(
  id bigserial NOT NULL, -- รหัสหน่วยของข้อมูล unit of data's serial number
  dataunit_name json, -- ชื่อหน่วยข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_dataunit PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_dataunit
  IS 'หน่วยของข้อมูล';
COMMENT ON COLUMN public.lt_dataunit.id IS 'รหัสหน่วยของข้อมูล unit of data''s serial number';
COMMENT ON COLUMN public.lt_dataunit.dataunit_name IS 'ชื่อหน่วยข้อมูล';
COMMENT ON COLUMN public.lt_dataunit.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_dataunit.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_dataunit.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_dataunit.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_dataunit.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_dataunit.deleted_at IS 'วันที่ลบข้อมูล deleted date';
