package tool

import (
	api_model "haii.or.th/api/server/model"
	model_setting "haii.or.th/api/server/model/setting"
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model "haii.or.th/api/thaiwater30/model/media"
	model_media_type "haii.or.th/api/thaiwater30/model/media_type"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	"strconv"
	"strings"
)

type Struct_OnLoadLastestImage struct {
	Agency    []*model_agency.Struct_Agency `json:"agency"`
	DateRange int64                         `json:"date_range"`
}

type Struct_DisplayImage struct {
	ImagePath string `json:"image_path"` //
}

type Struct_onLoadLastestImage struct {
	Agency    *Struct_onLoadLastestImage_Agency `json:"agency"`     // หน่วยงาน
	DateRange int64                             `json:"date_range"` // ช่วงวันที่ที่เลือกได้
}
type Struct_onLoadLastestImage_Agency struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`   // หน่วยงาน
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/lastest_image_load
// @Summary			เริ่มต้นหน้าตรวจสอบรูปภาพ
// @Method			GET
// @Produces		json
// @Response		200	Struct_onLoadLastestImage_Agency successful operation
func (srv *HttpService) onLoadLastestImage(ctx service.RequestContext) error {
	dataResult := &Struct_OnLoadLastestImage{}

	//Get DateRange Data
	dataResult.DateRange = model_setting.GetSystemSettingInt("bof.Tool.LastestImage.DateRange")
	if (dataResult.DateRange) == 0 {
		dataResult.DateRange = model_setting.GetSystemSettingInt("setting.Default.DateRange")
	}

	//Get List of AgencyInMediaTable Data
	resultAgency, err := model_agency.GetAgencyInMediaTable()
	if err != nil {
		return errors.Repack(err)
	}
	dataResult.Agency = resultAgency

	ctx.ReplyJSON(result.Result1(dataResult))
	return nil
}

type Struct_getLastestImage struct {
	Result string                `json:"result"` // example:`OK`
	Data   []*model.Struct_Media `json:"data"`   // ตารางตรวจสอบรูปภาพ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/lastest_image
// @Summary			ตารางตรวจสอบรูปภาพ
// @Method			GET
// @Parameter		-	query	model.Struct_Media_InputParam
// @Produces		json
// @Response		200	Struct_getLastestImage successful operation
func (srv *HttpService) getLastestImage(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_Media_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Get Data
	dataResult, err := model.GetMedia(p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

type Struct_getImageType struct {
	Result string                               `json:"result"` // example:`OK`
	Data   []*model_media_type.Struct_MediaType `json:"data"`   // ชนิดของภาพ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/tool/image_type
// @Summary			ชนิดของภาพทั้งหมด
// @Method			GET
// @Parameter		agency_id	query	string	example:`13` รหัสหน่วยงาน
// @Produces		json
// @Response		200	Struct_getImageType successful operation
func (srv *HttpService) getImageType(ctx service.RequestContext) error {
	//Map parameters
	p := &model.Struct_Media_InputParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Check id is not null
	if p.AgencyID == "" {
		dataResult, err := model_media_type.GetAllMediaType()
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		//Convert AgencyId type from string to int64
		intAgencyID, err := strconv.ParseInt(p.AgencyID, 10, 64)
		if err != nil {
			return errors.Repack(err)
		}

		dataResult, err := model_media_type.GetMediaTypeByAgency(intAgencyID)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	}

	return nil
}

func (srv *HttpService) displayImage(ctx service.RequestContext) error {
	var params struct {
		ImageID string `json:"image_id"`
	}
	ctx.GetRequestParams(&params)

	inf, err := api_model.GetCipher().DecryptText(params.ImageID)
	if err != nil {
		return rest.NewError(422, "invalid image id", err)
	}

	a := strings.Split(inf, ",")
	if len(a) != 2 {
		return rest.NewError(422, "invalid image id", nil)
	}

	fname := filepathx.JoinPath(api_model.GetStoragePath(), a[0], a[1])

	objResult := Struct_DisplayImage{}
	objResult.ImagePath = fname

	ctx.ReplyJSON(result.Result1(objResult))
	return nil
}
