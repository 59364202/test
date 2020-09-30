-- Table: public.m_swan_station

-- DROP TABLE public.m_swan_station;

CREATE TABLE public.m_swan_station
(
  id bigserial NOT NULL, -- รหัสสถานีสถานีคาดการณ์ความสูงคลื่น
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
  swan_name json, -- ชื่อสถานีโทรมาตร tele station's name
  swan_lat numeric(9,6), -- ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
  swan_long numeric(9,6), -- ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
  swan_oldcode text, -- รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station's serial number
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_swan_station PRIMARY KEY (id),
  CONSTRAINT fk_m_swan_s_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_swan_s_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_swan_station UNIQUE (swan_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_swan_station
  IS 'ข้อมูลพื้นฐานของสถานีคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.m_swan_station.id IS 'รหัสสถานีสถานีคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.m_swan_station.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_swan_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.m_swan_station.swan_name IS 'ชื่อสถานีโทรมาตร tele station''s name';
COMMENT ON COLUMN public.m_swan_station.swan_lat IS 'ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude';
COMMENT ON COLUMN public.m_swan_station.swan_long IS 'ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude';
COMMENT ON COLUMN public.m_swan_station.swan_oldcode IS 'รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station''s serial number';
COMMENT ON COLUMN public.m_swan_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่';
COMMENT ON COLUMN public.m_swan_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_swan_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_swan_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_swan_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_swan_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_swan_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_swan_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

