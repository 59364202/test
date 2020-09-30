-- Table: public.tr_avg_rainfall_y30

-- DROP TABLE public.tr_avg_rainfall_y30;

CREATE TABLE public.tr_avg_rainfall_y30
(
  reg_id character(2) NOT NULL, -- รหัสภาค (กรณีรวมทั้งประเทศใช้รหัส '00')
  month_id character(2) NOT NULL, -- รหัสเดือน
  volume double precision NOT NULL, -- ปริมาณฝนเฉลี่ย
  created_date date, -- วันที่สร้างข้อมูล
  created_by character varying(50), -- ผู้สร้างข้อมูล
  last_updated_date date, -- วันที่เปลี่ยนแปลงข้อมูลล่าสุด
  last_updated_by character varying(50), -- ผู้เปลี่ยนแปลงข้อมูล
  CONSTRAINT pk_tr_avg_rainfall_y30 PRIMARY KEY (reg_id, month_id)
)
WITH (
  OIDS=FALSE
);

COMMENT ON TABLE public.tr_avg_rainfall_y30
  IS 'ปริมาณฝนเฉลี่ยรายภาค-รายเดือน';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.reg_id IS 'รหัสภาค (กรณีรวมทั้งประเทศใช้รหัส ''00'')';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.month_id IS 'รหัสเดือน';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.volume IS 'ปริมาณฝนเฉลี่ย';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.created_date IS 'วันที่สร้างข้อมูล';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.created_by IS 'ผู้สร้างข้อมูล';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.last_updated_date IS 'วันที่เปลี่ยนแปลงข้อมูลล่าสุด';
COMMENT ON COLUMN public.tr_avg_rainfall_y30.last_updated_by IS 'ผู้เปลี่ยนแปลงข้อมูล';

