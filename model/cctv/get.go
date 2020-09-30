// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package cctv is a model for public.cctv table. This table store cctv information.
package cctv

import (
	"database/sql"
	//	"fmt"
	"encoding/json"
	"strconv"

	model_basin "haii.or.th/api/thaiwater30/model/basin"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
	"haii.or.th/api/util/pqx"
)

// Get infomation cctv return cctv details
//  Parameters:
//		None
//  Return:
//		Array information cctv
func GetDetailsCCTV() ([]*cctvOutput, error) {

	// open db
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	// get sql
	q := getCCTV
	p := []interface{}{}

	//	fmt.Println(q)

	// query data
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// make array output
	data := make([]*cctvOutput, 0)

	for rows.Next() {
		var (
			id                int64
			damID             sql.NullString
			tele_station_id   sql.NullString
			basin_id          sql.NullInt64
			lat               sql.NullFloat64
			long              sql.NullFloat64
			title             sql.NullString
			description       sql.NullString
			media_type        sql.NullString
			base_url          sql.NullString
			filename          sql.NullString
			tele_station_name sql.NullString
			basin_name        sql.NullString
			cctv_flash        sql.NullString
			cctv_quicktime    sql.NullString
			is_active		  sql.NullBool

			gecode_id     int64
			geocode       sql.NullString
			tumbon_code   sql.NullString
			tumbon_name   sql.NullString
			amphoe_code   sql.NullString
			amphoe_name   sql.NullString
			province_code sql.NullString
			province_name sql.NullString
			area_code     sql.NullString
			area_name     sql.NullString

			_basin_code sql.NullInt64
			_basin_name sql.NullString

			_subbasin_id sql.NullInt64
		)
		dataRow := &cctvOutput{}
		// scan data
		rows.Scan(&id, &tele_station_name, &basin_name, &gecode_id, &geocode, &tumbon_code, &tumbon_name, &amphoe_code, &amphoe_name,
			&province_code, &province_name, &area_code, &area_name, &damID, &tele_station_id, &basin_id, &lat, &long, &title,
			&description, &media_type, &base_url, &filename, &cctv_flash, &cctv_quicktime, &_basin_code, &_basin_name, &_subbasin_id, &is_active)

		if !tele_station_name.Valid || tele_station_name.String == "" {
			tele_station_name.String = "{}"
		}
		if !basin_name.Valid || basin_name.String == "" {
			basin_name.String = "{}"
		}

		// get data cctv by row
		dataRow.Id = id
		dataRow.DamID = damID.String
		dataRow.TeleStationName = json.RawMessage(tele_station_name.String)
		dataRow.BasinName = json.RawMessage(basin_name.String)
		dataRow.Lat = strconv.FormatFloat(lat.Float64, 'f', 6, 64)
		dataRow.Long = strconv.FormatFloat(long.Float64, 'f', 6, 64)
		dataRow.SubBasin_id = _subbasin_id.Int64
		dataRow.Title = title.String
		dataRow.Description = description.String
		dataRow.MediaType = media_type.String
		dataRow.URL = base_url.String + filename.String
		dataRow.CctvFlash = cctv_flash.String
		dataRow.CctvQuickTime = cctv_quicktime.String
		dataRow.IsActive = is_active.Bool

		geocodeObj := &model_lt_geocode.Struct_Geocode{}
		geocodeObj.Tumbon_code = tumbon_code.String
		geocodeObj.Tumbon_name = json.RawMessage(tumbon_name.String)
		geocodeObj.Amphoe_code = amphoe_code.String
		geocodeObj.Amphoe_name = json.RawMessage(amphoe_name.String)
		geocodeObj.Province_code = province_code.String
		geocodeObj.Province_name = json.RawMessage(province_name.String)
		geocodeObj.Area_code = area_code.String
		geocodeObj.Area_name = json.RawMessage(area_name.String)
		dataRow.Geocode = geocodeObj

		/*—– Basin —–*/
		dataRow.Basin = &model_basin.Struct_Basin{}
		dataRow.Basin.Id = basin_id.Int64
		dataRow.Basin.Basin_code = _basin_code.Int64
		if _basin_name.String == "" {
			_basin_name.String = "{}"
		}
		dataRow.Basin.Basin_name = json.RawMessage(_basin_name.String)

		data = append(data, dataRow)
	}

	return data, nil
}
