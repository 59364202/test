-- Table: public.lt_subcategory

-- DROP TABLE public.lt_subcategory;

CREATE TABLE public.lt_subcategory
(
  id bigserial NOT NULL, -- รหัสกลุ่มข้อมูลย่อย subcategory id
  category_id bigint NOT NULL, -- รหัสกลุ่มข้อมูลหลัก
  subcategory_name json, -- ชื่อกลุ่มข้อมูลย่อย
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_subcategory PRIMARY KEY (id),
  CONSTRAINT fk_lt_subca_reference_lt_categ FOREIGN KEY (category_id)
      REFERENCES public.lt_category (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_lt_subcategory UNIQUE (id, category_id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_subcategory
  IS 'กลุ่มข้อมูลย่อย ';
COMMENT ON COLUMN public.lt_subcategory.id IS 'รหัสกลุ่มข้อมูลย่อย subcategory id';
COMMENT ON COLUMN public.lt_subcategory.category_id IS 'รหัสกลุ่มข้อมูลหลัก';
COMMENT ON COLUMN public.lt_subcategory.subcategory_name IS 'ชื่อกลุ่มข้อมูลย่อย';
COMMENT ON COLUMN public.lt_subcategory.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_subcategory.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_subcategory.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_subcategory.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_subcategory.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_subcategory.deleted_at IS 'วันที่ลบข้อมูล deleted date';
