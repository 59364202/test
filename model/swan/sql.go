package swan

var SQL_SelectSwanCurrentDate = `
	SELECT     m.id , 
	m.swan_name , 
	m.swan_lat , 
	m.swan_long , 
	lg.province_code , 
	lg.province_name , 
	max_swan_highsig 
	FROM       ( 
			SELECT   Max(swan_highsig) AS max_swan_highsig , 
					swan_station_id 
			FROM     swan s 
			WHERE    swan_datetime >= $1 
			AND      swan_datetime < $2
			AND      deleted_at = '1970-01-01 07:00:00+07' 
			AND      swan_highsig > 2 
			GROUP BY swan_station_id) s 
	inner join m_swan_station m 
	ON         s.swan_station_id = m.id 
	AND        m.deleted_at = '1970-01-01 07:00:00+07' 
	inner join lt_geocode lg 
	ON         m.geocode_id = lg.id 
	WHERE      lg.province_code <> '99'
	`
