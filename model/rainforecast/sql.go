package rainforecast

import (
	"time"
)

func SQL_SelectRainforecastCurrentDate() string {

	now := time.Now().AddDate(0, 0, -1)

	var s = ` 
    SELECT l.province_code, l.province_name::text, max(r.rainforecast_value), l.tmd_area_code
	FROM rainforecast r 
	INNER JOIN lt_geocode l ON r.geocode_id::text = l.geocode 
	WHERE r.rainforecast_datetime >= ` + now.Format("2006-01-02") + `
	GROUP BY l.province_code, l.province_name::text, l.tmd_area_code 
	ORDER BY l.province_code `

	return s
}

var SQL_SelectRainforecast = `
SELECT
    province_code
  , province_name::jsonb
  , max(rainforecast_level) as level_max
FROM
    (
        SELECT
            lt.province_code
          , lt.province_name
          , r.rainforecast_level::int
        FROM
            rainforecast r
            LEFT JOIN
                lt_geocode lt
                ON
                    lt.id = r.geocode_id
        WHERE
            r.rainforecast_datetime >= $1
        order by
            r.created_at desc
          , r.geocode_id
    )
    a
WHERE
    rainforecast_level >= 4
GROUP BY
    province_code
  , province_name::jsonb
ORDER BY
    province_code
`
