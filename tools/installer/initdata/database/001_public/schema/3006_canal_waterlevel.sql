-- Table: public.canal_waterlevel

-- DROP TABLE public.canal_waterlevel;

CREATE TABLE public.canal_waterlevel
(
  id bigserial NOT NULL, -- รหัสระดับน้ำในคลอง canal water level's serial number
  canal_station_id bigint, -- รหัสสถานีวัดระดับน้ำในคลอง canal station's serial number
  canal_waterlevel_datetime timestamp with time zone NOT NULL, -- วันที่วัดระดับน้ำในคลอง record date
  canal_waterlevel_value double precision, -- ค่าระดับน้ำในคลอง canal water level value
  comm_status text, -- สถานะของเครื่องวัด meter status
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_canal_waterlevel PRIMARY KEY (id),
  CONSTRAINT fk_canal_wa_reference_m_canal_ FOREIGN KEY (canal_station_id)
      REFERENCES public.m_canal_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_canal_waterlevel UNIQUE (canal_station_id, canal_waterlevel_datetime, deleted_at),
  CONSTRAINT pt_canal_waterlevel_canal_waterlevel_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.canal_waterlevel
  IS 'ระดับน้ำในคลอง';
COMMENT ON COLUMN public.canal_waterlevel.id IS 'รหัสระดับน้ำในคลอง canal water level''s serial number';
COMMENT ON COLUMN public.canal_waterlevel.canal_station_id IS 'รหัสสถานีวัดระดับน้ำในคลอง canal station''s serial number';
COMMENT ON COLUMN public.canal_waterlevel.canal_waterlevel_datetime IS 'วันที่วัดระดับน้ำในคลอง record date';
COMMENT ON COLUMN public.canal_waterlevel.canal_waterlevel_value IS 'ค่าระดับน้ำในคลอง canal water level value';
COMMENT ON COLUMN public.canal_waterlevel.comm_status IS 'สถานะของเครื่องวัด meter status';
COMMENT ON COLUMN public.canal_waterlevel.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.canal_waterlevel.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.canal_waterlevel.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.canal_waterlevel.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.canal_waterlevel.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.canal_waterlevel.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.canal_waterlevel.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.canal_waterlevel.deleted_at IS 'วันที่ลบข้อมูล deleted date';

