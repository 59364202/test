// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package basin is a model for public.basin table. This table store basin.
package basin

import ()

var SQL_selectAllBasinFromMetaId = "SELECT b.id , b.basin_code , b.basin_name " +
	" FROM basin b " +
	" INNER JOIN agency a ON b.agency_id = a.id " +
	" INNER JOIN metadata m ON m.agency_id = a.id " +
	" WHERE m.id = $1 "

var SQL_selectBasinFromCode = "SELECT id , basin_code , basin_name FROM basin WHERE basin_code IN ("
var SQL_selectBasinFromCode_end = " )"

var SQL_selectAllBasin = "SELECT b.id , b.basin_code , b.basin_name " +
	" FROM basin b " +
	" WHERE b.deleted_by IS NULL "
