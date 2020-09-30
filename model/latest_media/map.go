// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package latest_media is a model for cache.latest_media This table store latest_media information.
package latest_media

import (
	//	"database/sql"
	//	"log"
	"strconv"
	"strings"
	//	"time"

	"haii.or.th/api/util/pqx"
)

type CfgMediaIndex struct {
	Function string
	Param1   interface{}
	Param2   interface{}

	SetIndex func(d *Struct_UpdateCacheMedia)
}

// config for create index when update latest_media
var mapCfgMediaIndex = map[int64]map[int64]*CfgMediaIndex{
	//	4: {
	//		130: &CfgMediaIndex{},
	//	},
	6: {
		//		22: &CfgMediaIndex{},
		23: &CfgMediaIndex{},
		24: &CfgMediaIndex{},
		25: &CfgMediaIndex{},
		26: &CfgMediaIndex{},
		27: &CfgMediaIndex{},
	},

	//	haii
	9: {
		//	Contour Images of Humidity
		1: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		//	Contour Images of Temperature
		2: &CfgMediaIndex{
			Function: "substr",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		//	Contour Images of Pressure
		3: &CfgMediaIndex{
			Function: "substr",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		//	Rain Accumulation Forecasts by WRF-ROMS Model Asia (27km x 27km) 7 days
		5: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				l := len(d.Filename.String) - 10
				last10 := d.Filename.String[l:]
				d.AddIndex(last10)
			},
		},
		//	Rain Accumulation Forecasts by WRF-ROMS Model Southeast Asia (9km x 9km) 7 days
		6: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				l := len(d.Filename.String) - 10
				last10 := d.Filename.String[l:]
				d.AddIndex(last10)
			},
		},
		7: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		8: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		9: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		//	WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level
		10: &CfgMediaIndex{
			Function: "substr",
			Param2:   10,
		},
		//	WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level
		11: &CfgMediaIndex{
			Function: "substr",
			Param2:   10,
		},

		//	WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level
		12: &CfgMediaIndex{
			Function: "substr",
			Param2:   10,
		},
		//	swan
		13: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				if t.Hour() != 7 && t.Hour() != 19 { // เอาแค่ภาพตอน 7, 19 โมง
					return
				}
				wd := t.Weekday()
				d.AddIndex(wd.String() + "_" + t.Format("15"))
				//				//	หา ว่าภาพนี้เป็นภาพที่เท่าไหร่ นับจาก วันปัจจุบัน
				//				now := time.Now()
				//				tnoTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
				//				nnoTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
				//				duration := tnoTime.Sub(nnoTime)
				//				dateDiff := int(duration.Hours() / 24)
				//				if dateDiff < 0 {
				//					return
				//				}
				//				d.AddIndex(strconv.Itoa(dateDiff) + "_" + t.Format("15")) // 0_02, 5_18
			},
		},
		//	sst_w
		14: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				_, w := t.ISOWeek()
				d.AddIndex(strconv.Itoa(w))
			},
		},
		//	Pressure 0.6km
		15: &CfgMediaIndex{
			Function: "substr",
			Param2:   5,
		},
		//	Pressure 1.5km
		16: &CfgMediaIndex{
			Function: "substr",
			Param2:   5,
		},
		//	WRF-ROMS Model, Thailand (3x3)
		17: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				if len(d.Filename.String) != 13 {
					return
				}
				d.AddIndex(d.Filename.String)
			},
		},
		//	WRF-ROMS Model, Asia (27km x 27km)
		18: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				if len(d.Filename.String) != 13 {
					return
				}
				d.AddIndex(d.Filename.String)
			},
		},
		//	WRF-ROMS Model, Southeast Asia (9km x 9km)
		19: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				//				l := len(d.Filename.String) - 10
				//				last10 := d.Filename.String[l:]
				//				index := last10[0:2]
				//				d.AddIndex(index)
				if len(d.Filename.String) != 13 {
					return
				}
				//				fnameDay := d.Filename.String[7:9]
				//				int_fnameDay, err := strconv.Atoi(fnameDay)
				//				if err != nil {
				//					return
				//				}
				//				if int_fnameDay < 4 { // d02_day04.jpg ขึ้นไป
				//					return
				//				}
				d.AddIndex(d.Filename.String)
			},
		},
		//	WRF-ROMS Model, Thailand Basin (3km x 3km)
		20: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				if len(d.Filename.String) != 19 {
					return
				}
				d.AddIndex(d.Filename.String)
			},
		},
		//	ภาพสมดุลน้ำ นอกเขตชลประทาน
		51: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(strings.Replace(d.Filename.String, t.Format("2006-01-02"), "", 1))
			},
		},
		//	ภาพสมดุลน้ำ นอกเขตชลประทาน รายเดือน
		52: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(strings.Replace(d.Filename.String, t.Format("2006_01"), "", 1))
			},
		},
		//	haii-swat-w-forecast
		57: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(strings.Replace(d.Filename.String, t.Format("2006-01-02"), "", 1))
			},
		},
		140: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		142: &CfgMediaIndex{
			Function: "substr",
			Param2:   8,
		},
		// SAT GSMaP (10km x 10km)
		157: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		// SAT PERSIANN (4km x 4km)
		158: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		// SAT GSMaP (24km x 24km)
		159: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(t.Format("15"))
			},
		},
		// swan latest
		170: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				switch d.Filename.String {
				case "wave_hr012.png", "wave_hr024.png":
					//					log.Println("set index 1")
					d.AddIndex("1")
				case "wave_hr036.png", "wave_hr048.png":
					d.AddIndex("2")
				case "wave_hr060.png", "wave_hr072.png":
					d.AddIndex("3")
				case "wave_hr084.png", "wave_hr096.png":
					d.AddIndex("4")
				case "wave_hr108.png", "wave_hr120.png":
					d.AddIndex("5")
				case "wave_hr132.png", "wave_hr144.png":
					d.AddIndex("6")
				case "wave_hr156.png", "wave_hr168.png":
					d.AddIndex("7")
				}
			},
		},
		//	WRF-ROM (thaiGeo)24hrs rain forecast 7day
		181: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				d.AddIndex(d.Filename.String)
			},
		},
	},
	// BMA, agency_id=10
	10: {
		// Radar nongjok, nongkham ,media_type_id = 30
		30: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				switch d.MediaDesc.String {
				case "bma-radar-nongjok":
					// njk = nongjok
					d.AddIndex("njk")
				case "bma-radar-nongkham":
					// nkm = nongkham
					d.AddIndex("nkm")
				}
			},
		},
	},
	11: {
	//		4: &CfgMediaIndex{},
	//		5: &CfgMediaIndex{},
	//		6: &CfgMediaIndex{},
	//		35: &CfgMediaIndex{},
	//		36: &CfgMediaIndex{},
	//		37: &CfgMediaIndex{},
	//		38: &CfgMediaIndex{},
	//		39: &CfgMediaIndex{},
	//		40: &CfgMediaIndex{},
	//		41: &CfgMediaIndex{},
	//		42: &CfgMediaIndex{},
	},
	13: {
		//		4: &CfgMediaIndex{},
		//		22: &CfgMediaIndex{},
		//	tmd-windmap850
		28: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(strings.Replace(d.Filename.String, t.Format("2006-01-02_15"), "", 1))
			},
		},
		//	tmd-windmap925
		29: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				t := pqx.NullStringToTime(d.MediaDatetime)
				d.AddIndex(strings.Replace(d.Filename.String, t.Format("2006-01-02_15"), "", 1))
			},
		},
		//  tmd-radar
		30: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		//	Himawarii - IR
		46: &CfgMediaIndex{
			Function: "substr",
			Param2:   2,
		},
		//	Himawarii - WV
		47: &CfgMediaIndex{
			Function: "substr",
			Param2:   2,
		},
		//		48: &CfgMediaIndex{},
		//	Himawarii - IR4
		49: &CfgMediaIndex{
			Function: "substr",
			Param2:   2,
		},
		//	Himawarii - S4
		63: &CfgMediaIndex{
			Function: "substr",
			Param2:   2,
		},
	},
	19: { //agency = drraa
		//media_type_id = 30
		30: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				// d.Filename.String = CAPPI240@chatu@171024063000.png
				// d.media_path.String = product/radar/history/drraa/cband/rongkwang/2018/05/04
				z := strings.Split(d.MediaPath.String, "/")
				d.AddIndex(z[5])
			},
		},
	},
	41: {
		//	storm_map
		62: &CfgMediaIndex{
			Function: "",
		},
	},
	50: {
		//	ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi
		141: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				if d.Filename.String == "00Latest.jpg" {
					d.AddIndex(d.Filename.String)
				}
			},
		},
	},
	51: {
		//	us_naval
		141: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				d.AddIndex("1")
			},
		},
	},
	52: {
		// digital_typhoon
		141: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				d.AddIndex("1")
			},
		},
	},
	55: {
		//		36: &CfgMediaIndex{},
		//		37: &CfgMediaIndex{},
		145: &CfgMediaIndex{
			Function: "custom",
			SetIndex: func(d *Struct_UpdateCacheMedia) {
				l := len(d.Filename.String) - 3
				d.AddIndex(d.Filename.String[l:])
			},
		},
		149: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		150: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		151: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		152: &CfgMediaIndex{
			Function: "substr",
			Param2:   6,
		},
		//		160: &CfgMediaIndex{},
		//		161: &CfgMediaIndex{},
		//		162: &CfgMediaIndex{},
		//		163: &CfgMediaIndex{},
	},
	56: {
	//	143: &CfgMediaIndex{},
	//	155: &CfgMediaIndex{},
	},
	57: {
		153: &CfgMediaIndex{
			Function: "substr",
			Param2:   3,
		},
	},

	// generate latest image (rain24, waterlevel, dam) for mobile
	64: {
		176: &CfgMediaIndex{},
		177: &CfgMediaIndex{},
		178: &CfgMediaIndex{},
	},
}

// Get config media index from mapCfgMediaIndex
//
//	Parameters:
//		media_type_id
//			media_type_id
//	Return:
//		*CfgMediaIndex
//		if no media_type_id in mapCfgMediaIndex will be nil
//
func GetCfg(agency_id, media_type_id int64) *CfgMediaIndex {
	s, ok := mapCfgMediaIndex[agency_id]
	if !ok {
		return nil
	}
	c, ok := s[media_type_id]
	if !ok {
		return nil
	}
	return c
}

// Convert interface to int
//
//	Parameters:
//		v
//			value type interface
//	Return:
//		int(v)
//		if nil will be 0
//
func GetInt(v interface{}) int {
	if v != nil {
		return v.(int)
	}
	return 0
}

// Convert interface to string
//
//	Parameters:
//		v
//			value type interface
//	Return:
//		string(v)
//		if nil will be ""
//
func GetString(v interface{}) string {
	if v != nil {
		return v.(string)
	}
	return ""
}
