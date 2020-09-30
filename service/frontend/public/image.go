// Author: Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media_animation "haii.or.th/api/thaiwater30/model/media_animation"
)

// @DocumentName	v1.public
// @Service			thaiwater30/public/water_balance_img
// @Summary			ภาพสมดุลน้ำจาก MODEL SWAT (Soil and Water Assessment Tool)
// @Method			GET
// @Produces		json
// @Response		200		getImgSwagger successful operation

type getImgSwagger struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Media `json:"data"`
}

type Rain7dayForecastStruct struct {
	PreRain          *result.Result `json:"pre_rain,omitempty"`
	PreRainSea       *result.Result `json:"pre_rain_sea,omitempty"`
	PreRainAsia      *result.Result `json:"pre_rain_asia,omitempty"`
	PreRainBasin     *result.Result `json:"pre_rain_basin,omitempty"`
	PreRainAnimation *result.Result `json:"pre_rain_animation,omitempty"`
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/rain7day_forecast
// @Summary			ภาพคาดการณ์ฝนล่วงหน้า 7 วัน
// @Method			GET
// @Produces		json
// @Response		200		Rain7dayForecastStruct successful operation

func (srv *HttpService) getRain7dayForecast(ctx service.RequestContext) error {

	rs := &Rain7dayForecastStruct{}

	// pre rain th
	rs.PreRain = buildResult(model_latest_media.GetPreRainTH())
	// pre rain asia
	rs.PreRainAsia = buildResult(model_latest_media.GetPreRainAsia())
	// pre rain south east asia
	rs.PreRainSea = buildResult(model_latest_media.GetPreRainSea())
	// pre rain basin
	rs.PreRainBasin = buildResult(model_latest_media.GetPreRainBasin())
	// pre rain animation
	rs.PreRainAnimation = buildResult(model_media_animation.GetPreRainAnimation())

	//	rs, err := model_latest_media.GetPreRain()
	//	if err != nil {
	//		ctx.ReplyJSON(result.Result0(err))
	//	} else {
	//		ctx.ReplyJSON(result.Result1(rs))
	//	}

	ctx.ReplyJSON(result.Result1(rs))

	return nil
}

func (srv *HttpService) getWaterBalanceImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetWaterBalance()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/water_balance_forcast_weekly_img
// @Summary			ภาพคาดการณ์สมดุลน้ำ รายสัปดาห์ จาก  MODEL SWAT(Soil and Water Assessment Tool)
// @Method			GET
// @Produces		json
// @Response		200		getImgSwagger successful operation

func (srv *HttpService) getPreWaterBalanceWeeklyImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetPreWaterBalanceWeekly()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/water_balance_forcast_monthly_img
// @Summary			ภาพคาดการณ์สมดุลน้ำ รายเดือน  จาก  MODEL SWAT(Soil and Water Assessment Tool)
// @Method			GET
// @Produces		json
// @Response		200		getImgSwagger successful operation

func (srv *HttpService) getPreWaterBalanceMonthlyImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetPreWaterBalanceMonthly()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
