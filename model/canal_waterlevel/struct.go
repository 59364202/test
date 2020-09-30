// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_waterlevel is a model for public.canal_waterlevel table. This table store canal_waterlevel.
package canal_waterlevel

import ()

type Param_CanalWaterlevel struct {
	Id            string `json:"id"`            //example:`1` ลำดับของรหัส
	Station_id    string `json:"station_id"`    // example:`132` รหัสของสถานีเป็นลำดับ
	Start_date    string `json:"start_date"`    // example:`2006-01-02` เวลาเริ่มต้น
	End_date      string `json:"end_date"`      // example:`2006-01-02` เวลาสิ้นสุด
	Subbasin_id   string `json:"subbasin_id"`   // example:`1` รหัส subbasin
	Agency_id     string `json:"agency_id"`     // example:`4` รหัส agency
	Province_code string `json:"province_code"` // example:`07` รหัส province
}

type GetCanalWaterlevelYearlyGraphInput struct {
	StationID   string `json:"station_id"`   // example:`1234` รหัสสถานี
	StationType string `json:"station_type"` // example:`W` ประเภทของสถานี
	Year        []int  `json:"year"`         // example:`2015,2016,2017` array ของปี
}
