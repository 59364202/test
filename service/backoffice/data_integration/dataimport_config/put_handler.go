package dataimport_config

import (
	"haii.or.th/api/server/model/cronjob"
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_download
// @Method			PUT
// @Summary			update data to table dataimport_download
// @Consumes		json
// @Parameter		-	body model_dataimport_config.DataDownloadConfigSwagger
// @Description		Return id of table api.dataimport_download after insert
// @Produces		json
// @Response		200		ResultDataimportDownloadPut   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDownloadPut struct {
	Result string                           `json:"result"` // example:`OK`
	Data   model_dataimport_config.NewRdlID `json:"data"`
}

// service update dataimport download config
func (srv *HttpService) updateDataimportDownloadConfig(ctx service.RequestContext) error {

	p := &model_dataimport_config.DataDownloadConfig{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	if p.DownloadSetting.SourceOptions[0] != nil {
		if p.DownloadSetting.SourceOptions[0].Details[0] != nil {
			if p.DownloadSetting.SourceOptions[0].Details[0].TimeoutSeconds <= 0 {
				p.DownloadSetting.SourceOptions[0].Details[0].TimeoutSeconds = 99 // default timeout 99 sec
			}
		}
	}

	rs, err := model_dataimport_config.UpdateDataImportConfig(p, ctx.GetUserID())
	if err != nil {
		return err
	}
	nodes, _ := model_dataimport_config.GetAllRDLNodeFromSetting()
	for _, v := range nodes {
		// ลบ cron ออกจากทุกเครื่อง กันไว้เผื่อกรณีเปลี่ยนเครื่อง
		rdl.RunRDLMgr(ctx, rdl.CmdCronDelete, datatype.MakeString(p.DownloadConfigID), datatype.MakeString(v.Text))
		// แก้ไข config ลงทุกเครื่อง กันไว้เผื่อกรณีเปลี่ยนเครื่อง
		rdl.RunRDLMgr(ctx, rdl.CmdDwonlaodEdit, datatype.MakeString(p.DownloadConfigID)+" "+rs.Agency, datatype.MakeString(v.Text))
	}

	if p.IsCronEnabled {
		// ถ้า cron เปิด ให้ใส่ cron ลงในเครื่องที่เลือก
		node, _ := model_dataimport_config.GetRDLNodeFromSetting(datatype.MakeInt(p.Node))
		rdl.RunRDLMgr(ctx, rdl.CmdCronEdit, datatype.MakeString(p.DownloadConfigID), node.Text)
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_dataset
// @Method			PUT
// @Summary			update dataimport_dataset in the server
// @Consumes		json
// @Parameter		-	body model_dataimport_config.DataImportDataSetConfigSwagger
// @Description		Return id of table api.dataimport_dataset after insert
// @Produces		json
// @Response		200		ResultDataimportDatasetPut   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDatasetPut struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`144` รหัส dataset
}

// service update dataimport dataset config
func (srv *HttpService) updateDataimportDatasetConfig(ctx service.RequestContext) error {

	p := &model_dataimport_config.DataImportDataSetConfig{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.UpdateDataImportConfigDataset(p, ctx.GetUserID())
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/metadata
// @Method			PUT
// @Summary			Update dataimport download id and dataimport dataset id to metadata by id
// @Parameter		-	form MetadataDescriptionSwagger
// @Description		Update metadata by id
// @Produces		json
// @Response		200		ResultNull   successful operation
// @Response		404			-		the request service name was not found

type MetadataDescriptionSwagger struct {
	MetadataID        int64    `json:"metadata_id"`        // รหัส metadata เช่น 1
	DownloadID        int64    `json:"download_id"`        // รหัส download เช่น 20
	DatasetID         int64    `json:"dataset_id"`         // รหัส dataset เช่น 24
	AdditionalDataset []string `json:"additional_dataset"` // รหัส dataset ที่เกี่ยวข้อง (ถ้ามี)
}

// service dataimport metadataa
func (srv *HttpService) updateMetadata(ctx service.RequestContext) error {

	p := &model_dataimport_config.MetadataDescription{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_dataimport_config.UpdateMetadata(p, "", ctx.GetUserID())
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/iscronenabled
// @Method			PUT
// @Summary			update field is cronenabled
// @Parameter		id	form    string  รหัส dataimport download เช่น 20
// @Parameter		is_cronenabled	form    bool  isCronenabled status
// @Description		Return id of table api.dataimport_download after insert
// @Produces		json
// @Response		200		ResultNull   successful operation
// @Response		404			-		the request service name was not found

// service update field cron enabled
func (srv *HttpService) cronEnabled(ctx service.RequestContext) error {
	var dataimport_download_id string
	dataimport_download_id = ctx.GetServiceParams("id")

	p := &model_dataimport_config.CronEnabled{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_dataimport_config.UpdateIsCronEnabled(dataimport_download_id, ctx.GetUserID(), p.IsCronenabled)
	if err != nil {
		return err
	}

	// nodes, _ := model_dataimport_config.GetAllRDLNodeFromSetting()
	// for _, v := range nodes {
	// 	runRDLMgr(ctx, CmdCronDelete, dataimport_download_id, datatype.MakeString(v.Text))
	// }
	// // Update cron on converter machine
	// if p.IsCronenabled {
	// 	model_dataimport_config.GetRDLNodeFromSetting()
	// 	err = runRDLMgr(ctx, CmdCronEdit, dataimport_download_id)
	// } else {

	// }

	if err != nil {
		return err
	}

	ctx.ReplyJSON(result.Result1(nil))
	return nil
}

func (srv *HttpService) updateDownloadCronList(ctx service.RequestContext) error {

	var dataimport_download_id string
	dataimport_download_id = ctx.GetServiceParams("id")

	p := &model_dataimport_config.DownloadCronSetting{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_dataimport_config.UpdateDownloadCronList(dataimport_download_id, ctx.GetUserID(), p.CrontabSetting)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

func (srv *HttpService) updateServerCronList(ctx service.RequestContext) error {

	var name string
	name = ctx.GetServiceParams("id")

	p := &model_dataimport_config.DownloadCronSetting{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_dataimport_config.UpdateApiCronList(name, p.CrontabSetting)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

type RunCronResult struct {
	LoGID int64 `json:"log_id"`
}

func (srv *HttpService) runCron(ctx service.RequestContext) error {
	name := ctx.GetServiceParams("id")
	if name == "" {
		return rest.NewError(422, "invalid cronjob setting name", nil)
	}

	job := cronjob.FindJobFromSettingName(name)
	if job == nil {
		return rest.NewError(422, "cronjob of the given setting name was not found", nil)
	}

	id, err := job.RunError(true)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(result.Result1(&RunCronResult{id}))
	return nil
}

type PutConfigVariableParams struct {
	VariableID   int64  `json:"variable_id"`
	Category     int64  `json:"category"`
	Name         string `json:"name"`
	VariableName string `json:"variable_name"`
	Value        string `json:"value"`
	//	DataGroupCode   string `json:"group_code"`
}

func (srv *HttpService) updateConfigVariable(ctx service.RequestContext) error {
	//Map parameters
	p := &PutConfigVariableParams{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	// p.UserId = ctx.GetUserID()
	//Put Agency
	//	rs, err := model.PutAgency(ctx.GetUserID(), ctx.GetServiceParams("id"), p.AgencyName, p.AgencyShortName, p.DepartmentId)
	rs, err := model_dataimport_config.UpdateConfigVariable(p.Category, p.Name, p.VariableName, p.Value, ctx.GetUserID(), p.VariableID)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
