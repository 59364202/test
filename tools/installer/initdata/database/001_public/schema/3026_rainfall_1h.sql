-- Table: public.rainfall_1h

-- DROP TABLE public.rainfall_1h;

CREATE TABLE public.rainfall_1h
(
  id bigserial NOT NULL, -- รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station number
  rainfall_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall_datetime_calc timestamp with time zone NOT NULL, -- วันที่คำนวณปริมาณน้ำฝน Rainfall date
  rainfall1h double precision, -- ปริมาณฝนรายชั่วโมง
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainfall_1h PRIMARY KEY (id),
  CONSTRAINT fk_rainfall_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainfall_1h UNIQUE (tele_station_id, rainfall_datetime, deleted_at),
  CONSTRAINT pt_rainfall_1h_rainfall_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainfall_1h
  IS 'ฝนรายชั่วโมง เป็นตารางทีได้จากการคำนวณทุกชั่วโมง';
COMMENT ON COLUMN public.rainfall_1h.id IS 'รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number';
COMMENT ON COLUMN public.rainfall_1h.tele_station_id IS 'รหัสสถานีโทรมาตร tele station number';
COMMENT ON COLUMN public.rainfall_1h.rainfall_datetime IS 'วันที่และเวลาที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_1h.rainfall_datetime_calc IS 'วันที่คำนวณปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_1h.rainfall1h IS 'ปริมาณฝนรายชั่วโมง';
COMMENT ON COLUMN public.rainfall_1h.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainfall_1h.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainfall_1h.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainfall_1h.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainfall_1h.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainfall_1h.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainfall_1h.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainfall_1h.deleted_at IS 'วันที่ลบข้อมูล deleted date';

