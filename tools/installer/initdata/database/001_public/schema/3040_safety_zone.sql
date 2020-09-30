-- Table: public.safety_zone

-- DROP TABLE public.safety_zone;

CREATE TABLE public.safety_zone
(
  id bigserial NOT NULL, -- รหัสจุดปลอดภัยจากดินถล่ม safety zone's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  safety_zone_name text, -- พื้นที่จุดปลอดภัยจากดินถล่ม safety zone's name
  safety_zone_oldcode text, -- รหัสเดิมของจุดปลอดภัยจากดินถล่ม old safety zone's serial number
  mooban_name text, -- ชือหมู่บ้าน address name
  mooban_id text, -- หมู่ที่ address number
  safety_zone_lat numeric(9,6), -- พิกัดละติจูด latitude
  safety_zone_long numeric(9,6), -- พิกัดลองติจูด longitude
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_safety_zone PRIMARY KEY (id),
  CONSTRAINT fk_safety_z_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_safety_z_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_safety_zone UNIQUE (safety_zone_oldcode, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.safety_zone
  IS 'จุดปลอดภัยจากดินถล่ม';
COMMENT ON COLUMN public.safety_zone.id IS 'รหัสจุดปลอดภัยจากดินถล่ม safety zone''s serial number';
COMMENT ON COLUMN public.safety_zone.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.safety_zone.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.safety_zone.safety_zone_name IS 'พื้นที่จุดปลอดภัยจากดินถล่ม safety zone''s name';
COMMENT ON COLUMN public.safety_zone.safety_zone_oldcode IS 'รหัสเดิมของจุดปลอดภัยจากดินถล่ม old safety zone''s serial number';
COMMENT ON COLUMN public.safety_zone.mooban_name IS 'ชือหมู่บ้าน address name';
COMMENT ON COLUMN public.safety_zone.mooban_id IS 'หมู่ที่ address number';
COMMENT ON COLUMN public.safety_zone.safety_zone_lat IS 'พิกัดละติจูด latitude';
COMMENT ON COLUMN public.safety_zone.safety_zone_long IS 'พิกัดลองติจูด longitude';
COMMENT ON COLUMN public.safety_zone.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.safety_zone.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.safety_zone.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.safety_zone.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.safety_zone.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.safety_zone.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.safety_zone.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.safety_zone.deleted_at IS 'วันที่ลบข้อมูล deleted date';

