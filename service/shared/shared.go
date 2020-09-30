package shared

import (
	"encoding/json"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/shared"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

type SharedParam struct {
	DataType string          `json:"data_type"`
	Name     string          `json:"name"`
	Data     json.RawMessage `json:"data"`
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
	//dpt.Register(ServiceVersion, service.MethodPOST, DataServiceName, srv.handlePostData)
	//dpt.Register(ServiceVersion, service.MethodPUT, DataServiceName, srv.handlePutData)
	//dpt.Register(ServiceVersion, service.MethodDELETE, DataServiceName, srv.handleDeleteData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "station":
		//Map parameters
		p := &SharedParam{}
		if err := ctx.GetRequestParams(p); err != nil {
			rest.NewError(404, "mismatch param name", nil)
		}
		ctx.LogRequestParams(p)

		switch p.DataType {
		case "rainfall":
			return srv.getTeleStationByDataType(ctx, p)
		case "dam_daily":
			return srv.getDamStationByDataType(ctx, p)
		case "dam_hourly":
			return srv.getDamStationByDataType(ctx, p)
		case "tele_waterlevel":
			return srv.getTeleStationByDataType(ctx, p)
		default:
			return rest.NewError(404, "mismatch param name", nil)
		}
	case "setting":
		//Map parameters
		p := &SharedParam{}
		if err := ctx.GetRequestParams(p); err != nil {
			rest.NewError(404, "mismatch param name", nil)
		}
		ctx.LogRequestParams(p)

		srv.getSystemSetting(ctx, p)
	/*
		switch p.Name {
		case "data_type":
			return srv.getSystemSetting(ctx, p)
		default:
			return rest.NewError(404, "mismatch param name", nil)
		}*/
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
