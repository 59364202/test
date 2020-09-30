package image

import (
	"encoding/json"
	"strings"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	//	model "haii.or.th/api/server/model"
	"haii.or.th/api/thaiwater30/util/b64"
)

const (
	DataServiceName = "thaiwater30/shared/image"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

type SharedParam struct {
	DataType string          `json:"data_type"`
	Name     string          `json:"name"`
	Data     json.RawMessage `json:"data"`
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {
	return srv.displayImage(ctx)
}

// @DocumentName 	v1.public
// @Service			thaiwater30/shared/image
// @Parameter		image	query	string	example: AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJHsgkwaxkBkU7K-JlYlJ0eL-kOj6BOobumHMJJ4nYHNIfCEaLG3zGgZ2LEuTaHsMjGebDH8UN encryptext
// @Method			GET
// @Summary			Get image from api service
// @Description		Return file
// @Produces		Binary
// @Response		200	file  successful operation
func (srv *HttpService) displayImage(ctx service.RequestContext) error {
	var params struct {
		ImageID string `json:"image"`
	}
	ctx.GetRequestParams(&params)

	//	inf, err := model.GetCipher().DecryptText(params.ImageID)
	//	if err != nil {
	//		return rest.NewError(422, "invalid image", err)
	//	}
	inf, err := b64.DecryptText(params.ImageID)
	if err != nil {
		return rest.NewError(422, "invalid image", err)
	}
	var a []string
	a = append(a, setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"))
	a = append(a, strings.Split(inf, ",")...)

	fname := filepathx.JoinPath(a...)
	//	ctx.ReplyJSON(fname)
	ctx.ReplyFile(fname)
	return nil
}
