// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Permporn Kuibumrung <permporn@haii.or.th>

package provinces

import (
	 model_last_media "haii.or.th/api/thaiwater30/model/latest_media"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Radar struct {
	Province_id	string `json:"province_id"`
	P_CODE		string `json:"P_CODE"`
	Name_t   	string `json:"name_t"`
	Name_e 		string `json:"name_e"`
	Radar_id    []string `json:"radar_id"` // รหัสสถานีคุณภาพน้ำ เช่น ["pkt240","krb240","stp240""]
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/radar
// @Method			GET
// @Summary			ข้อมูลเรดาร์ รายจังหวัด 
// @Parameter		province_code	query string example:`10` รหัสจังหวัด ไม่ใส่ = ทุกจังหวัด ,เลือกได้หลายจังหวัด เช่น 10,51,62
// @Produces		json
// @Response		200		RadarProvincestOutputSwagger successful operation
type RadarProvincestOutputSwagger struct {
	Result string                            `json:"result"` //example:`OK`
	Data   []model_last_media.Struct_Radar   `json:"data"`
}

func (srv *HttpService) getRadarProvinces(ctx service.RequestContext) error {

	p := &model_last_media.RadarProvincesInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	
	// read radar.json config
	radar := "/home/cim/go_local/src/haii.or.th/api/thaiwater30/service/provinces/radar.json" // server
	//radar :=  "./src/haii.or.th/api/thaiwater30/service/provinces/radar.json"  // local
	raw, err := ioutil.ReadFile(radar)
	if err != nil {
		fmt.Println("read station.json err : %s", err)
		return nil 
	}
	var radarArray []Radar
	if err = json.Unmarshal(raw, &radarArray); err != nil {
			fmt.Println("Unmarshal err : %s", err)
			fmt.Println("stations all: %v ", len(radarArray))
			return nil 
	}
	var code []string
	for _, radar := range radarArray {
		province_id := fmt.Sprint(radar.Province_id)
		if(province_id == p.Province_Code){
			code = radar.Radar_id
		}
	}
	rs, err := model_last_media.GetRadarByProvince(code)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}