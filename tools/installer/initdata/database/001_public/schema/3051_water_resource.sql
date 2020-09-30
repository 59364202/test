-- Table: public.water_resource

-- DROP TABLE public.water_resource;

CREATE TABLE public.water_resource
(
  id bigserial NOT NULL, -- รหัสแหล่งน้ำชุมชน water resource's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  water_resource_oldcode text, -- รหัสแหล่งน้ำชุมชน water resource's code
  projectname text, -- ชื่อโครงการ project name
  projecttype text, -- กิจกรรม project type
  fiscal_year text, -- ปีงบประมาณ fiscal year
  mooban text, -- หมู่บ้าน address number
  coordination text, -- ตำแหน่งโครงการ coordination
  benefit_household integer, -- ครัวเรือนที่ได้รับประโยชน์ benefit household
  benefit_area integer, -- พื้นที่ที่ได้รับประโยชน์ benefit area
  capacity double precision, -- ความจุ capacity
  standard_cost double precision, -- งบประมาณ standard cost
  budget real, -- ค่าใช้จ่าย budget
  contract_signdate text, -- วันที่ลงนามในสัญญา contract signdate
  contract_enddate text, -- วันที่สิ้นสุดสัญญา contract enddate
  rec_date date, -- วันที่บันทึกข้อมูล record date
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_water_resource PRIMARY KEY (id),
  CONSTRAINT fk_water_re_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_water_resource UNIQUE (agency_id, water_resource_oldcode, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.water_resource
  IS 'ข้อมูลแหล่งน้ำขนาดเล็ก';
COMMENT ON COLUMN public.water_resource.id IS 'รหัสแหล่งน้ำชุมชน water resource''s serial number';
COMMENT ON COLUMN public.water_resource.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.water_resource.water_resource_oldcode IS 'รหัสแหล่งน้ำชุมชน water resource''s code';
COMMENT ON COLUMN public.water_resource.projectname IS 'ชื่อโครงการ project name';
COMMENT ON COLUMN public.water_resource.projecttype IS 'กิจกรรม project type';
COMMENT ON COLUMN public.water_resource.fiscal_year IS 'ปีงบประมาณ fiscal year';
COMMENT ON COLUMN public.water_resource.mooban IS 'หมู่บ้าน address number';
COMMENT ON COLUMN public.water_resource.coordination IS 'ตำแหน่งโครงการ coordination';
COMMENT ON COLUMN public.water_resource.benefit_household IS 'ครัวเรือนที่ได้รับประโยชน์ benefit household';
COMMENT ON COLUMN public.water_resource.benefit_area IS 'พื้นที่ที่ได้รับประโยชน์ benefit area';
COMMENT ON COLUMN public.water_resource.capacity IS 'ความจุ capacity';
COMMENT ON COLUMN public.water_resource.standard_cost IS 'งบประมาณ standard cost';
COMMENT ON COLUMN public.water_resource.budget IS 'ค่าใช้จ่าย budget';
COMMENT ON COLUMN public.water_resource.contract_signdate IS 'วันที่ลงนามในสัญญา contract signdate';
COMMENT ON COLUMN public.water_resource.contract_enddate IS 'วันที่สิ้นสุดสัญญา contract enddate';
COMMENT ON COLUMN public.water_resource.rec_date IS 'วันที่บันทึกข้อมูล record date';
COMMENT ON COLUMN public.water_resource.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.water_resource.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.water_resource.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.water_resource.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.water_resource.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.water_resource.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.water_resource.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.water_resource.deleted_at IS 'วันที่ลบข้อมูล deleted date';

