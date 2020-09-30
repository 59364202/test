// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package flood_situation is a model for public.flood_situation table. This table store flood_situation.
package flood_situation

import ()

var SQL_GetLatestFloodSituation = `
SELECT lg.province_code,
       lg.province_name,
       fs.flood_datetime,
       fs.flood_name,
       fs.flood_link,
       fs.flood_description,
       fs.flood_author,
       fs.flood_colorlevel
FROM flood_situation fs
INNER JOIN lt_geocode lg ON fs.geocode_id = lg.id
AND fs.deleted_at = To_timestamp(0)
AND fs.flood_datetime =
  (SELECT Max(flood_datetime) AS flood_datetime
   FROM flood_situation
   WHERE agency_id = 16
     AND deleted_at = To_timestamp(0))
ORDER BY lg.province_code
`
