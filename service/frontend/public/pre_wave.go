// Edit by : Thitiporn Meeprasert <thitiporn@haii.or.th>
package public

import (
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"

	model_latest_media "haii.or.th/api/thaiwater30/model/latest_media"
	model_media_animation "haii.or.th/api/thaiwater30/model/media_animation"
)

// @DocumentName	v1.public
// @Service			thaiwater30/public/wave_forecast_img
// @Summary			ภาพคาดการณ์คลื่น
// @Method			GET
// @Produces		json
// @Response		200		getPreWaveImgSwagger successful operation

type getPreWaveImgSwagger struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Media `json:"data"`
}

func (srv *HttpService) getPreWaveImg(ctx service.RequestContext) error {

	rs, err := model_latest_media.GetPreWave()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}

// @DocumentName	v1.public
// @Service			thaiwater30/public/wave_forecast_animation_img
// @Summary			animation คาดการณ์คลื่น
// @Method			GET
// @Produces		json
// @Response		200		getPreWaveAnimationSwagger successful operation
type getPreWaveAnimationSwagger struct {
	Result string                             `json:"result"` // example:`OK`
	Data   []*model_latest_media.Struct_Media `json:"data"`
}

func (srv *HttpService) getPreWaveAnimation(ctx service.RequestContext) error {

	rs, err := model_media_animation.GetPreWaveAnimation()
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
