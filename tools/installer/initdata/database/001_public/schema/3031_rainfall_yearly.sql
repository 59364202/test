-- Table: public.rainfall_yearly

-- DROP TABLE public.rainfall_yearly;

CREATE TABLE public.rainfall_yearly
(
  id bigserial NOT NULL, -- รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  rainfall_datetime timestamp with time zone NOT NULL, -- วันที่เก็บปริมาณน้ำฝน Rainfall date
  rainfall_value double precision, -- ปริมาณฝนรายปี ถ้าปริมาณฝนมีไม่ครบทุกวันของแต่ละเดือนให้บันทึกว่าคำนวณมาจากกี่วัน
  day_count bigint, -- จำนวนวันที่ได้จากการคำนวณ
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainfall_yearly PRIMARY KEY (id),
  CONSTRAINT fk_rainfall_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainfall_yearly UNIQUE (tele_station_id, rainfall_datetime, deleted_at),
  CONSTRAINT pt_rainfall_yearly_rainfall_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainfall_yearly
  IS 'ฝนรายปี';
COMMENT ON COLUMN public.rainfall_yearly.id IS 'รหัสข้อมูลปริมาณน้ำฝนจากสถานีโทรมาตรอัตโนมัติ rainfall''s serial number';
COMMENT ON COLUMN public.rainfall_yearly.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.rainfall_yearly.rainfall_datetime IS 'วันที่เก็บปริมาณน้ำฝน Rainfall date';
COMMENT ON COLUMN public.rainfall_yearly.rainfall_value IS 'ปริมาณฝนรายปี ถ้าปริมาณฝนมีไม่ครบทุกวันของแต่ละเดือนให้บันทึกว่าคำนวณมาจากกี่วัน';
COMMENT ON COLUMN public.rainfall_yearly.day_count IS 'จำนวนวันที่ได้จากการคำนวณ';
COMMENT ON COLUMN public.rainfall_yearly.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainfall_yearly.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainfall_yearly.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainfall_yearly.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainfall_yearly.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainfall_yearly.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainfall_yearly.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainfall_yearly.deleted_at IS 'วันที่ลบข้อมูล deleted date';

