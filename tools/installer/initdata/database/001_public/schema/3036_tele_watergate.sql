-- Table: public.tele_watergate

-- DROP TABLE public.tele_watergate;

CREATE TABLE public.tele_watergate
(
  id bigserial NOT NULL, -- รหัสค่าระดับน้ำจากการวัดของสถานี
  tele_station_id bigint, -- รหัสสถานีโทรมาตร tele station's serial number
  watergate_datetime timestamp with time zone NOT NULL, -- วันที่ตรวจสอบค่าระดับน้ำ
  watergate_in double precision, -- ระดับน้ำด้านในประตูระบายน้ำ
  watergate_out double precision, -- ระดับน้ำนอกประตูระบายน้ำ
  watergate_out2 double precision, -- ระดับน้ำนอกประตูระบายน้ำ
  pump_on bigint, -- จำนวนเครื่องสูบน้ำที่ใช้งาน (เครื่อง)
  floodgate_open bigint, -- จำนวนประตูระบายน้ำที่เปิด (บาน)
  floodgate_height double precision, -- จำนวนความสูงของประตูระบายน้ำ (เมตร
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_tele_watergate PRIMARY KEY (id),
  CONSTRAINT fk_tele_wat_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_tele_watergate UNIQUE (tele_station_id, watergate_datetime, deleted_at),
  CONSTRAINT pt_tele_watergate_watergate_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.tele_watergate
  IS 'ค่าระดับน้ำของฝายและประตูน้ำจากโทรมาตร';
COMMENT ON COLUMN public.tele_watergate.id IS 'รหัสค่าระดับน้ำจากการวัดของสถานี';
COMMENT ON COLUMN public.tele_watergate.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.tele_watergate.watergate_datetime IS 'วันที่ตรวจสอบค่าระดับน้ำ';
COMMENT ON COLUMN public.tele_watergate.watergate_in IS 'ระดับน้ำด้านในประตูระบายน้ำ';
COMMENT ON COLUMN public.tele_watergate.watergate_out IS 'ระดับน้ำนอกประตูระบายน้ำ';
COMMENT ON COLUMN public.tele_watergate.watergate_out2 IS 'ระดับน้ำนอกประตูระบายน้ำ';
COMMENT ON COLUMN public.tele_watergate.pump_on IS 'จำนวนเครื่องสูบน้ำที่ใช้งาน (เครื่อง)';
COMMENT ON COLUMN public.tele_watergate.floodgate_open IS 'จำนวนประตูระบายน้ำที่เปิด (บาน)';
COMMENT ON COLUMN public.tele_watergate.floodgate_height IS 'จำนวนความสูงของประตูระบายน้ำ (เมตร';
COMMENT ON COLUMN public.tele_watergate.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.tele_watergate.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.tele_watergate.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.tele_watergate.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.tele_watergate.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.tele_watergate.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.tele_watergate.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.tele_watergate.deleted_at IS 'วันที่ลบข้อมูล deleted date';

