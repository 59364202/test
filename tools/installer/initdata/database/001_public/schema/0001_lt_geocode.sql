-- Table: public.lt_geocode

-- DROP TABLE public.lt_geocode;

CREATE TABLE public.lt_geocode
(
  id bigserial NOT NULL, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  geocode text, -- รหัสข้อมูลขอบเขตการปกครองของประเทศไทย
  area_code text, -- รหัสภาคของประเทศไทย
  province_code text, -- รหัสจังหวัดของประเทศไทย
  amphoe_code text, -- รหัสอำเภอของประเทศไทย
  tumbon_code text, -- รหัสตำบลของประเทศไทย
  area_name json, -- ชื่อภาคของประเทศไทย
  province_name json, -- ชื่อจังหวัดของประเทศไทย
  amphoe_name json, -- ชื่อตำบลของประเทศไทย
  tumbon_name json, -- ชื่อตำบลของประเทศไทย
  warning_zone text, -- โซนเตือนภัยพิจารณาตามพื้นที่
  zone_detail text, -- รายละเอียดของ Zone เตือนภัย
  tmd_area_code text, -- รหัสภาคของประเทศไทยตามกรมอุตุนิยมวิทยา
  tmd_area_name text, -- ชื่อภาคของประเทศไทยตามกรมอุตุนิยมวิทยา
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_geocode PRIMARY KEY (id),
  CONSTRAINT uk_lt_geocode UNIQUE (geocode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_geocode
  IS 'ขอบเขตการปกครองของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.lt_geocode.geocode IS 'รหัสข้อมูลขอบเขตการปกครองของประเทศไทย ';
COMMENT ON COLUMN public.lt_geocode.area_code IS 'รหัสภาคของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.province_code IS 'รหัสจังหวัดของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.amphoe_code IS 'รหัสอำเภอของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.tumbon_code IS 'รหัสตำบลของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.area_name IS 'ชื่อภาคของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.province_name IS 'ชื่อจังหวัดของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.amphoe_name IS 'ชื่อตำบลของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.tumbon_name IS 'ชื่อตำบลของประเทศไทย';
COMMENT ON COLUMN public.lt_geocode.warning_zone IS 'โซนเตือนภัยพิจารณาตามพื้นที่';
COMMENT ON COLUMN public.lt_geocode.zone_detail IS 'รายละเอียดของ Zone เตือนภัย';
COMMENT ON COLUMN public.lt_geocode.tmd_area_code IS 'รหัสภาคของประเทศไทยตามกรมอุตุนิยมวิทยา';
COMMENT ON COLUMN public.lt_geocode.tmd_area_name IS 'ชื่อภาคของประเทศไทยตามกรมอุตุนิยมวิทยา';
COMMENT ON COLUMN public.lt_geocode.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_geocode.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_geocode.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_geocode.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_geocode.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_geocode.deleted_at IS 'วันที่ลบข้อมูล deleted date';
