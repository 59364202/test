package cron

import (
	"haii.or.th/api/server/model/cronjob"
	"haii.or.th/api/server/model/setting"

	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
)

func RegisterCron() {
	cronjob.NewClusterFunc("thaiwater30.cron.UpdateCache.Latest.Media", model_latest_media.UpdateMediadDataCache)

	cronjob.NewClusterFunc("thaiwater30.cron.UpdateCache.Latest.Rain", model_rainfall24hr.UpdateRainfallThailandDataCache)
	cronjob.NewClusterFunc("thaiwater30.cron.UpdateCache.Latest.Waterlevel", model_tele_waterlevel.UpdateWaterLevelThailandDataCache)
	cronjob.NewClusterFunc("thaiwater30.cron.UpdateCache.Latest.Waterquality", model_waterquality.UpdateWaterQualityThailandDataCache)

	setting.SetSystemDefault("thaiwater30.service.metadata.MetadataDateRange", "0 3 10 * *")
	cronjob.NewClusterFunc("thaiwater30.cron.MetadataDateRange", model_metadata.ReCacheDataDateRange) // หา min, max date ของแต่ละ metadata

	setting.SetSystemDefault("thaiwater30.service.metadata.MetadataDateRange_Max", "30 */5 * * *")
	cronjob.NewClusterFunc("thaiwater30.cron.MetadataDateRange.Max", model_metadata.ReCacheDataDateRange_Max) // หา max date ของแต่ละ metadata
}
