-- Table: public.storm

-- DROP TABLE public.storm;

CREATE TABLE public.storm
(
  id bigserial NOT NULL, -- รหัสข้อมูลเส้นทางพายุ
  agency_id bigint, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
  storm_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บข้อมูลเส้นทางพายุ
  storm_lat numeric(9,6), -- ละติจูด
  storm_directionlat text, -- ทิศทาง N หรือ S
  storm_long numeric(9,6), -- ลองติจูด
  storm_directionlong text, -- ทิศทาง E หรือ W
  storm_name text, -- ชื่อพายุ
  storm_course text, -- Course of storm (Degrees, True)
  storm_speed text, -- ความเร็ว Speed of storm kt (speed is always in knots)
  storm_pressure text, -- ความกดอากาศ mb (pressure is always in mb)
  storm_wind text, -- Maximum sustained wind in storm (kt)
  storm_windgusts text, -- Maximum wind gusts in storm (kt)
  storm_observationtype text, -- Type of observation (ACT=actual, FOR=forecast)
  storm_act_for text, -- Actual Observations: Date record written (decimal years)
  storm_observationdate text, -- Date of observation (decimal years)
  storm_wmo text, -- WMO Header for observation
  storm_timeunit text, -- หน่วยเวลา
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_storm PRIMARY KEY (id),
  CONSTRAINT fk_storm_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_storm UNIQUE (storm_datetime, deleted_at, storm_name),
  CONSTRAINT pt_storm_storm_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.storm
  IS 'ข้อมูลเส้นทางพายุ';
COMMENT ON COLUMN public.storm.id IS 'รหัสข้อมูลเส้นทางพายุ';
COMMENT ON COLUMN public.storm.agency_id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency''s serial number';
COMMENT ON COLUMN public.storm.storm_datetime IS 'วันที่และเวลาที่เก็บข้อมูลเส้นทางพายุ';
COMMENT ON COLUMN public.storm.storm_lat IS 'ละติจูด';
COMMENT ON COLUMN public.storm.storm_directionlat IS 'ทิศทาง N หรือ S';
COMMENT ON COLUMN public.storm.storm_long IS 'ลองติจูด';
COMMENT ON COLUMN public.storm.storm_directionlong IS 'ทิศทาง E หรือ W';
COMMENT ON COLUMN public.storm.storm_name IS 'ชื่อพายุ';
COMMENT ON COLUMN public.storm.storm_course IS 'Course of storm (Degrees, True)';
COMMENT ON COLUMN public.storm.storm_speed IS 'ความเร็ว Speed of storm kt (speed is always in knots)';
COMMENT ON COLUMN public.storm.storm_pressure IS 'ความกดอากาศ mb (pressure is always in mb)';
COMMENT ON COLUMN public.storm.storm_wind IS 'Maximum sustained wind in storm (kt)';
COMMENT ON COLUMN public.storm.storm_windgusts IS 'Maximum wind gusts in storm (kt)';
COMMENT ON COLUMN public.storm.storm_observationtype IS 'Type of observation (ACT=actual, FOR=forecast)';
COMMENT ON COLUMN public.storm.storm_act_for IS 'Actual Observations: Date record written (decimal years)';
COMMENT ON COLUMN public.storm.storm_observationdate IS 'Date of observation (decimal years)';
COMMENT ON COLUMN public.storm.storm_wmo IS 'WMO Header for observation';
COMMENT ON COLUMN public.storm.storm_timeunit IS 'หน่วยเวลา';
COMMENT ON COLUMN public.storm.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.storm.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.storm.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.storm.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.storm.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.storm.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.storm.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.storm.deleted_at IS 'วันที่ลบข้อมูล deleted date';

