package event_log

var sqlSelectEventLog = ` SELECT b.id AS event_log_id
							    , b.created_at AS event_log_date
							    , b.event_data AS event_log_data
							    , b.event_message AS event_log_message
							    , t.sumrunscript AS event_log_duration
							    , mt.id AS metadata_id
							    , mt.metadataservice_name
							    , mt.metadataagency_name
							    , ag.id AS agency_id
							    , ag.agency_shortname
							    , ag.agency_name
							    , e.id AS event_type_id
							    , e.code AS event_type_code
							    , e.description AS event_type_desc
							    , d.id AS event_subtype_id
							    , d.code AS event_subtype_code
							    , d.description AS event_subtype_desc
					   FROM (SELECT b_1.id, x.convert_duration + x.import_duration AS sumrunscript
						     FROM api.event_log b_1
						     JOIN api.dataimport_dataset_log x ON x.id = (b_1.event_data->>'dataset_log_id')::bigint
						     WHERE (b_1.created_at >= $1 AND b_1.created_at <= $2)
					         UNION
					         SELECT b_1.id, z.download_duration AS sumrunscript
				             FROM api.event_log b_1
				             JOIN api.dataimport_download_log z ON z.id = (b_1.event_data ->'result'->>'id')::bigint
				             WHERE (b_1.created_at >= $1 AND b_1.created_at <= $2)) t
					 LEFT JOIN api.event_log b ON t.id = b.id
				     LEFT JOIN api.lt_event_code d ON b.event_code_id = d.id
				     LEFT JOIN api.dataimport_dataset_log ddl ON ddl.id = (b.event_data->>'dataset_log_id')::bigint
				     LEFT JOIN api.dataimport_dataset dd ON dd.id = (b.event_data ->>'dataset_id')::bigint
				     LEFT JOIN api.lt_event_log_category e ON e.id = b.event_log_category_id
				     LEFT JOIN metadata mt ON mt.dataimport_dataset_id = dd.id
				     LEFT JOIN agency ag ON ag.id = mt.agency_id   `

var sqlSelectEventLogWhere = ` WHERE (b.created_at >= $1 AND b.created_at <= $2) `

/*
var sqlSelectEventLogWhere = ` WHERE (b.created_at >= $1 AND b.created_at <= $2)
								   --AND d.deleted_by IS NULL AND (d.deleted_at IS NULL OR d.deleted_at = '1970-01-01 07:00:00+07')
								   --AND dd.deleted_by IS NULL AND (dd.deleted_at IS NULL OR dd.deleted_at = '1970-01-01 07:00:00+07')
								   --AND e.deleted_by IS NULL AND (e.deleted_at IS NULL OR e.deleted_at = '1970-01-01 07:00:00+07')
								   --AND mt.deleted_by IS NULL AND (mt.deleted_at IS NULL OR mt.deleted_at = '1970-01-01 07:00:00+07')
								   --AND ag.deleted_by IS NULL AND (ag.deleted_at IS NULL OR ag.deleted_at = '1970-01-01 07:00:00+07')  `
*/
var sqlSelectEventLogOrderBy = ` ORDER BY b.created_at DESC `

var sqlSelectEventLogReport = ` ORDER BY b.created_at DESC `

var SQL_EventReport = `
WITH event_dl 
     AS (SELECT l.id 
                , event_data ->> 'download_id'                            AS 
                  donwload_id 
                , event_data -> 'params' ->> 'dataimport_download_log_id' AS 
                  download_log_id 
                , u.account                                               AS 
                  agent_name 
                , u.id                                                    AS 
                  agent_id 
                , l.created_at 
                , e.code                                                  AS 
                  event_code 
                , l.event_code_id 
         FROM   api.event_log l 
                inner join api.lt_event_code e 
                        ON ( e.id = l.event_code_id ) 
                inner join api.USER u 
                        ON ( u.id = l.agent_user_id ) 
                inner join api.agent a 
                        ON ( a.user_id = u.id ) 
         WHERE  l.created_at >= Date_trunc('day', $1 :: timestamptz) 
                AND l.created_at < Date_trunc('day', $1 :: 
                                                     timestamptz + interval '1' 
                                                     day) 
                AND a.agent_type_id = 3 
                AND u.account = $2
                AND l.event_code_id = $3
                AND event_data -> 'download_id' IS NOT NULL), 
     event_ds 
     AS (SELECT l.id 
                , event_data ->> 'dataset_id'    AS dataset_id 
                , event_data ->> 'dataset_log_id'AS dataset_log_id 
                , u.account                      AS agent_name 
                , u.id                           AS agent_id 
                , l.created_at 
                , e.code                         AS event_code 
                , l.event_code_id 
         FROM   api.event_log l 
                inner join api.lt_event_code e 
                        ON ( e.id = l.event_code_id ) 
                inner join api.USER u 
                        ON ( u.id = l.agent_user_id ) 
                inner join api.agent a 
                        ON ( a.user_id = u.id ) 
         WHERE  l.created_at >= Date_trunc('day', $1 :: timestamptz) 
                AND l.created_at < Date_trunc('day', $1 :: 
                                                     timestamptz + interval '1' 
                                                     day) 
                AND a.agent_type_id = 3 
                AND u.account = $2
                AND l.event_code_id = $3
                AND event_data -> 'dataset_id' IS NOT NULL), 
        event_ac 
        AS (SELECT l.id 
                , event_message                 AS event_message 
                , event_data ->> 'access_log_id'AS access_log_id 
                , u.account                     AS agent_name 
                , u.id                          AS agent_id 
                , l.created_at 
                , e.code                        AS event_code 
                , l.event_code_id 
        FROM   api.event_log l 
                inner join api.lt_event_code e 
                        ON ( e.id = l.event_code_id ) 
                inner join api.USER u 
                        ON ( u.id = l.agent_user_id ) 
                inner join api.agent a 
                        ON ( a.user_id = u.id ) 
        WHERE  l.created_at >= Date_trunc('day', $1 :: timestamptz) 
                AND l.created_at < Date_trunc('day', $1 :: 
                                                     timestamptz + interval '1' 
                                                     day) 
                AND a.agent_type_id = 3 
                AND u.account = $2 
                AND l.event_code_id = $3 
                AND event_data -> 'access_log_id' IS NOT NULL), 
     event 
     AS (SELECT 'download'                        AS d_type 
                , edl.* 
                , donwload_id                     AS did 
                , dll.download_detail ->> 'error' AS download_detail 
                , dll.download_begin_at 
                , dl.download_name 
         FROM   event_dl edl 
                inner join api.dataimport_download_log dll 
                        ON edl.download_log_id = dll.id :: text 
                inner join api.dataimport_download dl 
                        ON dll.dataimport_download_id = dl.id 
         UNION 
         SELECT 'dataset'                           AS d_type 
                , eds.* 
                , dsl.dataimport_dataset_id :: text AS did 
                , dsl.data_path 
                , dsl.created_at 
                , ds.convert_name 
         FROM   event_ds eds 
                inner join api.dataimport_dataset_log dsl 
                        ON eds.dataset_log_id = dsl.id :: text 
                inner join api.dataimport_dataset ds 
                        ON dsl.dataimport_dataset_id = ds.id
         UNION 
         SELECT 'access' AS d_type 
                , esa.* 
                , '' 
                , esa.event_message 
                , esa.created_at 
                , '' 
         FROM   event_ac esa) 
SELECT d_type 
       , did               AS id 
       , download_log_id   AS log_id 
       , download_begin_at AS begin_at 
       , download_name     AS name 
       , download_detail   AS detail 
FROM   event 
`
