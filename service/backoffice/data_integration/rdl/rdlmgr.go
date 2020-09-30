package rdl

import (
	"encoding/json"
	"strings"
	"time"

	"haii.or.th/api/server/model"
	data "haii.or.th/api/server/model/dataimport"
	"haii.or.th/api/server/model/setting"
	model_dataimport_config "haii.or.th/api/thaiwater30/model/dataimport_config"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	"haii.or.th/api/util/shell"
)

const (
	// Default value of system_setting.
	DefaultDownloadRetryCount     = 3
	DefaultDownloadTimeoutSeconds = 30
	DefaultRDLMgrPath             = "dataimport/bin/rdlmgr"
	DefaultRDLPathPrefix          = "dataimport"
	DefaultRDLTimeoutSeconds      = 1
	RDLManagerProgram             = "bin/rdlmgr"

	CmdDwonlaodEdit = "download-edit"
	CmdCronEdit     = "cron-edit"
	CmdCronDelete   = "cron-delete"
	CmdPsRun        = "ps-run"
	CmdPsKill       = "ps-kill"
	CmdAgencyEdit   = "agency-edit"
)

const (
	// A system setting for this module.
	SettingPrefix                        = "server.service.dataimport"
	SettingDefaultDownloadRetryCount     = SettingPrefix + ".DefaultDownloadRetryCount"
	SettingDefaultDownloadTimeoutSeconds = SettingPrefix + ".DefaultDownloadTimeoutSeconds"
	SettingRDLNodes                      = SettingPrefix + ".RDLNodes"
	SettingDataPathPrefix                = SettingPrefix + ".DataPathPrefix"
	SettingUploadPathName                = SettingPrefix + ".UploadPathName"
)

func RunRDLMgr(ctx service.RequestContext, cmd string, idstr string, node string) error {
	inf, err := model_dataimport_config.GetDataImportDownloadConfig(idstr)
	if err != nil {
		return errors.Repackf(err, "invalid download id (%s)", idstr)
	}

	var params []string

	switch cmd {
	// case CmdDwonlaodEdit:
	// 	nrdi, _ := model_dataimport_config.NewRunDownloadID(inf.AgentUserID, datatype.MakeInt(idstr))
	// 	cmd = "download-edit"
	// 	params = append(params, nrdi.Agency)
	case CmdCronEdit:
		cmd = "cron-edit-with-detail"
		params = append(params, inf.DownloadName)
		params = append(params, inf.Description)
		params = append(params, "-")
		params = append(params, inf.CrontabSetting)
		params = append(params, datatype.MakeString(inf.MaxProcess))
		params = append(params, inf.DownloadScript)
		// "download_id", "name:", "description:", "agency", "interval", "max_process:0", "download_script"
	case CmdPsRun:
		params = append(params, inf.DownloadScript)
	case CmdCronDelete, CmdPsKill, CmdAgencyEdit, CmdDwonlaodEdit:
		// No additional parameter is required
	default:
		return errors.New("unknown rdlmgr command (" + cmd + ")")
	}
	//default first parameter = download_id
	cmd += " " + idstr
	for _, p := range params {
		cmd += " " + shell.QuoteArg(p)
	}

	if node == "" {
		rdlNode, _ := model_dataimport_config.GetRDLNodeFromSetting(datatype.MakeInt(inf.Node))
		node = rdlNode.Text
	}

	bx, err := runRemoteRDLCommand(ctx, node, cmd)
	if err != nil {
		return errors.Repack(err)
	}
	ctx.Logf("RDLMgr %s return %s", cmd, string(bx))
	return nil
}

// runRemoteRDLCommand runs a RDL command on a remote dataimport node from the
// parameters in the given RequestContext.
//
//  Parameters:
//		ctx
//			A pointer to service.RequestContext
//		nodeid
//			A remote dataimport node ID.
//		params
//			A rdlmgr command parameters
//  Return:
//		A JSON result of the command or nil if there is an error.
//
func runRemoteRDLCommand(ctx service.RequestContext, nodeid, params string) ([]byte, error) {
	var ninf data.RDLNodeInfo
	sname := SettingRDLNodes + "." + nodeid
	if err := json.Unmarshal([]byte(setting.GetSystemSetting(sname)), &ninf); err != nil {
		return nil, rest.NewError(422, "unknown remote rdl node", nil)
	}

	if ninf.User == "" || ninf.Host == "" || ninf.Key == "" {
		return nil, rest.NewError(500, "invalid remote rdl node setting", nil)
	}

	pp := ninf.PathPrefix
	if pp == "" {
		pp = DefaultRDLPathPrefix
	}
	if strings.Contains(pp, "..") {
		return nil, rest.NewError(500, "invalid remote rdl program path", nil)
	}
	cmd := filepathx.JoinPath("~", pp, RDLManagerProgram)

	tos := ninf.TimeoutSeconds
	if tos == 0 {
		tos = DefaultRDLTimeoutSeconds
	}
	if tos < 0 {
		tos = 0
	}

	key, err := model.GetCipher().DecryptText(ninf.Key)
	if err != nil || key == "" {
		key = ninf.Key
	}

	rssh, err := shell.NewRemoteSSH(ninf.Host, ninf.User, key,
		time.Second*time.Duration(tos))
	if err != nil {
		return nil, rest.NewError(500, "can not connect to remote rdl node", err)
	}
	defer rssh.Close()

	cmd += " -json " + params
	o, e, err := rssh.Run(cmd, time.Minute)

	rCode := 500
	if err != nil {
		msg := "can not execute remote command " + ninf.User + "@" + ninf.Host +
			" " + cmd + "\nError: " + err.Error()
		if len(o) > 0 {
			s := strings.Join(o, "\n")
			if p := strings.Index(s, "usage: "); p >= 0 {
				if p == 0 || s[p-1] == '\n' {
					rCode = 422
				}
			}

			msg += "\nSTDOUT:" + s
		}
		if len(e) > 0 {
			msg += "\nSTDERR:" + strings.Join(e, "\n")
		}
		log.Logf(msg)

		return nil, rest.NewError(rCode, msg, err)
	}

	s := strings.Join(o, "\n")
	if s == "" {
		s = "{}"
	}

	return []byte(s), nil
}
