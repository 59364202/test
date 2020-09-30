// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agency is a model for public.agency table. This table store agency.
package agency

import ()

var SQL_SelectAgency = ` SELECT a.id, a.agency_name, a.agency_shortname, d.id, d.department_name, m.id, m.ministry_name 
	FROM agency a
	INNER JOIN lt_department d ON a.department_id = d.id AND d.deleted_at = '1970-01-01 07:00:00+07'
	INNER JOIN lt_ministry m ON d.ministry_id = m.id AND m.deleted_at = '1970-01-01 07:00:00+07' 
	WHERE a.deleted_at = '1970-01-01 07:00:00+07' `

var SQL_SelectAgencyShopping = `
SELECT ag.id, ag.agency_name, md.count, msc.count_status_connect, mswu.count_status_wait_update, mswc.count_status_wait_connect, ds.count_dataservice
FROM agency ag
LEFT JOIN
  ( SELECT count(agency_id) AS COUNT,
           agency_id
   FROM metadata
   GROUP BY agency_id) md ON ag.id = md.agency_id
LEFT JOIN
  ( SELECT COUNT (metadatastatus_id) AS count_status_connect,
                 agency_id
   FROM metadata
   WHERE metadatastatus_id = 1
   AND deleted_at = to_timestamp(0)
   GROUP BY agency_id) msc ON ag.id = msc.agency_id
LEFT JOIN
  ( SELECT COUNT (metadatastatus_id) AS count_status_wait_update,
                 agency_id
   FROM metadata
   WHERE metadatastatus_id = 3
   AND deleted_at = to_timestamp(0)
   GROUP BY agency_id) mswu ON ag.id = mswu.agency_id
LEFT JOIN
  ( SELECT COUNT (metadatastatus_id) AS count_status_wait_connect,
                 agency_id
   FROM metadata
   WHERE metadatastatus_id = 2
   AND deleted_at = to_timestamp(0)
   GROUP BY agency_id) mswc ON ag.id = mswc.agency_id
LEFT JOIN
   ( SELECT COUNT(od.id) AS count_dataservice, m.agency_id
   FROM dataservice.order_detail od
   INNER JOIN metadata m ON od.metadata_id = m.id
   GROUP BY m.agency_id
   ) ds ON ag.id = ds.agency_id
LEFT JOIN api.user u ON ag.id = u.agency_id
AND u.id = $1
WHERE ag.deleted_at = '1970-01-01 07:00:00+07'
ORDER BY u.agency_id, ag.id`

var SQL_SelectAgencyMetadataSummary = `SELECT a.id, a.agency_shortname->>'en' AS agency_shortname, b.total, c.count_in_month
FROM agency a
LEFT JOIN
  (SELECT count(id) AS total, agency_id
   FROM metadata m
   WHERE deleted_at = '1970-01-01 07:00:00+07'
   GROUP BY agency_id) b ON a.id= b.agency_id
LEFT JOIN
  (SELECT count(dataimport_dataset_id) AS count_in_month, agency_id
   FROM
     (SELECT DISTINCT(ddl.dataimport_dataset_id), m.agency_id
      FROM metadata m
      INNER JOIN api.dataimport_dataset_log ddl ON m.dataimport_dataset_id = ddl.dataimport_dataset_id
      WHERE m.deleted_at = '1970-01-01 07:00:00+07'
        AND date_part('year',ddl.created_at) = $1
        AND date_part('month',ddl.created_at) = $2
      GROUP BY m.agency_id,
               ddl.dataimport_dataset_id) a
   GROUP BY agency_id) c ON a.id= c.agency_id
WHERE a.deleted_at = '1970-01-01 07:00:00+07'
ORDER BY a.id`

var SQL_CheckChild = ` SELECT agency_id FROM metadata WHERE agency_id = $1
	  UNION
	  SELECT agency_id FROM lt_hydroinfo_agency WHERE agency_id = $1
	  AND deleted_at = to_timestamp(0)
	  AND 	deleted_by is not null`

/* =================================================================== INSERT ====================================================================== */
var SQL_InsertAgency = `INSERT INTO agency (agency_name, agency_shortname, department_id, created_by, updated_by, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $4, NOW(), NOW()) RETURNING id `

/* =================================================================== UPDATE ====================================================================== */
var SQL_UpdateAgency = ` UPDATE agency SET agency_name = $2, agency_shortname = $3, department_id = $4, updated_by = $5, updated_at = NOW() WHERE id = $1 `
var SQL_UpdateAgencyToDelete = ` UPDATE agency SET deleted_by = $2, deleted_at = NOW(), updated_by = $2, updated_at = NOW() WHERE id = $1 `

/* =================================================================== DELETE ====================================================================== */
