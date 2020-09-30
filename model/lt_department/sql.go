package lt_department

import ()

var SQL_selectDepartment = " SELECT id , department_name FROM lt_department "

var sqlGetDepartment = " SELECT d.id, d.department_code, d.department_shortname, d.department_name, m.id, m.ministry_name " +
	" FROM lt_department d " +
	" INNER JOIN lt_ministry m ON d.ministry_id = m.id AND m.deleted_at = '1970-01-01 07:00:00+07' "

//var sqlGetDepartmentWhereId = " WHERE d.ministry_id IN ($1) AND d.deleted_by IS NULL "
var sqlGetDepartmentWhereDeletedIsNull = " WHERE d.deleted_at = '1970-01-01 07:00:00+07' "
var sqlGetDepartmentOrderBy = " ORDER BY d.id "

var sqlInsertDepartment = " INSERT INTO lt_department(department_code, department_shortname, department_name, ministry_id, created_by) VALUES($1, $2, $3, $4, $5) RETURNING id "

var sqlUpdateDepartment = " UPDATE lt_department SET department_code = $1, department_shortname = $2, department_name = $3, ministry_id = $4 , updated_by = $5 , updated_at = NOW() WHERE id = $6 "

//var sqlInserDepartmentTran = " INSERT INTO lt_department_translate(department_id, language_id, department_name, department_shortname, created_by ) VALUES ($1, $2, $3, $4, $5) "
//var sqlUpsertDepartmentTran = sqlInserDepartmentTran +
//	" ON CONFLICT (department_id , language_id) " +
//	" DO UPDATE SET department_name = $3 , updated_by = $5 , updated_at = NOW() "

var sqlCheckChild = " SELECT id FROM agency WHERE department_id = $1 AND deleted_by IS NULL LIMIT 1 "
var sqlUpdateToDeleteDepartment = " UPDATE lt_department SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "

//var sqlUpdateToDeleteDepartmentTrans = " UPDATE lt_department_translate SET deleted_by = $2 , deleted_at = NOW() WHERE department_id = $1 "

var sqlDeleteDepartmentTran = " DELETE FROM lt_department_translate WHERE department_id = $1 "
