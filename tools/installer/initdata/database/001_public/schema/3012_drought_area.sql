-- Table: public.drought_area

-- DROP TABLE public.drought_area;

CREATE TABLE public.drought_area
(
  id bigserial NOT NULL, -- รหัสพื้นที่ภัยแล้ง
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  drought_datetime timestamp with time zone, -- วันที่และเวลาประกาศพื้นที่ประกาศภัยแล้ง
  remark text, -- หมายเหตุ remark
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_drought_area PRIMARY KEY (id),
  CONSTRAINT fk_drought__reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_drought__reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_drought_area UNIQUE (geocode_id, agency_id, drought_datetime, deleted_at),
  CONSTRAINT pt_drought_area_drought_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.drought_area
  IS 'พื้นที่ภัยแล้ง';
COMMENT ON COLUMN public.drought_area.id IS 'รหัสพื้นที่ภัยแล้ง';
COMMENT ON COLUMN public.drought_area.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.drought_area.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.drought_area.drought_datetime IS 'วันที่และเวลาประกาศพื้นที่ประกาศภัยแล้ง';
COMMENT ON COLUMN public.drought_area.remark IS 'หมายเหตุ remark';
COMMENT ON COLUMN public.drought_area.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.drought_area.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.drought_area.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.drought_area.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.drought_area.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.drought_area.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.drought_area.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.drought_area.deleted_at IS 'วันที่ลบข้อมูล deleted date';

