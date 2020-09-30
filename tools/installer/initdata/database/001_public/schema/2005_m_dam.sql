-- Table: public.m_dam

-- DROP TABLE public.m_dam;

CREATE TABLE public.m_dam
(
  id bigserial NOT NULL, -- รหัสข้อมูลเขื่อน dam's serial number
  subbasin_id bigint, -- ลำดับลุ่มน้ำสาขา subbasin's serial number
  agency_id bigint, -- รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
  geocode_id bigint, -- ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode's serial number
  dam_name json, -- ชื่อเขื่อน
  dam_oldcode text, -- รหัสเดิมของเขื่อนขนาดใหญ่ old dam's serial number
  dam_lat numeric(9,6), -- พิกัดของเขื่อน latitude
  dam_long numeric(9,6), -- พิกัดของเขื่อน longitude
  max_water_level double precision, -- ระดับกักเก็บสูงสุด [ม.(รทก.)] max water level
  normal_water_level double precision, -- ระดับกักเก็บปกติ [ม.(รทก.)] normal water level
  min_water_level double precision, -- ระดับกักเก็บต่ำสุด [ม.(รทก.)] min water level
  normal_watergate_level double precision, -- เริ่มเปิดบานระบายน้ำที่ระดับ normal  [ม.(รทก.)] normal watergate level
  emer_watergate_level double precision, -- เริ่มเปิดบานระบายน้ำที่ระดับ Emergency  [ม.(รทก.)] mergency watergate level
  service_watergate_level double precision, -- เริ่มเปิดบานระบายน้ำที่ระดับ service [ม.(รทก.)] service watergate level
  max_old_storage double precision, -- ระดับน้ำสูงสุดที่เคยเก็บกัก[ม.(รทก.)] max old storage
  maxos_date date, -- วันที่เริ่มจัดเก็บระดับน้ำสูงสุดที่เคยเก็บกัก max old storage date
  min_old_storage double precision, -- ระดับน้ำต่ำสุดที่เคยเก็บกัก [ม.(รทก.)] min old storage
  minos_date date, -- วันที่เริ่มจัดเก็บระดับน้ำต่ำสุดที่เคยเก็บกัก min old storage date
  top_spillway_level double precision, -- ระดับขอบบนของบานประตู spillway [ม.(รทก.)] top spillway level
  ridge_spillway_level double precision, -- ระดับสันของบานประตู spillway  [ม.(รทก.)] ridge spillway level
  max_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.] max storage
  normal_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.] normal storage
  min_storage double precision, -- ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด [ล้าน ลบ.ม.] min storage
  uses_water double precision, -- ปริมาตรน้ำที่ใช้งานได้ [ล้าน ลบ.ม.] uses water
  avg_inflow double precision, -- ปริมาณน้ำไหลเข้าเฉลี่ย [ล้าน ลบ.ม.] average inflowing water
  avg_inflow_intyear text, -- ปีเริ่มต้นที่บันทึกปริมาณน้ำไหลเข้าเฉลี่ย start year (average inflowing water)
  avg_inflow_endyear text, -- ปีปัจจุบันปริมาณน้ำไหลเข้าเฉลี่ย end year (average inflowing water)
  max_inflow double precision, -- ปริมาณน้ำไหลเข้าสูงสุด [ล้าน ลบ.ม.] max inflowing water
  max_inflow_date date, -- วันที่เริ่มบันทึกปริมาณน้ำไหลเข้าสูงสุด start date (max inflowing water)
  downstream_storage double precision, -- ความจุท้ายน้ำ [ลบ.ม./วินาที] downstream storage
  water_shed double precision, -- พื้นที่รับน้ำ [ตร. กม.] water shed
  rainfall_yearly double precision, -- ปริมาณน้ำฝนเฉลี่ยต่อปี [มม.] rainfall (yearly)
  power_install double precision, -- กำลังผลิตติดตั้ง [MW]  power install
  power_intake_storage double precision, -- Storage at power intake sill level  [ล้าน ลบ.ม.] storage at power intake sill level
  power_intake_level double precision, -- Power intake sill Level [ม.(รทก.)] power intake sill level
  tailrace_level double precision, -- Tailace Normal Level  [ม.(รทก.)] tailace normal level
  used_genpower double precision, -- การใช้น้ำในการผลิตต่อหน่วย [cms/kwhr] used water power
  is_ignore boolean DEFAULT false, -- สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)
  dataimport_log_id bigint, -- Define import id รหัสของการ Import ข้อมูล importing data's ID
  created_by bigint, -- ชื่อผู้ใช้งานที่สร้างข้อมูล created user
  created_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่สร้างข้อมูล created date
  updated_by bigint, -- ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user
  updated_at timestamp with time zone NOT NULL DEFAULT now(), -- วันที่ปรับปรุงข้อมูลล่าสุด updated date
  deleted_by bigint, -- ชื่อผู้ใช้งานที่ลบข้อมูล deleted user
  deleted_at timestamp with time zone NOT NULL DEFAULT '1970-01-01 07:00:00+07'::timestamp with time zone, -- วันที่ลบข้อมูล deleted date
  CONSTRAINT pk_m_dam PRIMARY KEY (id),
  CONSTRAINT fk_m_dam_reference_agency FOREIGN KEY (agency_id)
      REFERENCES public.agency (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_dam_reference_lt_geoco FOREIGN KEY (geocode_id)
      REFERENCES public.lt_geocode (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT fk_m_dam_reference_subbasin FOREIGN KEY (subbasin_id)
      REFERENCES public.subbasin (id) MATCH SIMPLE
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT uk_m_dam UNIQUE (dam_oldcode, agency_id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.m_dam
  IS 'ข้อมูลพืนฐานเขื่อน';
COMMENT ON COLUMN public.m_dam.id IS 'รหัสข้อมูลเขื่อน dam''s serial number';
COMMENT ON COLUMN public.m_dam.subbasin_id IS 'ลำดับลุ่มน้ำสาขา subbasin''s serial number';
COMMENT ON COLUMN public.m_dam.agency_id IS 'รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency''s serial number';
COMMENT ON COLUMN public.m_dam.geocode_id IS 'ลำดับข้อมูลขอบเขตการปกครองของประเทศไทย Geocode''s serial number';
COMMENT ON COLUMN public.m_dam.dam_name IS 'ชื่อเขื่อน';
COMMENT ON COLUMN public.m_dam.dam_oldcode IS 'รหัสเดิมของเขื่อนขนาดใหญ่ old dam''s serial number';
COMMENT ON COLUMN public.m_dam.dam_lat IS 'พิกัดของเขื่อน latitude';
COMMENT ON COLUMN public.m_dam.dam_long IS 'พิกัดของเขื่อน longitude';
COMMENT ON COLUMN public.m_dam.max_water_level IS 'ระดับกักเก็บสูงสุด [ม.(รทก.)] max water level';
COMMENT ON COLUMN public.m_dam.normal_water_level IS 'ระดับกักเก็บปกติ [ม.(รทก.)] normal water level';
COMMENT ON COLUMN public.m_dam.min_water_level IS 'ระดับกักเก็บต่ำสุด [ม.(รทก.)] min water level';
COMMENT ON COLUMN public.m_dam.normal_watergate_level IS 'เริ่มเปิดบานระบายน้ำที่ระดับ normal  [ม.(รทก.)] normal watergate level';
COMMENT ON COLUMN public.m_dam.emer_watergate_level IS 'เริ่มเปิดบานระบายน้ำที่ระดับ Emergency  [ม.(รทก.)] mergency watergate level';
COMMENT ON COLUMN public.m_dam.service_watergate_level IS 'เริ่มเปิดบานระบายน้ำที่ระดับ service [ม.(รทก.)] service watergate level';
COMMENT ON COLUMN public.m_dam.max_old_storage IS 'ระดับน้ำสูงสุดที่เคยเก็บกัก[ม.(รทก.)] max old storage';
COMMENT ON COLUMN public.m_dam.maxos_date IS 'วันที่เริ่มจัดเก็บระดับน้ำสูงสุดที่เคยเก็บกัก max old storage date';
COMMENT ON COLUMN public.m_dam.min_old_storage IS 'ระดับน้ำต่ำสุดที่เคยเก็บกัก [ม.(รทก.)] min old storage';
COMMENT ON COLUMN public.m_dam.minos_date IS 'วันที่เริ่มจัดเก็บระดับน้ำต่ำสุดที่เคยเก็บกัก min old storage date';
COMMENT ON COLUMN public.m_dam.top_spillway_level IS 'ระดับขอบบนของบานประตู spillway [ม.(รทก.)] top spillway level';
COMMENT ON COLUMN public.m_dam.ridge_spillway_level IS 'ระดับสันของบานประตู spillway  [ม.(รทก.)] ridge spillway level';
COMMENT ON COLUMN public.m_dam.max_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.] max storage';
COMMENT ON COLUMN public.m_dam.normal_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.] normal storage';
COMMENT ON COLUMN public.m_dam.min_storage IS 'ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด [ล้าน ลบ.ม.] min storage';
COMMENT ON COLUMN public.m_dam.uses_water IS 'ปริมาตรน้ำที่ใช้งานได้ [ล้าน ลบ.ม.] uses water';
COMMENT ON COLUMN public.m_dam.avg_inflow IS 'ปริมาณน้ำไหลเข้าเฉลี่ย [ล้าน ลบ.ม.] average inflowing water';
COMMENT ON COLUMN public.m_dam.avg_inflow_intyear IS 'ปีเริ่มต้นที่บันทึกปริมาณน้ำไหลเข้าเฉลี่ย start year (average inflowing water)';
COMMENT ON COLUMN public.m_dam.avg_inflow_endyear IS 'ปีปัจจุบันปริมาณน้ำไหลเข้าเฉลี่ย end year (average inflowing water)';
COMMENT ON COLUMN public.m_dam.max_inflow IS 'ปริมาณน้ำไหลเข้าสูงสุด [ล้าน ลบ.ม.] max inflowing water';
COMMENT ON COLUMN public.m_dam.max_inflow_date IS 'วันที่เริ่มบันทึกปริมาณน้ำไหลเข้าสูงสุด start date (max inflowing water)';
COMMENT ON COLUMN public.m_dam.downstream_storage IS 'ความจุท้ายน้ำ [ลบ.ม./วินาที] downstream storage';
COMMENT ON COLUMN public.m_dam.water_shed IS 'พื้นที่รับน้ำ [ตร. กม.] water shed';
COMMENT ON COLUMN public.m_dam.rainfall_yearly IS 'ปริมาณน้ำฝนเฉลี่ยต่อปี [มม.] rainfall (yearly)';
COMMENT ON COLUMN public.m_dam.power_install IS 'กำลังผลิตติดตั้ง [MW]  power install';
COMMENT ON COLUMN public.m_dam.power_intake_storage IS 'Storage at power intake sill level  [ล้าน ลบ.ม.] storage at power intake sill level';
COMMENT ON COLUMN public.m_dam.power_intake_level IS 'Power intake sill Level [ม.(รทก.)] power intake sill level';
COMMENT ON COLUMN public.m_dam.tailrace_level IS 'Tailace Normal Level  [ม.(รทก.)] tailace normal level';
COMMENT ON COLUMN public.m_dam.used_genpower IS 'การใช้น้ำในการผลิตต่อหน่วย [cms/kwhr] used water power';
COMMENT ON COLUMN public.m_dam.is_ignore IS 'สถานีที่ต้องแสดงบนหน้าจอแสดงผลหรือไม่ (true/false)';
COMMENT ON COLUMN public.m_dam.dataimport_log_id IS 'Define import id รหัสของการ Import ข้อมูล importing data''s ID';
COMMENT ON COLUMN public.m_dam.created_by IS 'ชื่อผู้ใช้งานที่สร้างข้อมูล created user';
COMMENT ON COLUMN public.m_dam.created_at IS 'วันที่สร้างข้อมูล created date';
COMMENT ON COLUMN public.m_dam.updated_by IS 'ชื่อผู้ใช้งานที่ปรับปรุงข้อมูล updated user';
COMMENT ON COLUMN public.m_dam.updated_at IS 'วันที่ปรับปรุงข้อมูลล่าสุด updated date';
COMMENT ON COLUMN public.m_dam.deleted_by IS 'ชื่อผู้ใช้งานที่ลบข้อมูล deleted user';
COMMENT ON COLUMN public.m_dam.deleted_at IS 'วันที่ลบข้อมูล deleted date';

