-- Table: public.ground_waterquality

-- DROP TABLE public.ground_waterquality;

CREATE TABLE public.ground_waterquality
(
  id bigserial NOT NULL, -- ลำดับข้อมูลคุณภาพน้ำบาดาล
  ground_id bigint, -- รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_conductivity double precision, -- ค่านำกระแสไฟฟ้าแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_ph double precision, -- ค่าความเป็นกรด - ด่างแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_temp double precision, -- อุณหภูมิของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_color double precision, -- สีของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_scent double precision, -- กลิ่นของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_turbid double precision, -- ความขุ่นของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_status text, -- สถานะการใช้งานของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_ground_waterquality PRIMARY KEY (id),
  CONSTRAINT fk_ground_w_reference_m_ground FOREIGN KEY (ground_id)
      REFERENCES public.m_ground (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_ground_waterquality UNIQUE (ground_id, deleted_at, id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.ground_waterquality
  IS 'คุณภาพแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล (เก็บเฉพาะที่โอนย้ายมาจากฐาน NHC ส่วนระบบใหม่เก็บเป็น media)';
COMMENT ON COLUMN public.ground_waterquality.id IS 'ลำดับข้อมูลคุณภาพน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_id IS 'รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_conductivity IS 'ค่านำกระแสไฟฟ้าแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_ph IS 'ค่าความเป็นกรด - ด่างแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_temp IS 'อุณหภูมิของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_color IS 'สีของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_scent IS 'กลิ่นของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_turbid IS 'ความขุ่นของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.ground_status IS 'สถานะการใช้งานของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterquality.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.ground_waterquality.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.ground_waterquality.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.ground_waterquality.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.ground_waterquality.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.ground_waterquality.deleted_at IS 'วันที่ลบข้อมูล deleted date';

