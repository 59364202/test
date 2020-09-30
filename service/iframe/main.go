package iframe

import (
	//	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
)

const (
	DataServiceName = "thaiwater30/iframe"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)
}

func (srv *HttpService) handleGetData(ctx service.RequestContext) error {

	service_id := ctx.GetServiceParams("service")

	switch service_id {
	case "dam":
		return srv.getDamIframe(ctx)
	case "dam_graph":
		return srv.getDamGraph(ctx)
	case "rain24":
		return srv.getRain24(ctx)
		// graph rain hour ย้อนหลัง 24 ชม. จากเวลาปัจจุบัน
	case "rain24_graph":
		return srv.getRain24Graph(ctx)
	case "waterlevel":
		return srv.getWaterlevel(ctx)
	case "waterlevel_graph": // สำหรับหน้า main ข้อมูลกราฟระดับน้ำ เฉพาะข้อมูล 3 วันล่าสุด ไม่สามารถใส่วันที่ได้
		return srv.getWaterlevelGraph(ctx)
	case "waterquality":
		return srv.getWaterQuality(ctx)
	case "waterquality_graph":
		return srv.getWaterQualityGraph(ctx)
	default:
		return rest.NewError(404, "Unknown service id", nil)
	}

	return nil
}
