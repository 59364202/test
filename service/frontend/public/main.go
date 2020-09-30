package public

import (
	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/public"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	sub_service_id := ctx.GetServiceParams("id")

	switch service_id {
	case "thaiwater_main":
		return srv.getThaiwaterMain(ctx)
	case "thaiwater":
		switch sub_service_id {
		case "weather":
			return srv.getThaiwaterWeather(ctx)
		case "temperature":
			return srv.getThaiwaterTemperature(ctx)
		default:
			return rest.NewError(404, "Unknown service id", nil)
		}
	case "temperature_graph":
		return srv.getTemperatureGraphByStationAndDate(ctx)
	case "weather_graph":
		return srv.getWeatherGraphByStationAndDate(ctx)
	//======= thaiwater30 =======//
	case "thailand":
		return srv.getThailand(ctx)
	case "thailand_main":
		return srv.getThailandMain(ctx)
	case "thailand_main_rain":
		return srv.getThailandRain(ctx)
	case "thailand_main_waterlevel":
		return srv.getThailandWaterlevel(ctx)
	case "thailand_main_dam":
		return srv.getThailandDam(ctx)
	case "weather_img":
		return srv.getWeatherImageLatest(ctx)
	case "rain7day_forecast":
		return srv.getRain7dayForecast(ctx)
	case "weather_history_img":
		return srv.getWeatherImage(ctx)
	case "wave_forecast_img":
		return srv.getPreWaveImg(ctx)
	case "wave_forecast_animation_img":
		return srv.getPreWaveAnimation(ctx)
	case "water_balance_img":
		return srv.getWaterBalanceImg(ctx)
	case "water_balance_forcast_weekly_img":
		return srv.getPreWaterBalanceWeeklyImg(ctx)
	case "water_balance_forcast_monthly_img":
		return srv.getPreWaterBalanceMonthlyImg(ctx)
	case "storm_data":
		return srv.getStormData(ctx)
	//======= Rain =======//
	case "rain_load":
		return srv.getRainOnLoad(ctx)
		//		กราฟฝนรายชม. ย้อนหลัง ถ้าไม่ใส่วันที่ จะย้อนหลัง 3 วัน
	case "rain_hour_graph":
		return srv.getRainHourGraph(ctx)
	case "rain_24h":
		return srv.getRain24Hr(ctx)
	case "rain_24h_graph":
		return srv.getRain24HrGraph(ctx)
	case "rain_today":
		return srv.getRainToday(ctx)
	case "rain_today_graph":
		return srv.getRainTodayGraph(ctx)
	case "rain_yesterday":
		return srv.getRainDaily(ctx)
	case "rain_yesterday_graph":
		return srv.getRainDailyGraph(ctx)
	case "rain_monthly":
		return srv.getRainMonthly(ctx)
	case "rain_monthly_graph":
		return srv.getRainMonthlyGraph(ctx)
	case "rain_yearly":
		return srv.getRainYearly(ctx)
	case "rain_yearly_graph":
		return srv.getRainYearlyGraph(ctx)
	//======= Waterlevel =======//
	case "waterlevel_load":
		return srv.onLoadWaterlevel(ctx)
	case "waterlevel_graph": // กราฟนักวิเคราะห์ ใส่เงื่อนไข รหัสสถานี วันที่เริ่มต้น วันที่สิ้นสุด
		return srv.getWaterlevelGraphByStationAndDateAnalyst(ctx)
	case "waterlevel":
		return srv.getWaterlevelAnalyst(ctx)
	case "waterlevel_yearly_graph":
		return srv.getWaterlevelYearlyGraphAnalyst(ctx)
	case "watergate_load":
		return srv.onLoadWaterlevelInOut(ctx)
	case "watergate_graph":
		return srv.getWaterlevelInOutGraphAnalyst(ctx)
	case "flood_forecast_monitoring":
		return srv.getFloodForecastMonitoring(ctx)
	case "flood_forecast_data":
		return srv.getCpyForecastWaterlevel(ctx)
	case "wave_forecast":
		return srv.getSwanForecast(ctx)
	case "swan_station":
		return srv.getSwanStation(ctx)
		//======= advance =======//
	case "advance_rain_monthly_station_graph":
		return srv.getAdvRainMonthStationGraph(ctx)
	case "advance_rain_monthly_graph":
		return srv.getAdvRainMonthGraph(ctx)
	case "advance_rain_yearly_graph":
		return srv.getAdvRainYearGraph(ctx)
	case "advance_rain_monthly_region_graph":
		return srv.getAdvRainMonthlyAreaGraph(ctx)
	case "advance_rain_distribution":
		return srv.getAdvRainDiagram(ctx)
	case "advance_waterlevel_basin_graph":
		return srv.getAdvanceWaterlevelBasinGraph(ctx)
	case "advance_waterlevel_basin_24h_graph":
		return srv.getAdvanceWaterlevelBasin24hGraph(ctx)
	case "advance_graph_load":
		return srv.getAdvRainGraphOnload(ctx)
	case "advance_load":
		return srv.getAdvLoad(ctx)
	case "advance_rain_sum":
		return srv.getAdvRainSum(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
