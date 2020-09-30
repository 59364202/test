package rdl

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type ResultNull struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`null`
}

//deletePs terminate a download process on remote dataimport node.
// @Service			thaiwater30/backoffice/dataimport_config/ps
// @Method			DELETE
// @Summary			terminate a download process on remote dataimport node.
// @Parameter		id	query    string  example:1 dataimport download id
// @Produces		json
// @Response		200		result.Result   successful operation
// @Response		404			-		the request service name was not found

func (srv *HttpService) deletePs(ctx service.RequestContext) error {
	dataimport_download_id := ctx.GetServiceParams("id")
	err := RunRDLMgr(ctx, CmdPsKill, dataimport_download_id, "")
	if err != nil {
		return errors.Repack(err)
	}
	ctx.ReplyJSON(result.Result1(true))
	return nil
}
