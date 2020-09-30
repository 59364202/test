package public

import (
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.public
// @Service			thaiwater30/public/thailand
// @Summary			หน้าภาพรวมประเทศไทย
// @Description
// @				* ข้อมูลฝน 24 ชัวโมง
// @				* ข้อมูลเขื่อน
// @				* ข้อมูลระดับน้ำ
// @				* ข้อมูลคุณภาพน้ำ
// @				* พายุ
// @				* คาดการณ์ฝนล่วงหน้า
// @				* คาดการณ์คลื่นล่วงหน้า
// @				* พื้นที่ประกาศภัย
// @Method			GET
// @Produces		json
// @Response		200		Index successful operation
func (srv *HttpService) getThailand(ctx service.RequestContext) error {
	return replyWithHourlyCache(ctx, thailandPageCacheName, getThailandPageCacheBuildData)
}
