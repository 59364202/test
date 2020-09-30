package dba

import (
	"haii.or.th/api/server/model/cronjob"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName                  = "thaiwater30/backoffice/dba"
	ServiceVersion                   = service.APIVersion1
	SettingCreateYearlyPartitionCron = "thaiwater30.service.backoffice.dba.CreateYearlyPartitionCron"
	DefaultCreateYearlyPartitionCron = "0 0 25,26,27,28,29,30,31 12 *"
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
	dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handleDeleteData)

	setting.SetSystemDefault(SettingCreateYearlyPartitionCron, DefaultCreateYearlyPartitionCron)

	if _, err := cronjob.NewClusterFunc(SettingCreateYearlyPartitionCron,
		createYearlyPartitionCronJob); err != nil {
		log.Locationf("Can not add %s ...%v", SettingCreateYearlyPartitionCron, err)
	}

}

func createYearlyPartitionCronJob() error {
	return postYearlyPartition(0)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "delete_data_load":
		return srv.onLoadDeleteData(ctx)
	case "delete_data":
		return srv.getDeleteData(ctx)
	case "partition_table":
		return srv.getPartitionTable(ctx)
	case "partition_history":
		return srv.getPartitionHistory(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "partition":
		return srv.postPartition(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "delete_data":
		return srv.deleteDeleteData(ctx)
	case "partition":
		return srv.deletePartition(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
