-- Table: public.temperature

-- DROP TABLE public.temperature;

CREATE TABLE public.temperature
(
  id bigserial NOT NULL, -- รหัสข้อมูลค่าอุณหภูมิจากสถานีโทรมาตรอัตโนมัติ temperature's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  temp_datetime timestamp with time zone NOT NULL, -- วันที่เก็บค่าอุณหภูมิ record date
  temp_value double precision NOT NULL, -- ค่าอุณหภูมิ temperature value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_temperature PRIMARY KEY (id),
  CONSTRAINT fk_temperat_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_temperature UNIQUE (tele_station_id, temp_datetime, deleted_at),
  CONSTRAINT pt_temperature_temp_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.temperature
  IS 'อุณหภูมิ';
COMMENT ON COLUMN public.temperature.id IS 'รหัสข้อมูลค่าอุณหภูมิจากสถานีโทรมาตรอัตโนมัติ temperature''s serial number';
COMMENT ON COLUMN public.temperature.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.temperature.temp_datetime IS 'วันที่เก็บค่าอุณหภูมิ record date';
COMMENT ON COLUMN public.temperature.temp_value IS 'ค่าอุณหภูมิ temperature value';
COMMENT ON COLUMN public.temperature.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.temperature.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.temperature.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.temperature.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.temperature.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.temperature.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.temperature.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.temperature.deleted_at IS 'วันที่ลบข้อมูล deleted date';

