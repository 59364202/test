-- Table: public.tr_regional_avg_rainfall_y30

-- DROP TABLE public.tr_regional_avg_rainfall_y30;

CREATE TABLE public.tr_regional_avg_rainfall_y30
(
  reg_id character(2) NOT NULL, -- รหัสภาค (กรณีรวมทั้งประเทศใช้รหัส '00')
  volume double precision NOT NULL, -- ค่าฝนเฉลี่ย
  created_date date, -- วันที่สร้างข้อมูล
  created_by character varying(50), -- ผู้สร้างข้อมูล
  last_updated_date date, -- วันที่เปลี่ยนแปลงข้อมูลล่าสุด
  last_updated_by character varying(50), -- ผู้เปลี่ยนแปลงข้อมูล
  CONSTRAINT pk_tr_regional_avg_rainfall_y3 PRIMARY KEY (reg_id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.tr_regional_avg_rainfall_y30
  IS 'ค่าฝนเฉลี่ยรายภาค-รายปี';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.reg_id IS 'รหัสภาค (กรณีรวมทั้งประเทศใช้รหัส ''00'')';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.volume IS 'ค่าฝนเฉลี่ย';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.created_date IS 'วันที่สร้างข้อมูล';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.created_by IS 'ผู้สร้างข้อมูล';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.last_updated_date IS 'วันที่เปลี่ยนแปลงข้อมูลล่าสุด';
COMMENT ON COLUMN public.tr_regional_avg_rainfall_y30.last_updated_by IS 'ผู้เปลี่ยนแปลงข้อมูล';

