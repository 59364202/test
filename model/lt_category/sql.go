// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package category is a model for public.lt_category table. This table store lt_category information.
package lt_category

import ()

var SQL_selectCategory = "SELECT id , category_name FROM lt_category c WHERE c.deleted_at = '1970-01-01 07:00:00+07' ORDER BY c.id"

var sqlInsertCategory = " INSERT INTO lt_category (category_name , created_by) VALUES ( $1,$2 ) RETURNING id "

var sqlUpdateCategory = "UPDATE lt_category SET category_name = $1 , updated_by = $2 WHERE id = $3 "

var sqlCheckChild = " SELECT id FROM lt_subcategory WHERE category_id = $1 AND deleted_by IS NULL LIMIT 1"

//var sqlDeleteCategory = " DELETE FROM lt_category WHERE id = $1 "
var sqlUpdateToDeleteCategory = " UPDATE lt_category SET deleted_by = $2 , deleted_at = NOW() WHERE id = $1 "
