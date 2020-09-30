-- Table: public.geohazard_situation

-- DROP TABLE public.geohazard_situation;

CREATE TABLE public.geohazard_situation
(
  id bigserial NOT NULL, -- รหัสสถานการณ์ธรณีพิบัติภัย geohazard's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  agency_id bigint, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
  geohazard_datetime timestamp with time zone NOT NULL, -- วันและเวลาที่ประกาศสถานการณ์พิบัติภัย DateTime of geohazard
  geohazard_name text, -- ชื่อสถานการณ์พิบัติภัย name of geohazard
  geohazard_link text, -- ลิ้งที่แสดงสถานการณ์พิบัติภัย geohazard link
  geohazard_description text, -- รายละเอียดสถานการณ์พิบัติภัย description of geohazard
  geohazard_author text, -- ผู้รายงานสถานการณ์ author
  geohazard_colorlevel text, -- ระดับสีของเกณฑ์พิบัติภัย color level of geohazard
  geohazard_remark text, -- หมายเหตุ Remark
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 00:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_geohazard_situation PRIMARY KEY (id),
  CONSTRAINT fk_geohazar_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_geohazar_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_geohazard_situation UNIQUE (geohazard_link, deleted_at, geohazard_datetime),
  CONSTRAINT pt_geohazard_situation_geohazard_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.geohazard_situation
  IS 'สถานการณ์ธรณีพิบัติภัย';
COMMENT ON COLUMN public.geohazard_situation.id IS 'รหัสสถานการณ์ธรณีพิบัติภัย geohazard''s serial number';
COMMENT ON COLUMN public.geohazard_situation.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.geohazard_situation.agency_id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency''s serial number';
COMMENT ON COLUMN public.geohazard_situation.geohazard_datetime IS 'วันและเวลาที่ประกาศสถานการณ์พิบัติภัย DateTime of geohazard';
COMMENT ON COLUMN public.geohazard_situation.geohazard_name IS 'ชื่อสถานการณ์พิบัติภัย name of geohazard';
COMMENT ON COLUMN public.geohazard_situation.geohazard_link IS 'ลิ้งที่แสดงสถานการณ์พิบัติภัย geohazard link';
COMMENT ON COLUMN public.geohazard_situation.geohazard_description IS 'รายละเอียดสถานการณ์พิบัติภัย description of geohazard ';
COMMENT ON COLUMN public.geohazard_situation.geohazard_author IS 'ผู้รายงานสถานการณ์ author';
COMMENT ON COLUMN public.geohazard_situation.geohazard_colorlevel IS 'ระดับสีของเกณฑ์พิบัติภัย color level of geohazard';
COMMENT ON COLUMN public.geohazard_situation.geohazard_remark IS 'หมายเหตุ Remark';
COMMENT ON COLUMN public.geohazard_situation.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.geohazard_situation.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.geohazard_situation.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.geohazard_situation.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.geohazard_situation.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.geohazard_situation.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.geohazard_situation.deleted_at IS 'วันที่ลบข้อมูล deleted date';

