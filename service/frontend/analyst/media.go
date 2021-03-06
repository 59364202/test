package analyst

import (
	model_media "haii.or.th/api/thaiwater30/model/media"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	model_media_animation "haii.or.th/api/thaiwater30/model/media_animation"
)

type UpperWindHistoryInput struct {
	Height	string `json:"height"` // เช่น 06, 15 (0.6km or 1.5km)
	Year    string `json:"year"` // ปี เช่น 2016
	Month   string `json:"month"` // เดือนเช่น 04
}

type SwatHistoryInput struct {
	Year     string `json:"year"`	// ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
	DataType string `json:"datatype"` // ประเภทรูปภาพ เช่น animation thailand,animation southeast asia,animation asia,asia,southeast asia,thailand,thailand basin
}

type Wind10mHistoryInput struct {
	Year     string `json:"year"`	// ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
	TinitTime    string `json:"tinit_time"` // รอบของภาพ เช่น  07 หรือ 19
}
type RainAccumulatInput struct {
	Year     string `json:"year"`	// ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
}
type WaveHistoryInput struct {
	Year     string `json:"year"` // ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
	Day      string `json:"day"` // วันเช่น 08
	DataType string `json:"datatype"` // ประเภทรูปภาพเช่น animation,image
}

type VerticalWindHistoryInput struct {
	Year     string `json:"year"` // ปี เช่น 2016
	Month    string `json:"month"` // เดือนเช่น 04
}

type ReportHistoryInput struct {
	Limit    	string `json:"limit"`		// limit จำนวนข้อมูล เช่น 10
	AgencyId    string `json:"agency_id"` 	// รหัสหน่วยงาน เช่น 9
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/upper_wind_img
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/upper_wind_history_img
// @Method			GET
// @Summary			ภาพลมและความกดอากาศที่ความสูง 0.6 ย้อนหลัง
// @Parameter		-	query UpperWindHistoryInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getUpperWindHistory(ctx service.RequestContext) error {
	p := &UpperWindHistoryInput{}
	
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetUpperWindHistory(p.Height, p.Year, p.Month)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/rain_forecast_history_img
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/wind10m_forecast_history_img
// @Method			GET
// @Summary			ประวัติภาพคาดการณ์ลม 10 เมตร
// @Parameter		- query Wind10mHistoryInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getWind10mHistory(ctx service.RequestContext) error {

	p := &Wind10mHistoryInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetWind10mHistory(p.Year, p.Month, p.TinitTime)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/wind10m_forecast_animation_img
// @Summary			animation คาดการณ์ฝนล่วงหน้า
// @Method			GET
// @Produces		json
// @Response		200	MediaFileOutputSwagger successful operation
func (srv *HttpService) getPreWindAnimation(ctx service.RequestContext) error {

	rs, err := model_media_animation.GetPreWindAnimation()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/rainaccumulat_img
// @Method			GET
// @Summary			แผนภาพฝนสะสมรายวัน (USNRL)
// @Parameter		- query RainAccumulatInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) rainAccumulatHistory(ctx service.RequestContext) error {

	p := &RainAccumulatInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetRainAccumulatHistory(p.Year, p.Month)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}


// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/wave_forecast_history_img
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/vertical_wind_history_img
// @Method			GET
// @Summary			ภาพลมแนวดิ่ง 5km ย้อนหลัง
// @Parameter		-	query VerticalWindHistoryInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getVerticalWindHistory(ctx service.RequestContext) error {
	p := &VerticalWindHistoryInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetVerticalWindHistory(p.Year, p.Month)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/swat_img
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/swat_history_img
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/water_situation
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

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/report_history
// @Method			GET
// @Summary			แผนภาพสมดุลน้ำย้อนหลัง
// @Parameter		-	query SwatHistorySwaggerInput
// @Produces		json
// @Response		200		MediaFileOutputSwagger successful operation
func (srv *HttpService) getReportHistory(ctx service.RequestContext) error {
	p := &ReportHistoryInput{}
	
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	rs, err := model_media.GetReportHistory(p.Limit, p.AgencyId)

	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}