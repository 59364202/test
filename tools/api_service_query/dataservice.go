package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/pqx"

	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	"haii.or.th/api/thaiwater30/model/order_detail"
)

func testDataserviceById(service_id int64) {

	serviceId := []int{1, 2, 4}
	fpath := "src/haii.or.th/api/thaiwater30/tools/api_service_query/service"
	slowQueryFileNmae := "slowquery.log"

	slowQuery, err := os.Create(fpath + "/" + slowQueryFileNmae)
	defer slowQuery.Close()
	handlerError(err)

	sid := leadZero(service_id)
	log.Println("start metadata : ", sid)
	sd, err := order_detail.GetMetadataByMetadata(service_id)
	sd.Province = sql.NullString{Valid: false, String: ""} // ไม่ใช้ province ในการคิวรี่
	sd.Basin = sql.NullString{Valid: false, String: ""}    // ไม่ใช้ basin ในการคิวรี่
	handlerError(err)

	mpath := fpath + "/" + sid
	err = os.MkdirAll(mpath, 755)
	handlerError(err)

	detail, err := os.Create(mpath + "/detail.log")
	defer detail.Close()
	handlerError(err)

	for _, sid := range serviceId {

		sd.Service_id = int64(sid)
		switch sid {
		case 1: // service ที่ระบบเวลาย้อนหลังได้ default 3 วัน
			sd.Detail_fromdate = sql.NullString{Valid: true, String: time.Now().AddDate(0, 0, -7).Format("2006-01-02")}
			sd.Detail_todate = sql.NullString{Valid: true, String: time.Now().Format("2006-01-02")}
		case 2: // ให้บริการแบบ download
			sd.Detail_fromdate = sql.NullString{Valid: true, String: "2017-01-01"}
			sd.Detail_todate = sql.NullString{Valid: true, String: "2017-01-07"}
		default: //  webserivce ข้อมูลล่าสุด
			sd.Detail_fromdate = sql.NullString{Valid: false, String: ""}
			sd.Detail_todate = sql.NullString{Valid: false, String: ""}
		}

		row, err := getMetadataQueryResult(slowQuery, detail, sd) // get query result
		//			fmt.Println("getMetadataQueryResult")
		handlerError(err)

		data, err := order_detail.ScanData(row)
		handlerError(err)

		f, err := os.Create(mpath + "/" + datatype.MakeString(sid) + ".json")
		defer f.Close()
		handlerError(err)
		writeData(detail, f, data.Data)
	}
}

func testDataservice() {

	serviceId := []int{1, 2, 4}
	fpath := "src/haii.or.th/api/thaiwater30/tools/api_service_query/service"
	slowQueryFileNmae := "slowquery.log"

	p := &model_metadata.Param_Metadata{}
	rs_metadata, err := model_metadata.GetMetadataShoppingTable(p)
	handlerError(err)

	err = os.MkdirAll(fpath, 755)
	handlerError(err)

	slowQuery, err := os.Create(fpath + "/" + slowQueryFileNmae)
	defer slowQuery.Close()
	handlerError(err)

	for _, v := range rs_metadata {
		sid := leadZero(v.Id)
		log.Println("start metadata : ", sid)
		sd, err := order_detail.GetMetadataByMetadata(v.Id)
		sd.Province = sql.NullString{Valid: false, String: ""} // ไม่ใช้ province ในการคิวรี่
		sd.Basin = sql.NullString{Valid: false, String: ""}    // ไม่ใช้ basin ในการคิวรี่
		handlerError(err)

		mpath := fpath + "/" + sid
		err = os.MkdirAll(mpath, 755)
		handlerError(err)

		detail, err := os.Create(mpath + "/detail.log")
		defer detail.Close()
		handlerError(err)

		for _, sid := range serviceId {

			sd.Service_id = int64(sid)
			switch sid {
			case 1: // service ที่ระบบเวลาย้อนหลังได้ default 3 วัน
				sd.Detail_fromdate = sql.NullString{Valid: true, String: time.Now().AddDate(0, 0, -7).Format("2006-01-02")}
				sd.Detail_todate = sql.NullString{Valid: true, String: time.Now().Format("2006-01-02")}
			case 2: // ให้บริการแบบ download
				sd.Detail_fromdate = sql.NullString{Valid: true, String: "2017-01-01"}
				sd.Detail_todate = sql.NullString{Valid: true, String: "2017-01-07"}
			default: //  webserivce ข้อมูลล่าสุด
				sd.Detail_fromdate = sql.NullString{Valid: false, String: ""}
				sd.Detail_todate = sql.NullString{Valid: false, String: ""}
			}

			row, err := getMetadataQueryResult(slowQuery, detail, sd) // get query result
			//			fmt.Println("getMetadataQueryResult")
			handlerError(err)

			data, err := order_detail.ScanData(row)
			handlerError(err)

			f, err := os.Create(mpath + "/" + datatype.MakeString(sid) + ".json")
			defer f.Close()
			handlerError(err)
			writeData(detail, f, data.Data)
		}

	}
}

func getMetadataQueryResult(slowQuery, f *os.File, p *order_detail.Strct_Data) (*sql.Rows, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
		itf []interface{}
	)
	itf = make([]interface{}, 0)

	if order_detail.IsMedia(p.Table_name.String) && p.Fields.Valid && p.Fields.String != "" && p.Fields.String != "{}" {
		// media หา media_type_id จาก import_setting
		err = order_detail.FindMediaTypeId(p)
		if err != nil {
			return nil, err
		}
	}

	if p.Detail_fromdate.Valid {
		p.Detail_fromdate.String = p.Detail_fromdate.String[0:10]
	}
	if p.Detail_todate.Valid {
		p.Detail_todate.String = p.Detail_todate.String[0:10]
	}

	//  สร้าง sql
	p.Sql, itf = order_detail.SQL_GenSQLSelectDataservice_All(p)
	f.WriteString("================================\n")
	f.WriteString(fmt.Sprintf("Service id : %d\n", p.Service_id))
	f.WriteString("Query : " + p.Sql + "\n")
	f.WriteString(fmt.Sprintf("Param : %v \n", itf))
	tn := time.Now()
	row, err = db.Query(p.Sql, itf...)
	d := time.Now().Sub(tn)
	f.WriteString(fmt.Sprintf("(%0.5f sec)\n", d.Seconds()))
	if d.Seconds() > 3 {
		slowQuery.WriteString("================================\n")
		slowQuery.WriteString(fmt.Sprintf("metadata id : %d , service id : %d \n", p.Metadata_id, p.Service_id))
		slowQuery.WriteString("Query : " + p.Sql + "\n")
		slowQuery.WriteString(fmt.Sprintf("Param : %v \n", itf))
		slowQuery.Sync()
	}
	if err != nil {
		return nil, err
	}
	f.Sync()
	return row, nil
}

func writeData(detail, f *os.File, p []map[string]interface{}) {

	//	f.WriteString(fmt.Sprintf("%s", p))
	str, err := json.Marshal(p)
	if err != nil {
		log.Println("Error encoding JSON")
	} else {
		f.WriteString(string(str))
	}

	f.Sync()
	detail.WriteString(fmt.Sprintln("rows", len(p)))
	detail.Sync()
}

func leadZero(i int64) string {
	s := datatype.MakeString(i)
	s = "000" + s
	return s[len(s)-4:]
}
