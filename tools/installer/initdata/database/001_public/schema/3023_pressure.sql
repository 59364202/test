-- Table: public.pressure

-- DROP TABLE public.pressure;

CREATE TABLE public.pressure
(
  id bigserial NOT NULL, -- รหัสข้อมูลค่าความกดอากาศจากสถานีโทรมาตรอัตโนมัติ pressure's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  pressure_datetime timestamp with time zone NOT NULL, -- วันที่เก็บค่าความกดอากาศ record date
  pressure_value double precision NOT NULL, -- ค่าความกดอากาศ pressure value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_pressure PRIMARY KEY (id),
  CONSTRAINT fk_pressure_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_pressure UNIQUE (tele_station_id, pressure_datetime, deleted_at),
  CONSTRAINT pt_pressure_pressure_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.pressure
  IS 'ค่าความกดอากาศ หน่วย : hPa';
COMMENT ON COLUMN public.pressure.id IS 'รหัสข้อมูลค่าความกดอากาศจากสถานีโทรมาตรอัตโนมัติ pressure''s serial number';
COMMENT ON COLUMN public.pressure.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.pressure.pressure_datetime IS 'วันที่เก็บค่าความกดอากาศ record date';
COMMENT ON COLUMN public.pressure.pressure_value IS 'ค่าความกดอากาศ pressure value';
COMMENT ON COLUMN public.pressure.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.pressure.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.pressure.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.pressure.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.pressure.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.pressure.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.pressure.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.pressure.deleted_at IS 'วันที่ลบข้อมูล deleted date';

