package analyst

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media_animation "haii.or.th/api/thaiwater30/model/media_animation"
)

type Struct_getPreRainImg struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Media `json:"data"`   // ภาพ
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/rain_forecast_img
// @Summary			ภาพคาดการณ์ฝนล่วงหน้า
// @Method			GET
// @Produces		json
// @Response		200		[]model_latest_media.Struct_Media successful operation
func (srv *HttpService) getPreRainImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetPreRain()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

type Struct_getPreRainAnimation struct {
	Result string                                `json:"result"` // example:`OK`
	Data   []*model_media_animation.Struct_Media `json:"data"`   // animation
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/rain_forecast_animation_img
// @Summary			animation คาดการณ์ฝนล่วงหน้า
// @Method			GET
// @Produces		json
// @Response		200	Struct_getPreRainAnimation successful operation
func (srv *HttpService) getPreRainAnimation(ctx service.RequestContext) error {

	rs, err := model_media_animation.GetPreRainAnimation()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
