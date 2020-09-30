-- Table: public.ground_waterlevel

-- DROP TABLE public.ground_waterlevel;

CREATE TABLE public.ground_waterlevel
(
  id bigserial NOT NULL, -- ลำดับข้อมูลคุณภาพน้ำบาดาล
  ground_id bigint, -- รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_size double precision, -- ขนาดของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล
  ground_depth double precision, -- ความลึกแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล (เมตร)
  ground_waterlevel double precision, -- ระดับน้ำแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล (เมตร)
  ground_aquifer text, -- ชั้นหินในน้ำ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_ground_waterlevel PRIMARY KEY (id),
  CONSTRAINT fk_ground_w_reference_m_ground FOREIGN KEY (id)
      REFERENCES public.m_ground (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_ground_waterlevel UNIQUE (ground_id, id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.ground_waterlevel
  IS 'ระดับน้ำแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล  (เก็บเฉพาะที่โอนย้ายมาจากฐาน NHC ส่วนระบบใหม่เก็บเป็น media)';
COMMENT ON COLUMN public.ground_waterlevel.id IS 'ลำดับข้อมูลคุณภาพน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterlevel.ground_id IS 'รหัสแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterlevel.ground_size IS 'ขนาดของแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล';
COMMENT ON COLUMN public.ground_waterlevel.ground_depth IS 'ความลึกแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล (เมตร)';
COMMENT ON COLUMN public.ground_waterlevel.ground_waterlevel IS 'ระดับน้ำแหล่งน้ำผิวดิน / แหล่งน้ำบาดาล (เมตร)';
COMMENT ON COLUMN public.ground_waterlevel.ground_aquifer IS 'ชั้นหินในน้ำ';
COMMENT ON COLUMN public.ground_waterlevel.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.ground_waterlevel.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.ground_waterlevel.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.ground_waterlevel.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.ground_waterlevel.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.ground_waterlevel.deleted_at IS 'วันที่ลบข้อมูล deleted date';

