package dataimport_download_log



var sqlSelectOverAllDownloadCount_Online = ` SELECT mt.agency_id
											, mt.id AS metadata_id
											, mt.metadataagency_name::jsonb
											, mt.metadataservice_name::jsonb
											, mt.dataimport_download_id
											, mt.dataimport_dataset_id
											, (CASE WHEN mt.import_count IS NULL
													THEN 0
													ELSE mt.import_count END) $sqlCmdDateRange AS expected_download_count
											, SUM( CASE WHEN imd.download_record_count IS NULL
														THEN 0
														ELSE imd.download_record_count END) AS actual_download_count
											, SUM( CASE WHEN (CASE WHEN imd.download_record_count IS NULL THEN 0 ELSE imd.download_record_count END) > (CASE WHEN mt.import_count IS NULL THEN 0 ELSE mt.import_count END)
														THEN (CASE WHEN mt.import_count IS NULL THEN 0 ELSE mt.import_count END)
														ELSE (CASE WHEN imd.download_record_count IS NULL THEN 0 ELSE imd.download_record_count END) END) AS download_count
											, SUM( CASE WHEN imd.download_files_count IS NULL
														THEN 0
														ELSE imd.download_files_count END) AS download_files_count
											, SUM( CASE WHEN imd.import_record_count IS NULL
														THEN 0
														ELSE imd.import_record_count END) AS import_record_count
											, MAX(lastest_import) AS lastest_import
									FROM metadata mt
									LEFT JOIN ( SELECT dll.dataimport_download_id
													 , dsl.dataimport_dataset_id
													 , dll.download_begin_at::date AS download_date
													 , COUNT(*) AS download_record_count
													 , SUM(dll.download_files_count) AS download_files_count
													 , SUM(dsl.import_success_row) AS import_record_count
													 , MAX(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS lastest_import
											    FROM api.dataimport_download_log dll
											    LEFT JOIN api.dataimport_dataset_log dsl ON dll.id = dsl.dataimport_download_log_id
											    WHERE dll.download_path IS NOT NULL
											      $sqlCmdWhereOnline
											    GROUP BY dll.dataimport_download_id, dsl.dataimport_dataset_id, dll.download_begin_at::date) imd ON mt.dataimport_download_id = imd.dataimport_download_id AND mt.dataimport_dataset_id = imd.dataimport_dataset_id
									WHERE (mt.deleted_by IS NULL AND (mt.deleted_at IS NULL OR mt.deleted_at = '1970-01-01 07:00:00'))
									  -- AND mt.metadata_status = 'เชื่อมโยง'
									  AND mt.metadatastatus_id = 1
									  AND mt.connection_format = 'Online'
									GROUP BY mt.id, mt.agency_id, mt.import_count, mt.metadataagency_name::jsonb, mt.metadataservice_name::jsonb
									`

// -- Offline Metadata --
var sqlSelectOverAllDownloadCount_Offline = `	
SELECT mt.agency_id
		, mt.id AS metadata_id
		, mt.metadataagency_name::jsonb
		, mt.metadataservice_name::jsonb
		, mt.dataimport_download_id
		, mt.dataimport_dataset_id
		, (CASE WHEN import_count IS NULL
				THEN 0
				ELSE import_count END) AS expected_download_count
		, (CASE WHEN actual_download_count IS NULL
				THEN 0
				ELSE actual_download_count END) AS actual_download_count
		, (CASE WHEN (CASE WHEN actual_download_count IS NULL THEN 0 ELSE actual_download_count END) > (CASE WHEN import_count IS NULL THEN 0 ELSE import_count END)
				THEN (CASE WHEN import_count IS NULL THEN 0 ELSE import_count END)
				ELSE (CASE WHEN actual_download_count IS NULL THEN 0 ELSE actual_download_count END) END) AS download_count
		, (CASE WHEN download_files_count IS NULL
				THEN 0
				ELSE download_files_count END) AS download_files_count
		, (CASE WHEN import_record_count IS NULL
				THEN 0
				ELSE import_record_count END) AS import_record_count
		, lastest_import
FROM metadata mt
LEFT JOIN ( SELECT dll.dataimport_download_id
			, dsl.dataimport_dataset_id
			, COUNT(*) AS actual_download_count
			, SUM(dll.download_files_count) AS download_files_count
			, SUM(dsl.import_success_row) AS import_record_count
			, MAX(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS lastest_import
	FROM api.dataimport_download_log dll 
	LEFT JOIN api.dataimport_dataset_log dsl ON dll.id = dsl.dataimport_download_log_id
	WHERE dll.download_path IS NOT NULL
			$sqlCmdWhereOffline
		GROUP BY dll.dataimport_download_id, dsl.dataimport_dataset_id) sum_log ON mt.dataimport_download_id = sum_log.dataimport_download_id AND mt.dataimport_dataset_id = sum_log.dataimport_dataset_id
WHERE (mt.deleted_by IS NULL AND (mt.deleted_at IS NULL OR mt.deleted_at = '1970-01-01 07:00:00'))
-- AND mt.metadata_status = 'เชื่อมโยง'
AND mt.metadatastatus_id = 1
AND mt.connection_format = 'Offline' 
`

var sqlSumDownloadSizeByYear = `SELECT	extract(year FROM tall.date_time) AS date_time , 
SUM(tall.download_size_mb) AS download_size_mb , 
SUM(tall.download_files_count) AS download_files_count , 
SUM(tall.import_success_row) AS import_success_row
FROM
( SELECT    dts.date_series AS date_time , 
(CASE 
		  WHEN tdl.download_size_mb IS NULL THEN 0 
		  ELSE tdl.download_size_mb 
END) AS download_size_mb , 
 (CASE 
		  WHEN tdl.download_files_count IS NULL THEN 0 
		  ELSE tdl.download_files_count 
END) AS download_files_count , 
 (CASE 
		  WHEN tdl.import_success_row IS NULL THEN 0 
		  ELSE tdl.import_success_row 
END) AS import_success_row 
FROM      ( 
	   SELECT Generate_series($3::DATE, ($3::DATE + '12 MONTH'::interval - '1 DAY'::interval)::DATE, '1 day')::DATE AS date_series) dts
left join 
( 
		  SELECT    dll.download_begin_at::                 DATE    AS download_date ,
					SUM(dll.download_bytes_count)/ 1048576::NUMERIC AS download_size_mb
					-- / 1073741824::numeric AS download_size_gb 
					, 
					SUM(dll.download_files_count) AS download_files_count , 
					SUM(dsl.import_success_row)   AS import_success_row 
		  FROM      api.dataimport_download_log dll 
		  left join api.dataimport_dataset_log dsl 
		  ON        dll.id = dsl.dataimport_download_log_id 
		  left join metadata mt 
		  ON        mt.dataimport_download_id = dll.dataimport_download_id 
		  AND       mt.dataimport_dataset_id = dsl.dataimport_dataset_id 
		  left join agency agt 
		  ON        agt.id = mt.agency_id 
		  WHERE     dll.download_path IS NOT NULL 
		  AND       extract(month FROM dll.download_begin_at) BETWEEN 1 AND 12 
		  AND       extract(year FROM dll.download_begin_at) = $2
					-- AND mt.metadata_status = 'เชื่อมโยง' 
		  AND       mt.metadatastatus_id = 1 
		  AND       mt.agency_id = $1 
		  GROUP BY  dll.download_begin_at::DATE ) tdl 
ON        dts.date_series = tdl.download_date 
ORDER BY  dts.date_series ) tall
GROUP BY extract(year FROM tall.date_time)`

var sqlSelectMonthlyDownloadSizeByAgency = ` SELECT dts.date_series
												   , (CASE WHEN tdl.download_size_mb IS NULL THEN 0 ELSE tdl.download_size_mb END) AS download_size_mb
												   , (CASE WHEN tdl.download_files_count IS NULL THEN 0 ELSE tdl.download_files_count END) AS download_files_count
												   , (CASE WHEN tdl.import_success_row IS NULL THEN 0 ELSE tdl.import_success_row END) AS import_success_row
											  FROM ( SELECT GENERATE_SERIES($4::DATE, ($4::DATE + '1 MONTH'::INTERVAL - '1 DAY'::INTERVAL)::DATE, '1 day')::DATE AS date_series) dts
											  LEFT JOIN ( SELECT dll.download_begin_at::date AS download_date
																, SUM(dll.download_bytes_count)/ 1048576::numeric AS download_size_mb
																-- / 1073741824::numeric AS download_size_gb
																, SUM(dll.download_files_count) AS download_files_count
																, SUM(dsl.import_success_row) AS import_success_row
														  FROM api.dataimport_download_log dll
														  LEFT JOIN api.dataimport_dataset_log dsl ON dll.id = dsl.dataimport_download_log_id
														  LEFT JOIN metadata mt ON mt.dataimport_download_id = dll.dataimport_download_id AND mt.dataimport_dataset_id = dsl.dataimport_dataset_id
														  LEFT JOIN agency agt ON agt.id = mt.agency_id
														  WHERE dll.download_path IS NOT NULL
														    AND EXTRACT(MONTH FROM dll.download_begin_at) = $1
														    AND EXTRACT(YEAR FROM dll.download_begin_at) = $2
														    -- AND mt.metadata_status = 'เชื่อมโยง'
														    AND mt.metadatastatus_id = 1
														    AND mt.agency_id = $3
														  GROUP BY dll.download_begin_at::date ) tdl ON dts.date_series = tdl.download_date
											  ORDER BY dts.date_series `

//var sqlSelectCompareDownloadCount = ` 	SELECT download_year
//												, download_month
//												, SUM(expected_download_count) AS expected_download_count
//												, SUM(actual_download_count) AS actual_download_count
//												, SUM(download_count) AS download_count
//												, SUM(download_files_count) AS download_files_count
//												, SUM(import_record_count) AS import_record_count
//												, CASE WHEN SUM(expected_download_count) = 0 THEN 0 ELSE (SUM(download_count)/SUM(expected_download_count))*100 END AS percent_download
//										FROM(
//											) aa
//									GROUP BY download_year, download_month
//									ORDER BY download_year, download_month `
var sqlSelectCompareDownloadCount_Online = `
SELECT mt.id AS metadata
		, mt.connection_format
		, EXTRACT('month' FROM tmp_date.dt_date) AS download_month
		, EXTRACT('year' FROM tmp_date.dt_date) AS download_year
		, SUM(CASE WHEN mt.import_count IS NULL THEN 0 ELSE mt.import_count END) AS expected_download_count
		, SUM(CASE WHEN imd.download_record_count IS NULL THEN 0 ELSE imd.download_record_count END) AS actual_download_count
		, SUM(CASE WHEN (CASE WHEN imd.download_record_count IS NULL THEN 0 ELSE imd.download_record_count END) > (CASE WHEN mt.import_count IS NULL THEN 0 ELSE mt.import_count END)
			THEN (CASE WHEN mt.import_count IS NULL THEN 0 ELSE mt.import_count END)
			ELSE (CASE WHEN imd.download_record_count IS NULL THEN 0 ELSE imd.download_record_count END)
			END) AS download_count
		, SUM(CASE WHEN imd.download_files_count IS NULL THEN 0 ELSE imd.download_files_count END)AS download_files_count
		, SUM(CASE WHEN imd.import_record_count IS NULL THEN 0 ELSE imd.import_record_count END) AS import_record_count
		, MAX(lastest_import) AS lastest_import
FROM (	$sqlCmdSelectTempDate ) tmp_date
CROSS JOIN metadata mt
LEFT JOIN (	SELECT dll.dataimport_download_id
				 , dsl.dataimport_dataset_id
				 , dll.download_begin_at::date AS download_date
				 , COUNT(*) AS download_record_count
				 , SUM(dll.download_files_count) AS download_files_count
				 , SUM(dsl.import_success_row) AS import_record_count
				 --, MIN(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS first_import
				 , MAX(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS lastest_import
			FROM api.dataimport_download_log dll
			LEFT JOIN api.dataimport_dataset_log dsl ON dll.id = dsl.dataimport_download_log_id
			WHERE dll.download_path IS NOT NULL
			  $sqlCmdWhere
			GROUP BY dll.dataimport_download_id, dsl.dataimport_dataset_id, dll.download_begin_at::date
		) imd ON mt.dataimport_download_id = imd.dataimport_download_id AND mt.dataimport_dataset_id = imd.dataimport_dataset_id AND tmp_date.dt_date = imd.download_date
WHERE (mt.deleted_by IS NULL AND (mt.deleted_at IS NULL OR mt.deleted_at = '1970-01-01 07:00:00'))
  -- AND mt.metadata_status = 'เชื่อมโยง'
  AND mt.metadatastatus_id = 1
  AND mt.agency_id = $1
  AND mt.connection_format = 'Online'
GROUP BY mt.id, mt.connection_format, EXTRACT('month' FROM tmp_date.dt_date), EXTRACT('year' FROM tmp_date.dt_date)
`
var sqlSelectCompareDownloadCount_Offline = `
SELECT mt.id AS metadata
	, mt.connection_format
	, tmp_date.month AS download_month
	, tmp_date.year AS download_year
	, (CASE WHEN import_count IS NULL
			THEN 0
			ELSE import_count END) AS expected_download_count
	, (CASE WHEN actual_download_count IS NULL
			THEN 0
			ELSE actual_download_count END) AS actual_download_count
	, (CASE WHEN (CASE WHEN actual_download_count IS NULL THEN 0 ELSE actual_download_count END) > (CASE WHEN import_count IS NULL THEN 0 ELSE import_count END)
			THEN (CASE WHEN import_count IS NULL THEN 0 ELSE import_count END)
			ELSE (CASE WHEN actual_download_count IS NULL THEN 0 ELSE actual_download_count END) END) AS download_count
	, (CASE WHEN download_files_count IS NULL
			THEN 0
			ELSE download_files_count END) AS download_files_count
	, (CASE WHEN import_record_count IS NULL
			THEN 0
			ELSE import_record_count END) AS import_record_count
	, lastest_import												
FROM ( SELECT * FROM ($sqlCmdSelectTempDateOffline) y CROSS JOIN ( SELECT generate_series(1, 12, 1) as month ) m ) tmp_date
CROSS JOIN metadata mt
LEFT JOIN (	SELECT dll.dataimport_download_id
			, dsl.dataimport_dataset_id
			, dll.download_begin_at::date AS download_date
			, COUNT(*) AS actual_download_count
			, SUM(dll.download_files_count) AS download_files_count
			, SUM(dsl.import_success_row) AS import_record_count
			--, MIN(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS first_import
			, MAX(CASE WHEN dsl.import_success_row > 0 THEN dsl.updated_at ELSE NULL END) AS lastest_import
	FROM api.dataimport_download_log dll
	LEFT JOIN api.dataimport_dataset_log dsl ON dll.id = dsl.dataimport_download_log_id
	WHERE dll.download_path IS NOT NULL
			$sqlCmdWhereOffline
		GROUP BY dll.dataimport_download_id, dsl.dataimport_dataset_id, dll.download_begin_at::date
	) imd ON mt.dataimport_download_id = imd.dataimport_download_id AND mt.dataimport_dataset_id = imd.dataimport_dataset_id 
			AND EXTRACT(YEAR FROM imd.download_date) = tmp_date.year AND EXTRACT(MONTH FROM imd.download_date) = tmp_date.month 
WHERE (mt.deleted_by IS NULL AND (mt.deleted_at IS NULL OR mt.deleted_at = '1970-01-01 07:00:00'))
-- AND mt.metadata_status = 'เชื่อมโยง'
AND mt.metadatastatus_id = 1
AND mt.connection_format= 'Offline'
AND mt.agency_id = $1		
`
