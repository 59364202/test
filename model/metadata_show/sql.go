// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show

var sqlSelect = `SELECT md.id,m.id as metadata_id,metadataservice_name->> 'th' as metadata_name,agency.id as agency_id,agency_name->> 'th' as agency_name,md.metadata_show_system_id,metadata_show_system ->> 'th' as metadata_show_system,md.subcategory_id,subcategory_name ->> 'th' as subcategory_name,connection_format,metadata_method 
					FROM metadata m 
					LEFT JOIN metadata_show md ON m.id  =  md.metadata_id
					LEFT JOIN agency ON m.agency_id = agency.id
					LEFT JOIN lt_metadata_show_system ms ON md.metadata_show_system_id = ms.id					
					LEFT JOIN lt_subcategory s ON md.subcategory_id = s.id
					ORDER BY metadataservice_name->>'th'
					`

//WHERE md.metadata_id = $1 AND m.deleted_by IS NULL AND h.deleted_by IS NULL `

var sqlInsert = ` INSERT INTO metadata_show (metadata_id, metadata_show_system_id,subcategory_id, created_by, updated_by, created_at, updated_at)
				  VALUES ($1, $2, $3, $4, $4, NOW(), NOW()) RETURNING id `

var sqlDelete = "DELETE from  metadata_show WHERE id=$1"
