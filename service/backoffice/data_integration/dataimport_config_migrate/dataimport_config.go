package dataimport_config

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/dataimport_config_migrate"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
	dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlePutData)
	dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handleDeleteData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "dataimport_download":
		return srv.getDataimportConfig(ctx)
	case "metadata":
		return srv.getDataimportConfigList(ctx)
	case "dataimport_dataset":
		return srv.getDataimportDatasetConfig(ctx)
	case "history_page":
		return srv.getDataimportHistoryDetails(ctx)
	case "history":
		return srv.getDataimportHistoryData(ctx)
	case "metadata_provision":
		return srv.getMetadataPrivision(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "dataimport_download":
		return srv.addDataimportDownloadConfig(ctx)
	case "dataimport_dataset":
		return srv.addDataimportDatasetConfig(ctx)
	case "dataimport_download_copy":
		return srv.copyDataimportDownloadConfig(ctx)
	case "dataimport_dataset_copy":
		return srv.copyDataimportDatasetConfig(ctx)
	case "metadata_provision":
		return srv.postMetadataProvision(ctx)	
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "dataimport_download":
		return srv.updateDataimportDownloadConfig(ctx)
	case "dataimport_dataset":
		return srv.updateDataimportDatasetConfig(ctx)
	case "metadata":
		return srv.updateMetadata(ctx)
	case "iscronenabled":
		return srv.cronEnabled(ctx)
	case "metadata_provision":
		return srv.putMetadataProvision(ctx)	
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}
	return nil
}

func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "dataimport_download":
		return srv.deleteDataimportDownloadConfig(ctx)
	case "dataimport_dataset":
		return srv.deleteDataimportDatasetConfig(ctx)
	case "metadata_provision":
		return srv.deleteMetadataProvision(ctx)	
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
