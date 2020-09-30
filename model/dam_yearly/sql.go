package dam_yearly

var (
	// temporary by pass QC Rule
	// AND ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
	getDamYearlyGraph = `
	SELECT gs.date,sum(dd.dam_storage),sum(dd.dam_inflow),sum(dd.dam_released),sum(dd.dam_spilled),sum(dd.dam_losses),sum(dd.dam_evap),sum(dd.dam_uses_water),sum(dam_inflow_avg),sum(dam_released_acc),sum(dam_inflow_acc)
	FROM  (SELECT generate_series($1::date,$2, '1 day') AS date) gs
	LEFT JOIN public.dam_daily dd ON gs.date=dd.dam_date AND dd.deleted_at=to_timestamp(0)
	`

	getMediumDamGraph = `
	SELECT gs.date,sum(dd.mediumdam_storage),sum(dd.mediumdam_inflow),sum(dd.mediumdam_released),sum(dd.mediumdam_uses_water)
	FROM  (SELECT generate_series($1::date,$2, '1 day') AS date) gs
	LEFT JOIN public.medium_dam dd ON gs.date=dd.mediumdam_date AND dd.deleted_at=to_timestamp(0)
	WHERE  ( qc_status IS NULL OR qc_status->>'is_pass' = 'true' )
	`

	// getRuleCurveSql = "SELECT rc_datetime,urc_new,lrc_new FROM public.rulecurve WHERE rc_unit='mcm' AND deleted_at=to_timestamp(0) AND rc_datetime BETWEEN $1 AND $2 "
	getRuleCurveSql = "SELECT date, revised_urc AS upper_rule_curve, revised_lrc AS lower_rule_curve FROM public.dam_rulecurve"

	getMaxMin = "SELECT sum(max_storage),sum(min_storage), sum(normal_storage) FROM public.m_dam WHERE deleted_at=to_timestamp(0) "

	getMaxMinMedium = "SELECT sum(max_storage),sum(min_storage), sum(normal_storage) FROM public.m_medium_dam WHERE deleted_at=to_timestamp(0) "

	getGenSeries = "SELECT generate_series($1::date,$2, '1 day') AS date"
)
