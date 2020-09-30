// Edit by : Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media "haii.or.th/api/thaiwater30/model/media"
	model_weather_image "haii.or.th/api/thaiwater30/model/weather_image"
)

type AnimationInput struct {
	Agency    int64 `json:"agency_id"`     // example:`51` รหัสหน่วยงาน เช่น 51
	MediaType int64 `json:"media_type_id"` // example:`141` รหัสประเภทสื่อ เช่น 141
}

type HistoryDateInput struct {
	Agency    int64  `json:"agency_id"`     // example:`51` รหัสหน่วยงาน เช่น 51
	MediaType int64  `json:"media_type_id"` // example:`141` รหัสประเภทสื่อ เช่น 141
	Date      string `json:"date"`          // example:`2017-07-30` วันที่ เช่น 2017-07-30
}

type HistoryDateRangeInput struct {
	Agency    int64  `json:"agency_id"`     // example:`13` รหัสหน่วยงาน เช่น 13
	MediaType int64  `json:"media_type_id"` // example:`29` รหัสประเภทสื่อ เช่น 29
	StartDate string `json:"start_date"`    // example:`2013-05-04` วันเริ่มต้นของข้อมูลสื้อ เช่น 2013-05-04
	EndDate   string `json:"end_date"`      // example:`2013-05-10` วันสิ้นสุดของข้อมูลสื้อ เช่น 2013-05-10
}

type HistoryDateMonthInput struct {
	Agency    int64  `json:"agency_id"`     // example:`57` รหัสหน่วยงาน เช่น 57
	MediaType int64  `json:"media_type_id"` // example:`153` รหัสประเภทสื่อ เช่น 153
	Year      string `json:"year"`          // example:`2017` ปี เช่น 2017
	Month     string `json:"month"`         // example:`08` month เช่น 08
}

type HistoryDateYearInput struct {
	Agency    int64  `json:"agency_id"`     // example:`9` รหัสหน่วยงาน เช่น 9
	MediaType int64  `json:"media_type_id"` // example:`14` รหัสประเภทสื่อ เช่น 14
	Year      string `json:"year"`          // ปี เช่น 2016
}

func (srv *HttpService) getWeatherImageLatest(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("id")
	p := &model_weather_image.WeatherImageHistoryAllParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	//model_latest_media.GetLatestMedia(agency_id, media_id)

	var rs interface{}
	var err error
	arrayResult := make([]interface{}, 0)
	switch service_id {
	case "storm_history": // link to www.thaiwater.net
		rs, err = model_weather_image.GetStormHistory()
	case "india_ocean_ucl": // storm
		rs, err = model_weather_image.GetIndiaOceanUCLLatest()
	case "pacific_ocean_ucl": // storm
		rs, err = model_weather_image.GetPacificOceanUCLLatest()
	case "weather_map_tmd": //
		rs1, _ := model_weather_image.GetWeatherMapTMDLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetWeatherMapWind850TMDLatest()
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_weather_image.GetWeatherMapWind925TMDLatest()
		arrayResult = append(arrayResult, rs3)
		rs = arrayResult
	case "weather_map_hd": //
		rs, err = model_weather_image.GetWeatherMapHDLatest()
	case "cloud": //
		//  Kochi cloud himawari-8
		rs1, _ := model_latest_media.GetLatestMedia(50, 141)
		//		ตรวจสอบภาพถ้าไม่ update ใน 6 ชม. ถือว่าไม่มีภาพ
		if !model_latest_media.ValidStorm(rs1) {
			rs1 = nil
		}
		//Digital Typhoon cloud himawari-8
		rs2, _ := model_latest_media.GetLatestMedia(52, 141)
		// Naval Research Laboratory cloud himawari-8
		rs3, _ := model_latest_media.GetLatestMedia(51, 141)
		rs = &Struct_Cloud_Data{K: rs1, T: rs2, U: rs3}
	case "cloud_kochi": // himawari-8
		rs, err = model_weather_image.GetCloudKochiLatest()
	case "cloud_us_naval_research_lab": // himawari-8
		rs, err = model_weather_image.GetCloudUSNRLLatest()
	case "cloud_digital_typhoon": // himawari-8
		rs, err = model_weather_image.GetCloudDTLatest()
	case "satellite_image_gistda": //
		rs1, _ := model_weather_image.GetSatelliteComsAsiaLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetSatelliteComsSoutheastAsiaLatest()
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_weather_image.GetSatelliteCcomsThailandLatest()
		arrayResult = append(arrayResult, rs3)
		rs = arrayResult
	case "rain_image_us_naval_research_lab": // link to www.thaiwater.net trmmImages
		rs, err = model_weather_image.GetRainImageDailyUSNRLLatest()
	case "satellite_image_gsmaps":
		rs1, _ := model_weather_image.GetRainImageDailyGSmapsLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetRainImageDailyTRMMLatest()
		arrayResult = append(arrayResult, rs2)
		rs = arrayResult
	case "contour_image": //
		rs1, _ := model_weather_image.GetContourTemperatureLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetContourHumidityLatest()
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_weather_image.GetContourPressureLatest()
		arrayResult = append(arrayResult, rs3)
		rs = arrayResult
	case "modis_ndvi_usda": //terra aqua
		rs1, _ := model_weather_image.GetVegetationindexTerraUSDALatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetVegetationindexAquaUSDALatest()
		arrayResult = append(arrayResult, rs2)
		rs = arrayResult
	case "modis_soil_moisture_usda": //
		//	AFWA Surface Soil Moisuture
		rs1, _ := model_weather_image.GetSoilMoitsureUSDALatest()
		arrayResult = append(arrayResult, rs1)
		//	AFWA Sub Surface Soil Moisuture
		rs2, _ := model_weather_image.GetSoilMoitsureUSDAAFWASubSurfaceLatest()
		arrayResult = append(arrayResult, rs2)
		//	WMO Surface Soil Moisuture
		rs3, _ := model_weather_image.GetSoilMoitsureUSDAWMOSurfaceLatest()
		arrayResult = append(arrayResult, rs3)
		//	WMO Sub Surface Soil Moisuture
		rs4, _ := model_weather_image.GetSoilMoitsureUSDAWMOSubSurfaceLatest()
		arrayResult = append(arrayResult, rs4)
		rs = arrayResult
	case "modis_precipitation_usda": //
		rs1, _ := model_weather_image.GetPrecipitationUSDAWMODecadalPercentNormalPrecipitationLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetPrecipitationUSDAWMOPrecipitationLatest()
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_weather_image.GetPrecipitationUSDAAFWADecadalPercentNormalPrecipitationLatest()
		arrayResult = append(arrayResult, rs3)
		rs4, _ := model_weather_image.GetPrecipitationUSDAAFWAPrecipitationLatest()
		arrayResult = append(arrayResult, rs4)
		rs = arrayResult
	case "sst_ocean_weather": //
		rs4, _ := model_weather_image.GetWaveOceanWeatherThailandLatest()
		arrayResult = append(arrayResult, rs4)
		rs3, _ := model_weather_image.GetTemperatureOceanWeatherNpwLatest()
		arrayResult = append(arrayResult, rs3)
		rs2, _ := model_weather_image.GetTemperatureOceanWeatherIndLatest()
		arrayResult = append(arrayResult, rs2)
		rs1, _ := model_weather_image.GetTemperatureOceanWeatherGlobalLatest()
		arrayResult = append(arrayResult, rs1)
		rs = arrayResult
	case "wave_height_ocean_weather": //
		rs4, _ := model_weather_image.GetWaveOceanWeatherThailandLatest()
		arrayResult = append(arrayResult, rs4)
		rs3, _ := model_weather_image.GetWaveOceanWeatherNpwLatest()
		arrayResult = append(arrayResult, rs3)
		rs2, _ := model_weather_image.GetWaveOceanWeatherIndLatest()
		arrayResult = append(arrayResult, rs2)
		rs1, _ := model_weather_image.GetWaveOceanWeatherGlobalLatest()
		arrayResult = append(arrayResult, rs1)
		rs = arrayResult
	case "gssh_aviso": //
		rs, err = model_weather_image.GetMapHighSDAvisoJason1Latest()
	case "gsmap_hii": // แผนภาพฝนสะสมรายวัน  GSMaP โดย HII
		rs1, _ := model_weather_image.GetMapGsmap10kmHiiLatest()
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_weather_image.GetMapGsmapPersiann4kmHiiLatest()
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_weather_image.GetMapGsmap25kmHiiLatest()
		arrayResult = append(arrayResult, rs3)
		rs = arrayResult
	case "sst_2w_haii":
		rs, err = model_weather_image.GetMapSeaTemperatureHAIILatest()
	case "ssh_event_haii":
		rs, err = model_weather_image.GetMapSSHEventLatest()
	case "ssh_w_haii": //
		rs, err = model_weather_image.GetMapSeaWaterlevelWeekLatest()
	case "animation":
		rs, err = model_weather_image.GetWeatherAnimation(p.Agency, p.MediaType)
	case "image_generate":
		//	 generate latest thailand map image from data (rain24,waterlevel,dam)
		//		model_latest_media.GetLatestMedia(media_type_id,agency_id)
		//		rs1, _ := model_latest_media.GetLatestMedia(64, 176)
		//		arrayResult = append(arrayResult, rs1)
		//		rs2, _ := model_latest_media.GetLatestMedia(64, 177)
		//		arrayResult = append(arrayResult, rs2)
		//		rs3, _ := model_latest_media.GetLatestMedia(64, 178)
		//		arrayResult = append(arrayResult, rs3)

		//		model_media.GetMediaLatestByMediaTypeID(media_type_id,limit)
		//		2020-04-16 พบบางช่วง refresh แล้วข้อมูลเป็น null ทั้งหมด
		rs1, _ := model_media.GetMediaLatestByMediaTypeID(176, 1)
		arrayResult = append(arrayResult, rs1)
		rs2, _ := model_media.GetMediaLatestByMediaTypeID(177, 1)
		arrayResult = append(arrayResult, rs2)
		rs3, _ := model_media.GetMediaLatestByMediaTypeID(178, 1)
		arrayResult = append(arrayResult, rs3)
		rs = arrayResult
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getWeatherImage(ctx service.RequestContext) error {

	p := &model_weather_image.WeatherImageHistoryAllParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	service_id := ctx.GetServiceParams("id")
	var rs interface{}
	var err error
	switch service_id {
	case "date":
		rs, err = model_weather_image.GetWeatherHistoryDate(p.Agency, p.MediaType, p.Date)
	case "date_range":
		rs, err = model_weather_image.GetWeatherHistoryStartDateEndDate(p.Agency, p.MediaType, p.StartDate, p.EndDate)
	case "month":
		rs, err = model_weather_image.GetWeatherHistoryYearMonth(p.Agency, p.MediaType, p.Year, p.Month)
	case "year":
		rs, err = model_weather_image.GetWeatherHistoryYear(p.Agency, p.MediaType, p.Year)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type AWeatherImageDataSwagger struct {
	Result string                                        `json:"result"` //example:`OK`
	Data   []model_weather_image.WeatherImageDataSwagger `json:"data"`
}
type WeatherImageDataSwagger struct {
	Result string                                      `json:"result"` //example:`OK`
	Data   model_weather_image.WeatherImageDataSwagger `json:"data"`
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/storm_history
// @Method			GET
// @Summary			ภาพพายุ
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/india_ocean_ucl
// @Method			GET
// @Summary			แผนที่วิเคราะห์เส้นทางและความรุนแรงงของพายุ ในมหาสมุทรอินเดีย
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/pacific_ocean_ucl
// @Method			GET
// @Summary			แผนที่วิเคราะห์เส้นทางและความรุนแรงงของพายุ ในมหาสมุทรแปซิฟิก
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/weather_map_tmd
// @Method			GET
// @Summary			แผนที่อากาศ
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/weather_map_hd
// @Method			GET
// @Summary			แผนที่ลมฟ้าอากาศ
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/cloud
// @Method			GET
// @Summary			ภาพเมฆจากจากดาวเทียม Himawari-8 หน่วยงาน kochi, digital typhoon,us_naval_research_lab
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/cloud_kochi
// @Method			GET
// @Summary			ภาพเมฆจากจากดาวเทียม Himawari-8 จาก kochi
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/cloud_us_naval_research_lab
// @Method			GET
// @Summary			ภาพเมฆ จาก us_naval_research_lab
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/cloud_digital_typhoon
// @Method			GET
// @Summary			ภาพเมฆ จาก digital typhoon
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/satellite_image_gistda
// @Method			GET
// @Summary			ข้อมูลปริมาณน้ำฝนจากภาพถ่ายดาวเทียม COMS
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/rain_image_us_naval_research_lab
// @Method			GET
// @Summary			แผนภาพฝนสะสมรายวัน
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/satellite_image_gsmaps
// @Method			GET
// @Summary			แผนภาพฝนสะสมรายวัน GSmaps
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/contour_image
// @Method			GET
// @Summary			แผนที่แสดงการกระจายตัวความกดอากาศ
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/modis_ndvi_usda
// @Method			GET
// @Summary			ค่าดัชนีพืชพรรณจากภาพถ่ายดาวเทียม
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/modis_soil_moisture_usda
// @Method			GET
// @Summary			ค่าความชื้นในดิน
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/modis_precipitation_usda
// @Method			GET
// @Summary			ปริมาณและการกระจายตัวของฝน
// @Produces		json
// @Response		200		AWeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/sst_ocean_weather
// @Method			GET
// @Summary			แผนภาพอุณหภูมิผิวน้ำทะเลประเทศไทย
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/wave_height_ocean_weather
// @Method			GET
// @Summary			ภาพความสูงคลื่น
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/gssh_aviso
// @Method			GET
// @Summary			แผนภาพค่าเบี่ยงเบนความสูงระดับน้ำทะเล
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/sst_2w_haii
// @Method			GET
// @Summary			แผนภาพการเปลี่ยนแปลงของอุณหภูมิผิวน้ำทะเล ราย 2 สัปดาห์
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/gsmap_hii
// @Method			GET
// @Summary			แผนภาพฝนสะสมรายวัน GsMap จาก  HII
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/ssh_event_haii
// @Method			GET
// @Summary			แผนภาพการศึกษาการเกิดพายุ SSH Event
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/ssh_w_haii
// @Method			GET
// @Summary			แผนภาพการเปลี่ยนแปลงความสูงของระดับน้ำทะเลรายสัปดาห์
// @Produces		json
// @Response		200		WeatherImageDataSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_img/animation
// @Method			GET
// @Summary			ภาพ animation
// @Parameter		-	query AnimationInput
// @Produces		json
// @Response		200		model_weather_image.WeatherAnimationOutput successful operation

////////////////////////////////////////////////////////////////////////////////////////////

type WeatherHistoryDataOutputSwagger struct {
	Result string                                         `json:"result"` //example:`OK`
	Data   []model_weather_image.WeatherHistoryDataOutput `json:"data"`
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_history_img/date
// @Method			GET
// @Summary			ภาพย้อนหลังรายวัน
// @Parameter		-	query HistoryDateInput
// @Produces		json
// @Response		200		WeatherHistoryDataOutputSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_history_img/date_range
// @Method			GET
// @Summary			ภาพย้อนหลังตามช่วงเวลา
// @Parameter		-	query HistoryDateRangeInput
// @Produces		json
// @Response		200		WeatherHistoryDataOutputSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_history_img/month
// @Method			GET
// @Summary			ภาพย้อนหลังรายเดือน
// @Parameter		-	query HistoryDateMonthInput
// @Produces		json
// @Response		200		WeatherHistoryDataOutputSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/weather_history_img/year
// @Method			GET
// @Summary			ภาพย้อนหลังรายปี
// @Parameter		-	query HistoryDateYearInput
// @Produces		json
// @Response		200		WeatherHistoryDataOutputSwagger successful operation

// @DocumentName	v1.public
// @Service			thaiwater30/public/storm_scale
// @Method			GET
// @Summary			Scale พายุ
// @Produces		json
// @Response		200		StormScaleSwagger successful operation
type StormScaleSwagger struct {
	Result string       `json:"result"` //example:`OK`
	Data   []StormScale `json:"data"`
}

type StormScale struct {
	Operator string `json:"operator"`   // example:`>`
	Term     string `json:"term"`       // example:`135`
	Color    string `json:"color"`      // example:`#CC00CC`
	Knots    string `json:"knots_text"` // example:`>135`
	MpH      string `json:"mph_text"`   // example:`>155`
	KmH      string `json:"kmh_text"`   // example:`>250`
	Category string `json:"category"`   // example:`Cat 5`
	Strength string `json:"strength"`   // example:`Typhoon Cat 5`
	Scale    string `json:"scale_text"` // example:`รุนแรงมาก`
}

type Struct_Cloud_Data struct {
	K []*model_latest_media.Struct_Media `json:"kochi"`   //  kochi
	T []*model_latest_media.Struct_Media `json:"typhoon"` //  Digital Typhoon
	U []*model_latest_media.Struct_Media `json:"us"`      //  US Naval Research Laboratory
}
