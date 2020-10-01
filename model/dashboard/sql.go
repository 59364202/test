package dashboard

var (
	sqlSelectConvertMinute = "SELECT m.id, convert_minute, mf.datafrequency FROM public.metadata m LEFT JOIN public.lt_frequencyunit lfu ON m.frequencyunit_id=lfu.id LEFT JOIN public.metadata_frequency mf ON mf.metadata_id=m.id WHERE m.deleted_at=to_timestamp(0)"

	sqlUpdateMetadataOfflineDate = "UPDATE public.metadata SET metadata_offline_date=$1 WHERE id=$2"

	sqlLastLongin = "SELECT u.account, u.full_name , ls.created_at FROM api.login_session ls LEFT JOIN api.user u ON ls.user_id=u.id WHERE ls.deleted_at IS NULL AND  ls.created_at > CURRENT_TIMESTAMP - interval '1 day'"

	sqlLastDataImport = `
		SELECT m.metadataservice_name 
			, a.id
			, m.id
			, f.datafrequency
			, a.agency_shortname 
			, a.agency_name
			, ddl.import_begin_at 
		FROM   api.dataimport_dataset_log ddl 
   		INNER JOIN (SELECT dataimport_dataset_id
					  , Max(id) AS id 
			   FROM   api.dataimport_dataset_log ddl
			   GROUP  BY dataimport_dataset_id) ddl2 
		   	ON ddl.id = ddl2.id 
   			INNER JOIN public.metadata m 
		   	ON ddl.dataimport_dataset_id = m.dataimport_dataset_id 
   			INNER JOIN public.agency a 
			ON m.agency_id = a.id 
	   		INNER JOIN public.metadata_frequency f
			ON m.id = f.metadata_id
		WHERE  m.metadataservice_name IS NOT NULL
		ORDER  BY ddl.import_begin_at DESC`
)
