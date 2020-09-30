// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package provinces

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/provinces"
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
	case "rain24":
		return srv.getRain24(ctx)
	case "rain3d":
		return srv.getRain3d(ctx)
	case "temperature_maxmin_thisweek":
		return srv.getMaxMinTemperatureThisWeek(ctx)
	case "waterlevel":
		return srv.getWaterlevel(ctx)
	//======= Waterquality =======//
	case "waterquality":
		return srv.onLoadWaterQuality(ctx)
	case "waterquality_station_graph":
		return srv.getWaterQualityGraphCompareAnalyst(ctx)
	//======= Radar =======//
	case "radar":
		return srv.getRadarProvinces(ctx)
	//====== Rain 7 day ====//
	case "rain7d":
		return srv.getRain7d(ctx)
	//====== Rain today ====//
	case "rain1d":
		return srv.getRain1d(ctx)
	//====== Dam uses water ====//
	case "dam_uses_water":
		return srv.getDamUsesWater(ctx)
	//====== Rain Month ======//
	case "rainfall_month":
		return srv.GetRainfallMonth(ctx)
	//====== RainForcase ======//
	case "wrfrom_rainforcase7d":
		return srv.getRainforcaseProvinces(ctx)
	//====== FloodForecast ======//
	case "floodforecast_waterlevel":
		return srv.getFloodForecastWaterlevel(ctx)
	//======= Dam =======//
	case "dam":
		return srv.getDamNear(ctx)
	case "dam_daily_graph":
		return srv.getDamGraphDaily(ctx)
	case "dam_yearly_graph":
		return srv.getDamGraphYearly(ctx)
	case "dam_medium_graph":
		return srv.getDamMediumGraph(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
