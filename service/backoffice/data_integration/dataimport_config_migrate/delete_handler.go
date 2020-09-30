package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	model_metadata_provision "haii.or.th/api/thaiwater30/model/metadata_provision"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
)

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

func (srv *HttpService) deleteMetadataProvision(ctx service.RequestContext) error {

	p := &model_metadata_provision.MetadataProvisionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_provision.DelMetadataProvision(ctx.GetUserID(), p.ID)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
