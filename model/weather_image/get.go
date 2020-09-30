// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package weather_image is a model for public.media table. This table store media information.
package weather_image

import (
	"database/sql"
	//	"haii.or.th/api/server/model"
	//	"encoding/json"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/util/b64"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/thumbnail"
	//	"os"
	//	"path/filepath"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	prefixWebImage       = "public/resource/image/" //default public image web
	prefixAPIStaticImage = ""
	prefixAnimation      = "product/animation/" // "product/images/" // default animation
)

//Get sea waterlvel image latest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapSeaWaterlevelWeekLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพการเปลี่ยนแปลงความสูงของระดับน้ำทะเลรายสัปดาห์", "", "", "", "_blank")
	data.Agency = addAgency(9, "https://www.haii.or.th/", "_blank")
	data.CoverImage = addCoverLatest(9, 140, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ดูข้อมูลย้อนหลัง", "", "", "http://nhc.in.th/web/index.php?model=ssh&view=ssh_w", "_blank", 140, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  Get map ssh event
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapSSHEventLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพการศึกษาการเกิดพายุ SSH Event", "", "", "", "_blank")
	data.Agency = addAgency(9, "https://www.haii.or.th/", "_blank")
	data.CoverImage = addCoverLatest(9, 142, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ดูข้อมูลย้อนหลัง", "", "", "http://nhc.in.th/web/index.php?model=ssh&view=ssh_event", "_blank", 142, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetMapSeaTemperatureHAIILatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapSeaTemperatureHAIILatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพการเปลี่ยนแปลงของอุณหภูมิผิวน้ำทะเล ราย 2 สัปดาห์ โดย สสนก.", "", "", "", "_blank")
	data.Agency = addAgency(9, "https://www.haii.or.th/", "_blank")
	data.CoverImage = addCoverLatest(9, 14, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/Tracking/Now/SST_W/gallery/", "_blank", 14, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetMapHighSDAvisoJason1Latest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapHighSDAvisoJason1Latest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพค่าเบี่ยงเบนความสูงระดับน้ำทะเล ดาวเทียม Jason-1 GFO และ ENVISAT โดย  Aviso", "", "", "", "_blank")
	data.Agency = addAgency(57, "http://www.aviso.oceanobs.com/en/data/products.html", "_blank")
	data.CoverImage = addCoverLatest(57, 153, "", "modal", "")

	// add details
	animation, _ := getAnimation("image/ssh/global/aviso/media", "A_GSSH_merge_latest.gif")
	details := make([]*Detail, 0)
	detail := addDetail("ภาพเคลื่อนไหว", "", "", animation.MediaPath, "_blank", 153, "", "")
	details = append(details, detail)
	detail = addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/show_ssh.php", "_blank", 153, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetMapGsmap10kmHiiLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapGsmap10kmHiiLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพฝนสะสมรายวัน  GSMaP (10km x 10km) โดย  HII", "", "", "", "_blank")
	data.Agency = addAgency(9, "http://www.thaiwater.net", "_blank")
	data.CoverImage = addCoverLatest(9, 157, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net", "_blank", 157, "", "")
	details = append(details, detail)

	return data, nil
}

//  GetMapGsmapPersiann4kmHiiLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapGsmapPersiann4kmHiiLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพฝนสะสมรายวัน  PERSIANN (4km x 4km) โดย  HII", "", "", "", "_blank")
	data.Agency = addAgency(9, "http://www.thaiwater.net", "_blank")
	data.CoverImage = addCoverLatest(9, 158, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net", "_blank", 158, "", "")
	details = append(details, detail)

	return data, nil
}

//  GetMapGsmap25kmHiiLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetMapGsmap25kmHiiLatest() (*WeatherImageData, error) {
	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพฝนสะสมรายวัน  GSMaP (25km x 25km) โดย  HII", "", "", "", "_blank")
	data.Agency = addAgency(9, "http://www.thaiwater.net", "_blank")
	data.CoverImage = addCoverLatest(9, 159, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net", "_blank", 159, "", "")
	details = append(details, detail)

	return data, nil
}

//  GetWaveOceanWeatherThailandLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetWaveOceanWeatherThailandLatest() (*WeatherImageData, error) {
	var media_type_id int64 = 24
	var agency_id int64 = 56
	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพความสูงและทิศทางของคลื่นทะเลประเทศไทย โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(agency_id, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(agency_id, media_type_id, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ประเทศไทย", "", "", "", "_blank", media_type_id, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetWaveOceanWeatherNpwLatest
//  Parameters:
//    None
//  Return:
//    WeatherImageData
func GetWaveOceanWeatherNpwLatest() (*WeatherImageData, error) {
	var media_type_id int64 = 156
	var agency_id int64 = 56
	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพความสูงและทิศทางของคลื่นทะเลมหาสมุทรแปซิฟิก โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(agency_id, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(agency_id, media_type_id, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("มหาสมุทรแปซิฟิก", "", "", "", "_blank", media_type_id, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetWaveOceanWeatherIndLatest
//  Parameters:
//    None
//  Return:
//    WeatherImageData
func GetWaveOceanWeatherIndLatest() (*WeatherImageData, error) {
	var media_type_id int64 = 23
	var agency_id int64 = 56
	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพความสูงและทิศทางของคลื่นทะเลมหาสมุทรอินเดีย โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(agency_id, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(agency_id, media_type_id, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("มหาสมุทรอินเดีย", "", "", "", "_blank", media_type_id, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetWaveOceanWeatherGlobalLatest
//  Parameters:
//    None
//  Return:
//    WeatherImageData
func GetWaveOceanWeatherGlobalLatest() (*WeatherImageData, error) {
	var media_type_id int64 = 155
	var agency_id int64 = 56
	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพความสูงและทิศทางของคลื่นทะเลทั่วโลก โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(agency_id, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(agency_id, media_type_id, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ทั่วโลก", "", "", "", "_blank", media_type_id, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetPrecipitationUSDAWMODecadalPercentNormalPrecipitationLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetPrecipitationUSDAWMODecadalPercentNormalPrecipitationLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ปริมาณและการกระจายตัวของฝน WMO Decadal Percent Normal โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 152, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 152, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetPrecipitationUSDAWMOPrecipitationLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetPrecipitationUSDAWMOPrecipitationLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ปริมาณและการกระจายตัวของฝน WMO โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 151, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 151, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetPrecipitationUSDAAFWADecadalPercentNormalPrecipitationLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetPrecipitationUSDAAFWADecadalPercentNormalPrecipitationLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ปริมาณและการกระจายตัวของฝน AFWA Decadal Percent Normal โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 150, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 150, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetPrecipitationUSDAAFWAPrecipitationLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetPrecipitationUSDAAFWAPrecipitationLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ปริมาณและการกระจายตัวของฝน AFWA โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 149, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 149, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

func GetTemperatureOceanWeatherThailandLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("อุณหภูมิผิวน้ำทะเลประเทศไทย โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(56, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(56, 143, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ประเทศไทย", "", "", "", "_blank", 143, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

func GetTemperatureOceanWeatherGlobalLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("อุณหภูมิผิวน้ำทะเลทั่วโลก โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(56, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(56, 145, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ทั่วโลก", "", "", "", "_blank", 145, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

func GetTemperatureOceanWeatherIndLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("อุณหภูมิผิวน้ำทะเลมหาสมุทรอินเดีย โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(56, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(56, 146, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("มหาสมุทรอินเดีย", "", "", "", "_blank", 146, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

func GetTemperatureOceanWeatherNpwLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("อุณหภูมิผิวน้ำทะเลมหาสมุทรแปซิฟิก โดย Ocean Weather inc.", "", "", "", "_blank")
	data.Agency = addAgency(56, "http://www.oceanweather.com/", "_blank")
	data.CoverImage = addCoverLatest(56, 144, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("มหาสมุทรแปซิฟิก", "", "", "", "_blank", 144, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetSoilMoitsureUSDAWMOSubSurfaceLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSoilMoitsureUSDAWMOSubSurfaceLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าความชื้นในดิน WMO Sub-Surface โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 163, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 161, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetSoilMoitsureUSDAWMOSurfaceLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSoilMoitsureUSDAWMOSurfaceLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าความชื้นในดิน WMO Surface โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 162, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 160, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetSoilMoitsureUSDAAFWASubSurfaceLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSoilMoitsureUSDAAFWASubSurfaceLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าความชื้นในดิน AFWA Sub-Surface โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 161, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 161, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

//  GetSoilMoitsureUSDALatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSoilMoitsureUSDALatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าความชื้นในดิน AFWA Surface โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 160, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#soilmoisture", "_blank", 160, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//  GetVegetationindexAquaUSDALatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetVegetationindexAquaUSDALatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าดัชนีพืชพรรณจากภาพถ่ายดาวเทียม AQUA โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 36, "next.gif", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#ndvi", "_blank", 36, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//  GetVegetationindexTerraUSDALatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetVegetationindexTerraUSDALatest() (*WeatherImageData, error) {
	//agencyID = 55
	//mediaType = 37

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ค่าดัชนีพืชพรรณจากภาพถ่ายดาวเทียม TERRA โดย  United States Department of Agriculture", "", "", "", "_blank")
	data.Agency = addAgency(55, "https://www.pecad.fas.usda.gov/", "_blank")
	data.CoverImage = addCoverLatest(55, 37, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/main.php#ndvi", "_blank", 37, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//1.12
//  GetContourPressureLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetContourPressureLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่แสดงการกระจายตัวความกดอากาศ", "", "", "", "_blank")
	data.Agency = addAgency(9, "", "_blank")
	data.CoverImage = addCoverLatest(9, 3, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/radar/show_pressImg.php", "_blank", 3, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.12
//  GetContourHumidityLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetContourHumidityLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่แสดงการกระจายตัวความชื้น", "", "", "", "_blank")
	data.Agency = addAgency(9, "", "_blank")
	data.CoverImage = addCoverLatest(9, 1, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/radar/show_humidImg.php", "_blank", 1, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.12
//  GetContourTemperatureLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetContourTemperatureLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่แสดงการกระจายตัวอุณหภูมิ", "", "", "", "_blank")
	data.Agency = addAgency(9, "", "_blank")
	data.CoverImage = addCoverLatest(9, 2, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/radar/show_tempImg.php", "_blank", 2, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.11
//  GetRainImageDailyTRMMLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetRainImageDailyTRMMLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("Persiann", "", "", "", "_blank")
	data.Agency = addAgency(54, "", "_blank")
	data.CoverImage = addCoverLatest(54, 6, "", "modal", "")
	t := time.Now()
	t1 := t.AddDate(0, 0, -1)
	t2 := t.AddDate(0, 0, -7)
	p := []interface{}{54, 6, t2.Format("2006-01-02") + "07:00", t1.Format("2006-01-02") + "07:00"}
	data.GroupIMG, _ = addCoverHistoryForLatest(sqlSelectHistory7Day, p, 54, 6, "modal", "")
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/v3/gsmap#db", "_blank", 6, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.11
//  GetRainImageDailyGSmapsLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetRainImageDailyGSmapsLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพฝนสะสมรายวัน GSmaps", "", "", "", "_blank")
	data.Agency = addAgency(53, "", "_blank")
	data.CoverImage = addCoverLatest(53, 6, "", "modal", "")
	t := time.Now()
	t1 := t.AddDate(0, 0, -1)
	t2 := t.AddDate(0, 0, -7)
	p := []interface{}{53, 6, t2.Format("2006-01-02") + " 07:00", t1.Format("2006-01-02") + " 07:00"}
	data.GroupIMG, _ = addCoverHistoryForLatest(sqlSelectHistory7Day, p, 53, 6, "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/v3/gsmap#db", "_blank", 6, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//1.10
//  GetRainImageDailyUSNRLLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetRainImageDailyUSNRLLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนภาพฝนสะสมรายวัน", "", "", "", "_blank")
	data.Agency = addAgency(51, "https://www.nrlmry.navy.mil/", "_blank")
	data.CoverImage = addCoverLatest(51, 6, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ภาพเคลื่อนไหว", "", "", "http://www.thaiwater.net/TyphoonTracking/trmmImages/ssmi_trmm_amsub_accumulations/A_SSTA_latest.gif", "_blank", 6, "", "")
	details = append(details, detail)
	detail = addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/show_ssta.php", "_blank", 6, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//1.9
//  GetSatelliteCcomsThailandLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSatelliteCcomsThailandLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ข้อมูลปริมาณน้ำฝนจากภาพถ่ายดาวเทียม COMS, Thailand", "", "", "", "_blank")
	data.Agency = addAgency(11, "http://eo.gistda.or.th:8090", "_blank")
	data.CoverImage = addCoverLatest(11, 4, "", "modal", "")
	return data, nil
}

//1.9
//  GetSatelliteComsSoutheastAsiaLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSatelliteComsSoutheastAsiaLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ข้อมูลปริมาณน้ำฝนจากภาพถ่ายดาวเทียม COMS, SoutheastAsia", "", "", "", "_blank")
	data.Agency = addAgency(11, "http://eo.gistda.or.th:8090", "_blank")
	data.CoverImage = addCoverLatest(11, 6, "", "modal", "")

	return data, nil
}

//1.9
//  GetSatelliteComsAsiaLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetSatelliteComsAsiaLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ข้อมูลปริมาณน้ำฝนจากภาพถ่ายดาวเทียม COMS, Asia", "", "", "", "_blank")
	data.Agency = addAgency(11, "http://eo.gistda.or.th:8090", "_blank")
	data.CoverImage = addCoverLatest(11, 5, "", "modal", "")

	return data, nil
}

//1.8
//  GetCloudDTLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetCloudDTLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ภาพเมฆ", "", "", "", "_blank")
	data.Agency = addAgency(52, "http://agora.ex.nii.ac.jp/digital-typhoon/region/SEasia/2/", "_blank")
	data.CoverImage = addCoverLatest(52, 141, "", "modal", "")
	return data, nil
}

//1.7
//  GetCloudUSNRLLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetCloudUSNRLLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ภาพเมฆ", "", "", "", "_blank")
	data.Agency = addAgency(51, "https://www.nrl.navy.mil/", "_blank")
	data.CoverImage = addCoverLatest(51, 141, "", "modal", "")

	return data, nil
}

//1.6
//  GetCloudKochiLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetCloudKochiLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8", "", "", "http://www.thaiwater.net/~gms/", "_blank")
	data.Agency = addAgency(50, "http://weather.is.kochi-u.ac.jp/index-e.html", "_blank")
	data.CoverImage = addCoverLatest(50, 141, "ql.%", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ภาพเคลื่อนไหวจากต้นทาง", "", "", "http://www.thaiwater.net/gms/weather/A24_GMS_lastest_large.gif", "_blank", 141, "", "")
	details = append(details, detail)
	animation, _ := getAnimation("himawari-8/thailand/kochi/media/", "00Movie.mp4")

	detail = addDetail("ภาพเคลื่อนไหว", "", "", animation.MediaPath, "_blank", 141, "", "")
	details = append(details, detail)
	data.Details = details

	return data, nil
}

//1.5
//  GetWeatherMapHDLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetWeatherMapHDLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่ลมฟ้าอากาศ", "", "", "", "_blank")
	data.Agency = addAgency(6, "http://www.hydro.navy.mi.th/", "_blank")
	data.CoverImage = addCoverLatest(6, 27, "", "modal", "")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ดูข้อมูลย้อนหลัง", "", "", "http://www.nhc.in.th/web/index.php?model=weather_map_rtn", "_blank", 27, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.4
//  GetWeatherMapTMDLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetWeatherMapTMDLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่อากาศ", "", "", "http://www.thaiwater.net/TyphoonTracking/prewc1.php", "_blank")
	data.Agency = addAgency(13, "http://www.tmd.go.th/", "_blank")
	data.CoverImage = addCoverLatest(13, 22, "", "modal", "http://www.thaiwater.net/TyphoonTracking/wc.php?imgwc=lastest_wc.jpg")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/hmain.php?page=/TyphoonTracking/show_weather_map.php", "_blank", 22, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.4
//  GetWeatherMapWind850TMDLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetWeatherMapWind850TMDLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่ลมชั้นบนระดับ 850 ฟุต กรมอุตุนิยมวิทยา", "", "", "http://www.thaiwater.net/TyphoonTracking/prewc1.php", "_blank")
	data.Agency = addAgency(13, "http://www.tmd.go.th/", "_blank")
	data.CoverImage = addCoverLatest(13, 28, "", "modal", "http://www.thaiwater.net/TyphoonTracking/wc.php?imgwc=lastest_wc.jpg")

	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/hmain.php?page=/TyphoonTracking/show_weather_map.php", "_blank", 28, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.4
//  GetWeatherMapWind925TMDLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetWeatherMapWind925TMDLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่ลมชั้นบนระดับ 925 ฟุต กรมอุตุนิยมวิทยา", "", "", "http://www.thaiwater.net/TyphoonTracking/prewc1.php", "_blank")
	data.Agency = addAgency(13, "http://www.tmd.go.th/", "_blank")
	data.CoverImage = addCoverLatest(13, 29, "", "modal", "http://www.thaiwater.net/TyphoonTracking/wc.php?imgwc=lastest_wc.jpg")

	details := make([]*Detail, 0)
	detail := addDetail("รูปภาพจากคลัง", "", "", "http://www.thaiwater.net/DATA/REPORT/php/hmain.php?page=/TyphoonTracking/show_weather_map.php", "_blank", 29, "", "")
	details = append(details, detail)

	data.Details = details

	return data, nil
}

//1.3
//  GetPacificOceanUCLLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetPacificOceanUCLLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่วิเคราะห์เส้นทางและความรุนแรงงของพายุ ในมหาสมุทรแปซิฟิก จัดทำโดย University College London", "", "", "", "_blank")
	data.Agency = addAgency(41, "http://www.tropicalstormrisk.com/", "_blank")
	data.CoverImage = addCoverLatest(41, 62, "W.png", "modal", "")

	j := setting.GetSystemSettingJSON("Frontend.public.storm_setting")
	data.Scale = &j
	return data, nil
}

//1.2
//  GetIndiaOceanUCLLatest
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetIndiaOceanUCLLatest() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("แผนที่วิเคราะห์เส้นทางและความรุนแรงงของพายุ ในมหาสมุทรอินเดีย จัดทำโดย University College London", "", "", "", "_blank")
	data.Agency = addAgency(41, "http://www.tropicalstormrisk.com/", "_blank")
	data.CoverImage = addCoverLatest(41, 62, "I.png", "modal", "")

	j := setting.GetSystemSettingJSON("Frontend.public.storm_setting")
	data.Scale = &j
	return data, nil
}

//1.1
//  GetStormHistory
//  Parameters:
//		None
//  Return:
//		WeatherImageData
func GetStormHistory() (*WeatherImageData, error) {

	// add cover
	data := &WeatherImageData{}
	data.Description = addDescription("ภาพพายุหมุนในอดีตตั้งแต่ พ.ศ.2494 จนถึงปัจจุบัน", "", "", "http://www.thaiwater.net/TyphoonTracking/main.html", "_blank")
	//	data.Agency = addAgency(49, "")
	data.CoverImage = addCoverStatic(prefixWebImage+"weather", "storm.jpg", "http://www.thaiwater.net/TyphoonTracking/storm/storm.html", "_blank")

	// add details
	details := make([]*Detail, 0)
	detail := addDetail("ชมพายุ", "storm", "", "http://www.thaiwater.net/TyphoonTracking/storm/storm.html", "_blank", 0, "", "")
	details = append(details, detail)

	data.Details = details
	return data, nil
}

// func add detail description lang ,link,linktype , mediatype, icon path, icon name
//  Parameters:
//		descTH
//			description lang th
//		descEN
//			description lang en
//		descJP
//			description lang jp
//		link
//			url image
//		linkType
//			link type
//		mediaTypeID
//			media type id
//		iconPath
//			icon path
//		iconName
//			icon name
//  Return:
//		Detail
func addDetail(descTH, descEN, descJP, link, linkType string, mediaTypeID int64, iconPath, iconName string) *Detail {

	// define struct for data
	detail := &Detail{}

	// add data
	desc := &DescriptionNameLang{}
	desc.TH = descTH
	desc.EN = descEN
	desc.JP = descJP

	detail.Description = desc
	detail.Link = link
	detail.LinkType = linkType

	if mediaTypeID != 0 {
		detail.MediaTypeID = mediaTypeID
	}
	icon := &Icon{}
	if iconPath+iconName != "" {
		icon.FilePath = iconPath
		icon.Filename = iconName
		icon.MediaPath, _ = b64.EncryptText(iconPath + "/" + iconName)
	}
	detail.Icon = icon

	// return data
	return detail
}

// func add static cover
//  Parameters:
//		path
//			filepath
//		filename
//			filename
//		link
//			url image
//		linkType
//			url type
//  Return:
//		CoverImage
func addCoverStatic(path, filename, link, linkType string) *CoverImage {

	// define struct for data
	cover := &CoverImage{}
	cover.MediaPath, _ = b64.EncryptText(path + "/" + filename)
	cover.FilePath = path
	cover.Filename = filename
	// find thumbnail and add to struct
	tbname := thumbnail.GetThumbName(filename, "", "")
	cover.ThumbnailFilePath = path
	cover.ThumbnailFilename = strings.Replace(tbname, "/", "", 1)
	cover.ThumbnailMediaPath, _ = b64.EncryptText(path + "/" + tbname)

	cover.IsStatic = true
	cover.CoverLink = link
	if linkType == "" {
		linkType = "_blank"
	}
	cover.LinkType = linkType

	//return data
	return cover
}

// func add cover latest
//  Parameters:
//		agencyID
//			agency id
//		mediaType
//			media type id
//		fname
//			filename
//		linkType
//			url type
//		coverLink
//			cover url
//  Return:
//		CoverImage
func addCoverLatest(agencyID, mediaType int64, fname, linkType, coverLink string) *CoverImage {

	// define stuct for data
	cover := &CoverImage{}
	db, err := pqx.Open()
	if err != nil {
		return cover
	}

	// query image latest
	q := sqlSelectMediaImageLatest
	p := []interface{}{mediaType, agencyID}
	if fname != "" {
		q += " AND filename like $3"
		p = append(p, fname)
	}

	q += " ORDER BY media_type_id,agency_id, media_datetime DESC"

	//	fmt.Println(q)
	rows, err := db.Query(q, p...)
	if err != nil {
		return cover
	}
	defer rows.Close()

	for rows.Next() {
		var (
			agency   pqx.JSONRaw
			datetime time.Time
			path     sql.NullString
			filename sql.NullString
			ref      sql.NullString
		)
		// scan data from table
		rows.Scan(&datetime, &agency, &path, &filename, &ref)
		// add cover image
		cover.MediaPath, _ = b64.EncryptText(path.String + "/" + filename.String)
		cover.FilePath = path.String
		cover.Filename = filename.String
		src := filepathx.JoinPath(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), path.String, filename.String)
		tbsrc := thumbnail.GetThumbName(src, "", "")
		// check thumbnail
		if filepathx.ValidateFile(tbsrc) == nil {
			// add thumbnail
			fthumb := filepathx.Base(tbsrc)
			cover.ThumbnailMediaPath, _ = b64.EncryptText(path.String + "/" + fthumb)
			cover.ThumbnailFilePath = path.String
			cover.ThumbnailFilename = strings.Replace(fthumb, "/", "", 1)
		}
		cover.MediaDatetime = udt.DatetimeFormat(datetime, "datetime")
	}
	cover.CoverLink = coverLink
	if linkType == "" {
		linkType = "_blank"
	}
	cover.LinkType = linkType

	// return data
	return cover
}

// func add agency to struct
//  Parameters:
//		agencyID
//			agency id
//		agencyLink
//			agency link
//		linkType
//			url type
//  Return:
//		CoverImage
func addAgency(agencyID int64, agencyLink, linkType string) *Agency {

	db, err := pqx.Open()
	agency := &Agency{}
	if linkType == "" {
		linkType = "_blank"
	}
	agency.AgencyLink = agencyLink
	agency.LinkType = linkType

	if err != nil {
		return agency
	}
	// sql get agency name by id
	q := sqlSelectAgencyByID
	p := []interface{}{agencyID}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return agency
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int64
			agencyName pqx.JSONRaw
		)
		// scan data
		rows.Scan(&id, &agencyName)
		agency.AgencyID = id
		agency.AgencyName = agencyName.JSON()
		// add agency name
	}

	return agency
}

// func add description struct for data output
//  Parameters:
//		descTH
//			description lang th
//		descEN
//			description lang en
//		descJP
//			description lang jp
//		link
//			url image
//		linkType
//			link type
//		mediaTypeID
//			media type id
//  Return:
//		Description
func addDescription(descTH, descEN, descJP, descLink, linkType string) *Description {
	// define struct
	desc := &Description{}
	// add description name
	descName := &DescriptionNameLang{}
	descName.TH = descTH
	descName.EN = descEN
	descName.JP = descJP
	desc.DescriptionName = descName
	desc.DescriptionLink = descLink
	if linkType == "" {
		linkType = "_blank"
	}
	desc.LinkType = linkType
	// return data
	return desc
}

// add cover history for latest
//  Parameters:
//		q
//			sql get history
//		p
//			value for sql
//		agencyID
//			agency id
//		mediaType
//			media type id
//		linkType
//			url type
//		coverLink
//			cover url
//  Return:
//		Array CoverImage
func addCoverHistoryForLatest(q string, p []interface{}, agencyID, mediaType int64, linkType, coverLink string) ([]*CoverImage, error) {
	covers := make([]*CoverImage, 0)
	// open db
	db, err := pqx.Open()
	if err != nil {
		return covers, err
	}
	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return covers, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			agency   pqx.JSONRaw
			datetime time.Time
			path     sql.NullString
			filename sql.NullString
			ref      sql.NullString
		)
		rows.Scan(&datetime, &agency, &path, &filename, &ref)
		cover := &CoverImage{}
		// check valid path
		if path.Valid {
			cover.MediaPath, _ = b64.EncryptText(path.String + "/" + filename.String)
			cover.FilePath = path.String
			cover.Filename = filename.String
			cover.MediaDatetime = udt.DatetimeFormat(datetime, "datetime")
		}
		src := filepathx.JoinPath(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), path.String, filename.String)
		tbsrc := thumbnail.GetThumbName(src, "", "")
		// check thumbnail
		if filepathx.ValidateFile(tbsrc) == nil {
			fthumb := filepathx.Base(tbsrc)
			cover.ThumbnailMediaPath, _ = b64.EncryptText(path.String + "/" + fthumb)
			cover.ThumbnailFilePath = path.String
			cover.ThumbnailFilename = fthumb
		}
		// add cover link
		cover.CoverLink = coverLink
		if linkType == "" {
			linkType = "_blank"
		}
		cover.LinkType = linkType
		covers = append(covers, cover)
	}
	// return data
	return covers, nil
}

// func weather img by datetime
//  Parameters:
//		dataInput
//			WeatherImageInput
//  Return:
//		Array WeatherImageDataOutput
func GetWeatherImageDatetime(dataInput *WeatherImageInput) ([]*WeatherImageDataOutput, error) {

	if dataInput.Agency == 0 || dataInput.MediaType == 0 || dataInput.Date == "" {
		return nil, rest.NewError(422, "Invalid Parameter", nil)
	}

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := sqlSelectMediaImageDatetime
	p := []interface{}{dataInput.Agency, dataInput.MediaType, dataInput.Date + " 00:00:00", dataInput.Date + " 23:59:59"}

	//query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dateOutput := make([]*WeatherImageDataOutput, 0)
	for rows.Next() {
		var (
			agency   pqx.JSONRaw
			datetime time.Time
			path     sql.NullString
			filename sql.NullString
		)
		// scan data
		rows.Scan(&datetime, &agency, &path, &filename)

		// define struct for image output
		dd := &WeatherImageDataOutput{}
		dd.Agency = agency.JSON()
		dd.FilePath, _ = b64.EncryptText(path.String + "/" + filename.String)
		dd.Filename = filename.String
		dd.Datetime = udt.DatetimeFormat(datetime, "datetime")

		dateOutput = append(dateOutput, dd)
	}
	// return data
	return dateOutput, nil
}

// func get weather animation by agency id and mediatype id
//  Parameters:
//		agency
//			agency id
//		mediaType
//			media type id
//  Return:
//		WeatherAnimationOutput

func GetWeatherAnimation(agency, mediaType int64) (interface{}, error) {
	var rs interface{}
	var err error
	switch mediaType {
	case 141:
		rs, err = getAnimation("himawari-8/thailand/kochi/media/", "00Movie.mp4")
	case 143:
		rs, err = getAnimation("sst/south_china_sea/owi/media/", "SST.GIF")
	case 144:
		rs, err = getAnimation("sst/pacific/owi/media/", "SST.GIF")
	case 145:
		rs, err = getAnimation("sst/global/owi/media/", "SST.GIF")
	case 146:
		rs, err = getAnimation("sst/indian_northern/owi/media/", "SST.GIF")
	case 24:
		rs, err = getAnimation("wave_height/south_china_sea/owi/media/", "A_S_WAVE_latest.gif")
	case 25:
		rs, err = getAnimation("wave_height/indian_northern/owi/media/", "A_IND_WAVE_latest.gif")
	case 156:
		rs, err = getAnimation("wave_height/north_pacific_western/owi/media/", "A_N_WAVE_latest.gif")
	case 155:
		rs, err = getAnimation("wave_height/global/owi/media/", "A_G_WAVE_latest.gif")
	case 153:
		rs, err = getAnimation("ssh/global/aviso/media/", "A_GSSH_merge_latest.gif")
	case 70:
		rs, err = getAnimation("upper_wind_600m/asia/haii/media/", "upper_0.6km_large.mp4")
	case 71:
		rs, err = getAnimation("upper_wind_1500m/asia/haii/media/", "upper_1.5km_large.mp4")
	case 72:
		rs, err = getAnimation("upper_wind_5000m/asia/haii/media/", "upper_5.0km_large.mp4")
	default:
		return nil, rest.NewError(404, "Unknown media type ID", nil)
	}
	return rs, err
}

// func make struct for path animation
//  Parameters:
//		path
//			filepath
//		filename
//			filename
//  Return:
//		WeatherAnimationOutput
func getAnimation(path, filename string) (*WeatherAnimationOutput, error) {

	cover := &WeatherAnimationOutput{}
	//	src := filepathx.JoinPath(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), prefixAnimation, path, filename)
	src := filepathx.JoinPath("", prefixAnimation, path, filename)
	//	fmt.Println(src)
	cover.MediaPath, _ = b64.EncryptText(src)
	//	cover.MediaPath = src
	//	fmt.Println(mediaPath)
	cover.FilePath = filepathx.JoinPath(prefixAnimation, path)
	cover.Filename = filename
	return cover, nil
}

// func get animtion latest
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//  Return:
//		Array CoverImage
func getAnimationLatest(agencyID, mediaType int64, fname string) ([]*CoverImage, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql
	q := sqlSelectMediaImageLatest

	p := []interface{}{mediaType, agencyID}
	if fname != "" {
		q += " AND filename=$3"
		p = append(p, fname)
	}

	q += " ORDER BY media_type_id,agency_id, media_datetime DESC LIMIT 7"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*CoverImage, 0)
	for rows.Next() {
		var (
			agency   pqx.JSONRaw
			datetime time.Time
			path     sql.NullString
			filename sql.NullString
			ref      sql.NullString
		)
		// scan
		rows.Scan(&datetime, &agency, &path, &filename, &ref)
		// define cover image
		cover := &CoverImage{}
		cover.MediaPath, _ = b64.EncryptText(path.String + "/" + filename.String)
		cover.FilePath = path.String
		cover.Filename = filename.String
		data = append(data, cover)
	}

	return data, nil
}

// func get animation agency kochi
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//  Return:
//		CoverImage
func getAnimationKochiLatest(agencyID, mediaType int64, fname string) (*CoverImage, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql
	q := sqlSelectMediaImageLatest

	p := []interface{}{mediaType, agencyID}
	if fname != "" {
		q += " AND filename=$3"
		p = append(p, fname)
	}

	q += " ORDER BY media_type_id,agency_id, media_datetime"

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	cover := &CoverImage{}
	for rows.Next() {
		var (
			agency   pqx.JSONRaw
			datetime time.Time
			path     sql.NullString
			filename sql.NullString
			ref      sql.NullString
		)
		// scan
		rows.Scan(&datetime, &agency, &path, &filename, &ref)
		// add cover
		cover.MediaPath, _ = b64.EncryptText(path.String + "/" + filename.String)
		cover.FilePath = path.String
		cover.Filename = filename.String
	}
	// return data
	return cover, nil
}

// get weather history by start date and end date
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		sDate
//			start date
//		eDate
//			end date
//  Return:
//		Array WeatherHistoryDataOutput
func GetWeatherHistoryStartDateEndDate(agencyID, mediaType int64, sDate, eDate string) (interface{}, error) {

	var rs interface{}
	var err error
	switch mediaType {
	case 22:
		rs, err = getHistoryWeatherTMD(agencyID, mediaType, "", sDate, eDate)
	case 28:
		rs, err = getHistoryCommon(agencyID, mediaType, "", sDate, eDate)
	case 29:
		rs, err = getHistoryCommon(agencyID, mediaType, "", sDate, eDate)
	case 143:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 144:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 145:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 146:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 24:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 25:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 155:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	case 156:
		rs, err = getHistoryWeather6AMPM(agencyID, mediaType, "", sDate, eDate)
	default:
		return nil, rest.NewError(404, "Unknown media type ID", nil)
	}

	return rs, err
}

// get weather history by date
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		date
//			date media
//  Return:
//		Array WeatherHistoryDataOutput
func GetWeatherHistoryDate(agencyID, mediaType int64, date string) (interface{}, error) {

	var rs interface{}
	var err error
	// comment สำหรับให้ส่ง mediaType type id มาได้ทุกประเภท ไม่ต้อง fix ราย mediaType id เพราะ ทุก mediaType id เรียกใช้งาน func เดียวกัน
	//	switch mediaType {
	//	case 141:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 2:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 1:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 3:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 17:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 18:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 19:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	case 20:
	//		rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	//	default:
	//		return nil, rest.NewError(404, "Unknown media type ID", nil)
	//	}
	rs, err = getHistoryWeatherDate(agencyID, mediaType, "", date)
	if err != nil {
		return nil, rest.NewError(404, "Error Query", nil)
	}

	return rs, err
}

// func get weather history by month
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		year
//			year for get image
//		month
//			month for get image
//  Return:
//		Array WeatherHistoryDataOutput
func GetWeatherHistoryYearMonth(agencyID, mediaType int64, year, month string) (interface{}, error) {

	var rs interface{}
	var err error
	t1, err := time.Parse("20060102", year+month+"01")
	if err != nil {
		return nil, err
	}
	t2 := t1.AddDate(0, 1, -1)
	//	switch mediaType {
	//	case 153:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 142:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 140:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 14:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 160:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 161:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 162:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 163:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 149:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 150:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 151:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 152:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 173:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 62:
	//		//	storm url
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 27:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 141:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 157:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 158:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 159:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 143:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 5:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 6:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 4:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 155:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 17:
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 36: // modis ndvi terra
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 37: // modis ndvi aqua
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	case 182: // wind10m
	//		rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))
	//	default:
	//		return nil, rest.NewError(404, "Unknown media type ID", nil)
	//	}
	rs, err = getHistoryCommon(agencyID, mediaType, "", t1.Format("2006-01-02"), t2.Format("2006-01-02"))

	return rs, err
}

// get weather history by year
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		year
//			year for get image
//  Return:
//		Array WeatherHistoryDataOutput
func GetWeatherHistoryYear(agencyID, mediaType int64, year string) (interface{}, error) {

	var rs interface{}
	var err error
	y, _ := strconv.Atoi(year)
	switch mediaType {
	case 14:
		// sst 2w hii แผนภาพการเปลี่ยนแปลงของอุณหภูมิผิวน้ำทะเล ราย 2 สัปดาห์
		// comment เนื่องจาก ใช้ query บนข้อมูลออกมา 1 record
		//		rs, err = getHistoryWeatherYear(agencyID, mediaType, "", year+"-01-01", year+"-12-31", y, year+"-01-07")
		rs, err = getHistoryWeatherYear(agencyID, mediaType, "", year+"-01-01", year+"-12-31", y, year+"-12-31")
	default:
		return nil, rest.NewError(404, "Unknown media type ID", nil)
	}

	return rs, err
}

// sub function get weather history by year
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		year
//			year for get image
//		fname
//			filename
//		sDate
//			start date
//		eDate
//			end date
//		monFirst
//			month first
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryWeatherYear(agencyID, mediaType int64, fname, sDate, eDate string, year int, monFirst string) ([]*WeatherHistoryDataOutput, error) {

	if agencyID == 0 || mediaType == 0 || sDate == "" || eDate == "" || year == 0 {
		return nil, rest.NewError(422, "Null input parameterF", nil)
	}

	q := sqlHistoryWeatherYear1
	p := []interface{}{agencyID, mediaType, sDate, eDate, year, monFirst}

	return getHistoryOutputArrayWithSql(q, p)
}

// get weather history by date, agency id, mediatype
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//		date
//			date get image
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryWeatherDate(agencyID, mediaType int64, fname, date string) ([]*WeatherHistoryDataOutput, error) {

	if agencyID == 0 || mediaType == 0 || date == "" || date == "" {
		return nil, rest.NewError(422, "Null input parameterF", nil)
	}

	//	q := sqlHistoryWeatherDate
	//	q := "SELECT g.datetime,a.agency_name,m.media_path,m.filename,m.media_desc,$2 as id, mt.media_type_name, mt.media_subtype_name,mt.media_category " +
	//		"FROM (SELECT generate_series($3::timestamp,$4, '1 hour') AS datetime) g  " +
	//		"LEFT JOIN public.media m ON m.media_datetime >= $3 AND m.media_datetime <= $4 "
	//
	//	// storm ucl
	//	if mediaType == 62 {
	//		// เนื่องจาก นาทีที่ stamp ข้อมูลไม่ได้ ลงที่ 00 เลยต้องตัดเฉพาะ ชม.ที่เท่ากันมา join กันแทน
	//		q += " AND date(m.media_datetime) = date(g.datetime)  AND (EXTRACT(hour from m.media_datetime) = EXTRACT(hour from g.datetime)) "
	//	} else {
	//		q += "AND m.media_datetime=g.datetime "
	//	}
	//
	//	q += "AND m.media_type_id=$2 LEFT JOIN public.agency a ON m.agency_id=a.id AND a.id=$1 " +
	//		" LEFT JOIN public.lt_media_type mt ON m.media_type_id=mt.id " +
	//		"WHERE (m.deleted_at=to_timestamp(0) OR m.deleted_at IS NULL)"
	//
	//	p := []interface{}{agencyID, mediaType, date + " 00:00:00", date + " 23:59:59"}
	//
	//	if fname != "" {
	//		q += " AND filename=$5"
	//		p = append(p, fname)
	//	}

	// update by permporn (13/12/2019)
	q := "WITH Me AS (SELECT" +
		"(concat(m.media_datetime::date , ' ' , extract('hour' from m.media_datetime), ':00', ':00'))::timestamp as new_date," +
		"m.media_datetime," +
		"a.agency_name," +
		"m.media_path," +
		"m.filename," +
		"m.media_desc," +
		"mt.media_type_name," +
		"mt.media_subtype_name," +
		"mt.media_category " +
		" FROM " +
		"PUBLIC.media m	" +
		"LEFT JOIN PUBLIC.agency a ON m.agency_id = a.ID " +
		"LEFT JOIN PUBLIC.lt_media_type mt ON m.media_type_id = mt.ID " +
		" WHERE " +
		"(M.deleted_at = to_timestamp( 0 ) " +
		"OR m.deleted_at IS NULL) " +
		"AND m.media_type_id = $2 " +
		"AND m.agency_id = $1 " +
		"AND m.media_datetime >= $3" +
		"AND m.media_datetime <= $4" +
		") "

	q += " SELECT " +
		"g.datetime, " +
		"m.agency_name, " +
		"m.media_path, " +
		"m.filename, " +
		"m.media_desc, " +
		"$2 as id, " +
		"m.media_type_name, " +
		"m.media_subtype_name, " +
		"m.media_category " +
		"FROM " +
		"(SELECT generate_series($3::timestamp, $4::timestamp, '1 hour') AS datetime ) g "

	//  ภาพรายชั่วโมง  (storm ucl, digital-typhoon)
	if mediaType == 62 || (agencyID == 52 && mediaType == 141) {
		q += "LEFT JOIN Me m ON m.new_date = g.datetime "
	} else {
		q += "LEFT JOIN Me m ON m.media_datetime = g.datetime "
	}

	p := []interface{}{agencyID, mediaType, date + " 00:00:00", date + " 23:59:59"}

	//	fmt.Println(q)

	return getHistoryOutputArrayWithSql(q, p)
}

// get weather history time 6am
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//		sDate
//			start date
//		eDate
//			end date
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryWeather6AMPM(agencyID, mediaType int64, fname, sDate, eDate string) ([]*WeatherHistoryDataOutput, error) {

	if agencyID == 0 || mediaType == 0 || sDate == "" || eDate == "" {
		return nil, rest.NewError(422, "Null input parameterF", nil)
	}

	// sql
	q := sqlHistoryWeatherDaily
	p := []interface{}{agencyID, mediaType, sDate + " 06:00:00", eDate + " 23:59:59"}

	if fname != "" {
		q += " AND filename=$5"
		p = append(p, fname)
	}

	return getHistoryOutputArrayWithSql(q, p)
}

// func get weather history for agency TMD
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//		sDate
//			start date
//		eDate
//			end date
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryWeatherTMD(agencyID, mediaType int64, fname, sDate, eDate string) ([]*WeatherHistoryDataOutput, error) {

	if agencyID == 0 || mediaType == 0 || sDate == "" || eDate == "" {
		return nil, rest.NewError(422, "Null input parameterF", nil)
	}
	t, _ := time.Parse("2006-01-02", eDate)
	t = t.AddDate(0, 0, -1)
	q := sqlHistoryWeatherTMD
	p := []interface{}{agencyID, mediaType, sDate + " 07:00:00", t.Format("2006-01-02") + " 01:00:00"}

	if fname != "" {
		q += " AND filename=$5"
		p = append(p, fname)
	}

	return getHistoryOutputArrayWithSql(q, p)
}

// get history by common condition
//  Parameters:
//		agencyID
//			 agency id
//		mediaType
//			media type id
//		fname
//			filename
//		sDate
//			start date
//		eDate
//			end date
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryCommon(agencyID, mediaType int64, fname, sDate, eDate string) ([]*WeatherHistoryDataOutput, error) {

	q := sqlSelectMediaImageDatetime
	p := []interface{}{agencyID, mediaType, sDate + " 00:00:00", eDate + " 23:59:59"}

	if fname != "" {
		q += " AND filename=$5"
		p = append(p, fname)
	}

	return getHistoryOutputArrayWithSql(q, p)
}

// get history for output array with sql
//  Parameters:
//		q
//			sql get history image
//		p
//			value for sql
//  Return:
//		Array WeatherHistoryDataOutput
func getHistoryOutputArrayWithSql(q string, p []interface{}) ([]*WeatherHistoryDataOutput, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	fmt.Println(q)
	fmt.Println(p...)

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dateOutput := make([]*WeatherHistoryDataOutput, 0)
	for rows.Next() {
		var (
			agency      pqx.JSONRaw
			datetime    time.Time
			path        sql.NullString
			filename    sql.NullString
			mTypeId     int64
			mType       sql.NullString
			mSubType    sql.NullString
			mCategory   sql.NullString
			description sql.NullString
		)
		rows.Scan(&datetime, &agency, &path, &filename, &description, &mTypeId, &mType, &mSubType, &mCategory)

		// check valid path
		dd := &WeatherHistoryDataOutput{}
		if path.Valid {
			dd.MediaPath, _ = b64.EncryptText(path.String + "/" + filename.String)
			dd.FilePath = path.String
			dd.Filename = filename.String
		}
		dd.MediaDatetime = udt.DatetimeFormat(datetime, "datetime")
		dd.MediaTypeID = mTypeId
		dd.MediaType = mType.String
		dd.MediaSubType = mSubType.String
		dd.MediaCategory = mCategory.String
		dd.Description = description.String

		// check thumbnail from storage
		src := filepathx.JoinPath(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), path.String, filename.String)
		tbsrc := thumbnail.GetThumbName(src, "", "")
		if filepathx.ValidateFile(tbsrc) == nil {
			fthumb := filepathx.Base(tbsrc)
			dd.ThumbnailMediaPath, _ = b64.EncryptText(path.String + "/" + fthumb)
			dd.ThumbnailFilePath = path.String
			dd.ThumbnailFilename = fthumb

			dd.URLThumb, _ = b64.EncryptText(path.String + "/" + fthumb)
			dd.FilenameThumb = fthumb
			dd.FilepathThumb = path.String
		}

		dateOutput = append(dateOutput, dd)
	}

	return dateOutput, nil
}
