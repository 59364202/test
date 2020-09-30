package metadata

import (
	model_metadata_method "haii.or.th/api/thaiwater30/model/metadata_method"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_getMetadataMethod struct {
	Result string                                        `json:"result"` // example:`OK`
	Data   []*model_metadata_method.MetadataMethodParams `json:"data"`   // วิธีการได้มาของข้อมูล
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_method
// @Summary			วิธีการได้มาซึ่งข้อมูลทั้งหมด
// @Method			GET
// @Produces		json
// @Response		200	Struct_getMetadataMethod successful operation
func (srv *HttpService) getMetadataMethod(ctx service.RequestContext) error {
	rs, err := model_metadata_method.GetMetadataMethod()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_postMetadataMethod struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`20` รหัสวิธีการได้มาของข้อมูล ที่เพิ่มเข้าไป
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_method
// @Summary			เพิ่มวิธีการได้มาซึ่งข้อมูล
// @Method			POST
// @Parameter		metadata_method_name	query	string	example: "Web Service"  ชื่อวิธีการได้มาซึ่งข้อมูล
// @Produces		json
// @Response		200	Struct_postMetadataMethod successful operation
func (srv *HttpService) postMetadataMethod(ctx service.RequestContext) error {
	p := &model_metadata_method.MetadataMethodParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_method.PostMetadataMethod(ctx.GetUserID(), p.MethodName)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putMetadataMethod struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`20` รหัสวิธีการได้มาของข้อมูล ที่แก้ไข
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_method/{id}
// @Summary			แก้ไขวิธีการได้มาซึ่งข้อมูล
// @Method			PUT
// @Parameter		id	path	string	example: tci3CPw5YTdhuV9PuN8oh3e-I9GySj9SZmpYRIbct4iUiSSjyOqsVPewcVLgVHgDDtmzcX35dU5bkjt4qtzD_g  รหัสวิธีการได้มาของข้อมูลแบบเข้ารหัส
// @Parameter		metadata_method_name	form	string	example: "Web Service"  ชื่อวิธีการได้มาซึ่งข้อมูล
// @Produces		json
// @Response		200	Struct_putMetadataMethod successful operation
func (srv *HttpService) putMetadataMethod(ctx service.RequestContext) error {
	p := &model_metadata_method.MetadataMethodParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_method.PutMetadataMethod(ctx.GetUserID(), ctx.GetServiceParams("id"), p.MethodName)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteMetadataMethod struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_method
// @Summary			ลบวิธีการได้มาซึ่งข้อมูล
// @Method			DELETE
// @Parameter		metadata_method_id	query	string	example: tci3CPw5YTdhuV9PuN8oh3e-I9GySj9SZmpYRIbct4iUiSSjyOqsVPewcVLgVHgDDtmzcX35dU5bkjt4qtzD_g  รหัสวิธีการได้มาของข้อมูลแบบเข้ารหัส
// @Produces		json
// @Response		200	Struct_deleteMetadataMethod successful operation
func (srv *HttpService) deleteMetadataMethod(ctx service.RequestContext) error {

	p := &model_metadata_method.MetadataMethodParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_method.DelMetadataMethod(ctx.GetUserID(), p.MethodID)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
