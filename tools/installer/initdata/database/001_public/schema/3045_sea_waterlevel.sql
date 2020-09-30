-- Table: public.sea_waterlevel

-- DROP TABLE public.sea_waterlevel;

CREATE TABLE public.sea_waterlevel
(
  id bigserial NOT NULL, -- รหัสค่าระดับน้ำจากการวัดของสถานีบริเวณอ่าวไทยและอันดามัน
  sea_station_id bigint, -- รหัสสถานีสำหรับระดับน้ำบริเวณอ่าวไทยและอันดามัน
  waterlevel_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาของค่าระดับน้ำบริเวณอ่าวไทยและอันดามัน
  waterlevel_value double precision, -- ระดับน้ำ
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_sea_waterlevel PRIMARY KEY (id),
  CONSTRAINT fk_sea_wate_reference_m_sea_st FOREIGN KEY (sea_station_id)
      REFERENCES public.m_sea_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_sea_waterlevel UNIQUE (waterlevel_datetime, deleted_at),
  CONSTRAINT pt_sea_waterlevel_waterlevel_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.sea_waterlevel
  IS 'ค่าระดับน้ำบริเวณอ่าวไทยและอันดามัน';
COMMENT ON COLUMN public.sea_waterlevel.id IS 'รหัสค่าระดับน้ำจากการวัดของสถานีบริเวณอ่าวไทยและอันดามัน';
COMMENT ON COLUMN public.sea_waterlevel.sea_station_id IS 'รหัสสถานีสำหรับระดับน้ำบริเวณอ่าวไทยและอันดามัน';
COMMENT ON COLUMN public.sea_waterlevel.waterlevel_datetime IS 'วันที่และเวลาของค่าระดับน้ำบริเวณอ่าวไทยและอันดามัน';
COMMENT ON COLUMN public.sea_waterlevel.waterlevel_value IS 'ระดับน้ำ ';
COMMENT ON COLUMN public.sea_waterlevel.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.sea_waterlevel.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.sea_waterlevel.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.sea_waterlevel.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.sea_waterlevel.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.sea_waterlevel.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.sea_waterlevel.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.sea_waterlevel.deleted_at IS 'วันที่ลบข้อมูล deleted date';

