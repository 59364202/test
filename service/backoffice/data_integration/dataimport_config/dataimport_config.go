package dataimport_config

import (
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
)

// @DocumentName 	v1.webservice
//
// @Module		thaiwater30/backoffice/dataimport_config
// @Description	Dataimport config management

const (
	DataServiceName = "thaiwater30/backoffice/dataimport_config"
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
	case "download_cron_list":
		return srv.getDownloadCronList(ctx)
	case "api_cron_list":
		return srv.getServerCronList(ctx)
	case "config_variable":
		return srv.getConfigVariable(ctx)
	case "list_category_variable":
		return srv.getListVariable(ctx)
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
	case "ps":
		return rdl.HandlePostData(ctx)
	case "config_variable":
		return srv.postConfigCategory(ctx)
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
	case "download_cron_list":
		return srv.updateDownloadCronList(ctx)
	case "api_cron_list":
		return srv.updateServerCronList(ctx)
	case "api_cron_run":
		return srv.runCron(ctx)
	case "config_variable":
		return srv.postConfigCategory(ctx)
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
	case "ps":
		return rdl.HandleDeleteData(ctx)
	case "config_variable":
		return srv.deleteAgency(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
