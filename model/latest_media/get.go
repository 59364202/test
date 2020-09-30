// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package latest_media is a model for cache.latest_media This table store latest_media information.
package latest_media

import (
	"database/sql"
	//"encoding/json"
	//"io/ioutil"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/thumbnail"

	"haii.or.th/api/thaiwater30/util/array"
	"haii.or.th/api/thaiwater30/util/b64"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
	"haii.or.th/api/thaiwater30/util/sqltime"
)

//	query sql
//	Paramters:
//		strSQL
//			sql query
//		args
//			parameter ที่ใช้ร่วมกับ strSQL
//	Return:
//		[]Struct_Media
func scanSql(strSQL string, args ...interface{}) ([]*Struct_Media, error) {
	//log.Println(strSQL)
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		obj *Struct_Media

		_filepath   sql.NullString
		_filename   sql.NullString
		_media_time sqltime.NullTime

		_filename_thumb string
	)
	data := make([]*Struct_Media, 0)

	fmt.Println(strSQL)

	row, err := db.Query(strSQL, args...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_filepath, &_filename, &_media_time)
		if err != nil {
			return nil, err
		}

		obj = &Struct_Media{}
		//		obj.Path, _ = model.GetCipher().EncryptText(filepath.Join(_filepath.String, _filename.String))
		obj.Path, _ = b64.EncryptText(filepath.Join(_filepath.String, _filename.String))
		obj.Filename = _filename.String
		obj.FilePath = _filepath.String
		obj.Datetime = _media_time.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))

		_filename_thumb = thumbnail.GetThumbName(_filename.String, "", "")

		//		/data/thaiwater/thaiwaterdata/data
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), _filepath.String, _filename_thumb)); err == nil {
			//			obj.PathThumb, _ = model.GetCipher().EncryptText(filepath.Join(_filepath.String, _filename_thumb))
			obj.PathThumb, _ = b64.EncryptText(filepath.Join(_filepath.String, _filename_thumb))
			obj.FilenameThumb = _filename_thumb[1:]
		}

		fmt.Println("_filename_thumb ", _filename_thumb)
		fmt.Println("th file ", filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), _filepath.String, _filename_thumb))
		fmt.Println(_filename_thumb[1:])
		fmt.Println(err)

		obj.Dt = _media_time.Time

		data = append(data, obj)
	}
	return data, nil
}

//GetStorm141 ภาพพายุ
//kochi -> digital typhoon -> reaseach labotary เรียงตามลำดับถ้าข้อมูลไม่อัพเดท
func GetStorm141() ([]*Struct_Media, error) {
	s, err := GetLatestMedia(50, 141)
	if !ValidStorm(s) || err != nil {
		log.Println(err)
		s, err = GetLatestMedia(52, 141)
		if !ValidStorm(s) || err != nil {
			log.Println(err)
			s, err = GetLatestMedia(51, 141)
			if err != nil {
				return nil, err
			}
		}
	}
	return s, nil
}

//ValidStorm เช็คข้อมูลว่าอัพเดทไม่เกิน 6 ชม
//return true ถ้าไม่เกิน 6ชม
func ValidStorm(s []*Struct_Media) bool {
	if len(s) != 1 {
		return false
	}
	storm := s[0]

	// ถ้เวลาของภาพ ห่างจากปัจจุบัน เกิน x ชม. ให้ถือว่าเป็นภาพไม่อัพเดท
	diffDuration := time.Since(storm.Dt)
	if diffDuration.Hours() >= 6 {
		return false
	}

	return true
}

//	Get data from cache.latest_media
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//		media_type_id
//			รหัสชนิดข้อมูลสื่อ
//	Return:
//		[]Struct_Media
func GetLatestMedia(agency_id, media_type_id int64) ([]*Struct_Media, error) {
	return scanSql(SQL_Select, agency_id, media_type_id)
}

//	get รูปภาพคาดการณ์ฝน ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreRain() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreRain)
}

//	get รูปภาพคาดการณ์ฝน thailand ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreRainTH() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 17)
}

//	get รูปภาพคาดการณ์ฝน เอเชีย ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreRainAsia() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 18)
}

//	get รูปภาพคาดการณ์ฝน Southeast Asia ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreRainSea() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 19)
}

//	get รูปภาพคาดการณ์ฝน ลุ่มน้ำ ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreRainBasin() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 20)
}

//	get รูปภาพคาดการณ์คลื่น ล่าสุด
//	Return:
//		[]Struct_Media
func GetPreWave() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreWave)
}

//	get รูปภาพสมดุลน้ำ
//	Return:
//		[]Struct_Media
func GetWaterBalance() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 51)
}

//	get รูปภาพคาดการณ์สมดุลน้ำ รายสัปดาห์
//	Return:
//		[]Struct_Media
func GetPreWaterBalanceWeekly() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 57)
}

//	get รูปภาพคาดการณ์สมดุลน้ำ รายเดือน
//	Return:
//		[]Struct_Media
func GetPreWaterBalanceMonthly() ([]*Struct_Media, error) {
	return GetLatestMedia(9, 52)
}

//	Get Rainforcase_Province data from cache.latest_media
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//		media_type_id
//			รหัสชนิดข้อมูลสื่อ
//	Return:
//		[]Struct_Media
func GetLatestMediaRainforcaseProvince(agency_id, media_type_id int64) ([]*Struct_Media, error) {
	return scanSql(SQL_Select_Rainforcase_Province, agency_id, media_type_id)
}

//	get รูปภาพเรดาร์ ล่าสุด
//	Parameters:
//		agency_id
//			รหัสหน่วยงาน
//	Return:
//		[]Struct_Media
func GetRadar() ([]*Struct_Radar, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var sql, p = Gen_SQL_Radar()

	row, err := db.Query(sql, p...)
	if err != nil {
		return nil, err
	}

	var (
		obj       *Struct_Radar
		data      []*Struct_Radar = make([]*Struct_Radar, 0)
		dataIndex []int           = make([]int, 0)
	)
	//	var _setting = make([]map[string]interface{}, 0)
	//	json.Unmarshal(setting.GetSystemSettingJSON("Frontend.analyst.Radar.RadarTypeOrder"), &_setting)

	_setting, err := uSetting.New_Struct_RadarTypeOrder()
	if err != nil {
		return nil, err
	}

	// ตรึง radar ตาม setting พร้อมใส่ _error เป็น deault
	for index, v := range _setting {
		v.DefaultErrorFile()

		obj = &Struct_Radar{}
		obj.RadarName = v.Radar_name
		obj.RadarType = v.Radar_type
		obj.Agency = v.Agency
		obj.Timezone = v.Timezone
		obj.FilePath = v.FilePath
		obj.Filename = v.FileName
		obj.MediaPath = v.MediaPath
		obj.FilenameThumb = v.FileName
		obj.MediaPathThumb = v.MediaPath

		data = append(data, obj)
		dataIndex = append(dataIndex, index)
	}
	for row.Next() {
		var (
			_media_datetime time.Time
			_media_path     string
			_filename       string
			_short_filename string
		)
		err = row.Scan(&_media_datetime, &_media_path, &_filename, &_short_filename)
		if err != nil {
			return nil, err
		}

		//		if _short_filename[len(_short_filename)-1:] == "_" {
		//			strings.Replace(_short_filename, "_", "", 1)
		//		}
	L:
		for ind, v := range dataIndex {
			obj := data[v]
			if obj == nil {
				continue
			}
			if obj.RadarType != _short_filename {
				continue
			}
			obj.MediaDatetime = _media_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
			obj.FilePath = _media_path
			obj.Filename = _filename
			obj.MediaPath, _ = b64.EncryptText(filepath.Join(_media_path, _filename))

			// ถ้เวลาของภาพ ห่างจากปัจจุบัน เกิน x ชม. ให้  thumbnail เป็น _error
			// แต่ถ้าไม่เกิน x ชม. ให้ใช้ thumbnail เป็น thumbnail ของมันเอง
			diffDuration := time.Since(_media_datetime)
			if diffDuration.Hours() <= datatype.MakeFloat(setting.GetSystemSettingInt("Frontend.analyst.Radar.DiffHour")) {
				filenameThumb := thumbnail.GetThumbName(_filename, "", "")

				full_path := filepath.Join(_media_path, filenameThumb)

				if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
					obj.MediaPathThumb, _ = b64.EncryptText(filepath.Join(_media_path, filenameThumb))
					obj.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
				}

			}

			dataIndex = array.RemoveIntArrayByIndex(dataIndex, ind)
			break L

		}

	}

	return data, nil
}

type Struct_RadarTypeOrder struct {
	Radar_type      string `json:"radar_type"`      // example:`cri240` ประเภทเรดาร์
	Radar_name      string `json:"radar_name"`      // example:`เรดาห์เชียงราย รัศมี 240 กม.` ชื่อเรดาร์
	Radar_frequency int    `json:"radar_frequency"` // example:`60`
	Agency          string `json:"agency"`          // example:`tmd` หน่วยงาน
	Timezone        string `json:"timezone"`        // example:`GMT` zone เวลา
	Band            string `json:"band"`            // example:`sband`

	FileName  string // error file name
	FilePath  string // error file path
	MediaPath string // error file path (encrypt)
}

//	get รูปภาพเรดาร์ ล่าสุด รายจังหวัด
//	Parameters:
//		province_id
//			รหัสจังหวัด
//	Return:
//		[]Struct_Media
// 	Author:
//		permporn@haii.or.th
func GetRadarByProvince(inputData []string) ([]*Struct_Radar, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var sql, p = Gen_SQL_Radar_By_Code(inputData)

	row, err := db.Query(sql, p...)
	if err != nil {
		return nil, err
	}

	var (
		obj       *Struct_Radar
		data      []*Struct_Radar = make([]*Struct_Radar, 0)
		dataIndex []int           = make([]int, 0)
	)

	_setting, err := uSetting.New_Struct_RadarTypeOrder()
	if err != nil {
		return nil, err
	}

	//---test get datas systemsetting :"Frontend.analyst.Radar.RadarTypeOrder"---//
	//	_setting := make([]*Struct_RadarTypeOrder, 0)
	//
	//	// read radar.json config
	//	//radar := "/home/cim/go_local/src/haii.or.th/api/thaiwater30/service/provinces/radar.json" // server
	//	radar2 :=  "./src/haii.or.th/api/thaiwater30/service/provinces/radar2.json"  // local
	//	raw2, err := ioutil.ReadFile(radar2)
	//	if err != nil {
	//		fmt.Println("read station.json err : %s", err)
	//		//return nil
	//	}
	//	if err = json.Unmarshal(raw2, &_setting); err != nil {
	//			fmt.Println("Unmarshal err : %s", err)
	//			fmt.Println("stations all: %v ", len(_setting))
	//			//return nil
	//	}

	index := 0
	for _, v := range _setting {
		v.DefaultErrorFile()
		obj = &Struct_Radar{}
		if len(inputData) > 0 {
			for _, c := range inputData {
				if v.Radar_type == c {
					obj.RadarName = v.Radar_name
					obj.RadarType = v.Radar_type
					obj.Agency = v.Agency
					obj.Timezone = v.Timezone
					obj.FilePath = v.FilePath
					obj.Filename = v.FileName
					obj.MediaPath = v.MediaPath
					obj.FilenameThumb = v.FileName
					obj.MediaPathThumb = v.MediaPath

					data = append(data, obj)
					dataIndex = append(dataIndex, index)
					index++
				}
			}
		} else {
			obj.RadarName = v.Radar_name
			obj.RadarType = v.Radar_type
			obj.Agency = v.Agency
			obj.Timezone = v.Timezone
			obj.FilePath = v.FilePath
			obj.Filename = v.FileName
			obj.MediaPath = v.MediaPath
			obj.FilenameThumb = v.FileName
			obj.MediaPathThumb = v.MediaPath

			data = append(data, obj)
			dataIndex = append(dataIndex, index)
			index++
		}
	}

	for row.Next() {
		var (
			_media_datetime time.Time
			_media_path     string
			_filename       string
			_short_filename string
		)
		err = row.Scan(&_media_datetime, &_media_path, &_filename, &_short_filename)
		if err != nil {
			return nil, err
		}
	L:
		for ind, v := range dataIndex {
			obj := data[v]
			if obj == nil {
				continue
			}
			if obj.RadarType != _short_filename {
				continue
			}
			obj.MediaDatetime = _media_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
			obj.FilePath = _media_path
			obj.Filename = _filename
			obj.MediaPath, _ = b64.EncryptText(filepath.Join(_media_path, _filename))

			// ถ้เวลาของภาพ ห่างจากปัจจุบัน เกิน x ชม. ให้  thumbnail เป็น _error
			// แต่ถ้าไม่เกิน x ชม. ให้ใช้ thumbnail เป็น thumbnail ของมันเอง
			diffDuration := time.Since(_media_datetime)
			if diffDuration.Hours() <= datatype.MakeFloat(setting.GetSystemSettingInt("Frontend.analyst.Radar.DiffHour")) {
				filenameThumb := thumbnail.GetThumbName(_filename, "", "")

				full_path := filepath.Join(_media_path, filenameThumb)

				if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
					obj.MediaPathThumb, _ = b64.EncryptText(filepath.Join(_media_path, filenameThumb))
					obj.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
				}
			}
			dataIndex = array.RemoveIntArrayByIndex(dataIndex, ind)
			break L
		}
	}
	return data, nil
}
