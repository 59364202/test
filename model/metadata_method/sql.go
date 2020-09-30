package metadata_method

import ()

var (
	getMetadataMethod = "SELECT id, metadata_method_name FROM public.lt_metadata_method WHERE deleted_at=to_timestamp(0)"

	insMetadataMethod = "INSERT INTO public.lt_metadata_method(metadata_method_name, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at) VALUES ($1, $2, NOW(), $2, NOW(), $2, to_timestamp(0)) RETURNING id"

	updateMetadataMethod = "UPDATE public.lt_metadata_method SET metadata_method_name=$3 ,updated_by=$1 ,updated_at=NOW() WHERE id=$2"

	delMetadataMethod = "UPDATE public.lt_metadata_method SET deleted_by=$1 ,deleted_at=NOW() WHERE id=$2"
)
