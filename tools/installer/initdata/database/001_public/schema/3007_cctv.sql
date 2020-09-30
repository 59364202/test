-- Table: public.cctv

-- DROP TABLE public.cctv;

CREATE TABLE public.cctv
(
  id bigserial NOT NULL, -- รหัสข้อมูล cctv
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  dam_id text, -- รหัสเขื่อน
  basin_id text, -- รหัสลุ่มน้ำ
  cctv_lat numeric(9,6), -- พิกัดละติจูด
  cctv_long numeric(9,6), -- พิกัดลองติจูด
  cctv_flash text, -- วิดีโอ flash
  cctv_quicktime text, -- วิดีโอ quick time
  cctv_title text, -- ชื่อสถานที่
  cctv_description text, -- คำอธิบายสถานที่
  cctv_strem text, -- ชื่อวิดีโอสตรีม
  tele_station_id text, -- รหัสสถานีโทรมาตร...
  cctv_mediatype text, -- ประเภทข้อมูล
  cctv_url text, -- ที่อยู่วิดีโอ
  cctv_filename text, -- ชื่อไฟล์
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_cctv PRIMARY KEY (id),
  CONSTRAINT uk_cctv UNIQUE (cctv_url, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.cctv
  IS 'ข้อมูลพื้นฐาน cctv';
COMMENT ON COLUMN public.cctv.id IS 'รหัสข้อมูล cctv';
COMMENT ON COLUMN public.cctv.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.cctv.dam_id IS 'รหัสเขื่อน';
COMMENT ON COLUMN public.cctv.basin_id IS 'รหัสลุ่มน้ำ';
COMMENT ON COLUMN public.cctv.cctv_lat IS 'พิกัดละติจูด';
COMMENT ON COLUMN public.cctv.cctv_long IS 'พิกัดลองติจูด';
COMMENT ON COLUMN public.cctv.cctv_flash IS 'วิดีโอ flash';
COMMENT ON COLUMN public.cctv.cctv_quicktime IS 'วิดีโอ quick time';
COMMENT ON COLUMN public.cctv.cctv_title IS 'ชื่อสถานที่';
COMMENT ON COLUMN public.cctv.cctv_description IS 'คำอธิบายสถานที่';
COMMENT ON COLUMN public.cctv.cctv_strem IS 'ชื่อวิดีโอสตรีม';
COMMENT ON COLUMN public.cctv.tele_station_id IS 'รหัสสถานีโทรมาตรรหัส';
COMMENT ON COLUMN public.cctv.cctv_mediatype IS 'ประเภทข้อมูล';
COMMENT ON COLUMN public.cctv.cctv_url IS 'ที่อยู่วิดีโอ';
COMMENT ON COLUMN public.cctv.cctv_filename IS 'ชื่อไฟล์';
COMMENT ON COLUMN public.cctv.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.cctv.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.cctv.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.cctv.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.cctv.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.cctv.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.cctv.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.cctv.deleted_at IS 'วันที่ลบข้อมูล deleted date';