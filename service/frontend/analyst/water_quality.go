package analyst

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_canal_station "haii.or.th/api/thaiwater30/model/canal_station"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"
	"haii.or.th/api/thaiwater30/util/result"
	//"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//"strings"
	"encoding/json"
)

type Struct_OnLoadWaterQuality struct {
	Data      *result.Result `json:"data"`
	WQParam   *result.Result `json:"waterquality_param"`
	WQStation *result.Result `json:"waterquality_station"`
	WLStation *result.Result `json:"waterlevel_station"`
	Scale     *result.Result `json:"scale,omitempty"`
}

type Struct_station struct {
	Tele  []*model_tele_station.Struct_Station       `json:"tele_waterlevel"`
	Canal []*model_canal_station.Struct_CanalStation `json:"canal_waterlevel"`
}

type Struct_onLoadWaterQuality struct {
	Data      *Struct_onLoadWaterQuality_Data      `json:"data"`                 // คุณภาพน้ำในตาราง
	WQParam   *Struct_onLoadWaterQuality_WQParam   `json:"waterquality_param"`   // ชนิดข้อมูล
	WQStation *Struct_onLoadWaterQuality_WQStation `json:"waterquality_station"` // สถานีคุณภาพน้ำ
	WLStation *Struct_onLoadWaterQuality_WLStation `json:"waterlevel_station"`   // สถานีระดับน้ำ
	Scale     *Struct_onLoadWaterQuality_Scale     `json:"scale,omitempty"`      // เกณฑ์
}
type Struct_onLoadWaterQuality_Data struct {
	Result string                                    `json:"result"` // example:`OK`
	Data   []*model_waterquality.Struct_WaterQuality `json:"data"`   // คุณภาพน้ำในตาราง
}
type Struct_onLoadWaterQuality_WQParam struct {
	Result string          `json:"result"` // example:`OK`
	Data   json.RawMessage `json:"data"`   // example:`[{"id":"1","value":"salinity","text":"ความเค็ม(g/L)","name":{"th":"ความเค็ม(g/L)","en":"salinity(g/L)"}}]` ชนิดข้อมูล
}
type Struct_onLoadWaterQuality_WQStation struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model_waterquality_station.Struct_Station `json:"data"`   // สถานีคุณภาพน้ำ
}
type Struct_onLoadWaterQuality_WLStation struct {
	Result string         `json:"result"` // example:`OK`
	Data   Struct_station `json:"data"`   // สถานีระดับน้ำ
}
type Struct_onLoadWaterQuality_Scale struct {
	Result string         `json:"result"` // example:`OK`
	Data   Struct_station `json:"data"`   // example:`{"scale":{"salinity":[{"operator":">","term":"2","color":"#EE141F","text":"> 2","trans":"salinity_2","inGraph":true}]}}` เกณฑ์
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_load
// @Summary			เริ่มต้นหน้าคุณภาพน้ำ
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadWaterQuality successful operation
func (srv *HttpService) onLoadWaterQuality(ctx service.RequestContext) error {

	objResult := &Struct_OnLoadWaterQuality{}

	//=== Waterquality Param ===//
	resultParam := model_setting.GetSystemSettingJson("Frontend.public.waterquality_data_type")
	objResult.WQParam = result.Result1(&resultParam)

	//=== Scale ===//
	rsScale := model_setting.GetSystemSettingJSON("Frontend.public.waterquality_setting")
	objResult.Scale = result.Result1(&rsScale)

	//=== Waterquality Station ===//
	resultWQStation, err := model_waterquality_station.Get_AllWaterQualityStaion()
	if err != nil {
		objResult.WQStation = result.Result0(err)
	} else {
		objResult.WQStation = result.Result1(resultWQStation)
	}

	//=== Waterlevel Station ===//
	//	resultWLStation, err := model_tele_station.GetWaterlevelCanalStation("")
	//	if err != nil {
	//		objResult.WLStation = result.Result0(err)
	//	} else {
	//		objResult.WLStation = result.Result1(resultWLStation)
	//	}

	resultWLSation := &Struct_station{}
	//=== Tele Station ===//
	rsTeleStation, errT := model_tele_station.GetTeleStation()
	if errT != nil {
		err = errT
	} else {
		resultWLSation.Tele = make([]*model_tele_station.Struct_Station, 0)
		resultWLSation.Tele = rsTeleStation
	}
	//=== Canal Station ===//
	rsCanalStation, errC := model_canal_station.GetCanalStation()
	if errC != nil {
		err = errC
	} else {
		resultWLSation.Canal = make([]*model_canal_station.Struct_CanalStation, 0)
		resultWLSation.Canal = rsCanalStation
	}

	if err != nil {
		objResult.WLStation = result.Result0(err)
	} else {
		objResult.WLStation = result.Result1(resultWLSation)
	}

	//=== Data ===//
	resultData, err := model_waterquality.GetWaterQualityThailandDataCache(&model_waterquality.Param_WaterQualityCache{})
	if err != nil {
		objResult.Data = result.Result0(err)
	} else {
		objResult.Data = result.Result1(resultData)
	}

	ctx.ReplyJSON(objResult)
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_compare_station_graph
// @Method			GET
// @Summary			เปรียบเทียบคุณภาพน้ำรายสถานีสำหรับกราฟ
// @Parameter		-	query model_waterquality.WaterQualityGraphCompareAnalystInput
// @Produces		json
// @Response		200		WaterQualityGraphCompareAnalystOutputSwagger successful operation
type WaterQualityGraphCompareAnalystOutputSwagger struct {
	Result string                                                      `json:"result"` //example:`OK`
	Data   []model_waterquality.WaterQualityGraphCompareAnalystOutput2 `json:"data"`
}

func (srv *HttpService) getWaterQualityGraphCompareAnalyst(ctx service.RequestContext) error {

	p := &model_waterquality.WaterQualityGraphCompareAnalystInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_waterquality.GetWaterQualityGraphCompareAnalyst(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_compare_param_graph
// @Method			GET
// @Summary			เปรียบเทียบคุณภาพน้ำ พารามิเตอร์ สำหรับกราฟ
// @Parameter		-	query model_waterquality.WaterQualityGraphParamsAnalystInput
// @Produces		json
// @Response		200		[]model_waterquality.WaterQualityGraphParamsAnalystOutput2 successful operation
type WaterQualityGraphParamsAnalystOutputSwagger struct {
	Result string                                                     `json:"result"` //example:`OK`
	Data   []model_waterquality.WaterQualityGraphParamsAnalystOutput2 `json:"data"`
}

func (srv *HttpService) getWaterQualityGraphParamsAnalyst(ctx service.RequestContext) error {

	p := &model_waterquality.WaterQualityGraphParamsAnalystInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_waterquality.GetWaterQualiyGraphParamsAnalyst(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_compare_waterlevel_graph
// @Method			GET
// @Summary			เปรียบเทียบคุณภาพน้ำกับระดับน้ำสำหรับกราฟ
// @Parameter		-	query model_waterquality.WaterQualityGraphWaterlevelAnalystInput
// @Produces		json
// @Response		200		WaterQualityWaterlevelOutputSwagger successful operation

type WaterQualityWaterlevelOutputSwagger struct {
	Result string                                            `json:"result"` //example:`OK`
	Data   []model_waterquality.WaterQualityWaterlevelOutput `json:"data"`
}

func (srv *HttpService) getWaterQualityWaterlevelAnalyst(ctx service.RequestContext) error {

	p := &model_waterquality.WaterQualityGraphWaterlevelAnalystInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_waterquality.GetWaterQualiyGraphWaterlevel(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_compare_datetime_graph
// @Method			GET
// @Summary			เปรียบเทียบคุณภาพน้ำรายวันที่ สำหรับกราฟ
// @Parameter		-	query model_waterquality.WaterQualityGraphCompareDatetimeInput
// @Produces		json
// @Response		200		WaterQualityCompareDatetimeOutputSwagger successful operation

type WaterQualityCompareDatetimeOutputSwagger struct {
	Result string                                                 `json:"result"` //example:`OK`
	Data   []model_waterquality.WaterQualityCompareDatetimeOutput `json:"data"`
}

func (srv *HttpService) getWaterQualityDatetimeStationAnalyst(ctx service.RequestContext) error {

	p := &model_waterquality.WaterQualityGraphCompareDatetimeInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_waterquality.GetWaterQualityGraphCompareDatetimeAndStation(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/waterquality_monitoring
// @Summary			สถานีดูค่าความเค็ม
// @Method			GET
// @Produces		json
// @Response		200		WaterQualityCompareDatetimeOutputSwagger2 successful operation
type WaterQualityCompareDatetimeOutputSwagger2 struct {
	Result string                                            `json:"result"` //example:`OK`
	Data   []model_waterquality.MonitoringWaterqualityOutput `json:"data"`
}

func (srv *HttpService) getWaterQualitySalinityAnalyst(ctx service.RequestContext) error {

	rs, err := model_waterquality.GetWaterQualitySalinityAnalyst()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
