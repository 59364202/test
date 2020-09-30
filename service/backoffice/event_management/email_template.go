package event_management

import (
	model_template "haii.or.th/api/thaiwater30/model/event_log_sink_template"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_template
// @Method			GET
// @Summary			return list email template id
// @Produces		json
// @Response		200		EmailTemplateSwagger successful operation
// @Response		404			-		the request service name was not found

type EmailTemplateSwagger struct {
	Result string                        `json:"result"` //example:`OK`
	Data   []model_template.TemplateList `json:"data"`   // template
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_template/{id}
// @Method			GET
// @Parameter		id	path    string  template id ex. 1
// @Summary			return email template
// @Produces		json
// @Response		200		ResultTemplateDetails successful operation
type ResultTemplateDetails struct {
	Result string                          `json:"result"` // example:`OK`
	Data   *model_template.TemplateDetails `json:"data"`
}

func (srv *HttpService) getEmailTemplate(ctx service.RequestContext) error {
	var templateID string
	templateID = ctx.GetServiceParams("id")
	if templateID != "" {
		dataResult, err := model_template.GetTemplate(templateID)
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		dataResult, err := model_template.GetTemplates()
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	}
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_template
// @Method			POST
// @Consumes		json
// @Parameter		-	body model_template.TemplatInputSwagger
// @Summary			return email template
// @Produces		json
// @Response		200		ResultEmailTemplateID successful operation
type ResultEmailTemplateID struct {
	Result string `json:"result"` // example:`OK`
	Data   int64  `json:"data"`   // example:`34` รหัส template
}

func (srv *HttpService) postEmailTemplate(ctx service.RequestContext) error {

	p := &model_template.TemplatInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_template.AddTemplate(ctx.GetUserID(), p.Name, p.MessageSubject, p.MessageBody)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_template
// @Method			PUT
// @Consumes		json
// @Parameter		-	body model_template.TemplatInput
// @Summary			return email template
// @Produces		json
// @Response		200		ResultEmailTemplateID successful operation

func (srv *HttpService) putEmailTemplate(ctx service.RequestContext) error {

	p := &model_template.TemplatInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)

	dataResult, err := model_template.UpdateTemplate(ctx.GetUserID(), p.ID, p.Name, p.MessageSubject, p.MessageBody)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_template
// @Method			DELETE
// @Parameter		-	query TemplatInputDel
// @Summary			return email template
// @Produces		json
// @Response		200		ResultEmailTemplateID successful operation

type TemplatInputDel struct {
	ID int64 `json:"id"` // example:`1` รหัส template เช่น 8
}

func (srv *HttpService) deletedEmailTemplate(ctx service.RequestContext) error {
	p := &model_template.TemplatInput{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	if p.ID != 0 {
		dataResult, err := model_template.DeleteTemplate(p.ID, ctx.GetUserID())
		if err != nil {
			ctx.ReplyError(err)
		} else {
			ctx.ReplyJSON(result.Result1(dataResult))
		}
	} else {
		ctx.ReplyJSON(result.Result0(nil))
	}

	return nil
}
