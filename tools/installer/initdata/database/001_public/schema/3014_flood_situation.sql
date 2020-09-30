-- Table: public.flood_situation

-- DROP TABLE public.flood_situation;

CREATE TABLE public.flood_situation
(
  id bigserial NOT NULL, -- รหัสสถานการณ์น้ำ flood situation serial number
  agency_id bigint NOT NULL, -- รหัสหน่วยงาน agency number
  flood_datetime timestamp with time zone NOT NULL, -- วันและเวลาที่ประกาศสถานการณ์น้ำ DateTime of flood
  flood_name text, -- ชื่อสถานการณ์น้ำ name of flood
  flood_link text, -- ลิ้งที่แสดงสถานการณ์น้ำ flood link
  flood_description text, -- รายละเอียดสถานการณ์น้ำ description of flood
  flood_author text, -- ผู้รายงานสถานการณ์ author
  flood_colorlevel text, -- ระดับสีเกณฑืเตือนภัย
  flood_remark text, -- หมายเหตุ
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_flood_situation PRIMARY KEY (id),
  CONSTRAINT fk_flood_si_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_flood_situation UNIQUE (flood_datetime, flood_link, deleted_at),
  CONSTRAINT pt_flood_situation_flood_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.flood_situation
  IS 'สถานการณ์น้ำ';
COMMENT ON COLUMN public.flood_situation.id IS 'รหัสสถานการณ์น้ำ flood situation serial number';
COMMENT ON COLUMN public.flood_situation.agency_id IS 'รหัสหน่วยงาน agency number';
COMMENT ON COLUMN public.flood_situation.flood_datetime IS 'วันและเวลาที่ประกาศสถานการณ์น้ำ DateTime of flood';
COMMENT ON COLUMN public.flood_situation.flood_name IS 'ชื่อสถานการณ์น้ำ name of flood';
COMMENT ON COLUMN public.flood_situation.flood_link IS 'ลิ้งที่แสดงสถานการณ์น้ำ flood link';
COMMENT ON COLUMN public.flood_situation.flood_description IS 'รายละเอียดสถานการณ์น้ำ description of flood';
COMMENT ON COLUMN public.flood_situation.flood_author IS 'ผู้รายงานสถานการณ์ author';
COMMENT ON COLUMN public.flood_situation.flood_colorlevel IS 'ระดับสีเกณฑืเตือนภัย';
COMMENT ON COLUMN public.flood_situation.flood_remark IS 'หมายเหตุ';
COMMENT ON COLUMN public.flood_situation.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.flood_situation.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.flood_situation.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.flood_situation.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.flood_situation.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.flood_situation.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.flood_situation.deleted_at IS 'วันที่ลบข้อมูล deleted date';

