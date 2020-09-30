-- Table: public.floodforecast

-- DROP TABLE public.floodforecast;

CREATE TABLE public.floodforecast
(
  id bigserial NOT NULL, -- รหัสข้อมูลคาดการณ์น้ำท่วม
  floodforecast_station_id bigint, -- รหัสสถานีคาดการณ์น้ำท่วม
  floodforecast_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
  floodforecast_value double precision, -- ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_floodforecast PRIMARY KEY (id),
  CONSTRAINT fk_floodfor_reference_m_floodf FOREIGN KEY (floodforecast_station_id)
      REFERENCES public.m_floodforecast_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_floodforcast UNIQUE (floodforecast_station_id, floodforecast_datetime, deleted_at),
  CONSTRAINT pt_floodforecast_floodforecast_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.floodforecast
  IS 'ข้อมูลคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.floodforecast.id IS 'รหัสข้อมูลคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.floodforecast.floodforecast_station_id IS 'รหัสสถานีคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.floodforecast.floodforecast_datetime IS 'วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.floodforecast.floodforecast_value IS 'ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี';
COMMENT ON COLUMN public.floodforecast.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.floodforecast.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.floodforecast.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.floodforecast.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.floodforecast.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.floodforecast.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.floodforecast.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.floodforecast.deleted_at IS 'วันที่ลบข้อมูล deleted date';

