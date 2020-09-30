// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package sea_station is a model for public.sea_waterlevel table. This table store sea waterlevel data.
package sea_station

import (
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

// get sea station order by id , select by station type
//  Parameters:
//		stationType
//			station observe or forecast
//  Return:
//		Array Agency
func SeaStationByAgency(stationType string) ([]*Agency, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql get sea waterlevel by station and dates
	var q string
	if stationType == "observe" {
		q = sqlSeaStationByAgency
	} else {
		q = sqlSeaForecastStationByAgency
	}

	p := []interface{}{}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// define value
	var (
		aid        sql.NullInt64
		aName      pqx.JSONRaw
		aShortName pqx.JSONRaw
		mid        sql.NullInt64
		lat        sql.NullFloat64
		long       sql.NullFloat64
		oldcode    sql.NullString
		sName      pqx.JSONRaw
	)

	// data observe
	agency := make([]*Agency, 0)
	station := make([]*Station, 0)
	aData := &Agency{}
	var IdTemp int64
	for rows.Next() {
		// scan data
		rows.Scan(&aid, &aName, &aShortName, &mid, &lat, &long, &oldcode, &sName)
		// check duplicate data for add data to array
		if IdTemp != 0 {
			if IdTemp != aid.Int64 {
				// new array index
				agency = append(agency, aData)
				station = make([]*Station, 0)

				IdTemp = aid.Int64
				aData = &Agency{}
				aData.ID = aid.Int64
				aData.Name = aName.JSON()
				aData.ShortName = aShortName.JSON()
				s := &Station{}
				s.ID = mid.Int64
				s.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
				s.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
				s.Oldcode = oldcode.String
				s.Name = sName.JSON()
				station = append(station, s)
				aData.Station = station
			} else {
				// add data to array
				s := &Station{}
				s.ID = mid.Int64
				s.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
				s.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
				s.Oldcode = oldcode.String
				s.Name = sName.JSON()
				station = append(station, s)
				aData.Station = station
			}
		} else {
			// first row of data
			IdTemp = aid.Int64
			aData.ID = aid.Int64
			aData.Name = aName.JSON()
			aData.ShortName = aShortName.JSON()
			s := &Station{}
			s.ID = mid.Int64
			s.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
			s.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
			s.Oldcode = oldcode.String
			s.Name = sName.JSON()
			station = append(station, s)
			aData.Station = station
		}
	}

	if aData.ID != 0 {
		// add data
		agency = append(agency, aData)
	}

	return agency, nil
}
