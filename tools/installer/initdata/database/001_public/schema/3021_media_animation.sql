-- Table: public.media_animation

-- DROP TABLE public.media_animation;

CREATE TABLE public.media_animation
(
  id bigserial NOT NULL, -- รหัสแสดงข้อมูลสื่อ media's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  media_type_id bigint, -- รหัสแสดงชนิดข้อมูลสื่อ
  media_datetime timestamp with time zone, -- วันที่เก็บข้อมูลสื่อ record date
  media_path text NOT NULL, -- ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
  media_desc text, -- รายละเอียดของข้อมูลสื่อ description
  filename text, -- ชื่อไฟล์ของข้อมูลสื่อ file name
  refer_source text, -- แหล่งข้อมูลสำหรับอ้างอิง reference source
  remark text, -- หมายเหตุ
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_media_animation PRIMARY KEY (id),
  CONSTRAINT fk_media_an_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_media_an_reference_lt_media FOREIGN KEY (media_type_id)
      REFERENCES public.lt_media_type (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_media_animation UNIQUE (media_datetime, deleted_at, agency_id, media_type_id, filename),
  CONSTRAINT pt_media_animation_media_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.media_animation
  IS 'ข้อมูลสื่อส่วนของภาพเคลื่อนไหว (แยกตารางกับ media เนื่องจากการจัดเก็บภาพ Animation จะเรียกใช้งานเฉพาะ Latest) แยกตาม media_category';
COMMENT ON COLUMN public.media_animation.id IS 'รหัสแสดงข้อมูลสื่อ media''s serial number';
COMMENT ON COLUMN public.media_animation.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.media_animation.media_type_id IS 'รหัสแสดงชนิดข้อมูลสื่อ';
COMMENT ON COLUMN public.media_animation.media_datetime IS 'วันที่เก็บข้อมูลสื่อ record date';
COMMENT ON COLUMN public.media_animation.media_path IS 'ที่อยู่ของไฟล์ข้อมูลสื่อ file path location';
COMMENT ON COLUMN public.media_animation.media_desc IS 'รายละเอียดของข้อมูลสื่อ description';
COMMENT ON COLUMN public.media_animation.filename IS 'ชื่อไฟล์ของข้อมูลสื่อ file name';
COMMENT ON COLUMN public.media_animation.refer_source IS 'แหล่งข้อมูลสำหรับอ้างอิง reference source';
COMMENT ON COLUMN public.media_animation.remark IS 'หมายเหตุ';
COMMENT ON COLUMN public.media_animation.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.media_animation.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.media_animation.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.media_animation.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.media_animation.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.media_animation.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.media_animation.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.media_animation.deleted_at IS 'วันที่ลบข้อมูล deleted date';

