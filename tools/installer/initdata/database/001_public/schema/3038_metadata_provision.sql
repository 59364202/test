-- Table: public.metadata_provision

-- DROP TABLE public.metadata_provision;

CREATE TABLE public.metadata_provision
(
  id bigserial NOT NULL, -- รหัสบัญชีข้อมูล metadata serial number
  agency_id bigint, -- รหัสหน่วยงานเจ้าของข้อมูล agency id
  metadataservice_name json, -- ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
  dataimport_download_id bigint, -- รหัสเลข download configuration
  dataimport_dataset_id bigint, -- รหัสเลข dataset configuration
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_metadata_provision PRIMARY KEY (id),
  CONSTRAINT uk_metadata_provision UNIQUE (agency_id, dataimport_download_id, dataimport_dataset_id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata_provision
  IS 'รายการบัญชีข้อมูลสำหรับคำนวณฝน';
COMMENT ON COLUMN public.metadata_provision.id IS 'รหัสบัญชีข้อมูล metadata serial number';
COMMENT ON COLUMN public.metadata_provision.agency_id IS 'รหัสหน่วยงานเจ้าของข้อมูล agency id';
COMMENT ON COLUMN public.metadata_provision.metadataservice_name IS 'ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล';
COMMENT ON COLUMN public.metadata_provision.dataimport_download_id IS 'รหัสเลข download configuration';
COMMENT ON COLUMN public.metadata_provision.dataimport_dataset_id IS 'รหัสเลข dataset configuration';
COMMENT ON COLUMN public.metadata_provision.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata_provision.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata_provision.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata_provision.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata_provision.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata_provision.deleted_at IS 'วันที่ลบข้อมูล deleted date';

