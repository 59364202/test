package rainforecast_7day_province

var SQL_SelectRainforecast = `
WITH datas AS (SELECT
	rank() OVER (ORDER BY r.rainforecast_datetime DESC) AS day,
	r.rainforecast_datetime AS forecast_datetime,
	g.province_code AS province_id,
	g.province_name ->> 'th' AS province_name,
	r.rainforecast_value AS Rainfall,
	r.rainforecast_leveltext AS Rainfall_text
FROM
	rainforecast_7day r
LEFT JOIN lt_geocode g ON r.geocode_id = g."id"
WHERE
	g.province_code = $1
ORDER BY
	r.rainforecast_datetime DESC
LIMIT 7 )
SELECT rank() OVER (ORDER BY forecast_datetime ASC) AS day, forecast_datetime, province_id,province_name, Rainfall, Rainfall_text FROM datas 
`
