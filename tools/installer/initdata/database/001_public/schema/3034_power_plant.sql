-- Table: public.power_plant

-- DROP TABLE public.power_plant;

CREATE TABLE public.power_plant
(
  id bigserial NOT NULL, -- รหัสโรงไฟฟ้า power plant's serial number
  agency_id bigint, -- รหัสหน่วยงาน agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  power_plant_oldcode text, -- รหัสเดิมของโรงไฟฟ้า old power plant's serial number
  power_plant_name text, -- ชื่อโรงไฟฟ้า power plant's name
  power_plant_location text, -- สถานที่ตั้งโรงไฟฟ้า location
  power_plant_lat numeric(9,6), -- ละติจูดของโรงไฟฟ้า latitude
  power_plant_long numeric(9,6), -- ลองติจูดของโรงไฟฟ้า longitude
  power_plant_type text, -- ประเภทของโรงไฟฟ้า เช่น โรงไฟฟ้าพลังความร้อนร่วม โรงไฟฟ้าพลังความร้อน โรงไฟฟ้ากังหันแก๊ส หรือโรงไฟฟ้าพลังน้ำ power plant type
  power_producer_status text, -- ประเภทสถานะของผู้ผลิตไฟฟ้า power producer status
  capacity_mw double precision, -- กำลังผลิตติดตั้งของโรงไฟฟ้า (Installed Capacity) หน่วย: Mega Watt / installed capacity
  sold_mw double precision, -- กำลังไฟฟ้าที่ขายเข้าระบบส่ง (Sold to grid) หน่วย : Mega Watt / sold to grid
  fuel text, -- ประเภทเชื้อเพลิงหลักที่ใช้ในการผลิตไฟฟ้า fuel type
  secon_fuel text, -- เชื้อเพลิงสำรองที่ใช้ในการผลิตไฟฟ้า secondary fuel
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  qc_status json, -- สถานะของการตรวจคุณภาพข้อมูล quality control status
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_power_plant PRIMARY KEY (id),
  CONSTRAINT fk_power_pl_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_power_pl_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_power_plant UNIQUE (power_plant_oldcode)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.power_plant
  IS 'โรงไฟฟ้า  เป็นข้อมูลที่โอนถ่ายมาจาก nhc ซึ่งในปัจจุบันจะจัดเก็บเป็นข้อมูลสื่อ Shape File';
COMMENT ON COLUMN public.power_plant.id IS 'รหัสโรงไฟฟ้า power plant''s serial number';
COMMENT ON COLUMN public.power_plant.agency_id IS 'รหัสหน่วยงาน agency''s serial number';
COMMENT ON COLUMN public.power_plant.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.power_plant.power_plant_oldcode IS 'รหัสเดิมของโรงไฟฟ้า old power plant''s serial number';
COMMENT ON COLUMN public.power_plant.power_plant_name IS 'ชื่อโรงไฟฟ้า power plant''s name';
COMMENT ON COLUMN public.power_plant.power_plant_location IS 'สถานที่ตั้งโรงไฟฟ้า location';
COMMENT ON COLUMN public.power_plant.power_plant_lat IS 'ละติจูดของโรงไฟฟ้า latitude';
COMMENT ON COLUMN public.power_plant.power_plant_long IS 'ลองติจูดของโรงไฟฟ้า longitude';
COMMENT ON COLUMN public.power_plant.power_plant_type IS 'ประเภทของโรงไฟฟ้า เช่น โรงไฟฟ้าพลังความร้อนร่วม โรงไฟฟ้าพลังความร้อน โรงไฟฟ้ากังหันแก๊ส หรือโรงไฟฟ้าพลังน้ำ power plant type';
COMMENT ON COLUMN public.power_plant.power_producer_status IS 'ประเภทสถานะของผู้ผลิตไฟฟ้า power producer status';
COMMENT ON COLUMN public.power_plant.capacity_mw IS 'กำลังผลิตติดตั้งของโรงไฟฟ้า (Installed Capacity) หน่วย: Mega Watt / installed capacity';
COMMENT ON COLUMN public.power_plant.sold_mw IS 'กำลังไฟฟ้าที่ขายเข้าระบบส่ง (Sold to grid) หน่วย : Mega Watt / sold to grid';
COMMENT ON COLUMN public.power_plant.fuel IS 'ประเภทเชื้อเพลิงหลักที่ใช้ในการผลิตไฟฟ้า fuel type';
COMMENT ON COLUMN public.power_plant.secon_fuel IS 'เชื้อเพลิงสำรองที่ใช้ในการผลิตไฟฟ้า secondary fuel';
COMMENT ON COLUMN public.power_plant.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.power_plant.qc_status IS 'สถานะของการตรวจคุณภาพข้อมูล quality control status';
COMMENT ON COLUMN public.power_plant.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.power_plant.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.power_plant.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.power_plant.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.power_plant.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.power_plant.deleted_at IS 'วันที่ลบข้อมูล deleted date';

