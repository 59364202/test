// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Permporn Kuibumrung <permporn@haii.or.th>

package provinces

import (
	model_last_media "haii.or.th/api/thaiwater30/model/latest_media"
	modal_rainforecast "haii.or.th/api/thaiwater30/model/rainforecast_7day_province"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"strings"
	"fmt"
)

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/wrfrom_rainforcase7d
// @Method			GET
// @Summary			ภาพคาดการณ์ฝน 7 วัน รายจังหวัด (WRF-ROM 24hrs) 
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด 
// @Produces		json
// @Response		200		RainforcaseProvincestOutputSwagger successful operation
type RainforcaseProvincestOutputSwagger struct {
	Result string                            `json:"result"` //example:`OK`
	Data   []model_last_media.Struct_Media_Rainforcase   `json:"data"`
}

func (srv *HttpService) getRainforcaseProvinces(ctx service.RequestContext) error {
	
	var objResult model_last_media.Struct_Media_Rainforcase7day
	var rs []interface{}

	p := &model_last_media.RainforcaseProvincesInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	
	media , err  :=  model_last_media.GetLatestMediaRainforcaseProvince(9, 181)
	data, _ := modal_rainforecast.GetRainForecast3Day(p.Province_Code)
	var o = 0;
	for _ , m := range media {
		bb := strings.Split(m.Filename, ".")
		if len(p.Province_Code) > 0 {
			if bb[0] == p.Province_Code {
			 objResult.FilePath = m.FilePath
			 objResult.Filename = m.Filename
			 objResult.Path = m.Path
			 objResult.Datetime = m.Datetime
			 objResult.ProvinceId = bb[0]
			 for _ , d := range data {
				var n int
				fmt.Sscan(d.Forecast_day, &n);
				if o == (n-1) { 
				 objResult.Forecast_day = d.Forecast_day
				 objResult.Forecast_datetime = d.Forecast_datetime
				 objResult.Rainfall = d.Rainfall
				 objResult.Rainfall_text = d.Rainfall_text
				 objResult.Province_name = d.Province_name
				 rs = append(rs,objResult)
				}	
			 }	 
			 o++;	
			}
		}else{
			objResult.FilePath = m.FilePath	
			objResult.Filename = m.Filename
			objResult.Path = m.Path
			objResult.Datetime = m.Datetime
			objResult.ProvinceId = bb[0]
			rs = append(rs,objResult)	
		}
	}
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}