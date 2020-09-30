package metadata

import (
	model_metadata_status "haii.or.th/api/thaiwater30/model/metadata_status"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_getMetadataStatus struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   []*model_metadata_status.MetadataStatusParams `json:"data"`   // สถานะบัญชีข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_status
// @Summary			สถานะบัญชีข้อมูลทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getMetadataStatus successful operation
func (srv *HttpService) getMetadataStatus(ctx service.RequestContext) error {
	rs, err := model_metadata_status.GetMetadataStatus()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_postMetadataStatus struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`20` รหัสสถานะบัญชีข้อมูล ที่เพิ่มเข้าไป
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_status
// @Summary			เพิ่มสถานะบัญชีข้อมูล
// @Method			POST
// @Consumes		json
// @Parameter		- body model_metadata_status.DataMetadataStatusPostParams
// @Produces		json
// @Response		200	Struct_postMetadataStatus successful operation
func (srv *HttpService) postMetadataStatus(ctx service.RequestContext) error {
	p := &model_metadata_status.DataMetadataStatusParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_status.PostMetadataStatus(ctx.GetUserID(), p.StatusName)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putMetadataStatus struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`20` รหัสสถานะบัญชีข้อมูล ที่แก้ไข
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_status
// @Summary			แก้ไขสถานะบัญชีข้อมูล
// @Method			PUT
// @Consumes		json
// @Parameter		-	body	model_metadata_status.DataMetadataStatusParams
// @Produces		json
// @Response		200	Struct_putMetadataStatus successful operation
func (srv *HttpService) putMetadataStatus(ctx service.RequestContext) error {
	p := &model_metadata_status.DataMetadataStatusParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_status.PutMetadataStatus(ctx.GetUserID(), p.StatusID, p.StatusName)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteMetadatastatus struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_status
// @Summary			แก้ไขสถานะบัญชีข้อมูล
// @Method			DELETE
// @Parameter		id	query	string	 รหัสสถานะของบัญชีข้อมูลแบบเข้ารหัส example IJPRn6cYcJZ3KFrDxEEnnJ4d8_FlXM1X_l5oN4AnPxY3MvRzkxDEZsx-ZZ48swiGmrUxi40TGQp0T-VlwmtN-g
// @Produces		json
// @Response		200	Struct_deleteMetadatastatus successful operation
func (srv *HttpService) deleteMetadatastatus(ctx service.RequestContext) error {
	p := &model_metadata_status.DataMetadataStatusParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_status.DelMetadataStatus(ctx.GetUserID(), p.StatusID)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
