package weather_image

import (
	"encoding/json"
	tw30_setting "haii.or.th/api/thaiwater30/util/setting"
)

//type WeatherImageOutput struct {
//	Storm                          []*WeatherImageData `json:"storm"`
//	IndiaOceanUCL                  *WeatherImageData   `json:"india_ocean_ucl"`
//	PacificOceanUCL                *WeatherImageData   `json:"pacific_ocean_ucl"`
//	WeatherMapTMD                  *WeatherImageData   `json:"weather_map_tmd"`
//	WeatherMapHD                   *WeatherImageData   `json:"weather_map_hd"`
//	CloudKochi                     *WeatherImageData   `json:"cloud_kochi"`
//	CloudUSNRL                     *WeatherImageData   `json:"cloud_us_naval_research_lab"`
//	CloudDT                        *WeatherImageData   `json:"cloud_digital_typhoon"`
//	SatelliteComsAsia              *WeatherImageData   `json:"satellite_image_asia"`
//	SatelliteComsSoutheastAsia     *WeatherImageData   `json:"satellite_image_southeast_asia"`
//	SatelliteCcomsThailand         *WeatherImageData   `json:"satellite_image_thailan"`
//	RainImageDailyUSNavi           *WeatherImageData   `json:"rain_image_us_navi"`
//	RainImageDailyGSmaps           *WeatherImageData   `json:"rain_image_gsmaps"`
//	RainImageDailyTRMM             *WeatherImageData   `json:"rain_image_trmm"`
//	ContourTemperature             *WeatherImageData   `json:"contour_temperature"`
//	ContourHumidity                *WeatherImageData   `json:"contour_humidity"`
//	ContourPressure                *WeatherImageData   `json:"contour_pressure"`
//	VegetationindexTerraUSDA       *WeatherImageData   `json:"vegetationindex_terra_usda"`
//	VegetationindexAquaUSDA        *WeatherImageData   `json:"vegetationindex_aqua_usda"`
//	SoilMoitsureUSDA               *WeatherImageData   `json:"soil_moitsure_usda_afwa_surface"`
//	SoilMoitsureUSDAAFWASubSurface *WeatherImageData   `json:"soil_moitsure_usda_afwa_subsurface"`
//	SoilMoitsureUSDAWMOSurface     *WeatherImageData   `json:"soil_moitsure_usda_who_surface"`
//	SoilMoitsureUSDAWMOSubSurface  *WeatherImageData   `json:"soil_moitsure_usda_who_subsurface"`
//	//	DistributionRainUSDA               *WeatherImageData   `json:"distribution_rain_usda"`
//	MapTemperatureOceanWeatherThailand *WeatherImageData `json:"map_temperature_ocean_weather_thailand"`
//	MapTemperatureOceanWeatherPacific  *WeatherImageData `json:"map_temperature_ocean_weather_pacific"`
//	MapTemperatureOceanWeatherWorld    *WeatherImageData `json:"map_temperature_ocean_weather_global"`
//	MapTemperatureOceanWeatherIndia    *WeatherImageData `json:"map_temperature_ocean_weather_india"`
//	MapWaveOceanWeatherThailand        *WeatherImageData `json:"map_wave_ocean_weather_thailand"`
//	MapWaveOceanWeatherPacific         *WeatherImageData `json:"map_wave_ocean_weather_pacific"`
//	MapWaveOceanWeatherWorld           *WeatherImageData `json:"map_wave_ocean_weather_global"`
//	MapWaveOceanWeatherIndia           *WeatherImageData `json:"map_wave_ocean_weather_india"`
//	MapHighSDAvisoJason1               *WeatherImageData `json:"map_highSD_aviso_jason1"`
//	MapSeaTemperatureHAII              *WeatherImageData `json:"map_sea_temperature_haii"`
//	MapSSHEvent                        *WeatherImageData `json:"map_ssh_event"`
//	MapSeaWaterlevelWeek               *WeatherImageData `json:"map_seawaterlevel_week"`
//}

type WeatherImageData struct {
	Description *Description     `json:"description"`           // คำอธิบาย
	Agency      *Agency          `json:"agency"`                // หน่วยงาน
	Scale       *json.RawMessage `json:"scale,omitempty"`       // example:`{"operator": "<","term": "34","color": "#00CCFF","knots_text": "<34","mph_text": "<39","kmh_text": "<63","category": "TD", "strength": "Tropical Depression","scale_text": "รุนแรงน้อย"}` scale
	CoverImage  *CoverImage      `json:"cover_image"`           // ภาพ cover
	Details     []*Detail        `json:"detail"`                // รายละเอียด
	GroupIMG    []*CoverImage    `json:"group_image,omitempty"` // กรุ๊ปภาพ
}

type WeatherImageDataSwagger struct {
	Description *DescriptionSwagger               `json:"description"`           // คำอธิบาย
	Agency      *Agency                           `json:"agency"`                // หน่วยงาน
	Scale       *tw30_setting.Struct_StormSetting `json:"scale,omitempty"`       // scale
	CoverImage  *CoverImage                       `json:"cover_image"`           // ภาพ cover
	Details     []*DetailSwagger                  `json:"detail"`                // รายละเอียด
	GroupIMG    []*CoverImage                     `json:"group_image,omitempty"` // กรุ๊ปภาพ
}
type DetailSwagger struct {
	Icon        *Icon           `json:"icon"`          // icon
	Description json.RawMessage `json:"description"`   // example:`{"th":"คำอธิบายภาษาไทย"}`คำอธิบาย
	Link        string          `json:"link"`          // example:`http://web.thaiwater.net/thaiwater30/` ลิ้งค์
	LinkType    string          `json:"link_type"`     // example:`_blank`ชนิดของลิ้งค์
	MediaTypeID interface{}     `json:"media_type_id"` // example:`134` รหัสประเภทสื้อ
}

type DescriptionSwagger struct {
	DescriptionName json.RawMessage `json:"description_name"` // example:`{"th":"คำอธิบายภาษาไทย"}`คำอธิบาย
	DescriptionLink string          `json:"description_link"` // example:`http://web.thaiwater.net/thaiwater30/`ลิ้งค์
	LinkType        string          `json:"link_type"`        // example:`_blank`ชนิดของลิ้งค์
}

type Description struct {
	DescriptionName *DescriptionNameLang `json:"description_name"` // คำอธิบาย
	DescriptionLink string               `json:"description_link"` // example:`http://web.thaiwater.net/thaiwater30/`ลิ้งค์
	LinkType        string               `json:"link_type"`        // example:`_blank`ชนิดของลิ้งค์
}

type DescriptionNameLang struct {
	TH string `json:"th"` // example:`ไทย`
	EN string `json:"en"` // example:`english`
	JP string `json:"jp"` // example:`japan`
}

type Agency struct {
	AgencyName json.RawMessage `json:"agency_name"` // example:`{"th":"ชื่อหน่วยงาน"}`ชื่อหน่ยงาน
	AgencyID   int64           `json:"agency_id"`   // example:`9`รหัสหน่วยงาน
	AgencyLink string          `json:"agency_link"` // example:`http://web.thaiwater.net/thaiwater30/` ลิ้งค์หน่วยงาน
	LinkType   string          `json:"link_type"`   // example:`_blank`ชนิดของลิ้งค์
}

type CoverImage struct {
	FilePath           string `json:"filepath"`             // example:`/product/img/`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename           string `json:"filename"`             // example:`filename.jpg`ชื่อของไฟล์ข้อมูลสื่อ
	MediaPath          string `json:"media_path"`           // example:`QWe2131DAqweqEQw142`ลิ้งค์ของไฟล์ข้อมูลสื่อ
	ThumbnailFilePath  string `json:"thumbnail_filepath"`   // example:`/product/img/` ที่อยู่ของไฟล์ thumb ข้อมูลสื่อ
	ThumbnailFilename  string `json:"thumbnail_filename"`   // example:`thumb-filename.jpg` ชื่อไฟล์ thumb ของข้อมูลสื่อ
	ThumbnailMediaPath string `json:"thumbnail_media_path"` // example:`QWe2131DAqweqEQw142` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	IsStatic           bool   `json:"is_static"`            // example:`false`เป็น static?
	CoverLink          string `json:"cover_link"`           // example:`http://web.thaiwater.net/thaiwater30/`ลิ้งค์ภาพ cover
	LinkType           string `json:"link_type"`            // example:`_blank` ชนิดลิ้งค์
	MediaDatetime      string `json:"media_datetime"`       // example:`2006-01-02 15:04` วันที่เก็บข้อมูลสื่อ
}

type Detail struct {
	Icon        *Icon                `json:"icon"`          // icon
	Description *DescriptionNameLang `json:"description"`   // คำอธิบาย
	Link        string               `json:"link"`          // example:`http://web.thaiwater.net/thaiwater30/` ลิ้งค์
	LinkType    string               `json:"link_type"`     // example:`_blank`ชนิดของลิ้งค์
	MediaTypeID interface{}          `json:"media_type_id"` // example:`134` รหัสประเภทสื้อ
}

type Icon struct {
	FilePath  string `json:"filepath"`   // example:`/product/img/`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename  string `json:"filename"`   // example:`filename.jpg`ชื่อของไฟล์ข้อมูลสื่อ
	MediaPath string `json:"media_path"` // example:`QWe2131DAqweqEQw142`ลิ้งค์ของไฟล์ข้อมูลสื่อ
}

type WeatherImageHistoryAnimation struct {
	Agency    int64 `json:"agency_id"`     // example:`9` รหัสหน่วยงาน
	MediaType int64 `json:"media_type_id"` // example:`15` รหัสประเภทสื้อ
}

type WeatherImageHistorySEDate struct {
	Agency    int64  `json:"agency_id"`     // รหัสหน่วยงาน
	MediaType int64  `json:"media_type_id"` // รหัสประเภทสื้อ
	StartDate string `json:"start_date"`    // วันเริ่มต้นของข้อมูลสื้อ
	EndDate   string `json:"end_date"`      // วันสิ้นสุดของข้อมูลสื้อ
}

type WeatherImageHistoryDate struct {
	Agency    int64  `json:"agency_id"`     // รหัสหน่วยงาน
	MediaType int64  `json:"media_type_id"` // รหัสประเภทสื้อ
	Date      string `json:"date"`          // วันที่เก็บข้อมูลสื่อ
}

type WeatherImageHistoryAllParams struct {
	Agency    int64  `json:"agency_id"`     // รหัสหน่วยงาน เช่น 9
	MediaType int64  `json:"media_type_id"` // รหัสประเภทสื้อ เช่น 141
	Date      string `json:"date"`          // วันที่เก็บข้อมูลสื่อ เช่น 2006-01-02
	StartDate string `json:"start_date"`    // วันเริ่มต้นของข้อมูลสื้อ เช่น 2006-01-02
	EndDate   string `json:"end_date"`      // วันสิ้นสุดของข้อมูลสื้อ เช่น 2006-01-02
	Year      string `json:"year"`          // ปี เช่น 2016
	Month     string `json:"month"`         // month เช่น 07
}

type WeatherImageInput struct {
	Agency    int64  `json:"agency_id"`     // รหัสหน่วยงาน
	MediaType int64  `json:"media_type_id"` // รหัสประเภทสื้อ
	Date      string `json:"date"`          // วันที่เก็บข้อมูลสื่อ
}

type WeatherImageDataOutput struct {
	Datetime string          `json:"datetime"` // วันที่เก็บข้อมูลสื่อ
	FilePath string          `json:"filepath"` // ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename string          `json:"filename"` // ชื่อของไฟล์ข้อมูลสื่อ
	Agency   json.RawMessage `json:"agency"`   // หน่วยงาน
}

type WeatherAnimationOutput struct {
	FilePath  string `json:"filepath"`   // example:`/product/img/`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename  string `json:"filename"`   // example:`filename.mp4`ชื่อของไฟล์ข้อมูลสื่อ
	MediaPath string `json:"media_path"` // example:`QWe2131DAqweqEQw142`ลิ้งค์ของไฟล์ข้อมูลสื่อ
}

type WeatherHistoryDataOutput struct {
	// เพิ่ม struc ส่วนนี้มาเพิ่มแสดงผลข้ัอมูล รายละเอียดที่มากกขึ้น
	MediaTypeID   int64       `json:"media_type_id"`    // example: 18  รหัสชนิดข้อมูลสื้อ
	MediaType     string      `json:"media_type"`       // example:`Precipitation` ชนิดของสื่อ
	MediaSubType  string      `json:"media_subtype"`    // example:`Thailand` ชนิดย่อยข้อมูลสื่อ
	MediaCategory string      `json:"media_category"`   // example:`image` ประเภทของ media เช่น image, animation, excel etc
	Description   string      `json:"description"`      // example:`description` รายละเอียดของข้อมูลสื่อ
	URLThumb      interface{} `json:"media_path_thumb"` // example:`QWE1QTH3AD@LKH1238D` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	FilenameThumb interface{} `json:"filename_thumb"`   // example:`thumb-thailand.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	FilepathThumb interface{} `json:"filepath_thumb"`   // example:`/product/precipitation/thailand/2016/06/10` ที่อยู่ของไฟล์ thumb ข้อมูลสื่อ

	FilePath           string `json:"filepath"`             // example:`/product/img/`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename           string `json:"filename"`             // example:`filename.jpg`ชื่อของไฟล์ข้อมูลสื่อ
	MediaPath          string `json:"media_path"`           // example:`QWe2131DAqweqEQw142`ลิ้งค์ของไฟล์ข้อมูลสื่อ
	ThumbnailFilePath  string `json:"thumbnail_filepath"`   // example:`/product/img/` ที่อยู่ของไฟล์ thumb ข้อมูลสื่อ
	ThumbnailFilename  string `json:"thumbnail_filename"`   // example:`thumb-filename.jpg` ชื่อของไฟล์ thumb ข้อมูลสื่อ
	ThumbnailMediaPath string `json:"thumbnail_media_path"` // example:`QWe2131DAqweqEQw142` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	MediaDatetime      string `json:"media_datetime"`       // example:`2006-01-02 15:04`วันที่เก็บข้อมูลสื่อ
}
