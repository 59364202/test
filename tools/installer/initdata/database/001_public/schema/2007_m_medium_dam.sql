-- Table: public.m_medium_dam

-- DROP TABLE public.m_medium_dam;

CREATE TABLE public.m_medium_dam
(
  id bigserial NOT NULL, -- รหัสข้อมูลเขื่อนขนาดกลาง mediumdam's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  subbasin_id bigint, -- ลำดับลุ่มน้ำสาขา subbasin's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย
  mediumdam_name json, -- ชื่อเขื่อน
  mediumdam_oldcode text, -- รหัสเดิมของเขื่อนขนาดใหญ่ old dam's serial number
  mediumdam_lat numeric(9,6), -- พิกัดของเขื่อน latitude
  mediumdam_long numeric(9,6), -- พิกัดของเขื่อน longitude
  normal_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.] normal storage
  min_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด [ล้าน ลบ.ม.] min storage
  max_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักสูงสุด [ล้าน ลบ.ม.] max storage
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_medium_dam PRIMARY KEY (id),
  CONSTRAINT fk_m_medium_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_medium_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_medium_reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_medium_dam UNIQUE (mediumdam_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_medium_dam
  IS 'ข้อมูลพืนฐานเขื่อนขนาดกลาง';
COMMENT ON COLUMN public.m_medium_dam.id IS 'รหัสข้อมูลเขื่อนขนาดกลาง mediumdam''s serial number';
COMMENT ON COLUMN public.m_medium_dam.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.m_medium_dam.subbasin_id IS 'ลำดับลุ่มน้ำสาขา subbasin''s serial number';
COMMENT ON COLUMN public.m_medium_dam.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.m_medium_dam.mediumdam_name IS 'ชื่อเขื่อน';
COMMENT ON COLUMN public.m_medium_dam.mediumdam_oldcode IS 'รหัสเดิมของเขื่อนขนาดใหญ่ old dam''s serial number';
COMMENT ON COLUMN public.m_medium_dam.mediumdam_lat IS 'พิกัดของเขื่อน latitude';
COMMENT ON COLUMN public.m_medium_dam.mediumdam_long IS 'พิกัดของเขื่อน longitude';
COMMENT ON COLUMN public.m_medium_dam.normal_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.] normal storage';
COMMENT ON COLUMN public.m_medium_dam.min_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด [ล้าน ลบ.ม.] min storage';
COMMENT ON COLUMN public.m_medium_dam.max_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักสูงสุด [ล้าน ลบ.ม.] max storage';
COMMENT ON COLUMN public.m_medium_dam.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_medium_dam.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_medium_dam.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_medium_dam.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_medium_dam.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_medium_dam.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_medium_dam.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_medium_dam.deleted_at IS 'วันที่ลบข้อมูล deleted date';

