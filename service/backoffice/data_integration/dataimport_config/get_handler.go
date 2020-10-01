package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	//	"haii.or.th/api/server/model/dataimport"
)

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_download/{id}
// @Method			GET
// @Summary			Show data of module dataimport in the server
// @Parameter		id	path    string  รหัส dataimport download เช่น 20
// @Description		Return data api.dataimport_download by id
// @Produces		json
// @Response		200		ResultDataimportDownloadID   successful operation
// @Response		404			-		the request service name was not found

type ResultDataimportDownloadID struct {
	Result string                                            `json:"result"` // example:`OK`
	Data   *model_dataimport_config.DataImportDownloadConfig `json:"data"`
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_download
// @Method			GET
// @Summary			Show data of module dataimport in the server
// @Description		Return data of table api.dataimport_download
// @Produces		json
// @Response		200		ResultDataimportDownloadList   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDownloadList struct {
	Result string                                                  `json:"result"` // example:`OK`
	Data   *model_dataimport_config.DataDetailsDownloadListSwagger `json:"data"`
}

// service get list dataimport download config or one dataimport download config depend on params
func (srv *HttpService) getDataimportConfig(ctx service.RequestContext) error {
	var dataimport_download_id string

	//	dataimport.GenerateDataimportDailyReportEvent()

	dataimport_download_id = ctx.GetServiceParams("id")
	if dataimport_download_id != "" {
		rs, err := model_dataimport_config.GetDataImportDownloadConfig(dataimport_download_id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		rs, err := model_dataimport_config.GetDataImportDownloadList("")
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_dataset/{id}
// @Method			GET
// @Summary			Show data of module dataimport dataset
// @Parameter		id	path    string  รหัส dataimport dataset เช่น 32
// @Description		Return data of table api.dataimport_download
// @Produces		json
// @Response		200		ResultDataimportDatasetID   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDatasetID struct {
	Result string                                            `json:"result"` // example:`OK`
	Data   *model_dataimport_config.DataImportDataSetConfig1 `json:"data"`
}

// @Service			thaiwater30/backoffice/dataimport_config/dataimport_dataset
// @Method			GET
// @Summary			Show list of module dataimport dataset
// @Description		Return data of table api.dataimport_dataset
// @Produces		json
// @Response		200		ResultDataimportDatasetList   successful operation
// @Response		404			-		the request service name was not found
type ResultDataimportDatasetList struct {
	Result string                                                 `json:"result"` // example:`OK`
	Data   *model_dataimport_config.DataDetailsDatasetListSwagger `json:"data"`
}

// service get list dataimport dataset config or one dataimport dataset config depend on params
func (srv *HttpService) getDataimportDatasetConfig(ctx service.RequestContext) error {
	var dataimport_download_id, download_type string
	dataimport_download_id = ctx.GetServiceParams("id")
	if dataimport_download_id != "" {
		rs, err := model_dataimport_config.GetDataImportDatasetConfig(dataimport_download_id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		rs, err := model_dataimport_config.GetDataImportDatasetList("",download_type)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/metadata
// @Method			GET
// @Summary			Show data of module dataimport in the server
// @Description		Return data of table api.dataimport_download
// @Produces		json
// @Response		200		ResultDataDetailsSwagger   successful operation
// @Response		404			-		the request service name was not found
type ResultDataDetailsSwagger struct {
	Result string                                      `json:"result"` // example:`OK`
	Data   *model_dataimport_config.DataDetailsSwagger `json:"data"`
}

//  service list dataimport metadata list
func (srv *HttpService) getDataimportConfigList(ctx service.RequestContext) error {
	rs, err := model_dataimport_config.GetDataImportConfigMetadataList("")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/history_page
// @Method			GET
// @Summary			Show data of module dataimport in the server
// @Description		Return data of table api.dataimport_download
// @Produces		json
// @Response		200		ResultHistorySelectOptListSwagger   successful operation
// @Response		404			-		the request service name was not found
type ResultHistorySelectOptListSwagger struct {
	Result string                                               `json:"result"` // example:`OK`
	Data   *model_dataimport_config.HistorySelectOptListSwagger `json:"data"`
}

// service dropdown list dataimport history
func (srv *HttpService) getDataimportHistoryDetails(ctx service.RequestContext) error {

	rs, err := model_dataimport_config.GetDataimportHistorySelectOptList("")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

// @Service			thaiwater30/backoffice/dataimport_config/history
// @Method			GET
// @Summary			Show data of module dataimport in the server
// @Parameter		-	query model_dataimport_config.HistoryDataSelect
// @Description		Return data of table api.dataimport_download
// @Produces		json
// @Response		200		DataimportHistorySwagger   successful operation
// @Response		404			-		the request service name was not found

type DataimportHistorySwagger struct {
	Result string                                    `json:"result"` //example:`OK`
	Data   []model_dataimport_config.HistoryMetadata `json:"data"`
}

// service list dataimport history list
func (srv *HttpService) getDataimportHistoryData(ctx service.RequestContext) error {

	p := &model_dataimport_config.HistoryDataSelect{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.GetDataimportHistoryList(p.AgencyID, p.MetadataID, p.ProcessStatus, p.BeginAt, p.EndAt, "")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getDownloadCronList(ctx service.RequestContext) error {

	rs, err := model_dataimport_config.GetDownloadCronList()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getServerCronList(ctx service.RequestContext) error {

	rs := model_dataimport_config.GetServerCronList()

	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getConfigVariable(ctx service.RequestContext) error {

	rs, err := model_dataimport_config.GetConfigVariable()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

type ImportListResult struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

func (srv *HttpService) getListVariable(ctx service.RequestContext) error {

	rs, err := model_dataimport_config.GetListCatVariable()
	result := &ImportListResult{}
	result.Data = rs
	if err != nil {
		result.Result = "NO"
		result.Data = errors.Repack(err)
		ctx.ReplyJSON(result)
	} else {
		result.Result = "OK"
		ctx.ReplyJSON(result)
	}

	return nil
}

