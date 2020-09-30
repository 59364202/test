-- Table: public.m_ford_station

-- DROP TABLE public.m_ford_station;

CREATE TABLE public.m_ford_station
(
  id bigserial NOT NULL, -- รหัสสถานีวัดระดับน้ำบนถนน ford station's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
  agency_id bigint, -- รหัสหน่วยงาน
  ford_station_name json, -- ชื่อสถานีวัดระดับน้ำระดับถนน
  ford_station_lat numeric(9,6), -- ละติจูดสถานีวัดระดับน้ำระดับถนน
  ford_station_long numeric(9,6), -- ลองติจูดสถานีวัดระดับน้ำระดับถนน
  ford_station_oldcode text, -- รหัสเดิมของสถานีวัดระดับน้ำระดับถนน
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_ford_station PRIMARY KEY (id),
  CONSTRAINT fk_m_ford_s_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_ford_station UNIQUE (ford_station_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_ford_station
  IS 'ข้อมูลพื้นฐานของเครื่องวัดน้ำท่วมถนน';
COMMENT ON COLUMN public.m_ford_station.id IS 'รหัสสถานีวัดระดับน้ำบนถนน ford station''s serial number';
COMMENT ON COLUMN public.m_ford_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.m_ford_station.agency_id IS 'รหัสหน่วยงาน ';
COMMENT ON COLUMN public.m_ford_station.ford_station_name IS 'ชื่อสถานีวัดระดับน้ำระดับถนน';
COMMENT ON COLUMN public.m_ford_station.ford_station_lat IS 'ละติจูดสถานีวัดระดับน้ำระดับถนน';
COMMENT ON COLUMN public.m_ford_station.ford_station_long IS 'ลองติจูดสถานีวัดระดับน้ำระดับถนน';
COMMENT ON COLUMN public.m_ford_station.ford_station_oldcode IS 'รหัสเดิมของสถานีวัดระดับน้ำระดับถนน';
COMMENT ON COLUMN public.m_ford_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_ford_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_ford_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_ford_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_ford_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_ford_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_ford_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_ford_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

