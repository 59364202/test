package data_management

import (
	model_import_data "haii.or.th/api/thaiwater30/model/impdata"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/impdata_option_list
// @Summary			get command for download
// @Description		command import data
// @Method			GET
// @Produces		json
// @Response		200		ImportDataOptionListSwagger successful operation

type ImportDataOptionListSwagger struct {
	Result string `json:"result"`  //example:`OK`
	Data []model_import_data.ImportDataOptionListSwagger `json:"data"`
}

func (srv *HttpService) importDataOptionList(ctx service.RequestContext) error {

	//Get Data
	dataResult, err := model_import_data.GetImportDataSelectOption()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}
