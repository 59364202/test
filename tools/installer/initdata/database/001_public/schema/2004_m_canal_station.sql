-- Table: public.m_canal_station

-- DROP TABLE public.m_canal_station;

CREATE TABLE public.m_canal_station
(
  id bigserial NOT NULL, -- รหัสสถานีวัดระดับน้ำในคลอง canal station's serial number
  subbasin_id bigint, -- รหัสลุ่มน้ำสาขา basin's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  canal_station_name json, -- ชื่อสถานีวัดระดับน้ำในคลอง canal station's name
  canal_station_lat numeric(9,6), -- ละติจูดสถานีวัดระดับน้ำในคลอง latitude
  canal_station_long numeric(9,6), -- ลองติจูดสถานีวัดระดับน้ำในคลอง longitude
  canal_station_oldcode text, -- รหัสเดิมของสถานีวัดระดับน้ำในคลอง old canal station's code
  sort_order bigint, -- เพิ่มเรียงลำดับสถานีของคลอง
  distance double precision, -- ระบะทางการไหลของน้ำระหว่างสถานี
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_canal_station PRIMARY KEY (id),
  CONSTRAINT fk_m_canal__reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_canal__reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_canal_station UNIQUE (canal_station_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_canal_station
  IS 'ข้อมูลพื้นฐานของคลอง';
COMMENT ON COLUMN public.m_canal_station.id IS 'รหัสสถานีวัดระดับน้ำในคลอง canal station''s serial number';
COMMENT ON COLUMN public.m_canal_station.subbasin_id IS 'รหัสลุ่มน้ำสาขา basin''s serial number';
COMMENT ON COLUMN public.m_canal_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.m_canal_station.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_canal_station.canal_station_name IS 'ชื่อสถานีวัดระดับน้ำในคลอง canal station''s name';
COMMENT ON COLUMN public.m_canal_station.canal_station_lat IS 'ละติจูดสถานีวัดระดับน้ำในคลอง latitude';
COMMENT ON COLUMN public.m_canal_station.canal_station_long IS 'ลองติจูดสถานีวัดระดับน้ำในคลอง longitude';
COMMENT ON COLUMN public.m_canal_station.canal_station_oldcode IS 'รหัสเดิมของสถานีวัดระดับน้ำในคลอง old canal station''s code';
COMMENT ON COLUMN public.m_canal_station.sort_order IS 'เพิ่มเรียงลำดับสถานีของคลอง';
COMMENT ON COLUMN public.m_canal_station.distance IS 'ระบะทางการไหลของน้ำระหว่างสถานี ';
COMMENT ON COLUMN public.m_canal_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_canal_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_canal_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_canal_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_canal_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_canal_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_canal_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_canal_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

