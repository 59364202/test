package shared

import (
	"strings"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	model "haii.or.th/api/server/model"
)

const (
	DataServiceName = "thaiwater30/shared/file"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {
	return srv.returnFile(ctx)
}

// @DocumentName 	v1.public
// @Service			thaiwater30/shared/file
// @Parameter		file	query	string example:`maWEYa7CRiNzyi-8uUL-D68Q-1RG0oy4T-UpDMaAwzfhwZ7SrIsXYluVMAD9XaadUcLnm9xodn2DhcYGA1YbzQ` encryptext
// @Method			GET
// @Summary			Get file from api service
// @Description		Return file
// @Produces		Binary
// @Response		200	file successful operation
func (srv *HttpService) returnFile(ctx service.RequestContext) error {
	var params struct {
		File string `json:"file"`
	}
	ctx.GetRequestParams(&params)

	inf, err := model.GetCipher().DecryptText(params.File) // DecryptText
	if err != nil {
		//		return err
		return rest.NewError(422, "invalid file "+params.File, err)
	}

	var a []string
	a = append(a, setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"))
	a = append(a, strings.Split(inf, ",")...) // เอาทุกอย่างใส่่ array

	fname := filepathx.JoinPath(a...) // เอา array มา join ด้วย '/'
	ctx.ReplyFile(fname)
	return nil
}
