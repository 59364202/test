-- Table: public.m_tele_station

-- DROP TABLE public.m_tele_station;

CREATE TABLE public.m_tele_station
(
  id bigserial NOT NULL, -- รหัสสถานีโทรมาตร tele station's serial number
  subbasin_id bigint, -- รหัสลุ่มน้ำสาขา subbasin number
  agency_id bigint, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode  number
  tele_station_name json, -- ชื่อสถานีโทรมาตร tele station's name
  tele_station_lat numeric(9,6), -- ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
  tele_station_long numeric(9,6), -- ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
  tele_station_oldcode text, -- รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
  tele_station_type text, -- ชนิดของโทรมาตร (เช่น ระดับน้ำ)
  left_bank double precision, -- ระดับตลิ่ง (ซ้าย) left bank level
  right_bank double precision, -- ระดับตลิ่ง (ขวา) right bank level
  ground_level double precision, -- ระดับท้องน้ำ ม.รทก ground water level
  riverbank double precision, -- ความสูงของตลิ่ง (เมตร) riverbank height
  water_storage_station double precision, -- ความจุ (ล.ม. ต่อวินาที) water storage
  max_waterlevel_20y double precision, -- ระดับน้ำสูงสุดย้อนหลัง 20 ปี (เมตร) last 20 years of max water level
  sort_order bigint, -- เรียงลำดับลุ่มน้ำ
  pump bigint, -- จำนวนของเครื่องสูบน้ำ (เครื่อง)
  floodgate bigint, -- จำนวนของประตูระบายน้ำ (บาน)
  waterflow_time double precision, -- ระยะทางเดินของน้ำในแต่ละสถานี หน่วยเป็นชั่วโมง
  distance double precision, -- ระบะทางการไหลของน้ำระหว่างสถานี
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_tele_station PRIMARY KEY (id),
  CONSTRAINT uk_m_tele_station UNIQUE (tele_station_oldcode, agency_id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_tele_station
  IS 'ข้อมูลพื้นฐานของโทรมาตร';
COMMENT ON COLUMN public.m_tele_station.id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.m_tele_station.subbasin_id IS 'รหัสลุ่มน้ำสาขา subbasin number';
COMMENT ON COLUMN public.m_tele_station.agency_id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number';
COMMENT ON COLUMN public.m_tele_station.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode  number';
COMMENT ON COLUMN public.m_tele_station.tele_station_name IS 'ชื่อสถานีโทรมาตร tele station''s name';
COMMENT ON COLUMN public.m_tele_station.tele_station_lat IS 'ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude';
COMMENT ON COLUMN public.m_tele_station.tele_station_long IS 'ลองติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude';
COMMENT ON COLUMN public.m_tele_station.tele_station_oldcode IS 'รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number';
COMMENT ON COLUMN public.m_tele_station.tele_station_type IS 'ชนิดของโทรมาตร (เช่น ระดับน้ำ)';
COMMENT ON COLUMN public.m_tele_station.left_bank IS 'ระดับตลิ่ง (ซ้าย) left bank level';
COMMENT ON COLUMN public.m_tele_station.right_bank IS 'ระดับตลิ่ง (ขวา) right bank level';
COMMENT ON COLUMN public.m_tele_station.ground_level IS 'ระดับท้องน้ำ ม.รทก ground water level';
COMMENT ON COLUMN public.m_tele_station.riverbank IS 'ความสูงของตลิ่ง (เมตร) riverbank height';
COMMENT ON COLUMN public.m_tele_station.water_storage_station IS 'ความจุ (ล.ม. ต่อวินาที) water storage';
COMMENT ON COLUMN public.m_tele_station.max_waterlevel_20y IS 'ระดับน้ำสูงสุดย้อนหลัง 20 ปี (เมตร) last 20 years of max water level';
COMMENT ON COLUMN public.m_tele_station.sort_order IS 'เรียงลำดับลุ่มน้ำ';
COMMENT ON COLUMN public.m_tele_station.pump IS 'จำนวนของเครื่องสูบน้ำ (เครื่อง)';
COMMENT ON COLUMN public.m_tele_station.floodgate IS 'จำนวนของประตูระบายน้ำ (บาน)';
COMMENT ON COLUMN public.m_tele_station.waterflow_time IS 'ระยะทางเดินของน้ำในแต่ละสถานี หน่วยเป็นชั่วโมง';
COMMENT ON COLUMN public.m_tele_station.distance IS 'ระบะทางการไหลของน้ำระหว่างสถานี';
COMMENT ON COLUMN public.m_tele_station.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่';
COMMENT ON COLUMN public.m_tele_station.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_tele_station.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_tele_station.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_tele_station.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_tele_station.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_tele_station.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_tele_station.deleted_at IS 'วันที่ลบข้อมูล deleted date';

