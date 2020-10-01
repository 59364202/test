package metadata

import (
	//"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	//	"haii.or.th/api/server/model/cronjob"
	//	"haii.or.th/api/server/model/setting"
	//
	//	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	//	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	//	model_rainfall24hr "haii.or.th/api/thaiwater30/model/rainfall24hr"
	//	model_tele_waterlevel "haii.or.th/api/thaiwater30/model/tele_waterlevel"
	//	model_waterquality "haii.or.th/api/thaiwater30/model/waterquality"
)

const (
	DataServiceName = "thaiwater30/backoffice/metadata"
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

	// register geocode service
	dpt.Register(ServiceVersion, service.MethodGET, "thaiwater30/backoffice/metadata/geocode", srv.getGeocodeFromLatLon)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	//==== Metadata ====//
	case "metadata_load":
		return srv.onLoadMetadata(ctx)
	case "metadata_table":
		return srv.getMetadataTable(ctx)
	case "metadata":
		return srv.getMetadata(ctx)
	//=================//
	case "summary_load":
		return srv.onLoadSummaryMetadata(ctx)
	case "summary":
		return srv.getSummaryMetadata(ctx)
	case "summary_by_agency":
		return srv.getSummaryMetadataImported(ctx)
	case "summary_by_category":
		return srv.getSummaryMetadataImported(ctx)
	case "summary_detail":
		return srv.getMetadataImported(ctx)
	case "category":
		return srv.getCategory(ctx)
	case "subcategory":
		return srv.getSubCategory(ctx)
	case "frequencyunit":
		return srv.getFrequencyUnit(ctx)
	case "ministry":
		return srv.getMinistry(ctx)
	case "department":
		return srv.getDepartment(ctx)
	case "agency":
		return srv.getAgency(ctx)
	case "agency_onload":
		return srv.getAgencyOnLoad(ctx)
	case "metadata_method":
		return srv.getMetadataMethod(ctx)
	case "metadata_status":
		return srv.getMetadataStatus(ctx)
	case "servicemethod":
		return srv.getServicemethod(ctx)
	case "hydroinfo":
		return srv.getHydroinfo(ctx)
	case "dataformat":
		return srv.getDataformat(ctx)
	case "select_option_dataformat":
		return srv.getSelectOption(ctx)
	case "dataunit":
		return srv.getDataunit(ctx)
	case "metadata_show":
		return srv.getMetadataShow(ctx)
	case "show_system":
		return srv.getMetadataShowSystem(ctx)
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
	case "agency":
		return srv.postAgency(ctx)
	case "category":
		return srv.postCategory(ctx)
	case "subcategory":
		return srv.postSubCategory(ctx)
	case "frequencyunit":
		return srv.postFrequencyUnit(ctx)
	case "department":
		return srv.postDepartment(ctx)
	case "ministry":
		return srv.postMinistry(ctx)
	case "metadata_method":
		return srv.postMetadataMethod(ctx)
	case "metadata_status":
		return srv.postMetadataStatus(ctx)
	case "servicemethod":
		return srv.postServiceMethod(ctx)
	case "hydroinfo":
		return srv.postHydroInfo(ctx)
	case "dataformat":
		return srv.postDataformat(ctx)
	case "dataunit":
		return srv.postDatainit(ctx)
	case "metadata_show":
		return srv.postMetadataShow(ctx)
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
	case "agency":
		return srv.putAgency(ctx)
	case "category":
		return srv.putCategory(ctx)
	case "subcategory":
		return srv.putSubCategory(ctx)
	case "frequencyunit":
		return srv.putFrequencyUnit(ctx)
	case "department":
		return srv.putDepartment(ctx)
	case "ministry":
		return srv.putMinistry(ctx)
	case "metadata_method":
		return srv.putMetadataMethod(ctx)
	case "metadata_status":
		return srv.putMetadataStatus(ctx)
	case "servicemethod":
		return srv.putServiceMethod(ctx)
	case "hydroinfo":
		return srv.putHydroinfo(ctx)
	case "offlinedate":
		return srv.putMetadataOfflineDate(ctx)
	case "dataformat":
		return srv.putDataformat(ctx)
	case "dataunit":
		return srv.putDataunit(ctx)
	case "metadata_show":
		return srv.putMetadataShow(ctx)
	case "upload_img":
		return srv.putLogoAgency(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
func (srv *HttpService) handleDeleteData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "metadata":
		return srv.deleteMetadata(ctx)
	case "agency":
		return srv.deleteAgency(ctx)
	case "category":
		return srv.deleteCategory(ctx)
	case "subcategory":
		return srv.deleteSubCategory(ctx)
	case "frequencyunit":
		return srv.deleteFrequencyUnit(ctx)
	case "department":
		return srv.deleteDepartment(ctx)
	case "ministry":
		return srv.deleteMinistry(ctx)
	case "metadata_method":
		return srv.deleteMetadataMethod(ctx)
	case "metadata_status":
		return srv.deleteMetadatastatus(ctx)
	case "servicemethod":
		return srv.deleteServiceMethod(ctx)
	case "hydroinfo":
		return srv.deleteHydroinfo(ctx)
	case "dataformat":
		return srv.deleteDataformat(ctx)
	case "dataunit":
		return srv.deleteDataunit(ctx)
	case "metadata_show":
		return srv.deleteMetadataShow(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

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
