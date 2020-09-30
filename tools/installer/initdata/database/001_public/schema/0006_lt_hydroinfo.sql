-- Table: public.lt_hydroinfo

-- DROP TABLE public.lt_hydroinfo;

CREATE TABLE public.lt_hydroinfo
(
  id bigserial NOT NULL, -- รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo 's serial number
  hydroinfo_number smallint, -- ลำดับของชือ เพื่อใช้ในการนำเสนอ
  hydroinfo_name json, -- ชื่อด้าน
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_hydroinfo PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_hydroinfo
  IS ' ข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ แสดงหน่วยงานรับผิดชอบการใช้ข้อมูล 9 ด้าน (Main Responsible Agencies in “9 Aspects of Hydroinformatics”)';
COMMENT ON COLUMN public.lt_hydroinfo.id IS 'รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo ''s serial number';
COMMENT ON COLUMN public.lt_hydroinfo.hydroinfo_number IS 'ลำดับของชือ เพื่อใช้ในการนำเสนอ';
COMMENT ON COLUMN public.lt_hydroinfo.hydroinfo_name IS 'ชื่อด้าน';
COMMENT ON COLUMN public.lt_hydroinfo.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_hydroinfo.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_hydroinfo.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_hydroinfo.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_hydroinfo.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_hydroinfo.deleted_at IS 'วันที่ลบข้อมูล deleted date';

