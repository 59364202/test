package mobile

import (
	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/lt_geocode"
)

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/province
// @Summary			ข้อมูลจังหวัด
// @Description		ข้อมูลจังหวัด
// @Method			GET
// @Produces		json
// @Response		200	lt_geocode.Struct_ProvinceRegion successful operation
func (srv *HttpService) handlerGetProvince(ctx service.RequestContext) error {
	rs, err := lt_geocode.GetProvinceRegion()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(rs)
	return nil
}
