-- Table: public.metadata_frequency

-- DROP TABLE public.metadata_frequency;

CREATE TABLE public.metadata_frequency
(
  id bigserial NOT NULL, -- รหัสความถี่ของข้อมูล metadata frequency serial number
  metadata_id bigint, -- รหัสบัญชีข้อมูล metadata number
  datafrequency text, -- ความถี่ของข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_metadata_frequency PRIMARY KEY (id),
  CONSTRAINT fk_metadata_reference_metadata FOREIGN KEY (metadata_id)
      REFERENCES public.metadata (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_metadata_frequency UNIQUE (metadata_id, deleted_at, datafrequency)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata_frequency
  IS 'ความถี่ของข้อมูล';
COMMENT ON COLUMN public.metadata_frequency.id IS 'รหัสความถี่ของข้อมูล metadata frequency serial number';
COMMENT ON COLUMN public.metadata_frequency.metadata_id IS 'รหัสบัญชีข้อมูล metadata number';
COMMENT ON COLUMN public.metadata_frequency.datafrequency IS 'ความถี่ของข้อมูล';
COMMENT ON COLUMN public.metadata_frequency.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata_frequency.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata_frequency.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata_frequency.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata_frequency.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata_frequency.deleted_at IS 'วันที่ลบข้อมูล deleted date';

