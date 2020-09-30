package test

import (
	model_a "haii.or.th/api/thaiwater30/model/a"
	"haii.or.th/api/util/service"
)

type Juiss struct {
}

// @DocumentName	v1.test
// @Service 		test/juiss
// @Summary 		ทดสอบ service
// @Parameter		- query model_a.Param_handlerGetJuiss
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 model_a.jtest_province successful operation

func (srv *Juiss) handlerGetJuiss(ctx service.RequestContext) error {
	param := &model_a.Param_handlerGetJuiss{}
	err := ctx.GetRequestParams(param)
	if err != nil {
		return err
	}

	province, err := model_a.Jtest_Province(param)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(province)
	}

	return nil
}
