// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package air_station is a model for public.air_station table. This table store air_station.
package air_station

import (
	"encoding/json"
)

var AirStationForCheckMetadataColumn = []string{
	"id",
	"station_oldcode",
	"station_name",
	"agency_name",
	"lat",
	"long",
	"geocode",
}

type Struct_AirStation_ForCheckMetadata struct {
	Id      int64           `json:"id"`
	Name    json.RawMessage `json:"station_name"`
	OldCode string          `json:"station_oldcode"`
	Lat     interface{}     `json:"lat"`
	Long    interface{}     `json:"long"`
	Geocode string          `json:"geocode"`
	//Subbasin_name		json.RawMessage	`json:"subbasin_name"`
	AgencyShortname json.RawMessage `json:"agency_shortname"`
	AgencyName      json.RawMessage `json:"agency_name"`
}

type Struct_AirStation_InputParam struct {
	ColumnName string `json:"column_name"`
	AgencyID   string `json:"agency_id"`
}
