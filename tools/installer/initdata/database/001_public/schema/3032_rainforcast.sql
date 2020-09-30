-- Table: public.rainforecast

-- DROP TABLE public.rainforecast;

CREATE TABLE public.rainforecast
(
  id bigserial NOT NULL, -- รหัสข้อมูลคาดการณ์ฝน
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  agency_id bigint, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
  rainforecast_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บข้อมูลคาดการณ์ฝน
  rainforecast_value double precision, -- ข้อมูลคาดการณ์ฝน
  rainforecast_level text, -- เกณฑ์ของการคาดการณ์
  rainforecast_leveltext text, -- รายละเอียดเกณฑ์ของการคาดการณ์
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rainforecast PRIMARY KEY (id),
  CONSTRAINT fk_rainforc_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_rainforc_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_rainforecast UNIQUE (geocode_id, rainforecast_datetime, deleted_at),
  CONSTRAINT pt_rainforecast_rainforecast_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rainforecast
  IS 'ข้อมูลคาดการณ์ฝน';
COMMENT ON COLUMN public.rainforecast.id IS 'รหัสข้อมูลคาดการณ์ฝน';
COMMENT ON COLUMN public.rainforecast.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.rainforecast.agency_id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency''s serial number';
COMMENT ON COLUMN public.rainforecast.rainforecast_datetime IS 'วันที่และเวลาที่เก็บข้อมูลคาดการณ์ฝน';
COMMENT ON COLUMN public.rainforecast.rainforecast_value IS 'ข้อมูลคาดการณ์ฝน';
COMMENT ON COLUMN public.rainforecast.rainforecast_level IS 'เกณฑ์ของการคาดการณ์';
COMMENT ON COLUMN public.rainforecast.rainforecast_leveltext IS 'รายละเอียดเกณฑ์ของการคาดการณ์';
COMMENT ON COLUMN public.rainforecast.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rainforecast.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.rainforecast.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rainforecast.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rainforecast.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rainforecast.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rainforecast.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rainforecast.deleted_at IS 'วันที่ลบข้อมูล deleted date';

