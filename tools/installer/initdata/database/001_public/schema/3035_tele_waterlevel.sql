-- Table: public.tele_waterlevel

-- DROP TABLE public.tele_waterlevel;

CREATE TABLE public.tele_waterlevel
(
  id bigserial NOT NULL, -- รหัสค่าระดับน้ำจากการวัดของสถานี
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  waterlevel_datetime timestamp with time zone NOT NULL, -- วันที่ตรวจสอบค่าระดับน้ำ
  waterlevel_m double precision, -- ระดับน้ำ เมตร
  waterlevel_msl double precision, -- ระดับน้ำ รทก
  flow_rate double precision, -- อัตราการไหล
  discharge double precision, -- ปริมาณการระบายน้ำ discharge
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_tele_waterlevel PRIMARY KEY (id),
  CONSTRAINT fk_tele_wat_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_tele_waterlevel UNIQUE (tele_station_id, waterlevel_datetime, deleted_at),
  CONSTRAINT pt_tele_waterlevel_waterlevel_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.tele_waterlevel
  IS 'ค่าระดับน้ำจากโทรมาตร';
COMMENT ON COLUMN public.tele_waterlevel.id IS 'รหัสค่าระดับน้ำจากการวัดของสถานี';
COMMENT ON COLUMN public.tele_waterlevel.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.tele_waterlevel.waterlevel_datetime IS 'วันที่ตรวจสอบค่าระดับน้ำ';
COMMENT ON COLUMN public.tele_waterlevel.waterlevel_m IS 'ระดับน้ำ เมตร';
COMMENT ON COLUMN public.tele_waterlevel.waterlevel_msl IS 'ระดับน้ำ รทก';
COMMENT ON COLUMN public.tele_waterlevel.flow_rate IS 'อัตราการไหล';
COMMENT ON COLUMN public.tele_waterlevel.discharge IS 'ปริมาณการระบายน้ำ discharge';
COMMENT ON COLUMN public.tele_waterlevel.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.tele_waterlevel.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.tele_waterlevel.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.tele_waterlevel.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.tele_waterlevel.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.tele_waterlevel.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.tele_waterlevel.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.tele_waterlevel.deleted_at IS 'วันที่ลบข้อมูล deleted date';

