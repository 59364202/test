package test

import (
	"haii.or.th/api/util/service"

	model_a "haii.or.th/api/thaiwater30/model/a"
)

type Werawan struct {
}

// @DocumentName	v1.test
// @Service 		test/werawan
// @Summary 		ทดสอบ service
// @Parameter		- query
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 model_a.werawan_province successful operation

func (srv *Werawan) handlerGetWerawan(ctx service.RequestContext) error {
	// call model
	province, err := model_a.Werawan_Province()

	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(province)
	}

	return nil
}
