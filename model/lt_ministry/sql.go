package lt_ministry

import ()

var SQL_selectMinistry = "SELECT id , ministry_name FROM lt_ministry "

var sqlGetMinistry = " SELECT m.id, m.ministry_code, m.ministry_shortname, m.ministry_name " +
	" FROM lt_ministry m " +
	" WHERE m.deleted_at = '1970-01-01 07:00:00+07' " +
	" ORDER BY m.id "

var sqlInsertMinistry = " INSERT INTO lt_ministry (ministry_code, ministry_shortname, ministry_name, created_by) VALUES ($1, $2, $3, $4) RETURNING id "
var sqlUpdateMinistry = " UPDATE lt_ministry SET ministry_code = $1, ministry_shortname = $2, ministry_name = $3, updated_by = $4, updated_at = NOW() WHERE id = $5 "

var sqlCheckChild = " SELECT id FROM lt_department WHERE ministry_id = $1 AND deleted_by IS NULL LIMIT 1"

var sqlUpdateToDeleteMinistry = " UPDATE lt_ministry SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "
