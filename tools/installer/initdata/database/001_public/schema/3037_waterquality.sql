-- Table: public.waterquality

-- DROP TABLE public.waterquality;

CREATE TABLE public.waterquality
(
  id bigserial NOT NULL, -- รหัสข้อมูลคุณภาพน้ำจากการวัดของสถานี
  waterquality_id bigint, -- รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ
  waterquality_datetime timestamp with time zone NOT NULL, -- วันที่ตรวจสอบค่าคุณภาพน้ำอัตโนมัติ
  waterquality_do double precision, -- ออกซิเจนละลายในน้ำ หน่วย mg/l
  waterquality_conductivity double precision, -- ความนำไฟฟ้าในน้ำ หน่วย uS/cm ชื่อเต็ม The Electrical Conductivity (ec)
  waterquality_ph double precision, -- ความเป็นกรด-ด่าง
  waterquality_temp double precision, -- อุณหภูมิน้ำ หน่วย ?C
  waterquality_turbid double precision, -- ค่าความขุ่นในน้ำ หน่วย NTU
  waterquality_bod double precision, -- ค่าความสกปรกในรูปสารอินทรีย์ หน่วย mg/l
  waterquality_tcb double precision, -- ปริมาณแบคทีเรียในรูปโคลิฟอร์มทั้งหมด หน่วย MPN/100 ml
  waterquality_fcb double precision, -- ปริมาณแบคทีเรียในรูปฟีคลอโคลิฟอร์ม หน่วย MPN/100 ml
  waterquality_nh3n double precision, -- ปริมาณแอมโมเนีย-ไนโตรเจน หน่วย mg/l
  waterquality_wqi double precision, -- ช่วงคะแนน WQI
  waterquality_ammonium double precision, -- ปริมาณแอมโมเนีย
  waterquality_nitrate double precision, -- ไนโตรเจน
  waterquality_salinity double precision, -- ค่าความเค็ม
  waterquality_tds double precision, -- ค่า tds
  waterquality_chlorophyll double precision, -- คลอโรฟิลด์
  waterquality_colorstatus text, -- สถานะของสี
  waterquality_status text, -- สถานะของคุณภาพน้ำ
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_waterquality PRIMARY KEY (id),
  CONSTRAINT fk_waterqua_reference_m_waterq FOREIGN KEY (waterquality_id)
      REFERENCES public.m_waterquality_station (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_waterquality UNIQUE (waterquality_id, waterquality_datetime, deleted_at),
  CONSTRAINT pt_waterquality_waterquality_datetime CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);
COMMENT ON TABLE public.waterquality
  IS 'ค่าตรวจวัดคุณภาพน้ำ';
COMMENT ON COLUMN public.waterquality.id IS 'รหัสข้อมูลคุณภาพน้ำจากการวัดของสถานี';
COMMENT ON COLUMN public.waterquality.waterquality_id IS 'รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ ';
COMMENT ON COLUMN public.waterquality.waterquality_datetime IS 'วันที่ตรวจสอบค่าคุณภาพน้ำอัตโนมัติ ';
COMMENT ON COLUMN public.waterquality.waterquality_do IS 'ออกซิเจนละลายในน้ำ หน่วย mg/l';
COMMENT ON COLUMN public.waterquality.waterquality_conductivity IS 'ความนำไฟฟ้าในน้ำ หน่วย uS/cm ชื่อเต็ม The Electrical Conductivity (ec)';
COMMENT ON COLUMN public.waterquality.waterquality_ph IS 'ความเป็นกรด-ด่าง ';
COMMENT ON COLUMN public.waterquality.waterquality_temp IS 'อุณหภูมิน้ำ หน่วย ?C';
COMMENT ON COLUMN public.waterquality.waterquality_turbid IS 'ค่าความขุ่นในน้ำ หน่วย NTU';
COMMENT ON COLUMN public.waterquality.waterquality_bod IS 'ค่าความสกปรกในรูปสารอินทรีย์ หน่วย mg/l';
COMMENT ON COLUMN public.waterquality.waterquality_tcb IS 'ปริมาณแบคทีเรียในรูปโคลิฟอร์มทั้งหมด หน่วย MPN/100 ml';
COMMENT ON COLUMN public.waterquality.waterquality_fcb IS 'ปริมาณแบคทีเรียในรูปฟีคลอโคลิฟอร์ม หน่วย MPN/100 ml';
COMMENT ON COLUMN public.waterquality.waterquality_nh3n IS 'ปริมาณแอมโมเนีย-ไนโตรเจน หน่วย mg/l';
COMMENT ON COLUMN public.waterquality.waterquality_wqi IS 'ช่วงคะแนน WQI';
COMMENT ON COLUMN public.waterquality.waterquality_ammonium IS 'ปริมาณแอมโมเนีย';
COMMENT ON COLUMN public.waterquality.waterquality_nitrate IS 'ไนโตรเจน';
COMMENT ON COLUMN public.waterquality.waterquality_salinity IS 'ค่าความเค็ม';
COMMENT ON COLUMN public.waterquality.waterquality_tds IS 'ค่า tds';
COMMENT ON COLUMN public.waterquality.waterquality_chlorophyll IS 'คลอโรฟิลด์';
COMMENT ON COLUMN public.waterquality.waterquality_colorstatus IS 'สถานะของสี';
COMMENT ON COLUMN public.waterquality.waterquality_status IS 'สถานะของคุณภาพน้ำ';
COMMENT ON COLUMN public.waterquality.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.waterquality.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.waterquality.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.waterquality.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.waterquality.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.waterquality.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.waterquality.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.waterquality.deleted_at IS 'วันที่ลบข้อมูล deleted date';

