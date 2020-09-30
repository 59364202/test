package mobile

import (
	//	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	"haii.or.th/api/server/model/setting"

	"strings"
	"unicode"
)

const (
	// List of API service in this package.
	ModuleName     = "mobile"
	ServiceName    = ModuleName + "/"
	ServiceVersion = service.APIVersion1

	// dam
	DamSummaryLatest = ServiceName + "dam_summary_latest"       // ข้อมูลน้ำในเขื่อนหลัก ล่าสุด
	DamList          = ServiceName + "dam_list"                 // ข้อมูลน้ำในเขื่อนหลัก ล่าสุด
	DamGraphHistory  = ServiceName + "dam_graph/history/dam_id" // ข้อมูลกราฟเขื่อน ย้อนหลัง 2 ปี
	DamGraphCureemt  = ServiceName + "dam_graph/current/dam_id" // ข้อมูลกราฟเขื่อน ปีปัจจุบัน

	// favorite
	FavoriteProvince = ServiceName + "favorite/province" // ข้อมูลสถานที่โปรด
	FavoriteLatLong  = ServiceName + "favorite/latlong"  // ข้อมูลสถานที่โปรด จากพิกัดตำแหน่งที่อยู่ปัจจุบัน

	// forecast
	RainForecast     = ServiceName + "rain_forecast"      // คาดการณ์ฝน 3 วัน ล่าสุด
	RainForecastData = ServiceName + "rain_forecast_data" // ข้อมูลคาดการณ์ฝน 3 วัน ล่าสุด
	WaveForecast     = ServiceName + "wave_forecast"      // คาดการณ์คลื่น 3 วัน ล่าสุด
	Rain7dayForecast = ServiceName + "rain7day_forecast"  // คาดการณ์ฝน 7 วัน
	Wave7dayForecast = ServiceName + "wave7day_forecast"  // คาดการณ์คลื่น 7 วัน
	StormForecast    = ServiceName + "storm_forecast"     // คาดการณ์พายุ ล่าสุด

	// infomation
	Province = ServiceName + "province" // ข้อมูลจังหวัด

	//media
	MediaLatest = ServiceName + "media/radar/tmd/latest" // ข้อมูลภาพเรดาร์จากกรมอุตุนิยมวิทยา

	// provinces
	WaterqualityLatest = ServiceName + "waterquality_latest" // waterquality_latest
	RainForecastProv   = ServiceName + "rain_forecast_prov"  // ข้อมูลคาดการณ์ฝน รายจังหวัด

	//Rainfall
	Rainfall24h_latest             = ServiceName + "rainfall24h_latest"            //ข้อมูลปริมาณฝนสะสม 24 ชม.ล่าสุด
	Rainfall_latest_list           = ServiceName + "rainfall_latest_list"          //ข้อมูลปริมาณฝนสะสม 24 ชม.ล่าสุด
	Rainfall_latest_list_provincce = ServiceName + "rainfall_latest_list_province" //ข้อมูลปริมาณฝน ล่าสุด จังหวัด
	Rainfall_lateest_station_graph = ServiceName + "rainfall_station_graph"

	//Storm
	Storm = ServiceName + "storm" //ข้อมูลพายุ ล่าสุด

	//Weather
	Weather_today = ServiceName + "weather_today"

	//Temperature
	Temperature = ServiceName + "temp_latest_prov"

	//warning
	BKK_rainfall24h = ServiceName + "warning/bkk/rainfall24h" // ข้อมูลเตือนภัยกรุงเทพมหานคร
	BKK_rainfall6h  = ServiceName + "warning/bkk/rainfall6h"  // ข้อมูลเตือนภัยกรุงเทพมหานคร

	//Waterlevel
	Wl_basin_latest  = ServiceName + "wl_basin_latest"     //ข้อมูลระดับน้ำ ล่าสุด
	Wl_latest_list   = ServiceName + "wl_latest_list"      //ข้อมูลระดับน้ำสูงสุด และต่ำสุด ล่าสุด 20 อันดับ
	Wl_latest_prov   = ServiceName + "wl_latest_list_prov" //ข้อมูลระดับน้ำ จังหวัด ล่าสุด 20 อันดับ
	Wl_station_graph = ServiceName + "wl_station_graph"    //ข้อมูลกราฟ ระดับน้ำ 7 วัน, 24 ชั่วโมง
)

type HttpService struct {
}

type mobileErr struct {
	Error string `json:"error"`
}

//	json error for mobile api
//	Parameters:
//		msg
//			error massage
//	Return:
//		mobileErr
func MobileErr(msg string) *mobileErr {
	if msg == "" {
		msg = "-"
	}
	return &mobileErr{msg}
}

//	ครอบ service.Dispatcher เพื่อทำฟังค์ชั่นเช็ค token ก่อน
type Dispatcher struct {
	dpt service.Dispatcher
}

func (d Dispatcher) Register(version int64, method int64, serviceName string, fn service.HandlerFunc) {
	d.dpt.Register(version, method, serviceName, func(ctx service.RequestContext) error {
		return d.checkToken(fn, ctx)
	})
}

//	https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
func (d Dispatcher) RemoveAllSpac(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

//	ฟังค์ชั่นเช็ค token ถ้าไม่มีสิทธิ์ จะขึ้น "Unauthorized Access."
func (d Dispatcher) checkToken(fn service.HandlerFunc, ctx service.RequestContext) error {
	token := ctx.GetServiceParams("token")
	allowToken := d.RemoveAllSpac(setting.GetSystemSetting("service.mobile.token"))
	if !strings.Contains(allowToken, `"token":"`+token+`"`) {
		//  เช็ค token แล้วไม่มีใน api.system_setting service.mobile.token
		//	reply json error (*mobileErr)
		ctx.ReplyJSON(MobileErr("Unauthorized Access."))
		return nil
	}

	return fn(ctx)
}

func RegisterService(_dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt := &Dispatcher{_dpt}
	// dam
	dpt.Register(ServiceVersion, service.MethodGET, DamSummaryLatest, srv.handlerGetDamSummaryLatest)
	dpt.Register(ServiceVersion, service.MethodGET, DamList, srv.handlerGetDamList)
	dpt.Register(ServiceVersion, service.MethodGET, DamGraphHistory, srv.handlerGetDamGraphHistory)
	dpt.Register(ServiceVersion, service.MethodGET, DamGraphCureemt, srv.handlerGetDamGraphCurrent)

	// favorite
	dpt.Register(ServiceVersion, service.MethodGET, FavoriteProvince, srv.handlerGetFavoriteProvince)
	dpt.Register(ServiceVersion, service.MethodGET, FavoriteLatLong, srv.handlerGetFavoriteLatLong)

	// forecast
	forecast := &Forecast{}
	dpt.Register(ServiceVersion, service.MethodGET, RainForecast, forecast.handlerGetRainForecast)
	dpt.Register(ServiceVersion, service.MethodGET, RainForecastData, forecast.handlerGetRainForecastData)
	dpt.Register(ServiceVersion, service.MethodGET, WaveForecast, forecast.handlerGetWaveForecast)
	dpt.Register(ServiceVersion, service.MethodGET, Rain7dayForecast, forecast.handlerGetRain7dayForecast)
	dpt.Register(ServiceVersion, service.MethodGET, Wave7dayForecast, forecast.handlerGetWave7dayForecast)
	dpt.Register(ServiceVersion, service.MethodGET, StormForecast, forecast.handlerGetStormForecast)

	// infomation
	dpt.Register(ServiceVersion, service.MethodGET, Province, srv.handlerGetProvince)

	// Media latest
	dpt.Register(ServiceVersion, service.MethodGET, MediaLatest, srv.handlerGetMediaLatest)

	//	provinces
	//	ไม่ได้อยู่บน mobile ไม่ต้องทำ
	provinces := &Provinces{}
	dpt.Register(ServiceVersion, service.MethodGET, WaterqualityLatest, provinces.handlerGetWaterqualityLatest)

	// Warning
	// rainfal24h
	dpt.Register(ServiceVersion, service.MethodGET, BKK_rainfall24h, srv.handlerGetBKK_Rainfall24h)

	// rainfal6h
	dpt.Register(ServiceVersion, service.MethodGET, BKK_rainfall6h, srv.handlerGet_bkk_Rainfall6h)

	//Rainfall
	dpt.Register(ServiceVersion, service.MethodGET, Rainfall24h_latest, srv.handlerGet_Rainfall24h_latest)
	dpt.Register(ServiceVersion, service.MethodGET, Rainfall_latest_list, srv.handlerGet_Rainfall_latest_list)
	dpt.Register(ServiceVersion, service.MethodGET, Rainfall_latest_list_provincce, srv.handlerGet_rainfall_province)
	dpt.Register(ServiceVersion, service.MethodGET, Rainfall_lateest_station_graph, srv.handlerGet_rainfall_lateest_station_graph)

	//Storm
	dpt.Register(ServiceVersion, service.MethodGET, Storm, srv.handlerGet_storm)

	//weather
	dpt.Register(ServiceVersion, service.MethodGET, Weather_today, srv.handlerGet_weather_today)

	//temperature
	dpt.Register(ServiceVersion, service.MethodGET, Temperature, srv.handlerGet_temperature)

	//Waterlevel
	dpt.Register(ServiceVersion, service.MethodGET, Wl_basin_latest, srv.handlerGet_wl_basin_data)
	dpt.Register(ServiceVersion, service.MethodGET, Wl_latest_list, srv.handlerGet_wl_latest_list)
	dpt.Register(ServiceVersion, service.MethodGET, Wl_latest_prov, srv.handlerGet_wl_latest_prov)
	dpt.Register(ServiceVersion, service.MethodGET, Wl_station_graph, srv.handlerGet_wl_station_graph)
}
