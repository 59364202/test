package rdl

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

type ResultDataimportDownloadPost struct {
	Result string                           `json:"result"` // example:`OK`
	Data   model_dataimport_config.NewRdlID `json:"data"`
}

//postPs Run a download process on remote dataimport node.
// @Service			thaiwater30/backoffice/dataimport_config/ps
// @Method			POST
// @Summary			Run a download process on remote dataimport node.
// @Parameter		id	query    string  example:1 dataimport download id
// @Produces		json
// @Response		200		result.Result   successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) postPs(ctx service.RequestContext) error {

	dataimport_download_id := ctx.GetServiceParams("id")
	if dataimport_download_id == "" {
		return rest.NewError(422, "invalid download id", nil)
	}
	RunRDLMgr(ctx, CmdPsRun, dataimport_download_id, "")

	ctx.ReplyJSON(result.Result1(true))

	return nil
}
