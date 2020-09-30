package storm

import "time"
import "strconv"

// ออ้มเป็นคนให้เกณฑ์ ซึ่งใช้จากเกณฑ์เตือนภัย
// เดิม >= 91 and < 115.3
//2018-07-17 อ้อมแจ้งปรับเกณฑ์เฝ้าระวัง พายุ 109.4<x<=125.3
// ข้อมูลจาก wunderground
// ปัญหา query ข้อมูลย้อนหลัง 1 วัน เนื่องจากข้อมุล พายุ update ข้อมูบทุก 6 ชม. เวลา เป็น GMT ทำให้ ถ้า query ข้อมูลย้อนหลังน้อย กรณีข้อมูลพายุมาช้า จะทำให้ไม่แสดงข้อมูลพายุหน้าเว็บ
//2019-01-07 พบปัญหา query ได้ชื่อพายุมา ที่เป็นปัจจุบันมา แต่เงื่อนไขการดึงข้อมูลพายุ ไม่ได้ใส่เงื่อนไขวันที่ ทำให้ query ดึงข้อมูลพายุขื่อปัจจุบันแต่ข้อมุลเป็นข้อมูลของปีเก่า
//แก้ไข เพิ่มเงื่อนไขเวลาในการดึงข้อมูลเส้นทางพายุ
//AND storm_datetime > '2019-01-06' AND deleted_at = To_timestamp(0)
// ถ้าใช้วันที่เป็นเงื่อนไข จะไปผิดเคสมีพายุ ข้อมูลเส้นพายุจะไม่ครบ เปลี่ยนเป็นเงื่อนไขปีแทน

// เปบี่ยนให้เ็ป็น func เพื่อจะได้สร้าง วันที่ มาใช้ในคิวรี่
func SQL_GetStormCurrentDate() string {
	dt := time.Now().AddDate(0, 0, -1) // CURRENT_TIMESTAMP - interval '1 day'

	var SQL_GetStormCurrentDate = `
	SELECT 	storm_datetime,
		 storm_lat,
		 storm_directionlat,
		 storm_long,
		 storm_directionlong,
		 storm_name,
		 storm_pressure,
		 storm_wind
FROM (
	SELECT 
		 ROW_NUMBER() OVER(partition by storm_datetime,storm_lat,storm_long ORDER BY storm_datetime DESC,ID DESC ) as row_number,*
	FROM storm
	WHERE deleted_at = To_timestamp(0)
		AND storm_name :: text IN (
									SELECT DISTINCT( storm_name :: text )
									FROM   storm
									WHERE  storm_datetime > '` + dt.Format(time.RFC3339) + `' AND deleted_at = To_timestamp(0)
										AND (storm_long >= 91 AND storm_long <= 125.3)
										AND (storm_lat >= 5.4 AND storm_lat <= 21.9)									
						)
		AND (storm_long >= 91 AND storm_long <= 125.3)
		AND (storm_lat >= 5.4 AND storm_lat <= 21.9)
		AND extract(year from storm_datetime) = '` + strconv.Itoa(dt.Year()) + `'
		AND deleted_at = To_timestamp(0)
	ORDER BY storm_name :: text,storm_datetime
) storm_all
WHERE row_number = 1 
`
	// 2018-11-26
	//	เงื่อนไขการดึงข้อมูลพายุมีการระบุว่าให้ดึงข้อมูลพายุย้อนหลัง 1 วันเท่านั้น
	//	var SQL_GetStormCurrentDate = `
	//	SELECT 	storm_datetime,
	//		 storm_lat,
	//		 storm_directionlat,
	//		 storm_long,
	//		 storm_directionlong,
	//		 storm_name,
	//		 storm_pressure,
	//		 storm_wind
	//FROM (
	//	SELECT
	//		 ROW_NUMBER() OVER(partition by storm_datetime,storm_lat,storm_long ORDER BY storm_datetime DESC,ID DESC ) as row_number,*
	//	FROM storm
	//	WHERE
	//		storm_datetime > '` + dt.Format(time.RFC3339) + `'
	//		AND deleted_at = To_timestamp(0)
	//		AND storm_name :: text IN (SELECT DISTINCT( storm_name :: text )
	//											FROM storm
	//											WHERE storm_datetime > '` + dt.Format(time.RFC3339) + `' AND deleted_at = To_timestamp(0)
	//											)
	//		AND (storm_long >= 91 AND storm_long <= 125.3)
	//		AND (storm_lat >= 5.4 AND storm_lat <= 21.9)
	//	ORDER BY storm_name :: text,storm_datetime
	//) storm_all
	//WHERE row_number = 1
	//`

	return SQL_GetStormCurrentDate
}

// Last sql 18/07/2018
//SELECT storm_datetime,
//       storm_lat,
//       storm_directionlat,
//       storm_long,
//       storm_directionlong,
//       storm_name,
//       storm_pressure,
//       storm_wind
//FROM   storm
//WHERE  deleted_at = To_timestamp(0)
//       AND storm_name :: text IN (SELECT DISTINCT( storm_name :: text )
//                                  FROM   storm
//                                  WHERE  storm_datetime > CURRENT_TIMESTAMP - interval '1 day'
//                                         AND deleted_at = To_timestamp(0)
//																	)
//       AND ( storm_long > 109.4 AND storm_long <= 125.3)
//	   AND ( storm_lat >= 5.4 AND storm_lat <= 21.9)
//ORDER  BY storm_name :: text,
//          storm_datetime

// get latest storm (limit 5 records)
var sqlGetStormPeriod = `SELECT
		storm_name_alias
	FROM mv_storm_period`

// get storm history
var sqlGetStormHistory = `SELECT id
		, storm_datetime
		, storm_lat
		, storm_long
		, storm_directionlat
		, storm_directionlong
		, storm_pressure
		, storm_wind
		, color
		, to_lat
		, to_long
	 FROM v_storm_history `