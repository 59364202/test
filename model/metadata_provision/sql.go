package metadata_provision

import ()

var (
	sqlSelectMetadataProvision = "SELECT md.id, md.metadataservice_name, dd.download_name, dds.convert_name " +
		"FROM public.metadata_provision md " +
		"LEFT JOIN api.dataimport_download dd ON md.dataimport_download_id=dd.id " +
		"LEFT JOIN api.dataimport_dataset dds ON md.dataimport_dataset_id=dds.id " +
		"WHERE md.deleted_at=to_timestamp(0)"
		
	sqlInsMetadataProvision = "INSERT INTO public.metadata_provision(metadataservice_name, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at) VALUES ($1, $2, NOW(), $2, NOW(), $2, to_timestamp(0)) RETURNING id"

	sqlUpdateMetadataMethod = "UPDATE public.metadata_provision SET metadataservice_name=$3 ,updated_by=$1 ,updated_at=NOW() WHERE id=$2"

	sqlDelMetadataMethod = "UPDATE public.metadata_provision SET deleted_by=$1 ,deleted_at=NOW() WHERE id=$2"

)
