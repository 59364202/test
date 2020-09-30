package manipulate

import (
	"encoding/json"
	model "haii.or.th/api/thaiwater30/model/manipulate/hydroinfo"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// SubCategory //
type hydroinfoParam struct {
	Id            string          `json:"id"`
	HydroInfoName json.RawMessage `json:"hydroinfo_name"`
	AgencyId      string          `json:"agency_id"`
}

func (srv *HttpService) getHydroinfo(ctx service.RequestContext) error {
	p := &hydroinfoParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	result, err := model.GetHydroinfo(ctx.GetServiceParams("id"), p.AgencyId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) postHydroInfo(ctx service.RequestContext) error {
	p := &hydroinfoParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model.PostHydroInfo(ctx.GetUserID(), p.HydroInfoName, p.AgencyId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) putHydroinfo(ctx service.RequestContext) error {
	p := &hydroinfoParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model.PutHydroinfo(ctx.GetUserID(), ctx.GetServiceParams("id"), p.HydroInfoName, p.AgencyId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}

func (srv *HttpService) deleteHydroinfo(ctx service.RequestContext) error {
	result, err := model.DeleteHydroinfo(ctx.GetUserID(), ctx.GetServiceParams("id"))
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result)
	}

	return nil
}
