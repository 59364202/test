-- Table: public.wind

-- DROP TABLE public.wind;

CREATE TABLE public.wind
(
  id bigserial NOT NULL, -- รหัสข้อมูลความเร็วลมจากสถานีโทรมาตรอัตโนมัติ wind's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  wind_datetime timestamp with time zone NOT NULL, -- วันที่เก็บค่าความเร็วลม record date
  wind_speed double precision, -- ค่าความเร็วลม wind value
  wind_dir_value double precision, -- ค่าองศาของทิศทางลม wind direction value (degree)
  wind_dir text, -- ทิศทางลม wind direction
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_wind PRIMARY KEY (id),
  CONSTRAINT fk_wind_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_wind UNIQUE (tele_station_id, wind_datetime, deleted_at),
  CONSTRAINT pt_wind_wind_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.wind
  IS 'ความเร็วและทิศทางลม';
COMMENT ON COLUMN public.wind.id IS 'รหัสข้อมูลความเร็วลมจากสถานีโทรมาตรอัตโนมัติ wind''s serial number';
COMMENT ON COLUMN public.wind.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.wind.wind_datetime IS 'วันที่เก็บค่าความเร็วลม record date';
COMMENT ON COLUMN public.wind.wind_speed IS 'ค่าความเร็วลม wind value';
COMMENT ON COLUMN public.wind.wind_dir_value IS 'ค่าองศาของทิศทางลม wind direction value (degree)';
COMMENT ON COLUMN public.wind.wind_dir IS 'ทิศทางลม wind direction';
COMMENT ON COLUMN public.wind.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.wind.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.wind.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.wind.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.wind.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.wind.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.wind.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.wind.deleted_at IS 'วันที่ลบข้อมูล deleted date';

