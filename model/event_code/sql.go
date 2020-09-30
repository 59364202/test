package event_code

var sqlGetEventCode = ` SELECT subevent.id
							 , subevent.code
							 , subevent.description
							 , subevent.is_autoclose
							 , subevent.troubleshoot
							 , subevent.subtype_category
							 , event.id AS event_id
							 , event.code AS event_code
							 , event.description AS event_desc
					    FROM api.lt_event_code subevent
					    LEFT JOIN api.lt_event_log_category event ON event.id = subevent.event_log_category_id
					    WHERE subevent.deleted_at IS NULL AND subevent.deleted_by IS NULL 
						  AND event.deleted_at IS NULL AND event.deleted_by IS NULL `

var sqlGetEventCodeOrderby = ` ORDER BY event.description->>'en', subevent.description->>'en' `

var sqlUpdateEventCode = ` UPDATE api.lt_event_code
								  SET event_log_category_id = $2
								    , code = $3
								    , description = $4
								    , is_autoclose = $5
								    , troubleshoot = $6
								    , subtype_category = $7
								    , updated_by = $8
								    , updated_at = NOW()
								  WHERE id = $1
								    AND deleted_at IS NULL AND deleted_by IS NULL `

var sqlInsertEventCode = ` INSERT INTO api.lt_event_code(
									    event_log_channel_id
									  , event_log_category_id
									  , code
									  , description
									  , is_autoclose
									  , troubleshoot
									  , subtype_category
									  , created_by, updated_by, created_at, updated_at)
							  VALUES (6, $1, $2, $3, $4, $5, $6, $7, $7, NOW(), NOW()) RETURNING id `

var sqlDeleteEventCode = ` DELETE FROM api.lt_event_code WHERE id = $1 AND deleted_at IS NULL AND deleted_by IS NULL `

var sqlCheckEventCodeChild = ` SELECT id FROM api.event_log WHERE event_code_id = $1 LIMIT 1 `

var sqlEventCategoryEventCode = `
SELECT c.id 
       , c.code 
       , e.id 
       , e.description 
FROM   api.lt_event_log_category c 
       INNER JOIN api.lt_event_code e 
               ON c.id = e.event_log_category_id 
ORDER  BY c.code 
          , e.id 
`
