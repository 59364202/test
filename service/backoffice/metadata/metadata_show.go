// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata

// เก็บข้อมูลว่าระบบที่นำข้อมูลไปแสดงระบบไหนบ้าง และแสดงในส่วนไหน

import (
	"strconv"

	model_metadata_show "haii.or.th/api/thaiwater30/model/metadata_show"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Struct_getMetadataShow struct {
	Result string                                           `json:"result"` // example:`OK`
	Data   []*model_metadata_show.Struct_MetadataShow_Param `json:"data"`   // ข้อมูลไปแสดงที่ระบบใดบ้าง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_show
// @Summary			ข้อมูลไปแสดงที่ระบบใดบ้าง
// @Method			GET
// @Produces		json
// @Response		200	Struct_getMetadataShow successful operation
func (srv *HttpService) getMetadataShow(ctx service.RequestContext) error {
	rs, err := model_metadata_show.GetMetadataShow()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_postMetadataShow struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Insert Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_show
// @Summary			เพิ่มข้อมูลไปแสดงที่ระบบใดบ้าง
// @Method			POST
// @Consumes		json
// @Parameter		-	body	model_metadata_show.Struct_MetadataShow_Param
// @Produces		json
// @Response		200	Struct_postMetadataShow successful operation
func (srv *HttpService) postMetadataShow(ctx service.RequestContext) error {
	p := &model_metadata_show.Struct_MetadataShow_Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_show.PostMetadataShow(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_putMetadataShow struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_show/{id}
// @Summary			แก้ไขข้อมูลไปแสดงที่ระบบใดบ้าง
// @Method			PUT
// @Parameter		id	form	string	รหัส example 1
// @Produces		json
// @Response		200	Struct_putMetadataShow successful operation
func (srv *HttpService) putMetadataShow(ctx service.RequestContext) error {
	p := &model_metadata_show.Struct_MetadataShow_Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_show.PutMetadataShow(p, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_deleteMetadataShow struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Delete Successful`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/metadata_show
// @Summary			ลบข้อมูลไปแสดงที่ระบบใดบ้าง
// @Method			DELETE
// @Parameter		id	query	string	รหัส example 1
// @Produces		json
// @Response		200	Struct_deleteMetadataShow successful operation
func (srv *HttpService) deleteMetadataShow(ctx service.RequestContext) error {

	p := &model_metadata_show.Struct_MetadataShow_Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_show.Delete(ctx.GetUserID(), strconv.Itoa(int(p.ID)))
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
