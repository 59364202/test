-- Table: public.dbamodule_history

-- DROP TABLE public.dbamodule_history;

CREATE TABLE public.dbamodule_history
(
  id bigserial NOT NULL, -- รหัสประวัติของการจัดการโมดูล DBA
  table_name text, -- ชื่อตารางที่จัดการโมดูล DBA
  year text, -- ปีที่การจัดการโมดูล DBA
  month text, -- เดือนที่การจัดการโมดูล DBA
  dba_remark text, -- รายละเอียดการจัดการโมดูล DBA
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_dbamodule_history PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.dbamodule_history
  IS 'ประวัติของการจัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.id IS 'รหัสประวัติของการจัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.table_name IS 'ชื่อตารางที่จัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.year IS 'ปีที่การจัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.month IS 'เดือนที่การจัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.dba_remark IS 'รายละเอียดการจัดการโมดูล DBA';
COMMENT ON COLUMN public.dbamodule_history.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.dbamodule_history.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.dbamodule_history.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.dbamodule_history.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.dbamodule_history.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.dbamodule_history.deleted_at IS 'วันที่ลบข้อมูล deleted date';

