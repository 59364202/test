package mobile

import (
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/media"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/media/radar/tmd/latest
// @Summary			ข้อมูลภาพเรดาร์จากกรมอุตุนิยมวิทยา
// @Description		ข้อมูลภาพเรดาร์จากกรมอุตุนิยมวิทยา
// @Method			GET
// @Produces		json
// @Response		200	media.Struct_MediaLatest successful operation
func (srv *HttpService) handlerGetMediaLatest(ctx service.RequestContext) error {
	media_url := ctx.BuildURL(0, "thaiwater30/shared/image", true)
	rs, err := media.GetMediaLatest(media_url)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
