// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Permporn Kuibumrung <permporn@haii.or.th>
package provinces

import (
	model_setting "haii.or.th/api/server/model/setting"
	model_canal_station "haii.or.th/api/thaiwater30/model/canal_station"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"encoding/json"
)

type Struct_OnLoadWaterQuality struct {
	Data      *result.Result `json:"data"`
	Scale     *result.Result `json:"scale,omitempty"`
}

type Struct_station struct {
	Tele  []*model_tele_station.Struct_Station       `json:"tele_waterlevel"`
	Canal []*model_canal_station.Struct_CanalStation `json:"canal_waterlevel"`
}

type Struct_onLoadWaterQuality struct {
	Data      *Struct_onLoadWaterQuality_Data      `json:"data"`                 // คุณภาพน้ำในตาราง
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

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/waterquality
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Parameter		data_limit	query int example:`20` จำนวน records ที่ต้องการ
// @Summary			คุณภาพน้ำ รายจังหวัด 
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadWaterQuality successful operation
func (srv *HttpService) onLoadWaterQuality(ctx service.RequestContext) error {

	objResult := &Struct_OnLoadWaterQuality{}

	//=== Scale ===//
	rsScale := model_setting.GetSystemSettingJSON("Frontend.public.waterquality_setting")
	objResult.Scale = result.Result1(&rsScale)

	//=== Data ===//
	param_waterquality := &model_waterquality.Param_WaterQualityCache{}
	if err := ctx.GetRequestParams(param_waterquality); err != nil {
		return err
	}
	ctx.LogRequestParams(param_waterquality)
	
	resultData, err := model_waterquality.GetWaterQualityThailandDataCache(param_waterquality)
	if err != nil {
		objResult.Data = result.Result0(err)
	} else {
		objResult.Data = result.Result1(resultData)
	}

	ctx.ReplyJSON(objResult)
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/waterquality_station_graph
// @Method			GET
// @Summary			คุณภาพน้ำรายสถานีสำหรับกราฟ 
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