package subcategory

import ()

//var sqlGetSubCategory = " SELECT  t.subcategory_id,t.subcategory_name ,tc.category_id ,tc.category_name ,l.code  " +
//	" FROM language l  " +
//	" LEFT JOIN lt_category_translate tc ON tc.language_id = l.id " +
//	" LEFT JOIN lt_subcategory lt ON lt.category_id = tc.category_id " +
//	" LEFT JOIN lt_subcategory_translate t ON t.subcategory_id = lt.id AND t.language_id = l.id "

var sqlGetSubCategory = "SELECT s.id , s.subcategory_name, c.id , c.category_name " +
	" FROM lt_subcategory s " +
	" INNER JOIN lt_category c ON s.category_id = c.id and c.deleted_at = '1970-01-01 07:00:00+07' "

//var sqlGetSubCategoryWhereCategoryId = " WHERE s.category_id IN ($1) AND s.deleted_by IS NULL "
var sqlGetSubCategoryWhereDeletedIsNull = " WHERE s.deleted_at = '1970-01-01 07:00:00+07' "
var sqlGetSubCategoryOrderBy = " ORDER BY s.id , c.id "

var sqlInsertSubCategory = " INSERT INTO lt_subcategory (category_id , created_by , subcategory_name) VALUES($1 , $2 , $3) RETURNING id "

var sqlUpdateSubCategory = " UPDATE lt_subcategory SET category_id = $1 , updated_by = $2 , subcategory_name = $3 , updated_at = NOW() WHERE id = $4 "

var sqlCheckChild = " SELECT id FROM metadata WHERE subcategory_id = $1 AND deleted_by IS NULL LIMIT 1 "
var sqlUpdateToDeleteSubCategory = " UPDATE lt_subcategory SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "
