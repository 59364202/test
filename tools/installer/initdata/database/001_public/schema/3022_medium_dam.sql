-- Table: public.medium_dam

-- DROP TABLE public.medium_dam;

CREATE TABLE public.medium_dam
(
  id bigserial NOT NULL, -- รหัสข้อมูลเขื่อนขนาดกลาง รายวัน medium dam (daily)'s serial number
  mediumdam_id bigint, -- รหัสข้อมูลเขื่อนขนาดกลาง  medium dam number
  mediumdam_date date NOT NULL, -- วันที่เก็บข้อมูล record date
  mediumdam_storage double precision, -- ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
  mediumdam_inflow double precision, -- ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
  mediumdam_released double precision, -- ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
  mediumdam_uses_water double precision, -- ปริมาณน้ำที่ใช้ได้ uses water volume
  mediumdam_storage_percent double precision, -- เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_medium_dam PRIMARY KEY (id),
  CONSTRAINT fk_medium_d_reference_m_medium FOREIGN KEY (mediumdam_id)
      REFERENCES public.m_medium_dam (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_medium_dam UNIQUE (mediumdam_id, mediumdam_date, deleted_at),
  CONSTRAINT pt_medium_dam_mediumdam_date CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.medium_dam
  IS 'ข้อมูลเขื่อนขนาดกลาง';
COMMENT ON COLUMN public.medium_dam.id IS 'รหัสข้อมูลเขื่อนขนาดกลาง รายวัน medium dam (daily)''s serial number';
COMMENT ON COLUMN public.medium_dam.mediumdam_id IS 'รหัสข้อมูลเขื่อนขนาดกลาง  medium dam number';
COMMENT ON COLUMN public.medium_dam.mediumdam_date IS 'วันที่เก็บข้อมูล record date';
COMMENT ON COLUMN public.medium_dam.mediumdam_storage IS 'ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume';
COMMENT ON COLUMN public.medium_dam.mediumdam_inflow IS 'ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume';
COMMENT ON COLUMN public.medium_dam.mediumdam_released IS 'ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume';
COMMENT ON COLUMN public.medium_dam.mediumdam_uses_water IS 'ปริมาณน้ำที่ใช้ได้ uses water volume';
COMMENT ON COLUMN public.medium_dam.mediumdam_storage_percent IS 'เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume';
COMMENT ON COLUMN public.medium_dam.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.medium_dam.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.medium_dam.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.medium_dam.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.medium_dam.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.medium_dam.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.medium_dam.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.medium_dam.deleted_at IS 'วันที่ลบข้อมูล deleted date';

