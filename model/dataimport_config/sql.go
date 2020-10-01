// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dataimport_config is a model for api.dataimport_download table. This table store dataimport config information.
package dataimport_config

var (
	getDownloadConfig = "SELECT id, agent_user_id, download_setting, download_script, crontab_setting, download_name, is_cronenabled, description, node, max_process FROM api.dataimport_download " +
		"WHERE id = $1 AND deleted_at IS NULL"

	getDownloadConfigList = "SELECT dd.id, dd.download_script, dd.download_name, dd.is_cronenabled, dd.download_setting#>>'{source_options,0,details,0,host}', dd.description, dds.dataimport_download_id, dd.node " +
		"FROM api.dataimport_download dd LEFT JOIN api.dataimport_dataset dds ON dd.id=dds.dataimport_download_id " +
		"WHERE dd.deleted_at IS NULL GROUP BY  dd.id, dd.download_script, dd.download_name, dd.is_cronenabled, dd.download_setting#>>'{source_options,0,details,0,host}', dd.description, dds.dataimport_download_id ORDER BY dd.id"

	getDownloadConfigNameForDataset = "SELECT id, download_name, download_setting->>'result_file', download_setting#>>'{source_options,0,details,0,files,0,destination}', " +
		"download_setting#>>'{source_options,0,details,0,files,0,source}', download_setting#>>'{source_options,0,details,0,params,output_filename}', " +
		"download_setting#>>'{source_options,0,details,0,params,table_name}', download_setting#>>'{source_options,0,details,0,params,old_table_name}', download_script " +
		"FROM api.dataimport_download WHERE deleted_at IS NULL "

	getDownloadConfigNameForDatasetOrderBy = " ORDER BY id "

	getDatasetConfigList = `SELECT ds.id, dl.download_name, ds.convert_name,ds.dataimport_download_id,dl.download_script, replace((((((((import_setting #> '{configs}'::text[]) ->> 0)::json) -> 'imports'::text) ->> 0)::json) -> 'destination'::text)::text, '"'::text, ''::text) AS table_name FROM api.dataimport_dataset ds LEFT JOIN api.dataimport_download dl ON ds.dataimport_download_id=dl.id WHERE ds.deleted_at IS NULL `

	getDatasetConfigListOrderBy = " ORDER BY ds.id"

	getDatasetConfig = `SELECT id, agent_user_id, dataimport_download_id, convert_setting, import_setting, lookup_table, import_table, convert_script, import_script, convert_name 
		FROM api.dataimport_dataset WHERE id = $1 AND deleted_at IS NULL`

	//	getMonitorScript = "SELECT scriptname FROM public.metadata WHERE id = $1 AND deleted_at = '1970-01-01 07:00:00+07' ORDER BY scriptname"

	insDownloadConfig = `INSERT INTO api.dataimport_download(agent_user_id, download_setting, created_by, created_at, updated_by, updated_at, download_script, crontab_setting, download_name, description, node, max_process, is_cronenabled) 
	VALUES ($1, $2, $3, NOW(), $3, NOW(), $4, $5, $6, $7, $8, $9, $10) returning id`

	insDatasetConfig = "INSERT INTO api.dataimport_dataset(agent_user_id, dataimport_download_id, convert_setting , import_setting, lookup_table, import_table, created_by, created_at, updated_by, updated_at, convert_script, import_script, convert_name) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), $7, NOW(), $8, $9, $10) returning id"

	cpDownloadConfig = "INSERT INTO api.dataimport_download (agent_user_id	,download_setting	,created_by	,created_at	,updated_by	,updated_at	,deleted_by	,deleted_at	,download_script	,crontab_setting	,download_name	) SELECT agent_user_id	,download_setting	,created_by	,created_at	,updated_by	,updated_at	,deleted_by	,deleted_at	,download_script	,crontab_setting	,download_name||'-copy'    FROM  api.dataimport_download WHERE id=$1 returning id "

	cpDatasetConfig = "INSERT INTO api.dataimport_dataset (    agent_user_id , dataimport_download_id , convert_setting , import_setting , lookup_table , import_table , created_by , created_at , updated_by , updated_at , deleted_by , deleted_at , convert_script , import_script , convert_name ) SELECT agent_user_id , dataimport_download_id , convert_setting , import_setting , lookup_table , import_table , created_by , created_at , updated_by , updated_at , deleted_by , deleted_at , convert_script , import_script , convert_name||'-copy'   FROM  api.dataimport_dataset WHERE id=$1 returning id"

	updateMetaData = "UPDATE public.metadata SET updated_by=$2, updated_at=NOW(), dataimport_download_id=$3, dataimport_dataset_id=$4, additional_dataset=$5 WHERE id = $1"

	updateMetaDataProvision = "UPDATE public.metadata_provision SET updated_by=$2, updated_at=NOW(), dataimport_download_id=$3, dataimport_dataset_id=$4 WHERE id = $1"

	updateDownloadConfig = `UPDATE api.dataimport_download SET agent_user_id=$1, download_setting=$2, updated_by=$3, updated_at=NOW(), download_script=$4, crontab_setting=$6, 
	 download_name=$7, description=$8, node=$9, max_process=$10, is_cronenabled=$11 WHERE id = $5`

	updateDatasetConfig = "UPDATE api.dataimport_dataset SET agent_user_id=$1, dataimport_download_id=$2, convert_setting=$3, import_setting=$4, lookup_table=$5, import_table=$6, updated_by=$7, updated_at=NOW(), convert_script=$8, import_script = $9, convert_name = $10 WHERE id = $11"

	getMetaDataList = "SELECT a.id, a.agency_name->>'th', md.id, md.metadataservice_name->>'th' AS metadata_name, dl.id AS download_id,dl.download_name, ds.id AS dataset_id, ds.convert_name, additional_dataset FROM public.metadata md LEFT JOIN api.dataimport_download dl ON md.dataimport_download_id=dl.id LEFT JOIN api.dataimport_dataset ds ON md.dataimport_dataset_id=ds.id LEFT JOIN public.agency a ON md.agency_id=a.id where md.deleted_at = '1970-01-01 07:00:00+07' order by md.id"

	getMetaDataListProvision = "SELECT a.id, a.agency_name->>'th', md.id, md.metadataservice_name->>'th' AS metadata_name, dl.id AS download_id,dl.download_name, ds.id AS dataset_id, ds.convert_name, null FROM public.metadata_provision md LEFT JOIN api.dataimport_download dl ON md.dataimport_download_id=dl.id LEFT JOIN api.dataimport_dataset ds ON md.dataimport_dataset_id=ds.id LEFT JOIN public.agency a ON md.agency_id=a.id where md.deleted_at = '1970-01-01 07:00:00+07' order by md.id"

	//	getParentTable = "SELECT p.relname AS parent,isc.column_name FROM pg_inherits JOIN pg_class as p ON (inhparent=p.oid) LEFT JOIN information_schema.columns isc ON p.relname=isc.table_name GROUP BY parent,isc.column_name ORDER BY parent, isc.column_name"
	//	getParentTable = "SELECT parent,column_name FROM ((SELECT p.relname AS parent,isc.column_name FROM pg_inherits JOIN pg_class as p ON (inhparent=p.oid) LEFT JOIN information_schema.columns isc ON p.relname=isc.table_name GROUP BY parent,isc.column_name ORDER BY parent, isc.column_name) UNION " +
	//		"(SELECT a.table_name,isc.column_name FROM information_schema.tables as a LEFT JOIN information_schema.columns isc ON a.table_name=isc.table_name WHERE a.table_schema = 'public' and not exists ( SELECT distinct(p.relname) FROM pg_inherits  JOIN pg_class p ON (inhparent=p.oid) where  a.table_name = p.relname ) and a.table_name not in ( select relname from pg_inherits i join pg_class c on c.oid = inhrelid ) order by a.table_name, isc.column_name))t ORDER BY 1"

	getParentTable = "SELECT cls.relname, attr.attname FROM pg_attribute attr JOIN pg_class cls ON " +
		" (cls.oid = attrelid AND cls.relkind = 'r' AND cls.relname NOT SIMILAR TO '%_y[0-9]{4}m[0-9]{2}') " +
		" JOIN pg_namespace nsp ON (nsp.oid = cls.relnamespace AND nsp.nspname = 'public') WHERE attr.attnum > 0 ORDER BY relname, attname"

	getMasterTable = "SELECT a.table_name,isc.column_name FROM information_schema.tables as a LEFT JOIN information_schema.columns isc ON a.table_name=isc.table_name WHERE a.table_schema = 'public' and not exists ( SELECT distinct(p.relname) FROM pg_inherits  JOIN pg_class p ON (inhparent=p.oid) where  a.table_name = p.relname ) and a.table_name not in ( select relname from pg_inherits i join pg_class c on c.oid = inhrelid ) order by a.table_name, isc.column_name"

	getAgentDataimport = "SELECT u.id,u.account FROM api.user u left join api.agent a on u.id=a.user_id WHERE a.agent_type_id = 3 AND a.deleted_at IS NULL order by u.account"

	getRDLNodesFromSystemSetting = `
	SELECT Replace(name, 'server.service.dataimport.RDLNodes.', '') AS name 
		   , id
	FROM   api.system_setting 
	WHERE  name LIKE 'server.service.dataimport.RDLNodes.%' 
	`

	getDownloadNameList = "SELECT id, download_name,download_setting#>>'{source_options,0,details,0,host}' FROM api.dataimport_download WHERE deleted_at IS NULL ORDER BY id"

	getDatasetNameList = "SELECT id, convert_name, dataimport_download_id FROM api.dataimport_dataset WHERE deleted_at IS NULL ORDER BY id"

	getDownloadDatasetListID = "SELECT id, dataimport_download_id FROM api.dataimport_dataset WHERE deleted_at IS NULL ORDER BY dataimport_download_id"

	deleteDownloadConfig = "UPDATE api.dataimport_download SET is_cronenabled=false, updated_by=$1, updated_at=NOW(), deleted_by=$1, deleted_at=NOW() WHERE id = $2"

	deleteDatasetConfigByDownloadID = "UPDATE api.dataimport_dataset SET updated_by=$1, updated_at=NOW(), deleted_by=$1, deleted_at=NOW() WHERE dataimport_download_id = $2"

	deleteDownloadIDMetadata = "UPDATE public.metadata SET updated_by=$1, updated_at=NOW(), dataimport_download_id=NULL, dataimport_dataset_id=NULL WHERE dataimport_download_id = $2"

	deleteDatasetConfig = "UPDATE api.dataimport_dataset SET updated_by=$1, updated_at=NOW(), deleted_by=$1, deleted_at=NOW() WHERE id = $2"

	deleteDatasetIDMetadata = "UPDATE public.metadata SET updated_by=$1, updated_at=NOW(), dataimport_dataset_id=NULL WHERE dataimport_dataset_id = $2"

	updateIsCronenabled = "UPDATE api.dataimport_download SET updated_by=$1, updated_at=NOW(), is_cronenabled = $3 WHERE id = $2"

	getAgencyList = "SELECT id,agency_name->>'th' AS agency_name FROM  public.agency WHERE deleted_at = to_timestamp(0) ORDER BY agency_name"

	getMetadataList = "SELECT id,metadata.metadataservice_name->>'th',agency_id AS metadata_name FROM public.metadata WHERE deleted_at = to_timestamp(0) ORDER BY metadata_name"

	getHistoryList = "SELECT ddl.dataimport_download_id, ddl.id, ddsl.dataimport_dataset_id, ddsl.id, ddl.download_begin_at,ddl.download_duration,ddsl.convert_begin_at, " +
		"ddsl.convert_duration,ddsl.import_begin_at,ddsl.import_duration,ddsl.process_status,ddl.download_bytes_count/1024 as filesize, " +
		"md.metadataservice_name,md.metadata_convertfrequency,md.metadata_channel,dd.download_script,a.agency_name,download_event_code_id, " +
		"ddsl.convert_event_code_id , ddsl.import_event_code_id, ec.description,ec1.description,ec2.description " +
		"FROM api.dataimport_download_log ddl " +
		"LEFT JOIN api.dataimport_dataset_log ddsl ON ddl.id=ddsl.dataimport_download_log_id " +
		"LEFT JOIN api.dataimport_download dd ON ddl.dataimport_download_id=dd.id " +
		"LEFT JOIN public.metadata md ON ddl.dataimport_download_id=md.dataimport_download_id AND " +
		"			     CASE " +
		"				WHEN ddsl.dataimport_dataset_id IS NOT NULL THEN  ddsl.dataimport_dataset_id=md.dataimport_dataset_id " +
		"				WHEN ddsl.dataimport_dataset_id IS NULL THEN TRUE " +
		"			     END " +
		"LEFT JOIN public.agency a ON md.agency_id=a.id " +
		"LEFT JOIN api.lt_event_code ec ON ddl.download_event_code_id=ec.id " +
		"LEFT JOIN api.lt_event_code ec1 ON ddsl.convert_event_code_id=ec1.id " +
		"LEFT JOIN api.lt_event_code ec2 ON ddsl.import_event_code_id=ec2.id "

	getMetadataListProvision = "SELECT id,metadata_provision.metadataservice_name->>'th',agency_id AS metadata_name FROM public.metadata_provision WHERE deleted_at = to_timestamp(0) ORDER BY metadata_name"

	getHistoryListProvision = "SELECT ddl.dataimport_download_id, ddl.id, ddsl.dataimport_dataset_id, ddsl.id, ddl.download_begin_at,ddl.download_duration,ddsl.convert_begin_at, " +
		"ddsl.convert_duration,ddsl.import_begin_at,ddsl.import_duration,ddsl.process_status,ddl.download_bytes_count/1024 as filesize, " +
		"md.metadataservice_name,dd.download_script,a.agency_name,download_event_code_id, " +
		"ddsl.convert_event_code_id , ddsl.import_event_code_id, ec.description,ec1.description,ec2.description " +
		"FROM api.dataimport_download_log ddl " +
		"LEFT JOIN api.dataimport_dataset_log ddsl ON ddl.id=ddsl.dataimport_download_log_id " +
		"LEFT JOIN api.dataimport_download dd ON ddl.dataimport_download_id=dd.id " +
		"LEFT JOIN public.metadata_provision md ON ddl.dataimport_download_id=md.dataimport_download_id AND " +
		"			     CASE " +
		"				WHEN ddsl.dataimport_dataset_id IS NOT NULL THEN  ddsl.dataimport_dataset_id=md.dataimport_dataset_id " +
		"				WHEN ddsl.dataimport_dataset_id IS NULL THEN TRUE " +
		"			     END " +
		"LEFT JOIN public.agency a ON md.agency_id=a.id " +
		"LEFT JOIN api.lt_event_code ec ON ddl.download_event_code_id=ec.id " +
		"LEFT JOIN api.lt_event_code ec1 ON ddsl.convert_event_code_id=ec1.id " +
		"LEFT JOIN api.lt_event_code ec2 ON ddsl.import_event_code_id=ec2.id "

	getAgentName = "SELECT account FROM api.user WHERE id=$1 AND deleted_at IS NULL"

	getDownloadCronList = "SELECT dl.id, download_name, crontab_setting, download_script, is_cronenabled, node, u.account, description FROM api.dataimport_download dl LEFT JOIN api.user u ON u.id = dl.agent_user_id WHERE dl.deleted_at IS NULL"

	updateDownloadCron = "UPDATE api.dataimport_download SET updated_by=$1, updated_at=NOW(), crontab_setting = $3 WHERE id = $2"

	getConfigVariable = "SELECT aa.id , cc.name_cat , aa.config_name, aa.variable_name, aa.value from api.config_variable aa left join api.config_variable_category cc on cc.id = aa.category Where deleted_at IS NULL"

	getListVariable = "SELECT id , name_cat from api.config_variable_category"
	// getListVariable = "SELECT id , name_cat from api.config_variable_category"

)
