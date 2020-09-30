-- Table: public.swan

-- DROP TABLE public.swan;

CREATE TABLE public.swan
(
  id bigserial NOT NULL, -- รหัสข้อมูลคาดการณ์ความสูงคลื่น
  swan_station_id bigint, -- รหัสสถานีคาดการณ์ความสูงคลื่น
  swan_datetime timestamp with time zone NOT NULL, -- วันที่และเวลาที่เก็บข้อมูลคาดการณ์ความสูงคลื่น
  swan_depth double precision, -- ข้อมูลคาดการณ์ความลึกของคลื่น หน่วย m
  swan_highsig double precision, -- ข้อมูลคาดการณ์ความสูงของคลื่น หน่วย m
  swan_direction double precision, -- ข้อมูลคาดการณ์ทิศทางของคลื่น หน่วย degree
  swan_period_top double precision, -- ข้อมูลคาดการณ์คาบคลื่นสูงสุด หน่วย sec
  swan_period_average double precision, -- ข้อมูลคาดการณ์คาบคลื่นเฉลี่ย หน่วย sec
  swan_windx double precision, -- ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศตะวันออกและทิศตะวันตก หน่วย m/s
  swan_windy double precision, -- ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศเเหนือและใต้ หน่วย m/s
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_swan PRIMARY KEY (id),
  CONSTRAINT fk_swan_reference_m_swan_s FOREIGN KEY (swan_station_id)
      REFERENCES public.m_swan_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_swan UNIQUE (swan_station_id, swan_datetime, deleted_at),
  CONSTRAINT pt_swan_swan_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.swan
  IS 'ข้อมูลคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.swan.id IS 'รหัสข้อมูลคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.swan.swan_station_id IS 'รหัสสถานีคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.swan.swan_datetime IS 'วันที่และเวลาที่เก็บข้อมูลคาดการณ์ความสูงคลื่น';
COMMENT ON COLUMN public.swan.swan_depth IS 'ข้อมูลคาดการณ์ความลึกของคลื่น หน่วย m';
COMMENT ON COLUMN public.swan.swan_highsig IS 'ข้อมูลคาดการณ์ความสูงของคลื่น หน่วย m';
COMMENT ON COLUMN public.swan.swan_direction IS 'ข้อมูลคาดการณ์ทิศทางของคลื่น หน่วย degree';
COMMENT ON COLUMN public.swan.swan_period_top IS 'ข้อมูลคาดการณ์คาบคลื่นสูงสุด หน่วย sec';
COMMENT ON COLUMN public.swan.swan_period_average IS 'ข้อมูลคาดการณ์คาบคลื่นเฉลี่ย หน่วย sec';
COMMENT ON COLUMN public.swan.swan_windx IS 'ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศตะวันออกและทิศตะวันตก หน่วย m/s';
COMMENT ON COLUMN public.swan.swan_windy IS 'ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศเเหนือและใต้ หน่วย m/s';
COMMENT ON COLUMN public.swan.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.swan.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.swan.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.swan.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.swan.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.swan.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.swan.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.swan.deleted_at IS 'วันที่ลบข้อมูล deleted date';

