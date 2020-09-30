package data_service

import (
	"haii.or.th/api/server/model/setting"

	util_file "haii.or.th/api/thaiwater30/util/file"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/order_detail"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/reupload
// @Summary			อัพโหลดให้บริการข้อมูลอีกครั้ง
// @Method			GET
// @Parameter		od_id	query	string	example:`721` รหัส order_detail ที่ต้องการให้อัพโหลดอีกครั้ง
// @Produces		json
// @Response		200	- successful operation
type Param_reUpload struct {
	Od string `json:"od_id"`
}

func (srv *HttpService) reUpload(ctx service.RequestContext) error {
	p := &Param_reUpload{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	folderPath := util_file.UploadPath + "/" + util_file.DataserviceUploadLetterPath + setting.GetSystemSetting("dataservice.TempZipFile") + "/" + p.Od
	err := order_detail.UploadFileToWeb(folderPath, p.Od)
	return err
}
