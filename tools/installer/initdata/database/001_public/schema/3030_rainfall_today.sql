-- Table: public.rainfall_today

-- DROP TABLE public.rainfall_today;

CREATE TABLE public.rainfall_today
(
  id bigserial NOT NULL, -- รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station number
  rainfall_datetime timestamp with time zone NOT NULL, -- วันที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall_datetime_calc timestamp with time zone NOT NULL, -- วันที่คำนวณปริมาณน้ำฝน Rainfall date
  rainfall_value double precision, -- ปริมาณฝนวันนี้ (เวลา 7.01 - เวลาปัจจุบัน)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainfall_today PRIMARY KEY (id),
  CONSTRAINT fk_rainfall_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainfall_today UNIQUE (tele_station_id, rainfall_datetime, deleted_at),
  CONSTRAINT pt_rainfall_today_rainfall_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainfall_today
  IS 'ฝนวันนี้ (เวลา 7.01 - เวลาปัจจุบัน)';
COMMENT ON COLUMN public.rainfall_today.id IS 'รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall''s serial number';
COMMENT ON COLUMN public.rainfall_today.tele_station_id IS 'รหัสสถานีโทรมาตร tele station number';
COMMENT ON COLUMN public.rainfall_today.rainfall_datetime IS 'วันที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_today.rainfall_datetime_calc IS 'วันที่คำนวณปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_today.rainfall_value IS 'ปริมาณฝนวันนี้ (เวลา 7.01 - เวลาปัจจุบัน)';
COMMENT ON COLUMN public.rainfall_today.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainfall_today.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainfall_today.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainfall_today.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainfall_today.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainfall_today.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainfall_today.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainfall_today.deleted_at IS 'วันที่ลบข้อมูล deleted date';

