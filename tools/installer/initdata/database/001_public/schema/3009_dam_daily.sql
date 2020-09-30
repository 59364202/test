-- Table: public.dam_daily

-- DROP TABLE public.dam_daily;

CREATE TABLE public.dam_daily
(
  id bigserial NOT NULL, -- รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. รายวัน dam (daily)'s serial number
  dam_id bigint, -- รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. dam's serial number
  dam_date date NOT NULL, -- วันที่เก็บข้อมูล record date
  dam_level double precision, -- ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.) last water level
  dam_storage double precision, -- ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
  dam_inflow double precision, -- ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
  dam_released double precision, -- ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
  dam_spilled double precision, -- ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง spilled water volume
  dam_losses double precision, -- ปริมาณน้ำที่สูญเสีย loss water volume
  dam_evap double precision, -- ปริมาณน้ำที่ระเหย evaporated water volume
  dam_uses_water double precision, -- ปริมาณน้ำที่ใช้ได้ uses water volume
  dam_storage_percent double precision, -- เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume
  dam_uses_water_percent double precision, -- เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.) data form rid not cal/ percent of uses water volume
  dam_inflow_avg double precision, -- ปริมาตรน้ำไหลลงอ่างเก็บน้ำสะสมตั้งแต่ต้นปี
  dam_released_acc double precision, -- ปริมาตรน้ำระบายสะสมตั้งแต่ต้นปี
  dam_inflow_acc double precision, -- ปริมาตรน้ำไหลลงอ่างเก็บน้ำเฉลี่ยทั้งปี
  dam_inflow_acc_percent double precision, -- เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%) data form rid not cal/ percent of inflowing water volume
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_dam_daily PRIMARY KEY (id),
  CONSTRAINT fk_dam_dail_reference_m_dam FOREIGN KEY (dam_id)
      REFERENCES public.m_dam (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_dam_daily UNIQUE (dam_id, dam_date, deleted_at),
  CONSTRAINT pt_dam_daily_dam_date CHECK (false) NO INHERIT
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.dam_daily
  IS 'เขื่อนขนาดใหญ่ รายวัน';
COMMENT ON COLUMN public.dam_daily.id IS 'รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. รายวัน dam (daily)''s serial number';
COMMENT ON COLUMN public.dam_daily.dam_id IS 'รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. dam''s serial number';
COMMENT ON COLUMN public.dam_daily.dam_date IS 'วันที่เก็บข้อมูล record date';
COMMENT ON COLUMN public.dam_daily.dam_level IS 'ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.) last water level';
COMMENT ON COLUMN public.dam_daily.dam_storage IS 'ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume';
COMMENT ON COLUMN public.dam_daily.dam_inflow IS 'ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume';
COMMENT ON COLUMN public.dam_daily.dam_released IS 'ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume';
COMMENT ON COLUMN public.dam_daily.dam_spilled IS 'ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง spilled water volume';
COMMENT ON COLUMN public.dam_daily.dam_losses IS 'ปริมาณน้ำที่สูญเสีย loss water volume';
COMMENT ON COLUMN public.dam_daily.dam_evap IS 'ปริมาณน้ำที่ระเหย evaporated water volume';
COMMENT ON COLUMN public.dam_daily.dam_uses_water IS 'ปริมาณน้ำที่ใช้ได้ uses water volume';
COMMENT ON COLUMN public.dam_daily.dam_storage_percent IS 'เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume';
COMMENT ON COLUMN public.dam_daily.dam_uses_water_percent IS 'เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.) data form rid not cal/ percent of uses water volume';
COMMENT ON COLUMN public.dam_daily.dam_inflow_avg IS 'ปริมาตรน้ำไหลลงอ่างเก็บน้ำสะสมตั้งแต่ต้นปี';
COMMENT ON COLUMN public.dam_daily.dam_released_acc IS 'ปริมาตรน้ำระบายสะสมตั้งแต่ต้นปี';
COMMENT ON COLUMN public.dam_daily.dam_inflow_acc IS 'ปริมาตรน้ำไหลลงอ่างเก็บน้ำเฉลี่ยทั้งปี';
COMMENT ON COLUMN public.dam_daily.dam_inflow_acc_percent IS 'เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%) data form rid not cal/ percent of inflowing water volume';
COMMENT ON COLUMN public.dam_daily.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.dam_daily.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.dam_daily.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.dam_daily.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.dam_daily.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.dam_daily.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.dam_daily.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.dam_daily.deleted_at IS 'วันที่ลบข้อมูล deleted date';

