-- Table: public.lt_category

-- DROP TABLE public.lt_category;

CREATE TABLE public.lt_category
(
  id bigserial NOT NULL, -- รหัสกลุ่มข้อมูลหลัก
  category_name json, -- ชื่อกลุ่มข้อมูลหลัก
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_category PRIMARY KEY (id),
  CONSTRAINT uk_lt_category UNIQUE (id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_category
  IS 'กลุ่มข้อมูลหลัก';
COMMENT ON COLUMN public.lt_category.id IS 'รหัสกลุ่มข้อมูลหลัก';
COMMENT ON COLUMN public.lt_category.category_name IS 'ชื่อกลุ่มข้อมูลหลัก';
COMMENT ON COLUMN public.lt_category.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_category.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_category.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_category.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_category.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_category.deleted_at IS 'วันที่ลบข้อมูล deleted date';