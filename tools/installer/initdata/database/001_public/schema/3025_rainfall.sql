-- Table: public.rainfall

-- DROP TABLE public.rainfall;

CREATE TABLE public.rainfall
(
  id bigserial NOT NULL, -- รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station  number
  rainfall_datetime timestamp with time zone NOT NULL, -- วันที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall5m double precision, -- ปริมาณน้ำฝนทุก 5 นาที Rainfall Every 5 minute
  rainfall_date_calc date, -- วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall
  rainfall10m double precision, -- ปริมาณน้ำฝนทุก 10 นาที Rainfall Every 10 minute
  rainfall15m double precision, -- ปริมาณน้ำฝนทุก 15 นาที Rainfall Every 15 minute
  rainfall30m double precision, -- ปริมาณน้ำฝนทุก 30 นาที Rainfall Every 30  minute
  rainfall1h double precision, -- ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
  rainfall3h double precision, -- ปริมาณน้ำฝนทุก 3 ชั่วโมง Rainfall Every 3 hours
  rainfall6h double precision, -- ปริมาณน้ำฝนทุก 6 ชั่วโมง Rainfall Every 6 hours
  rainfall12h double precision, -- ปริมาณน้ำฝนทุก 12 ชั่วโมง Rainfall Every 12  hours
  rainfall24h double precision, -- ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24 hours
  rainfall_acc double precision, -- ปริมาณน้ำฝนสะสม Rainfall Accumulate
  rainfall_today double precision, -- ปริมาณน้ำฝนสะสมวันนี้
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainfall PRIMARY KEY (id),
  CONSTRAINT fk_rainfall_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainfall UNIQUE (tele_station_id, rainfall_datetime, deleted_at),
  CONSTRAINT pt_rainfall_rainfall_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainfall
  IS 'ฝน';
COMMENT ON COLUMN public.rainfall.id IS 'รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall number';
COMMENT ON COLUMN public.rainfall.tele_station_id IS 'รหัสสถานีโทรมาตร tele station  number';
COMMENT ON COLUMN public.rainfall.rainfall_datetime IS 'วันที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall.rainfall5m IS 'ปริมาณน้ำฝนทุก 5 นาที Rainfall Every 5 minute';
COMMENT ON COLUMN public.rainfall.rainfall_date_calc IS 'วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall';
COMMENT ON COLUMN public.rainfall.rainfall10m IS 'ปริมาณน้ำฝนทุก 10 นาที Rainfall Every 10 minute';
COMMENT ON COLUMN public.rainfall.rainfall15m IS 'ปริมาณน้ำฝนทุก 15 นาที Rainfall Every 15 minute';
COMMENT ON COLUMN public.rainfall.rainfall30m IS 'ปริมาณน้ำฝนทุก 30 นาที Rainfall Every 30  minute';
COMMENT ON COLUMN public.rainfall.rainfall1h IS 'ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour';
COMMENT ON COLUMN public.rainfall.rainfall3h IS 'ปริมาณน้ำฝนทุก 3 ชั่วโมง Rainfall Every 3 hours';
COMMENT ON COLUMN public.rainfall.rainfall6h IS 'ปริมาณน้ำฝนทุก 6 ชั่วโมง Rainfall Every 6 hours';
COMMENT ON COLUMN public.rainfall.rainfall12h IS 'ปริมาณน้ำฝนทุก 12 ชั่วโมง Rainfall Every 12  hours';
COMMENT ON COLUMN public.rainfall.rainfall24h IS 'ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24 hours';
COMMENT ON COLUMN public.rainfall.rainfall_acc IS 'ปริมาณน้ำฝนสะสม Rainfall Accumulate';
COMMENT ON COLUMN public.rainfall.rainfall_today IS 'ปริมาณน้ำฝนสะสมวันนี้';
COMMENT ON COLUMN public.rainfall.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainfall.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainfall.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainfall.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainfall.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainfall.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainfall.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainfall.deleted_at IS 'วันที่ลบข้อมูล deleted date';

