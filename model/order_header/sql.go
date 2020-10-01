// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_header is a model for dataservice.order_header table. This table store order_header.
package order_header

// ------------------------------ insert ------------------------------
var SQL_InsertOrderHeader = "INSERT INTO dataservice.order_header (user_id, order_status_id, order_datetime, order_quality, order_purpose, order_forexternal) VALUES ($1, $5, NOW(), $2, $3, $4) RETURNING id;"

// ------------------------------ select ------------------------------
var SQL_SelectOrderHeaderByUserId = "SELECT oh.id , oh.order_datetime , oh.order_quality , os.id ,os.order_status " +
	" FROM dataservice.order_header oh " +
	" INNER JOIN dataservice.order_status os ON oh.order_status_id = os.id"
	// +" WHERE user_id = $1"

var SQL_SelectOrderById = "SELECT order_status_id, order_quality, order_purpose , order_forexternal " +
	" FROM dataservice.order_header " +
	" WHERE id = $1"

var SQL_SelectOrderForExternal = `
SELECT id, 
       order_datetime 
FROM   dataservice.order_header 
WHERE  order_forexternal = true 
       AND id IN (SELECT DISTINCT( order_header_id ) 
                  FROM   dataservice.order_detail od 
                         INNER JOIN dataservice.order_header oh 
                                 ON od.order_header_id = oh.id 
                                    AND oh.order_forexternal = true 
                  WHERE  od.detail_letterpath IS NULL) 
       AND order_status_id IN (1,2)
`

var SQL_SelectOrderHeader = "SELECT oh.id, oh.order_datetime, u.id, u.full_name, ohs.id, ohs.order_status, a.id, a.agency_name " +
	" FROM dataservice.order_header oh " +
	" INNER JOIN dataservice.order_status ohs ON oh.order_status_id = ohs.id  " +
	" INNER JOIN api.user u ON oh.user_id = u.id " +
	" LEFT JOIN agency a ON u.agency_id = a.id "
var SQL_SelectOrderHeader_OrderById = " ORDER BY oh.id DESC "

// ------------------------------ update ------------------------------
var SQL_UpdateOrderHeaderStatusToCancelByOrderHeaderId = "UPDATE dataservice.order_header SET order_status_id = 3 , updated_at = NOW() , updated_by = $2 WHERE id = $1"

// ------------------------------ delete ------------------------------
var SQL_DeleteOrderHeaderByOrderHeaderId = "DELETE FROM dataservice.order_header WHERE id = $1"
