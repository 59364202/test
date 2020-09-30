package subbasin

import (

)

var (
	sqlGetSubbasin = "SELECT b.basin_code,b.basin_name,sb.subbasin_code,subbasin_name,sb.id FROM public.subbasin sb LEFT JOIN public.basin b ON sb.basin_id=b.id WHERE sb.deleted_at=to_timestamp(0) AND b.deleted_at=to_timestamp(0) AND b.basin_code != '99' ORDER BY basin_code"
)