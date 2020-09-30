package iframe

import (
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"time"

	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"

	uSetting "haii.or.th/api/thaiwater30/util/setting"
)

type waterQualityIframeStruct struct {
	Province *result.Result     `json:"province"` // จังหวัด
	Datatype *result.ResultJson `json:"datatype"` // ประเภทข้อมูล
	Setting  *result.ResultJson `json:"setting"`  // ตั้งค่าการแสดงผล
}
type Struct_Iframe_Waterquality struct {
	Province *Struct_Iframe_Province              `json:"province"` // จังหวัด
	Datatype *Struct_Iframe_Waterquality_Datatype `json:"datatype"` // ประเภทข้อมูล
	Setting  *Struct_Iframe_Waterquality_Setting  `json:"setting"`  // เกณฑ์
}
type Struct_Iframe_Waterquality_Datatype struct {
	Result string                                 `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_WaterqulityDataType `json:"data"`   // example:`[{"id":"1","value":"","text":"ความเค็ม(g/L)","name":{"th":"ความเค็ม(g/L)","en":"salinity(g/L)"}},{"id":"2","value":"ph","text":"กรด-ด่าง(pH)","name":{"th":"กรด-ด่าง(pH)","en":"pH"}}]` ประเภทข้อมูล
}
type Struct_Iframe_Waterquality_Setting struct {
	Result string                                 `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_WaterqualitySetting `json:"data"`   //  เกณฑ์
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/waterquality
// @Summary			รายชื่อจังหวัด, ประเภทข้อมูลคุณภาพน้ำ, การแสดงผล
// @Description		เริ่มต้นหน้า iframe waterquality
// @				* รายชื่อจังหวัด
// @				* ประเภทข้อมูลคุณภาพน้ำ
// @				* การแสดงผล
// @Method			GET
// @Produces		json
// @Response		200	Struct_Iframe_Waterquality successful operation
func (srv *HttpService) getWaterQuality(ctx service.RequestContext) error {
	rs := &waterQualityIframeStruct{}
	rs_province, err := model_lt_geocode.GetAllProvince()
	if err != nil {
		rs.Province = result.Result0(err)
	} else {
		rs.Province = result.Result1(rs_province)
	}

	rs.Datatype = &result.ResultJson{Result: "OK", Data: setting.GetSystemSettingJson("Frontend.public.waterquality_data_type")}
	rs.Setting = &result.ResultJson{Result: "OK", Data: setting.GetSystemSettingJson("Frontend.public.waterquality_setting")}

	ctx.ReplyJSON(rs)
	return nil
}

type Struct_Iframe_Waterquality_Graph struct {
	Result string                                                       `json:"result"` // example:`OK`
	Data   []*model_waterquality.WaterQualityGraphCompareAnalystOutput2 `json:"data"`   // ข้อมูลกราฟ
}

type Param_getWaterQualityGraph struct {
	WaterQualityStation int64  `json:"waterquality_station_id"` // รหัสสถานีคุณภาพน้ำ เช่น [1,2,3,4]
	Param               string `json:"param"`                   // ชื่อ field ที่ต้องการ เช่น ph
	DatetimeStart       string `json:"start_datetime"`          // เวลาเริ่มต้น เช่น 2016-01-02 15:04
	DatetimeEnd         string `json:"end_datetime"`            // เวลาสิ้นสุด เช่น 2016-01-02 15:04
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/waterquality_graph
// @Summary			ข้อมูลกราฟคุณภาพน้ำ
// @Method			GET
// @Parameter		waterquality_station_id	query	int64	example:104 รหัสสถานี
// @Parameter		param	query	string	enum:[salinity,ph,turbid,conductivity,tds,chlorophyll,do,temp] example:salinity ประเภทข้อมูล
// @Produces		json
// @Response		200	Struct_Iframe_Waterquality_Graph successful operation
func (srv *HttpService) getWaterQualityGraph(ctx service.RequestContext) error {
	p := &Param_getWaterQualityGraph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	layout := "2006-01-02 15:00"

	pa := &model_waterquality.WaterQualityGraphCompareAnalystInput{}

	now := time.Now()
	pa.WaterQualityStation = []int64{p.WaterQualityStation}
	pa.Param = p.Param
	pa.DatetimeEnd = now.Format(layout)
	pa.DatetimeStart = now.AddDate(0, 0, -3).Format(layout)

	rs_, err := model_waterquality.GetWaterQualityGraphCompareAnalyst(pa)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs_))

	return nil
}
