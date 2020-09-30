package event_management

import (
	"encoding/json"
	"haii.or.th/api/server/model"
	"haii.or.th/api/server/model/eventlog/sender"
	"haii.or.th/api/server/model/eventlog/sender/sms"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/thaiwater30/model/event_log_sink_method"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	"strings"
)

const (
	SettingSmsServerPrefix = "thaiwater30.service.event_management.smsserver."
)

type smsServerConfigInfo struct {
	ConfigName string `json:"config_name"` // example:`config name` config name ex. cfg_sms_sv1
	sms.Params
}

type smsServerGetResult struct {
	Result string                 `json:"result"` // enum: ["NO","OK"]
	Data   []*smsServerConfigInfo `json:"data,omitempty"`
}

type smsServerPostResult struct {
	Result string               `json:"result"` // enum: ["NO","OK"]
	Data   *smsServerConfigInfo `json:"data,omitempty"`
}

type smsServerDeleteResult struct {
	Result string `json:"result"` // enum: ["NO","OK"]
}

type smsServerConfigInfoSwagger struct {
	ConfigName string `json:"config_name"` // example:`config name` config name ex. cfg_sms_sv1
	// SMS server
	Server string `json:"server"` // example:`sms.server` sms server ex. sms.server1
	// User
	User string `json:"user"` // example:`user` username ex. usercim
	// Password
	Password string `json:"password"` // example:`password` password ex. xxxx
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sms_server
// @Method			GET
// @Summary			Retrive the SMTP server name and information
// @Produces		json
// @Response		200		smsServerGetResult successful operation
func (srv *HttpService) getSmsServer(ctx service.RequestContext) error {
	var r []*smsServerConfigInfo

	c := model.GetCipher()
	for k, v := range setting.FindSystemSetting(SettingSmsServerPrefix) {
		var a smsServerConfigInfo

		if err := json.Unmarshal([]byte(v), &a.Params); err != nil {
			ctx.Logf("invalid smtp server configuration at %s ...%v", k, err)
			continue
		}

		if s, err := c.DecryptText(a.User); err == nil {
			a.User = s
		}
		if s, err := c.DecryptText(a.Password); err == nil {
			a.Password = s
		}

		n := strings.TrimPrefix(k, SettingSmsServerPrefix)
		a.ConfigName = n
		r = append(r, &a)
	}

	ctx.ReplyJSON(&smsServerGetResult{"OK", r})
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sms_server
// @Method			POST
// @Summary			Add new SMTP server configuration
// @Produces		json
// @Parameter		-	 form smsServerConfigInfo
// @Response		200		smsServerPostResult successful operation
func (srv *HttpService) postSmsServer(ctx service.RequestContext) error {
	var p smsServerConfigInfo

	if err := ctx.GetRequestParams(&p); err != nil {
		return errors.Repack(err)
	}

	if p.ConfigName == "" {
		return rest.NewError(422, "missing server config name", nil)
	}

	return errors.Repack(srv.editSmsServer(&p, ctx))
}

func (srv *HttpService) editSmsServer(cfg *smsServerConfigInfo, ctx service.RequestContext) error {
	settingName := SettingSmsServerPrefix + cfg.ConfigName
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
		settingName: &cfg.Params,
	}
	if err = setting.SetSystemSetting(ctx.GetUserID(), v, false); err != nil {
		return errors.Repack(err)
	}
	cfg.User = puser
	cfg.Password = ppassword

	ctx.ReplyJSON(&smsServerPostResult{"OK", cfg})
	return nil
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sms_server/{config_name}
// @Method			PUT
// @Summary			Edit the SMTP server configuration
// @Produces		json
// @Parameter		config_name	path  string  email server config name ex. cfg_sms_sv1
// @Parameter		-	 form smsServerConfigInfoSwagger
// @Response		200		smsServerPostResult successful operation
// @Response		422		- 				  config was not found
func (srv *HttpService) putSmsServer(ctx service.RequestContext) error {
	var p smsServerConfigInfo

	p.ConfigName = ctx.GetServiceParams("id")
	if p.ConfigName == "" {
		return rest.NewError(422, "missing mail server config name", nil)
	}

	if err := ctx.GetRequestParams(&p.Params); err != nil {
		return errors.Repack(err)
	}

	return errors.Repack(srv.editSmsServer(&p, ctx))
}

// @DocumentName	v1.webservice
// @Service			/thaiwater30/backoffice/event_management/sms_server/{config_name}
// @Method			DELETE
// @Summary			Delete the SMTP server configuration
// @Produces		json
// @Parameter		config_name	path string	sms server config name ex. sms1
// @Response		200		smsServerDeleteResult successful operation
// @Response	    422		-   config was not found or in use.

func (srv *HttpService) deleteSmsServer(ctx service.RequestContext) error {
	pname := ctx.GetServiceParams("id")

	settingName := SettingSmsServerPrefix + pname
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
	ctx.ReplyJSON(&smsServerDeleteResult{"OK"})
	return nil
}
