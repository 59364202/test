-- Table: public.m_waterquality_station

-- DROP TABLE public.m_waterquality_station;

CREATE TABLE public.m_waterquality_station
(
  id bigserial NOT NULL, -- รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ water quality number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode number
  agency_id bigint, -- รหัสหน่วยงาน agency number
  subbasin_id bigint, -- ลำดับลุ่มน้ำสาขา subbasin number
  waterquality_station_name json, -- ชื่อสถานีโทรมาตร water station's name
  waterquality_station_lat numeric(9,6), -- ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
  waterquality_station_long numeric(9,6), -- ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
  waterquality_station_oldcode text, -- รหัสเครื่องตรวจวัดคุณภาพน้ำอัตโนมัติเดิม old water station number
  sort_order bigint, -- เรียงลำดับลุ่มน้ำ
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 00:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_waterquality_station PRIMARY KEY (id),
  CONSTRAINT fk_m_waterq_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_waterq_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_waterq_reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_waterquality_station UNIQUE (agency_id, waterquality_station_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_waterquality_station
  IS 'ข้อมูลพื้นฐานสถานีวัดคุณภาพน้ำอัตโนมัติ';
COMMENT ON COLUMN public.m_waterquality_station.id IS 'รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ water quality number';
COMMENT ON COLUMN public.m_waterquality_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode number';
COMMENT ON COLUMN public.m_waterquality_station.agency_id IS 'รหัสหน่วยงาน agency number';
COMMENT ON COLUMN public.m_waterquality_station.subbasin_id IS 'ลำดับลุ่มน้ำสาขา subbasin number';
COMMENT ON COLUMN public.m_waterquality_station.waterquality_station_name IS 'ชื่อสถานีโทรมาตร water station''s name';
COMMENT ON COLUMN public.m_waterquality_station.waterquality_station_lat IS 'ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude';
COMMENT ON COLUMN public.m_waterquality_station.waterquality_station_long IS 'ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude';
COMMENT ON COLUMN public.m_waterquality_station.waterquality_station_oldcode IS 'รหัสเครื่องตรวจวัดคุณภาพน้ำอัตโนมัติเดิม old water station number';
COMMENT ON COLUMN public.m_waterquality_station.sort_order IS 'เรียงลำดับลุ่มน้ำ';
COMMENT ON COLUMN public.m_waterquality_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่';
COMMENT ON COLUMN public.m_waterquality_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_waterquality_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_waterquality_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_waterquality_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_waterquality_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_waterquality_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_waterquality_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

