-- Table: public.ignore_history

-- DROP TABLE public.ignore_history;

CREATE TABLE public.ignore_history
(
  id bigserial NOT NULL, -- รหัสประวัติของการ ignore ข้อมูล
  ignore_datetime timestamp with time zone, -- วันและเวลาที่ Ignore ข้อมูล
  data_category text, -- ประเภขข้อมูล คือ ฝน เขื่อน ระดับน้ำ และคุณภาพน้ำ
  station_id text, -- รหัสสถานี
  station_oldcode text, -- รหัสสถานีเดิม
  station_name text, -- ชื่อของสถานี
  station_province text, -- ชื่อจังหวัดของสถานีที่ติดตั้ง
  agency_shortname text, -- ชื่อย่อหน่วยงานของสถานีที่ติดตั้ง
  data_datetime timestamp with time zone, -- วันและเวลาของข้อมูล
  data_value double precision, -- ค่าของข้อมูล
  data_id bigint, -- รหัสของข้อมูลที่ ignore
  remark text, -- หมายเหตุ
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_ignore_history PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.ignore_history
  IS 'ประวัติของการ ignore ข้อมูล';
COMMENT ON COLUMN public.ignore_history.id IS 'รหัสประวัติของการ ignore ข้อมูล';
COMMENT ON COLUMN public.ignore_history.ignore_datetime IS 'วันและเวลาที่ Ignore ข้อมูล';
COMMENT ON COLUMN public.ignore_history.data_category IS 'ประเภขข้อมูล คือ ฝน เขื่อน ระดับน้ำ และคุณภาพน้ำ';
COMMENT ON COLUMN public.ignore_history.station_id IS 'รหัสสถานี';
COMMENT ON COLUMN public.ignore_history.station_oldcode IS 'รหัสสถานีเดิม';
COMMENT ON COLUMN public.ignore_history.station_name IS 'ชื่อของสถานี';
COMMENT ON COLUMN public.ignore_history.station_province IS 'ชื่อจังหวัดของสถานีที่ติดตั้ง';
COMMENT ON COLUMN public.ignore_history.agency_shortname IS 'ชื่อย่อหน่วยงานของสถานีที่ติดตั้ง';
COMMENT ON COLUMN public.ignore_history.data_datetime IS 'วันและเวลาของข้อมูล';
COMMENT ON COLUMN public.ignore_history.data_value IS 'ค่าของข้อมูล';
COMMENT ON COLUMN public.ignore_history.data_id IS 'รหัสของข้อมูลที่ ignore';
COMMENT ON COLUMN public.ignore_history.remark IS 'หมายเหตุ';
COMMENT ON COLUMN public.ignore_history.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.ignore_history.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.ignore_history.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.ignore_history.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.ignore_history.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.ignore_history.deleted_at IS 'วันที่ลบข้อมูล deleted date';

