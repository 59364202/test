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

func (srv *HttpService) addDataimportDownloadConfig(ctx service.RequestContext) error {

	p := &model_dataimport_config.DataDownloadConfig{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

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

func (srv *HttpService) postMetadataProvision(ctx service.RequestContext) error {
	p := &model_metadata_provision.MetadataProvisionInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	rs, err := model_metadata_provision.PostMetadataProvision(ctx.GetUserID(), p.Name)
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	}
	ctx.ReplyJSON(result.Result1(rs))

	return nil
}
