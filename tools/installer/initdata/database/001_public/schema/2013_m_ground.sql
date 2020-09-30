-- Table: public.m_ground

-- DROP TABLE public.m_ground;

CREATE TABLE public.m_ground
(
  id bigserial NOT NULL, -- รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
  ground_oldcode text, -- รหัสเดิมของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  well_ownner text, -- เจ้าของบ่อน้ำบาดาล
  mooban_id text, -- รหัสหมู่บ้าน
  mooban text, -- ชื่อหมู่บ้าน
  ground_lat numeric(9,6), -- ละติจูดแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_long numeric(9,6), -- ลองติจูดแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_location text, -- สถานที่ของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  map_sheet text, -- ระวางแผนที่
  map_zone text, -- เส้นกริดแบ่งโซน
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_ground PRIMARY KEY (id),
  CONSTRAINT fk_m_ground_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_ground UNIQUE (ground_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_ground
  IS 'ข้อมูลพื้นฐานของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.id IS 'รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_ground.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.m_ground.ground_oldcode IS 'รหัสเดิมของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.well_ownner IS 'เจ้าของบ่อน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.mooban_id IS 'รหัสหมู่บ้าน';
COMMENT ON COLUMN public.m_ground.mooban IS 'ชื่อหมู่บ้าน';
COMMENT ON COLUMN public.m_ground.ground_lat IS 'ละติจูดแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.ground_long IS 'ลองติจูดแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.ground_location IS 'สถานที่ของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.m_ground.map_sheet IS 'ระวางแผนที่';
COMMENT ON COLUMN public.m_ground.map_zone IS 'เส้นกริดแบ่งโซน';
COMMENT ON COLUMN public.m_ground.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_ground.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_ground.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_ground.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_ground.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_ground.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_ground.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_ground.deleted_at IS 'วันที่ลบข้อมูล deleted date';

