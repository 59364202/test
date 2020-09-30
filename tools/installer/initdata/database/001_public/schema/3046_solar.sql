-- Table: public.solar

-- DROP TABLE public.solar;

CREATE TABLE public.solar
(
  id bigserial NOT NULL, -- รหัสข้อมูลความเข้มแสงจากสถานีโทรมาตรอัตโนมัติ solar's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  solar_datetime timestamp with time zone NOT NULL, -- วันที่เก็บค่าความเข้มแสง record date
  solar_value double precision NOT NULL, -- ค่าความเข้มแสง solar's value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_solar PRIMARY KEY (id),
  CONSTRAINT fk_solar_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_solar UNIQUE (tele_station_id, solar_datetime, deleted_at),
  CONSTRAINT pt_solar_solar_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.solar
  IS 'ความเข้มแสง';
COMMENT ON COLUMN public.solar.id IS 'รหัสข้อมูลความเข้มแสงจากสถานีโทรมาตรอัตโนมัติ solar''s serial number';
COMMENT ON COLUMN public.solar.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.solar.solar_datetime IS 'วันที่เก็บค่าความเข้มแสง record date';
COMMENT ON COLUMN public.solar.solar_value IS 'ค่าความเข้มแสง solar''s value';
COMMENT ON COLUMN public.solar.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.solar.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.solar.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.solar.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.solar.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.solar.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.solar.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.solar.deleted_at IS 'วันที่ลบข้อมูล deleted date';

