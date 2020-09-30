package lt_servicemethod

import ()

var sqlGetServiceMethod = " SELECT s.id, s.servicemethod_name FROM lt_servicemethod s WHERE s.deleted_by IS NULL "
var sqlGetServiceMethodOrderBy = " ORDER BY s.servicemethod_name->>'th' "

var sqlUpdateServiceMethod = ` UPDATE lt_servicemethod
								  SET updated_by = $2 
								  , servicemethod_name = $3
								  WHERE id = $1 `
