-- Table: public.metadata_hydroinfo

-- DROP TABLE public.metadata_hydroinfo;

CREATE TABLE public.metadata_hydroinfo
(
  id bigserial NOT NULL, -- รหัสการจับคู่ระหว่างบัญชีข้อมูลกับข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo 's serial number
  metadata_id bigint, -- รหัสบัญชีข้อมูล metadata serial number
  hydroinfo_id bigint, -- รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_metadata_hydroinfo PRIMARY KEY (id),
  CONSTRAINT fk_metadata_reference_lt_hydro FOREIGN KEY (hydroinfo_id)
      REFERENCES public.lt_hydroinfo (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_metadata_reference_metadata FOREIGN KEY (metadata_id)
      REFERENCES public.metadata (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_metadata_hydroinfo UNIQUE (metadata_id, hydroinfo_id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.metadata_hydroinfo
  IS ' การจับคู่ระหว่างบัญชีข้อมูลกับข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ แสดงหน่วยงานรับผิดชอบการใช้ข้อมูล 9 ด้าน (Main Responsible Agencies in “9 Aspects of Hydroinformatics”)';
COMMENT ON COLUMN public.metadata_hydroinfo.id IS 'รหัสการจับคู่ระหว่างบัญชีข้อมูลกับข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ hydroinfo ''s serial number';
COMMENT ON COLUMN public.metadata_hydroinfo.metadata_id IS 'รหัสบัญชีข้อมูล metadata serial number';
COMMENT ON COLUMN public.metadata_hydroinfo.hydroinfo_id IS 'รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ';
COMMENT ON COLUMN public.metadata_hydroinfo.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.metadata_hydroinfo.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.metadata_hydroinfo.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.metadata_hydroinfo.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.metadata_hydroinfo.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.metadata_hydroinfo.deleted_at IS 'วันที่ลบข้อมูล deleted date';

