package mobile

import (
	"log"
	"time"

	"haii.or.th/api/util/service"

	"haii.or.th/api/thaiwater30/model/waterquality"
)

type Provinces struct{}

func (srv *Provinces) handlerGetWaterqualityLatest(ctx service.RequestContext) error {
	result, err := waterquality.Get_WaterqualityLatest()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	//	go testSleep()
	return nil
}

func testSleep() {
	var z *favoriteLatLongItf
	if z == nil {
		time.Sleep(5 * time.Second)
		log.Println(z.Flag)
	}
}
