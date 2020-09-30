package metadata

import (
	model "haii.or.th/api/thaiwater30/model/hydroinfo"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_getHydroinfo struct {
	Result string                    `json:"result"` // example:`OK`
	Data   []*model.Struct_Hydroinfo `json:"data"`   // กลุ่มข้อมูลด้านน้ำและภูมิอากาศทั้งหมด
}

// @DocumentName	v1.webservice
//กลุ่มข้อมูลด้านน้ำและภูมิอากาศทั้งหมด @Service			thaiwater30/backoffice/metadata/hydroinfo
// @Summary			กลุ่มข้อมูลด้านน้ำและภูมิอากาศทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getHydroinfo successful operation
func (srv *HttpService) getHydroinfo(ctx service.RequestContext) error {
	resultData, err := model.GetAllHydroinfo()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}

	return nil
}

type Struct_postHydroInfo struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Inserted Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/hydroinfo
// @Summary			เพิ่มกลุ่มข้อมูลด้านน้ำและภูมิอากาศ
// @Method			POST
// @Consumes		json
// @Parameter		-	body	model.Struct_Hydroinfo_InputParam
// @Produces		json
// @Response		200	Struct_postHydroInfo successful operation
func (srv *HttpService) postHydroInfo(ctx service.RequestContext) error {
	param := &model.Struct_Hydroinfo_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	resultData, err := model.PostHydroInfo(ctx.GetUserID(), param)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(resultData)
	}

	return nil
}

type Struct_putHydroinfo struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Updated Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/hydroinfo
// @Summary			แก้ไขกลุ่มข้อมูลด้านน้ำและภูมิอากาศ
// @Method			PUT
// @Consumes		json
// @Parameter		-	body	model.Struct_Hydroinfo_InputParam
// @Produces		json
// @Response		200	Struct_putHydroinfo successful operation
func (srv *HttpService) putHydroinfo(ctx service.RequestContext) error {
	param := &model.Struct_Hydroinfo_InputParam{}
	if err := ctx.GetRequestParams(param); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(param)

	resultData, err := model.PutHydroinfo(ctx.GetUserID(), param)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}

	return nil
}

type Struct_deleteHydroinfo struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/hydroinfo
// @Summary			แก้ไขกลุ่มข้อมูลด้านน้ำและภูมิอากาศ
// @Method			DELETE
// @Parameter		id	query	int64	รหัสข้อมูลด้านเพื่อสนับสนุนการบริหารจัดการน้ำ
// @Produces		json
// @Response		200	Struct_deleteHydroinfo successful operation
func (srv *HttpService) deleteHydroinfo(ctx service.RequestContext) error {
	p := &model.Struct_Hydroinfo_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	resultData, err := model.DeleteHydroinfo(ctx.GetUserID(), p.ID)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(resultData))
	}

	return nil
}
