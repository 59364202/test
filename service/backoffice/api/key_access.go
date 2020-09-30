package api

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_key_access "haii.or.th/api/thaiwater30/model/agent"
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/thaiwater30/service/backoffice/data_integration/rdl"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
	"haii.or.th/api/util/shell"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/gen_key
// @Method			PUT
// @Summary			Update data of module APIS in the server
// @Description		Update secret_key of table api.agent
// @Parameter		-	form KeyAccessSwaggerGenKey
// @Produces		json
// @Response		200		ResultGenKey   successful operation
// @Response		404			-		the request service name was not found

type KeyAccessSwaggerGenKey struct {
	ID int64 `json:"id"` // example:`20` id ของ agent
}

type ResultGenKey struct {
	Result string                    `json:"result"` // example:`OK`
	Data   model_key_access.NewAgent `json:"data"`
}

func (srv *HttpService) genKey(ctx service.RequestContext) error {

	p := &KeyAccess{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	data, err := model_key_access.UpdateKey(ctx.GetUserID(), p.ID)
	if err != nil {
		return err
	}

	ctx.ReplyJSON(data)
	return nil
}

type KeyAccess struct {
	ID              int64  `json:"id"`
	UserAccount     string `json:"account"`
	AgentType       int64  `json:"agent_type"`
	PermissionRealm int64  `json:"permission_realm"`
	CallbackURL     string `json:"callback_url"`
	RequestOrigin   string `json:"request_origin"`
	KeyAccess       string `json:"key_access"`
	Agency          int64  `json:"agency_id"`
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/key_access
// @Method			POST
// @Summary			Add data to module APIS in the server
// @Description		Add data to table api.agent for new Agency.
// @Parameter		-	form KeyAccessSwagger
// @Produces		json
// @Response		200		ResultKeyAccess   successful operation
// @Response		404			-		the request service name was not found

type KeyAccessSwagger struct {
	UserAccount     string `json:"account"`          // example:`dataimport-haii` ชื่อผู้ใช้
	AgentType       int64  `json:"agent_type"`       // example:`1` ประเภท agent
	PermissionRealm int64  `json:"permission_realm"` // example:`3` รหัส realm
	CallbackURL     string `json:"callback_url"`     // example:`http://web.thaiwater.net/thaiwater30/apicb` callback url สำหรับ agent www
	RequestOrigin   string `json:"request_origin"`   // example:`http://web.thaiwater.net` request origin สำหรับ agent www
	KeyAccess       string `json:"key_access"`       // example:`117ad549211be3479086978c0a59a8385aa24b9d74` secret_key ของ agent ในการติดต่อ api
}

type ResultKeyAccess struct {
	Result string                    `json:"result"` // example:`OK`
	Data   model_key_access.NewAgent `json:"data"`
}

func (srv *HttpService) newKeyAccess(ctx service.RequestContext) error {

	p := &KeyAccess{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model_key_access.AddKeyAccess(ctx.GetUserID(), p.UserAccount, p.CallbackURL, p.RequestOrigin, p.AgentType, p.KeyAccess, p.PermissionRealm)

	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

type paramEditAgency struct {
	Agency    string `json:"agency"`
	AgentName string `json:"agent_name"`
	AgentKey  string `json:"agent_key"`
}

func (srv *HttpService) editAgency(ctx service.RequestContext) error {

	p := &paramEditAgency{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	ctx.LogRequestParams(p)
	rdls, _ := model_dataimport_config.GetAllRDLNodeFromSetting()
	idstr := ""
	params := []string{p.Agency, p.AgentName, p.AgentKey}
	for _, p := range params {
		idstr += " " + shell.QuoteArg(p)
	}

	for _, v := range rdls {
		rdl.RunRDLMgr(ctx, rdl.CmdAgencyEdit, idstr, v.Text.(string))
	}

	ctx.ReplyJSON(result.Result1(true))
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/key_access
// @Method			PUT
// @Summary			Update data of module APIS in the server
// @Description		Update data of table api.agent
// @Parameter		-	form KeyAccessSwaggerEdit
// @Produces		json
// @Response		200		ResultKeyAccess   successful operation
// @Response		404			-		the request service name was not found

type KeyAccessSwaggerEdit struct {
	ID              int64  `json:"id"`               // example:`20` id ของ agent
	UserAccount     string `json:"account"`          // example:`dataimport-haii` ชื่อผู้ใช้
	AgentType       int64  `json:"agent_type"`       // example:`1` ประเภท agent
	PermissionRealm int64  `json:"permission_realm"` // example:`3` รหัส realm
	CallbackURL     string `json:"callback_url"`     // example:`http://web.thaiwater.net/thaiwater30/apicb` callback url สำหรับ agent www
	RequestOrigin   string `json:"request_origin"`   // example:`http://web.thaiwater.net` request origin สำหรับ agent www
	KeyAccess       string `json:"key_access"`       // example:`117ad549211be3479086978c0a59a8385aa24b9d74` secret_key ของ agent ในการติดต่อ api
}

func (srv *HttpService) editKeyAccess(ctx service.RequestContext) error {

	p := &KeyAccess{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model_key_access.EditKeyAccess(ctx.GetUserID(), p.ID, p.CallbackURL, p.RequestOrigin, p.AgentType, p.KeyAccess, p.PermissionRealm, p.UserAccount)

	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/del_key
// @Method			PUT
// @Summary			Update data of module APIS in the server
// @Description		Remove secret_key of table api.agent
// @Parameter		-	form KeyAccessSwaggerDeletedKey
// @Produces		json
// @Response		200		DelKeySwaggerOutput   successful operation
// @Response		404			-		the request service name was not found
type KeyAccessSwaggerDeletedKey struct {
	ID int64 `json:"id"` // example:`20` id ของ agent
}

type DelKeySwaggerOutput struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`null`
}

func (srv *HttpService) delKey(ctx service.RequestContext) error {

	p := &KeyAccess{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)
	result, err := model_key_access.DeleteKey(ctx.GetUserID(), p.ID)

	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/key_access
// @Method			DELETE
// @Summary			Remove data of module APIS in the server
// @Description		Soft deleted data of table api.agent
// @Parameter		-	query KeyAccessSwaggerDeletedAgent
// @Produces		json
// @Response		200		DelKeySwaggerOutput   successful operation
// @Response		404			-		the request service name was not found

type KeyAccessSwaggerDeletedAgent struct {
	ID int64 `json:"id"` // example:`20` id ของ agent เช่น 20
}

func (srv *HttpService) deletedKeyAccess(ctx service.RequestContext) error {

	p := &KeyAccess{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}

	ctx.LogRequestParams(p)

	result, err := model_key_access.DeleteKeyAccess(p.ID)

	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/api/key_access
// @Method			GET
// @Summary			Show data of module APIS in the server
// @Description		Return data of table api.agent condition agent type as dataimport.
// @Produces		json
// @Response		200  ResultKeyAccessGetParams successful operation
// @Response		404			-		the request service name was not found

type ResultKeyAccessGetParams struct {
	Result string              `json:"result"` // example:`OK`
	Data   *KeyAccessGetParams `json:"data"`
}

type KeyAccessGetParams struct {
	KeyAccess       ResultKeyAccessSwgger             `json:"key_access"`
	AgentType       ResultAgentTypeOutPutSwgger       `json:"agent"`
	PermissionRealm ResultPermissionRealmOutputSwgger `json:"realm"`
	Agency          ResultStruct_AgencySwgger         `json:"agency"`
}

type ResultStruct_AgencySwgger struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_agency.Struct_Agency `json:"data"`
}

type ResultPermissionRealmOutputSwgger struct {
	Result string                   `json:"result"` // example:`OK`
	Data   []*PermissionRealmOutput `json:"data"`
}

type ResultAgentTypeOutPutSwgger struct {
	Result string             `json:"result"` // example:`OK`
	Data   []*AgentTypeOutPut `json:"data"`
}

type ResultKeyAccessSwgger struct {
	Result string                        `json:"result"` // example:`OK`
	Data   []*model_key_access.KeyAccess `json:"data"`
}

type AgentTypeOutPut struct {
	ID   int64  `json:"text"`  // example:`1` รหัสของประเภท agent
	Name string `json:"value"` // example:`www` ชื่อประเภท agent เช่น service,www,dataimport
}

type PermissionRealmOutput struct {
	ID   int64  `json:"text"`  // example:`1` รหัสของ permission realm
	Name string `json:"value"` // example:`haii-thaiwater30` ชื่อ permission realm
}

func (srv *HttpService) combineServiceKeyAccess(ctx service.RequestContext) error {

	type Result struct {
		KeyAccess       interface{} `json:"key_access"`
		AgentType       interface{} `json:"agent"`
		PermissionRealm interface{} `json:"realm"`
		Agency          interface{} `json:"agency"`
	}
	var err error

	data := &Result{}
	data.KeyAccess, err = model_key_access.GetKeyAccessTable()
	data.AgentType, err = model_key_access.GetAgentType()
	data.PermissionRealm, err = model_key_access.GetPermissionRealm()
	data.Agency, err = model_agency.GetAllAgency()
	if err != nil {
		return err
	}

	ctx.ReplyJSON(result.Result1(data))
	return nil
}
