-- Table: public.sea_water_forecast

-- DROP TABLE public.sea_water_forecast;

CREATE TABLE public.sea_water_forecast
(
  id bigserial NOT NULL, -- รหัสข้อมูลคาดการณ์น้ำท่วม
  sea_station_id bigint, -- รหัสสถานีคาดการณ์น้ำท่วม
  seaforecast_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
  seaforecast_value double precision, -- ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_sea_water_forecast PRIMARY KEY (id),
  CONSTRAINT fk_sea_wate_reference_m_sea_st FOREIGN KEY (sea_station_id)
      REFERENCES public.m_sea_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_seaforecast UNIQUE (sea_station_id, seaforecast_datetime, deleted_at),
  CONSTRAINT pt_sea_water_forecast_seaforecast_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.sea_water_forecast
  IS 'ข้อมูลทำนายระดับน้ำบริเวณอ่าวไทยและอันดามัน';
COMMENT ON COLUMN public.sea_water_forecast.id IS 'รหัสข้อมูลคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.sea_water_forecast.sea_station_id IS 'รหัสสถานีคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.sea_water_forecast.seaforecast_datetime IS 'วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม';
COMMENT ON COLUMN public.sea_water_forecast.seaforecast_value IS 'ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี';
COMMENT ON COLUMN public.sea_water_forecast.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.sea_water_forecast.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.sea_water_forecast.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.sea_water_forecast.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.sea_water_forecast.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.sea_water_forecast.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.sea_water_forecast.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.sea_water_forecast.deleted_at IS 'วันที่ลบข้อมูล deleted date';

