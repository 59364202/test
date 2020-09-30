// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package air_station is a model for public.air_station table. This table store air_station.
package air_station

import ()

var sqlGetAirStation = ` SELECT a.id
							, a.air_station_oldcode
							, a.air_station_name
							, agt.agency_name
							, agt.agency_shortname
							, a.air_station_lat
							, a.air_station_long
							, g.geocode
						FROM m_air_station a
						LEFT JOIN agency agt ON agt.id = a.agency_id
						LEFT JOIN lt_geocode g ON g.id = a.geocode_id `
var sqlGetAirStationOrderBy = `  ORDER BY a.air_station_oldcode `
