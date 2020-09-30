package metadata

import (
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
	"haii.or.th/api/util/rest"
	model_spatial_thailand_boundary "haii.or.th/api/thaiwater30/model/spatial_thailand_boundary"
	model_lt_geocode "haii.or.th/api/thaiwater30/model/lt_geocode"
)

type Struct_Geocode struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*model_lt_geocode.Struct_Geocode_Id `json:"data"`   // กระทรวง
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/metadata/geocode
// @Summary			Get geocode from latitude, longitude
// @Method			GET
// @Parameter		lat	query	string	required:true example:`13.483333` Latitude
// @Parameter		lng	query	string	required:true example:`101.006389` Longitude
// @Produces		json
// @Response		200	Struct_Geocode successful operation
func (srv *HttpService) getGeocodeFromLatLon(ctx service.RequestContext) error {
	// Map parameters
	lat := ctx.GetServiceParams("lat")
	lon := ctx.GetServiceParams("lon")
	
	if lat != "" && lon != "" {
		// query geocode
		rs, err := model_spatial_thailand_boundary.GetGeocode(lat, lon)
		if err != nil {
			return rest.NewError(422, "No rows in result set", err)
		} else {
			ctx.ReplyJSON(result.Result1(rs))
		}
		return nil
	}

	return rest.NewError(422, "No latitude or longitude provided.", nil)
}