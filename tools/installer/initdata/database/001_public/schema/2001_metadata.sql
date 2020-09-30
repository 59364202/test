-- Table: public.metadata

-- DROP TABLE public.metadata;

CREATE TABLE public.metadata
(
  id bigserial NOT NULL, -- รหัสบัญชีข้อมูล metadata serial number
  subcategory_id bigint, -- รหัสหมวดย่อยของข้อมูล subcategory number
  agency_id bigint, -- รหัสหน่วยงานเจ้าของข้อมูล agency id
  dataformat_id bigint, -- รหัสรูปแบบของข้อมูล Format Data s serial number
  dataunit_id bigint, -- รหัสหน่วยของข้อมูล unit of data's serial number
  metadatastatus_id bigint, -- รหัสสถานะการเชื่อมโยง
  connection_format text, -- รูปแบบการเชื่อมโยง (Online/Offline)
  metadataagency_name json, -- ชือบัญชีข้อมูลที่ยืนยันกับหน่วยงาน
  metadataservice_name json, -- ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
  metadata_tag json, -- tag ของบัญชีข้อมูล
  metadata_description json, -- รายละเอียดของบัญชีข้อมูล
  metadata_method text, -- วิธีการเชื่อมโยง (download/webservice/static)
  metadata_channel text, -- ช่องทางการเชื่อมโยงข้อมูล : รับข้อมูลจากโมดูลเชื่อมโยง
  metadata_convertfrequency text, -- ความถึ่การเชื่อมโยงของบัญชีข้อมูล : รับข้อมูลจากโมดูลเชื่อมโยง
  metadata_contact text, -- ติดต่อเจ้าของข้อมูล
  metadata_agencystoredate date, -- วันที่หน่วยงานเริ่มจัดเก็บข้อมูล
  metadata_startdatadate date, -- วันที่เริ่มใช้งานข้อมูล
  metadata_update_plan bigint, -- ระยะเวลาการปรับปรุงข้อมูล หน่วยเป็นนาที
  metadata_laws text, -- ข้อจำกัดทางกฎหมาย
  metadata_remark text, -- หมายเหตุ
  dataimport_download_id bigint, -- รหัสเลข download configuration
  dataimport_dataset_id bigint, -- รหัสเลข dataset configuration
  import_count smallint, -- จำนวนที่ใช้ในการคำนวณเปอร์เซนต์นำเข้าข้อมูล หน่วย : ครั้งต่อวัน
  frequencyunit_id bigint, -- รหัสของหน่วยข้อมูล
  metadata_offline_date timestamp with time zone, -- สำหรับเก็บวันที่แจ้งเตือนข้อมูล offline บน dashbord
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  metadata_frequency bigint,
  metadata_status text,
  data_min_date date, -- วันที่น้อยที่สุดของข้อมูล
  data_max_date date, -- วันที่มากที่สุดของข้อมูล
  data_last_check date, -- วันที่เช็ค min, max date ครั้งสุดท้าย
  CONSTRAINT pk_metadata PRIMARY KEY (id),
  CONSTRAINT fk_metadata_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_lt_dataf FOREIGN KEY (dataformat_id)
      REFERENCES public.lt_dataformat (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_lt_datau FOREIGN KEY (dataunit_id)
      REFERENCES public.lt_dataunit (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_lt_metad FOREIGN KEY (metadatastatus_id)
      REFERENCES public.lt_metadata_status (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_lt_subca FOREIGN KEY (subcategory_id)
      REFERENCES public.lt_subcategory (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata
  IS 'รายการบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata.id IS 'รหัสบัญชีข้อมูล metadata serial number';
COMMENT ON COLUMN public.metadata.subcategory_id IS 'รหัสหมวดย่อยของข้อมูล subcategory number';
COMMENT ON COLUMN public.metadata.agency_id IS 'รหัสหน่วยงานเจ้าของข้อมูล agency id';
COMMENT ON COLUMN public.metadata.dataformat_id IS 'รหัสรูปแบบของข้อมูล Format Data s serial number';
COMMENT ON COLUMN public.metadata.dataunit_id IS 'รหัสหน่วยของข้อมูล unit of data''s serial number';
COMMENT ON COLUMN public.metadata.metadatastatus_id IS 'รหัสสถานะการเชื่อมโยง';
COMMENT ON COLUMN public.metadata.connection_format IS 'รูปแบบการเชื่อมโยง (Online/Offline)';
COMMENT ON COLUMN public.metadata.metadataagency_name IS 'ชือบัญชีข้อมูลที่ยืนยันกับหน่วยงาน ';
COMMENT ON COLUMN public.metadata.metadataservice_name IS 'ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_tag IS 'tag ของบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_description IS 'รายละเอียดของบัญชีข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_method IS 'วิธีการเชื่อมโยง (download/webservice/static)';
COMMENT ON COLUMN public.metadata.metadata_channel IS 'ช่องทางการเชื่อมโยงข้อมูล : รับข้อมูลจากโมดูลเชื่อมโยง';
COMMENT ON COLUMN public.metadata.metadata_convertfrequency IS 'ความถึ่การเชื่อมโยงของบัญชีข้อมูล : รับข้อมูลจากโมดูลเชื่อมโยง';
COMMENT ON COLUMN public.metadata.metadata_contact IS 'ติดต่อเจ้าของข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_agencystoredate IS 'วันที่หน่วยงานเริ่มจัดเก็บข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_startdatadate IS 'วันที่เริ่มใช้งานข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_update_plan IS 'ระยะเวลาการปรับปรุงข้อมูล หน่วยเป็นนาที';
COMMENT ON COLUMN public.metadata.metadata_laws IS 'ข้อจำกัดทางกฎหมาย';
COMMENT ON COLUMN public.metadata.metadata_remark IS 'หมายเหตุ';
COMMENT ON COLUMN public.metadata.dataimport_download_id IS 'รหัสเลข download configuration';
COMMENT ON COLUMN public.metadata.dataimport_dataset_id IS 'รหัสเลข dataset configuration';
COMMENT ON COLUMN public.metadata.import_count IS 'จำนวนที่ใช้ในการคำนวณเปอร์เซนต์นำเข้าข้อมูล หน่วย : ครั้งต่อวัน';
COMMENT ON COLUMN public.metadata.frequencyunit_id IS 'รหัสของหน่วยข้อมูล';
COMMENT ON COLUMN public.metadata.metadata_offline_date IS 'สำหรับเก็บวันที่แจ้งเตือนข้อมูล offline บน dashbord';
COMMENT ON COLUMN public.metadata.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata.data_min_date IS 'วันที่น้อยที่สุดของข้อมูล';
COMMENT ON COLUMN public.metadata.data_max_date IS 'วันที่มากที่สุดของข้อมูล';
COMMENT ON COLUMN public.metadata.data_last_check IS 'วันที่เช็ค min, max date ครั้งสุดท้าย';
