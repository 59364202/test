-- Table: public.landslide_area

-- DROP TABLE public.landslide_area;

CREATE TABLE public.landslide_area
(
  id bigserial NOT NULL, -- รหัสพื้นที่ดินถล่ม landslide area's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  subbasin_id bigint, -- รหัสลุ่มน้ำสาขา basin's serial number
  mooban_no text, -- หมู่ที่ address number
  mooban_name text, -- ชื่อหมู่บ้าน address name
  landslide_area_oldcode text NOT NULL, -- รหัสเดิมพื้นที่ดินถล่ม landslide area code
  landslide_area_lat numeric(9,6), -- พิกัดละติจูด latitude
  landslide_area_long numeric(9,6), -- พิกัดลองติจูด longitude
  contact_name text, -- ชื่อบุคคลติดต่อ contact name
  contact_tel1 text, -- เบอร์โทรศัพท์บุคคลติดต่อ1 contact number (1)
  contact_tel2 text, -- เบอร์โทรศัพท์บุคคลติดต่อ2 contact number (2)
  risk_level text, -- ระดับความเสี่ยง risk level
  risk_detail text, -- รายละเอียดความเสี่ยงที่เกิดขึ้น risk detail
  remark text, -- หมายเหตุ remark
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_landslide_area PRIMARY KEY (id),
  CONSTRAINT fk_landslid_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_landslid_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_landslid_reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_landslide_area UNIQUE (landslide_area_oldcode, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.landslide_area
  IS 'พื้นที่ดินถล่ม';
COMMENT ON COLUMN public.landslide_area.id IS 'รหัสพื้นที่ดินถล่ม landslide area''s serial number';
COMMENT ON COLUMN public.landslide_area.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.landslide_area.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.landslide_area.subbasin_id IS 'รหัสลุ่มน้ำสาขา basin''s serial number';
COMMENT ON COLUMN public.landslide_area.mooban_no IS 'หมู่ที่ address number';
COMMENT ON COLUMN public.landslide_area.mooban_name IS 'ชื่อหมู่บ้าน address name';
COMMENT ON COLUMN public.landslide_area.landslide_area_oldcode IS 'รหัสเดิมพื้นที่ดินถล่ม landslide area code';
COMMENT ON COLUMN public.landslide_area.landslide_area_lat IS 'พิกัดละติจูด latitude';
COMMENT ON COLUMN public.landslide_area.landslide_area_long IS 'พิกัดลองติจูด longitude';
COMMENT ON COLUMN public.landslide_area.contact_name IS 'ชื่อบุคคลติดต่อ contact name';
COMMENT ON COLUMN public.landslide_area.contact_tel1 IS 'เบอร์โทรศัพท์บุคคลติดต่อ1 contact number (1)';
COMMENT ON COLUMN public.landslide_area.contact_tel2 IS 'เบอร์โทรศัพท์บุคคลติดต่อ2 contact number (2)';
COMMENT ON COLUMN public.landslide_area.risk_level IS 'ระดับความเสี่ยง risk level';
COMMENT ON COLUMN public.landslide_area.risk_detail IS 'รายละเอียดความเสี่ยงที่เกิดขึ้น risk detail';
COMMENT ON COLUMN public.landslide_area.remark IS 'หมายเหตุ remark';
COMMENT ON COLUMN public.landslide_area.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.landslide_area.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.landslide_area.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.landslide_area.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.landslide_area.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.landslide_area.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.landslide_area.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.landslide_area.deleted_at IS 'วันที่ลบข้อมูล deleted date';

