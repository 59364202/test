package metadata

import (
	"encoding/json"
	"fmt"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	model "haii.or.th/api/thaiwater30/model/lt_frequencyunit"
)

// FrequencyUnit //
type Param_postFrequencyUnit struct {
	FrequencyUnitName json.RawMessage `json:"frequencyunit_name,omitempty"` // example:`{"th": "ชั่วโมง"}` หน่วยความถี่
	ConvertMinute     string          `json:"convert_minute"`               // example: `60` แปลงค่าเป็นนาที
}

type Struct_getFrequencyUnit struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model.FrequencyUnit_struct `json:"data"`   // หน่วยความถี่ทั้งหมด
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/frequencyunit
// @Summary			หน่วยความถี่ทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getFrequencyUnit successful operation
func (srv *HttpService) getFrequencyUnit(ctx service.RequestContext) error {
	//Get Data
	rs, err := model.GetFrequencyUnit(ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_postFrequencyUnit struct {
	Result string                      `json:"result"` // example:`OK`
	Data   *model.FrequencyUnit_struct `json:"data"`   // หน่วยความถี่
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/frequencyunit
// @Summary			เพิ่มหน่วยความถี่
// @Method			POST
// @Consumes		json
// @Parameter		-	body	Param_postFrequencyUnit
// @Produces		json
// @Response		200	Struct_postFrequencyUnit successful operation
func (srv *HttpService) postFrequencyUnit(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_postFrequencyUnit{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Check input parameter
	if p.ConvertMinute == "" {
		ctx.ReplyError(fmt.Errorf("Can not get 'convert_minute'"))
		return nil
	}

	//Post Data
	rs, err := model.PostFrequencyUnit(ctx.GetUserID(), p.FrequencyUnitName, p.ConvertMinute)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putFrequencyUnit struct {
	Struct_postFrequencyUnit
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/frequencyunit/{id}
// @Summary			แก้ไขหน่วยความถี่
// @Method			PUT
// @Consumes		json
// @Parameter		id	path	string	รหัสหน่วยความถี่ example 1
// @Parameter		-	body	Param_postFrequencyUnit
// @Produces		json
// @Response		200	Struct_putFrequencyUnit successful operation
func (srv *HttpService) putFrequencyUnit(ctx service.RequestContext) error {
	//Map parameters
	p := &Param_postFrequencyUnit{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	//Check input parameter
	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get 'ID'"))
		return nil
	}
	if p.ConvertMinute == "" {
		ctx.ReplyError(fmt.Errorf("Can not get 'convert_minute'"))
		return nil
	}

	//Put Data
	rs, err := model.PutFrequencyUnit(ctx.GetUserID(), ctx.GetServiceParams("id"), p.FrequencyUnitName, p.ConvertMinute)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteFrequencyUnit struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/frequencyunit/{id}
// @Summary			ลบหน่วยความถี่
// @Method			DELETE
// @Parameter		id	path	string	รหัสหน่วยความถี่ example 1
// @Produces		json
// @Response		200	Struct_deleteFrequencyUnit successful operation
func (srv *HttpService) deleteFrequencyUnit(ctx service.RequestContext) error {

	//Check input parameter
	if ctx.GetServiceParams("id") == "" {
		ctx.ReplyError(fmt.Errorf("Can not get 'ID'"))
		return nil
	}

	//Delete Data
	rs, err := model.DeleteFrequencyUnit(ctx.GetUserID(), ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
