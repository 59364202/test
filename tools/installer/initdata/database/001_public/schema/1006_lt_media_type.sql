-- Table: public.lt_media_type

-- DROP TABLE public.lt_media_type;

CREATE TABLE public.lt_media_type
(
  id bigserial NOT NULL, -- รหัสแสดงชนิดข้อมูลสื่อ
  media_type_name text, -- ชื่อแสดงชนิดข้อมูลสื่อ
  media_subtype_name text, -- ชื่อแสดงชนิดย่อยข้อมูลสื่อ
  media_category text, -- กลุ่มของสื่อ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_media_type PRIMARY KEY (id),
  CONSTRAINT uk_lt_media_type UNIQUE (media_type_name, media_subtype_name, media_category, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_media_type
  IS 'กลุ่มหลักของข้อมูลสื่อ เช่น รูปภาพ ';
COMMENT ON COLUMN public.lt_media_type.id IS 'รหัสแสดงชนิดข้อมูลสื่อ';
COMMENT ON COLUMN public.lt_media_type.media_type_name IS 'ชื่อแสดงชนิดข้อมูลสื่อ';
COMMENT ON COLUMN public.lt_media_type.media_subtype_name IS 'ชื่อแสดงชนิดย่อยข้อมูลสื่อ';
COMMENT ON COLUMN public.lt_media_type.media_category IS 'กลุ่มของสื่อ';
COMMENT ON COLUMN public.lt_media_type.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_media_type.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_media_type.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_media_type.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_media_type.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_media_type.deleted_at IS 'วันที่ลบข้อมูล deleted date';