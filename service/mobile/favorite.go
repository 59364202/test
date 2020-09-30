package mobile

import (
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/rainfall24hr"
	"haii.or.th/api/thaiwater30/model/rainforecast"
	"haii.or.th/api/thaiwater30/model/spatial_province"
	"haii.or.th/api/thaiwater30/model/tele_waterlevel"
	"haii.or.th/api/thaiwater30/model/tmd_report"
)

// struct for apidocs
type favoriteProvince struct {
	Rain_forecast []*rainforecast.Struct_RainForecast3Day  `json:"rain_forecast"` // array ข้อมูลคาดการณ์ฝน 3 วัน
	Rainfall      *rainfall24hr.Struct_RainfallMaxMin      `json:"rainfall"`      // ข้อมูลฝน ล่าสุด ในจังหวัด
	Water_level   *tele_waterlevel.Struct_WaterlevelMinMax `json:"water_level"`   // ข้อมูลระดับน้ำ ล่าสุด ในจังหวัด
	Temperature   *tmd_report.Struct_Temperature           `json:"temperature"`   // ข้อมูลอุณหภูมิ ล่าสุด ในจังหวัด
}

//struct for service
type favoriteProvinceItf struct {
	Rain_forecast interface{} `json:"rain_forecast"` // array ข้อมูลคาดการณ์ฝน 3 วัน
	Rainfall      interface{} `json:"rainfall"`      // ข้อมูลฝน ล่าสุด ในจังหวัด
	Water_level   interface{} `json:"water_level"`   // ข้อมูลระดับน้ำ ล่าสุด ในจังหวัด
	Temperature   interface{} `json:"temperature"`   // ข้อมูลอุณหภูมิ ล่าสุด ในจังหวัด
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/favorite/{prov_id}
// @Summary			ข้อมูลสถานที่โปรด
// @Description		ข้อมูลสถานที่โปรด
// @Method			GET
// @Parameter		prov_id	path string example:`1จ` รหัสจังหวัด
// @Produces		json
// @Response		200	favoriteProvince successful operation
func (srv *HttpService) handlerGetFavoriteProvince(ctx service.RequestContext) error {
	prov_id := ctx.GetServiceParams("prov_id")
	rs := &favoriteProvinceItf{}

	rs.LoadData(prov_id)

	ctx.ReplyJSON(rs)
	return nil
}

// struct for apidocs
type favoriteLatLong struct {
	Flag     bool        `json:"flag"`     // ส่งค่าเป็น true, false หมายถึง มี/ไม่มี ข้อมูลสถานที่โปรด
	Location interface{} `json:"location"` // array ข้อมูลสถานที่
	favoriteProvince
}

//struct for service
type favoriteLatLongItf struct {
	Flag     bool        `json:"flag"`     // ส่งค่าเป็น true, false หมายถึง มี/ไม่มี ข้อมูลสถานที่โปรด
	Location interface{} `json:"location"` // array ข้อมูลสถานที่
	favoriteProvinceItf
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/favorite/latlong/{lat}/{long}
// @Summary			ข้อมูลสถานที่โปรด จากพิกัดตำแหน่งที่อยู่ปัจจุบัน
// @Description		ข้อมูลสถานที่โปรด จากพิกัดตำแหน่งที่อยู่ปัจจุบัน
// @Method			GET
// @Parameter		lat	path string example:`13.15437605541853` พิกัดละติจูด
// @Parameter		long	path string example:`101.46972656250001` พิกัดลองจิจูด
// @Produces		json
// @Response		200	favoriteLatLong successful operation
func (srv *HttpService) handlerGetFavoriteLatLong(ctx service.RequestContext) error {
	lat := ctx.GetServiceParams("lat")
	long := ctx.GetServiceParams("long")

	rs := &favoriteLatLongItf{}

	prov, err := spatial_province.GerProv(lat, long)
	if err != nil || prov == nil {
		rs.Flag = false
	} else {
		rs.Flag = true
		rs.Location = prov
		prov_id := datatype.MakeString(prov.Province_id)

		rs.LoadData(prov_id)
	}
	ctx.ReplyJSON(rs)
	return nil
}

//	load data from another model to favoriteProvinceItf
func (s *favoriteProvinceItf) LoadData(prov_id string) {
	_rain_forecast, err_rain_forecast := rainforecast.GetRainForecast3Day(prov_id)
	if err_rain_forecast == nil {
		s.Rain_forecast = _rain_forecast
	} else {
		s.Rain_forecast = err_rain_forecast.Error()
	}

	_water_level, err_water_level := tele_waterlevel.GetWaterLevelMinMax(prov_id)
	if err_water_level == nil {
		s.Water_level = _water_level
	} else {
		s.Water_level = err_rain_forecast.Error()
	}

	_rainfall24h, err_rainfall24h := rainfall24hr.GetRainfallMaxMin(prov_id)
	if err_rainfall24h == nil {
		s.Rainfall = _rainfall24h
	} else {
		s.Rainfall = err_rainfall24h.Error()
	}

	_tmd, _ := tmd_report.GetTemperature(prov_id)
	//	if err_tmd == nil {
	s.Temperature = _tmd
	//	} else {
	//		s.Temperature = err_tmd.Error()
	//	}
}
