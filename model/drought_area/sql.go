// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package drought_area is a model for public.drought_area table. This table store drought_area.
package drought_area

import ()

var SQL_GetLatestDrought = `
SELECT lg.province_code, 
       lg.province_name 
FROM   drought_area da 
       INNER JOIN lt_geocode lg 
               ON da.geocode_id = lg.id 
                  AND da.deleted_at = To_timestamp(0) 
                  AND da.drought_datetime = (SELECT 
                      Max(drought_datetime) AS drought_datetime 
                                             FROM   drought_area 
                                             WHERE  agency_id = 16 
                                                    AND deleted_at = 
                                                        To_timestamp(0)) 
ORDER  BY lg.province_code 
`
