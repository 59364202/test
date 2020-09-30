package impdata

import ()

var (
	getAgencyList = "SELECT a.id,a.agency_name->>'th' AS agency_name,md.dataimport_download_id,md.metadataservice_name->>'th' AS metadataservice_name, dd.download_script, md.id FROM public.agency a " +
		"LEFT JOIN public.metadata md ON a.id=md.agency_id LEFT JOIN api.dataimport_download dd ON md.dataimport_download_id=dd.id " +
		"WHERE md.dataimport_download_id IS NOT NULL AND md.dataimport_dataset_id IS NOT NULL ORDER BY agency_name,metadataservice_name"
)
