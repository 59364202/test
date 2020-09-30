-- Table: public.subbasin

-- DROP TABLE public.subbasin;

CREATE TABLE public.subbasin
(
  id bigserial NOT NULL, -- ลำดับลุ่มน้ำสาขา subbasin number
  basin_id bigint NOT NULL, -- รหัสลุ่มน้ำ basin number
  subbasin_code text, -- รหัสลุ่มน้ำสาขา subbasin code
  subbasin_name json, -- ชื่อลุ่มน้ำสาขา subbasin's name
  subbasin_area double precision, -- ความยาวของลุ่มน้ำ หน่วย: เมตร subbasin's area (square mater)
  subbasin_areakm double precision, -- ความยาวของลุ่มน้ำสาขา หน่วยเป็นกิโลเมตร subbasin's length
  subbasin_perimeter double precision, -- ความยาวของเส้นรอบรูป subbasin's perimeter
  subbasin_acres double precision, -- พื้นที่ของลุ่มน้ำสาขา หน่วยเป็น เอเคอร์ subbasin's area (square acre)
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_subbasin PRIMARY KEY (id),
  CONSTRAINT fk_subbasin_reference_basin FOREIGN KEY (basin_id)
      REFERENCES public.basin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_subbasin UNIQUE (subbasin_code, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.subbasin
  IS 'ลุ่มน้ำสาขา';
COMMENT ON COLUMN public.subbasin.id IS 'ลำดับลุ่มน้ำสาขา subbasin number';
COMMENT ON COLUMN public.subbasin.basin_id IS 'รหัสลุ่มน้ำ basin number';
COMMENT ON COLUMN public.subbasin.subbasin_code IS 'รหัสลุ่มน้ำสาขา subbasin code';
COMMENT ON COLUMN public.subbasin.subbasin_name IS 'ชื่อลุ่มน้ำสาขา subbasin''s name';
COMMENT ON COLUMN public.subbasin.subbasin_area IS 'ความยาวของลุ่มน้ำ หน่วย: เมตร subbasin''s area (square mater)';
COMMENT ON COLUMN public.subbasin.subbasin_areakm IS 'ความยาวของลุ่มน้ำสาขา หน่วยเป็นกิโลเมตร subbasin''s length';
COMMENT ON COLUMN public.subbasin.subbasin_perimeter IS 'ความยาวของเส้นรอบรูป subbasin''s perimeter';
COMMENT ON COLUMN public.subbasin.subbasin_acres IS 'พื้นที่ของลุ่มน้ำสาขา หน่วยเป็น เอเคอร์ subbasin''s area (square acre)';
COMMENT ON COLUMN public.subbasin.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.subbasin.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.subbasin.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.subbasin.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.subbasin.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.subbasin.deleted_at IS 'วันที่ลบข้อมูล deleted date';