package public

import (
	"haii.or.th/api/server/model/setting"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	result "haii.or.th/api/thaiwater30/util/result"
	//	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	//	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_drought_area "haii.or.th/api/thaiwater30/model/drought_area"
	model_flood_situation "haii.or.th/api/thaiwater30/model/flood_situation"
	model_geocode "haii.or.th/api/thaiwater30/model/geocode"
	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media_animation "haii.or.th/api/thaiwater30/model/media_animation"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_rainforecast "haii.or.th/api/thaiwater30/model/rainforecast"
	model_storm "haii.or.th/api/thaiwater30/model/storm"
	model_swan "haii.or.th/api/thaiwater30/model/swan"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_temperature "haii.or.th/api/thaiwater30/model/temperature"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
)

const (
	thailandMainCacheName     = "public.thailand.main"
	thailandPageCacheName     = "public.thailand.index"
	rainCacheName             = "public.main.rain_"
	waterLevelCacheName       = "public.main.waterlevel_"
	thaiwaterMainCacheName    = "public.thaiwater.main"
	thaiwaterWeatherCacheName = "public.thaiwater.weather"
)

func buildResult(data interface{}, err error) *result.Result {
	if err != nil {
		return result.Result0(err.Error())
	}
	return result.Result1(data)
}

// get storm image
func getStromData() (interface{}, error) {
	//UCL storm image
	storm1, err := model_latest_media.GetLatestMedia(41, 62)
	if err != nil {
		return nil, err
	}
	//  Kochi cloud himawari-8
	storm2, err := model_latest_media.GetLatestMedia(50, 141)
	if err != nil {
		return nil, err
	}
	if !model_latest_media.ValidStorm(storm2) {
		storm2 = nil
	}
	//Digital Typhoon cloud himawari-8
	storm3, err := model_latest_media.GetLatestMedia(52, 141)
	if err != nil {
		return nil, err
	}
	// if !model_latest_media.ValidStorm(storm3) {
	// 	storm3 = nil
	// }

	// Naval Research Laboratory cloud himawari-8
	storm4, err := model_latest_media.GetLatestMedia(51, 141)
	if err != nil {
		return nil, err
	}

	return &Struct_Storm_Data{C: storm1, K: storm2, T: storm3, U: storm4}, nil
}

// main
func getThailandMainCacheBuildData() (interface{}, error) {
	rs := &ThailandStruct{}

	//province
	rs.Province = buildResult(model_geocode.GetProvinceAll())

	// rain
	rs.Rain = new(ThailandModuleStruct)
	rs.Rain.Data = buildResult(model_rainfall24hr.GetRainfallThailandDataCache(&model_rainfall24hr.Param_Rainfall24{}))

	// radar
	rs.Radar = new(ThailandModuleStruct)
	rs.Radar.Data = buildResult(model_latest_media.GetRadar())

	rs.Dam = new(ThailandModuleStruct)

	p := &model_dam_daily.Struct_DamDailyLastest_InputParam{Agency_id: "12"}
	rs.Dam.Data = buildResult(model_dam_daily.GetDamDailyLastest(p))

	// water_level
	rs.WaterLevel = new(ThailandModuleStruct)
	rs.WaterLevel.Data = buildResult(
		model_tele_waterlevel.GetWaterLevelThailandDataCache(
			&model_tele_waterlevel.Waterlevel_InputParam{IsMain: true}))

	// water quality
	rs.WaterQuality = new(ThailandModuleStruct)
	rs.WaterQuality.Data = buildResult(model_waterquality.GetWaterQualityThailandDataCache(&model_waterquality.Param_WaterQualityCache{IsMain: true}))

	// strom
	rs.Storm = new(ThailandModuleStruct)
	rs.Storm.Data = buildResult(getStromData())

	// pre rain th
	rs.PreRain = new(ThailandModuleStruct)
	rs.PreRain.Data = buildResult(model_latest_media.GetPreRainTH())
	// pre rain asia
	rs.PreRainAsia = new(ThailandModuleStruct)
	rs.PreRainAsia.Data = buildResult(model_latest_media.GetPreRainAsia())
	// pre rain sea
	rs.PreRainSea = new(ThailandModuleStruct)
	rs.PreRainSea.Data = buildResult(model_latest_media.GetPreRainSea())
	// pre rain basin
	rs.PreRainBasin = new(ThailandModuleStruct)
	rs.PreRainBasin.Data = buildResult(model_latest_media.GetPreRainBasin())
	// pre rain animation
	rs.PreRainAnimation = buildResult(model_media_animation.GetPreRainAnimation())

	// wave
	rs.Wave = new(ThailandModuleStruct)
	rs.Wave.Data = buildResult(model_latest_media.GetPreWave())
	rs.WaveAnimation = buildResult(model_media_animation.GetPreWaveAnimation())

	// warning
	rs.Warning = new(ThailandModuleStruct)

	rs.Warning.TempData = buildResult(model_rainfall24hr.GetRainfallThailandDataCache(&model_rainfall24hr.Param_Rainfall24{IsHourly: true}))
	rs.Warning.TempData2 = buildResult(model_rainforecast.GetRainforecastWarning())
	rs.Warning.Drought = buildResult(model_drought_area.GetLatestDrought())
	rs.Warning.Flood = buildResult(model_flood_situation.GetLatestFloodSituation())

	//Fill setting
	rs.Rain.Setting = setting.GetSystemSettingJson("Frontend.public.rain_setting")
	rs.Dam.Setting = setting.GetSystemSettingJson("Frontend.public.dam_scale_color")
	rs.WaterLevel.Setting = setting.GetSystemSettingJson("Frontend.public.waterlevel_setting")
	rs.WaterQuality.Setting = setting.GetSystemSettingJSON("Frontend.public.waterquality_setting")
	rs.Storm.Setting = setting.GetSystemSettingJSON("Frontend.public.storm_setting")
	rs.Warning.Setting = setting.GetSystemSettingJSON("Frontend.public.warning_setting")
	return rs, nil
}

// แผนที่ประเทศไทย
func getThailandPageCacheBuildData() (interface{}, error) {
	rs := &ThailandStruct{}

	//province
	rs.Province = buildResult(model_geocode.GetProvinceAll())

	// rain
	rs.Rain = new(ThailandModuleStruct)
	rs.Rain.Data = buildResult(model_rainfall24hr.GetRainfallThailandDataCache(&model_rainfall24hr.Param_Rainfall24{IsHourly: true}))
	rs.Rain.Setting = setting.GetSystemSettingJson("Frontend.public.rain_setting")

	// dam
	rs.Dam = new(ThailandModuleStruct)

	pdam := &model_dam_daily.Struct_DamDailyLastest_InputParam{
		Agency_id: "12",
	}
	rs.Dam.Data = buildResult(model_dam_daily.GetDamDailyLastest(pdam))
	rs.Dam.Setting = setting.GetSystemSettingJson("Frontend.public.dam_scale_color")

	// water_level
	rs.WaterLevel = new(ThailandModuleStruct)
	rs.WaterLevel.Data = buildResult(model_tele_waterlevel.GetWaterLevelThailandDataCache(&model_tele_waterlevel.Waterlevel_InputParam{IsHourly: true}))
	rs.WaterLevel.Setting = setting.GetSystemSettingJson("Frontend.public.waterlevel_setting")

	// water_quality
	rs.WaterQuality = new(ThailandModuleStruct)
	rs.WaterQuality.Data = buildResult(model_waterquality.GetWaterQualityThailandDataCache(&model_waterquality.Param_WaterQualityCache{IsMain: true}))
	rs.WaterQuality.Setting = setting.GetSystemSettingJSON("Frontend.public.waterquality_setting")

	// pre_rain
	rs.PreRain = new(ThailandModuleStruct)
	rs.PreRain.Data = buildResult(model_rainforecast.GetRainforecastWarning())
	rs.PreRain.Setting = setting.GetSystemSettingJSON("Frontend.public.pre_rain_setting")

	// wave
	rs.Wave = new(ThailandModuleStruct)
	rs.Wave.Data = buildResult(model_swan.GetSwanCurrentDate())
	rs.Wave.Setting = setting.GetSystemSettingJSON("Frontend.public.wave_setting")

	// strom + list เส้นทางพายุ
	rs.Storm = new(ThailandModuleStruct)
	rs.Storm.Data = buildResult(model_storm.GetStormCurrentDate())
	rs.Storm.Setting = setting.GetSystemSettingJSON("Frontend.public.storm_setting")

	// warning
	rs.Warning = new(ThailandModuleStruct)
	rs.Warning.Setting = setting.GetSystemSettingJSON("Frontend.public.warning_setting")
	rs.Warning.Drought = buildResult(model_drought_area.GetLatestDrought())
	rs.Warning.Flood = buildResult(model_flood_situation.GetLatestFloodSituation())

	return rs, nil
}

// Provinces : Radar
func getProvincesCacheBuildData() (interface{}, error) {
	rs := &ThailandStruct{}

	// radar
	rs.Radar = new(ThailandModuleStruct)
	rs.Radar.Data = buildResult(model_latest_media.GetRadar())
	return rs, nil
}

// ================= thaiwater main
func getThaiwaterMainCacheBuildData() (interface{}, error) {
	rs := &ThaiwaterMainStruct{}

	// dam
	rs.Dam = new(ThailandModuleStruct)
	p := &model_dam_daily.Struct_DamDailyLastest_InputParam{Agency_id: "12"}
	rs.Dam.Data = buildResult(model_dam_daily.GetDamDailyLastest(p))
	rs.Dam.Summary = buildResult(model_dam_daily.GetDamDailyLastestSummary())
	rs.Dam.Summary4Dam = buildResult(model_dam_daily.GetDamDailySummary4Dam())

	// Temperature
	rs.Temperature = new(ThailandModuleStruct)
	temperatureParam := &model_temperature.Param_TemperatureProvinces{}
	temperatureParam.Data_Limit = 10
	rs.Temperature.Data = buildResult(model_temperature.GetTemperatureLatest(temperatureParam))
	rs.Temperature.Summary = buildResult(model_temperature.GetTemperatureMaxMinByRegion())

	// rain
	rs.Rain = new(ThailandModuleStruct)
	rs.Rain.Data = buildResult(model_rainfall24hr.GetRainfallThailandDataCache(&model_rainfall24hr.Param_Rainfall24{}))

	// strom image
	rs.Storm = new(ThailandModuleStruct)
	rs.Storm.Data = buildResult(getStromData())

	//Fill setting
	rs.Dam.Setting = setting.GetSystemSettingJson("Frontend.public.dam_scale_color")
	rs.Rain.Setting = setting.GetSystemSettingJson("Frontend.public.rain_setting")
	rs.Storm.Setting = setting.GetSystemSettingJSON("Frontend.public.storm_setting")

	return rs, nil
}
