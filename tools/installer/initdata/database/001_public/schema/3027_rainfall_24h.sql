-- Table: public.rainfall_24h

-- DROP TABLE public.rainfall_24h;

CREATE TABLE public.rainfall_24h
(
  id bigserial NOT NULL, -- รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station number
  rainfall_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall_datetime_calc timestamp with time zone NOT NULL, -- วันที่และเวลาจริงที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall24h double precision, -- ปริมาณฝนย้อนหลัง 24 ชั่วโมง
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainfall_24h PRIMARY KEY (id),
  CONSTRAINT fk_rainfall_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainfall_24h UNIQUE (tele_station_id, rainfall_datetime, deleted_at),
  CONSTRAINT pt_rainfall_24h_rainfall_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainfall_24h
  IS 'ฝนสะสม 24 ชั่วโมง เป็นตารางทีได้จากการคำนวณทุกชั่วโมง';
COMMENT ON COLUMN public.rainfall_24h.id IS 'รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number';
COMMENT ON COLUMN public.rainfall_24h.tele_station_id IS 'รหัสสถานีโทรมาตร tele station number';
COMMENT ON COLUMN public.rainfall_24h.rainfall_datetime IS 'วันที่และเวลาที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_24h.rainfall_datetime_calc IS 'วันที่และเวลาจริงที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_24h.rainfall24h IS 'ปริมาณฝนย้อนหลัง 24 ชั่วโมง';
COMMENT ON COLUMN public.rainfall_24h.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainfall_24h.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainfall_24h.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainfall_24h.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainfall_24h.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainfall_24h.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainfall_24h.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainfall_24h.deleted_at IS 'วันที่ลบข้อมูล deleted date';