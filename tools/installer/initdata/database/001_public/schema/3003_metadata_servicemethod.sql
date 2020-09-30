-- Table: public.metadata_servicemethod

-- DROP TABLE public.metadata_servicemethod;

CREATE TABLE public.metadata_servicemethod
(
  id bigserial NOT NULL, -- รหัสรูปแบบการให้บริการข้อมูลของแต่ละบัญชีข้อมูล
  metadata_id bigint, -- รหัสบัญชีข้อมูล metadata serial number
  servicemethod_id bigint, -- รหัสวิธีการให้บริการข้อมูล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_metadata_servicemethod PRIMARY KEY (id),
  CONSTRAINT fk_metadata_reference_lt_servi FOREIGN KEY (servicemethod_id)
      REFERENCES public.lt_servicemethod (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_metadata FOREIGN KEY (metadata_id)
      REFERENCES public.metadata (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_metadata_servicemethod UNIQUE (metadata_id, servicemethod_id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata_servicemethod
  IS 'รูปแบบการให้บริการข้อมูลของแต่ละบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_servicemethod.id IS 'รหัสรูปแบบการให้บริการข้อมูลของแต่ละบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata_servicemethod.metadata_id IS 'รหัสบัญชีข้อมูล metadata serial number';
COMMENT ON COLUMN public.metadata_servicemethod.servicemethod_id IS 'รหัสวิธีการให้บริการข้อมูล';
COMMENT ON COLUMN public.metadata_servicemethod.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata_servicemethod.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata_servicemethod.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata_servicemethod.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata_servicemethod.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata_servicemethod.deleted_at IS 'วันที่ลบข้อมูล deleted date';

