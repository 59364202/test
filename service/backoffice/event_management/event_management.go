package event_management

import (
	//"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/backoffice/event_management"
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
	case "event_load":
		return srv.getEvent(ctx)
	case "event":
		return srv.getEvent(ctx)
	case "subevent_load":
		return srv.onLoadSubEvent(ctx)
	case "subevent":
		return srv.getSubEvent(ctx)
	case "sink_method_load":
		return srv.onLoadSinkMethod(ctx)
	case "sink_method":
		return srv.getSinkMethod(ctx)
	case "email_server":
		return srv.getEmailServer(ctx)
	case "email_template":
		return srv.getEmailTemplate(ctx)
	case "sink_condition":
		return srv.getSinkCondition(ctx)
	case "sink_condition_option":
		return srv.getSinkConditionSelectOption(ctx)
	case "sink_target":
		return srv.getSinkTarget(ctx)
	case "sink_target_option":
		return srv.getSinkTargetSelectOption(ctx)
	case "lt_sink_method":
		return srv.getLtSinkMethodType(ctx)
	case "sms_server":
		return srv.getSmsServer(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "event":
		return srv.postEvent(ctx)
	case "subevent":
		return srv.postSubEvent(ctx)
	case "sink_method":
		return srv.postSinkMethod(ctx)
	case "sink_method_system_setting":
		return srv.postSinkMethodSystemSetting(ctx)
	case "email_server":
		return srv.postEmailServer(ctx)
	case "email_template":
		return srv.postEmailTemplate(ctx)
	case "sink_condition":
		return srv.postSinkCondition(ctx)
	case "sink_target":
		return srv.postSinkTarget(ctx)
	case "lt_sink_method":
		return srv.postLtSinkMethodType(ctx)
	case "sms_server":
		return srv.postSmsServer(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "event":
		return srv.putEvent(ctx)
	case "subevent":
		return srv.putSubEvent(ctx)
	case "sink_method":
		return srv.putSinkMethod(ctx)
	case "sink_method_system_setting":
		return srv.putSinkMethodSystemSetting(ctx)
	case "email_server":
		return srv.putEmailServer(ctx)
	case "email_template":
		return srv.putEmailTemplate(ctx)
	case "sink_condition":
		return srv.putSinkCondition(ctx)
	case "sink_target":
		return srv.putSinkTarget(ctx)
	case "lt_sink_method":
		return srv.putLtSinkMethodType(ctx)
	case "sms_server":
		return srv.putSmsServer(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "event":
		return srv.deleteEvent(ctx)
	case "subevent":
		return srv.deleteSubEvent(ctx)
	case "sink_method":
		return srv.deleteSinkMethod(ctx)
	case "email_server":
		return srv.deleteEmailServer(ctx)
	case "email_template":
		return srv.deletedEmailTemplate(ctx)
	case "sink_condition":
		return srv.deletedSinkCondition(ctx)
	case "sink_target":
		return srv.deletedSinkTarget(ctx)
	case "sms_server":
		return srv.deleteSmsServer(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
