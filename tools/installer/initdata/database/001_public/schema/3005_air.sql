-- Table: public.air

-- DROP TABLE public.air;

CREATE TABLE public.air
(
  id bigserial NOT NULL, -- รหัสค่าคุณภาพอากาศจากการวัดของสถานี
  air_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  air_datetime timestamp with time zone NOT NULL, -- วันที่ตรวจสอบค่าคุณภาพอากาศ
  air_so2 double precision, -- ก๊าซซัลเฟอร์ไดออกไซด์ unit= ppb
  air_no2 double precision, -- ก๊าซไนโตรเจนไดออกไซด์ unit= ppb
  air_co double precision, -- ก๊าซคาร์บอนมอนนอกไซด์ unit=ppm
  air_o3 double precision, -- ก๊าซโอโซน unit= ppb
  air_pm10 double precision, -- ฝุ่นละอองขนาดไม่เกิน 10 ไมครอน unit=?g/m?
  air_pm25 double precision, -- ฝุ่นละอองขนาดไม่เกิน 2.5 ไมครอน
  air_aqi double precision, -- ดัชนีคุณภาพอากาศ (Air quality Index)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_air PRIMARY KEY (id),
  CONSTRAINT fk_air_reference_m_air_st FOREIGN KEY (air_station_id)
      REFERENCES public.m_air_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_air UNIQUE (air_station_id, air_datetime, deleted_at),
  CONSTRAINT pt_air_air_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.air
  IS 'ค่าตรวจวัดคุณภาพอากาศ';
COMMENT ON COLUMN public.air.id IS 'รหัสค่าคุณภาพอากาศจากการวัดของสถานี';
COMMENT ON COLUMN public.air.air_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.air.air_datetime IS 'วันที่ตรวจสอบค่าคุณภาพอากาศ';
COMMENT ON COLUMN public.air.air_so2 IS 'ก๊าซซัลเฟอร์ไดออกไซด์ unit= ppb';
COMMENT ON COLUMN public.air.air_no2 IS 'ก๊าซไนโตรเจนไดออกไซด์ unit= ppb';
COMMENT ON COLUMN public.air.air_co IS 'ก๊าซคาร์บอนมอนนอกไซด์ unit=ppm';
COMMENT ON COLUMN public.air.air_o3 IS 'ก๊าซโอโซน unit= ppb';
COMMENT ON COLUMN public.air.air_pm10 IS 'ฝุ่นละอองขนาดไม่เกิน 10 ไมครอน unit=?g/m?';
COMMENT ON COLUMN public.air.air_pm25 IS 'ฝุ่นละอองขนาดไม่เกิน 2.5 ไมครอน';
COMMENT ON COLUMN public.air.air_aqi IS 'ดัชนีคุณภาพอากาศ (Air quality Index)';
COMMENT ON COLUMN public.air.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.air.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.air.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.air.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.air.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.air.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.air.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.air.deleted_at IS 'วันที่ลบข้อมูล deleted date';

