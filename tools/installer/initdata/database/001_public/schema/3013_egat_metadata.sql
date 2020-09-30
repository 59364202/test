-- Table: public.egat_metadata

-- DROP TABLE public.egat_metadata;

CREATE TABLE public.egat_metadata
(
  id bigserial NOT NULL, -- รหัสรายละเอียดสถานีของโทรมาตร กฟผ. Egat metadata 's serial number
  tele_station_id bigint NOT NULL, -- รหัสสถานีโทรมาตร tele station's serial number
  egat_deviceid smallint NOT NULL, -- รหัส Device ตาม Web service
  egat_stationsi smallint, -- รหัสสถานีตาม Web service
  egat_filename text, -- ชื่อกลุ่มโทรมาตร โดยตัดคำจากชื่อไฟล์ ได้แก่ bhu sk mk ur pm rp
  egat_namenv text, -- รายละเอียดสถานีตาม Service
  egat_devicesi smallint, -- ลำดับ Device ตาม Web service
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_egat_metadata PRIMARY KEY (id),
  CONSTRAINT fk_egat_met_reference_m_tele_s FOREIGN KEY (tele_station_id)
      REFERENCES public.m_tele_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_egat_metadata UNIQUE (tele_station_id, egat_deviceid, egat_stationsi, deleted_at)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.egat_metadata
  IS 'รายละเอียดสถานีของโทรมาตร กฟผ.';
COMMENT ON COLUMN public.egat_metadata.id IS 'รหัสรายละเอียดสถานีของโทรมาตร กฟผ. Egat metadata ''s serial number';
COMMENT ON COLUMN public.egat_metadata.tele_station_id IS 'รหัสสถานีโทรมาตร tele station''s serial number';
COMMENT ON COLUMN public.egat_metadata.egat_deviceid IS 'รหัส Device ตาม Web service';
COMMENT ON COLUMN public.egat_metadata.egat_stationsi IS 'รหัสสถานีตาม Web service';
COMMENT ON COLUMN public.egat_metadata.egat_filename IS 'ชื่อกลุ่มโทรมาตร โดยตัดคำจากชื่อไฟล์ ได้แก่ bhu sk mk ur pm rp';
COMMENT ON COLUMN public.egat_metadata.egat_namenv IS 'รายละเอียดสถานีตาม Service';
COMMENT ON COLUMN public.egat_metadata.egat_devicesi IS 'ลำดับ Device ตาม Web service';
COMMENT ON COLUMN public.egat_metadata.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.egat_metadata.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.egat_metadata.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.egat_metadata.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.egat_metadata.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.egat_metadata.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.egat_metadata.deleted_at IS 'วันที่ลบข้อมูล deleted date';

