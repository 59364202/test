package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
)

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_download
// @Method			POST
// @Summary			Add data to table dataimport_download  in the server
// @Consumes		json
// @Parameter		-	body model_dataimport_config.DataDownloadConfigSwagger
// @Description		Return id of table api.dataimport_download after insert
// @Produces		json
// @Response		200		ResultDataimportDownloadPost   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDownloadPost struct {
	Result string                           `json:"result"` // example:`OK`
	Data   model_dataimport_config.NewRdlID `json:"data"`
}

// service add dataimport download config
func (srv *HttpService) addDataimportDownloadConfig(ctx service.RequestContext) error {

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

	rs, err := model_dataimport_config.AddDataImportDownloadConfig(p, ctx.GetUserID())
	if err != nil {
		return err
	}

	nodes, _ := model_dataimport_config.GetAllRDLNodeFromSetting()
	for _, v := range nodes {
		// เพิ่ม config ลงทุกเครื่อง กันไว้เผื่อกรณีเปลี่ยนเครื่อง
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
// @Method			POST
// @Summary			Add data to table dataimport_dataset in the server
// @Consumes		json
// @Parameter		-	body model_dataimport_config.DataImportDataSetConfigSwagger
// @Description		Return id of table api.dataimport_dataset after insert
// @Produces		json
// @Response		200		ResultDataimportDatasetPost  successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDatasetPost struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`144` รหัส dataset
}

type ResultDataimportDownloadCopy struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`144` รหัส download
}

// service add dataimport dataset config
func (srv *HttpService) addDataimportDatasetConfig(ctx service.RequestContext) error {

	p := &model_dataimport_config.DataImportDataSetConfig{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.AddDataImportDatasetConfig(p, ctx.GetUserID())
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_dataset_copy
// @Method			POST
// @Summary			copy dataimport dataset config to table dataimport_dataset
// @Parameter		-	form DatasetIDSwagger
// @Description		copy dataimport dataset config
// @Produces		json
// @Response		200		ResultDataimportDatasetPost   successful operation
// @Response		404			-		the request service name was not found

type DatasetIDSwagger struct {
	DatasetID int64 `json:"dataset_id"` // รหัส  dataset เช่น 1
}

// copy dataimport dataset config
func (srv *HttpService) copyDataimportDatasetConfig(ctx service.RequestContext) error {

	type DatasetID struct {
		DatasetID int64 `json:"dataset_id"`
	}
	p := &DatasetID{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.CopyDataimportDatasetConfig(p.DatasetID)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_download_copy
// @Method			POST
// @Summary			copy dataimport download config to table dataimport_download
// @Parameter		-	form DownloadIDSwagger
// @Description		copy dataimport download config
// @Produces		json
// @Response		200		ResultDataimportDownloadCopy   successful operation
// @Response		404			-		the request service name was not found

type DownloadIDSwagger struct {
	DownloadID int64 `json:"download_id"` // รหัส download เช่น 1
}

// copy dataimport download config
func (srv *HttpService) copyDataimportDownloadConfig(ctx service.RequestContext) error {

	type DownloadID struct {
		DownloadID int64 `json:"download_id"`
	}
	p := &DownloadID{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.CopyDataimportDownloadConfig(p.DownloadID)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}
