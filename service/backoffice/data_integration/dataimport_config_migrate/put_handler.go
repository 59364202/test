package dataimport_config

import (
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	model_metadata_provision "haii.or.th/api/thaiwater30/model/metadata_provision"
	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

func (srv *HttpService) updateDataimportDownloadConfig(ctx service.RequestContext) error {

	p := &model_dataimport_config.DataDownloadConfig{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

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

func (srv *HttpService) updateMetadata(ctx service.RequestContext) error {

	p := &model_dataimport_config.MetadataDescription{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	err := model_dataimport_config.UpdateMetadata(p, "migrate", ctx.GetUserID())
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result.Result1(nil))

	return nil
}

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
	ctx.ReplyJSON(result.Result1(nil))
	return nil
}

func (srv *HttpService) putMetadataProvision(ctx service.RequestContext) error {
	p := &model_metadata_provision.MetadataProvisionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_provision.PutMetadataPrivision(ctx.GetUserID(), p.ID, p.Name)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
