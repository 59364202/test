-- Table: public.lt_department

-- DROP TABLE public.lt_department;

CREATE TABLE public.lt_department
(
  id bigserial NOT NULL, -- ลำดับข้อมูลกรม
  ministry_id bigint NOT NULL, -- ลำดับข้อมูลกระทรวง
  department_code text NOT NULL, -- รหัสกรม
  department_shortname json, -- ชื่อย่อกรม
  department_name json NOT NULL, -- ชื่อกรม
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_department PRIMARY KEY (id),
  CONSTRAINT fk_lt_depar_reference_lt_minis FOREIGN KEY (ministry_id)
      REFERENCES public.lt_ministry (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_lt_department UNIQUE (department_code, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_department
  IS 'ข้อมูลกรม';
COMMENT ON COLUMN public.lt_department.id IS 'ลำดับข้อมูลกรม';
COMMENT ON COLUMN public.lt_department.ministry_id IS 'ลำดับข้อมูลกระทรวง';
COMMENT ON COLUMN public.lt_department.department_code IS 'รหัสกรม';
COMMENT ON COLUMN public.lt_department.department_shortname IS 'ชื่อย่อกรม';
COMMENT ON COLUMN public.lt_department.department_name IS 'ชื่อกรม';
COMMENT ON COLUMN public.lt_department.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_department.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_department.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_department.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_department.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_department.deleted_at IS 'วันที่ลบข้อมูล deleted date';