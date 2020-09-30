-- Table: public.crosssection

-- DROP TABLE public.crosssection;

CREATE TABLE public.crosssection
(
  id bigserial NOT NULL, -- รหัสข้อมูลภาพตัดลำน้ำ canal station's serial number
  section_station_id bigint, -- รหัสสถานีของข้อมูลภาพตัดลำน้ำ section station's serial number
  point_id text, -- ตำแหน่งของการจุดลำน้ำ
  section_lat numeric(9,6), -- ตำแหน่งละติจูดของการวัด
  section_long numeric(9,6), -- ตำแหน่งลองติจูดของการวัด
  distance text NOT NULL, -- ระยะทาง หน่วย : เมตร distance
  water_level_msl text, -- ระดับน้ำ หน่วย : ม.รทก. water level (msl)
  water_level_m text, -- ระดับน้ำ หน่วย : ม water level (m)
  remark text, -- ตำแหน่งที่วัด LB : left bank RB : right bank  CL / location
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_crosssection PRIMARY KEY (id),
  CONSTRAINT fk_crosssec_reference_m_crosss FOREIGN KEY (section_station_id)
      REFERENCES public.m_crosssection_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_crosssection UNIQUE (section_station_id, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.crosssection
  IS 'ข้อมูลภาพตัดลำน้ำ';
COMMENT ON COLUMN public.crosssection.id IS 'รหัสข้อมูลภาพตัดลำน้ำ canal station''s serial number';
COMMENT ON COLUMN public.crosssection.section_station_id IS 'รหัสสถานีของข้อมูลภาพตัดลำน้ำ section station''s serial number';
COMMENT ON COLUMN public.crosssection.point_id IS 'ตำแหน่งของการจุดลำน้ำ';
COMMENT ON COLUMN public.crosssection.section_lat IS 'ตำแหน่งละติจูดของการวัด';
COMMENT ON COLUMN public.crosssection.section_long IS 'ตำแหน่งลองติจูดของการวัด';
COMMENT ON COLUMN public.crosssection.distance IS 'ระยะทาง หน่วย : เมตร distance';
COMMENT ON COLUMN public.crosssection.water_level_msl IS 'ระดับน้ำ หน่วย : ม.รทก. water level (msl)';
COMMENT ON COLUMN public.crosssection.water_level_m IS 'ระดับน้ำ หน่วย : ม water level (m)';
COMMENT ON COLUMN public.crosssection.remark IS 'ตำแหน่งที่วัด LB : left bank RB : right bank  CL / location';
COMMENT ON COLUMN public.crosssection.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.crosssection.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.crosssection.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.crosssection.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.crosssection.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.crosssection.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.crosssection.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.crosssection.deleted_at IS 'วันที่ลบข้อมูล deleted date';

