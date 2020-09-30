package data_service

import (
	data "haii.or.th/api/server/model/eventlog"
	model_user "haii.or.th/api/server/model/user"
	model_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_header "haii.or.th/api/thaiwater30/model/order_header"

	"time"

	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/data_service/data_service
// @Summary			ส่งคำขอบริการข้อมูล
// @Consumes		json
// @Method			POST
// @Parameter		-	body	model_header.Param_OrderHeader required:true
// @Produces		json
// @Response		200		-	successful operation
func (srv *HttpService) postShopping(ctx service.RequestContext) error {
	p := &model_header.Param_OrderHeader{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if p.User_Id == 0 {
		p.User_Id = ctx.GetUserID()
	}
	p, err := model_header.InsertOrder(p) // insert ลง db
	if err != nil {
		ctx.ReplyError(err)
	} else {
		// เตรียมข้อมูลสำหรับส่ง mail
		mData := &model_detail.MailData{}
		mData.Data, err = model_detail.GetOrderDetailByOrderHeaderId(p.Id, ctx)
		if err != nil {
			return errors.Repack(err)
		}
		mData.UserName = model_user.GetUser(p.User_Id).FullName
		mData.UserId = p.User_Id
		mData.Date = time.Now().Format("2006-01-02")
		mData.ServiceId = ctx.GetServiceID()
		mData.AgentUserId = ctx.GetAgentUserID()
		mData.IsInit = true

		media_url := ctx.BuildURL(0, "thaiwater30/shared/file", true)
		log.WrapPanicGO(func() { model_detail.OnCreateDataService(mData, media_url) }) // ส่งเมลล์

		data.LogSystemEvent(ctx.GetServiceID(), ctx.GetAgentUserID(), p.User_Id, eventcode.EventDataServiceSendMail, "POST : data_service shopping", mData)
		ctx.ReplyJSON(result.Result1(""))
	}

	return nil
}
