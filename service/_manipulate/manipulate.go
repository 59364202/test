package manipulate

import (
	model_amphoe "haii.or.th/api/thaiwater30/model/manipulate/amphoe"
	model_province "haii.or.th/api/thaiwater30/model/manipulate/province"
	model_region "haii.or.th/api/thaiwater30/model/manipulate/region"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/manipulate"
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
	case "region":
		return srv.getRegion(ctx)
	case "province":
		return srv.getProvince(ctx)
	case "amphoe":
		return srv.getAmphoe(ctx)

	case "hydroinfo":
		return srv.getHydroinfo(ctx)
	//==== Metadata ====//
	case "metadata_table":
		return srv.getMetadataTable(ctx)
	case "metadata":
		return srv.getMetadata(ctx)
	//=================//
	case "dataunit":
		return srv.getDataunit(ctx)
	case "servicemethod":
		return srv.getServiceMethod(ctx)
	case "dataformat":
		return srv.getDataformat(ctx)
	case "select_option_dataformat":
		return srv.getSelectOption(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePostData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {

	case "metadata":
		return srv.postMetadata(ctx)
	case "dataunit":
		return srv.postDatainit(ctx)
	case "servicemethod":
		return srv.postServiceMethod(ctx)
	case "hydroinfo":
		return srv.postHydroInfo(ctx)
	case "dataformat":
		return srv.postDataformat(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

func (srv *HttpService) handlePutData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {

	case "metadata":
		return srv.putMetadata(ctx)
	case "dataunit":
		return srv.putDataunit(ctx)
	case "servicemethod":
		return srv.putServiceMethod(ctx)
	case "hydroinfo":
		return srv.putHydroinfo(ctx)
	case "dataformat":
		return srv.putDataformat(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {

	case "metadata":
		srv.deleteMetadata(ctx)
	case "dataunit":
		return srv.deleteDataunit(ctx)

	case "servicemethod":
		return srv.deleteServiceMethod(ctx)
	case "hydroinfo":
		return srv.deleteHydroinfo(ctx)
	case "dataformat":
		return srv.deleteDataformat(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

// get

type Param struct {
	Region      string `json:"region_id"`
	Province    string `json:"province_id"`
	Amphoe      string `json:"amphoe_id"`
	Ministry    string `json:"ministry_id"`
	Category    string `json:"category_id"`
	Department  string `json:"department_id"`
	Subcategory string `json:"subcategory_id"`
	Agency      string `json:"agency_id"`
	Hydroinfo   string `json:"hydroinfo_id"`

	Text map[string]string `json:"text"`
}

// region
func (srv *HttpService) getRegion(ctx service.RequestContext) error {

	result, err := model_region.GetRegion()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// province
func (srv *HttpService) getProvince(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	result, err := model_province.GetProvince(p.Region)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// Amphoe
func (srv *HttpService) getAmphoe(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	result, err := model_amphoe.GetAmphoe(p.Province)
	if err != nil {
		return err
	} else {
		ctx.ReplyJSON(result)
	}
	return nil
}
