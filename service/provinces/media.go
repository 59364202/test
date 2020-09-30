package provinces

import (
	model_media "haii.or.th/api/thaiwater30/model/media"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type SwatHistoryInput struct {
	Year     string `json:"year"`	// ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
	DataType string `json:"datatype"` // ประเภทรูปภาพ เช่น animation thailand,animation southeast asia,animation asia,asia,southeast asia,thailand,thailand basin
}

type WaveHistoryInput struct {
	Year     string `json:"year"` // ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
	Day      string `json:"day"` // วันเช่น 08
	DataType string `json:"datatype"` // ประเภทรูปภาพเช่น animation,image
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/upper_wind_img
// @Method			GET
// @Summary			ภาพลม
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getUpperWindLatest(ctx service.RequestContext) error {

	rs, err := model_media.GetUpperWindLatest()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/rain_forecast_history_img
// @Method			GET
// @Summary			ประวัติภาพคาดการณ์ฝน
// @Parameter		- query SwatHistoryInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation

type MediaFileOutputSwagger struct {
	Result string `json:"result"`  //example:`OK`
	Data []model_media.MediaFileOutput `json:"data"`
}

func (srv *HttpService) getPrecipitationRainHistory(ctx service.RequestContext) error {

	p := &SwatHistoryInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetPrecipitationRainHistory(p.Year, p.Month, p.DataType)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/wave_forecast_history_img
// @Method			GET
// @Summary			ภาพคาดการณ์คลื่นย้อนหลัง
// @Parameter		-	query WaveHistoryInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getWaveHistory(ctx service.RequestContext) error {

	p := &WaveHistoryInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetWaveHistory(p.Year, p.Month, p.Day, p.DataType)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/swat_img
// @Method			GET
// @Summary			แผนภาพสมดุลน้ำ
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getSwatLatest(ctx service.RequestContext) error {

	rs, err := model_media.GetSwatLatest()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/swat_history_img
// @Method			GET
// @Summary			แผนภาพสมดุลน้ำย้อนหลัง
// @Parameter		-	query SwatHistorySwaggerInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
type SwatHistorySwaggerInput struct {
	Year     string `json:"year"`	// ปี เช่น 2017
	Month    string `json:"month"` // เดือนเช่น 08
	DataType string `json:"datatype"` // ประเภทรูปภาพ เช่น swat-w-forecast หรือ swat-w-back หรือ swat-m
}

func (srv *HttpService) getSwatHistory(ctx service.RequestContext) error {

	p := &SwatHistorySwaggerInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	rs, err := model_media.GetSwatHistory(p.Year, p.Month, p.DataType)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.provinces
// @Service			thaiwater30/provinces/water_situation
// @Method			GET
// @Summary			สถานการณ์น้ำ
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getWaterSituation(ctx service.RequestContext) error {

	rs, err := model_media.GetWaterSituation()

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
