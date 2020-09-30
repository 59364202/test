-- Table: public.basin

-- DROP TABLE public.basin;

CREATE TABLE public.basin
(
  id bigserial NOT NULL, -- รหัสลุ่มน้ำ basin's number
  agency_id bigint, -- รหัสหน่วยงาน agency's number
  basin_code smallint, -- รหัสลุ่มน้ำ basin's code
  basin_name json, -- ชื่อลุ่มน้ำ basin's name
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_basin PRIMARY KEY (id),
  CONSTRAINT fk_basin_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_basin UNIQUE (basin_code, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.basin
  IS 'ข้อมูลลุ่มน้ำ';
COMMENT ON COLUMN public.basin.id IS 'รหัสลุ่มน้ำ basin''s number';
COMMENT ON COLUMN public.basin.agency_id IS 'รหัสหน่วยงาน agency''s number';
COMMENT ON COLUMN public.basin.basin_code IS 'รหัสลุ่มน้ำ basin''s code';
COMMENT ON COLUMN public.basin.basin_name IS 'ชื่อลุ่มน้ำ basin''s name';
COMMENT ON COLUMN public.basin.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.basin.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.basin.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.basin.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.basin.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.basin.deleted_at IS 'วันที่ลบข้อมูล deleted date';
