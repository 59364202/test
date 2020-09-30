package manipulate

import (
	model "haii.or.th/api/thaiwater30/model/manipulate/metadata"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// SubCategory //
//type subCategoryParam struct {
//	Text       map[string]string `json:"text"`
//	CategoryId int               `json:"category_id"`
//}

func (srv *HttpService) GetMetadataHydro(ctx service.RequestContext) error {

	return nil
}

func (srv *HttpService) getMetadataTable(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	result, err := model.GetMatadataTable(p.Subcategory, p.Agency)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) getMetadata(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	if ctx.GetServiceParams("id") == "" {
		return srv.getMetadataTable(ctx)
	}

	result, err := model.GetMetadata(ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) postMetadata(ctx service.RequestContext) error {
	p, err := ctx.GetPostParamsMap()
	if err != nil {
		return errors.Repack(err)
	}

	result, err := model.PostMetadata(ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}
	return nil
}

func (srv *HttpService) putMetadata(ctx service.RequestContext) error {
	p, err := ctx.GetPostParamsMap()
	if err != nil {
		return errors.Repack(err)
	}

	result, err := model.PutMetadata(ctx.GetServiceParams("id"), ctx.GetUserID(), p)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}
	return nil
}

func (srv *HttpService) deleteMetadata(ctx service.RequestContext) error {
	p := &Param{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	result, err := model.DeleteMetadata(ctx.GetServiceParams("id"), ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}
	return nil
}
