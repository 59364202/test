-- Table: public.lt_servicemethod

-- DROP TABLE public.lt_servicemethod;

CREATE TABLE public.lt_servicemethod
(
  id bigserial NOT NULL, -- รหัสวิธีการให้บริการข้อมูล sevice method s serial number
  servicemethod_name json, -- ชื่อวิธีการให้บริการข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_servicemethod PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_servicemethod
  IS 'วิธีการให้บริการข้อมูล';
COMMENT ON COLUMN public.lt_servicemethod.id IS 'รหัสวิธีการให้บริการข้อมูล sevice method s serial number';
COMMENT ON COLUMN public.lt_servicemethod.servicemethod_name IS 'ชื่อวิธีการให้บริการข้อมูล';
COMMENT ON COLUMN public.lt_servicemethod.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_servicemethod.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_servicemethod.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_servicemethod.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_servicemethod.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_servicemethod.deleted_at IS 'วันที่ลบข้อมูล deleted date';

