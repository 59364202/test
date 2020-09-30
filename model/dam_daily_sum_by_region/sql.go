package dam_daily_sum_by_region

// get total sum by region
var sqlGetDamSumByRegion = `SELECT dam_date
	, total_dam_inflow, total_dam_released, total_dam_storage, total_dam_uses_water
	, total_max_storage, total_min_storage, total_normal_storage
	, total_dam_inflow_acc, total_dam_released_acc
	FROM dam_daily_sum_by_region
	WHERE area_code = $1
	AND dam_date BETWEEN $2 AND $3
	AND deleted_at = '1970-01-01 07:00:00+07' 
	ORDER BY dam_date`

// get total all region
var sqlGetDamSumByRegionAll = `SELECT dam_date
	, sum(total_dam_inflow) AS total_dam_inflow, sum(total_dam_released) AS total_dam_released
	, sum(total_dam_storage) AS total_dam_storage, sum(total_dam_uses_water) AS total_dam_uses_water
	, sum(total_max_storage) AS total_max_storage, sum(total_min_storage) AS total_min_storage
	, sum(total_normal_storage) AS total_normal_storage
	, sum(total_dam_inflow_acc) AS total_dam_inflow_acc
	, sum(total_dam_released_acc) AS total_dam_released_acc
	FROM dam_daily_sum_by_region
	WHERE dam_date BETWEEN $1 AND $2
	AND deleted_at = '1970-01-01 07:00:00+07'
	GROUP BY dam_date
	ORDER BY dam_date`

// get compare by year same date by region
var sqlGetDamCompareSumByRegion = `SELECT EXTRACT(YEAR FROM dam_date) AS dam_year, dam_date
	, total_dam_inflow, total_dam_released, total_dam_storage, total_dam_uses_water
	, total_dam_inflow_acc, total_dam_released_acc
	FROM dam_daily_sum_by_region
	WHERE area_code = $1
	AND EXTRACT(DAY FROM dam_date) = $2 
	AND EXTRACT(MONTH FROM dam_date) = $3
	AND deleted_at = '1970-01-01 07:00:00+07' 
	ORDER BY dam_date`

// get compare by year total all region
var sqlGetDamCompareSumByRegionAll = `SELECT EXTRACT(YEAR FROM dam_date) AS dam_year, dam_date
	, sum(total_dam_inflow) AS total_dam_inflow, sum(total_dam_released) AS total_dam_released
	, sum(total_dam_storage) AS total_dam_storage, sum(total_dam_uses_water) AS total_dam_uses_water
	, sum(total_dam_inflow_acc) AS total_dam_inflow_acc
	, sum(total_dam_released_acc) AS total_dam_released_acc
	FROM dam_daily_sum_by_region
	WHERE EXTRACT(DAY FROM dam_date) = $1 
	AND EXTRACT(MONTH FROM dam_date) = $2
	AND deleted_at = '1970-01-01 07:00:00+07'
	GROUP BY dam_date
	ORDER BY dam_date`
