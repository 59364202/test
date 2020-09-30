// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package water_map is a model for public.dam_daily and public.tele_waterlevel table. This table store tele waterlevel data.
package water_map

import (
	"database/sql"
	"haii.or.th/api/util/pqx"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"strconv"
	"time"
)

// get dam for water map
//  Parameters:
//		None
//  Return:
//		WaterMapOutput
func GetWaterMap() (*WaterMapOutput, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := getDamDaily
	p := []interface{}{}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataDam := make([]*DamDailyOutput, 0)
	//	listBasinDam := make([]*BasinDam, 0)
	//	basinDam := &BasinDam{}

	//	var damRow string
	for rows.Next() {
		var (
			damName          pqx.JSONRaw
			agencyName       pqx.JSONRaw
			basinName        pqx.JSONRaw
			subbasinName     pqx.JSONRaw
			areaName         pqx.JSONRaw
			provinceName     pqx.JSONRaw
			amphoeName       pqx.JSONRaw
			tumbonName       pqx.JSONRaw
			lat              sql.NullFloat64
			long             sql.NullFloat64
			damOldCode       sql.NullString
			maxWaterlevel    sql.NullFloat64
			normalWaterlevel sql.NullFloat64
			minWaterlevel    sql.NullFloat64
			damDate          sql.NullString
			damInflow        sql.NullFloat64
			damReleased      sql.NullFloat64
			damStorage       sql.NullFloat64
		)
		rows.Scan(&damName, &agencyName, &basinName, &subbasinName, &areaName, &provinceName, &amphoeName, &tumbonName, &lat, &long, &damOldCode, &maxWaterlevel,
			&normalWaterlevel, &minWaterlevel, &damDate, &damInflow, &damReleased, &damStorage)
		damRow := &DamDailyOutput{}
		// define data output by row
		damRow.DamName = damName.JSON()
		damRow.AgencyName = agencyName.JSON()
		damRow.BasinName = basinName.JSON()
		damRow.SubbasinName = subbasinName.JSON()
		damRow.AreaName = areaName.JSON()
		damRow.ProvinceName = provinceName.JSON()
		damRow.AmphoeName = amphoeName.JSON()
		damRow.TumbonName = tumbonName.JSON()
		damRow.Lat = ValidData(lat.Valid, strconv.FormatFloat(lat.Float64, 'f', 6, 64))
		damRow.Long = ValidData(long.Valid, strconv.FormatFloat(long.Float64, 'f', 6, 64))
		damRow.DamOldCode = damOldCode.String
		damRow.MaxWaterlevel = ValidData(maxWaterlevel.Valid, maxWaterlevel.Float64)
		damRow.NormalWaterlevel = ValidData(normalWaterlevel.Valid, normalWaterlevel.Float64)
		damRow.MinWaterlevel = ValidData(minWaterlevel.Valid, minWaterlevel.Float64)
		dt,_ := time.Parse(time.RFC3339,damDate.String)
		
		damRow.DamDate = udt.DatetimeFormat(dt,"date")
		damRow.DamInflow = ValidData(damInflow.Valid, damInflow.Float64)
		damRow.DamReleased = ValidData(damReleased.Valid, damReleased.Float64)
		damRow.DamStorage = ValidData(damStorage.Valid, damStorage.Float64)
		dataDam = append(dataDam, damRow)
	}

	//	if damRow != "" {
	//		basinDam.DamDaily = dataDam
	//		listBasinDam = append(listBasinDam, basinDam)
	//	}
	
	// data from tele waterlevel 
	q = getTeleWaterLevel
	p = []interface{}{}
	rows, err = db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataWaterlevel := make([]*TeleWaterlevel, 0)
	//	listBasinWaterlevel := make([]*BasinWaterlevel, 0)
	//	basinWaterlevel := &BasinWaterlevel{}

	//	var waterlevelRow string
	for rows.Next() {
		var (
			teleStationName    pqx.JSONRaw
			agencyName         pqx.JSONRaw
			basinName          pqx.JSONRaw
			subbasinName       pqx.JSONRaw
			areaName           pqx.JSONRaw
			provinceName       pqx.JSONRaw
			amphoeName         pqx.JSONRaw
			tumbonName         pqx.JSONRaw
			lat                sql.NullFloat64
			long               sql.NullFloat64
			teleStationOldcode sql.NullString
			rightBank          sql.NullFloat64
			leftBank           sql.NullFloat64
			waterlevelDatetime time.Time
			waterlevelM        sql.NullFloat64
			waterlevelMSL      sql.NullFloat64
//			waterlevelIn       sql.NullFloat64
//			waterlevelOut      sql.NullFloat64
//			waterlevelOut2     sql.NullFloat64
			flowRate           sql.NullFloat64
			Discharge          sql.NullFloat64
		)
		rows.Scan(&teleStationName, &agencyName, &basinName, &subbasinName, &areaName, &provinceName, &amphoeName, &tumbonName, &lat, &long, &teleStationOldcode,
			&rightBank, &leftBank, &waterlevelDatetime, &waterlevelM, &waterlevelMSL, &flowRate, &Discharge)

		dataWaterlevelRow := &TeleWaterlevel{}
		// waterlevel by row
		dataWaterlevelRow.TeleStationName = teleStationName.JSON()
		dataWaterlevelRow.AgencyName = agencyName.JSON()
		dataWaterlevelRow.BasinName = basinName.JSON()
		dataWaterlevelRow.SubbasinName = subbasinName.JSON()
		dataWaterlevelRow.AreaName = areaName.JSON()
		dataWaterlevelRow.ProvinceName = provinceName.JSON()
		dataWaterlevelRow.AmphoeName = amphoeName.JSON()
		dataWaterlevelRow.TumbonName = tumbonName.JSON()
		dataWaterlevelRow.Lat = ValidData(lat.Valid, strconv.FormatFloat(lat.Float64, 'f', 6, 64))
		dataWaterlevelRow.Long = ValidData(long.Valid, strconv.FormatFloat(long.Float64, 'f', 6, 64))
		dataWaterlevelRow.TeleStationOldCode = teleStationOldcode.String
		dataWaterlevelRow.RightBank = ValidData(rightBank.Valid, rightBank.Float64)
		dataWaterlevelRow.LeftBank = ValidData(leftBank.Valid, leftBank.Float64)
		dataWaterlevelRow.WaterlevelDatetime = udt.DatetimeFormat(waterlevelDatetime,"datetime")
		dataWaterlevelRow.WaterlevelM = ValidData(waterlevelM.Valid, waterlevelM.Float64)
		dataWaterlevelRow.WaterlevelMSL = ValidData(waterlevelMSL.Valid, waterlevelMSL.Float64)
//		dataWaterlevelRow.WaterlevelIn = ValidData(waterlevelIn.Valid, waterlevelIn.Float64)
//		dataWaterlevelRow.WaterlevelOut = ValidData(waterlevelOut.Valid, waterlevelOut.Float64)
//		dataWaterlevelRow.WaterlevelOut2 = ValidData(waterlevelOut2.Valid, waterlevelOut2.Float64)
		dataWaterlevelRow.FlowRate = ValidData(flowRate.Valid, flowRate.Float64)
		dataWaterlevelRow.Discharge = ValidData(Discharge.Valid, Discharge.Float64)
		dataWaterlevel = append(dataWaterlevel, dataWaterlevelRow)
	}

	//	if waterlevelRow != "" {
	//		basinWaterlevel.TeleWL = dataWaterlevel
	//		listBasinWaterlevel = append(listBasinWaterlevel, basinWaterlevel)
	//	}

	data := &WaterMapOutput{}
	data.DamDaily = dataDam
	data.TeleWL = dataWaterlevel
	// return data for water map
	return data, nil
}

// check valid data by field
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}
