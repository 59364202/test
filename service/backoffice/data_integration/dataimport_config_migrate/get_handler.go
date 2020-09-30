package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	model_metadata_provision "haii.or.th/api/thaiwater30/model/metadata_provision"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

func (srv *HttpService) getDataimportConfig(ctx service.RequestContext) error {
	var dataimport_download_id string
	dataimport_download_id = ctx.GetServiceParams("id")
	if dataimport_download_id != "" {
		rs, err := model_dataimport_config.GetDataImportDownloadConfig(dataimport_download_id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		rs, err := model_dataimport_config.GetDataImportDownloadList("migrate")
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

func (srv *HttpService) getDataimportDatasetConfig(ctx service.RequestContext) error {
	var dataimport_download_id string
	type DownloadName struct {
		DownloadName string `json:"download_name"`
	}
	
	dataimport_download_id = ctx.GetServiceParams("id")
	
	p := &DownloadName{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	
	
	if dataimport_download_id != "" {
		rs, err := model_dataimport_config.GetDataImportDatasetConfig(dataimport_download_id)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	} else {
		rs, err := model_dataimport_config.GetDataImportDatasetList("migrate",p.DownloadName)
		if err != nil {
			return err
		}
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

func (srv *HttpService) getDataimportConfigList(ctx service.RequestContext) error {
	rs, err := model_dataimport_config.GetDataImportConfigMetadataList("migrate")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getDataimportHistoryDetails(ctx service.RequestContext) error {

	rs, err := model_dataimport_config.GetDataimportHistorySelectOptList("migrate")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getDataimportHistoryData(ctx service.RequestContext) error {

	p := &model_dataimport_config.HistoryDataSelect{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	rs, err := model_dataimport_config.GetDataimportHistoryList(p.AgencyID, p.MetadataID, p.ProcessStatus, p.BeginAt, p.EndAt, "migrate")
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getMetadataPrivision(ctx service.RequestContext) error {

	var metadata_provision_id string
	metadata_provision_id = ctx.GetServiceParams("id")

	rs, err := model_metadata_provision.GetMetadataProvision(metadata_provision_id)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
