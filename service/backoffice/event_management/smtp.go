package event_management

import (
	"encoding/json"
	"haii.or.th/api/server/model"
	"haii.or.th/api/server/model/eventlog/sender"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/model/event_log_sink_method"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/sendmail"
	"haii.or.th/api/util/service"
	"strings"
)

const (
	SettingSMTPServerPrefix = "thaiwater30.service.event_management.smtpserver."
)

type emailServerConfigInfo struct {
	ConfigName string `json:"config_name"` //example:`smtp` config name
	sendmail.Config
}

type emailServerGetResult struct {
	Result string                   `json:"result"` // enum: ["OK","NO"]
	Data   []*emailServerConfigInfo `json:"data,omitempty"`
}

type emailServerPostResult struct {
	Result string                 `json:"result"` // enum: ["OK","NO"]
	Data   *emailServerConfigInfo `json:"data,omitempty"`
}

type emailServerDeleteResult struct {
	Result string `json:"result"` // enum: ["OK","NO"]
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_server
// @Method			GET
// @Summary			Retrive the SMTP server name and information
// @Produces		json
// @Response		200		emailServerGetResult successful operation
func (srv *HttpService) getEmailServer(ctx service.RequestContext) error {
	var r []*emailServerConfigInfo

	c := model.GetCipher()
	for k, v := range setting.FindSystemSetting(SettingSMTPServerPrefix) {
		var a emailServerConfigInfo

		if err := json.Unmarshal([]byte(v), &a.Config); err != nil {
			ctx.Logf("invalid smtp server configuration at %s ...%v", k, err)
			continue
		}

		if s, err := c.DecryptText(a.User); err == nil {
			a.User = s
		}
		if s, err := c.DecryptText(a.Password); err == nil {
			a.Password = s
		}

		n := strings.TrimPrefix(k, SettingSMTPServerPrefix)
		a.ConfigName = n
		r = append(r, &a)
	}

	ctx.ReplyJSON(&emailServerGetResult{"OK", r})
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_server
// @Method			POST
// @Summary			Add new SMTP server configuration
// @Produces		json
// @Parameter		-	 form emailServerConfigInfo
// @Response		200		emailServerPostResult successful operation
func (srv *HttpService) postEmailServer(ctx service.RequestContext) error {
	var p emailServerConfigInfo

	if err := ctx.GetRequestParams(&p); err != nil {
		return errors.Repack(err)
	}

	if p.ConfigName == "" {
		return rest.NewError(422, "missing server config name", nil)
	}

	return errors.Repack(srv.editEmailServer(&p, ctx))
}

func (srv *HttpService) editEmailServer(cfg *emailServerConfigInfo, ctx service.RequestContext) error {
	settingName := SettingSMTPServerPrefix + cfg.ConfigName
	if ctx.GetRequestMethod() == service.MethodPUT {
		if setting.GetSystemSetting(settingName) == "" {
			return rest.NewError(422, "unknown mail server config", nil)
		}
	}

	// Clear user/password before log parameter
	puser := cfg.User
	cfg.User = ""
	ppassword := cfg.Password
	cfg.Password = ""

	ctx.LogRequestParams(&cfg)

	c := model.GetCipher()
	var err error
	if cfg.User, err = c.EncryptText(puser); err != nil {
		return errors.Repack(err)
	}
	if cfg.Password, err = c.EncryptText(ppassword); err != nil {
		return errors.Repack(err)
	}

	var v = map[string]interface{}{
		settingName: &cfg.Config,
	}
	if err = setting.SetSystemSetting(ctx.GetUserID(), v, false); err != nil {
		return errors.Repack(err)
	}
	cfg.User = puser
	cfg.Password = ppassword

	ctx.ReplyJSON(&emailServerPostResult{"OK", cfg})
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_server/{config_name}
// @Method			PUT
// @Summary			Edit the SMTP server configuration
// @Produces		json
// @Parameter		config_name	path  string  email server config name ex. smtp1
// @Parameter		-	 form sendmail.Config
// @Response		200		emailServerPostResult successful operation
// @Response		422		- 				  config was not found
func (srv *HttpService) putEmailServer(ctx service.RequestContext) error {
	var p emailServerConfigInfo

	p.ConfigName = ctx.GetServiceParams("id")
	if p.ConfigName == "" {
		return rest.NewError(422, "missing mail server config name", nil)
	}

	if err := ctx.GetRequestParams(&p.Config); err != nil {
		return errors.Repack(err)
	}

	return errors.Repack(srv.editEmailServer(&p, ctx))
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/email_server/{config_name}
// @Method			DELETE
// @Summary			Delete the SMTP server configuration
// @Produces		json
// @Parameter		config_name	path string	email server config name ex. smtp1
// @Response		200		emailServerDeleteResult successful operation
// @Response	    422		-   config was not found or in use.

func (srv *HttpService) deleteEmailServer(ctx service.RequestContext) error {
	pname := ctx.GetServiceParams("id")

	settingName := SettingSMTPServerPrefix + pname
	if setting.GetSystemSetting(settingName) == "" {
		return rest.NewError(422, "unknown mail server config", nil)
	}

	//Search for sink method that use this setting
	var mp event_log_sink_method.Struct_EventLogSinkMethod_InputParam
	ms, err := event_log_sink_method.GetEventLogSinkMethod(&mp)
	if err != nil {
		return errors.Repack(err)
	}

	for _, m := range ms {
		var cfg sender.StdConfig
		if err := json.Unmarshal(m.Sink_params, &cfg); err == nil && cfg.SystemSettingName == settingName {
			return rest.NewError(422, "mail server config is in used", nil)
		}
	}

	setting.DeleteSystemSetting(ctx.GetUserID(), []string{settingName})
	ctx.ReplyJSON(&emailServerDeleteResult{"OK"})
	return nil
}
