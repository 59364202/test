-- Table: public.m_floodforecast_station

-- DROP TABLE public.m_floodforecast_station;

CREATE TABLE public.m_floodforecast_station
(
  id bigserial NOT NULL, -- รหัสสถานีสถานีคาดการณ์ระดับน้ำท่วม
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
  subbasin_id bigint, -- รหัสลุ่มน้ำสาขา subbasin's serial number
  floodforecast_station_name json, -- ชื่อสถานีโทรมาตร tele station's name
  floodforecast_station_lat numeric(9,6), -- ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
  floodforecast_station_long numeric(9,6), -- ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
  floodforecast_station_oldcode text, -- รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station's serial number
  floodforecast_station_type text, -- ชนิดของโทรมาตร (เช่น ระดับน้ำ)
  floodforecast_station_alarm double precision, -- ระดับเตือนภัย
  floodforecast_station_warning double precision, -- ระดับเตือนภัย
  floodforecast_station_critical double precision, -- ระดับเตือนภัย
  floodforecast_station_unit text, -- หน่วยของข้อมูล
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_floodforecast_station PRIMARY KEY (id),
  CONSTRAINT fk_m_floodf_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_floodf_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_floodf_reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_floodforcast_station UNIQUE (floodforecast_station_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_floodforecast_station
  IS 'ข้อมูลพื้นฐานของสถานีคาดการณ์ระดับน้ำท่วม';
COMMENT ON COLUMN public.m_floodforecast_station.id IS 'รหัสสถานีสถานีคาดการณ์ระดับน้ำท่วม';
COMMENT ON COLUMN public.m_floodforecast_station.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_floodforecast_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.m_floodforecast_station.subbasin_id IS 'รหัสลุ่มน้ำสาขา subbasin''s serial number';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_name IS 'ชื่อสถานีโทรมาตร tele station''s name';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_lat IS 'ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_long IS 'ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_oldcode IS 'รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station''s serial number';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_type IS 'ชนิดของโทรมาตร (เช่น ระดับน้ำ)';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_alarm IS 'ระดับเตือนภัย';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_warning IS 'ระดับเตือนภัย';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_critical IS 'ระดับเตือนภัย';
COMMENT ON COLUMN public.m_floodforecast_station.floodforecast_station_unit IS 'หน่วยของข้อมูล';
COMMENT ON COLUMN public.m_floodforecast_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_floodforecast_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_floodforecast_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_floodforecast_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_floodforecast_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_floodforecast_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_floodforecast_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_floodforecast_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

