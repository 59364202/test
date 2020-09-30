package shared

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_canal_station "haii.or.th/api/thaiwater30/model/canal_station"
	model_lt_department "haii.or.th/api/thaiwater30/model/lt_department"
	model_lt_subcategory "haii.or.th/api/thaiwater30/model/lt_subcategory"
	model_tele_station "haii.or.th/api/thaiwater30/model/tele_station"
	model_waterquality_station "haii.or.th/api/thaiwater30/model/waterquality_station"
)

const (
	DataServiceName = "thaiwater30/frontend/shared"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")
	switch service_id {
	case "station_all":
		return srv.getAllStation(ctx)
	//		m_tele_station R,A
	case "station":
		return srv.getRainStation(ctx)
		//		m_tele_station W,A + m_canal_station
	case "tele_canal_station":
		return srv.getTeleCanalStation(ctx)
	case "subcategory":
		return srv.getSubcategory(ctx)
	case "department":
		return srv.getDepartment(ctx)
	case "agency":
		return srv.getAgency(ctx)
	case "waterquality_station":
		return srv.getWaterqualityStation(ctx)
	case "watergate_station":
		return srv.getWaterGateStation(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}

type Param struct {
	Category_Id   int64  `json:"category_id"`
	Ministry_id   int64  `json:"ministry_id"`
	Province_Id   string `json:"province_id"`
	Province_Code string `json:"province_code"`
}

func (srv *HttpService) getSubcategory(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	var (
		rs  []*model_lt_subcategory.Struct_subcategory
		err error
	)
	if p.Category_Id != 0 {
		rs, err = model_lt_subcategory.GetSubcategoryFromCategoryId(p.Category_Id)

	} else {
		rs, err = model_lt_subcategory.GetAllSubcategory()
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

func (srv *HttpService) getDepartment(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	var (
		rs  []*model_lt_department.Struct_Department
		err error
	)
	if p.Ministry_id != 0 {
		rs, err = model_lt_department.GetDepartmentFromMinistryId(p.Ministry_id)
	} else {
		rs, err = model_lt_department.GetAllDepartment()
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getAgency(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	var (
		rs  []*model_agency.Struct_Agency
		err error
	)
	if p.Ministry_id != 0 {
		rs, err = model_agency.GetAgencyByMinistryId(p.Ministry_id)
	} else {
		rs, err = model_agency.GetAllAgency()
	}

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_RainStation struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_tele_station.Struct_Station `json:"data"`   // ข้อมูลสถานี
}

// @DocumentName	v1.public
// @Service			thaiwater30/frontend/shared/station
// @Summary 		สถานีฝนรายจังหวัด
// @Parameter		province_code	query string	example: 10 รหัสจังหวัด
// @Method			GET
// @Produces		json
// @Response		200		Struct_RainStation	successful operation
func (srv *HttpService) getRainStation(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	rs_station, err := model_tele_station.GetTeleStationFromProvinceCode(p.Province_Code, "R,A")
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs_station))
	}

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/frontend/shared/station_all
// @Summary 		สถานีทั้งหมดรายจังหวัด
// @Parameter		province_code	query string	example: 10 รหัสจังหวัด
// @Method			GET
// @Produces		json
// @Response		200		Struct_RainStation	successful operation
func (srv *HttpService) getAllStation(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	rs_station, err := model_tele_station.GetTeleStationFromProvinceCode(p.Province_Code, "")
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs_station))
	}

	return nil
}

type Struct_Station struct {
	Tele  []*model_tele_station.Struct_Station       `json:"tele_waterlevel"`  // สถานีโทรมาตร
	Canal []*model_canal_station.Struct_CanalStation `json:"canal_waterlevel"` // สถานีคลอง
}
type Struct_TeleCanalStation struct {
	Result string          `json:"result"` // example:`OK`
	Data   *Struct_Station `json:"data"`   // ข้อมูลสถานี
}

// @DocumentName	v1.public
// @Service			thaiwater30/frontend/shared/tele_canal_station
// @Summary			สถานีระดับน้ำรายจังหวัด
// @Parameter		province_code	query string	example: 10 รหัสจังหวัด
// @Method			GET
// @Produces		json
// @Response		200		Struct_TeleCanalStation	successful operation
func (srv *HttpService) getTeleCanalStation(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	var (
		rs  *Struct_Station = new(Struct_Station)
		err error
	)

	rs_station, err_station := model_tele_station.GetTeleStationFromProvinceCode(p.Province_Code, "W,A")
	if err_station != nil {
		err = err_station
	} else {
		rs.Tele = make([]*model_tele_station.Struct_Station, 0)
		rs.Tele = append(rs.Tele, rs_station...)
	}
	rs_canal, err_canal := model_canal_station.GetCanalStationFromProvinceCode(p.Province_Code)
	if err_canal != nil {
		err = err_canal
	} else {
		rs.Canal = make([]*model_canal_station.Struct_CanalStation, 0)
		rs.Canal = append(rs.Canal, rs_canal...)
	}

	if err != nil {
		return err
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_WaterQualityStation struct {
	Result string                                       `json:"result"` // example:`OK`
	Data   []*model_waterquality_station.Struct_Station `json:"data"`   // ข้อมูลสถานี
}

// @DocumentName	v1.public
// @Service			thaiwater30/frontend/shared/waterquality_station
// @Summary			สถานีคุณภาพน้ำรายจังหวัด
// @Parameter		province_code	query string	example: 10 รหัสจังหวัด
// @Method			GET
// @Produces		json
// @Response		200		Struct_WaterQualityStation	successful operation
func (srv *HttpService) getWaterqualityStation(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	rs, err := model_waterquality_station.Get_AllWaterQualityStaion_From_ProvinceCode(p.Province_Code)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_WaterGateStation struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_tele_station.Struct_Station `json:"data"`   // ข้อมูลสถานี
}

// @DocumentName	v1.public
// @Service			thaiwater30/frontend/shared/watergate_station
// @Summary			สถานีวัดระดับน้ำปตร./ฝาย รายจังหวัด
// @Parameter		province_code	query string	example: 10 รหัสจังหวัด
// @Method			GET
// @Produces		json
// @Response		200		Struct_WaterGateStation	successful operation
func (srv *HttpService) getWaterGateStation(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_tele_station.GetWeirStationFromProvinceCode(p.Province_Code)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
