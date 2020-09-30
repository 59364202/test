package mobile

import (
	"time"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"

	uDatetime "haii.or.th/api/thaiwater30/util/datetime"

	"haii.or.th/api/thaiwater30/model/latest_media"
	"haii.or.th/api/thaiwater30/model/media_animation"
	"haii.or.th/api/thaiwater30/model/rainforecast"
	"haii.or.th/api/thaiwater30/model/swan"
)

type Forecast struct{}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rain_forecast
// @Summary			คาดการณ์ฝน 3 วัน ล่าสุด
// @Description		คาดการณ์ฝน 3 วัน ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	rainforecast.Struct_RainForcastRegion successful operation
func (srv *Forecast) handlerGetRainForecast(ctx service.RequestContext) error {
	result, err := rainforecast.GetRainForcastRegion()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rain_forecast_data
// @Summary			ข้อมูลคาดการณ์ฝน 3 วัน ล่าสุด
// @Description		ข้อมูลคาดการณ์ฝน 3 วัน ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	rainforecast.Struct_RainForecastData successful operation
func (srv *Forecast) handlerGetRainForecastData(ctx service.RequestContext) error {
	result, err := rainforecast.GetRainForecastData()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wave_forecast
// @Summary			คาดการณ์คลื่น 3 วัน ล่าสุด
// @Description		คาดการณ์คลื่น 3 วัน ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	latest_media.Struct_Media successful operation
func (srv *Forecast) handlerGetWaveForecast(ctx service.RequestContext) error {
	result, err := swan.Get_WaveForecast()
	if err != nil {
		return err
	}
	ctx.ReplyJSON(result)
	return nil
}

type Struct_Rain7day_forecast struct {
	Image_url string `json:"image_url,omitempty"` // example:`http://live1.haii.or.th/product/latest/wrfroms/v2/mobile/d03_day01.jpg`
	Date      string `json:"date,omitempty"`      // example:`25 พ.ค.`
	Video_url string `json:"video_url,omitempty"` // example:`http://www.nhc.in.th/product/latest/wrfroms/v2/ani_d01_large.mp4`
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/rain7day_forecast
// @Summary			คาดการณ์ฝน 7 วัน ล่าสุด
// @Description		คาดการณ์ฝน 7 วัน ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	Struct_Rain7day_forecast successful operation
func (srv *Forecast) handlerGetRain7dayForecast(ctx service.RequestContext) error {

	// ภาพ คาดการณ์ฝน
	_prerain, err := latest_media.GetPreRain()
	if err != nil {
		return err
	}

	// วิดีโอ คาดการณ์ฝน
	_animation, err := media_animation.GetPreRainAnimation()
	if err != nil {
		return err
	}
	// ใช้ ตัวที่ 0 ของ array animation เพราะ มัน sort ตามชื่อวิดีโอ ทำให้ตัวแรกเป็น "ani_d01_large.mp4" (27x27)
	return Data7DayForecast(ctx, _prerain, _animation[0])
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/wave7day_forecast
// @Summary			คาดการณ์คลื่น 7 วัน ล่าสุด
// @Description		คาดการณ์คลื่น 7 วัน ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	Struct_Rain7day_forecast successful operation
func (srv *Forecast) handlerGetWave7dayForecast(ctx service.RequestContext) error {

	// ภาพ คาดการณ์ฝน
	_prerain, err := latest_media.GetPreWave()
	if err != nil {
		return err
	}

	// วิดีโอ คาดการณ์ฝน
	_animation, err := media_animation.GetPreWaveAnimationMp4()
	if err != nil {
		return err
	}
	// ใช้ ตัวที่ 0 ของ array animation เพราะ มัน มันมีตัวเดียว
	return Data7DayForecast(ctx, _prerain, _animation[0])
}

func Data7DayForecast(ctx service.RequestContext, _img []*latest_media.Struct_Media, _animation *media_animation.Struct_Media) error {
	dt := time.Now()
	image_url := ctx.BuildURL(0, "thaiwater30/shared/image?dt=" + dt.String() + "&image=", true)

	// ภาพ
	img := make([]*Struct_Rain7day_forecast, 0)
	for _, p := range _img {
		dt := p.Dt
		d := &Struct_Rain7day_forecast{
			Image_url: image_url + datatype.MakeString(p.Path),
			Date:      dt.Format("2 ") + uDatetime.MonthTHShort(dt.Month()),
		}
		img = append(img, d)
	}

	// วีดีโอ
	animation := &Struct_Rain7day_forecast{
		Video_url: image_url + datatype.MakeString(_animation.Path),
	}

	result := []interface{}{
		img,
		animation,
	}
	ctx.ReplyJSON(result)
	return nil
}

type Struct_StormForecast struct {
	Image_full_size_url string `json:"image_full_size_url"` // example:`http://www.thaiwater.net/Tracking/Now/wp.png` URL รูปภาพพายุขนาดใหญ่
	Image_thumb_url     string `json:"image_thumb_url"`     // example:`http://www.thaiwater.net/Tracking/Now/wp_latest245.jpg` URL รูปภาพพายุเล็ก
}

// @DocumentName	v1.mobile
// @Service			mobile/{token}/th/storm_forecast
// @Summary			คาดการณ์พายุ ล่าสุด
// @Description		คาดการณ์พายุ ล่าสุด
// @Method			GET
// @Produces		json
// @Response		200	Struct_StormForecast successful operation
func (srv *Forecast) handlerGetStormForecast(ctx service.RequestContext) error {
	dt := time.Now()
	image_url := ctx.BuildURL(0, "thaiwater30/shared/image?dt=" + dt.String() + "&image=", true)

	_storm1, err := latest_media.GetLatestMedia(41, 62)
	if err != nil {
		return err
	}
	if len(_storm1) != 2 { // ต้องได้แค่ W.png กับ I.png
		return errors.New("_storm1 error")
	}
	var stormI, stormW *Struct_StormForecast
	if _storm1[0].Filename == "I.png" {
		// array 0 = I.png
		stormW = &Struct_StormForecast{
			Image_full_size_url: image_url + datatype.MakeString(_storm1[1].Path),
			Image_thumb_url:     image_url + datatype.MakeString(_storm1[1].PathThumb),
		}
		stormI = &Struct_StormForecast{
			Image_full_size_url: image_url + datatype.MakeString(_storm1[0].Path),
			Image_thumb_url:     image_url + datatype.MakeString(_storm1[0].PathThumb),
		}
	} else {
		// array 0 = W.png
		stormW = &Struct_StormForecast{
			Image_full_size_url: image_url + datatype.MakeString(_storm1[0].Path),
			Image_thumb_url:     image_url + datatype.MakeString(_storm1[0].PathThumb),
		}
		stormI = &Struct_StormForecast{
			Image_full_size_url: image_url + datatype.MakeString(_storm1[1].Path),
			Image_thumb_url:     image_url + datatype.MakeString(_storm1[1].PathThumb),
		}
	}

	_storm2, err := latest_media.GetLatestMedia(50, 141)
	if err != nil {
		return err
	}
	if len(_storm2) != 1 { // ต้องได้แค่ 00Latest.jpg
		return errors.New("_storm2 error")
	}
	var stormKochi *Struct_StormForecast = &Struct_StormForecast{
		Image_full_size_url: image_url + datatype.MakeString(_storm2[0].Path),
		Image_thumb_url:     image_url + datatype.MakeString(_storm2[0].PathThumb),
	}

	// เรียง W, I. kochi
	result := []*Struct_StormForecast{
		stormW,
		stormI,
		stormKochi,
	}
	ctx.ReplyJSON(result)
	return nil
}
