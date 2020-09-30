-- Table: public.rulecurve

-- DROP TABLE public.rulecurve;

CREATE TABLE public.rulecurve
(
  id bigserial NOT NULL, -- รหัสชนิดของ Rule Curves / rule curve type's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  dam_id bigint, -- รหัสของเขื่อน
  dam_name text, -- ชื่อเขื่อน dam's name
  rc_datetime timestamp with time zone, -- วันเดือนปีที่วัดค่า Rule Curves / rule curve measure year
  rc_unit text, -- หน่วยการวัด RuleCurve โดยมีหน่วยเป็น ล้าน ลูกบากศ์เมตร (mcm) หรือหน่วยเป็น เมตร (msl)
  urc_old double precision, -- upper rule curve เดิม
  lrc_old double precision, -- lower rule curve เดิม
  urc_new double precision, -- upper rule curve ปรับปรุง
  lrc_new double precision, -- lower rule curve ปรับปรุง
  rc_lat numeric(9,6), -- ละติจูดของการวัด rule curve
  rc_long numeric(9,6), -- ลองติจูดของการวัด rule curve
  rc_filepath text, -- ที่อยู่ของไฟล์ file path location
  rc_offset double precision, -- ระยะห่างของการวัดแต่ละจุด เมตร
  rc_remark text, -- หมายเหตุ เช่น LB RB
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_rulecurve PRIMARY KEY (id),
  CONSTRAINT fk_rulecurv_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT pt_rulecurve_rc_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.rulecurve
  IS 'เกณฑ์ควบคุมระดับน้ำ (Rule Curve)';
COMMENT ON COLUMN public.rulecurve.id IS 'รหัสชนิดของ Rule Curves / rule curve type''s serial number';
COMMENT ON COLUMN public.rulecurve.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.rulecurve.dam_id IS 'รหัสของเขื่อน';
COMMENT ON COLUMN public.rulecurve.dam_name IS 'ชื่อเขื่อน dam''s name';
COMMENT ON COLUMN public.rulecurve.rc_datetime IS 'วันเดือนปีที่วัดค่า Rule Curves / rule curve measure year';
COMMENT ON COLUMN public.rulecurve.rc_unit IS 'หน่วยการวัด RuleCurve โดยมีหน่วยเป็น ล้าน ลูกบากศ์เมตร (mcm) หรือหน่วยเป็น เมตร (msl)';
COMMENT ON COLUMN public.rulecurve.urc_old IS 'upper rule curve เดิม';
COMMENT ON COLUMN public.rulecurve.lrc_old IS 'lower rule curve เดิม';
COMMENT ON COLUMN public.rulecurve.urc_new IS 'upper rule curve ปรับปรุง';
COMMENT ON COLUMN public.rulecurve.lrc_new IS 'lower rule curve ปรับปรุง';
COMMENT ON COLUMN public.rulecurve.rc_lat IS 'ละติจูดของการวัด rule curve';
COMMENT ON COLUMN public.rulecurve.rc_long IS 'ลองติจูดของการวัด rule curve';
COMMENT ON COLUMN public.rulecurve.rc_filepath IS 'ที่อยู่ของไฟล์ file path location';
COMMENT ON COLUMN public.rulecurve.rc_offset IS 'ระยะห่างของการวัดแต่ละจุด เมตร';
COMMENT ON COLUMN public.rulecurve.rc_remark IS 'หมายเหตุ เช่น LB RB';
COMMENT ON COLUMN public.rulecurve.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.rulecurve.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.rulecurve.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.rulecurve.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.rulecurve.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.rulecurve.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.rulecurve.deleted_at IS 'วันที่ลบข้อมูล deleted date';

