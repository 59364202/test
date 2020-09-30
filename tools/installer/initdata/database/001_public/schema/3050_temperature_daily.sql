-- Table: public.temperature_daily

-- DROP TABLE public.temperature_daily;

CREATE TABLE public.temperature_daily
(
  id bigserial NOT NULL, -- รหัสข้อมูลค่าอุณหภูมิจากสถานีโทรมาตรอัตโนมัติ temperature's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  temperature_date timestamp with time zone NOT NULL, -- วันที่เก็บค่าอุณหภูมิ record date
  temperature_value double precision NOT NULL, -- ค่าอุณหภูมิ temperature value
  maxtemperature double precision, -- ค่าอุณหภูมิสูงสุด max temperature value
  diffmaxtemperature double precision, -- ส่วนต่างค่าอุณหภูมิสูงสุด difference max temperature value
  mintemperature double precision, -- ค่าอุณหภูมิต่ำสุด min temperature value
  diffmintemperature double precision, -- ส่วนต่างค่าอุณหภูมิต่ำสุด difference min  temperature value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_temperature_daily PRIMARY KEY (id),
  CONSTRAINT fk_temperat_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_temperature_daily UNIQUE (tele_station_id, temperature_date, deleted_at),
  CONSTRAINT pt_temperature_daily_temperature_date CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.temperature_daily
  IS 'อุณหภูมิรายวัน';
COMMENT ON COLUMN public.temperature_daily.id IS 'รหัสข้อมูลค่าอุณหภูมิจากสถานีโทรมาตรอัตโนมัติ temperature''s serial number';
COMMENT ON COLUMN public.temperature_daily.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.temperature_daily.temperature_date IS 'วันที่เก็บค่าอุณหภูมิ record date';
COMMENT ON COLUMN public.temperature_daily.temperature_value IS 'ค่าอุณหภูมิ temperature value';
COMMENT ON COLUMN public.temperature_daily.maxtemperature IS 'ค่าอุณหภูมิสูงสุด max temperature value';
COMMENT ON COLUMN public.temperature_daily.diffmaxtemperature IS 'ส่วนต่างค่าอุณหภูมิสูงสุด difference max temperature value';
COMMENT ON COLUMN public.temperature_daily.mintemperature IS 'ค่าอุณหภูมิต่ำสุด min temperature value';
COMMENT ON COLUMN public.temperature_daily.diffmintemperature IS 'ส่วนต่างค่าอุณหภูมิต่ำสุด difference min  temperature value';
COMMENT ON COLUMN public.temperature_daily.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.temperature_daily.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.temperature_daily.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.temperature_daily.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.temperature_daily.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.temperature_daily.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.temperature_daily.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.temperature_daily.deleted_at IS 'วันที่ลบข้อมูล deleted date';

