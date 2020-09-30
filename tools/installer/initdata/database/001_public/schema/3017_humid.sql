-- Table: public.humid

-- DROP TABLE public.humid;

CREATE TABLE public.humid
(
  id bigserial NOT NULL, -- รหัสข้อมูลค่าความชื้นสัมพัทธ์จากสถานีโทรมาตรอัตโนมัติ humid's serial number
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  humid_datetime timestamp with time zone NOT NULL, -- วันที่เก็บค่าความชื้นสัมพัทธ์ record date
  humid_value double precision NOT NULL, -- ค่าความชื้นสัมพัทธ์ humid value
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_humid PRIMARY KEY (id),
  CONSTRAINT fk_humid_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_humid UNIQUE (tele_station_id, humid_datetime, deleted_at),
  CONSTRAINT pt_humid_humid_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.humid
  IS 'ความชื้นสัมพัทธ์';
COMMENT ON COLUMN public.humid.id IS 'รหัสข้อมูลค่าความชื้นสัมพัทธ์จากสถานีโทรมาตรอัตโนมัติ humid''s serial number';
COMMENT ON COLUMN public.humid.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.humid.humid_datetime IS 'วันที่เก็บค่าความชื้นสัมพัทธ์ record date';
COMMENT ON COLUMN public.humid.humid_value IS 'ค่าความชื้นสัมพัทธ์ humid value';
COMMENT ON COLUMN public.humid.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.humid.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.humid.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.humid.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.humid.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.humid.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.humid.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.humid.deleted_at IS 'วันที่ลบข้อมูล deleted date';

