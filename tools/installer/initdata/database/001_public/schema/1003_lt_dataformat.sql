-- Table: public.lt_dataformat

-- DROP TABLE public.lt_dataformat;

CREATE TABLE public.lt_dataformat
(
  id bigserial NOT NULL, -- รหัสรูปแบบของข้อมูล Format Data s serial number
  metadata_method_id bigint NOT NULL, -- รหัสวิธีการได้มาของข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  dataformat_name json, -- ชื่อรูปแบบข้อมูล
  CONSTRAINT pk_lt_dataformat PRIMARY KEY (id),
  CONSTRAINT fk_lt_dataf_reference_lt_metad FOREIGN KEY (metadata_method_id)
      REFERENCES public.lt_metadata_method (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.lt_dataformat
  IS 'รูปแบบของข้อมูล';
COMMENT ON COLUMN public.lt_dataformat.id IS 'รหัสรูปแบบของข้อมูล Format Data s serial number';
COMMENT ON COLUMN public.lt_dataformat.metadata_method_id IS 'รหัสวิธีการได้มาของข้อมูล ';
COMMENT ON COLUMN public.lt_dataformat.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.lt_dataformat.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.lt_dataformat.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.lt_dataformat.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.lt_dataformat.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.lt_dataformat.deleted_at IS 'วันที่ลบข้อมูล deleted date';
COMMENT ON COLUMN public.lt_dataformat.dataformat_name IS 'ชื่อรูปแบบข้อมูล';