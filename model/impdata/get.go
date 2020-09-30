package impdata

import (
	"database/sql"
	"haii.or.th/api/thaiwater30/util/selectoption"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"
	"strconv"
)

//  Parameters:
//		None
//  Return:
//		Array ImportDataOptionList
func GetImportDataSelectOption() ([]*ImportDataOptionList, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := getAgencyList
	p := []interface{}{}

	// query
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// define value output
	data := make([]*ImportDataOptionList, 0)
	dataAgency := &ImportDataOptionList{}
	dataMetadata := make([]*MetadataOptionList, 0)
	for rows.Next() {
		opt := &selectoption.Option{}
		optM := &MetadataOptionList{}
		var (
			agencyID       sql.NullInt64
			agencyName     sql.NullString
			downloadID     sql.NullInt64
			downloadIDName sql.NullString
			downloadScript sql.NullString
			metadataID     sql.NullInt64
		)
		// scan data
		rows.Scan(&agencyID, &agencyName, &downloadID, &downloadIDName, &downloadScript, &metadataID)
		// check agency
		if dataAgency.Agency == nil {
			// add data
			opt.Text = agencyName.String
			opt.Value = agencyID.Int64
			dataAgency.Agency = opt
			optM = &MetadataOptionList{}
			optM.Text = downloadIDName.String
			optM.Value = downloadID.Int64
			optM.MetadataID = metadataID.Int64
			optM.DownloadScript = downloadScript.String
			optM.DownloadCommand = "bin/rdl " + strconv.FormatInt(downloadID.Int64, 10) + " " + downloadScript.String
			// add data to array
			dataMetadata = append(dataMetadata, optM)
		} else if dataAgency.Agency.Text.(string) != agencyName.String {
			// add data
			dataAgency.Metadata = dataMetadata
			data = append(data, dataAgency)
			dataAgency = &ImportDataOptionList{}
			dataMetadata = make([]*MetadataOptionList, 0)
			opt := &selectoption.Option{}
			opt.Text = agencyName.String
			opt.Value = agencyID.Int64
			dataAgency.Agency = opt
			optM = &MetadataOptionList{}
			optM.Text = downloadIDName.String
			optM.Value = downloadID.Int64
			optM.MetadataID = metadataID.Int64
			optM.DownloadScript = downloadScript.String
			optM.DownloadCommand = "bin/rdl " + strconv.FormatInt(downloadID.Int64, 10) + " " + downloadScript.String
			// add data to array
			dataMetadata = append(dataMetadata, optM)
		} else {
			optM = &MetadataOptionList{}
			optM.Text = downloadIDName.String
			optM.Value = downloadID.Int64
			optM.MetadataID = metadataID.Int64
			optM.DownloadScript = downloadScript.String
			optM.DownloadCommand = "bin/rdl " + strconv.FormatInt(downloadID.Int64, 10) + " " + downloadScript.String
			// add data to array
			dataMetadata = append(dataMetadata, optM)
		}
	}
	if dataAgency.Agency != nil {
		dataAgency.Metadata = dataMetadata
		// add data to array
		data = append(data, dataAgency)
	}
	//return data
	return data, nil
}
