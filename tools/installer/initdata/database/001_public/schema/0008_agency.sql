-- Table: public.agency

-- DROP TABLE public.agency;

CREATE TABLE public.agency
(
  id bigserial NOT NULL, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
  department_id bigint NOT NULL, -- รหัสกรม
  agency_shortname json, -- ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ) abbreviation (English)
  agency_name json NOT NULL, -- ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_agency PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.agency
  IS 'หน่วยงานที่เชื่อมโยงข้อมูลกับคลังข้อมูลฯ';
COMMENT ON COLUMN public.agency.id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency''s serial number';
COMMENT ON COLUMN public.agency.department_id IS 'รหัสกรม';
COMMENT ON COLUMN public.agency.agency_shortname IS 'ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ) abbreviation (English)';
COMMENT ON COLUMN public.agency.agency_name IS 'ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ';
COMMENT ON COLUMN public.agency.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.agency.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.agency.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.agency.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.agency.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.agency.deleted_at IS 'วันที่ลบข้อมูล deleted date';