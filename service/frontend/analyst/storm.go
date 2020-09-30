// Author: Manorot Tangseveephan <manorot@hii.or.th>
package analyst

// ข้อมูลขื่อพายุในอดีต
import (
	"haii.or.th/api/util/service"
	"haii.or.th/api/util/errors"
	
	model_storm "haii.or.th/api/thaiwater30/model/storm"
	result "haii.or.th/api/thaiwater30/util/result"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/analyst/storm_history
// @Summary			ข้อมูลพายุในอดีต
// @Parameter		-	query model_storm.StructStormHistoryParam
// @Method			GET
// @Produces		json
// @Response		200		StructStormHistory successful operation

type StormHistorySwagger struct {
	Result string                      				`json:"result"`
	Data   []*model_storm.StructStormHistory 		`json:"data"`    // ข้อมูลพายุในอดีต
}

// get storm history data, if no criteria will get latest 5 storms
func (srv *HttpService) getStormHistory(ctx service.RequestContext) error {
	// get params from querystring
	p := &model_storm.StructStormHistoryParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	
	// get data
	rs, err := model_storm.GetStormHistory(p)
	
	// reply result
	if err != nil {
		ctx.ReplyJSON(result.Result0(err))
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}
