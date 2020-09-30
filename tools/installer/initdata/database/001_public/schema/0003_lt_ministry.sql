-- Table: public.lt_ministry

-- DROP TABLE public.lt_ministry;

CREATE TABLE public.lt_ministry
(
  id bigserial NOT NULL, -- ลำดับข้อมูลกระทรวง ministry number
  ministry_code text NOT NULL, -- รหัสกระทรวง ministry code
  ministry_shortname json, -- ชื่อย่อกระทรวง
  ministry_name json NOT NULL, -- ชื่อกระทรวง
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_lt_ministry PRIMARY KEY (id),
  CONSTRAINT uk_lt_ministry UNIQUE (ministry_code, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_ministry
  IS 'กระทรวง';
COMMENT ON COLUMN public.lt_ministry.id IS 'ลำดับข้อมูลกระทรวง ministry number';
COMMENT ON COLUMN public.lt_ministry.ministry_code IS 'รหัสกระทรวง ministry code';
COMMENT ON COLUMN public.lt_ministry.ministry_shortname IS 'ชื่อย่อกระทรวง';
COMMENT ON COLUMN public.lt_ministry.ministry_name IS 'ชื่อกระทรวง';
COMMENT ON COLUMN public.lt_ministry.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_ministry.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_ministry.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_ministry.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_ministry.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_ministry.deleted_at IS 'วันที่ลบข้อมูล deleted date';
