package subbasin

import (
	"database/sql"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
)

// get sub basin
//  Parameters:
//		None
//  Return:
//		Array Basin
func GetSubbasin() ([]*Basin, error) {

	// connect database server
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}
	// sql subbasin
	q := sqlGetSubbasin
	p := []interface{}{}

	// query
	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()
	data := make([]*Basin, 0)
	sBasin := make([]*Subbasin, 0)
	basin := &Basin{}
	sBasinCode := ""
	
	for rows.Next() {
		var (
			basincode    sql.NullString
			basinname    pqx.JSONRaw
			subbasincode sql.NullString
			subbasinname pqx.JSONRaw
			sbID         sql.NullInt64
		)
		// scan data
		rows.Scan(&basincode, &basinname, &subbasincode, &subbasinname, &sbID)
		
		// check subbasin and basin duplicate
		if sBasinCode != "" {
			// new basin
			if sBasinCode != basincode.String {
				basin.Subbasin = sBasin
				data = append(data, basin)

				sBasin = make([]*Subbasin, 0)
				basin = &Basin{}
				sBasinCode = basincode.String
				basin.BasinCode = basincode.String
				basin.BasinName = basinname.JSON()
				sb := &Subbasin{}
				sb.SubbasinCode = subbasincode.String
				sb.SubbasinName = subbasinname.JSON()
				sb.ID = sbID.Int64
				sBasin = append(sBasin, sb)
				basin.Subbasin = sBasin
			} else {
				// add subbasin to basin
				sb := &Subbasin{}
				sb.SubbasinCode = subbasincode.String
				sb.SubbasinName = subbasinname.JSON()
				sb.ID = sbID.Int64
				sBasin = append(sBasin, sb)
				basin.Subbasin = sBasin
			}
		} else {
			// new basin if first row
			sBasinCode = basincode.String
			basin.BasinCode = basincode.String
			basin.BasinName = basinname.JSON()
			sb := &Subbasin{}
			sb.SubbasinCode = subbasincode.String
			sb.SubbasinName = subbasinname.JSON()
			sb.ID = sbID.Int64
			sBasin = append(sBasin, sb)
			basin.Subbasin = sBasin
		}
	}

	if len(sBasin) > 0 {
		basin.Subbasin = sBasin
		// add data to array
		data = append(data, basin)
	}

	return data, nil
}
