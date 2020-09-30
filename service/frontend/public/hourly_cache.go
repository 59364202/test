package public

import (
	"haii.or.th/api/server/model/datacache"
	"haii.or.th/api/util/service"
	"time"
)

type hourlyCacheValidator struct {
}

func (v *hourlyCacheValidator) IsValid(lastupdate time.Time) bool {
	t := time.Now()
	// Cache must be update within the current hour
	return t.Sub(lastupdate) < time.Hour && t.Hour() == lastupdate.Hour()
}

func (v *hourlyCacheValidator) GetDescription() string {
	return "refresh every hour"
}

type getHourlyCacheBuildDataFn func() (interface{}, error)
type hourlyCache struct {
	buildDataFn getHourlyCacheBuildDataFn
}

func (c *hourlyCache) BuildData() (interface{}, error) {
	if c == nil || c.buildDataFn == nil {
		return nil, nil
	}
	return c.buildDataFn()
}

func getHourlyCache(cname string, fn getHourlyCacheBuildDataFn) ([]byte, time.Time, error) {
	if !datacache.IsRegistered(cname) {
		c := &hourlyCache{buildDataFn: fn}
		tb := []string{"cache.latest_rainfall24h", "cache.latest_waterlevel", "cache.latest_media", "cache.latest_waterquality", "public.m_tele_station", "public.m_waterquality_station", "public.m_dam"}
		datacache.RegisterDataCache(cname, c, tb, &hourlyCacheValidator{}, "1,11,21,31,41,51 * * * *")
	}

	return datacache.GetGZJSON(cname)
}

func replyWithHourlyCache(ctx service.RequestContext, cname string, fn getHourlyCacheBuildDataFn) error {
	b, t, err := getHourlyCache(cname, fn)
	if err != nil {
		return err
	}

	r := service.NewCachedResult(200, service.ContentJSON, b, t)
	ctx.Reply(r)
	return nil
}
