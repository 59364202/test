package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/util/pqx"

	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

const basePath = "src/haii.or.th/api/thaiwater30/service/api_service/"

var SQL_SelectMetadata = `
SELECT     
           m.id, 
           m.metadataservice_name,
           dd.import_setting#>'{configs,0,imports,0}'->>'destination' as table_name, 
           CASE 
                      WHEN dd.import_table->> 'tables' != '{}' THEN dd.import_table->'tables'-> json_object_keys(dd.import_table-> 'tables') 
                      ELSE NULL 
           END AS import_f
FROM       metadata m 
INNER JOIN api.dataimport_dataset dd 
ON         m.dataimport_dataset_id = dd.id 
order by m.id
`

var SQL_SelectComment = `
SELECT c.table_name,c.column_name,pgd.description, c.data_type
FROM pg_catalog.pg_statio_all_tables as st
  inner join pg_catalog.pg_description pgd on (pgd.objoid=st.relid)
  inner join information_schema.columns c on (pgd.objsubid=c.ordinal_position
    and  c.table_schema=st.schemaname and c.table_name=st.relname and c.table_schema = 'public')
`

// map datatype to go language type
var M_DataType map[string]string = map[string]string{
	"json": "*json.RawMessage",

	"bigint":   "int64",
	"smallint": "int64",
	"integer":  "int64",

	"double precision": "float64",
	"real":             "float64",
	"numeric":          "float64",

	"date": "string",
	"timestamp with time zone": "string",
	"text": "string",

	"boolean": "bool",
}

const nl string = "\n"
const tab string = "\t"
const media_url string = "http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file"

var Log *os.File

const LogFileName string = "src/haii.or.th/api/thaiwater30/tools/api_service_docs/log.log"

func GenSwagger(outfile string) error {

	//
	//	log.Println(d, mapUseTable)

	f, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer f.Close()

	Log, err = os.Create(LogFileName)
	if err != nil {
		return err
	}
	defer Log.Close()
	f.WriteString(`package api_service

import (
	"encoding/json"
)
// @Document		v1.dataservice
// @Version			1.0
// @Title			WebService API for Data Services
// @Description    	WebService ในกลุ่มนี้ เป็น WebService ที่ใช้สำหรับ ให้บริการข้อมูล ที่ผู้ใช้ ร้องขอ ผ่าน ระบบ ให้บริการข้อมูลแก่บุคลภายนอก 
// @				
// @				เมื่อผู้ใช้ซึ่งเป็นบุคลภายนอก ร้องข้อข้อมูล ระบบ จะทำการจัดเตรียมข้อมูล จากนั้น ระบบ จะทำการสร้าง จดหมายอิเล็คโทรนิค เพื่อ
// @    			แจ้งผู้ใช้ถึง URL ที่ ผู้ใช้ สามารถใช้เพื่อเข้าถึงข้อมูลที่ร้องขอได้
// @				
// @				โดย WebService ที่ให้บริการข้อมูลเหล่านี้ ทุก service จะมี parameter พิเศษ หนึ่งตัวชื่อ eid ซึ่งจะถูกสร้าง จากระบบ
// @				โดยจะเป็น ค่าตัวอักษรที่ไม่ซ้ำกัน (unique RFC7515 Unpadded 'base64url') กับการร้องข้อข้อมูลอื่นก่อนหน้านี้ 
// @				เพื่อป้องกันไม่ให้คนอื่นซึ่งไม่ใช่ผู้ร้องขอข้อมูลเข้าถึงข้อมูลชุดนี้ได้
// @				
// @				สามารถดู eid ได้จาก จดหมายอิเล็คโทรนิค ที่อยู่ใน ช่องทางการรับข้อมูล
// @TermsOfService 	http://www.haii.or.th/tos
// @ContactEmail    api@haii.or.th
// @License      	http://www.haii.or.th/license HAII License
// @ExternalDoc		http://swagger.io/swagger-ui/ Find out more about Swagger-UI

// @DocumentName	v1.dataservice
// @Module			thaiwater30
// @Description		ระบบให้บริการข้อมูล Api Services

`)

	d, mapUseTable := getMd()
	genfile(f, d, mapUseTable)
	f.Sync()
	log.Println("gen file success")
	Log.WriteString(" =================================== \n")
	Log.WriteString("gen file SUCCESS ")
	Log.Sync()
	return nil
}

// generate file
func genfile(f *os.File, d []*md, mapUseTable map[string]bool) {
	mapTableDesc := findTableComment(mapUseTable)
	noData := make([]int64, 0)
	var service = "thaiwater30/api_service?mid="
	for _, v := range d {
		var id string = fmt.Sprintf("%03d", v.Id)

		s := `
// @DocumentName	v1.dataservice
// @Service 		` + service + id + `
// @Summary 		` + v.Name + `
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
`
		s += `// @Response		200 Metadata_` + id + ` successful operation
type	Metadata_` + id + ` struct{
`
		//		st, err, _noData := FindSturct(v, mapTableDesc)
		//		if err == nil {
		//			s += st
		//		}
		//		if _noData != 0 {
		//			noData = append(noData, _noData)
		//		}

		if model_detail.IsMedia(v.TableName) {
			// เป็น media ใช้ struct Metadata_Media
			//			s += `// @Response		200 Metadata_Media successful operation
			//			`
			st, err, _noData := FindSturct(v, mapTableDesc)
			if err == nil {
				s += st
			}
			if _noData != 0 {
				noData = append(noData, _noData)
			}

		} else {
			//			s += `// @Response		200 Metadata_` + id + ` successful operation
			//type	Metadata_` + id + ` struct{
			//`
			for _, field := range v.Fields {
				d := getDT(mapTableDesc, v.TableName, field)
				ex := getExample(v.TableName, field)
				s += tab + strings.ToUpper(field) + tab + getType(d.Type) + tab + getJson(field) + tab + "// " + ex + d.Desc + nl
			}
			//			s += "}" + nl
		}
		s += "}" + nl
		f.WriteString(s)
	}
	log.Println(" NO data [", len(noData), "]  MID : ", noData)
	Log.WriteString(fmt.Sprintf(" NO data [%d] MID : %v \n", len(noData), noData))
}

// find struct  value from db (only media)
func FindSturct(arg *md, mapTableDesc map[string]map[string]*dt) (string, error, int64) {
	p, err := model_detail.GetMetadataByMetadata(arg.Id)
	if err != nil {
		log.Println(err, " from FindSturctMedia ", arg)
		Log.WriteString(fmt.Sprintf("%v from FindSturctMedia %v \n", err, arg))
		return "", err, 0
	}
	p.Service_id = 4

	if model_detail.IsMedia(arg.TableName) {
		err = model_detail.FindMediaTypeId(p)
		if err != nil {
			log.Println(err, " from FindMediaTypeId ", arg)
			Log.WriteString(fmt.Sprintf("%v from FindMediaTypeId %v \n", err, arg))
			return "", err, 0
		}

		if p.MediaTypeId == "" {
			log.Println(" no media_type_id ?", arg.Id)
			Log.WriteString(fmt.Sprintf("no media_type_id %v \n", arg.Id))
		}
	}

	itf := make([]interface{}, 0)
	p.Sql, itf = model_detail.SQL_GenSQLSelectDataservice_All(p)

	p.Sql += " LIMIT 1"
	log.Println("MID :", arg.Id, ", Query : ", p.Sql, itf)
	Log.WriteString(fmt.Sprintf("MID %v, Query : %v  %v \n", arg.Id, p.Sql, itf))

	tn := time.Now()
	db, err := pqx.Open()
	row, err := db.Query(p.Sql, itf...)
	d := time.Now().Sub(tn)
	log.Printf("(%0.3f ms)\n", d.Seconds()*1000)
	Log.WriteString(fmt.Sprintf("(%0.3f ms)\n", d.Seconds()*1000))

	if model_detail.IsMedia(arg.TableName) {
		return FindExampleMedia(row, arg, p)
	} else {
		return FindExample(row, arg, mapTableDesc)
	}

}

func FindExample(row *sql.Rows, arg *md, mapTableDesc map[string]map[string]*dt) (string, error, int64) {
	str := ""
	data, _ := model_detail.ScanData(row)
	if len(data.Data) < 1 {
		// ไม่มีข้อมูล ใช้ตัวอย่างข้อมูลจาก map
		log.Println(" no result, MID: ", arg)
		column := data.Columns
		for _, v := range column {
			d := getDT(mapTableDesc, arg.TableName, v)
			ex := getExample(arg.TableName, v)
			str += tab + strings.ToUpper(v) + tab + getType(d.Type) + tab + getJson(v) + tab + "// " + ex + d.Desc + nl
		}
		return str, nil, arg.Id
	} else {
		d := data.Data[0]
		for i, v := range d {
			d := getDT(mapTableDesc, arg.TableName, i)
			ex := example(converInterfaceToString(i, v))
			str += tab + strings.ToLower(i) + tab + getType(d.Type) + tab + getJson(i) + tab + "// " + ex + d.Desc + nl
		}
		return str, nil, 0
	}
}

// find struct and exmaple for media
func FindExampleMedia(row *sql.Rows, arg *md, p *model_detail.Strct_Data) (string, error, int64) {
	var dm *model_metadata.Struct_Data_Media
	var rid int64 = 0
	data_media, _ := model_metadata.ScanData_Media(row, media_url)
	if len(data_media) < 1 {
		log.Println(" no result, MID: ", arg)
		rid = arg.Id

		dm = &model_metadata.Struct_Data_Media{
			Agency_id:      p.Agency_id.String,
			Media_Type_id:  p.MediaTypeId,
			Media_Datetime: "2016-12-13T21:06:55+07:00",
			Path:           "custom_path",
			Filename:       arg.Name + ".pdf",
			Media_Desc:     arg.Name,
			Refer_Source:   " ",
		}
	} else {
		dm = data_media[0]
	}

	str := ""
	//	Agency_id      int64  `json:"agency_id"`      // example:`50` รหัสหน่วยงาน agency's serial number
	str += tab + strings.ToUpper("agency_id") + tab + "int64" + tab + getJson("agency_id") + tab + "// " + example(dm.Agency_id) + " รหัสหน่วยงาน agency's serial number" + nl

	//	Media_type_id  int64  `json:"media_type_id"`  // example:`141` รหัสแสดงชนิดข้อมูลสื่อ
	str += tab + strings.ToUpper("media_type_id") + tab + "int64" + tab + getJson("media_type_id") + tab + "// " + example(dm.Media_Type_id) + " รหัสแสดงชนิดข้อมูลสื่อ" + nl

	//	Media_datetime string `json:"media_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บข้อมูลสื่อ record date
	str += tab + strings.ToUpper("media_datetime") + tab + "string" + tab + getJson("media_datetime") + tab + "// " + example(dm.Media_Datetime) + " วันที่เก็บข้อมูลสื่อ record date" + nl

	//	Media_path     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=AAECAwQFBgcICQoLDA0ODz2sq-vcpj7lylQ7-UPJBMDrGmg4sd2EqTleGoMGzbfRwOOn9GdVm9blDTBH42TRFzCh4Sws-QyEtcntJfW62mIK0Q==` ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	str += tab + strings.ToUpper("media_path") + tab + "string" + tab + getJson("media_path") + tab + "// " + example(encryptUrl(dm.Path+"/"+dm.Filename.(string))) + " ที่อยู่ของไฟล์ข้อมูลสื่อ file path location" + nl

	//	Filename       string `json:"filename"`       // example:`00Latest.jpg` ชื่อไฟล์ของข้อมูลสื่อ file name
	str += tab + strings.ToUpper("filename") + tab + "string" + tab + getJson("filename") + tab + "// " + example(dm.Filename) + " ชื่อไฟล์ของข้อมูลสื่อ file name" + nl

	//	Media_desc     string `json:"media_desc"`     // example:`ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi` รายละเอียดของข้อมูลสื่อ description
	str += tab + strings.ToUpper("media_desc") + tab + "string" + tab + getJson("media_desc") + tab + "// " + example(dm.Media_Desc) + " รายละเอียดของข้อมูลสื่อ description" + nl

	//	Refer_source   string `json:"refer_source"`   // example:`http://weather.is.kochi-u.ac.jp/SE/00Latest.jpg` แหล่งข้อมูลสำหรับอ้างอิง reference source
	str += tab + strings.ToUpper("refer_source") + tab + "string" + tab + getJson("refer_source") + tab + "// " + example(dm.Refer_Source) + " แหล่งข้อมูลสำหรับอ้างอิง reference source" + nl
	return str, nil, rid
}

// convert datatype to swagger datattype
func getType(t string) string {
	if s, ok := M_DataType[t]; ok {
		return s
	}

	return "string"
}

//
func converInterfaceToString(col string, val interface{}) string {
	if val == nil {
		return ""
	}
	if float, ok := val.(float64); ok {
		// type float64
		return fmt.Sprintf("%f", float)
	} else if i, ok := val.(int64); ok {
		return fmt.Sprintf("%d", i)
	} else if t, ok := val.(time.Time); ok {
		// date
		var v string = t.Format(time.RFC3339)
		if len(col) > 4 {
			// column ยาวกว่า 4 ตัวอักษร
			if col[len(col)-4:] == "date" && col != "temperature_date" {
				// ต้องเช็คว่าเป็น date รึป่าว จะได้แปลงเป็น format 2006-01-02
				// ไม่เอา temperature_date เพราะในดาต้าเบส เก็บเป็น datetime
				// it is of type time.Time
				v = t.Format("2006-01-02")
			}
		}
		return v
	} else if js, ok := val.(map[string]interface{}); ok {
		// เป็น map[string]interface{}
		// วนลูป map แปลงเป็น string
		str := `{`
		for i, v := range js {
			if str != `{` {
				str += ", "
			}
			str += `"` + i + `" : ` + `"` + converInterfaceToString(col, v) + `"`
		}
		return str + `}`

	}

	return val.(string)
}

// get field description
func getDT(m map[string]map[string]*dt, tb, field string) *dt {
	if s, ok := m[tb][field]; ok {
		return s
	}
	log.Println("err no tb field", tb, field)
	return nil
}

// get json field
func getJson(s string) string {
	return "`json:\"" + s + "\"`"
}

// get field example
func getExample(tb, s string) string {
	ex, ok := MapFieldExample[tb][s]
	if !ok {
		// หาตาม table ไม่เจอ
		log.Println(tb, s, "not found in MapFieldExample")
		// หาจาก all
		ex, ok = MapFieldExample["all"][s]

		if !ok {
			log.Println(tb, s, "not found in MapFieldExample ALL")
			return ""
		}
		log.Println("found in all ", ex)

	}
	if ex == "" {
		log.Println(tb, s, " MapFieldExample no value")
		return ""
	}
	return example(ex)
}

//  interface to example:`interface`
func example(s interface{}) string {
	var str string
	if s != nil {
		if t, ok := s.(time.Time); ok {
			str = t.Format(time.RFC3339)
		} else if v, ok := s.(int64); ok {
			str = strconv.FormatInt(v, 10)
		} else if v, ok := s.(float64); ok {
			str = fmt.Sprintf("%f", v)
		} else {
			str = s.(string)
		}
	}

	return "example:`" + str + "` "
}

// encryptext
func encryptUrl(arg string) interface{} {
	chipper, _ := ZXC.GetCrypter()
	rs, _ := chipper.EncryptText(arg)
	rs = media_url + "?file=" + rs

	return rs
}

// หา table comment, datatype ทั้งหมด
func findTableComment(mapUseTable map[string]bool) map[string]map[string]*dt {
	rs := make(map[string]map[string]*dt)

	var sqlWhere string
	for table, _ := range mapUseTable {
		if sqlWhere != "" {
			sqlWhere += ", "
		}
		sqlWhere += "'" + table + "'"
	}

	db, err := pqx.Open()
	check(err)

	log.Println(SQL_SelectComment + " WHERE table_name IN (" + sqlWhere + ")")
	rows, err := db.Query(SQL_SelectComment + " WHERE table_name IN (" + sqlWhere + ")")
	check(err)
	for rows.Next() {
		var (
			_table_name  string
			_column_name string
			_desc        sql.NullString
			_data_type   sql.NullString
		)
		err = rows.Scan(&_table_name, &_column_name, &_desc, &_data_type)
		check(err)

		if _, ok := rs[_table_name]; !ok {
			rs[_table_name] = make(map[string]*dt)

			st := model_metadata.GetTable(_table_name)
			if st.HasProvince {
				// มี province
				rs[_table_name]["province_name"] = &dt{Desc: "ชื่อจังหวัดของประเทศไทย", Type: "json"}
				rs[_table_name]["amphoe_name"] = &dt{Desc: "ชื่ออำเภอของประเทศไทย", Type: "json"}
				rs[_table_name]["tumbon_name"] = &dt{Desc: "ชื่อตำบลของประเทศไทย", Type: "json"}
			}
		}
		rs[_table_name][_column_name] = &dt{Desc: _desc.String, Type: _data_type.String}
	}
	return rs
}

type dt struct {
	Desc string
	Type string
}

// 	get metadata to struct
func getMd() ([]*md, map[string]bool) {
	db, err := pqx.Open()
	check(err)

	rows, err := db.Query(SQL_SelectMetadata)
	check(err)

	var rs = make([]*md, 0)
	var rsMap = make(map[string]bool)
	for rows.Next() {
		var (
			_id         int64
			_name       sql.NullString
			_table_name sql.NullString
			_fields     sql.NullString
		)

		err = rows.Scan(&_id, &_name, &_table_name, &_fields)
		check(err)

		var _im_f map[string]interface{}
		// หา name
		var name string
		json.Unmarshal([]byte(_name.String), &_im_f)
		name = _im_f["th"].(string)

		// หา field
		var fields = make([]string, 0)
		json.Unmarshal([]byte(_fields.String), &_im_f)

		st := model_metadata.GetTable(_table_name.String)
		if st == nil {
			log.Println(_table_name.String)
		}
		if st.IsMaster {
			// เป็น master table เอา Id ด้วย
			fields = append(fields, "id")
		}

		for _, v := range _im_f["fields"].([]interface{}) {
			s := v.(string)
			if s == "" {
				log.Println(_id, _im_f["fields"])
				continue
			}
			if s[0:1] == "#" { // ไม่เอา ที่ขึ้นต้นด้วย #
				continue
			}
			if s == "geocode_id" && st.HasProvince {
				// มี geocode_id และ setting ไว้ว่า ตารางนี้ มีการ join table geocode
				fields = append(fields, "province_name")
				fields = append(fields, "amphoe_name")
				fields = append(fields, "tumbon_name")
			} else {
				fields = append(fields, s)
			}

		}

		o := &md{}

		o.Id = _id
		o.Name = name
		o.TableName = _table_name.String
		o.Fields = fields

		rs = append(rs, o)

		rsMap[_table_name.String] = true
	}
	return rs, rsMap
}

// check err != nil
func check(err error) {
	if err != nil {
		exit(0, err)
	}
}

//	exit programe and display err
func exit(code int, err error) {
	log.Println(err)
	os.Exit(code)
}

type md struct {
	Id        int64
	Name      string
	TableName string
	Fields    []string
	DataType  string
}
