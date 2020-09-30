package analyst

import (
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//"haii.or.th/api/util/log"
)

const (
	DataServiceName = "thaiwater30/analyst"
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

	switch service_id {
	//======= Dam =======//
	case "dam_load":
		return srv.onLoadDamLastest(ctx)
	case "dam":
		return srv.getDam(ctx)
	case "dam_graph":
		return srv.getDamGraph(ctx)
	case "dam_daily_graph":
		return srv.getDamGraphDaily(ctx)
	case "dam_yearly_graph":
		return srv.getDamGraphYearly(ctx)
	case "dam_medium_graph":
		return srv.getDamMediumGraph(ctx)
	case "dam_monitoring":
		return srv.getDam4Main(ctx)
	case "dam_daily_sum_region_rid":
		return srv.getDamSumByRegionRid(ctx)	
	case "dam_daily_compare_sum_region_rid":
		return srv.getDamCompareSumByRegionRid(ctx)	
	case "cctv":
		return srv.getUrlCCTV(ctx)
	case "water_situation":
		return srv.getWaterSituation(ctx)
	case "water_map":
		return srv.getWaterMap(ctx)
	//======= Waterquality =======//
	case "waterquality_load":
		return srv.onLoadWaterQuality(ctx)
	case "waterquality_compare_station_graph":
		return srv.getWaterQualityGraphCompareAnalyst(ctx)
	case "waterquality_compare_param_graph":
		return srv.getWaterQualityGraphParamsAnalyst(ctx)
	case "waterquality_compare_waterlevel_graph":
		return srv.getWaterQualityWaterlevelAnalyst(ctx)
	case "waterquality_compare_datetime_graph":
		return srv.getWaterQualityDatetimeStationAnalyst(ctx)
	case "waterquality_monitoring":
		return srv.getWaterQualitySalinityAnalyst(ctx)
	//======= sea_waterlevel =======//
	case "sea_waterlevel_load":
		return srv.getSeaWaterlevelAnalyst(ctx)
	case "sea_waterlevel_graph":
		return srv.getSeaWaterlevelRealAnalyst(ctx)
	case "sea_waterlevel_forecast_graph":
		return srv.getSeaWaterlevelForecastAnalyst(ctx)

	//======= Storm History =======//
	case "storm_history":
		return srv.getStormHistory(ctx)
	
	//======= Image =======//
	case "rain_forecast_img":
		return srv.getPreRainImg(ctx)
	case "rain_forecast_animation_img":
		return srv.getPreRainAnimation(ctx)
	case "radar_img":
		return srv.getRadarImg(ctx)
	case "radar_history_img":
		return srv.getRadarHistoryImg(ctx)
	case "rain_forecast_history_img":
		return srv.getPrecipitationRainHistory(ctx)
		
	//------ wind10m
	case "wind10m_forecast_history_img":
		return srv.getWind10mHistory(ctx)
	case "wind10m_forecast_animation_img":
		return srv.getPreWindAnimation(ctx)

	//------ rainaccumulat (USNRL)
	case "rainaccumulat_img":
		return srv.rainAccumulatHistory(ctx)
		
	case "upper_wind_img":
		return srv.getUpperWindLatest(ctx)
	case "upper_wind_history_img":
		return srv.getUpperWindHistory(ctx)
	case "vertical_wind_history_img":
		return srv.getVerticalWindHistory(ctx)
	case "wave_forecast_history_img":
		return srv.getWaveHistory(ctx)
	case "report_history":
		return srv.getReportHistory(ctx)
	//======= end Image  =======//
	case "swat_img":
		return srv.getSwatLatest(ctx)
	case "swat_history_img":
		return srv.getSwatHistory(ctx)
	case "storm_scale":
		ctx.ReplyJSON(result.ResultJson1(setting.GetSystemSettingJSON("Frontend.public.storm_setting")))
	default:
		return rest.NewError(404, "Unknown service id : " + service_id, nil)
	}

	return nil
}
