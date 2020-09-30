-- Table: public.ford_waterlevel

-- DROP TABLE public.ford_waterlevel;

CREATE TABLE public.ford_waterlevel
(
  id bigserial NOT NULL, -- รหัสระดับน่ำบนถนน ford water level's serial number
  ford_station_id bigint, -- รหัสสถานีวัดระดับน้ำบนถนน
  ford_waterlevel_datetime timestamp with time zone NOT NULL, -- วันที่วัดระดับน่ำบนถนน
  ford_waterlevel_value double precision, -- ค่าระดับน่ำบนถนน
  comm_status text, -- สถานะของเครื่องวัด meter status
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_ford_waterlevel PRIMARY KEY (id),
  CONSTRAINT fk_ford_wat_reference_m_ford_s FOREIGN KEY (ford_station_id)
      REFERENCES public.m_ford_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_ford_waterlevel UNIQUE (ford_station_id, ford_waterlevel_datetime, deleted_at),
  CONSTRAINT pt_ford_waterlevel_ford_waterlevel_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.ford_waterlevel
  IS 'ระดับน้ำบนผิวจราจร';
COMMENT ON COLUMN public.ford_waterlevel.id IS 'รหัสระดับน่ำบนถนน ford water level''s serial number';
COMMENT ON COLUMN public.ford_waterlevel.ford_station_id IS 'รหัสสถานีวัดระดับน้ำบนถนน ';
COMMENT ON COLUMN public.ford_waterlevel.ford_waterlevel_datetime IS 'วันที่วัดระดับน่ำบนถนน ';
COMMENT ON COLUMN public.ford_waterlevel.ford_waterlevel_value IS 'ค่าระดับน่ำบนถนน';
COMMENT ON COLUMN public.ford_waterlevel.comm_status IS 'สถานะของเครื่องวัด meter status';
COMMENT ON COLUMN public.ford_waterlevel.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.ford_waterlevel.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.ford_waterlevel.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.ford_waterlevel.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.ford_waterlevel.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.ford_waterlevel.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.ford_waterlevel.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.ford_waterlevel.deleted_at IS 'วันที่ลบข้อมูล deleted date';

