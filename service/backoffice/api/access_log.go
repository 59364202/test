package api

import (
	model_accessLog "haii.or.th/api/thaiwater30/model/accesslog"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

type AccessLogParam struct {
	DateStart string `json:"dateStart"` //เวลาเริ่มต้น เช่น 2006-01-02 15:04
	DateEnd   string `json:"dateEnd"`   // เวลาเริ่มต้น เช่น 2006-01-02 15:04
	Service   string `json:"service"`   // ชื่อ service id เช่น  105
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/access_log
// @Method			GET
// @Summary			Show data of module APIS in the server
// @Description		Return data of table api.access_log condition between datetime and service name
// @Parameter		-	query AccessLogParam
// @Produces		json
// @Response		200		AccessLogResultSwagger successful operation
// @Response		404			-		the request service name was not found

type AccessLogResultSwagger struct {
	Result string `json:"result"`  //example:`ok`
	Data []model_accessLog.ResultAccessLog 
}

func (srv *HttpService) accessLog(ctx service.RequestContext) error {

	p := &AccessLogParam{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model_accessLog.GetHistory(p.DateStart, p.DateEnd, p.Service)
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

var _dummy *result.Result

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/service_name
// @Method			GET
// @Summary			get service name from api server for dropdown box
// @Description		Return data of table api.access_log condition between datetime and service name
// @Produces		json
// @Response		200		ResultServiceNameSwagger successful operation
// @Response		404			-		the request service name was not found

type ResultServiceNameSwagger struct {
	Result string `json:"result"` // example:`OK`
	Data *ServiceNameSwagger `json:"data"`
}

type ServiceNameSwagger struct {
	DateRange     string                 `json:"date_range"` // example:`30` กำหนดช่วงวันในการเลือก
	ServiceOption []*ServiceNameSelectOptionSwagger `json:"service_name"`
}

type ServiceNameSelectOptionSwagger struct {
	Value int64  `json:"value"` // example:`1` รหัส service
	Text  string `json:"text"`  // example:`backoffice/dba/` ชื่อ service
}

func (srv *HttpService) serviceName(ctx service.RequestContext) error {

	result, err := model_accessLog.GetServiceName()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

func (srv *HttpService) agentName(ctx service.RequestContext) error {

	result, err := model_accessLog.GetAgentName()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}
