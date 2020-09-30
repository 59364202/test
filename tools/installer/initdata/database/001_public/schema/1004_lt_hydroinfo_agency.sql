-- Table: public.lt_hydroinfo_agency

-- DROP TABLE public.lt_hydroinfo_agency;

CREATE TABLE public.lt_hydroinfo_agency
(
  id bigserial NOT NULL, -- รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo 's serial number
  hydroinfo_id bigint, -- รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_hydroinfo_agency PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_hydroinfo_agency
  IS ' ข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ แสดงหน่วยงานรับผิดชอบการใช้ข้อมูล 9 ด้าน (Main Responsible Agencies in “9 Aspects of Hydroinformatics”) กับหน่วยงานที่เชื่อมโยง';
COMMENT ON COLUMN public.lt_hydroinfo_agency.id IS 'รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo ''s serial number';
COMMENT ON COLUMN public.lt_hydroinfo_agency.hydroinfo_id IS 'รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ';
COMMENT ON COLUMN public.lt_hydroinfo_agency.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.lt_hydroinfo_agency.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_hydroinfo_agency.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_hydroinfo_agency.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_hydroinfo_agency.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_hydroinfo_agency.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_hydroinfo_agency.deleted_at IS 'วันที่ลบข้อมูล deleted date';