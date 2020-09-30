package migrate

import ()

var (
	NhcMediaType1 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and media_type_id = '1' "

	NhcMediaType2 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and media_type_id = '2' "

	NhcMediaType3 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and (media_type_id = '3' or media_type_id = '41' ) "

	NhcMediaType4 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and media_type_id = '4' AND ( ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) OR (media_time >= '07:00:00' AND media_time <= '18:59:59') )"

	NhcMediaType5 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '5' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') ) or (anhc_id = '19012' and media_type_id = '5' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )"

	NhcMediaType6 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '6' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') ) or (anhc_id = '19012' and media_type_id = '6' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )"

	NhcMediaType7 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012'  AND media_type_id = '11' AND (media_time >= '07:00:00' AND media_time <= '18:59:59')) or (anhc_id = '19012'  AND media_type_id = '11' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59'))) "

	NhcMediaType8 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012'  AND media_type_id = '12' AND (media_time >= '07:00:00' AND media_time <= '18:59:59')) or (anhc_id = '19012'  AND media_type_id = '12' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59'))) "

	NhcMediaType9 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012'  AND media_type_id = '13' AND (media_time >= '07:00:00' AND media_time <= '18:59:59')) or (anhc_id = '19012'  AND media_type_id = '13' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59'))) "

	NhcMediaType10 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '14' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '14' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '19012' and media_type_id = '31' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '31' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '11004' and media_type_id = '57' AND media_time >= '01:00:00' AND media_time < '06:59:59' )
or (anhc_id = '11004' and media_type_id = '57' AND media_time >= '07:00:00' AND media_time < '12:59:59' )
or (anhc_id = '11004' and media_type_id = '57' AND media_time >= '13:00:00' AND media_time < '18:59:59' )
or (anhc_id = '11004' and media_type_id = '57' AND ((media_time >= '00:00:00' AND media_time <= '00:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`

	NhcMediaType11 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '15' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '15' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`

	NhcMediaType14 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and media_type_id = '26' "

	NhcMediaType15 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '21' AND (media_time >= '07:00:00' AND media_time <= '18:59:59'))
or (anhc_id = '19012' and media_type_id = '21' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`

	NhcMediaType16 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '22' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') ) or (anhc_id = '19012' and media_type_id = '22' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )"

	NhcMediaType17 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '25' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '25' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '19012' and media_type_id = '30' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '30' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`

	NhcMediaType18 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '19012' and media_type_id = '23' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '23' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '19012' and media_type_id = '28' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '28' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`
		
	NhcMediaType19 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE  (anhc_id = '19012' and media_type_id = '24' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '24' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '19012' and media_type_id = '29' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '19012' and media_type_id = '29' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`
		
	NhcMediaType23 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '02005' and media_type_id = '33' AND (media_time >= '07:00:00' AND media_time < '19:00:00') )
or (anhc_id = '02005' and media_type_id = '33' AND media_time  >=  '19:00:00' )`
		
	NhcMediaType24 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '02005' and media_type_id = '34' AND media_time >= '07:00:00' AND media_time < '19:00:00' )
or (anhc_id = '02005' and media_type_id = '34' AND media_time  >=  '19:00:00' )`
		
	NhcMediaType25 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '02005' and media_type_id = '35' AND media_time >= '07:00:00' AND media_time < '19:00:00' )
or (anhc_id = '02005' and media_type_id = '35' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
`	
	NhcMediaType26 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '02005' and media_type_id = '36' AND media_time >= '07:00:00' AND media_time < '19:00:00' )
or (anhc_id = '02005' and media_type_id = '36' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`
		
	NhcMediaType27 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '02005' and media_type_id = '32' AND (media_time >= '07:00:00' AND media_time <= '18:59:59') )
or (anhc_id = '02005' and media_type_id = '32' AND ((media_time >= '00:00:00' AND media_time <= '06:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )
or (anhc_id = '11004' and media_type_id = '37' AND media_time >= '01:00:00' AND media_time < '06:59:59' )
or (anhc_id = '11004' and media_type_id = '37' AND media_time >= '07:00:00' AND media_time < '12:59:59' )
or (anhc_id = '11004' and media_type_id = '37' AND media_time >= '13:00:00' AND media_time < '18:59:59' )
or (anhc_id = '11004' and media_type_id = '37' AND ((media_time >= '00:00:00' AND media_time <= '00:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) `
		
	NhcMediaType28 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '11004' and media_type_id = '38' AND media_time >= '01:00:00' AND media_time < '06:59:59' )
or (anhc_id = '11004' and media_type_id = '38' AND media_time >= '07:00:00' AND media_time < '12:59:59' )
or (anhc_id = '11004' and media_type_id = '38' AND media_time >= '13:00:00' AND media_time < '18:59:59' )
or (anhc_id = '11004' and media_type_id = '38' AND ((media_time >= '00:00:00' AND media_time <= '00:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`
		
	NhcMediaType29 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = '11004' and media_type_id = '39' AND media_time >= '07:00:00' AND media_time < '12:59:59' )
or (anhc_id = '11004' and media_type_id = '39' AND ((media_time >= '00:00:00' AND media_time <= '00:59:59') OR (media_time >= '19:00:00' AND media_time <= '23:59:59')) )`
		
	NhcMediaType30 = `SELECT count(*) FROM media WHERE (anhc_id = '11004' and media_type_id = '40' AND location = 'cmp' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'cri' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'hhn' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'kkn' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'krb' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'lmp' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'phb' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'phs' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'pkt' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'ryg' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'skn' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'srn' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'stp' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'svp' )
or (anhc_id = '11004' and media_type_id = '40' AND location = 'ubn' )
or (anhc_id = '15009' and media_type_id = '40' AND location = 'njk' )`
	
	NhcMediaType32 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = 'A0001' and media_type_id = '55' "
	
	NhcMediaType33 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19006' and media_type_id = '56' "
	
	NhcMediaType140 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012'  AND media_type_id = '16' "
	
	NhcMediaType141 = `SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE (anhc_id = 'A0003'  AND media_type_id = '19' )
or (anhc_id = 'A0004'  AND media_type_id = '19' )
or (anhc_id = 'A0005'  AND media_type_id = '19' )`
	
	NhcMediaType142 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012'  AND media_type_id = '17' "
	
	NhcMediaType149 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = 'A0002' and media_type_id = '7' "
	
	NhcMediaType150 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = 'A0002' and media_type_id = '8' "
	
	NhcMediaType151 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = 'A0002' and media_type_id = '9' "
	
	NhcMediaType152 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = 'A0002' and media_type_id = '10' "
	
	NhcMediaType157 = "SELECT 'http://www.nhc.in.th'||media_path||filename FROM media WHERE anhc_id = '19012' and media_type_id = '27' "
)
