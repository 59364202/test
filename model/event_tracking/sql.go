package event_tracking

import ()

var (
	getEventLogCategory = "SELECT id,description::jsonb,code FROM api.lt_event_log_category WHERE deleted_at IS NULL ORDER BY description::jsonb"

	getEventCode = "SELECT id,description::jsonb,event_log_category_id FROM api.lt_event_code WHERE deleted_at IS NULL AND is_autoclose=FALSE AND id !=1 ORDER BY event_log_category_id,description::jsonb"

	getEventCode2 = "SELECT id,description::jsonb,event_log_category_id FROM api.lt_event_code WHERE deleted_at IS NULL AND is_autoclose=FALSE AND (id !=42 AND id !=1) ORDER BY event_log_category_id,description::jsonb"

	getAgencyList = "SELECT id,agency_name::json AS agency_name FROM  public.agency WHERE deleted_at = to_timestamp(0) ORDER BY agency_name->>'th'"

	getEventInvalidDataSelectOption = `
			SELECT a.id,
		       a.agency_name::jsonb,
		       array_agg(DISTINCT el.created_at::date)
		FROM api.event_log el
		LEFT JOIN api.dataimport_dataset dd ON CAST(el.event_data->>'dataset_id' AS integer)=dd.id
		LEFT JOIN public.metadata md ON dd.id=md.dataimport_dataset_id
		LEFT JOIN public.agency a ON md.agency_id=a.id
		LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id
		LEFT JOIN api.dataimport_download ddl ON dd.dataimport_download_id=ddl.id
		WHERE el.event_code_id=42
		  AND el.solve_event_at IS NULL
		  AND a.agency_name->>'th' IS NOT NULL
		  AND el.event_data->'params'->>'data_path' != ''
		GROUP BY a.id,
		         a.agency_name::jsonb
		ORDER BY a.id,
		         a.agency_name::jsonb
         `

	//	getEventInvalidDataSelectOption = "SELECT DISTINCT el.created_at::date,a.id,a.agency_name::jsonb,a.agency_name->>'th' FROM api.event_log el  " +
	//		"LEFT JOIN api.dataimport_dataset dd ON CAST(el.event_data->>'dataset_id' AS integer)=dd.id  " +
	//		"LEFT JOIN public.metadata md ON dd.id=md.dataimport_dataset_id LEFT JOIN public.agency a ON md.agency_id=a.id  " +
	//		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id LEFT JOIN api.dataimport_download ddl ON dd.dataimport_download_id=ddl.id  " +
	//		"WHERE el.event_code_id=42 AND el.solve_event_at IS NULL AND a.agency_name->>'th' IS NOT NULL AND el.event_data->'params'->>'data_path' != '' " +
	//		"ORDER BY a.agency_name->>'th',el.created_at::date "

	getEventTracking = "SELECT el.id,ec.description::jsonb,el.created_at,a.agency_name::jsonb,elc.description::jsonb,el.event_message,lsq.sent_at,el.solve_event_at,el.event_data->'params'->>'data_path',elc.code, emd.metadataservice_name FROM  api.event_log el  " +
		"LEFT JOIN ( " +
		"SELECT el.id,md.metadataservice_name::text,md.agency_id " +
		"FROM api.event_log el  " +
		"INNER JOIN metadata md ON md.dataimport_dataset_id=(event_data->>'dataset_id')::bigint " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id != 1 " +
		"AND event_data->>'dataset_id' <> '' " +
		"UNION ALL  " +
		"SELECT el.id, string_agg(md.metadataservice_name::text, '|'),md.agency_id " +
		"FROM api.event_log el  " +
		"INNER JOIN metadata md ON md.dataimport_download_id=(event_data->>'download_id')::bigint " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id != 1 " +
		"AND event_data->>'download_id' <> '' " +
		"GROUP BY el.id,md.agency_id " +
		") emd ON el.id = emd.id " +
		"LEFT JOIN api.lt_event_log_category elc ON el.event_log_category_id=elc.id " +
		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id " +
		"LEFT JOIN public.agency a ON emd.agency_id=a.id " +
		"LEFT JOIN api.event_log_sink_queue lsq ON el.id=lsq.event_log_id " +
		"WHERE el.event_code_id != 1 AND ec.is_autoclose=FALSE AND el.created_at BETWEEN $1 AND $2 AND agency_name IS NOT NULL "

	getEventTrackingUpdate = "SELECT el.id,ec.description::jsonb,el.created_at,a.agency_name::jsonb,el.event_message,lsq.sent_at,el.solve_event_at,elc.description::jsonb,elc.code,emd.metadataservice_name FROM api.event_log el  " +
		"LEFT JOIN ( " +
		"SELECT el.id,md.metadataservice_name::text,md.agency_id " +
		"FROM api.event_log el  " +
		"INNER JOIN metadata md ON md.dataimport_dataset_id=(event_data->>'dataset_id')::bigint " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id != 42 AND el.event_code_id != 1 AND el.solve_event_at IS NULL " +
		"AND event_data->>'dataset_id' <> '' " +
		"UNION ALL  " +
		"SELECT el.id, string_agg(md.metadataservice_name::text, '|'),md.agency_id " +
		"FROM api.event_log el  " +
		"INNER JOIN metadata md ON md.dataimport_download_id=(event_data->>'download_id')::bigint " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id != 42 AND el.event_code_id != 1 AND el.solve_event_at IS NULL " +
		"AND event_data->>'download_id' <> '' " +
		"GROUP BY el.id,md.agency_id " +
		") emd ON el.id = emd.id " +
		"LEFT JOIN api.lt_event_log_category elc ON el.event_log_category_id=elc.id " +
		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id " +
		"LEFT JOIN public.agency a ON emd.agency_id=a.id " +
		"LEFT JOIN api.event_log_sink_queue lsq ON el.id=lsq.event_log_id " +
		"WHERE el.event_code_id != 42 AND ec.is_autoclose=FALSE AND el.event_code_id != 1 AND el.solve_event_at IS NULL AND el.created_at BETWEEN $1 AND $2 AND agency_name IS NOT NULL "

	getInvalidData = "SELECT el.id,ec.code,el.created_at,a.agency_name->>'th',dd.convert_name,el.event_data,el.event_data->'params'->>'data_path' AS data_path, md.metadata_method, " +
		"ddl.download_setting->>'result_file' AS result_file, ddl.download_setting#>>'{source_options,0,details,0,files,0,destination}' AS destination, " +
		"ddl.download_setting#>>'{source_options,0,details,0,files,0,source}' AS source, ddl.download_setting#>>'{source_options,0,details,0,params,output_filename}' AS output_filename, " +
		"ddl.download_setting#>>'{source_options,0,details,0,params,table_name}' AS table_name,ddl.download_script,md.metadataservice_name->>'th' FROM api.event_log el " +
		"LEFT JOIN api.dataimport_dataset dd ON CAST(el.event_data->>'dataset_id' AS integer)=dd.id " +
		"LEFT JOIN public.metadata md ON dd.id=md.dataimport_dataset_id LEFT JOIN public.agency a ON md.agency_id=a.id " +
		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id LEFT JOIN api.dataimport_download ddl ON dd.dataimport_download_id=ddl.id " +
		"WHERE el.event_code_id=42 AND el.solve_event_at IS NULL AND el.created_at BETWEEN $1 AND $2 AND a.agency_name->>'th' IS NOT NULL AND el.event_data->'params'->>'data_path' != '' "

	getSendInvalidData = "SELECT el.id,ec.code,el.created_at,dd.convert_name,a.agency_name->>'th',el.event_message,emd.metadataservice_name FROM api.event_log el  " +
		"INNER JOIN (  " +
		"SELECT el.id,md.metadataservice_name::text,md.agency_id,md.dataimport_dataset_id " +
		"FROM api.event_log el   " +
		"INNER JOIN metadata md ON md.dataimport_dataset_id=(event_data->>'dataset_id')::bigint  " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id=42 AND el.send_error_at IS NULL  " +
		"AND event_data->>'dataset_id' <> ''  " +
		"UNION ALL   " +
		"SELECT el.id, string_agg(md.metadataservice_name::text, '|'),md.agency_id,md.dataimport_dataset_id " +
		"FROM api.event_log el   " +
		"INNER JOIN metadata md ON md.dataimport_download_id=(event_data->>'download_id')::bigint  " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id=42 AND el.send_error_at IS NULL  " +
		"AND event_data->>'download_id' <> ''  " +
		"GROUP BY el.id,md.agency_id,md.dataimport_dataset_id " +
		") emd ON el.id = emd.id AND CAST(el.event_data->>'dataset_id' AS integer)=emd.dataimport_dataset_id " +
		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id  " +
		"LEFT JOIN api.dataimport_dataset dd ON emd.dataimport_dataset_id=dd.id " +
		"LEFT JOIN public.agency a ON emd.agency_id=a.id " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id=42 AND el.send_error_at IS NULL AND a.agency_name->>'th' IS NOT NULL AND el.is_downloaded=TRUE "

	getTrackingInvalidData = "SELECT el.id,ec.code,el.created_at,a.agency_name->>'th',el.event_message,el.send_error_at,el.solve_event_at,el.solve_event,md.metadataservice_name " +
		"FROM api.event_log el  " +
		"INNER JOIN metadata md ON md.dataimport_dataset_id=(event_data->>'dataset_id')::bigint  " +
		"LEFT JOIN api.lt_event_code ec ON el.event_code_id=ec.id " +
		"LEFT JOIN api.user u ON el.agent_user_id=u.id " +
		"LEFT JOIN public.agency a ON u.agency_id=a.id " +
		"WHERE el.created_at BETWEEN $1 AND $2 AND el.event_code_id=42 AND a.agency_name->>'th' IS NOT NULL AND el.is_downloaded=TRUE AND send_error_at IS NOT NULL " +
		"AND event_data->>'dataset_id' <> '' "

	updateEventTrackingMessage = "UPDATE api.event_log SET updated_by=$1, updated_at=NOW(), solve_event_at=NOW(), solve_event=$2"

	updateEventSendInvalidData = "UPDATE api.event_log SET updated_by=$1, updated_at=NOW(), send_error_at=$2"

	updateEventTrackingInvalidData = "UPDATE api.event_log SET updated_by=$1, updated_at=NOW(), solve_event_at=NOW(), solve_event=$2"

	updateEventTrackingIsDownloaded = "UPDATE api.event_log SET updated_by=$1, updated_at=NOW(), is_downloaded=TRUE WHERE id=$2"

	getInvalidDataPath = "SELECT event_data->'params'->>'data_path' as data_path FROM api.event_log WHERE id = $1 AND deleted_at IS NULL"
)
