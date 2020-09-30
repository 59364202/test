package dam_daily_sum_by_region

import (
	"database/sql"
	_ "log"
	"time"
	// "fmt"
	"strconv"
	
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	float "haii.or.th/api/thaiwater30/util/float"
)

// get dam data by region by date
func GetSumByRegionRid(param *DamSumByRegionInput) (*DamSumByRegionByYearOutput, error) {
	// validate input
	if len(param.Year) == 0 {
		return nil, rest.NewError(422, "No Year", nil)
	}
	
	// add sql condition
	p := []interface{}{}
	var q string
	
	// get sql
	if len(param.Region_id) == 0 || param.Region_id == "0" {	// all region
		q = sqlGetDamSumByRegionAll
		// return nil, rest.NewError(422, "No Region ID", nil)
	} else {	// specific region
		// region_id
		p = append(p, param.Region_id)
	
		q = sqlGetDamSumByRegion
	}
	
	// get year param
	minYear, _ := min(param.Year)
	maxYear, _ := max(param.Year)
	
	// set max and min date for year param
	minDate := strconv.FormatInt(minYear, 10) + "-01-01"
	maxDate := strconv.FormatInt(maxYear, 10) + "-12-31"
	// example Year = [2016, 2017]
	// fmt.Println(minDate)		2016-01-01
	// fmt.Println(maxDate)		2018-12-31
	
	// add parameters for query
	p = append(p, minDate)
	p = append(p, maxDate)
	
	// open db
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	
	defer rows.Close()
	
	// loop through data and add data array group by year
	data := make([]*DamSumByRegionByYear, 0)
	// rs := []*DamSumByRegionGraphData{}
	// rs := make([]*DamSumByRegionGraphData, 0)
	dataRow := make(map[int64][]*DamSumByRegionGraphData)	// map[KeyType]ValueType
	var upperBound, lowerBound, normalBound float64
	
	for rows.Next() {
		// define query output structure
		var (
			dam_date         		time.Time
			dam_inflow       		sql.NullFloat64
			dam_released     		sql.NullFloat64
			dam_storage      		sql.NullFloat64
			dam_uses_water   		sql.NullFloat64
			total_max_storage		sql.NullFloat64
			total_min_storage		sql.NullFloat64
			total_normal_storage	sql.NullFloat64
			dam_inflow_acc       	sql.NullFloat64
			dam_released_acc     	sql.NullFloat64
		)
		
		// scan query output
		err = rows.Scan(&dam_date, &dam_inflow, &dam_released, &dam_storage, &dam_uses_water, &total_max_storage,
			&total_min_storage, &total_normal_storage, &dam_inflow_acc, &dam_released_acc)
		
		// check errors
		if err != nil {
			return nil, err
		}
		
		// if first row get upper-lower bound
		if (upperBound == 0) {
			upperBound = float.TwoDigit(total_max_storage.Float64)
			lowerBound = float.TwoDigit(total_min_storage.Float64)
			normalBound = float.TwoDigit(total_normal_storage.Float64)
		}
		
		// check year data is equal to year param
		row := &DamSumByRegionGraphData{}
		row.DamDate = dam_date.Format("2006-01-02")
		row.DamInflow = float.TwoDigit(dam_inflow.Float64)
		row.DamReleased = float.TwoDigit(dam_released.Float64)
		row.DamStorage = float.TwoDigit(dam_storage.Float64)
		row.DamUsesWater = float.TwoDigit(dam_uses_water.Float64)
		row.DamInflowAcc = float.TwoDigit(dam_inflow_acc.Float64)
		row.DamReleasedAcc = float.TwoDigit(dam_released_acc.Float64)
		
		dy := dam_date.Year()
		
		dataRow[int64(dy)] = append(dataRow[int64(dy)], row)
	}
	
	data = make([]*DamSumByRegionByYear, 0)
	for _, y := range param.Year {
		m := &DamSumByRegionByYear{}
		// m.Region_id = param.Region_id
		m.Year = y
		m.Data = dataRow[int64(y)]
		
		data = append(data, m)
	}
	
	bound := &DamSumByRegionBound{
		Upper: upperBound,
		Lower: lowerBound,
		Normal: normalBound,
	}
	
	output := &DamSumByRegionByYearOutput{
		GraphData: data,
		Bound: bound,
	}
	
	return output, nil
}

// get dam data compare same date each year
func GetCompareSumByRegionRid(param *DamCompareSumByRegionInput) ([]*DamCompareSumByRegionOutput, error) {
	// validate input
	if len(param.Day) == 0 {
		return nil, rest.NewError(422, "No Day", nil)
	}
	
	if len(param.Month)== 0 {
		return nil, rest.NewError(422, "No Month", nil)
	}
	
	// add sql condition
	p := []interface{}{}
	var q string
	
	// get sql
	if len(param.Region_id) == 0 || param.Region_id == "0" {	// all region
		q = sqlGetDamCompareSumByRegionAll
	} else {	// specific region
		// region_id
		p = append(p, param.Region_id)
		q = sqlGetDamCompareSumByRegion
	}
	
	// add parameters for query
	p = append(p, param.Day)
	p = append(p, param.Month)
	
	// open db
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	
	defer rows.Close()
	
	// loop through data 
	data := make([]*DamCompareSumByRegionOutput, 0)
	
	for rows.Next() {
		// define query output structure
		var (
			dam_year				int64
			dam_date         		time.Time
			dam_inflow       		sql.NullFloat64
			dam_released     		sql.NullFloat64
			dam_storage      		sql.NullFloat64
			dam_uses_water   		sql.NullFloat64
			dam_inflow_acc       	sql.NullFloat64
			dam_released_acc     	sql.NullFloat64
		)
		
		// scan query output
		err = rows.Scan(&dam_year, &dam_date, &dam_inflow, &dam_released, &dam_storage, &dam_uses_water, 
			&dam_inflow_acc, &dam_released_acc)
		
		// check errors
		if err != nil {
			return nil, err
		}
		
		// check year data is equal to year param
		row := &DamCompareSumByRegionOutput{}
		row.DamYear = dam_year
		row.DamDate = dam_date.Format("2006-01-02")
		row.DamInflow = float.TwoDigit(dam_inflow.Float64)
		row.DamReleased = float.TwoDigit(dam_released.Float64)
		row.DamStorage = float.TwoDigit(dam_storage.Float64)
		row.DamUsesWater = float.TwoDigit(dam_uses_water.Float64)
		row.DamInflowAcc = float.TwoDigit(dam_inflow_acc.Float64)
		row.DamReleasedAcc = float.TwoDigit(dam_released_acc.Float64)
		
		data = append(data, row)
	}
	
	return data, nil
}

func min(values []int64) (min int64, e error) {
    if len(values) == 0 {
        return 0, errors.New("Cannot detect a minimum value in an empty slice")
    }

    min = values[0]
    for _, v := range values {
	    if (v < min) {
	        min = v
        }
    }

    return min, nil
}

func max(values []int64) (max int64, e error) {
    if len(values) == 0 {
        return 0, errors.New("Cannot detect a maximum value in an empty slice")
    }

    max = values[0]
    for _, v := range values {
	    if (v > max) {
	        max = v
        }
    }

    return max, nil
}