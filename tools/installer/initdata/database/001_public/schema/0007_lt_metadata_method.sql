-- Table: public.lt_metadata_method

-- DROP TABLE public.lt_metadata_method;

CREATE TABLE public.lt_metadata_method
(
  id bigserial NOT NULL, -- รหัสวิธีการได้มาของข้อมูล Format Data s serial number
  metadata_method_name text, -- ชื่อวิธีการได้มาซึ่งข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_metadata_method PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_metadata_method
  IS 'วิธีการได้มาของข้อมูล เช่น Download Webservice Web Extract ';
COMMENT ON COLUMN public.lt_metadata_method.id IS 'รหัสวิธีการได้มาของข้อมูล Format Data s serial number';
COMMENT ON COLUMN public.lt_metadata_method.metadata_method_name IS 'ชื่อวิธีการได้มาซึ่งข้อมูล';
COMMENT ON COLUMN public.lt_metadata_method.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_metadata_method.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_metadata_method.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_metadata_method.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_metadata_method.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_metadata_method.deleted_at IS 'วันที่ลบข้อมูล deleted date';