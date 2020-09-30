// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package dam_yearly is a model for public.dam_yearly table. This table store dam data.
package dam_yearly

import (
	model_setting "haii.or.th/api/server/model/setting"
	"database/sql"
	"haii.or.th/api/thaiwater30/util/datetime"
	tw30_sort "haii.or.th/api/thaiwater30/util/sort"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"sort"
	"strconv"
	"time"
	"fmt"
	"math"
)

// Get Dam Daily return value dam daily by dam id and date range
//  Parameters:
//		inputData
//				GraphDamYearlyInput
//  Return:
//		data output for graph
func GetDamGraphYearly(inputData *GraphDamYearlyInput) (*GraphDamOutput, error) {

	// sql get dam yearly graph
	q := getDamYearlyGraph
	inputDamID := ""
	damIDMxMi := ""
	damID := []interface{}{}
	// loop condition with dam id
	for i, v := range inputData.DamID {
		if i > 0 {
			inputDamID += " OR dd.dam_id=$" + strconv.Itoa(i+3)
			damIDMxMi += " OR id=$" + strconv.Itoa(i+1)
			damID = append(damID, v)
		} else {
			inputDamID = "dd.dam_id=$" + strconv.Itoa(i+3)
			damIDMxMi = "id=$" + strconv.Itoa(i+1)
			damID = append(damID, v)
		}
	}

	if inputDamID != "" {
		q += " AND (" + inputDamID + ")"
	}
	q += " GROUP BY gs.date ORDER BY gs.date ASC "

	// process sql get dataimport download
	data := &GraphDamOutput{}
	dataOutput := make([]*DataOutput, 0)
	dd := make(tw30_sort.DataRange, 0)
	dd = append(dd, inputData.Year...)
	sort.Sort(dd)
	inputData.Year = dd
	// loop get data by year
	for _, v := range inputData.Year {
		u, err := getDamGraphOneYear(v, q, damID, inputData.DataType)
		if err != nil {
			return nil, err
		}
		dataOutput = append(dataOutput, u)
	}

	// rule curve sql
	q = getRuleCurveSql
	if inputDamID != "" {
		q += " WHERE dam_id=" + strconv.FormatInt(inputData.DamID[0], 10)
	}
	q += " ORDER BY date"

	// debug sql
	fmt.Println(q)

	// get rule curve by damID
	u, l, err := getRuleCurve(q)
	if err != nil {
		return nil, err
	}

	// debug query result
	fmt.Println(u)

	data.UpperRuleCurve = u
	data.LowerRuleCurve = l

	// get lower bound and upper bound
	q = getMaxMin
	if damIDMxMi != "" {
		q += " AND (" + damIDMxMi + ")"
	}

	max, min, normal, err := getMaxMinStorage(q, damID)
	if err != nil {
		return nil, err
	}
	data.GraphData = dataOutput
	data.LowerBound = min
	data.UpperBound = max
	data.NormalBound = normal

	// return data
	return data, nil
}

// Get medium dam data return value medium dam by medium dam id and date range
//  Parameters:
//		inputData
//				GraphDamMediumInput
//  Return:
//		data output for graph
func GetDamMediumGraph(inputData *GraphDamMediumInput) (*GraphDamOutput, error) {

	// sql get dam yearly graph
	q := getMediumDamGraph
	inputDamID := ""
	damIDMxMi := ""
	damID := []interface{}{}
	// loop condition mediumdam id
	for i, v := range inputData.DamID {
		if i > 0 {
			inputDamID += " OR dd.mediumdam_id=$" + strconv.Itoa(i+3)
			damIDMxMi += " OR id=$" + strconv.Itoa(i+1)
			damID = append(damID, v)
		} else {
			inputDamID = "dd.mediumdam_id=$" + strconv.Itoa(i+3)
			damIDMxMi = "id=$" + strconv.Itoa(i+1)
			damID = append(damID, v)
		}
	}

	if inputDamID != "" {
		q += " AND (" + inputDamID + ")"
	}
	q += " GROUP BY gs.date ORDER BY gs.date DESC "

	// process sql get dataimport download
	data := &GraphDamOutput{}
	dataOutput := make([]*DataOutput, 0)
	dd := make(tw30_sort.DataRange, 0)
	dd = append(dd, inputData.Year...)
	// sort year
	sort.Sort(dd)
	inputData.Year = dd
	// loop get data mediumdam by year
	for _, v := range inputData.Year {
		u, err := getDamMediumGraph(v, q, damID, inputData.DataType)
		if err != nil {
			return nil, err
		}
		dataOutput = append(dataOutput, u)
	}

	// get lower bound and upper bound
	q = getMaxMinMedium
	if damIDMxMi != "" {
		q += " AND (" + damIDMxMi + ")"
	}

	max, min, normal, err := getMaxMinStorage(q, damID)
	if err != nil {
		return nil, err
	}
	data.GraphData = dataOutput
	data.LowerBound = min
	data.UpperBound = max
	data.NormalBound = normal
	// return data
	return data, nil
}

// Get medium dam by year
//  Parameters:
//		yearInt
//				year type int
//		q
//				sql get mediumdam
//		pDam
//				value params for sql
//		dataType
//				field mediumdam
//  Return:
//		data output for graph
func getDamMediumGraph(yearInt int64, q string, pDam []interface{}, dataType string) (*DataOutput, error) {
	year := strconv.FormatInt(yearInt, 10)

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// process sql get dam yearly
	p := []interface{}{year + "-01-01", year + "-12-31"}
	p = append(p, pDam...)
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := &DataOutput{}
	data.Year = yearInt
	dataRow := make([]*DataYear, 0)
	defer rows.Close()
	for rows.Next() {
		var (
			mediumdam_date       time.Time
			mediumdam_storage    sql.NullFloat64
			mediumdam_inflow     sql.NullFloat64
			mediumdam_released   sql.NullFloat64
			mediumdam_uses_water sql.NullFloat64
		)
		rows.Scan(&mediumdam_date, &mediumdam_storage, &mediumdam_inflow, &mediumdam_released, &mediumdam_uses_water)

		dd := &DataYear{}
		dd.DamDate = datetime.DatetimeFormat(mediumdam_date, "date")
		// select field for return
		if dataType == "mediumdam_storage" {
			dd.Value = ValidData(mediumdam_storage.Valid, mediumdam_storage.Float64)
		} else if dataType == "mediumdam_inflow" {
			dd.Value = ValidData(mediumdam_inflow.Valid, mediumdam_inflow.Float64)
		} else if dataType == "mediumdam_released" {
			dd.Value = ValidData(mediumdam_released.Valid, mediumdam_released.Float64)
		} else if dataType == "mediumdam_uses_water" {
			dd.Value = ValidData(mediumdam_uses_water.Valid, mediumdam_uses_water.Float64)
		} else {
			return nil, rest.NewError(422, "No Data Type", nil)
		}
		dataRow = append(dataRow, dd)
	}
	if len(dataRow) > 0 {
		data.Data = dataRow
	} else {
		data.Data, _ = generateAllDateInYear(year)
	}
	return data, nil
}

// get dam one year
//  Parameters:
//		yearInt
//				year type int
//		q
//				sql get mediumdam
//		pDam
//				value params for sql
//		dataType
//				field dam
//  Return:
//		data output for graph
func getDamGraphOneYear(yearInt int64, q string, pDam []interface{}, dataType string) (*DataOutput, error) {
	year := strconv.FormatInt(yearInt, 10)

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// process sql get dam yearly
	p := []interface{}{year + "-01-01", year + "-12-31"}
	p = append(p, pDam...)
	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	data := &DataOutput{}
	data.Year = yearInt
	dataRow := make([]*DataYear, 0)
	defer rows.Close()
	for rows.Next() {
		var (
			dam_date         time.Time
			dam_storage      sql.NullFloat64
			dam_inflow       sql.NullFloat64
			dam_released     sql.NullFloat64
			dam_spilled      sql.NullFloat64
			dam_losses       sql.NullFloat64
			dam_evap         sql.NullFloat64
			dam_uses_water   sql.NullFloat64
			dam_inflow_avg   sql.NullFloat64
			dam_released_acc sql.NullFloat64
			dam_inflow_acc   sql.NullFloat64
		)
		rows.Scan(&dam_date, &dam_storage, &dam_inflow, &dam_released, &dam_spilled, &dam_losses, &dam_evap, &dam_uses_water, &dam_inflow_avg, &dam_released_acc, &dam_inflow_acc)

		dd := &DataYear{}
		dd.DamDate = datetime.DatetimeFormat(dam_date, "date")

		// select field for return
		if dataType == "dam_storage" {
			dd.Value = ValidData(dam_storage.Valid, dam_storage.Float64)
		} else if dataType == "dam_inflow" {
			dd.Value = ValidData(dam_inflow.Valid, dam_inflow.Float64)
		} else if dataType == "dam_released" {
			dd.Value = ValidData(dam_released.Valid, dam_released.Float64)
		} else if dataType == "dam_spilled" {
			dd.Value = ValidData(dam_spilled.Valid, dam_spilled.Float64)
		} else if dataType == "dam_losses" {
			dd.Value = ValidData(dam_losses.Valid, dam_losses.Float64)
		} else if dataType == "dam_evap" {
			dd.Value = ValidData(dam_evap.Valid, dam_evap.Float64)
		} else if dataType == "dam_uses_water" {
			dd.Value = ValidData(dam_uses_water.Valid, dam_uses_water.Float64)
		} else if dataType == "dam_inflow_avg" {
			dd.Value = ValidData(dam_inflow_avg.Valid, dam_inflow_avg.Float64)
		} else if dataType == "dam_released_acc" {
			dd.Value = ValidData(dam_released_acc.Valid, dam_released_acc.Float64)
		} else if dataType == "dam_inflow_acc" {
			dd.Value = ValidData(dam_inflow_acc.Valid, dam_inflow_acc.Float64)
		} else {
			return nil, rest.NewError(422, "No Data Type", nil)
		}

		dataRow = append(dataRow, dd)
	}

	if len(dataRow) > 0 {
		data.Data = dataRow
	} else {
		data.Data, _ = generateAllDateInYear(year)
	}

	return data, nil
}

// get rule curve data
//  Parameters:
//		q
//				sql get mediumdam
//		pDam
//				value params for sql
//  Return:
//		array data rulecurve upper and lower
func getRuleCurve(q string) ([]*RuleCurve, []*RuleCurve, error) {
	var dataUpper = make([]*RuleCurve, 0)
	var dataLower = make([]*RuleCurve, 0)
	var strDateFormat = model_setting.GetSystemSetting("setting.Default.DateFormat")
	
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// query
	_result, err := db.Query(q)

	// check if error
	if err != nil {
		return nil, nil, pqx.GetRESTError(err)
	}

	// close connection
	defer _result.Close()

	// loop through result
	for _result.Next() {
		var (
			date             time.Time
			upper_rule_curve sql.NullFloat64
			lower_rule_curve sql.NullFloat64
		)
		
		// Scan to execute query with variables
		_result.Scan(&date, &upper_rule_curve, &lower_rule_curve)

		var objUpperRuleCurve = &RuleCurve{}
		objUpperRuleCurve.Date = date.Format(strDateFormat)
		objUpperRuleCurve.Value = math.Ceil(upper_rule_curve.Float64*100)/100
		
		var objLowerRuleCurve = &RuleCurve{}
		objLowerRuleCurve.Date = date.Format(strDateFormat)
		objLowerRuleCurve.Value = math.Ceil(lower_rule_curve.Float64*100)/100

		dataUpper = append(dataUpper, objUpperRuleCurve)
		dataLower = append(dataLower, objLowerRuleCurve)
	}

	// return upper and lower
	return dataUpper, dataLower, nil
}

// function get max min storage
//  Parameters:
//		q
//				sql get mediumdam
//		pDam
//				value params for sql
//  Return:
//		data max min storage
func getMaxMinStorage(q string, p []interface{}) (interface{}, interface{}, interface{}, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, nil, nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// process sql get max min storage
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, nil, nil, pqx.GetRESTError(err)
	}

	defer rows.Close()
	var (
		max sql.NullFloat64
		min sql.NullFloat64
		normal sql.NullFloat64
	)
	for rows.Next() {
		rows.Scan(&max, &min, &normal)
	}
	// chack valid data and reutrun
	return ValidData(max.Valid, max.Float64), ValidData(min.Valid, min.Float64), ValidData(normal.Valid, normal.Float64), nil
}


// chack valid data by field when select from database
func ValidData(valid bool, value interface{}) interface{} {
	if valid {
		return value
	}
	return nil
}

// generate year
//  Parameters:
//		year
//				year
//  Return:
//		generate day all year input
func generateAllDateInYear(year string) ([]*DataYear, error) {
	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// process sql gen series
	q := getGenSeries
	p := []interface{}{year + "-01-01", year + "-12-31"}
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	defer rows.Close()
	dataRow := make([]*DataYear, 0)
	for rows.Next() {
		var (
			date time.Time
		)
		rows.Scan(&date)
		dd := &DataYear{}
		dd.DamDate = datetime.DatetimeFormat(date, "date")
		dataRow = append(dataRow, dd)
	}

	// return date range
	return dataRow, nil
}
