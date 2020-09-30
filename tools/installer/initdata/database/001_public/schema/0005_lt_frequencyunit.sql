-- Table: public.lt_frequencyunit

-- DROP TABLE public.lt_frequencyunit;

CREATE TABLE public.lt_frequencyunit
(
  id bigserial NOT NULL, -- รหัสหน่วยของความถี่การเชื่อมโยง  unit of frequency's serial number
  convert_minute numeric, -- แปลงหน่วยเป็นนาที
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  frequencyunit_name json, -- ชื่อหน่วยของระยะเวลาของการเชื่อมโยง
  CONSTRAINT pk_lt_frequencyunit PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_frequencyunit
  IS 'หน่วยของความถี่ของการเชื่อมโยง';
COMMENT ON COLUMN public.lt_frequencyunit.id IS 'รหัสหน่วยของความถี่การเชื่อมโยง  unit of frequency''s serial number';
COMMENT ON COLUMN public.lt_frequencyunit.convert_minute IS 'แปลงหน่วยเป็นนาที';
COMMENT ON COLUMN public.lt_frequencyunit.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_frequencyunit.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_frequencyunit.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_frequencyunit.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_frequencyunit.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_frequencyunit.deleted_at IS 'วันที่ลบข้อมูล deleted date';
COMMENT ON COLUMN public.lt_frequencyunit.frequencyunit_name IS 'ชื่อหน่วยของระยะเวลาของการเชื่อมโยง';