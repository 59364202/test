// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	//	gLog "log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	util_file "haii.or.th/api/thaiwater30/util/file"
	"haii.or.th/api/util/errors"

	//	"haii.or.th/api/util/log"

	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/zip"

	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
)

// get metadata จาก order_detail id
func GetMetadataByOd(od_id int64) (*Strct_Data, error) {
	return scanMetadata(od_id, 0, "")
}

// get metadata จาก encrypt text
func GetMetadataByEId(e_id string) (*Strct_Data, error) {
	return scanMetadata(0, 0, e_id)
}

// get metadata จาก metadata id
func GetMetadataByMetadata(m_id int64) (*Strct_Data, error) {
	return scanMetadata(0, m_id, "")
}

// scan metadata
func scanMetadata(od_id, m_id int64, e_id string) (*Strct_Data, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		strSql string
		itf    []interface{}

		_service_id         sql.NullInt64
		_detail_fromdate    sql.NullString
		_detail_todate      sql.NullString
		_metadata_id        sql.NullInt64
		_agency_id          sql.NullString
		_connection_format  sql.NullString
		_table_name         sql.NullString
		_fields             sql.NullString
		_import_f           sql.NullString
		_orderdetail_id     sql.NullInt64
		_province           sql.NullString
		_basin              sql.NullString
		_is_enabled         sql.NullBool
		_created_at         sql.NullString
		_frequency          sql.NullString
		_dataset_id         int64
		_additional_dataset sql.NullString
	)
	itf = make([]interface{}, 0)
	// เลือกใช้ sql ตาม parameter
	if e_id != "" {
		strSql = SQL_SelectMetadata_FromEId
		itf = append(itf, e_id)
	} else if od_id != 0 {
		strSql = SQL_SelectMetadata
		itf = append(itf, od_id)
	} else if m_id != 0 {
		strSql = SQL_SelectMetadata_FromMetadata
		itf = append(itf, m_id)
	}
	err = db.QueryRow(strSql, itf...).
		Scan(&_service_id, &_detail_fromdate, &_detail_todate, &_metadata_id, &_agency_id, &_connection_format, &_table_name,
			&_fields, &_import_f, &_orderdetail_id, &_province, &_basin, &_is_enabled, &_created_at, &_frequency, &_dataset_id,
			&_additional_dataset)
		//	gLog.Println(strSql, itf)
		//	gLog.Println(err)
	if err != nil {
		return nil, err
	}

	data := &Strct_Data{}
	data.Service_id = _service_id.Int64
	data.Detail_fromdate = _detail_fromdate
	data.Detail_todate = _detail_todate
	data.Metadata_id = _metadata_id.Int64
	data.Agency_id = _agency_id
	data.Connection_format = _connection_format
	data.Table_name = _table_name
	data.Fields = _fields
	data.OrderDetail_id = _orderdetail_id.Int64
	data.Province = _province
	data.Basin = _basin
	data.IsEnabled = _is_enabled
	data.Frequency = _frequency
	data.CreateAt = pqx.NullStringToTime(_created_at)
	data.Dataset_id = _dataset_id
	data.AdditionalDataset = _additional_dataset

	var _im_f map[string]interface{}
	var temp_selecte_field string = ""
	json.Unmarshal([]byte(_import_f.String), &_im_f)

	// สร้าง select field ตาม dataset import_table
	var table = _table_name.String	
	st := model_metadata.GetTable(table)
	if st.IsMaster {
		// เป็น master table เอา Id ด้วย
		temp_selecte_field = "tt.id"
	}
	
	for _, v := range _im_f["fields"].([]interface{}) {
		s := v.(string)
		if strings.Trim(s, " ") == "" { // ชื่อฟิลด์ ต้องไม่เป็นช่องว่าง
			continue
		}
		//		if s == "#id" || s == "#row" || s == "qc_status" { // ไม่เอา #id, #row
		//			continue
		//		}
		if s[0:1] == "#" || s == "qc_status" || s == "refer_source" { // ไม่เอา ถ้าตัวหน้าสุดเป็น #, qc_status, refer_source
			if s == "qc_status" {
				data.HasQC = true
			}
			continue
		}
		//		if s[0:1] == "#" { // ถ้าตัวหน้าสุดเป็น # ให้เอา # ออก
		//			s = s[1:len(s)]
		//		}

		if temp_selecte_field != "" {
			temp_selecte_field += ", "
		}
		if s == "geocode_id" && st.HasProvince {
			// มี geocode_id และ setting ไว้ว่า ตารางนี้ มีการ join table geocode
			temp_selecte_field += "lg.province_name, lg.amphoe_name, lg.tumbon_name"
		} else {
			temp_selecte_field += "tt." + s + " "
		}

	}
	data.SelectFields = temp_selecte_field

	return data, nil
}

// get query result ของ SelectedFields
func GetMetadataQueryResult(p *Strct_Data) (*sql.Rows, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	var (
		row *sql.Rows
		itf []interface{}
	)
	itf = make([]interface{}, 0)

	if IsMedia(p.Table_name.String) && p.Fields.Valid && p.Fields.String != "" && p.Fields.String != "{}" {
		// media หา media_type_id จาก import_setting
		err = FindMediaTypeId(p)
		if err != nil {
			return nil, err
		}
	}

	if p.Detail_fromdate.Valid {
		//	2017-09-29T00:00:00Z เอา 10 ตัวแรก
		p.Detail_fromdate.String = p.Detail_fromdate.String[0:10]
	}
	if p.Detail_todate.Valid {
		//	2017-09-29T00:00:00Z เอา 10 ตัวแรก
		p.Detail_todate.String = p.Detail_todate.String[0:10]
	}

	//  สร้าง sql
	// p.Sql, itf = SQL_GenSQLSelectDataservice_All(datatype.MakeString(p.Service_id), p.SelectFields, p.Table_name.String, p.Agency_id.String, p.MediaTypeId,
	// 	p.Province.String, p.Basin.String, p.Detail_fromdate.String, p.Detail_todate.String, datatype.MakeString(p.Dataset_id), p.HasQC)
	p.Sql, itf = SQL_GenSQLSelectDataservice_All(p)
	//	log.Log(p.Sql, itf)
	row, err = db.Query(p.Sql, itf...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// media หา media_type_id จาก import_setting
func FindMediaTypeId(p *Strct_Data) error {
	var (
		_itfFields    []map[string]interface{}
		_mapField     map[string]interface{}
		media_type_id interface{}
	)
	err := json.Unmarshal([]byte(p.Fields.String), &_itfFields)
	if err != nil {
		return err
	}
L:
	for _, iv := range _itfFields {
		_mapField = iv
		if _mapField["name"].(string) == "media_type_id" {
			break L
		}
	}
	if _mapField["transform_method"] == "evaluate" {
		//	metadata id 525
		if p.Metadata_id == 525 {
			media_type_id = "160, 161, 162, 163"
		}

	} else {
		media_type_id = _mapField["transform_params"]
	}

	if str, ok := media_type_id.(string); ok {
		p.MediaTypeId = str
	} else {
		p.MediaTypeId = strconv.Itoa(media_type_id.(int))
	}

	return nil
}

// scan ด้าต้า จาก result query
func ScanData(row *sql.Rows) (*model_metadata.Struct_MetadataImportByAgency_Table, error) {
	rs := &model_metadata.Struct_MetadataImportByAgency_Table{}

	columns, _ := row.Columns()
	count := len(columns)
	values := make([]interface{}, count) // สำหรับ scan dynamic column
	valuePtrs := make([]interface{}, count)
	final_result := make([]map[string]interface{}, 0) // result ในรูปแบบ map[string]interface
	final_array := make([][]interface{}, 0)           // result ในรูปแบบ array

	for row.Next() {
		// scan data ลง array
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		row.Scan(valuePtrs...)
		tmp_struct := map[string]interface{}{}
		tmp_array := make([]interface{}, 0)

		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				var js map[string]interface{}
				// check json
				if json.Unmarshal(b, &js) == nil {
					tmp_struct[col] = js
					tmp_array = append(tmp_array, js) // json
				} else {
					tmp_struct[col] = string(b)
					tmp_array = append(tmp_array, string(b)) // string
				}
			} else if float, ok := val.(float64); ok {
				// type float64
				tmp_struct[col] = float
				tmp_array = append(tmp_array, float)
			} else {
				// default
				var v interface{} = val
				if len(col) > 4 {
					// column ยาวกว่า 4 ตัวอักษร
					if col[len(col)-4:] == "date" && col != "temperature_date" {
						// ต้องเช็คว่าเป็น date รึป่าว จะได้แปลงเป็น format 2006-01-02
						// ไม่เอา temperature_date เพราะในดาต้าเบส เก็บเป็น datetime
						if date, ok := val.(time.Time); ok {
							// it is of type time.Time
							v = date.Format("2006-01-02")
						}
					}
				}
				tmp_struct[col] = v
				tmp_array = append(tmp_array, v)
			}
		}
		final_result = append(final_result, tmp_struct)
		final_array = append(final_array, tmp_array)
	}
	rs.Columns = columns
	rs.Data = final_result
	rs.DataArray = final_array
	return rs, nil
}

//	สร้าง csv ไฟล์จาก struct
func GenerateFileFromData(od_id string, data *model_metadata.Struct_MetadataImportByAgency_Table, data_media []*model_metadata.Struct_Data_Media) (string, error) {
	if data_media != nil {
		return GenerateFileFromData_Media(od_id, data_media)
	}
	if data != nil {
		return GenerateFileFromData_Data(od_id, data)
	}
	return "", nil
}

// generate zip file from data เฉพาะที่ไม่ใช่ media
// return folder ที่เก้บ zip ไฟล์
func GenerateFileFromData_Data(id string, data *model_metadata.Struct_MetadataImportByAgency_Table) (string, error) {
	fx, err := ioutil.TempFile("", "order_detail")
	if err != nil {
		return "", err
	}
	defer func() { fx.Close(); os.Remove(fx.Name()) }()

	// จำนวน row สูงสุดต่อ 1 ไฟล์
	// -1 เพราะ มันมีหัวตาราง
	MaxRowInFile := setting.GetSystemSettingInt("dataservice.MaxRowInFile") - 1
	// จำนวนไฟล์ สูงสุดต่อ 1 zip
	MaxFileInZip := setting.GetSystemSettingInt("dataservice.MaxFileInZip")
	// folder ที่จะใช้เก็บ zip file
	folderZipFile := util_file.UploadPath + "/" + util_file.DataserviceUploadLetterPath + setting.GetSystemSetting("dataservice.TempZipFile") + "/" + id

	var writer *zip.ZipArchive
	var countZip int = 0

	zipName := ""
	dataArray := data.DataArray
	lenData := len(dataArray)
	var loopCountCsv int64 = 0 // นับจำนวน csv ไฟล์ใน zip
	var countFile int = 0      // นับจำนวนไฟล์ทั้งหมด (ใช้เป็นชื่อไฟล csv)

	if err != nil {
		return "", err
	}
	defer writer.Close()

	if _, err := os.Stat(folderZipFile); os.IsNotExist(err) { // สร้าง folder
		os.MkdirAll(folderZipFile, os.FileMode(util_file.UploadDirPerm))
	}

	for lenData >= 0 {
		countFile++
		if writer == nil || loopCountCsv >= MaxFileInZip { // เข้าลูปครั้งแรกหรือ วนลูป > MaxFileInZip สร้าง writer
			loopCountCsv = 0 // reset loopcount
			countZip++

			zipName = folderZipFile + "/" + id
			if countZip > 1 {
				zipName += "-" + strconv.Itoa(countZip)
			}
			zipName += ".zip"

			writer, err = zip.NewArchive(zipName)
			if err != nil {
				return "", err
			}
			defer writer.Close()
		}
		var arr int64
		if int64(lenData) > MaxRowInFile {
			arr = MaxRowInFile // ถ้า lenData > MaxRowInFile ใช้ MaxRowInFile
		} else {
			arr = int64(lenData) // ถ้า lenData <=  MaxRowInFile ใช้ lenData
		}

		err = writer.AddCSV(strconv.Itoa(countFile), data.Columns, dataArray[0:arr])
		if err != nil {
			return "", errors.New(err.Error())
		}
		dataArray = dataArray[arr:] // ตัด row ที่ใช้แล้วออกไป
		lenData = len(dataArray)    // นับ row ที่เหลือใหม่

		loopCountCsv++
		if lenData == 0 {
			lenData = -1 // ไม่มี row แล้ว ตั้งให้เป็นติดลบ เพื่อออกจากลูป
		}
	}

	return folderZipFile, nil
}

// generate zip file from data เฉพาะที่เป็น media
// return folder ที่เก้บ zip ไฟล์
func GenerateFileFromData_Media(id string, data_media []*model_metadata.Struct_Data_Media) (string, error) {
	// folder ที่จะใช้เก็บ zip file
	folderZipFile := util_file.UploadPath + "/" + util_file.DataserviceUploadLetterPath + setting.GetSystemSetting("dataservice.TempZipFile") + "/" + id
	if _, err := os.Stat(folderZipFile); os.IsNotExist(err) { // สร้าง folder
		os.MkdirAll(folderZipFile, os.ModePerm)
	}

	zipName := folderZipFile + "/" + id + ".zip"
	writer, err := zip.NewArchive(zipName)
	if err != nil {
		return "", err
	}
	defer writer.Close()

	for _, v := range data_media {
		fpath := filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), v.Path, v.Filename.(string))
		err = writer.AddPath(fpath)
		if err != nil {
			return "", err
		}
	}

	return folderZipFile, nil
}

// เช็คว่าเป็น เทเบิ้ล media รึปล่าว
func IsMedia(s string) bool {
	if s == "media" || s == "latest_media" || s == "media_other" || s == "media_animation" {
		return true
	}
	return false
}
