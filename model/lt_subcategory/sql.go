package lt_subcategory

import ()

var SQL_selectSubcategory = " SELECT id, subcategory_name FROM lt_subcategory "

var sqlGetSubCategory = "SELECT s.id , s.subcategory_name, c.id , c.category_name " +
	" FROM lt_subcategory s " +
	" INNER JOIN lt_category c ON s.category_id = c.id and c.deleted_at = '1970-01-01 07:00:00+07' "

var sqlGetSubCategoryWhereDeletedIsNull = " WHERE s.deleted_at = '1970-01-01 07:00:00+07' "
var sqlGetSubCategoryOrderBy = " ORDER BY s.id , c.id "

var sqlInsertSubCategory = " INSERT INTO lt_subcategory (category_id , created_by , subcategory_name) VALUES($1 , $2 , $3) RETURNING id "

var sqlUpdateSubCategory = " UPDATE lt_subcategory SET category_id = $1 , updated_by = $2 , subcategory_name = $3 , updated_at = NOW() WHERE id = $4 "

var sqlCheckChild = " SELECT id FROM metadata WHERE subcategory_id = $1 AND deleted_by IS NULL LIMIT 1 "
var sqlUpdateToDeleteSubCategory = " UPDATE lt_subcategory SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "
