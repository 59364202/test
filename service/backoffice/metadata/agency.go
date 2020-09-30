package metadata

import (
	"encoding/json"
	"strconv"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	//	model "haii.or.th/api/thaiwater30/model/manipulate/agency"
	model_department "haii.or.th/api/thaiwater30/model/lt_department"
	model_ministry "haii.or.th/api/thaiwater30/model/lt_ministry"
)

// Agency //
type agencyParam struct {
	Id              string          `json:"id"`
	AgencyName      json.RawMessage `json:"agency_name"`
	AgencyShortName json.RawMessage `json:"agency_shortname"`
	DepartmentId    string          `json:"department_id"`
	MinistryId      string          `json:"ministry_id"`
}

type Struct_getAgency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงานทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency
// @Summary			หน่วยงานทั้งหมด
// @Method			GET
// @Parameter		department_id	query	string	required:false example:`24` รหัสกระทรวง
// @Parameter		ministry_id	query	string	required:false example:`2` รหัสกรม
// @Produces		json
// @Response		200	Struct_getAgency successful operation
func (srv *HttpService) getAgency(ctx service.RequestContext) error {
	//Map parameters
	p := &model_agency.Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	var rs []*model_agency.Struct_Agency
	var err error
	if p.DepartmentId != "" {
		departmentId, err := strconv.ParseInt(p.DepartmentId, 10, 64)
		if err != nil {
			return err
		}
		rs, err = model_agency.GetAgencyByDepartmentId(departmentId)
		if err != nil {
			return err
		}
	} else if p.MinistryId != "" {
		ministryId, err := strconv.ParseInt(p.MinistryId, 10, 64)
		if err != nil {
			return err
		}
		rs, err = model_agency.GetAgencyByMinistryId(ministryId)
		if err != nil {
			return err
		}
	} else {
		rs, err = model_agency.GetAllAgency()
		if err != nil {
			return nil
		}
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

type Struct_AgencyOnLoad struct {
	Ministry   *result.Result `json:"ministry"`
	Department *result.Result `json:"department"`
	Agency     *result.Result `json:"agency"`
}

type Struct_getAgencyOnLoad struct {
	Agency     *Struct_getAgencyOnLoad_Agency     `json:"agency"`     // หน่วยงาน
	Department *Struct_getAgencyOnLoad_Department `json:"department"` // กรม
	Ministry   *Struct_getAgencyOnLoad_Ministry   `json:"ministry"`   // กระทรวง
}
type Struct_getAgencyOnLoad_Ministry struct {
	Result string                            `json:"result"` // example:`OK`
	Data   []*model_ministry.Ministry_struct `json:"data"`   // กระทรวง
}
type Struct_getAgencyOnLoad_Department struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_department.Department_struct `json:"data"`   // กรม
}
type Struct_getAgencyOnLoad_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency_onload
// @Summary			เริ่มต้นหน้าหน่วยงาน
// @Method			GET
// @Produces		json
// @Response		200	Struct_getAgencyOnLoad	successful operation
func (srv *HttpService) getAgencyOnLoad(ctx service.RequestContext) error {
	//Map parameters
	p := &model_agency.Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs := &Struct_AgencyOnLoad{}
	rs_department, err := model_department.GetDepartment(make([]int, 0))
	if err != nil {
		rs.Department = result.Result0(err)
	} else {
		rs.Department = result.Result1(rs_department)
	}

	rs_ministry, err := model_ministry.GetMinistry()
	if err != nil {
		rs.Ministry = result.Result0(err)
	} else {
		rs.Ministry = result.Result1(rs_ministry)
	}

	rs_agency, err := model_agency.GetAllAgency()
	if err != nil {
		rs.Agency = result.Result0(err)
	} else {
		rs.Agency = result.Result1(rs_agency)
	}
	ctx.ReplyJSON(rs)

	//Get Agency

	return nil
}

type Struct_postAgency struct {
	Result string                     `json:"result"` // example:`OK`
	Data   model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency
// @Summary			เพิ่มหน่วยงาน
// @Method			POST
// @Consumes		json
// @Parameter		- body model_agency.Param_PostAgency
// @Produces		json
// @Response		200	Struct_postAgency successful operation
func (srv *HttpService) postAgency(ctx service.RequestContext) error {
	//Map parameters
	p := &model_agency.Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	p.UserId = ctx.GetUserID()
	//Post Agency
	//	rs, err := model.PostAgency(ctx.GetUserID(), p.AgencyName, p.AgencyShortName, p.DepartmentId)
	rs, err := model_agency.InsertAgency(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putAgency struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency
// @Summary			แก้ไขหน่วยงาน
// @Method			PUT
// @Consumes		json
// @Parameter		-	body	model_agency.Param_Agency
// @Produces		json
// @Response		200	Struct_putAgency successful operation
func (srv *HttpService) putAgency(ctx service.RequestContext) error {
	//Map parameters
	p := &model_agency.Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	p.UserId = ctx.GetUserID()
	//Put Agency
	//	rs, err := model.PutAgency(ctx.GetUserID(), ctx.GetServiceParams("id"), p.AgencyName, p.AgencyShortName, p.DepartmentId)
	rs, err := model_agency.UpdateAgency(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteAgency struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/agency/{id}
// @Summary			ลบหน่วยงาน
// @Method			DELETE
// @Parameter		id	path	string example:`1`	รหัสหน่วยงาน
// @Produces		json
// @Response		200	Struct_deleteAgency successful operation
func (srv *HttpService) deleteAgency(ctx service.RequestContext) error {
	//Map parameters
	p := &model_agency.Param_Agency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if ctx.GetServiceParams("id") != "" {
		p.Id = ctx.GetServiceParams("id")
	}
	p.UserId = ctx.GetUserID()
	//Delete Agency
	//	rs, err := model.DeleteAgency(ctx.GetUserID(), ctx.GetServiceParams("id"))
	rs, err := model_agency.DeleteAgency(p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
