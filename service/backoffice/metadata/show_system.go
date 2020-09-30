// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata

//ระบบแสดงข้อมูล เช่น คลังข้อมูลน้ำฯ,เว็บจังหวัด

import (
	model_metadata_show_system "haii.or.th/api/thaiwater30/model/metadata_show_system"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

type Struct_getMetadataShowSystem struct {
	Result string                                                  `json:"result"` // example:`OK`
	Data   []*model_metadata_show_system.Struct_MetadataShowSystem `json:"data"`   // ข้อมูลไปแสดงที่ระบบใดบ้าง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/show_system
// @Summary			ระบบแสดงข้อมูล
// @Method			GET
// @Produces		json
// @Response		200	Struct_getMetadataShowSystem successful operation
func (srv *HttpService) getMetadataShowSystem(ctx service.RequestContext) error {
	rs, err := model_metadata_show_system.GetMetadataShowSystem()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
