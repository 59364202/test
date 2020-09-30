package main

import (
	"os"
	"os/user"

	"haii.or.th/api/server"

	dashboard "haii.or.th/api/thaiwater30/service/dashboard"
	migration_log "haii.or.th/api/thaiwater30/service/migration_log"

	backoffice_api "haii.or.th/api/thaiwater30/service/backoffice/api"
	backoffice_data_integration_report "haii.or.th/api/thaiwater30/service/backoffice/data_integration_report"
	backoffice_data_management "haii.or.th/api/thaiwater30/service/backoffice/data_management"
	backoffice_data_service "haii.or.th/api/thaiwater30/service/backoffice/data_service"
	backoffice_dba "haii.or.th/api/thaiwater30/service/backoffice/dba"
	backoffice_event_management "haii.or.th/api/thaiwater30/service/backoffice/event_management"
	backoffice_metadata "haii.or.th/api/thaiwater30/service/backoffice/metadata"
	backoffice_tool "haii.or.th/api/thaiwater30/service/backoffice/tool"

	frontend_agency "haii.or.th/api/thaiwater30/service/frontend/agency"
	frontend_data_service "haii.or.th/api/thaiwater30/service/frontend/data_service"
	frontend_public "haii.or.th/api/thaiwater30/service/frontend/public"
	frontend_shared "haii.or.th/api/thaiwater30/service/frontend/shared"

	backoffice_dataimport_config "haii.or.th/api/thaiwater30/service/backoffice/data_integration/dataimport_config"
	backoffice_dataimport_config_migrate "haii.or.th/api/thaiwater30/service/backoffice/data_integration/dataimport_config_migrate"

	api_service "haii.or.th/api/thaiwater30/service/api_service"

	shared "haii.or.th/api/thaiwater30/service/shared"
	shared_api_service "haii.or.th/api/thaiwater30/service/shared/api_service"
	shared_image "haii.or.th/api/thaiwater30/service/shared/file"
	shared_file "haii.or.th/api/thaiwater30/service/shared/image"

	analyst "haii.or.th/api/thaiwater30/service/frontend/analyst"

	provinces "haii.or.th/api/thaiwater30/service/provinces"

	cache_system "haii.or.th/api/server/service/datacache"

	thaiwater30_cron "haii.or.th/api/thaiwater30/cron"

	"haii.or.th/api/thaiwater30/service/iframe"
	//	"haii.or.th/api/thaiwater30/service/manipulate"
	"haii.or.th/api/util/perm"

	"haii.or.th/api/thaiwater30/service/mobile"

	"haii.or.th/api/thaiwater30/service/test"

	"haii.or.th/api/server/model/backgroundjob"
	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/util/filepathx"
)

// Program version information.
// These three vairables will be replaced with  branch/tag information from git
// at compile time using  -ldflags -X main.BuildVersion="XXX"
var BuildVersion = ""

const DefaultTW30BkJobRunner = "tw30bgjob"
const TW30BkJobRunnerSettingName = "thaiwater30.model.backgroundjob.Runner"

func main() {
	perm.Umask(0027)

	srv := server.New("TW30", BuildVersion)

	dpt, err := srv.GetServiceDispatcher()
	if err != nil {
		os.Exit(1)
	}

	var bkrunner = DefaultTW30BkJobRunner
	if ux, err := user.Current(); err == nil {
		bkrunner = filepathx.JoinPath(ux.HomeDir, "go_local", "bin", DefaultTW30BkJobRunner)
	}
	setting.SetSystemDefault(TW30BkJobRunnerSettingName, bkrunner)
	backgroundjob.SetBackgroundRunnerSettingName(TW30BkJobRunnerSettingName)

	//Backoffice
	backoffice_dataimport_config.RegisterService(dpt)
	backoffice_dataimport_config_migrate.RegisterService(dpt)
	backoffice_api.RegisterService(dpt)
	backoffice_dba.RegisterService(dpt)
	backoffice_data_service.RegisterService(dpt)
	backoffice_event_management.RegisterService(dpt)
	backoffice_data_management.RegisterService(dpt)
	backoffice_tool.RegisterService(dpt)
	backoffice_metadata.RegisterService(dpt)
	backoffice_data_integration_report.RegisterService(dpt)
	//	manipulate.RegisterService(dpt)

	//Frontend
	frontend_agency.RegisterService(dpt)
	frontend_shared.RegisterService(dpt)
	frontend_data_service.RegisterService(dpt)
	frontend_public.RegisterService(dpt)

	//Shared
	shared.RegisterService(dpt)
	shared_image.RegisterService(dpt)
	shared_file.RegisterService(dpt)
	shared_api_service.RegisterService(dpt)
	iframe.RegisterService(dpt)

	api_service.RegisterService(dpt)

	//Dashboard
	dashboard.RegisterService(dpt)

	//Analyst
	analyst.RegisterService(dpt)

	//provinces
	provinces.RegisterService(dpt)

	//System
	cache_system.RegisterService(dpt)

	//Migration Log
	migration_log.RegisterService(dpt)

	//Mobile
	mobile.RegisterService(dpt)

	// test
	test.RegisterService(dpt)

	//
	thaiwater30_cron.RegisterCron()

	srv.Start()
}
