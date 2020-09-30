package iframe

import (
	"time"
	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/service"

	model_dam "haii.or.th/api/thaiwater30/model/dam"
	model_dam_daily "haii.or.th/api/thaiwater30/model/dam_daily"
	model_dam_yearly "haii.or.th/api/thaiwater30/model/dam_yearly"

	result "haii.or.th/api/thaiwater30/util/result"
	uSetting "haii.or.th/api/thaiwater30/util/setting"
)

type Struct_Iframe_Dam struct {
	Dam      *Struct_Iframe_Dam_Dam      `json:"dam"`          // เขื่อน
	Datatype *Struct_Iframe_Dam_Datatype `json:"dam_datatype"` // ประเภทข้อมูล
}
type Struct_Iframe_Dam_Dam struct {
	Result string                     `json:"result"` // example:`OK`
	Data   []*model_dam.Struct_GetDam `json:"data"`   // เขื่อน
}
type Struct_Iframe_Dam_Datatype struct {
	Result string                         `json:"result"` // example:`OK`
	Data   []*uSetting.Struct_DamDataType `json:"data"`   // ประเภทข้อมูล
}

type Struct_If_Dam struct {
	Dam      *result.Result     `json:"dam"`          // เขื่อน
	Datatype *result.ResultJson `json:"dam_datatype"` // ประเภทข้อมูล
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/dam
// @Summary			รายชื่อเขื่อน, ประเภทข้อมูลเขื่อน
// @Description		เริ่มต้นหน้า iframe dam
// @				* รายชื่อเขื่อน
// @				* ประเภทข้อมูลเขื่อน
// @Method			GET
// @Produces		json
// @Response		200	Struct_Iframe_Dam successful operation
func (srv *HttpService) getDamIframe(ctx service.RequestContext) error {
	p := &model_dam_daily.Struct_DamDailyLastest_InputParam{}
	p.Agency_id = "12"
	rs := &Struct_If_Dam{}
	rs_dam, err := model_dam.GetDam("", p.Agency_id)
	if err != nil {
		rs.Dam = result.Result0(err.Error())
	} else {
		rs.Dam = result.Result1(rs_dam)
	}

	rs.Datatype = result.ResultJson1(setting.GetSystemSettingJson("Frontend.public.dam_data_type"))

	ctx.ReplyJSON(rs)

	return nil
}

type Struct_Iframe_Dam_Graph struct {
	Result string                           `json:"result"` // example:`OK`
	Data   *model_dam_yearly.GraphDamOutput `json:"data"`   // ข้อมูลกราฟ
}
type Param_Iframe_Dam_Graph struct {
	DataType string `json:"data_type"` // enum:[dam_storage,dam_inflow,dam_inflow_acc,dam_released,dam_released_acc] example: dam_storage ประเภทข้อมูล
	DamID    int64  `json:"dam_id"`    // example:10 หมายเลขรหัสของเขื่อนที่ต้องการข้อมูล
}

// @DocumentName	v1.public
// @Service			thaiwater30/iframe/dam_graph
// @Summary			กราฟเขื่อนย้อนหลัง 3 ปี
// @Method			GET
// @Parameter		-	query	Param_Iframe_Dam_Graph
// @Produces		json
// @Response		200	Struct_Iframe_Dam_Graph successful operation
func (srv *HttpService) getDamGraph(ctx service.RequestContext) error {
	p := &Param_Iframe_Dam_Graph{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)
	pa := &model_dam_yearly.GraphDamYearlyInput{}
	pa.DamID = []int64{p.DamID}
	pa.DataType = p.DataType
	var curYear int64 = int64(time.Now().Year())
	pa.Year = []int64{curYear, curYear - 1, curYear - 2}
	rs, err := model_dam_yearly.GetDamGraphYearly(pa)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(rs))
	}

	return nil
}
