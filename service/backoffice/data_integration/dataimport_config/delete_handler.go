package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @Service			thaiwater30/backoffice/api/dataimport_download
// @Method			DELETE
// @Summary			Remove data of module APIS in the server
// @Parameter		id	query    string  example:1 dataimport download id
// @Description		Soft deleted data of table api.agent
// @Produces		json
// @Response		200		ResultNull   successful operation
// @Response		404			-		the request service name was not found

// service soft deleted dataimport download config
type ResultNull struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`null`
}

func (srv *HttpService) deleteDataimportDownloadConfig(ctx service.RequestContext) error {

	var dataimport_download_id string
	dataimport_download_id = ctx.GetServiceParams("id")

	err := model_dataimport_config.DeleteDataimportDownload(dataimport_download_id, ctx.GetUserID())
	if err != nil {
		return err
	}

	rdlNodes, _ := model_dataimport_config.GetAllRDLNodeFromSetting()
	for _, v := range rdlNodes {
		// ลบ cron ออกจากทุกเครื่อง
		rdl.RunRDLMgr(ctx, rdl.CmdCronDelete, dataimport_download_id, datatype.MakeString(v.Text))
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

// @Service			thaiwater30/backoffice/api/dataimport_dataset
// @Method			DELETE
// @Summary			Remove data of module APIS in the server
// @Parameter		id	query    string  example:2 dataimport dataset id
// @Description		Soft deleted data of table api.agent
// @Produces		json
// @Response		200		ResultNull   successful operation
// @Response		404			-		the request service name was not found

// service soft deleted dataimport dataset config
func (srv *HttpService) deleteDataimportDatasetConfig(ctx service.RequestContext) error {

	var dataimport_dataset_id string
	dataimport_dataset_id = ctx.GetServiceParams("id")

	err := model_dataimport_config.DeleteDataimportDataset(dataimport_dataset_id, ctx.GetUserID())
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

func (srv *HttpService) deleteAgency(ctx service.RequestContext) error {
	//Map parameters
	p := &PutConfigVariableParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	// if ctx.GetServiceParams("variable_id") != "" {
	// 	p.VariableID = ctx.GetServiceParams("variable_id")
	// }
	// p.UserId = ctx.GetUserID()
	//Delete Agency
	//	rs, err := model.DeleteAgency(ctx.GetUserID(), ctx.GetServiceParams("id"))
	rs, err := model_dataimport_config.DeleteConfigVariable(p.VariableID, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
