// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package canal_station is a model for public.m_canal_station table. This table store m_canal_station information.
package canal_station

import ()

var SQL_GetCanalStaion = "SELECT mcs.id, mcs.canal_station_name, mcs.canal_station_lat, mcs.canal_station_long, lg.geocode, lg.province_name, lg.province_code " +
	" FROM m_canal_station mcs " +
	" INNER JOIN lt_geocode lg ON mcs.geocode_id = lg.id "
var SQL_GetCanalStaion_OrderBy = " ORDER BY mcs.canal_station_name->>'th' "
var SQL_GetCanalStaion_OrderByProvince = " ORDER BY province_code ASC "

var sqlGetCanalStation = ` SELECT a.id
							, a.canal_station_oldcode
							, a.canal_station_name
							, agt.agency_name
							, agt.agency_shortname
							, a.canal_station_lat
							, a.canal_station_long
							, g.geocode
							, sb.subbasin_name
						FROM m_canal_station a
						LEFT JOIN agency agt ON agt.id = a.agency_id
						LEFT JOIN lt_geocode g ON g.id = a.geocode_id
						LEFT JOIN subbasin sb ON sb.id = a.subbasin_id `

var sqlGetCanalStationOrderBy = `  ORDER BY a.canal_station_oldcode `
