package migration_log

import (
	model "haii.or.th/api/thaiwater30/model/migration"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type Result struct {
	Result     string      `json:"result"`
	LastUpdate interface{} `json:"last_update"`
	BgRunning  bool        `json:"bg_running"`
	Data       interface{} `json:"data"`
}

func (srv *HttpService) getSummaryData(ctx service.RequestContext) error {

	//	rs, err = model.GetDataformat(ctx.GetServiceParams("id"), p.MetdataMethodID)
	rsl := &Result{}

	rs, ls, isBgRunning, err := model.GetSummaryData()

	if err != nil {
		rsl.Result = "NO"
		rsl.Data = err.Error()
		ctx.ReplyJSON(rsl)
	} else {
		rsl.Result = "OK"
		rsl.LastUpdate = ls
		rsl.BgRunning = isBgRunning
		rsl.Data = rs
		ctx.ReplyJSON(rsl)
	}
	return nil
}

func (srv *HttpService) getSummaryMasterData(ctx service.RequestContext) error {
	//	rs, err = model.GetDataformat(ctx.GetServiceParams("id"), p.MetdataMethodID)
	rsl := &Result{}

	rs, ls, isBgRunning, err := model.GetSummaryMasterData()

	if err != nil {
		rsl.Result = "NO"
		rsl.Data = err.Error()
		ctx.ReplyJSON(rsl)
	} else {
		rsl.Result = "OK"
		rsl.LastUpdate = ls
		rsl.BgRunning = isBgRunning
		rsl.Data = rs
		ctx.ReplyJSON(rsl)
	}
	return nil
}

func (srv *HttpService) getSummaryImage(ctx service.RequestContext) error {

	rsl := &Result{}
	rs, ls, isBgRunning, err := model.GetSummaryImage()

	if err != nil {
		rsl.Result = "NO"
		rsl.Data = err.Error()
		ctx.ReplyJSON(rsl)
	} else {
		rsl.Result = "OK"
		rsl.LastUpdate = ls
		rsl.BgRunning = isBgRunning
		rsl.Data = rs
		ctx.ReplyJSON(rsl)
	}
	return nil
}

type DataParam struct {
	NHCTable  string `json:"nhc_table"`
	TW30Table string `json:"tw30_table"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (srv *HttpService) getDataByTable(ctx service.RequestContext) error {

	p := &DataParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	rs, err := model.GetDataByTable(p.NHCTable, p.StartDate, p.EndDate)

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenTable(ctx service.RequestContext) error {

	rs, err := model.RegenData()

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenMasterTable(ctx service.RequestContext) error {

	rs, err := model.RegenMasterData()

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenTableImg(ctx service.RequestContext) error {

	rs, err := model.RegenDataImg()

	if err != nil {
		ctx.ReplyJSON(result.Result1(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenTableImg1(ctx service.RequestContext) error {

	rs, err := model.RegenDataImg1()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenTableImg2(ctx service.RequestContext) error {

	rs, err := model.RegenDataImg2()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

func (srv *HttpService) getRegenTableImg3(ctx service.RequestContext) error {

	rs, err := model.RegenDataImg3()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type ImageParam struct {
	MediaTypeID int64 `json:"media_type_id"`
}

func (srv *HttpService) getImgByMedia(ctx service.RequestContext) error {
	p := &ImageParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	rs, err := model.GetImgByMedia(p.MediaTypeID)

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
