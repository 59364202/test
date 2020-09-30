-- Table: public.soilmoisture

-- DROP TABLE public.soilmoisture;

CREATE TABLE public.soilmoisture
(
  id bigserial NOT NULL, -- ข้อมูลความชื้นในดิน soil moisture's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  soil_datetime timestamp with time zone NOT NULL, -- วันที่ของค่าข้อมูลความชื้นในดิน record date
  soil_value double precision, -- ค่าข้อมูลความชื้นในดิน soil moisture value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_soilmoisture PRIMARY KEY (id),
  CONSTRAINT fk_soilmois_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_soilmoisture UNIQUE (tele_station_id, soil_datetime, deleted_at),
  CONSTRAINT pt_soilmoisture_soil_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.soilmoisture
  IS 'ข้อมูลความชื้นในดิน';
COMMENT ON COLUMN public.soilmoisture.id IS 'ข้อมูลความชื้นในดิน soil moisture''s serial number';
COMMENT ON COLUMN public.soilmoisture.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.soilmoisture.soil_datetime IS 'วันที่ของค่าข้อมูลความชื้นในดิน record date';
COMMENT ON COLUMN public.soilmoisture.soil_value IS 'ค่าข้อมูลความชื้นในดิน soil moisture value';
COMMENT ON COLUMN public.soilmoisture.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.soilmoisture.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.soilmoisture.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.soilmoisture.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.soilmoisture.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.soilmoisture.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.soilmoisture.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.soilmoisture.deleted_at IS 'วันที่ลบข้อมูล deleted date';

