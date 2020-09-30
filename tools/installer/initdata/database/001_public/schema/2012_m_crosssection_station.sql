-- Table: public.m_crosssection_station

-- DROP TABLE public.m_crosssection_station;

CREATE TABLE public.m_crosssection_station
(
  id bigserial NOT NULL, -- รหัสสถานีของข้อมูลภาพตัดลำน้ำ section station's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  survey_year text, -- ปีที่สำรวจ survey year
  section_location text, -- ตำแหน่งที่ตั้ง location
  section_oldcode text, -- รหัสรูปตัดเดิม old section station's serial number
  section_filepath text, -- ที่อยู่ของไฟล์ที่ได้รับจากแหล่งข้อมูล file path location
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_crosssection_station PRIMARY KEY (id),
  CONSTRAINT fk_m_crosss_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_crosssection_station UNIQUE (section_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_crosssection_station
  IS 'สถานีของข้อมูลภาพตัดลำน้ำ';
COMMENT ON COLUMN public.m_crosssection_station.id IS 'รหัสสถานีของข้อมูลภาพตัดลำน้ำ section station''s serial number';
COMMENT ON COLUMN public.m_crosssection_station.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_crosssection_station.survey_year IS 'ปีที่สำรวจ survey year';
COMMENT ON COLUMN public.m_crosssection_station.section_location IS 'ตำแหน่งที่ตั้ง location';
COMMENT ON COLUMN public.m_crosssection_station.section_oldcode IS 'รหัสรูปตัดเดิม old section station''s serial number';
COMMENT ON COLUMN public.m_crosssection_station.section_filepath IS 'ที่อยู่ของไฟล์ที่ได้รับจากแหล่งข้อมูล file path location';
COMMENT ON COLUMN public.m_crosssection_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_crosssection_station.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.m_crosssection_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_crosssection_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_crosssection_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_crosssection_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_crosssection_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_crosssection_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';