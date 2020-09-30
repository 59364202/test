package api_service

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

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=001
// @Summary 		ระดับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_001 successful operation
type Metadata_001 struct {
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64          `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=003
// @Summary 		แผนที่แม่น้ำสำคัญในประเทศที่กรมเจ้าท่าดูแล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_003 successful operation
type Metadata_003 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`1`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`54`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-05-31T17:59:11+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=cjWZb8e1MBwEnYhLuD69myBFAu1FBwVEQppcIHy4sm_N-2_Q4HUFI10UYuLMFQeUZ1RImtD6noXVuwfbCgrvT5zXFs6e5STrWZwi9iO_aca3ZHpeOddcmL1TcdxFmkhbhhS9kUd81ChLi6RQsW0c5g`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`test_1.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`md-river_map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=004
// @Summary 		ข้อมูลท่าเทียบเรือ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_004 successful operation
type Metadata_004 struct {
	ID          int64            `json:"id"`          // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	AGENCY_NAME *json.RawMessage `json:"agency_name"` // ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=005
// @Summary 		ข้อมูลพื้นฐานของสถานีวัด ระดับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_005 successful operation
type Metadata_005 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=006
// @Summary 		ข้อมูลร่องน้ำหลังการขุดลอก 16 ร่องน้ำเศรษฐกิจ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_006 successful operation
type Metadata_006 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`1`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`53`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=-XQb1ZCH76C2zTHfoQ6OlKd9jWm0oTxB88OyI54HVisLK4r8EzV_fSHV31MssgWD_hGHgKIdedeCWaQyH43r9vXS3NqSTrT1hn7OZqyxFKNVRfY6zw0MCoth4mw8AR5LVZBrZQSRxWnBAqL91G3c9A`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`plot.log`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`md-water_channel`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=007
// @Summary 		แผนการขุดลอกและบำรุงรักษาร่องน้ำภายในประเทศ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_007 successful operation
type Metadata_007 struct {
	ID          int64            `json:"id"`          // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	AGENCY_NAME *json.RawMessage `json:"agency_name"` // ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=008
// @Summary 		แผนที่เสี่ยงภัยดินถล่ม (Landslide Hazard map) (ศูนย์เตือนภัย)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_008 successful operation
type Metadata_008 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`2`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`131`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T13:34:37+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=DOUymAaNNyF1H2YgMiENQW6wNHNDZGLmT0e2EMRODLCxcFkrZwXkai-G8EZPmvdTCZ4fcSeekLwIpllqY1VxLLnsL939BgFlTCbEGeo6Oh65Qq-ifw3OMxkxac6_5jQCb-xlLpq_9s0FZZ1E6CyeKrZbgxzAYp7fr6BjBqqE1fc`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`arc0653r.001`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dmr-landslide_hazard_map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=009
// @Summary 		หมู่บ้านเสี่ยงภัยดินถล่ม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_009 successful operation
type Metadata_009 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`2`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`56`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T13:34:57+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=n902GCQHYGMZEo2FynKgiFwUlRpulBweamJE2fGGKyEeMx7Iz4Cq0g9_z6r_GOmdUd8_t6wM4ntub5Dp8GOV7Hb5KHtLgXIn59wRsoKl3bQy_OpaDHMP8AmJJh-sMqoMZAGtc1CUgeCChUr6BWE1dqIxRic0fTGxQ_d4VXHLLik`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`vrisk_yala.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dmr-landslide-village-risk`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=010
// @Summary 		พื้นที่เสี่ยงภัยระดับชุมชน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_010 successful operation
type Metadata_010 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`2`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`127`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T13:35:35+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=dl3Jx9_FdYgjRwh0CTOiOjdKjKFLMe038g9sLofZIUeIBozcy7BoKrRngF9XT7W65wMc4yOqTHOlApWpJ7j7GK-tQJl6ewW_jucgNC0DMtwF4oL5WjQWyjp0Hj38dpIrWwzESrAS6P4k7MogQk1h4QJBnTdeMKtsfA_cO6dxs3A`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`scar_nammun_uttaradit.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dmr-hazard_community_area`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=011
// @Summary 		ข้อมูลจุดปลอดภัยดินถล่ม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_011 successful operation
type Metadata_011 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`2`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`128`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T13:35:49+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=UrZXZJsPErLrDrW3pV9O6w380tRb2SSQ6ISalbmngMXkPbN6Z3TgB_yD2Tc2spfrlZA_t8ZvzCpZ7O_S5TgFHRHBR2tHKLJqPjNa6QUcSondFIeSAwaLL-7UIorhpyxAqqhwcUfJ1pt2gMNuJNunPd_pvAPmHVizpsxz5RaEDJTiv8nioNdjccpiMOmnVCky`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`safety_area_utharadit1.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dmr-safe_landsilde_area`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=016
// @Summary 		สถานการณ์ธรณีพิบัติภัย
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_016 successful operation
type Metadata_016 struct {
	PROVINCE_NAME         *json.RawMessage `json:"province_name"`         // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME           *json.RawMessage `json:"amphoe_name"`           // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME           *json.RawMessage `json:"tumbon_name"`           // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	AGENCY_ID             int64            `json:"agency_id"`             // example:`2` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	GEOHAZARD_DATETIME    string           `json:"geohazard_datetime"`    // example:`2006-01-02T15:04:05Z07:00` วันและเวลาที่ประกาศสถานการณ์พิบัติภัย DateTime of geohazard
	GEOHAZARD_NAME        string           `json:"geohazard_name"`        // example:`บ.ใหม่ ม.6 ต.สาริกา อ.เมือง จ.นครนายก_เช้านี้ท้องฟ้าโปร่ง อากาศแจ่มใส เมื่อวานนี้มีฝนตกเล็กน้อย` ชื่อสถานการณ์พิบัติภัย name of geohazard
	GEOHAZARD_LINK        string           `json:"geohazard_link"`        // example:`http://www.dmr.go.th/ewt_news.php?nid=102885` ลิ้งที่แสดงสถานการณ์พิบัติภัย geohazard link
	GEOHAZARD_DESCRIPTION string           `json:"geohazard_description"` // example:`null` รายละเอียดสถานการณ์พิบัติภัย description of geohazard
	GEOHAZARD_AUTHOR      string           `json:"geohazard_author"`      // example:`http://www.dmr.go.th/` ผู้รายงานสถานการณ์ author
	GEOHAZARD_COLORLEVEL  string           `json:"geohazard_colorlevel"`  // example:`null` ระดับสีของเกณฑ์พิบัติภัย color level of geohazard
	GEOHAZARD_REMARK      string           `json:"geohazard_remark"`      // example:`null` หมายเหตุ Remark
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=017
// @Summary 		ประกาศทรัพยากรธรณี
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_017 successful operation
type Metadata_017 struct {
	AGENCY_ID         int64  `json:"agency_id"`         // example:`2` รหัสหน่วยงาน agency number
	FLOOD_DATETIME    string `json:"flood_datetime"`    // example:`2006-01-02T15:04:05Z07:00` วันและเวลาที่ประกาศสถานการณ์น้ำ DateTime of flood
	FLOOD_NAME        string `json:"flood_name"`        // example:`รายงานสถานการณ์ธรณีพิบัติภัยประจำวัน วันอังคารที่ ๒๒ สิงหาคม พ.ศ. ๒๕๖๐` ชื่อสถานการณ์น้ำ name of flood
	FLOOD_LINK        string `json:"flood_link"`        // example:`http://www.dmr.go.th/ewt_news.php?nid=102862` ลิ้งที่แสดงสถานการณ์น้ำ flood link
	FLOOD_DESCRIPTION string `json:"flood_description"` // example:` ` รายละเอียดสถานการณ์น้ำ description of flood
	FLOOD_AUTHOR      string `json:"flood_author"`      // example:`http://www.dmr.go.th/` ผู้รายงานสถานการณ์ author
	FLOOD_COLORLEVEL  string `json:"flood_colorlevel"`  // example:` ` ระดับสีเกณฑืเตือนภัย
	FLOOD_REMARK      string `json:"flood_remark"`      // example:` ` หมายเหตุ
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=018
// @Summary 		ข้อมูลฝน จากโทรมาตรเตือนภัย (early warning)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_018 successful operation
type Metadata_018 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string           `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL15M       float64          `json:"rainfall15m"`       // example:`0` ปริมาณน้ำฝนทุก 15 นาที Rainfall Every 15 minute
	RAINFALL12H       float64          `json:"rainfall12h"`       // example:`4.5` ปริมาณน้ำฝนทุก 12 ชั่วโมง Rainfall Every 12  hours
	RAINFALL24H       float64          `json:"rainfall24h"`       // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=019
// @Summary 		ข้อมูลระดับน้ำจากโทรมาตรเตือนภัย (early warning)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_019 successful operation
type Metadata_019 struct {
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64          `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=020
// @Summary 		ข้อมูลความชื้นในดิน จากโทรมาตรเตือนภัย (early warning)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_020 successful operation
type Metadata_020 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`1969` รหัสสถานีโทรมาตร tele station's serial number
	SOIL_DATETIME   string           `json:"soil_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่ของค่าข้อมูลความชื้นในดิน record date
	SOIL_VALUE      float64          `json:"soil_value"`      // example:`61.6` ค่าข้อมูลความชื้นในดิน (มม.) soil moisture value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=023
// @Summary 		ขอบเขตลุ่มน้ำย่อยมาตรฐาน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_023 successful operation
type Metadata_023 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`3`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`85`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=vk6hWKtPPWLLbUri4h5Muzif_1VgZVwP7-ZLF_rgMrc5IIez7aU6b1bxoP4MAfx4IkSZdl5uj1G24yX49OTo4r0BFXFoP9ujdCpXo3rAGvjeda4uR2BuSpLd0H70uIV5AVI46s18lasM3ZZ6eqLY-m9zFkjF2AOpk_3nJFlK1MI5fSmqP8LSUZB1osaGvyZO`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ขอบเขตลุ่มน้ำย่อยมาตรฐาน.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ขอบเขตลุ่มน้ำย่อยมาตรฐาน`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=025
// @Summary 		รายงานการศึกษา/ข้อมูลพื้นฐานลุ่มน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_025 successful operation
type Metadata_025 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`3`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`85`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=TLi3ReNJECoO87ai5Uzr9E5RA20WwhhY5uIbQz2PCpVF92bzJrCMHgQa0babQ7lccEMd2E-sqUTb8nJvF-q5WKK6LYEDTKwU5fdnpu7BupoCyEVn9tS9rwZvOy6RtXW1K1bRGuCn_f9W6jOWASpRQk4sUDHD4qHL-adhe-f5qnRFXOW-hHR6DWAyqp26JQ693mRg2eyIhjUIf6t5FSqyYtkahyC217235Fm_8fXbB2s`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`รายงานการศึกษา/ข้อมูลพื้นฐานลุ่มน้ำ.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`รายงานการศึกษา/ข้อมูลพื้นฐานลุ่มน้ำ`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=026
// @Summary 		แผนที่ข้อมูลพื้นฐานอื่นๆ ของลุ่มน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_026 successful operation
type Metadata_026 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`3`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`85`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=LHefUV1-6h1BOZg6D-llU73G7zIXBw4IoUPuL4xR72ua3R7lwfYqKHVuf_DHVmXQD_g7XorBl6v2iPAKQTPMOKcwD8-Y0YOCQHBq6qR3oD1kLuPeLIm0Yxv0QHkrrY7MWblE-Nbs8N9RcTeJBBVPUgmO_XQsLiA8QjYchLNj_rFrtklT4NT41fMS_rAamPeZqK8SofMI0yQ7hZFknmU6al3Y46GDqRnR1fvVUsxTwhQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่ข้อมูลพื้นฐานอื่นๆ ของลุ่มน้ำ.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่ข้อมูลพื้นฐานอื่นๆ ของลุ่มน้ำ`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=027
// @Summary 		ข้อมูลพื้นฐานของโทรมาตรเตือนภัย
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_027 successful operation
type Metadata_027 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=032
// @Summary 		ตำแหน่งบ่อน้ำบาดาล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_032 successful operation
type Metadata_032 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`4`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`107`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-01-18T10:48:05+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=OuS_AKyAtFYHr3WqmDkDTZWejP-nySfLjjnlA1O48Rq3gAVekSWEG4B3Y_cDuC_F1wqjj1aYX9SV8zXsVdJmtdaICKRW9EA5YNNnJ0nP8VtOLyy-QIyWCJFcmI_F-z1jEKsJnMZbhLSvj2mvtBiyLzNPW_y9wEF_OKh915C5O50LLCW7oKVTklP3FWc4yl5l`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`Well_Lower_NE_Z48.dbf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dgr-groundwater_location`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=033
// @Summary 		ข้อมูลบ่อน้ำบาดาล-คุณภาพ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_033 successful operation
type Metadata_033 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`4`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`130`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-09-29T14:22:39+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=l24FNF5Vz5hWd9D93fq5RVkOtWcLYlQ7MpOoTUZFLMLNivZK_zNBu8unpTrq3CfJE9pwNzHHZoHIGxjIZAlbmm5qYlmPT1FhubPSQ3C-u3lXmxxhqby24voYF44xN4OkW2LECULwA9N53Uxh19svfuyxa0FRn5SyX7uiDG5HHGgToLTYjZUKZz4Tj0X_9yLZKBTmwrbizl3Hm5MHpLGlp59tt2F0mPAJegd3QuAZwiFMNAUFbt7TnnkNgHw4DGHj0qB9O9i-MduB9V-W3YdZTWXLVJzrAOrvTW0no5GoKk8PCAnEECdec4VSxgMv3NRB`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลบ่อน้ำบาดาล ตำแหน่ง ระดับ คุณภาพ details_point_well.xls`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`dgr-groundwater_data`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=034
// @Summary 		ข้อมูลคุณภาพน้ำบาดาล (เกณฑ์น้ำอุปโภคบริโภค)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_034 successful operation
type Metadata_034 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`4`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`108`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=H4I4DBgr5G61Wv0K1eaPYvHabmnYVXd65dsah19uO1MOV8vFGY_HH8Qxy22OkYMVuqYJepRrerRcXU5qMz5-P4T3iHM3CXZCyEHLmmyPN4J1zBVFqUtW57edf04Ki9mK0MZZFqZyjEVkTGgyc1HkZsGmJ3ThLyxnUXuvGrdmcuE_JJlkQ2Czx108N8w8NgNNETm3HZm8ecqrK8B7OY7v6AX5NukCP4PCWhxx27uZv3EQQYt3ZSN8j8PO4ws0d59e`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลคุณภาพน้ำบาดาล (เกณฑ์น้ำอุปโภคบริโภค).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลคุณภาพน้ำบาดาล (เกณฑ์น้ำอุปโภคบริโภค)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=036
// @Summary 		ข้อมูลแหล่งน้ำขนาดเล็กที่อยู่ในความรับผิดชอบ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_036 successful operation
type Metadata_036 struct {
	AGENCY_ID              int64   `json:"agency_id"`              // example:`5` รหัสหน่วยงาน agency's serial number
	WATER_RESOURCE_OLDCODE string  `json:"water_resource_oldcode"` // example:`5309` รหัสแหล่งน้ำชุมชน water resource's code
	PROJECTNAME            string  `json:"projectname"`            // example:`งานปรับปรุงพื้นที่และจัดทำระบบส่งน้ำในไร่นา ( 70 แห่ง ) ปี 2547` ชื่อโครงการ project name
	PROJECTTYPE            string  `json:"projecttype"`            // example:`คลองส่งน้ำ` กิจกรรม project type
	FISCAL_YEAR            string  `json:"fiscal_year"`            // example:`2547` ปีงบประมาณ fiscal year
	MOOBAN                 string  `json:"mooban"`                 // example:`เมืองแปง ม.1` หมู่บ้าน address number
	COORDINATION           string  `json:"coordination"`           // example:`MB338244` ตำแหน่งโครงการ พิกัด x-y coordination
	BENEFIT_HOUSEHOLD      int64   `json:"benefit_household"`      // example:`250` ครัวเรือนที่ได้รับประโยชน์ (ครอบครัว) benefit household
	BENEFIT_AREA           int64   `json:"benefit_area"`           // example:`1100` พื้นที่ที่ได้รับประโยชน์ (ไร่) benefit area
	CAPACITY               float64 `json:"capacity"`               // example:`3700` ความจุ (ลูกบาศก์เมตร) capacity
	STANDARD_COST          float64 `json:"standard_cost"`          // example:`8326000` งบประมาณ (บาท) standard cost
	BUDGET                 float64 `json:"budget"`                 // example:`0` ค่าใช้จ่าย (บาท) budget
	CONTRACT_SIGNDATE      string  `json:"contract_signdate"`      // example:`2004-03-01` วันที่ลงนามในสัญญา contract signdate
	CONTRACT_ENDDATE       string  `json:"contract_enddate"`       // example:`2004-07-29` วันที่สิ้นสุดสัญญา contract enddate
	REC_DATE               string  `json:"rec_date"`               // example:` ` วันที่บันทึกข้อมูล record date
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=037
// @Summary 		การใช้ที่ดินในแต่ละภาคส่วน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_037 successful operation
type Metadata_037 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`87`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-09-27T16:07:01+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=PedrsCF85Tz53pO4JCUqNST9VpreWqM7y2oEjG1o8PnCkerfphvBCRn-iA13qU2YjrN7zU8WuzU79tPYxo-Q42RSiKKDcI92UX4sLPQVwHMbl6CYv4gQsh3CCUMJNIc9HhXhrl-1fovYe5tgfEj9ApsoR1jIXb3ZZ8IgPovXI6lX_AhJsitLBHSA4uWboJ36IvooZ6ds8V8X79k2-jc9ug`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ใบรายชื่อ  21-25 สค.xls`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ldd-soil_usage`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=038
// @Summary 		สภาพภัยแล้งจากภาพถ่ายผ่านดาวเทียม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_038 successful operation
type Metadata_038 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`88`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=gkZku_MrBHOg6t6Oe-J8-uSFR4z9pqfDDtZnxWAOnbxdWda_8OW6umQAKdojonV9wCVHIkZ5aEAqrMGQGqfQQbIQgoDpzrXSlm1UptpQh-jsQGzGZhxxN9dG3TG04xJF2tlgyvEIDqyCtR4qbYnXyw13Sh_xw5bf7CNBBhOO0QYpia0PqMAZL62PwJ3lKtbHNioZFdU1t_rihkyGxUNq7aKSQFV82vCuL8zHcD07ALY`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`สภาพภัยแล้งจากภาพถ่ายผ่านดาวเทียม.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`สภาพภัยแล้งจากภาพถ่ายผ่านดาวเทียม`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=039
// @Summary 		หมู่บ้านที่เสี่ยงต่อดินถล่มและน้ำป่าไหลหลาก
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_039 successful operation
type Metadata_039 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`89`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=WMh9q2J6uq-iQ0ZPsq47AQHXmRX2dejR1Xh0dmXPSflEhY89GtmOxvnsemcagsrvgxzk6rQiiBCn0BwmBvznDnkPGvK3xhSpQNFq0u95dj9UilG0BfGxwlKSmUc7vsHULNHPnW86MRlZ4-NpT0L_ikCPmWlAIXYrlorSuusAsybFxB68xVtxtXsMiYHD6av75-EWLSsh38X0b4QLzq_QeLodX9NTzPIiAhnnGxi4Qv-A2DamdTVNDuI7Oq5ETddQXvtFkEdwccCUUUjpF59l1A`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`หมู่บ้านที่เสี่ยงต่อดินถล่มและน้ำป่าไหลหลาก.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`หมู่บ้านที่เสี่ยงต่อดินถล่มและน้ำป่าไหลหลาก`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=040
// @Summary 		พื้นที่น้ำท่วมซ้ำซาก
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_040 successful operation
type Metadata_040 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`90`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-09-28T16:12:26+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=qhLtB25CiOaV0yhrIsfwsfpgdyW3pGJ-XgDHf8o2nEyomP2fgPq0MojiJ-n5rFnZOmz0wXoN7CQbBrPd3s7QKNuIZV5JhaRJ4NwNXTZfmovxHrr6nTs6wjoAi-dkg7eg9BLYxW5b8CdetCObD6oEThkjmERYei3YeHTGMqWLUkQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`fl_th56.shp`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ldd-high_flooding_area`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=041
// @Summary 		พื้นที่เสี่ยงต่อการเกิดดินถล่ม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_041 successful operation
type Metadata_041 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`91`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=ZJZbyIdS1fSwr7ejAvOjbIM2Ep1ri8ieL1OvxM11fYyXmPVEi0b6jgYcf1aQyaNPDZtkocAStzpzINeqATV37iJjDJcq-ZND1R2uDPulHwbfckQ6cgOgSYhw5jH1hVI2se_XVvTkIxNgXIkB9LyM74YYXaJuXTzf6kGujYFV1R-3dOX1hXuM0kP8cSrIxXB3XW26aKDzFRBOS7p02I8QJQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`พื้นที่เสี่ยงต่อการเกิดดินถล่ม.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`พื้นที่เสี่ยงต่อการเกิดดินถล่ม`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=042
// @Summary 		พื้นที่เสี่ยงวิกฤตอุทกภัย
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_042 successful operation
type Metadata_042 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`92`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=iMIwNeE9m5cdAEdRTvdvEplWfm3G3OjS8anlNIygak5LFiKdrHe8VgJiCxd1o7mR9cassEVD2OodCjiv4CNL31Qw05kvrJIke9eczzVvbk1nfMnGa-oAmcp28hmK2NzRYloSDDbHrfrvKZXYBaWCHbo_cExH_U3fR4aUtPSAmHNyP2ip0vfdGWm64GJJDEKM`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`พื้นที่เสี่ยงวิกฤตอุทกภัย.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`พื้นที่เสี่ยงวิกฤตอุทกภัย`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=044
// @Summary 		แผนที่ชุดดิน (soil series) มาตราส่วน 1:100,000
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_044 successful operation
type Metadata_044 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`93`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=gguYDIpKGEpVqjyErYT1Vil6aQHZTrE32_o1HxtYfIExl0DpPrRiu1V7eLUC7n2aVCJH44EPVvcEjsRFyukH2bdajGe4S6ss65fpKUFfLgyqyGRCaz9FQwr-D7dRuY1Ift6W-sHcA4-22fO8mz8O_n3XOkf8_Y4O83l7qLZHzG5bTxgmLQEAm3dvRsEbA5IEym0HKO_TO5VxeDHg3XdToQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่ชุดดิน (soil series) มาตราส่วน 1:100,000.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่ชุดดิน (soil series) มาตราส่วน 1:100,000`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=045
// @Summary 		แผนที่ดินเค็ม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_045 successful operation
type Metadata_045 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`94`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=qP6rO82kA3n4lWU_iRuyYmYHWSCN746NfSpQH8DMuJUUShIUQILZkrI329Y7Zm9Yicv6Byc-GkKPrIz0RVv-Ww1Dol8i7xMgs5Dv4gGQBjDf15tFLQZYxywxOzb_JAJEn3dn4gePUKD-580sH_eQSA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่ดินเค็ม.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่ดินเค็ม`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=046
// @Summary 		แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน 1:50,000
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_046 successful operation
type Metadata_046 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`95`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=dGP8czmPq_m63aUzlGsTxllHJTYkukTsD2KYLcT9e2Ed7I0jG6IGfggwk2rPquyX-b4AvXicPb73VrSGxyO5Ju05Eftcryrnih5K3rW6_x8fzuzEMP06SFRPHQVh-LZ2a2EEwkPlRDyOtu1PaZimshMoIIAT7YhbkoOMpCuhCznx2fabVksiM8FosjZ_15flo-plC4dcdrEmB3zN26vW7vZVA8Y6uW8Ei8rpF836i44`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน 1:50,000 .pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน 1:50,000 `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=047
// @Summary 		แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน  1:25,000
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_047 successful operation
type Metadata_047 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`96`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=XrU6j87_CF2zrq3suLizCZM4Zx-ladnn1IYpUoW8bHLpXooSFSHfN4ecpP8lxTH-80rhZZ78iKNaxDjWIbc1Y6noVBDgnLWBTQSe8MMmUh-Fuj8TezZ6_C6ekyfAr8SsVk89VFH3uw4239Ymn8VLl4s4p6tetgxxRwM5fKQkWUJMNPP5e64pP_6uWC2wLDl_CfYGb3ilVCJiVw_T7ZNN7a3xk7SStTwFoVySK8YZwtU`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน  1:25,000.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่กลุ่มชุดดิน (soil group) มาตราส่วน  1:25,000`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=048
// @Summary 		รายงานชุดดินจัดตั้ง
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_048 successful operation
type Metadata_048 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`97`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=AuLuMYf11Wo8V7bPYSuwwVmycQWrj-qjihHlYpkkdOzbf7lhgfcXplpA9Gw0XqEacmz7mMalNNPVsKMR-3iXtNh01kv-gEvfYgLuxzctqeA3lVv9J3riXo7-V1KDQyWGThlxF_cy0Pl8NnadB5A5nxmcdgXWrgLpJl2_IJEXbrE`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`รายงานชุดดินจัดตั้ง.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`รายงานชุดดินจัดตั้ง`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=049
// @Summary 		แผนที่การชะล้างพังทลายของดิน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_049 successful operation
type Metadata_049 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`98`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=MUJ6CtsM7JOpYW5KJJ9gyihvW-RWrU41aB590yydSrF062T-zqH8r7GRT7szYBnv-sUW_Xb1I7qeAjbLdDse5xEd9LYofxlaKIdFpJJn64a98GaY_cVEy-8xtRcHwiPu3I3yFKOhex8Z_CVbUSHV5eitlSw8yLqJafWgiwY0_qjN8SthEbujtpIHTxfwJR5Z8mfy4i_ulWSbYTBfvjwtYw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่การชะล้างพังทลายของดิน.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่การชะล้างพังทลายของดิน`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=050
// @Summary 		เขตการใช้ที่ดินระดับตำบล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_050 successful operation
type Metadata_050 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`99`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=m2rZaswEEo4zluXpcnh_LWUvaPOm1dVbF-Fregf0iUy_J2KKS79_8Bq8yQPDZTUJ41FZzP7ouEk_AVsJa01yHVpPkWnrAkyfQVCerN6gvAt6jFhY9pSb4E3hRrRZHFVMAL2Fpl0KuqHFvqcL75aq_TY_Oc9HmL1Xv_DK-cxzHytPmNKAXRESBZ_7Ts10cqUg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`เขตการใช้ที่ดินระดับตำบล.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`เขตการใช้ที่ดินระดับตำบล`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=051
// @Summary 		แผนที่พื้นที่ชุ่มน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_051 successful operation
type Metadata_051 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`100`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-09-28T17:06:48+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=CeB_bF_my12P8wjQTK8cehk6XqpE7wUivJluP8vf7pQ6OYPmvE1aYV0QP3H_FD_yj7gL0of6pz2Q08BYFxdIht0hviuPTC1ya48hNsNT3fnHN3QdcM51SWPlDgHPzE13VGiISZOReJHAbS7TItR4rdk-7LBXl6ndcA2FPKMx5xC5ZOiKPDBQ55MogCNzaK81JYUpjXmVm2QtrxVXtyMFZg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`Point_อินทร์บุรี_Identity.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ldd-wetland_map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=052
// @Summary 		แผนการใช้ที่ดินลุ่มน้ำระดับสาขา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_052 successful operation
type Metadata_052 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`101`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=o0An2pXTq-VO8sIN2s4L9YA2Z8t5Gi2nocO41ER0ejX54DYjBNGAvLy_cfpg9umqRTQIUXJQOU1mppCIQNqVmxSQNgDKGJuKJEYNnn9q3tNUuNg5V6FiddvBubrQZERVGPIr1AswAWWrw04rTB4sngrGXwS_SRX10uH57h_wop-g4QbGSHRBRGD0NH1n6M2J6ZNiFcwz5bTUS4rCjTEXhA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนการใช้ที่ดินลุ่มน้ำระดับสาขา.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนการใช้ที่ดินลุ่มน้ำระดับสาขา`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=053
// @Summary 		เขตการใช้ที่ดินพืชเศรษฐกิจ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_053 successful operation
type Metadata_053 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`102`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=mjmhFcyuoYnNlHZNpUYctq35hRJYrU_l8h-Pd9IQnBdwuwSgz6p3_wmm3x-CNgpN3xLlgzYIvnS1LQlf6gqc1Yzgl9QY_lVZ1ef2Qv2PEWRy_7nwennnm8Aeurq1gS_0wf0uROsyYAwHB89gGvUZZRzPEz4LVc4cGyzhVcGjnx2TOH5OhjPNO8DFUrJ4qIHS`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`เขตการใช้ที่ดินพืชเศรษฐกิจ .pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`เขตการใช้ที่ดินพืชเศรษฐกิจ `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=054
// @Summary 		ข้อมูลเขตเหมาะสมสำหรับการปลูกพืชเศรษฐกิจ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_054 successful operation
type Metadata_054 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`5`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`103`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=TbLX80x6o0-pxmOMg5sN8hUAW6PYT9-JXZ4QLf65jNfP1_Cb-5YuD52weu0FZQWh8d0IZdHWvP4yEgNJpviSrJ15GYe2mHrJPUn53IQyOiPykTSLoFTFu3XO3ZSSmkANnBVyInGrDpt5P1UAFeT04q-IlbMzlQcQl6iivcXzNes82eSN586ee4mrk5O32HHyn_Mys72peAAhCCY7oJ9TWbsB8YD9A7cFdgved3nh6Ie39nWKZTfHtKm5hzX-GdPl`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลเขตเหมาะสมสำหรับการปลูกพืชเศรษฐกิจ.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลเขตเหมาะสมสำหรับการปลูกพืชเศรษฐกิจ`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=055
// @Summary 		ระดับความสูงน้ำทะเล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_055 successful operation
type Metadata_055 struct {
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_M        float64          `json:"waterlevel_m"`        // ระดับน้ำ เมตร รสม.
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=056
// @Summary 		แผนที่อากาศผิวพื้น
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_056 successful operation
type Metadata_056 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`6`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`27`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T12:41:35+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=XyTPeozBxNFR1G84FGxGq7jBWFdgD4mOCfraXiPZ6OvCdgy_ft3OB02jNPPiuzLHRxcpWVHMdO-Y-pKBDNwFiqdQtWC-Lm9gLgr3IGV_zShT1bj0EtdkrzZYYgzx5-k-H6gDtZPvaXMLLl3uY6kiCldt3cJem82Ip_DaHUcfzw5o_MyU48O8b76X08a6CNSn`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`20180817054135_2018081618.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`hd-weather_map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=057
// @Summary 		ข้อมูลการพยากรณ์คลื่นเชิงตัวเลขพื้นที่เอเชียตะวันออกเฉียงใต้
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_057 successful operation
type Metadata_057 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`6`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`22`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T10:12:59+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=7w-7HXfzXn6bdR4PpdKm9nqNp2mcf9_6MdBE_w-ovJPKrwpPTSL9nAh_LytjZV5S1SdWHz0wTLTZeH8muWcRhc4WClC0RhDvviVtadxruO1dmGywv_XIHAdqxVsXM8Q0ZN0fu2Mk95dza7WTEfc2B9Ua-kBXMfBOOr2lWLZ_2XU`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`thai048.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`rtn map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=058
// @Summary 		ข้อมูลการพยากรณ์คลื่นเชิงตัวเลขพื้นที่อ่าวไทยและทะเลอันดามัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_058 successful operation
type Metadata_058 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`6`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`22`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T10:12:59+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=GsHx_mQ50XCv4XGS3kEEWkgCUqtnbrK7l_wUoaQG_maKslKfdpH4jhTuqNhWIEK6iyT1P0ZMjlB264Z5OExmWVT5yURFtqOw2GlR2bh_6u_2CnW-jJPKW-TJyS_iGEvCUbXfO_1Pg4wYQ2bApmzGbr3qSlJ-XlbJ75wSbiIAZSY`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`thai048.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`rtn map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=059
// @Summary 		ข้อมูลการพยากรณ์คลื่นเชิงตัวเลขพื้นที่มหาสมุทรอินเดีย
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_059 successful operation
type Metadata_059 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`6`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`22`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T10:12:59+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=qNValrtwcAheYAl9lPbXQ_mkUHUnv7F66st-GsCkhFo6ukezVKBPgIpspJApnYZCLshom0bMFDA9DjBbLeYVL8r0rs5Gs5xRTrEruabWxNrsGS2rsIGugrIdintqmxzNp4wfU9eBIyteyyuJnqaefq1034pWkHWL9R_DDoH7idE`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`thai048.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`rtn map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=060
// @Summary 		ข้อมูลการพยากรณ์การเปลี่ยนแปลงระดับน้ำทะเล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_060 successful operation
type Metadata_060 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`6`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`22`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T10:12:59+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=QdsVD6OkA2LEQsJCFZJEBDBrZj0dBZnTFZIDkeIuNOXhxYrPvT_CR4AaTcpLqbF_A5jRkNZ8Is5OxInsjbxVfgHj1MNjNFkzJG-aboDVzYLYN33OkebMhMuXxha9VU6qkLEogFeOqb30xj__p87-PpCQ8QMWsH8fiou-eAqTU0k`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`thai048.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`rtn map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=061
// @Summary 		ฝน จากสถานีตรวจสารประกอบอุตุนิยมวิทยา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_061 successful operation
type Metadata_061 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string           `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL5M        float64          `json:"rainfall5m"`        // example:`0` ปริมาณน้ำฝนทุก 5 นาที Rainfall Every 5 minute
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=062
// @Summary 		อุณหภูมิ จากสถานีตรวจสารประกอบอุตุนิยมวิทยา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_062 successful operation
type Metadata_062 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	TEMP_DATETIME   string           `json:"temp_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าอุณหภูมิ record date
	TEMP_VALUE      float64          `json:"temp_value"`      // example:`26.46` ค่าอุณหภูมิ temperature value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=063
// @Summary 		ลม จากสถานีตรวจสารประกอบอุตุนิยมวิทยา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_063 successful operation
type Metadata_063 struct {
	TELE_STATION_ID int64   `json:"tele_station_id"` // example:`60` รหัสสถานีโทรมาตร tele station's serial number
	WIND_DATETIME   string  `json:"wind_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความเร็วลม record date
	WIND_DIR_VALUE  float64 `json:"wind_dir_value"`  // example:`228` ค่าองศาของทิศทางลม wind direction value (degree)
	WIND_SPEED      float64 `json:"wind_speed"`      // example:`8` ค่าความเร็วลม wind value
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=064
// @Summary 		ความชื้น จากสถานีตรวจสารประกอบอุตุนิยมวิทยา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_064 successful operation
type Metadata_064 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	HUMID_DATETIME  string           `json:"humid_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความชื้นสัมพัทธ์ record date
	HUMID_VALUE     float64          `json:"humid_value"`     // example:`95.33` ค่าความชื้นสัมพัทธ์ humid value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=065
// @Summary 		ความกดอากาศ จากสถานีตรวจสารประกอบอุตุนิยมวิทยา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_065 successful operation
type Metadata_065 struct {
	TELE_STATION_ID   int64   `json:"tele_station_id"`   // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	PRESSURE_DATETIME string  `json:"pressure_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความกดอากาศ record date
	PRESSURE_VALUE    float64 `json:"pressure_value"`    // example:`1012.67` ค่าความกดอากาศ pressure value
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=067
// @Summary 		ข้อมูลขอบเขตพื้นที่ให้บริการ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_067 successful operation
type Metadata_067 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`7`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`115`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=l8t67T2daam0005taP0DTGyz4YRU8T8S5o0LJUCYD5IYfqiVv6t80GA70yhUDh0eSJJ8ghKeqanTFi9WFLWAVvYDnox9b8D5GoY4mpQkn9fISPrYIpOb-r-aceNECAytobZK7Ge5kPz9838FmyP3oi6IxWRCYAVGu85_H08-N_d3FhMsgkB2lwyh4XXTGqPI6KlHK6lrqEE1BWxlKJ0wyQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลขอบเขตพื้นที่ให้บริการ.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลขอบเขตพื้นที่ให้บริการ`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=068
// @Summary 		ข้อมูลพื้นฐานสถานีประปา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_068 successful operation
type Metadata_068 struct {
	ID                           int64            `json:"รหัสสถานี"`					// example:`124`  รหัสสถานี serial number
	AGENCY_ID                    int64            `json:"agency_id"`                    // example:`7`  รหัสหน่วยงาน agency's serial number
	PROVINCE_NAME                *json.RawMessage `json:"province_name"`                // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME                  *json.RawMessage `json:"amphoe_name"`                  // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME                  *json.RawMessage `json:"tumbon_name"`                  // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	WATERQUALITY_STATION_NAME    *json.RawMessage `json:"waterquality_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	WATERQUALITY_STATION_LAT     float64          `json:"waterquality_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	WATERQUALITY_STATION_LONG    float64          `json:"waterquality_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	WATERQUALITY_STATION_OLDCODE string           `json:"waterquality_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=069
// @Summary 		แหล่งน้ำดิบ (ข้อมูลแหล่งน้ำ)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_069 successful operation
type Metadata_069 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`7`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`117`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=4HTgv1HpJiC0dRGVu8fjH0m5NFhcgzTpJVmKe6-pKAtKpsf7dKYC3VBHmlu8UkKgoIy-LP0xKoGkrRqYIFLGtZQ1-G3AzJunJ5PGwifpxeKVJ3yEtN2FpyAVLkGMiIosSP9OYXw0FTYxLA_qQSq1FQBf5PMNH8nOG30OwQVZI5nRaoETSS8m2pL3HqD5ndWe`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แหล่งน้ำดิบ (ข้อมูลแหล่งน้ำ).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แหล่งน้ำดิบ (ข้อมูลแหล่งน้ำ)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=074
// @Summary 		ข้อมูลอ่างเก็บน้ำขนาดใหญ่ (รายวัน)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_074 successful operation
type Metadata_074 struct {
	DAM_ID       int64            `json:"dam_id"`       // example:`9` รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. dam's serial number
	DAM_DATE     string           `json:"dam_date"`     // example:`2006-01-02` วันที่เก็บข้อมูล record date
	DAM_LEVEL    float64          `json:"dam_level"`    // example:`57.57` ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.) last water level
	DAM_STORAGE  float64          `json:"dam_storage"`  // example:`10.0249` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
	DAM_INFLOW   float64          `json:"dam_inflow"`   // example:`2.24` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
	DAM_RELEASED float64          `json:"dam_released"` // example:`0.09` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
	DAM_SPILLED  float64          `json:"dam_spilled"`  // example:`2.7479` ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง spilled water volume
	QC_STATUS    *json.RawMessage `json:"qc_status"`    // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=075
// @Summary 		ข้อมูลอ่างเก็บน้ำขนาดใหญ่ (รายชัวโมง)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_075 successful operation
type Metadata_075 struct {
	DAM_ID       int64            `json:"dam_id"`       // example:`9` รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. dam's serial number
	DAM_DATETIME string           `json:"dam_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บข้อมูล record date
	DAM_LEVEL    float64          `json:"dam_level"`    // example:`57.57` ระดับน้ำกักเก็บปัจจุบัน ม.(รทก.) last water level
	DAM_STORAGE  float64          `json:"dam_storage"`  // example:`10.0249` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
	DAM_INFLOW   float64          `json:"dam_inflow"`   // example:`2.24` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
	DAM_RELEASED float64          `json:"dam_released"` // example:`0.09` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
	DAM_SPILLED  float64          `json:"dam_spilled"`  // example:`2.7479` ปริมาณระบายน้ำผ่านทางน้ำล้น (ล้าน ลบ.ม.) ทุกชั่วโมง spilled water volume
	QC_STATUS    *json.RawMessage `json:"qc_status"`    // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=076
// @Summary 		ฝน จากระบบโทรมาตรเขื่อนภูมิพล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_076 successful operation
type Metadata_076 struct {
	TELE_STATION_ID   int64   `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string  `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL1H        float64 `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=077
// @Summary 		ระดับน้ำ จากระบบโทรมาตรเขื่อนภูมิพล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_077 successful operation
type Metadata_077 struct {
	TELE_STATION_ID     int64   `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string  `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64 `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=079
// @Summary 		พื้นฐานโรงไฟฟ้า
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_079 successful operation
type Metadata_079 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`8`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`61`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-15T16:09:22+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Cfuq3JZKW-bxX_CxBHMkr5RhGg1v4qMpCuEknr-bDSNcZ5211PX5ZjSUuUN_-tWuZBeiCyjz1laO4QtOQEg6aP759hX48Sg9_wgKRZFuQrk1chU4oDSo3s-oKM52LDY4_eE5puWp5gCkNpHGTN5FbebRr0PkwcFpGgsuXNHgh9U`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`EGAT_Power_Plant_Iden.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`egat-powerplant`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=083
// @Summary 		พื้นฐานโทรมาตรเขื่อนภูมิพล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_083 successful operation
type Metadata_083 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=084
// @Summary 		พื้นฐานเขื่อน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_084 successful operation
type Metadata_084 struct {
	ID                      int64            `json:"id"`                      // example:`49` รหัสข้อมูลเขื่อน dam's serial number
	AGENCY_ID               int64            `json:"agency_id"`               // example:`8` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	DAM_NAME                *json.RawMessage `json:"dam_name"`                // example:`{"th":"รัชชประภา","en":"RAJJAPRABHA DAM","jp":" "}` ชื่อเขื่อน
	MAX_WATER_LEVEL         float64          `json:"max_water_level"`         // example:`97.65` ระดับกักเก็บสูงสุด [ม.(รทก.)] max water level
	NORMAL_WATER_LEVEL      float64          `json:"normal_water_level"`      // example:`95` ระดับกักเก็บปกติ [ม.(รทก.)] normal water level
	MIN_WATER_LEVEL         float64          `json:"min_water_level"`         // example:`62` ระดับกักเก็บต่ำสุด [ม.(รทก.)] min water level
	NORMAL_WATERGATE_LEVEL  float64          `json:"normal_watergate_level"`  // example:`95` เริ่มเปิดบานระบายน้ำที่ระดับ normal  [ม.(รทก.)] normal watergate level
	EMER_WATERGATE_LEVEL    float64          `json:"emer_watergate_level"`    // เริ่มเปิดบานระบายน้ำที่ระดับ Emergency  [ม.(รทก.)] mergency watergate level
	SERVICE_WATERGATE_LEVEL float64          `json:"service_watergate_level"` // เริ่มเปิดบานระบายน้ำที่ระดับ service [ม.(รทก.)] service watergate level
	MAX_OLD_STORAGE         float64          `json:"max_old_storage"`         // example:`90.7` ระดับน้ำสูงสุดที่เคยเก็บกัก[ม.(รทก.)] max old storage
	MAXOS_DATE              string           `json:"maxos_date"`              // example:`1997-09-28` วันที่เริ่มจัดเก็บระดับน้ำสูงสุดที่เคยเก็บกัก max old storage date
	MIN_OLD_STORAGE         float64          `json:"min_old_storage"`         // example:`64.39` ระดับน้ำต่ำสุดที่เคยเก็บกัก [ม.(รทก.)] min old storage
	MINOS_DATE              string           `json:"minos_date"`              // example:`1992-07-25` วันที่เริ่มจัดเก็บระดับน้ำต่ำสุดที่เคยเก็บกัก min old storage date
	TOP_SPILLWAY_LEVEL      float64          `json:"top_spillway_level"`      // example:`95.5` ระดับขอบบนของบานประตู spillway [ม.(รทก.)] top spillway level
	RIDGE_SPILLWAY_LEVEL    float64          `json:"ridge_spillway_level"`    // example:`87.5` ระดับสันของบานประตู spillway  [ม.(รทก.)] ridge spillway level
	MAX_STORAGE             float64          `json:"max_storage"`             // example:`6144.38` ปริมาตรน้ำที่ระดับเก็บกักสูงสุด  [ล้าน ลบ.ม.] max storage
	NORMAL_STORAGE          float64          `json:"normal_storage"`          // example:`5638.84` ปริมาตรน้ำที่ระดับเก็บกักปกติ [ล้าน ลบ.ม.] normal storage
	MIN_STORAGE             float64          `json:"min_storage"`             // example:`1351.54` ปริมาตรน้ำที่ระดับเก็บกักต่ำสุด [ล้าน ลบ.ม.] min storage
	USES_WATER              float64          `json:"uses_water"`              // example:`4200` ปริมาตรน้ำที่ใช้งานได้ [ล้าน ลบ.ม.] uses water
	AVG_INFLOW              float64          `json:"avg_inflow"`              // example:`2578.85` ปริมาณน้ำไหลเข้าเฉลี่ย [ล้าน ลบ.ม.] average inflowing water
	AVG_INFLOW_INTYEAR      string           `json:"avg_inflow_intyear"`      // example:`1986` ปีเริ่มต้นที่บันทึกปริมาณน้ำไหลเข้าเฉลี่ย start year (average inflowing water)
	MAX_INFLOW              float64          `json:"max_inflow"`              // example:`161.62` ปริมาณน้ำไหลเข้าสูงสุด [ล้าน ลบ.ม.] max inflowing water
	MAX_INFLOW_DATE         string           `json:"max_inflow_date"`         // example:`1997-08-24` วันที่เริ่มบันทึกปริมาณน้ำไหลเข้าสูงสุด start date (max inflowing water)
	DOWNSTREAM_STORAGE      float64          `json:"downstream_storage"`      // example:`1100` ความจุท้ายน้ำ [ลบ.ม./วินาที] downstream storage
	WATER_SHED              float64          `json:"water_shed"`              // example:`1435` พื้นที่รับน้ำ [ตร. กม.] water shed
	RAINFALL_YEARLY         float64          `json:"rainfall_yearly"`         // example:`1967.46` ปริมาณน้ำฝนเฉลี่ยต่อปี [มม.] rainfall (yearly)
	POWER_INSTALL           float64          `json:"power_install"`           // example:`240` กำลังผลิตติดตั้ง [MW]  power install
	POWER_INTAKE_STORAGE    float64          `json:"power_intake_storage"`    // example:`721.17` Storage at power intake sill level  [ล้าน ลบ.ม.] storage at power intake sill level
	POWER_INTAKE_LEVEL      float64          `json:"power_intake_level"`      // example:`53` Power intake sill Level [ม.(รทก.)] power intake sill level
	TAILRACE_LEVEL          float64          `json:"tailrace_level"`          // example:`12` Tailace Normal Level  [ม.(รทก.)] tailace normal level
	USED_GENPOWER           float64          `json:"used_genpower"`           // example:`7.98` การใช้น้ำในการผลิตต่อหน่วย [cms/kwhr] used water power
	DAM_LAT                 float64          `json:"dam_lat"`                 // example:`8.966667` พิกัดของเขื่อน latitude
	DAM_LONG                float64          `json:"dam_long"`                // example:`98.783333` พิกัดของเขื่อน longitude
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=085
// @Summary 		ประกาศเตือนภัยจาก กฟผ.
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_085 successful operation
type Metadata_085 struct {
	AGENCY_ID         int64  `json:"agency_id"`         // example:`2` รหัสหน่วยงาน agency number
	FLOOD_DATETIME    string `json:"flood_datetime"`    // example:`2006-01-02T15:04:05Z07:00` วันและเวลาที่ประกาศสถานการณ์น้ำ DateTime of flood
	FLOOD_NAME        string `json:"flood_name"`        // example:`รายงานสถานการณ์ธรณีพิบัติภัยประจำวัน วันอังคารที่ ๒๒ สิงหาคม พ.ศ. ๒๕๖๐` ชื่อสถานการณ์น้ำ name of flood
	FLOOD_LINK        string `json:"flood_link"`        // example:`http://www.dmr.go.th/ewt_news.php?nid=102862` ลิ้งที่แสดงสถานการณ์น้ำ flood link
	FLOOD_DESCRIPTION string `json:"flood_description"` // example:` ` รายละเอียดสถานการณ์น้ำ description of flood
	FLOOD_REMARK      string `json:"flood_remark"`      // example:` ` หมายเหตุ
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=087
// @Summary 		ปริมาณการใช้น้ำจากเขื่อนในแต่ละฤดูกาล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_087 successful operation
type Metadata_087 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`8`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`110`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=ueXDaDdxOJqcs-RuebMRSCyj3HVrgCVpiY_SV50CQmZI9kaNVxBIaO0n-pnz8U8skqbFvd7hdEgQ7fMBIZNbPNLqkEw5ldphZYN-FW0kdjOe523AZEqpyp6_A8sfsOKm4E_7dmh-cgdvThM3AZYFHt6mkPLpWAYFYHSHLeJxIpIDr1KSdiGM6-FsjY176RhRTGzt08iIO9XmW0FiZFaUO3UA7kltB1jKJvr2IKP37l8`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ปริมาณการใช้น้ำจากเขื่อนในแต่ละฤดูกาล.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ปริมาณการใช้น้ำจากเขื่อนในแต่ละฤดูกาล`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=088
// @Summary 		ผลกระทบจากการกักเก็บน้ำที่ผ่านมา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_088 successful operation
type Metadata_088 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`8`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`111`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=iP0QShKVUtYiGwDK6nkky_-GBsjMIBpniARYzjJB9uFKskO6HMZ--oByA3I3QlVSobFrNADmivEjuHmE7zKZWcL5PDhhMlBSwDXg13syjaRWu7QKayvRtCCAqs1f8aaaOf-ozr16Q09ihmsjgOETqWl2WwIQsD1gteeYP-lGAbgoS5UsaPCHYM9Kh1B6bMb4LFA20CgoAZJdkWZVifgD9g`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ผลกระทบจากการกักเก็บน้ำที่ผ่านมา.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ผลกระทบจากการกักเก็บน้ำที่ผ่านมา`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=098
// @Summary 		ฝน รายชั่วโมง จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_098 successful operation
type Metadata_098 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME  string           `json:"rainfall_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL_DATE_CALC string           `json:"rainfall_date_calc"` // example:`2006-01-02` วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall
	RAINFALL10M        float64          `json:"rainfall10m"`        // example:`0` ปริมาณน้ำฝนทุก 10 นาที Rainfall Every 10 minute
	RAINFALL1H         float64          `json:"rainfall1h"`         // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
	RAINFALL24H        float64          `json:"rainfall24h"`        // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	RAINFALL_TODAY     float64          `json:"rainfall_today"`     // ปริมาณน้ำฝนสะสมวันนี้
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=099
// @Summary 		อุณหภูมิ จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_099 successful operation
type Metadata_099 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	TEMP_DATETIME   string           `json:"temp_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าอุณหภูมิ record date
	TEMP_VALUE      float64          `json:"temp_value"`      // example:`26.46` ค่าอุณหภูมิ temperature value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=100
// @Summary 		ความกดอากาศ จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_100 successful operation
type Metadata_100 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	PRESSURE_DATETIME string           `json:"pressure_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความกดอากาศ record date
	PRESSURE_VALUE    float64          `json:"pressure_value"`    // example:`1012.67` ค่าความกดอากาศ pressure value
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=101
// @Summary 		ความชื้นสัมพัทธ์ จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_101 successful operation
type Metadata_101 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	HUMID_DATETIME  string           `json:"humid_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความชื้นสัมพัทธ์ record date
	HUMID_VALUE     float64          `json:"humid_value"`     // example:`95.33` ค่าความชื้นสัมพัทธ์ humid value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=102
// @Summary 		ความเข้มแสง จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_102 successful operation
type Metadata_102 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`151` รหัสสถานีโทรมาตร tele station's serial number
	SOLAR_DATETIME  string           `json:"solar_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความเข้มแสง record date
	SOLAR_VALUE     float64          `json:"solar_value"`     // example:`120.37` ค่าความเข้มแสง solar's value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=103
// @Summary 		ระดับน้ำ จากข้อมูลโทรมาตร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_103 successful operation
type Metadata_103 struct {
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64          `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=105
// @Summary 		ข้อมูลพื้นฐานของสถานีโทรมาตร เช่น รหัส,ชื่อ,พิกัด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_105 successful operation
type Metadata_105 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	SUBBASIN_ID          int64            `json:"subbasin_id"`          // example:`198` รหัสลุ่มน้ำสาขา subbasin number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	TELE_STATION_TYPE    string           `json:"tele_station_type"`    // example:`W` ชนิดของโทรมาตร (เช่น ระดับน้ำ)
	LEFT_BANK            float64          `json:"left_bank"`            // example:`1.782` ระดับตลิ่ง (ซ้าย) left bank level
	RIGHT_BANK           float64          `json:"right_bank"`           // example:`0.646` ระดับตลิ่ง (ขวา) right bank level
	GROUND_LEVEL         float64          `json:"ground_level"`         // example:`-2.678` ระดับท้องน้ำ ม.รทก ground water level
	OFFSET               float64          `json:"offset"`               // ค่า offset ระดับน้ำ การแสดงผลระดับน้ำให้คำนวณค่านี้ด้วย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=106
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลระดับน้ำของลุ่มน้ำเจ้าพระยา สสนก. ล่าสุด (CPY)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_106 successful operation
type Metadata_106 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=107
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลอัตราการไหลของลุ่มน้ำเจ้าพระยา กรมชลฯ ล่าสุด (CPY)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_107 successful operation
type Metadata_107 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=108
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลระดับน้ำของลุ่มน้ำ ชี มูล สสนก. ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_108 successful operation
type Metadata_108 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=109
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลอัตราการไหลของลุ่มน้ำ ชี มูล กรมชลฯ ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_109 successful operation
type Metadata_109 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=110
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลระดับน้ำของลุ่มน้ำภาคตะวันออก สสนก. ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_110 successful operation
type Metadata_110 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=111
// @Summary 		คาดการณ์ระดับน้ำท่วมจากข้อมูลอัตราการไหลของลุ่มน้ำภาคตะวันออก กรมชลฯ ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_111 successful operation
type Metadata_111 struct {
	FLOODFORECAST_STATION_ID int64   `json:"floodforecast_station_id"` // example:`9` รหัสสถานีคาดการณ์น้ำท่วม
	FLOODFORECAST_DATETIME   string  `json:"floodforecast_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์น้ำท่วม
	FLOODFORECAST_VALUE      float64 `json:"floodforecast_value"`      // example:`31` ข้อมูลคาดการณ์น้ำท่วมจากระดับน้ำ (ม.รทก) และอัตราการไหล (m3/s) โดยดูที่หน่วยของแต่ละสถานี
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=119
// @Summary 		ผลสำรวจหน้าตัดและสภาพลำน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_119 successful operation
type Metadata_119 struct {
	SECTION_STATION_ID int64   `json:"section_station_id"` // example:`9` รหัสสถานีของข้อมูลภาพตัดลำน้ำ section station's serial number
	POINT_ID           string  `json:"point_id"`           // ตำแหน่งของการจุดลำน้ำ
	WATER_LEVEL_MSL    string  `json:"water_level_msl"`    // example:`12` ระดับน้ำ หน่วย : ม.รทก. water level (msl)
	DISTANCE           string  `json:"distance"`           // example:`23` ระยะทาง หน่วย : เมตร distance
	REMARK             string  `json:"remark"`             // example:`LB` ตำแหน่งที่วัด LB : left bank RB : right bank  CL / location
	SECTION_LAT        float64 `json:"section_lat"`        // ตำแหน่งละติจูดของการวัด
	SECTION_LONG       float64 `json:"section_long"`       // ตำแหน่งลองติจูดของการวัด
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=120
// @Summary 		ข้อมูลแหล่งน้ำเครือข่ายชุมชน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_120 successful operation
type Metadata_120 struct {
	MOOBAN     string  `json:"mooban"`     // example:`เมืองแปง ม.1` หมู่บ้าน address number
	CAPACITY   float64 `json:"capacity"`   // example:`3700` ความจุ (ลูกบาศก์เมตร) capacity
	AGENCY_ID  int64   `json:"agency_id"`  // example:`5` รหัสหน่วยงาน agency's serial number
	GEOCODE_ID int64   `json:"geocode_id"` // example:`4966` ขอบเขตการปกครองของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=121
// @Summary 		ข้อมูลพื้นฐาน 25 ลุ่มน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_121 successful operation
type Metadata_121 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`73`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-02-14T10:16:15+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=PdkvnZpoLzHT25nFh2xC1qzXdPp5ALUuF7e8uaYPjJKwOYFEKd9LShJM3E2Bh--c_v8Amg5nS8LFZhnujV0SC6pVnRwurMzR57oSr2J4ugopuweLEqn_uFA0DJHWydUQPgasRfib0soOIKCykqs4sQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`25-westside_south.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`haii-25_basin-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=122
// @Summary 		รายงานสถานการณ์น้ำ1 วัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_122 successful operation
type Metadata_122 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`74`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=zhm_g-cCimNPps3uXLoa7_6yVMtes0fupXuGBpCrbk9YN7OJok7xxW3i5AcguVWLN4f8D78XqmGMWbfmjjAjJIu4MK1SrqyD4clzaJBtS48ZYbvgu1--GJ74qcAwmlfnP-zwBBogRqLhHPLRQPVGCIUWpytui-GakzUOqDYbIiOiMKN4hPLYAkXBhwKewFyp`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`รายงานสถานการณ์น้ำ1 วัน.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`รายงานสถานการณ์น้ำ1 วัน`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=123
// @Summary 		รายงานสถานการณ์น้ำรายสัปดาห์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_123 successful operation
type Metadata_123 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`75`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-02-14T11:10:31+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=T5dbLmn-PLfleZTsH-6eu88Y_65414TwONl8WtKEsE4XRcaqrWZufQR4m7SqFbUSYBPnOuGERXuVOpwIg-4LqJoj3EDjSkUtLM6AMsZr9h_-VbnsVPPyMZuop09HVBr3rCxy8QGqviF4KTErbHUHbdkK0z5_gWsBoEPK77vt_2e4VO1ReNAGu4FD123eEi9x`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`751_60WeeklyReportNo06.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`haii-report_water_situation_weekly-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=124
// @Summary 		รายงานสถานการณ์น้ำรายเดือน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_124 successful operation
type Metadata_124 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`76`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=9KYaWUHmz14mYk0T2c6fElW8Bst7DC2BsO3OkOag44Ughb4iGd6msJ2OEAPG55bpULWKFmvLnYWUi62UVvX1UHcgIqdVVft7rkWnattdPS3RXqKcaGC07kPoxXJpRvn5cOXBnpyC5Hj3EGjzuBzR6vhrrwTrqU06mGytk4ZTOccPaKpdhPgf9c5e3M0Tbp7X`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`รายงานสถานการณ์น้ำรายเดือน.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`รายงานสถานการณ์น้ำรายเดือน`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=125
// @Summary 		รายงานสถานการณ์น้ำ1 ปี
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_125 successful operation
type Metadata_125 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`77`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-02-15T09:10:08+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=D7fy8fpjcci8AOSycfabpsoKtslyHPMxNxHPiiEvcXsS2ythM01nCHv_fQQoF015BXmC8q8xaGoYnih8wkQ6W-tlrmMJ3QrgNCMnBX_vZgGMDbKi5-0bx1T9IgaxdXjW8EZncQgjmWPkSV3sXArqnQ4ZQLPfhNm4DmpNzXJbnmqEke0v-fI-YAGUE7GPWh_r`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`waterreport_2556.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`haii-report_water_situation_yearly-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=126
// @Summary 		ข้อมูลพื้นฐานทรัพยากรน้ำจังหวัด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_126 successful operation
type Metadata_126 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`78`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-02-14T12:08:27+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=PulB5itvJtlcK2-N5LpR3-JVIoafu5jdzacM-yokINIqVTfR2McpfMHhb3kZg2gqMoMaIlq-Yj58VmSh8qBg_BftEsvakyAMg-fijRiXmrKBSC2N95LdMP5QApwmxDh9oE1DG5vs5_FokPlMEeoi568prjL0J5NVHoUAy_elwhk`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`central_final.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`haii-water_resource_province-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=127
// @Summary 		ข้อมูลพื้นฐานทรัพยากรน้ำตำบล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_127 successful operation
type Metadata_127 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`79`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=aoureq7FyMBSW7VF2EdqHauGyue8tHHR6TFnG6spM0us9vXTcLnQq88yinoeXkhptRrzyOZdgkPgbDagGVsIm4u_o6DWj3IAAfy0I_dT4GZ1QBxaBDzfgMpBDfwRawkHt1TgpuytDS6HyDMEe0tkd0EujM1-vLdSrz9IQA217HOkJZmxfEEzZqfTYsa2boIiWujLA-ocPodsWmTLGkFz8Q`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลพื้นฐานทรัพยากรน้ำตำบล.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลพื้นฐานทรัพยากรน้ำตำบล`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=128
// @Summary 		ข้อมูลตรวจอัตโนมัติฝน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_128 successful operation
type Metadata_128 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string           `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL15M       float64          `json:"rainfall15m"`       // example:`0` ปริมาณน้ำฝนทุก 15 นาที Rainfall Every 15 minute
	RAINFALL1H        float64          `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
	RAINFALL3H        float64          `json:"rainfall3h"`        // example:`3` ปริมาณน้ำฝนทุก 3 ชั่วโมง Rainfall Every 3 hours
	RAINFALL6H        float64          `json:"rainfall6h"`        // example:`3.5` ปริมาณน้ำฝนทุก 6 ชั่วโมง Rainfall Every 6 hours
	RAINFALL24H       float64          `json:"rainfall24h"`       // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=129
// @Summary 		ข้อมูลตรวจอัตโนมัติระดับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_129 successful operation
type Metadata_129 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`956` รหัสสถานีโทรมาตร tele station's serial number
	WATERGATE_DATETIME string           `json:"watergate_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERGATE_IN       float64          `json:"watergate_in"`       // example:`0.04` ระดับน้ำด้านในประตูระบายน้ำ
	WATERGATE_OUT      float64          `json:"watergate_out"`      // example:`0.44` ระดับน้ำนอกประตูระบายน้ำ
	WATERGATE_OUT2     float64          `json:"watergate_out2"`     // example:`0` ระดับน้ำนอกประตูระบายน้ำ
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=130
// @Summary 		ข้อมูลตรวจอัตโนมัติฝนเขต
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_130 successful operation
type Metadata_130 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string           `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL5M        float64          `json:"rainfall5m"`        // example:`0` ปริมาณน้ำฝนทุก 5 นาที Rainfall Every 5 minute
	RAINFALL30M       float64          `json:"rainfall30m"`       // example:`0` ปริมาณน้ำฝนทุก 30 นาที Rainfall Every 30  minute
	RAINFALL1H        float64          `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
	RAINFALL3H        float64          `json:"rainfall3h"`        // example:`3` ปริมาณน้ำฝนทุก 3 ชั่วโมง Rainfall Every 3 hours
	RAINFALL6H        float64          `json:"rainfall6h"`        // example:`3.5` ปริมาณน้ำฝนทุก 6 ชั่วโมง Rainfall Every 6 hours
	RAINFALL12H       float64          `json:"rainfall12h"`       // example:`4.5` ปริมาณน้ำฝนทุก 12 ชั่วโมง Rainfall Every 12  hours
	RAINFALL24H       float64          `json:"rainfall24h"`       // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	RAINFALL_ACC      float64          `json:"rainfall_acc"`      // example:`241` ปริมาณน้ำฝนสะสม Rainfall Accumulate
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=135
// @Summary 		ข้อมูลตรวจวัดอัตโนมัติน้ำท่วมผิวจราจร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_135 successful operation
type Metadata_135 struct {
	FORD_STATION_ID          int64            `json:"ford_station_id"`          // example:`39` รหัสสถานีวัดระดับน้ำบนถนน
	FORD_WATERLEVEL_DATETIME string           `json:"ford_waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่วัดระดับน่ำบนถนน
	FORD_WATERLEVEL_VALUE    float64          `json:"ford_waterlevel_value"`    // example:`39.3` ค่าระดับน่ำบนถนน
	COMM_STATUS              string           `json:"comm_status"`              // example:`Connect, Fail` สถานะของเครื่องวัด meter status
	QC_STATUS                *json.RawMessage `json:"qc_status"`                // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=139
// @Summary 		ข้อมูลตรวจวัดอัตโนมัติอัตราการไหล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_139 successful operation
type Metadata_139 struct {
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	FLOW_RATE           float64          `json:"flow_rate"`           // อัตราการไหล
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=140
// @Summary 		ข้อมูลพื้นฐานของสถานีตรวจวัดอัตโนมัติ SCADA
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_140 successful operation
type Metadata_140 struct {
	ID                   int64   `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string  `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_LAT     float64 `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64 `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=141
// @Summary 		ข้อมูลพื้นฐานของสถานีตรวจวัดอัตโนมัติฝนเขต
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_141 successful operation
type Metadata_141 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=142
// @Summary 		ข้อมูลพื้นฐานของจุดติดตั้งสถานีตรวจวัดอัตโนมัติระดับน้ำในคลองสายหลัก
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_142 successful operation
type Metadata_142 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=144
// @Summary 		ข้อมูลพื้นฐานของจุดติดตั้งสถานีตรวจวัดอัตโนมัติวัดน้ำท่วมผิวจราจร
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_144 successful operation
type Metadata_144 struct {
	ID                   int64   `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string  `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_LAT     float64 `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64 `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	AGENCY_ID            int64   `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=146
// @Summary 		ข้อมูลพื้นฐานจุดติดตั้งตำแหน่งสถานีสูบน้ำ และข้อมูลประสิทธิภาพการสูบน้ำ (นอกระบบ SCADA)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_146 successful operation
type Metadata_146 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`10`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`84`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=CslmMt5yJyZDqDEHhw9hZibucjG8FMLowtFtrhH2DLRHyJFSYZZCebSbhsSYxlk7p_16gltpdlFp4LGmM9MGV7Nd1zMbsC5octmE4oWXMzaucXgPcxFzkg1eNTi2tyuJ6cLlWq0eyvGBIPKIkRzDuwtuLj7DO6KBN744aXsEM2oI0qdCUiRZqx_4Q8zbjdgivL-7FfFSDGn1queuWIaaeP93tynKwA_DROcTdBb2wSqmqJcHJZKaqzb9zzYQnqMomVYDWcOYln8EVGAkpdN6TdSXRsyo7mdyUDP8N4u2bPK7Jnhf_F-7C5704fmjU42iHB3fhp-GjYSpgf6R2e4TILOEcmyNK6GSIdactYrs1IbTf07-AQY4zNjV4h9Ha6gz8rhI167LRn117LOEitMQKoqKTHRi9R9bSNp13HcZ1gA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลพื้นฐานจุดติดตั้งตำแหน่งสถานีสูบน้ำ และข้อมูลประสิทธิภาพการสูบน้ำ (นอกระบบ SCADA).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลพื้นฐานจุดติดตั้งตำแหน่งสถานีสูบน้ำ และข้อมูลประสิทธิภาพการสูบน้ำ (นอกระบบ SCADA)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=152
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียม Thaichote
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_152 successful operation
type Metadata_152 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`41`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-08T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=i9GFhyoJGfGp5wMdgMmw0LfymLU3gIMgmTA1t8HrKvK6HkcLlRTEkBbm2NGho4DZ7s0W_2VoQB7yA_7JjABqyFO2u1OOHQkM52iv1USPTEYpmJimDAizPfFrognhvLEzTzWp16acl6fVdsP5tJDvzxNGm3bQckhYqLpwrwZtU2Oo1IWwZ_wKnIS1a-cHGBFZ12NXndpFI_MEPpxyifbBmw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`th1_262_030454_20171008_032000_mss_2a_fc.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA thaichote`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=153
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียม Radarsat 1
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_153 successful operation
type Metadata_153 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`39`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=dQtkfX70GnpRayrgpZuxLKV7P3c91QqOLYQ56B4b0vaADOwkxn__uZvwDflDSW8DBKr0zZEt6AXBLXTby1IH0q4yu8VDWLixXMiWupZLhd3oZbVfq435IcGOBRmBd1ubwRvMlna6mKuHTkgaYsuLbiCKUO2MUmkQ4yBGMQh4q9uhawNL3t5klIqk-3Hzi6jW9E3UkB-_AjeNxwVhwERsHA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่น้ำท่วมภาพถ่ายดาวเทียม Radarsat 1.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่น้ำท่วมภาพถ่ายดาวเทียม Radarsat 1`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=154
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียม Radarsat 2
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_154 successful operation
type Metadata_154 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`40`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-13T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=BuZU_VF8WVKDgOYx5WPrIdBizspieJWoNC5FcHYHNoclKvnrhtpbeLhKv5L-g6kjf1d_G5ejbDpvoZKhzNzDzhuANDELVAxD7I-CJOvRRty6fPrT94vpjS63kPThalRNRgq3IRUFF6C_tpsrazXl0x-PdPYwORXZzUhSdkkRctPGa1VUH_GPVn9LbGVOo22a`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rd2_20171013_RAD2_61_0031_0045_3.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA radar-sat2`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=155
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียม Cosmos SkyMed
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_155 successful operation
type Metadata_155 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`38`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-15T02:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=xccTlcbI4SZiHwLL6VycomYM71x68EPqa3jluCVqyFB-0hFZ7xPCclV_JfkUaQBB8ISaGph7wIUeIDmr8M2b1l2_JwE0qqo0tsbDrXhmGwnHpS60zLfeTyov3_WhdF6suPcVCNcNrc-RY7u7Mwf90A7eZ5VpBjQrZBbwAW6eyfhXSLGXq4pYeP0TQYnrQKB2`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`cm2_20171015_02_hv.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA cosmos skymed`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=156
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียมแบบModis Ndvi
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_156 successful operation
type Metadata_156 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`35`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-04-29T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=wgU--kp1YCcbcnNmA1opTYirXq1x0by9DBH69HBg1OwAZeTPUMnXIm8sM1WSYd3OewUJa2iEW9NhLTXHhrup9pmpXA6syUCj3XsxhabKdsqnt7TQHZwu7NFmtAvOubeZrKo7lABuc5EIGgc9jNxJPO16fq8z0z6gqQ6v2jmY54I`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`1D_NDVI_20180429_0000_Modis.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA modis`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=157
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียมแบบ Aqua
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_157 successful operation
type Metadata_157 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`36`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-05-17T07:12:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=LjllmPWfALRMBrcp5GdtfrJCRX1uEiDHbFuybiMaIPYVAnl-ts5pWM5EZb_j8a9tC-8ymTzxwY5Lx9MqCbckfubZZyKAYR-CR8gEd1taoJ3h6rN5qpF8yUHChhp7ImTzlgneKR0PQDD_i3xj1iB_fjBnqmxAtbjvMfftG6spUcQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`1D_NDVI_20180517_0712_Aqua.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA Aqua`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=158
// @Summary 		แผนที่น้ำท่วมภาพถ่ายดาวเทียมแบบ Terra
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_158 successful operation
type Metadata_158 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`37`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-14T04:03:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=QiDNhZkwAKP6ILyZkO4Dqno2JpX5dF4iEI53_-HtFtCOZHWHdWTZ0oFvmj5JntieuV8kU7XsS8XD15U7JICgCO8iRmg-D20z4IC_0-V88RuYtC6l_xkodK_W21_A3qh8bMybTpAudIG3s5rFS8ahtomnhtla-38mggKCCoz0VTP0Zbuc3iXOYxJvQjaHD4bH`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`201808140403Terra_HK_NDVI1D.tif.aux.xml`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GISTDA terra`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=159
// @Summary 		ข้อมูลคลื่นชายฝั่งและข้อมูลเรด้าตรวจวัดบริเวณอ่าวไทย (CODAR)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_159 successful operation
type Metadata_159 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`42`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=BCQilk_Pp36oFDz1OochfOY0FzDFT1QYPEvtNDtJiViIEY3h7t82XjLl31BYWUeWIPNABoQlYXFBKmDlCNadkLZxfugJXTPl22ddnPO3sfcDnpHzJfVZf2RHni_UM1ykbf2BrrJqbSkUk_dbs1d2It2HpQzZ8P_iGqjoOZdjUgn4JHBa_73Omzft2l3psEbq10MOiGQLQgWQxpIY81ZlgaQcVeMhb5a1l3kVKaapGABDosR4tCqae12xQziJOFzVvsiv4sa2Bog69YLrooKD_poEcUoHiogGNxoSb-SndFdQV_3efawXDbcR9TEvybDo`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลคลื่นชายฝั่งและข้อมูลเรด้าตรวจวัดบริเวณอ่าวไทย (CODAR).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลคลื่นชายฝั่งและข้อมูลเรด้าตรวจวัดบริเวณอ่าวไทย (CODAR)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=160
// @Summary 		ข้อมูลพื้นที่น้ำท่วมปี 2548-2554
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_160 successful operation
type Metadata_160 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`132`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-12-28T09:58:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=vzcGLYpimedhThXNATD2-yCoffzw7Pp4uv1LL3W1hoRBzQdAjG7QmMeGLaEnKFAOcXzvlI7ox0L1WXg_fKRz8oscPttjU-0whLJiC7PFofaeEPxi4F8owqYETf7i9Lc6DlbQyvYtAlcNXOjwMGQ_VsxqMkArHAgYa9MRYCyPK3YUbuLpebr7b97qFWVl658u`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`flood7days_20171113_20171107.zip`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`gist-flood_area-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=161
// @Summary 		ข้อมูลการกัดเซาะชายฝั่ง
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_161 successful operation
type Metadata_161 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`133`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2013-02-13T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=tiKVbZzhpEaFzsKm_zoDty3oI-aDVu6t2X_RA7FzVN8KAa7L23JBcQthBqrLfBPBJ6K7MGZ4jtYr775mzjZNZvxCSik63BYaj-79XtHRQuMEy8tQ7bbuaxp9wJm4MRPb9OxxdKnS5oseJw9MMvf0t76qI0SIs3d12wYe8CxUfRczifLW2y0j6R6Tq9Dm9g5V`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`cur_gulf_20130213_0700.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`gist-coastal_erosion-offline`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=162
// @Summary 		LIDAR
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_162 successful operation
type Metadata_162 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`134`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=TCTO7u02WhmUl9APweSdqarM9vfbCHBKTw2DhRqkzKhw7YMim2AL0xFus50PQYEDLh36w7JkO-2l-cZFUqPx6KUkD_qAnOUMMNicAyTmW4c`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`LIDAR.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`LIDAR`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=163
// @Summary 		Thaichote
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_163 successful operation
type Metadata_163 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`135`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=yKxQ4fzKk622nIGJSqb-rV6eReBYns-xb3oM39F9iO73ZDMPbGCMv1feFDHKLBp5Lt3_UPSrTY363EMW1qtIHzacI0Ju4rs2iL_gg9OZL_s`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`Thaichote .pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Thaichote `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=222
// @Summary 		ข้อมูลน้ำในเขื่อนใหญ่
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_222 successful operation
type Metadata_222 struct {
	DAM_ID              int64   `json:"dam_id"`              // example:`9` รหัสข้อมูลเขื่อนขนาดใหญ่ ของ กฟผ. dam's serial number
	DAM_DATE            string  `json:"dam_date"`            // example:`2006-01-02` วันที่เก็บข้อมูล record date
	DAM_STORAGE         float64 `json:"dam_storage"`         // example:`10.0249` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
	DAM_USES_WATER      float64 `json:"dam_uses_water"`      // example:`40.560001` ปริมาณน้ำที่ใช้ได้ uses water volume
	DAM_INFLOW          float64 `json:"dam_inflow"`          // example:`2.24` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
	DAM_RELEASED        float64 `json:"dam_released"`        // example:`0.09` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
	DAM_STORAGE_PERCENT float64 `json:"dam_storage_percent"` // example:`26.959999` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume

	DAM_USES_WATER_PERCENT float64          `json:"dam_uses_water_percent"` // example:`17.059999` เปอร์เซนต์ปริมาตรน้ำใช้การได้ (% รนก.) data form rid not cal/ percent of uses water volume
	DAM_INFLOW_ACC         float64          `json:"dam_inflow_acc"`         // example:`130` ปริมาตรน้ำไหลลงอ่างเก็บน้ำเฉลี่ยทั้งปี
	DAM_INFLOW_AVG         float64          `json:"dam_inflow_avg"`         // example:`24.167` ปริมาตรน้ำไหลลงอ่างเก็บน้ำสะสมตั้งแต่ต้นปี
	DAM_INFLOW_ACC_PERCENT float64          `json:"dam_inflow_acc_percent"` // example:`48.529999` เปอร์เซนต์ปริมาณน้ำไหลเทียบกับปริมาณน้ำไหลลงเขื่อนขนาดใหญ่เฉลี่ยรวมทั้งปี (%) data form rid not cal/ percent of inflowing water volume
	DAM_RELEASED_ACC       float64          `json:"dam_released_acc"`       // example:`60.988896` ปริมาตรน้ำระบายสะสมตั้งแต่ต้นปี
	QC_STATUS              *json.RawMessage `json:"qc_status"`              // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=223
// @Summary 		ข้อมูลน้ำในเขื่อนขนาดกลาง
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_223 successful operation
type Metadata_223 struct {
	MEDIUMDAM_ID              int64   `json:"mediumdam_id"`              // example:`9` รหัสข้อมูลเขื่อนขนาดกลาง  medium dam number
	MEDIUMDAM_STORAGE         float64 `json:"mediumdam_storage"`         // example:`15.801` ปริมาณน้ำกักเก็บปัจจุบัน (ล้าน ลบ.ม.) last water storage volume
	MEDIUMDAM_STORAGE_PERCENT float64 `json:"mediumdam_storage_percent"` // example:`57.043` เปอร์เซนต์ปริมาตรน้ำข้อมูลเขื่อนขนาดใหญ่  (% รนก.) data form rid not ca / percent of storage volume

	MEDIUMDAM_INFLOW     float64 `json:"mediumdam_inflow"`     // example:`0.878` ปริมาณน้ำไหลเข้าอ่างทุกชั่วโมง (ล้าน ลบ.ม) inflowing water volume
	MEDIUMDAM_RELEASED   float64 `json:"mediumdam_released"`   // example:`1.76` ปริมาณการระบายผ่านเครื่องทุกชั่วโมง (ล้าน ลบ.ม.) released water volume
	MEDIUMDAM_USES_WATER float64 `json:"mediumdam_uses_water"` // example:`14.651` ปริมาณน้ำที่ใช้ได้ uses water volume
	MEDIUMDAM_DATE       string  `json:"mediumdam_date"`       // example:`2006-01-02` วันที่เก็บข้อมูล record date
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=228
// @Summary 		สถานีวัดน้ำท่าจากศูนย์อุทกวิทยา ภาค1-8
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_228 successful operation
type Metadata_228 struct {
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_M        float64          `json:"waterlevel_m"`        // ระดับน้ำ เมตร รสม.
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=242
// @Summary 		เกณฑ์การเตือนภัยฝน 24 ชม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_242 successful operation
type Metadata_242 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`12`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`129`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-03-14T11:55:20+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Ful8qiZ-45r2jvzQZwcpnxdzh3k68tAXjEVz8VnZdnca_RRY_v0S5tgABLK8JJL3pCXlIkN2u5_HG3dwPJAmhs_Dm8yTB5SgTQ0RFmB4Q3D4erZVR9HWYLwmy6YBSO5z`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`Z61.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`rid-rain_threshold`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=244
// @Summary 		ฝน จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_244 successful operation
type Metadata_244 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME  string           `json:"rainfall_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL3H         float64          `json:"rainfall3h"`         // example:`3` ปริมาณน้ำฝนทุก 3 ชั่วโมง Rainfall Every 3 hours
	RAINFALL24H        float64          `json:"rainfall24h"`        // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
	RAINFALL_DATE_CALC string           `json:"rainfall_date_calc"` // example:`2006-01-02` วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=245
// @Summary 		ทิศทางและความเร็วลม จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_245 successful operation
type Metadata_245 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`60` รหัสสถานีโทรมาตร tele station's serial number
	WIND_DATETIME   string           `json:"wind_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความเร็วลม record date
	WIND_SPEED      float64          `json:"wind_speed"`      // example:`8` ค่าความเร็วลม wind value
	WIND_DIR_VALUE  float64          `json:"wind_dir_value"`  // example:`228` ค่าองศาของทิศทางลม wind direction value (degree)
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=246
// @Summary 		อุณหภูมิ จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_246 successful operation
type Metadata_246 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	TEMP_DATETIME   string           `json:"temp_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าอุณหภูมิ record date
	TEMP_VALUE      float64          `json:"temp_value"`      // example:`26.46` ค่าอุณหภูมิ temperature value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=247
// @Summary 		ความกดอากาศ จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_247 successful operation
type Metadata_247 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	PRESSURE_DATETIME string           `json:"pressure_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความกดอากาศ record date
	PRESSURE_VALUE    float64          `json:"pressure_value"`    // example:`1012.67` ค่าความกดอากาศ pressure value
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=248
// @Summary 		ความชื้นสัมพัทธ์ จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_248 successful operation
type Metadata_248 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	HUMID_DATETIME  string           `json:"humid_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความชื้นสัมพัทธ์ record date
	HUMID_VALUE     float64          `json:"humid_value"`     // example:`95.33` ค่าความชื้นสัมพัทธ์ humid value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=249
// @Summary 		อุณหภูมิที่สูงสุดและต่ำสุดใน 24 ชั่วโมงที่ผ่านมา จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_249 successful operation
type Metadata_249 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`3634` รหัสสถานีโทรมาตร tele station's serial number
	TEMPERATURE_DATE   string           `json:"temperature_date"`   // example:`2006-01-02` วันที่เก็บค่าอุณหภูมิ record date
	TEMPERATURE_VALUE  float64          `json:"temperature_value"`  // example:`26.2` ค่าอุณหภูมิ temperature value
	MAXTEMPERATURE     float64          `json:"maxtemperature"`     // example:`34.5` ค่าอุณหภูมิสูงสุด max temperature value
	DIFFMAXTEMPERATURE float64          `json:"diffmaxtemperature"` // example:`-1.3` ส่วนต่างค่าอุณหภูมิสูงสุด difference max temperature value
	MINTEMPERATURE     float64          `json:"mintemperature"`     // example:`25.2` ค่าอุณหภูมิต่ำสุด min temperature value
	DIFFMINTEMPERATURE float64          `json:"diffmintemperature"` // example:`-0.7` ส่วนต่างค่าอุณหภูมิต่ำสุด difference min  temperature value
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=250
// @Summary 		ฝนสะสมรายวัน จากสถานีโทรมาตรอัตโนมัติ (WMO)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_250 successful operation
type Metadata_250 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`135` รหัสสถานีโทรมาตร tele station number
	RAINFALL_DATETIME string           `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL_VALUE    float64          `json:"rainfall_value"`    // example:`3.2` ปริมาณฝนรายวัน เวลา 7:01 - 7:00  Rainfall daily
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=252
// @Summary 		ข้อมูลเรดาร์ตรวจอากาศเรดาร์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_252 successful operation
type Metadata_252 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`30`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-11T12:45:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=lSLh7efSI_IeRWAG2lWfv-lIp1sHQVRbM9gjRCbtWjQtr4m-BWA_UEDxIUITTnCNVn-A3_5xsfW44nkL-A-P5h4T9A5NyTH1DSh-Q26iJcEKvZ1lUP5r_KL_Wwc7pnPoMQQaE8wnIxMuUqYIODgMnQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`skn240_201808111245.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd-radar`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=253
// @Summary 		ข้อมูลแผนที่อากาศ แผนที่อากาศผิวพื้น ภาพ ณ เวลา 01.00 07.00 13.00 19.00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_253 successful operation
type Metadata_253 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`22`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-11T13:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=oW-RLI24X9G-JFKzHwCurjs096LNxLHZaAi6kBcmd-s9e4NCLy6C5tYdHWk94n28SRI_gtBjGYHgyIFcvS3qiHOoXzZ-S3ww6UXQfEVfnlICIppDhe4UO5pUHXuYkafgVWGKbdV6YHl3yOtd4bQOZflscExhu4713KIC5OeLIRUFRDrHugINxrjd9JADuOpS`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`2018-08-11_TopChart_13.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd-weather_map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=254
// @Summary 		แผนที่ลมชั้นบน ระดับ 925 hPa ภาพ ณ เวลา 01.00 07.00 13.00 19.00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_254 successful operation
type Metadata_254 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`29`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-11T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=TjBvAZB8GoF16eW6V8D3CSiR9bN3WZpoGcaOvswFBzT3tQMLWRoa8aPetv2HVF8lMPvnqtWQskXphDrLfwgXVnfZsiM6HUMr0AWrvZcNFSstUiY5F0xnKFuv6v1qfiws0giJefLTbhsTjzFgdwaIMrk7fq562kHcR2gRj-39WvWpcefPioVl0hnvPvQl443d`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`2018-08-11_07_UpperWind925.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd-windmap925`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=255
// @Summary 		แผนที่ลมชั้นบน ระดับ 850 hPa ภาพ ณ เวลา 01.00 07.00 13.00 19.00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_255 successful operation
type Metadata_255 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`28`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-11T13:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=xtKfAZLd5IU3ssE1bhZ45LX2w0CVvbBPhW7Yfa4VGtZeTs_8H_3ZOlQEcVyROmsae0WBrknNWUi-IcOctEgaY46uuFoYvTmkMmkCfMw7BYr_3bp8dFaMYqTUGYwkdzx2axGTk0uteLdTJ80duhPa6BYpl6YrIoxR37lOzY-wKCSrssuXCsxjT8kf-x1Bv6V0`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`2018-08-11_13_UpperWind850.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd-windmap850`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=256
// @Summary 		ข้อมูลภาพถ่ายดาวเทียมสภาพภูมิอากาศ Himawarii IR
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_256 successful operation
type Metadata_256 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`46`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T15:45:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=afkeG31Ow-Vr2_ge7v7bmZE8XajUPMkce1h3AX6p8KxCZczrrskcdOON198A7oU-Vl6keu2QMUNKXQQFh7t5kkl71IJY30XPKecWc0WojIn0SBZKc4HddsDQRnz1DjNiGF_PaaerCwL3YSJrS7l27J5lD8tVWepD7o5oE4tJGmI`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`IR201708280850.JPG`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd himawari-ir`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=259
// @Summary 		ภาพแผนที่ฝนสะสม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_259 successful operation
type Metadata_259 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`4`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Hh4d8hOFrQIQUAN2nU_j-Jiyy73vvxW-0RHgc-ltqwpiKNWw5izSdMBFPy_E3hAJSqTtR_KHBmWQRRDG2q8owJLO8k4d6FMIv44AtTcXeHBq_-FDuiWm-ACQR24fZVKfTBp_dYqsNo1cNPASTp55xg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ภาพแผนที่ฝนสะสม.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพแผนที่ฝนสะสม`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=260
// @Summary 		รายงานสรุปลักษณะอากาศ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_260 successful operation
type Metadata_260 struct {
	AGENCY_ID       int64            `json:"agency_id"`       // รหัสหน่วยงาน agency number
	PROVINCE_NAME   *json.RawMessage `json:"province_name"`   // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME     *json.RawMessage `json:"amphoe_name"`     // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME     *json.RawMessage `json:"tumbon_name"`     // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	WEATHER_DATE    string           `json:"weather_date"`    // วันที่พยากรณ์สภาพอากาศ
	OVERALL_FORCAST *json.RawMessage `json:"overall_forcast"` // ภาพรวมการพยากรณ์สภาพอากาศ
	REGION_FORCAST  *json.RawMessage `json:"region_forcast"`  // พยากรณ์สภาพอากาศรายภาค
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=261
// @Summary 		แผนที่แสดงดรรชนีความแห้งแล้งของฝนที่ต่างจากค่าปกติ (Standardized Precipitation Index: SPI)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_261 successful operation
type Metadata_261 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`4`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=G8fV6LeVo042iwnN6_lmfOoKslyN9IjYQZbZfCnchr2g4ztfV03eGdOzNRAbDez7vbrp9jSLMt37XAiq4X_MMd8y2Xd6abzqBcP18k2lCs-pUCVkTIae_9llZI1oWPDtozh2HfV4aG7y54zPwmwmaYHKHR90uicoRfD8NEpxs1wCUvFMcsQXl6BL5iNzYHFU9khPVGDBWg6YPjTFmIrbnqBZGZ_IV4B4QEBTJ0-t2La3JH7UOgCG9rFqFLnU7LR6nLNPMx42_uuQEK-neQ94l96JkMsd0KOWHgU2byQAAX0AN8X2wmX0yVTSThwGtxqZ19ujYnpftXh_Nt8IjGX77w`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่แสดงดรรชนีความแห้งแล้งของฝนที่ต่างจากค่าปกติ (Standardized Precipitation Index: SPI).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่แสดงดรรชนีความแห้งแล้งของฝนที่ต่างจากค่าปกติ (Standardized Precipitation Index: SPI)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=262
// @Summary 		แผนที่แสดงดรรชนีความชื้นที่เป็นประโยชน์สำหรับพืช (Moisture Available Index: MAI)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_262 successful operation
type Metadata_262 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`4`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=VOOKpFMGDpH73eeS3RZmmdVinLJmsac_uYTTZuWWD5eD7YA_z6Cc7g4kTIpuMUp44Pah_F2UkV5bWI2k-vqAXL0_4_uhakPvtz8ezpWfmaWf5YbHnq1fHedbexcTB49RwXqoo2E7HXpwaIJ76N8-pAVUoIdRX5rLE9r2BCNBenPB7IerYOK8lGwJ7wRKO7wGoQM6aqjw1dG6R9FcK5oWX9veZzwtWU_wuGHGzrWyo0S0HcCtSQKIcSYIj-5_U3Y2knxyTPKLdlJr9rnieFqLA03WsPMy462I38eUF_nU6eHlGIou3gt-ZgrEUsa1TJjU`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนที่แสดงดรรชนีความชื้นที่เป็นประโยชน์สำหรับพืช (Moisture Available Index: MAI).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนที่แสดงดรรชนีความชื้นที่เป็นประโยชน์สำหรับพืช (Moisture Available Index: MAI)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=271
// @Summary 		ข้อมูลคุณภาพน้ำของสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ 5 พารามิเตอร์(PH,DO,EC,TEMP,Turbidy)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_271 successful operation
type Metadata_271 struct {
	WATERQUALITY_ID           int64            `json:"waterquality_id"`           // example:`117` รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ
	WATERQUALITY_DATETIME     string           `json:"waterquality_datetime"`     // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าคุณภาพน้ำอัตโนมัติ
	WATERQUALITY_STATUS       string           `json:"waterquality_status"`       // example:`ปกติ` สถานะของคุณภาพน้ำ
	WATERQUALITY_PH           float64          `json:"waterquality_ph"`           // example:`4` ความเป็นกรด-ด่าง
	WATERQUALITY_DO           float64          `json:"waterquality_do"`           // example:`0` ออกซิเจนละลายในน้ำ หน่วย mg/l
	WATERQUALITY_CONDUCTIVITY float64          `json:"waterquality_conductivity"` // example:`232` ความนำไฟฟ้าในน้ำ หน่วย uS/cm ชื่อเต็ม The Electrical Conductivity (ec)
	WATERQUALITY_TEMP         float64          `json:"waterquality_temp"`         // example:`30.22` อุณหภูมิน้ำ หน่วย ?C
	WATERQUALITY_SALINITY     float64          `json:"waterquality_salinity"`     // example:`0.09` ค่าความเค็ม
	QC_STATUS                 *json.RawMessage `json:"qc_status"`                 // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=278
// @Summary 		ข้อมูลพื้นฐานสถานีตรวจวัดคุณภาพอากาศ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_278 successful operation
type Metadata_278 struct {
	ID                  int64            `json:"id"`                  // example:`9` รหัสเครื่องวัดอากาศ air station's serial number
	AIR_STATION_OLDCODE string           `json:"air_station_oldcode"` // example:`59t` รหัสเครื่องวัดอากาศเดิมของแต่ละหน่วยงาน old tele station's serial number
	AIR_STATION_NAME    *json.RawMessage `json:"air_station_name"`    // example:`{"en":"The Government Public Relations Department","th":"กรมประชาสัมพันธ์"}` ชื่อสถานีโทรมาตร tele station's name
	AGENCY_ID           int64            `json:"agency_id"`           // example:`14` รหัสหน่วยงาน agency's serial number
	AIR_STAITON_TYPE    string           `json:"air_staiton_type"`    // example:`GROUND` ชนิดของเครื่องวัดอากาศ (m = mobile เคลื่อนที่ได้ )
	AIR_STATION_LAT     float64          `json:"air_station_lat"`     // example:`13.783143` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	AIR_STATION_LONG    float64          `json:"air_station_long"`    // example:`100.540529` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=279
// @Summary 		ข้อมูลคุณภาพอากาศ (ปริมาณฝุ่น PM10)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_279 successful operation
type Metadata_279 struct {
	AIR_STATION_ID int64            `json:"air_station_id"` // example:`9` รหัสสถานีโทรมาตร tele station's serial number
	AIR_DATETIME   string           `json:"air_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าคุณภาพอากาศ
	AIR_SO2        float64          `json:"air_so2"`        // example:`1` ก๊าซซัลเฟอร์ไดออกไซด์ unit= ppb
	AIR_NO2        float64          `json:"air_no2"`        // example:`1` ก๊าซไนโตรเจนไดออกไซด์ unit= ppb
	AIR_CO         float64          `json:"air_co"`         // example:`0.12` ก๊าซคาร์บอนมอนนอกไซด์ unit=ppm
	AIR_PM10       float64          `json:"air_pm10"`       // example:`1` ฝุ่นละอองขนาดไม่เกิน 10 ไมครอน unit=?g/m?
	AIR_AQI        float64          `json:"air_aqi"`        // example:`1` ดัชนีคุณภาพอากาศ (Air quality Index)
	QC_STATUS      *json.RawMessage `json:"qc_status"`      // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=280
// @Summary 		เส้นทางหลวง (roadnet)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_280 successful operation
type Metadata_280 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`15`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`59`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-14T21:24:56+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=vDKTcoT98s1wbPRit7SAJ9UgHlIDSoLfAp8hXk_WgarJAsTBB2Gy5LkDaLGX7DYHGtjF9kYwDJbcbbaNi8WToxvTxbbpYRQgegSnX8dY-GeyykZTL4cJG9pSM-sBmzMHuqpWQgke-v5ekkBXaY2j7g`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`DOHRoad.shp`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`doh-roadnet`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=284
// @Summary 		การวิเคราะห์จุดอ่อนทางหลวง VI
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_284 successful operation
type Metadata_284 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`15`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`58`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-14T21:22:43+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=jfW6qIvhop5rOY_ObKNQ7bBTkGz3fu9z51YVFdDdz3hMTR09k2KPGPstAjfPX1W2kHhzKW0PX5IXpZLiKwXsoHMILjc9EdKj12ebgJsmPjJOgBXy8TNPFPJAPuLhjXFoKgp8aZlu31Yk7CjGhi-YAuXGrqOZIjTwgnF2FYNVLJD1qL8GXM60iu8TqsLpFXsOMhUZKrJ6VGU9ev8IZFc23neUA8Rt3lECdaNvcImDR8RL5Ek4LKxhsnNHJMMLlfZr`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`6 การวิเคราะห์จุดอ่อนทางหลวง  VIVI Final Report.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`doh-vi`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=291
// @Summary 		ข้อมูลพื้นที่เสี่ยงอุทกภัยและโคลนถล่ม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_291 successful operation
type Metadata_291 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`16`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`104`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T14:11:49+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=QhOYGdS0shHYWgaBhmW-806LiePMX1F48LsEsjq4O69-mmc1oaPQIqPWkfm3Hc4z3NePWDJ3jmDsbBBcy3-HcVfHacXoKlhGc78m3b9otC_prDKaXWe4U640omTevwhfbqGJdISSsWCBLWlbobNTADlpBA7P2lqnVQi-4l8Y0bA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`flood9y_thai.shp`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`disaster-geohazard_area`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=298
// @Summary 		ข้อมูลพื้นที่เสี่ยงสึนามิ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_298 successful operation
type Metadata_298 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`16`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`105`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T14:06:11+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=vwcCguRybm2j06B8ByKYt8U-9lxXj682_kXNUitYo4P1dlW2cnWTK1UqA3CQN-z5zwFPzP7taQmk7wt9Mm4uD372QBn2RD_HYsUC94VTj1Q4rJGo6iQtG9CPbKT9jGNHIHF3fapKnJZnOd3jhXn5RfjL9PNZ3lLq6tbnu6wp5kIxmhsE0A_ecb80d7H9mWoK`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`utm_site_phase_i.shx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`disaster-tsunamihazard_area`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=332
// @Summary 		ข้อมูลปริมาณน้ำฝน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_332 successful operation
type Metadata_332 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME  string           `json:"rainfall_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL_DATE_CALC string           `json:"rainfall_date_calc"` // example:`2006-01-02` วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall
	RAINFALL10M        float64          `json:"rainfall10m"`        // example:`0` ปริมาณน้ำฝนทุก 10 นาที Rainfall Every 10 minute
	RAINFALL1H         float64          `json:"rainfall1h"`         // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
	RAINFALL24H        float64          `json:"rainfall24h"`        // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=334
// @Summary 		ข้อมูลการใช้น้ำของนิคมอุตสาหกรรม (รายงานสถิติ)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_334 successful operation
type Metadata_334 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`22`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`111`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=UWKX7DxqIhZZKCWiSut2d9vx8Pd6g01QDuulnMatiDDqBfv83UrQPORq7f9Rdl8mXwdSu3TcHIQ5aLxn2rQ3vg4OQ_rDthwh1Q7t5_u6uygLlT92nEmw1mtjrzvVaVTvstwZv8RZLD9TpO2MpPriicth4UK9uNfgzZcF8zVI2uR6GWYHf73kCqMd3Wn43_oGBs0GxqTVt-jGVPMuT6C1k0Wzk7Y1x7Psf6zTuQABlSKniQw0XXdRLdx4de4YR_yQK8fLBqNWktnHOslKPut0tg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลการใช้น้ำของนิคมอุตสาหกรรม (รายงานสถิติ).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลการใช้น้ำของนิคมอุตสาหกรรม (รายงานสถิติ)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=337
// @Summary 		ข้อมูลสถิติผลกระทบจากภัยพิบัติด้านต่างๆ เช่น อุทกภัยและภัยแล้ง
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_337 successful operation
type Metadata_337 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`22`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`113`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=OWgCYpBoHy9dHKZ0eP_6fvbE6FS0U1-DWOJ1rt7sJCs7MGNC7j5mq_SP5btKl-6u1kGV_XkVd_FL9u9l5iaAsYmbW9FF1bNCxgRobwj7p1s1EUceqSBnUY8P8F5r3S8y5d31BVVW_5IXbg1osXMkSqfr9nPPNcrXfbLMJctdit7ksp4tFYmkYxfn6ve-xqlU2vlpQOtFmv3uZ_AVO4VUXsO4Hasu4sNCrv7Rg67femDF9JPlhXgYr5hh9N2B6F8TDOcTqJtmymC6T2M9_Ia6zGdcLiaWVfHn_M3l8YPWeu8URW2riqvwfxwCHaVlBxl7-AU_WGwIR9B4QfVgymcQxA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลสถิติผลกระทบจากภัยพิบัติด้านต่างๆ เช่น อุทกภัยและภัยแล้ง.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลสถิติผลกระทบจากภัยพิบัติด้านต่างๆ เช่น อุทกภัยและภัยแล้ง`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=340
// @Summary 		ข้อมูลพื้นที่นิคมอุตสาหกรรมที่กำกับดูแล (แผนที่)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_340 successful operation
type Metadata_340 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`22`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`114`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=sSTRcrIC8IUkFVPw6kR3uXdLO5WRUlwrQPPXs6kGA2Bp_vBvlajPCKijEn030_XgLnxMgXBkO3DhVLKD-lLqIrwIWGIikgNMLBWlo555Xbgep5qtrVAMefPY-mEu05Y7E-FfK1lguuqtXGy70E9ubeyCo55nXjggcTvkmux4q4BXRIaNXIM6K50SgBLwXBd5sPr7wGlIQaEQecQZRDjN2RsgWZEOZvpqvHsuubLtWM8FMf1yCEpypG1G-vB6pCQvaXbAisxB-v_r9eR-SSGDJg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลพื้นที่นิคมอุตสาหกรรมที่กำกับดูแล (แผนที่).pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลพื้นที่นิคมอุตสาหกรรมที่กำกับดูแล (แผนที่)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=341
// @Summary 		ข้อมูลจากสถานีตรวจวัดคุณภาพน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_341 successful operation
type Metadata_341 struct {
	WATERQUALITY_ID           int64            `json:"waterquality_id"`           // example:`117` รหัสสถานีตรวจวัดคุณภาพน้ำอัตโนมัติ
	WATERQUALITY_DATETIME     string           `json:"waterquality_datetime"`     // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าคุณภาพน้ำอัตโนมัติ
	WATERQUALITY_TEMP         float64          `json:"waterquality_temp"`         // example:`30.22` อุณหภูมิน้ำ หน่วย ?C
	WATERQUALITY_CONDUCTIVITY float64          `json:"waterquality_conductivity"` // example:`232` ความนำไฟฟ้าในน้ำ หน่วย uS/cm ชื่อเต็ม The Electrical Conductivity (ec)
	WATERQUALITY_TDS          float64          `json:"waterquality_tds"`          // example:`130` ค่า tds
	WATERQUALITY_SALINITY     float64          `json:"waterquality_salinity"`     // example:`0.09` ค่าความเค็ม
	WATERQUALITY_DO           float64          `json:"waterquality_do"`           // example:`0` ออกซิเจนละลายในน้ำ หน่วย mg/l
	WATERQUALITY_PH           float64          `json:"waterquality_ph"`           // example:`4` ความเป็นกรด-ด่าง
	WATERQUALITY_CHLOROPHYLL  float64          `json:"waterquality_chlorophyll"`  // example:`6` คลอโรฟิลด์
	WATERQUALITY_TURBID       float64          `json:"waterquality_turbid"`       // example:`0` ค่าความขุ่นในน้ำ หน่วย NTU
	QC_STATUS                 *json.RawMessage `json:"qc_status"`                 // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=364
// @Summary 		แผนพัฒนาเศรษฐกิจและสังคม
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_364 successful operation
type Metadata_364 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`26`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`60`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-14T21:27:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=2dysggoWwwvwgoIBtA_fXLsMyrZZoeeTftigAgybfs6ozq0_vSiIUSjDpVyK35r-317ZmuUXR14jphorEGmONHy9AzC7Qg3ycIZmXbKpOZ7KnHJpH_z4RwPewJxinqXnoo79trke8zYKRr7ohZy3H1ohBnCr8Slv6eWa8iPtMMQKjmx2huvaACQx4qfujS4FRaIWxdsv33ZfWVbvBbyrRg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนพัฒนาฯ ฉบับที่10.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`nesdb-ecosocial-plan`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=366
// @Summary 		ข้อมูลของโครงการที่เกี่ยวข้องกับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_366 successful operation
type Metadata_366 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`27`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`168`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-14T21:51:06+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=co2DA-E53vAg0YzauxsS8vKa1tMNgF8CBr6sSDA10CXXhiaFzyziEI_fgssJktVE8X9R8Iu3PNVQ3gGYAXD989p5uev_CBA37vhpq-77Tff3aQX715eTZ2LsuGCjSXsKlIIIbr5AYIbpWhL1GG-cbf3XwVVG9oZk3A2aCWv700mt87LSNh5PeiMEGhS-Gh9K`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`bugdet_2558(phijit-nokonsawan-chaiyaphum-buri).xlsx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`bb-water_project`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=370
// @Summary 		ข้อมูลการติดตั้งโครงข่าย GIN ภายใต้โครงการ NHC
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_370 successful operation
type Metadata_370 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`31`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`169`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-06-19T13:44:52+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=rd5pGyryDd7AkKo4sYuSF3Pfq-1J73SftH3BYineofq1GZ1HQMGUoRqAH0Qq85mk-CMpQDQKJCF346qlmiTAEh4lKuIGgk0JkptClRJUeRwR_dvSVwJm4ZZULzwfG4Qo8e_frllza2eOX6G-it4KuCwLwYJOj0236bvktHk9s_1QThdKtfMHR8fgrSn1jEM61lJZ4-ZbsDf4MMZ8XiovJQhd5n0KA6JQTQ5bb4QQc9wUojO7h21dDhh1E_3WFmtdLqoWveMBrfvaTW6E60vvub9iCIN1cEyRdxZFHhUjv78frD0prJ_Sv91JDkDTy4XKWVIBikUF5A1djc7W7t4Wd39dO9fgidNmGDgAJip2FoI`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`หน่วยงานคลังข้อมูลน้ำและภูมิอากาศแห่งชาติ รายชื่อ 34 หน่วยงาน (1).xlsx`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ega-ginservice`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=406
// @Summary 		ข้อมูลอุณหภูมิ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_406 successful operation
type Metadata_406 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	TEMP_DATETIME   string           `json:"temp_datetime"`   // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าอุณหภูมิ record date
	TEMP_VALUE      float64          `json:"temp_value"`      // example:`26.46` ค่าอุณหภูมิ temperature value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=407
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ลม 5 km ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_407 successful operation
type Metadata_407 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`72`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=YLqEH7CijttfAKeKRC034V5aqijESyHLKOxblI9C7g1-ckbf2ccH-CZq1vqx-MD3td8Wmqvv5gFul34cU1DKrj-vxm5feG3BEXnqq0lbG8SiaeiEoMzdnW2vcVHej0mndsNBlK7i_OOg4uyR2iqM-1dOWWG3H2-iUppk33zp7nI`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_5.0km_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Vertical Wind at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=408
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ความกดอากาศและลมที่ระดับความสูง 0.6 km
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_408 successful operation
type Metadata_408 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`70`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=UPFxsbgs2wdUP654lKPmr6SU2B2ACKpTjCNxMR75RH1Z6XFvFLNW0VNw_yIW9eSJsNSon4OMDEUUjzGAeHcJi-eNBDM47dcbipGP-AH77K-hDTRkdDpzPcp5VH3kX7IKA5sQ4d7-xTfX6RmZpSoWoq1fRbI_Cu9xW346hEco2gs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_0.6km_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction and Atmospheric Pressure at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=409
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ความกดอากาศและลม ที่ระดับความสูง 1.5 km
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_409 successful operation
type Metadata_409 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`71`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=vsvjpfh7vHS1GqyIDFIREdxePQhkwsFkWvxWFrxXL6ndgRvbRGIucsnp4hHgKjwlQ5IY7kg1tthla9dDbROk8JRGOCjZSPa-TFpe3ynYPImfWoL-BpDUxX6lEr6w8awluKfB33P-YVeuDfiAefusUbXNdPkTdnK-DVJEkMq7uzs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_1.5km_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction and Atmospheric Pressure at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=410
// @Summary 		ภาพเส้นทางพายุ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_410 successful operation
type Metadata_410 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`41`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`62`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T08:59:52+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=YFyHTrUncIybWi9DbAjHAn_1GsHh2RlCBUvlfgi2HVHvKw5Y9igsAxsjA-3nWqpUd9Ro5aD6xxcgzO2m43n6j7rJhwX7hYHpRQVbEhI_yBooIhLptTwTIkAyCO_0cRTnhRL_RPgroVEyg2s8qRswqw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`I.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`UCL strom map`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=411
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ฝนจากแบบจำลองสภาพอากาศ WRF-ROMS Model, Thailand ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_411 successful operation
type Metadata_411 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`80`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=OktQFpjZ8AqbS4Qtj5sBaLgtBrYZSuSgBueMHcst7x6rLQgAzQbWGanF_MqdHuiWLthuKikRGWHDaDbO_1bqL5cIkPPOjf9imfeAgbUikx3dkPS-1TQGEfXapuH5zXQiWX4LXwUnjaLdtgwHgtEOAn9kkKzUtvRZrVNcgu49WjA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ani_d03_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Coupled Model (WRF-ROMS), Precipitation Map, Thailand(3kmx3km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=412
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ฝนจากแบบจำลองสภาพอากาศ WRF-ROMS Model, Southeast Asia ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_412 successful operation
type Metadata_412 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`81`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=cDpFoQU019a0ktDgJq1RARSQrUH-dMs3gXAGnj7MAIMoKAzUiVHtasG5fuMGJ4ssh0korNpWLcPnbzFKZlh6Kb4dRUnVeztXYjT0k-fpblIQVX4XdYxmfG8VeayfWdKqgfF_be2XtWmZ5VzLaXa-ZeS1drc7np9dHBB1Vypo8Zg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ani_d02_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Coupled Model (WRF-ROMS), Precipitation Map, SoutheastAsia(9kmx9km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=413
// @Summary 		ข้อมูลภาพเคลื่อนไหวคาดการณ์ฝนจากแบบจำลองสภาพอากาศ WRF-ROMS Model,  Asia ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_413 successful operation
type Metadata_413 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`82`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=iDnpV-xSft64lnjMSKlyi5UA3gZ_LGwFyRSo6D6ia51n-NzRWvzf7BRbxT_yV49nvg0PZ9dxHMBDRS3598hMfSmZFmhWTVWmyf0PAhW77cj1bQ2ooysuvdbPhasKtgVzrHXH-wmxl1wCjqlUkuGUiMM4mO9HY8bI_0QHKm1yzEs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ani_d01_large.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Coupled Model (WRF-ROMS), Precipitation Map, Asia(27kmx27km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=414
// @Summary 		ภาพเคลื่อนไหวคาดการณ์ความสูงคลื่น ( SWAN Model ) ล่าสุด
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_414 successful operation
type Metadata_414 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`83`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T11:11:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Rcf8gFTaje0VLtD-pDMGuPWpQQ48UIXFLoCeg41CPJ7MkVuE394mrrOkLgMKY74BbGDWFHZZDqv3wZVmTqwETjUPgUhD2Mb9xxNJ1CWiw28MrmpXg0Oy_Ku_UghqZ-3CVBl7F3MgrrCAs05gO0bQXg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`wave_168hr.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`THAILAND 168 HRS FORECASTED WAVE HEIGHT WITH DIRECTION (WRF-SWAN MODEL)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=415
// @Summary 		แผนภาพการเปลี่ยนแปลงของอุณหภูมิผิวน้ำทะเลรายปี
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_415 successful operation
type Metadata_415 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`14`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-01T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=fL72Yhn0BFiDegz0S57CXLbEiWNpVrO4gBPnbcDmrklVoY4BX4KSVn6ySicnknmaQiG4AtWzKGG8EuestBR4XwKXPKl_jzZ9GS43sktpx5poLj-ji-Mm1xnXEbuL9m03IIQyqOJ3tMUhgWLxJ42bMKLcwRM8qW4HKI0JkqC_hec`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ssh_diff_20171001_20171008.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Change in weekly Sea Surface Height Anomalies 2017`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=416
// @Summary 		แผนที่แสดงการกระจายตัวอุณหภูมิรายวัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_416 successful operation
type Metadata_416 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`2`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T08:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=r1LqwufPQ6CquPMz9Ign-Qkjx6j-1P5pAsf9-XFktUx32utXOSy8hyMXIOZMj8OMiNxW7bOaq_UYyogniKt6ySDWDWOOs8jf9bq3SjUxL9nGfpeqnROPaTmX3A9T4aZlcOpQkS8xBfYHLic1xWyGFMx2b61TTCeHeHhF8i8g83YHTOTdm7HdGNDvK2TUWD2d`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`hatempY2018M08D17T08.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Contour Images of Temperature`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=417
// @Summary 		แผนที่แสดงการกระจายตัวความชื้นรายวัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_417 successful operation
type Metadata_417 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`1`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T08:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=nhfmrBmmaOliqUTp3Bu21HdNyym82z1SwmEWBJMsFkD8roFzUzJt8gJU6bD2HC67sv6DkBU5W4knEB7952s7Q28G-KiVIRdz8h8I3tIsyo4MZFQcgDzNbq_6k1nWDN_UPdkNpmCk5oGhohc8Ms7MXLG6ruRZuaHE8hvDvTYxuoxWhR_btLPOz1LZ9PdYc30L`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`hahumidY2018M08D17T08.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Contour Images of Humidity`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=418
// @Summary 		แผนที่แสดงการกระจายตัวความกดอากาศรายวัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_418 successful operation
type Metadata_418 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`3`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T08:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=NC2jmiNw7ank22SsWSQvnbjVepItQ29ZBB8gadqFgn27NF3iBu5LD5K9Gqct5icAF0oRqZy7-xC2uGCcUYi6G48Gw5nvCJ-BhqjxZQZenyl2F5si-YmbHJSp2dnPgEatsSaKjnOiSrgqiJ3cdb1PHveXH38uXgLO1b8f46-ZRDY4O--NSulHWo_dZV6r4xXZ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`hapressY2018M08D17T08.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Contour Images of Pressure`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=419
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่ประเทศไทยรายวันเวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_419 successful operation
type Metadata_419 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`17`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Q4294SZtiCvzIrlxUKFe1FniwZJ-GUG8wMqE6_SnAO6q2w5Dl6FBOZYe1rM3wx0_gDD732MgUb0wmv3f707UvTd5fHhmTlnBwSPN5OJ_Rk6y6ZLWIG_OR5w-Ni_t9bxMweOn1Yc7VxT0PH9Fvb4RAA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`d03_day03.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand (3x3)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=420
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่ประเทศไทยรายวันเวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_420 successful operation
type Metadata_420 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`17`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=NDiw4PeO3NV9PGGk8BBO53k5O7-2vvEL0Tc8IddATwxjKI1D3NaBVglGKnz7C8vST8vj1BHfAiXqnnN6hkepMRI4pqs_USX3t4DO_mthk16qAkYauyVHoByOVtDQ9gFW40CzmMUzeD53e52WivgdIg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`d03_day03.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand (3x3)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=421
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่เอเชียตะวันออกเฉียงใต้รายวันเวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_421 successful operation
type Metadata_421 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`19`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-23T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=uId90XeyrichggirvByvUrIix1jOu7y4angSOQeu7WIqmaWs0fO-QKqJNujYfj2sviODze8Ai0MopkSI5XJkWYrjWFOmhHW2ZQx2tzjcDGFZmndM0qBHjbA_ogM4FV8l6chSZ-6V5fZEma3NB0eWRns2f-QIUYax27ziHGFtqXDC0qhTV9K1OfStriUtWEz0yVTFHCIq7vXHPBtm_9SpfQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161207.d02.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Southeast Asia (9km x 9km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=422
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่เอเชียตะวันออกเฉียงใต้รายวันเวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_422 successful operation
type Metadata_422 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`19`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-23T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=uAprqW5jeB6jeYhnF7-ePVs-Ck8FcLN3vmmDMqV0XBaR74rDrm9_htHQ8ev-lQ7ziR0_-8lmhjvLzF8DJhr5ESIRZUkKe2YsrqzfApZ3BtuYYpdvG9ahhNXd-faflfQujRlFTb8MHhYfpQY-PAgW7J2wXuunrKoZ_yhgYfPR-A8GuAWiiCKPAs6izYn93uo3irh3H8w6eq9BnFmduaSIPA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161207.d02.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Southeast Asia (9km x 9km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=423
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่เอเชียรายวันเวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_423 successful operation
type Metadata_423 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`81`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=C1xIqryIWj3c9bxf9LioFKgJLMePeGklFIAVKVEZVwY9qLahnmBUYd4ncJ5z9ZuCizW1X5AhoLZ8JnAlAXiQL0nOSEC_sUS7t1IczRK7ZWQdfqz5vuuHOTRNuCOHJ50gHSCXkkMar723DqFJ7stGZXePpFWSBXdZWxE4C8U3yvcDqw-dWWisCqje-yDdVbAIJ_tt6UL8JGd20uAgSO65qA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161203.d03_basin.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand Basin (3km x 3km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=424
// @Summary 		แผนที่คาดการณ์ฝนพื้นที่เอเชียรายวันเวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_424 successful operation
type Metadata_424 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`81`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=_0RCogMLrhiPygWoXTYlFJkCcLY_REsQTc9BmmjGwCZgG9Mi61LmyXLNzi1UW1KCEXH8sOHhE6baq22tYVa2w0huwBPUkFgZlf4LWzxQV8Zj38sokXSgJ-T-WcpwIoF6pLavl3HGIPkMTtsCFKlldBPrCOY26K4bjX2snd99KU9IJYDWH7PL2TaKEY47cyOle2BCLuYlMR4Qh9k20zx-wQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161203.d03_basin.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand Basin (3km x 3km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=425
// @Summary 		แผนที่คาดการณ์ฝนลุ่มน้ำในประเทศไทยรายวันเวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_425 successful operation
type Metadata_425 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`81`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=9qdBoJMQ9yWalHbUfWnmfAttK1qnpL5puDQEHREXaU6ROFuFjOiQToWEKEgXVUuqQoeANo1HyvZpLlt7q5pd15STni48A98qZynWhJZLpGf_C-UQ9w6aKjKWhhq-O03W47OKKMeibv5FLUrk63qoTivZKg5Rm4k3nRhZtFw3HxJ91tz6-UwiEPbnQnMhM3HBGbwxqNteRil66Qpwv7flbQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161203.d03_basin.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand Basin (3km x 3km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=426
// @Summary 		แผนที่คาดการณ์ฝนลุ่มน้ำในประเทศไทยรายวันเวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_426 successful operation
type Metadata_426 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`81`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-19T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=VKfYIJURGisXT2gy4hj03UiwAw3zLOw0eUj-8bGNTyarZDsOyP8QlxIMd3MscqcDyD2Qimv0oeCopOGw9Yw5PVfyc4EF0c38x2cbiiOng8Sbc_MgP5e8xOgYjRux7IvEC6ErDgwsmvHcg_ii276ObtUmqhqWHWK0U-bRoq2_jj4axO8umVcGyUzjSLPtaT1FHgs_DyMl3USSO3xJmENCaQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`rain_init_201808161203.d03_basin.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS Model, Thailand Basin (3km x 3km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=427
// @Summary 		ข้อมูสถานีคาดการณ์น้ำท่วม ลุ่มน้ำเจ้าพระยา (CPY)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_427 successful operation
type Metadata_427 struct {
	ID                             int64            `json:"id"`                             // example:`9` รหัสสถานีสถานีคาดการณ์ระดับน้ำท่วม
	FLOODFORECAST_STATION_OLDCODE  string           `json:"floodforecast_station_oldcode"`  // example:`NAN007` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station's serial number
	FLOODFORECAST_STATION_LAT      float64          `json:"floodforecast_station_lat"`      // example:`16.270330` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	FLOODFORECAST_STATION_LONG     float64          `json:"floodforecast_station_long"`     // example:`100.413960` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	FLOODFORECAST_STATION_TYPE     string           `json:"floodforecast_station_type"`     // example:`ระดับน้ำ` ชนิดของโทรมาตร (เช่น ระดับน้ำ) station type (water level,discharge)
	AGENCY_ID                      int64            `json:"agency_id"`                      // example:`9` รหัสหน่วยงาน agency's serial number
	PROVINCE_NAME                  *json.RawMessage `json:"province_name"`                  // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME                    *json.RawMessage `json:"amphoe_name"`                    // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME                    *json.RawMessage `json:"tumbon_name"`                    // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	SUBBASIN_ID                    int64            `json:"subbasin_id"`                    // example:`1` รหัสลุ่มน้ำสาขา subbasin's serial number
	FLOODFORECAST_STATION_WARNING  float64          `json:"floodforecast_station_warning"`  // example:`33.75` ระดับเตือนภัยแบบ warning
	FLOODFORECAST_STATION_ALARM    float64          `json:"floodforecast_station_alarm"`    // example:`33.15` ระดับเตือนภัยแบบ alarm warning
	FLOODFORECAST_STATION_UNIT     string           `json:"floodforecast_station_unit"`     // example:`ม.รทก.` หน่วยของข้อมูล (data unit)
	FLOODFORECAST_STATION_CRITICAL float64          `json:"floodforecast_station_critical"` // example:`34.34` ระดับเตือนภัยแบบวิกฤติ critical warning
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=428
// @Summary 		ข้อมูสถานีคาดการณ์น้ำท่วม ลุ่มน้ำชี-มูล (ESAN)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_428 successful operation
type Metadata_428 struct {
	ID                             int64            `json:"id"`                             // example:`9` รหัสสถานีสถานีคาดการณ์ระดับน้ำท่วม
	FLOODFORECAST_STATION_OLDCODE  string           `json:"floodforecast_station_oldcode"`  // example:`NAN007` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station's serial number
	FLOODFORECAST_STATION_LAT      float64          `json:"floodforecast_station_lat"`      // example:`16.270330` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	FLOODFORECAST_STATION_LONG     float64          `json:"floodforecast_station_long"`     // example:`100.413960` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	FLOODFORECAST_STATION_TYPE     string           `json:"floodforecast_station_type"`     // example:`ระดับน้ำ` ชนิดของโทรมาตร (เช่น ระดับน้ำ) station type (water level,discharge)
	AGENCY_ID                      int64            `json:"agency_id"`                      // example:`9` รหัสหน่วยงาน agency's serial number
	PROVINCE_NAME                  *json.RawMessage `json:"province_name"`                  // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME                    *json.RawMessage `json:"amphoe_name"`                    // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME                    *json.RawMessage `json:"tumbon_name"`                    // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	SUBBASIN_ID                    int64            `json:"subbasin_id"`                    // example:`1` รหัสลุ่มน้ำสาขา subbasin's serial number
	FLOODFORECAST_STATION_WARNING  float64          `json:"floodforecast_station_warning"`  // example:`33.75` ระดับเตือนภัยแบบ warning
	FLOODFORECAST_STATION_ALARM    float64          `json:"floodforecast_station_alarm"`    // example:`33.15` ระดับเตือนภัยแบบ alarm warning
	FLOODFORECAST_STATION_UNIT     string           `json:"floodforecast_station_unit"`     // example:`ม.รทก.` หน่วยของข้อมูล (data unit)
	FLOODFORECAST_STATION_CRITICAL float64          `json:"floodforecast_station_critical"` // example:`34.34` ระดับเตือนภัยแบบวิกฤติ critical warning
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=429
// @Summary 		ข้อมูสถานีคาดการณ์น้ำท่วม ลุ่มน้ำภาคตะวันออก (EAST)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_429 successful operation
type Metadata_429 struct {
	ID                             int64            `json:"id"`                             // example:`9` รหัสสถานีสถานีคาดการณ์ระดับน้ำท่วม
	FLOODFORECAST_STATION_OLDCODE  string           `json:"floodforecast_station_oldcode"`  // example:`NAN007` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station's serial number
	FLOODFORECAST_STATION_LAT      float64          `json:"floodforecast_station_lat"`      // example:`16.270330` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	FLOODFORECAST_STATION_LONG     float64          `json:"floodforecast_station_long"`     // example:`100.413960` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	FLOODFORECAST_STATION_TYPE     string           `json:"floodforecast_station_type"`     // example:`ระดับน้ำ` ชนิดของโทรมาตร (เช่น ระดับน้ำ) station type (water level,discharge)
	AGENCY_ID                      int64            `json:"agency_id"`                      // example:`9` รหัสหน่วยงาน agency's serial number
	PROVINCE_NAME                  *json.RawMessage `json:"province_name"`                  // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME                    *json.RawMessage `json:"amphoe_name"`                    // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME                    *json.RawMessage `json:"tumbon_name"`                    // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	SUBBASIN_ID                    int64            `json:"subbasin_id"`                    // example:`1` รหัสลุ่มน้ำสาขา subbasin's serial number
	FLOODFORECAST_STATION_WARNING  float64          `json:"floodforecast_station_warning"`  // example:`33.75` ระดับเตือนภัยแบบ warning
	FLOODFORECAST_STATION_ALARM    float64          `json:"floodforecast_station_alarm"`    // example:`33.15` ระดับเตือนภัยแบบ alarm warning
	FLOODFORECAST_STATION_UNIT     string           `json:"floodforecast_station_unit"`     // example:`ม.รทก.` หน่วยของข้อมูล (data unit)
	FLOODFORECAST_STATION_CRITICAL float64          `json:"floodforecast_station_critical"` // example:`34.34` ระดับเตือนภัยแบบวิกฤติ critical warning
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=430
// @Summary 		ภาพคาดการณ์ความสูงคลื่น ( SWAN Model ) รายวัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_430 successful operation
type Metadata_430 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`13`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-22T22:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=cebbAACSC27HiKBWXfPxoSTROeHhP24sBZ6sT6UT4ME4yJucjvBx1KUUUk9hfMqxH4SlWE68c9DrTUF4mEI9GIc4kinS5VKLTbQE1CjAqdqVxNFTw--7suM-EOJoECSN6VHa5iDzKvdpgWbmnX1lAq4bO4V5pOH8rvsa3eEEFWAM5OD-ApEb2RJ0_zsc3I2B`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`wave_ini_201808151900_20180822_220000.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`THAILAND 168 HRS FORECASTED WAVE HEIGHT WITH DIRECTION (WRF-SWAN MODEL)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=431
// @Summary 		ข้อมูลสถานีคาดการณ์ความสูงคลื่น ( SWAN Model)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_431 successful operation
type Metadata_431 struct {
	ID            int64            `json:"id"`            // example:`9` รหัสสถานีสถานีคาดการณ์ความสูงคลื่น
	AGENCY_ID     int64            `json:"agency_id"`     // example:`9` รหัสหน่วยงาน agency's serial number
	PROVINCE_NAME *json.RawMessage `json:"province_name"` // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME   *json.RawMessage `json:"amphoe_name"`   // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME   *json.RawMessage `json:"tumbon_name"`   // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	SWAN_NAME     *json.RawMessage `json:"swan_name"`     // example:`{"en":"Pattaya"}` ชื่อสถานีโทรมาตร tele station's name
	SWAN_LAT      float64          `json:"swan_lat"`      // example:`12.871000` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	SWAN_LONG     float64          `json:"swan_long"`     // example:`100.844000` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=432
// @Summary 		ข้อมูลคาดการณ์ความสูงคลื่น SWAN Model
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_432 successful operation
type Metadata_432 struct {
	SWAN_STATION_ID     int64   `json:"swan_station_id"`     // example:`1` รหัสสถานีคาดการณ์ความสูงคลื่น
	SWAN_DATETIME       string  `json:"swan_datetime"`       // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์ความสูงคลื่น
	SWAN_DEPTH          float64 `json:"swan_depth"`          // example:`2.3737` ข้อมูลคาดการณ์ความลึกของคลื่น หน่วย m
	SWAN_HIGHSIG        float64 `json:"swan_highsig"`        // example:`0.10686` ข้อมูลคาดการณ์ความสูงของคลื่น หน่วย m
	SWAN_DIRECTION      float64 `json:"swan_direction"`      // example:`88.637` ข้อมูลคาดการณ์ทิศทางของคลื่น หน่วย degree
	SWAN_PERIOD_TOP     float64 `json:"swan_period_top"`     // example:`3.3597` ข้อมูลคาดการณ์คาบคลื่นสูงสุด หน่วย sec
	SWAN_PERIOD_AVERAGE float64 `json:"swan_period_average"` // example:`2.6946` ข้อมูลคาดการณ์คาบคลื่นเฉลี่ย หน่วย sec
	SWAN_WINDX          float64 `json:"swan_windx"`          // example:`1.0075` ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศตะวันออกและทิศตะวันตก หน่วย m/s
	SWAN_WINDY          float64 `json:"swan_windy"`          // example:`0.3487` ข้อมูลคาดการณ์เวคเตอร์ลมในแนวทิศเเหนือและใต้ หน่วย m/s
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=433
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_433 successful operation
type Metadata_433 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=GJ0QPeEno24e6-eR89aEZKwOafgWQceAf6ArIJx8I3vv5V7dSSUXkJMCUW6JQVepUnc4jjx69ydug6k_v20ucixC3HtR2rW_iUruDL2Om_5vd2-NNPHAChC3iHqPWP0hWvPMu20nI3hxTRnZ7pCvKSbB7YLRTWBlOaBk7e3x_Fmmk9Cy5ett6wCr2-VI4Lf5`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=434
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_434 successful operation
type Metadata_434 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=S7EZhnCqNH81QP-1AFepXaEOezK1ivM0__GyRfYK6bOsvxQ5niSB5fHRuKe4frsPwJHsBPvUDMKEw1QssvojLX7CB5EvDmZzP1YMPh3drxCggyXu4HTJSIGkOJnMEOZE-Fjqn39-EJt3jK06zgGt4YbHWuUTJlQQ9fek6PR6P-077bZ1qPNa5p0J_sI6bh1N`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=435
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Southeast Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_435 successful operation
type Metadata_435 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=FFp9m8wOgibonss_anJi9GxkNPHG-VCWLWPQR1QQK75iw0jjC0tW7L2WFg7fp57hcXZh6nXN8mNo7r_eeWysjOb1eRlS5TFbJEytGsQKag-JUbKKcHVPqrZ8BYAseqTAS1N2pRuwYPMz0EVsrD0bJ7gW9Eavv7RG4EQgf6xEZSPgEr5qccNml3CRAyTauTJ8`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=436
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Southeast Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_436 successful operation
type Metadata_436 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=ut38j-LkTtwKGOZGZAB4UokQhlg6iLHlTPBE6qzEj8GRglLIrV_2EuGM5NyTsRAUBcjnYTM6WjsLsmlQlH7Xg2efzXio_7MW5Y6P82xr28FfIuX1ZbtY1NUEG6Y7oKoI6NAMmSBwGMkpH_aq4JfLcTxvGKpq1oq8xxufm4pBZBg9aU_4eXi3PyhT8mN8Db8z`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=437
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Thailand รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_437 successful operation
type Metadata_437 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=RKP9TFJFcyN_Un4Q-EmzBc-bKYUbRl3U4ZO9dGTx1oqyPgX0V5-d4XcSvyEe7ngpiaq4YFZp6DtGa4t7_1cUkc1gDjiobFzhBxayjpOI_9RCAa98bP0Iq6op55CuhbJc8m8w2OzBBt7F3p6E4k4AyXUQ_VzEqmgzAkX3rwhFzIytup59HH6nGs7Zta6tWflJ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=438
// @Summary 		ข้อมูลภาพคาดการณ์ลม 5 km Thailand รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_438 successful operation
type Metadata_438 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`12`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-01T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=HyWhfm7q_RF-4n_W2H-VWWeQ2UEXQvkQhlGLnXZ6hlUvZZ3afAGq-HEQ0kzMqmeoC_-4-hjXQvEvhzmpyotFpbbQENv8-pCTWtLloAIw8vCxcREV0eD1nUIDtm0JvTairIccuN5ENDU9Av1JMhc9D6BUrtbjKotuPOTqKw-_n-E-YHS48SgrbiAeg8jEbB8m`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_5.0km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 5.0km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=439
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_439 successful operation
type Metadata_439 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Bsl6Xh-xUxHByU0LlSgJT7-FFkAW_jT1Nhn5TneJytFQvrYh_d2UNJVS3ELA4I5AH9wlf7Ihf2Psi_-3mQqHNjINcA9qBBXI9Zgk1uRdZlLhUf2cNGi0FdWhIQACtpBm6HssKCMVgeFyLrnvOB0zpfuxb9kUPTuik0pkYe2P6LA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=440
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_440 successful operation
type Metadata_440 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=mjdg2dai4t3J5TieDSZoBBA6HiwrLT9z2h_JPO3T9xavVYmKeNqM08lCWLqb4hFaSZa8Ty6L254HMxFBpuG6_4D1wuQL1t1TARen3NIRNN_WZbBZVZt1JDpUN63JbOy1lF8XIil0xdJBkO0SwU9IOFdMB7Y9QDduIGc8IcGRFtI`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=441
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Southeast Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_441 successful operation
type Metadata_441 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=CL4vwn1fjEoJ_IMDqK4wyeeVy0ZvMV9T7KHZzZXok8nwWjxVTRHVtjrr7nEjWY4VXLipKok9sp1_NKp7xaaEkWT9aWj8xOf64NdCGaAP5JyqKXiNQvF-WMkfrT71la74Gbv4CHldg4lVyEJcx0G1oYJwsau8PEE4V-kyPRdseGs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=442
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Southeast Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_442 successful operation
type Metadata_442 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=9uPtpw4_VtW8vq8eLLAKB1Xu0HXKAfQ8zJ1Nf2FTlpW3TOYTpzT1e5vtODxHkVnsT0G24OEgxosXHdcyQQwTATgP53tFPvX87lUlRVAyMuQztlCa5Ztf10yHbtJZf3Gs8NArEO7J8IwaPZbcvU7Tu6RhNfFOfBUbC1HTHwDvJis`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=443
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Thailand รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_443 successful operation
type Metadata_443 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=adDNSTJOOTmhexNayMuhMotOVUS63bhw1wlo9j2qSHbww46TT3F0iUkgDGxgGX3wi15HV9hfaFuqvWhS-I1OPaa-UQ8YnRH5VZs6KjYWMn0OnROwXCiPZFQpT_JhRk_dbN8dSe6Wo2tcSzuoNON_YyQB7HYFKvxi-OnXaZJFTRg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=444
// @Summary 		ข้อมูลภาพคาดการณ์ลม 0.6 km Thailand รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_444 successful operation
type Metadata_444 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`10`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-06-22T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=S-NUYPrhJrgo2Pttq41jZAz1p9FaTQi0sqA07QTbUysnzF90V4zfrHWMUTKub2-kugiMN5vZzpjW-iipsJACeR9LlqWy0NV_SDMkXBqE6qeLm7rR82YXcrzpzTp4b-8q8WlO8a8gRXydLxrg2vT8Ti7b_BD9L06PXMrA9bs-2sg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 0.6km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=445
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_445 successful operation
type Metadata_445 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=hbGLNd1j7ooK4pTnBrBY9KA6qHWru01is-KCqHnIioPt4N6gSuuIp9U8uA88gGx-PikHqC1SicR8BRtbIAiX1n45o5dIslGYFE5oPomF93dClJuYNQECWUyk_NsQsWJsoZlAYHL-rPUN1kF_QZSGuw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=446
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_446 successful operation
type Metadata_446 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=LytkpNVdSKvBtgU8ITxMxzXrbIhwaJ6YgX7P2dqh9osmxQrLgHUiAKATAsa3OKECCOPN-NpXQsQFvND_sU6kTZe9we-f98Jhu9yKpYVNJeJVO3NyJu7I2GXMBjjRho6Ms_PitMj5geN6xI9elMGASw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=447
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Southeast Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_447 successful operation
type Metadata_447 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=9788-zEVNtEfQCxwwI2mdQXnqh6Sp9URmriIjUQ92MROk8oIiU9JfkeP8A6J1tAxzU0iaElXvr7GNH9NUM1NZH-bOeACvTL3GJjvMK33vu5otpFuKHIsfmptT8zKhRx8yFWOl2mi--tVT6WWX7BlQQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=448
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Southeast Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_448 successful operation
type Metadata_448 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=LmPyhltemAH1IMnaQq5p7a6Cs0fliTbaU7ESwbMF-o2TTlRIYcsW0kzZnayA35NvgcmIYw555U3rF7VNEANQ3WOXvv8ZmcNM0v2f4jigJljYZVMclHj9eZ85wXnZbSaS8TpfwwcwmEaeCCARkqmfWQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=449
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Thailand รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_449 successful operation
type Metadata_449 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Rp-pBWQM0C7x3gsrn2oZrpx0PyLCRWicfdKxu7zDTnjn21LvLn_Ewgfawg7-zS67HNoTLm-ddtF-f7GkhS9V_8YD7JRBfufOTEG_i2ncFZrhHd-AO7IfiB2qkqkwqY699kGuNY0eFuoAgpyfRAe7rA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=450
// @Summary 		ข้อมูลภาพคาดการณ์ลม 1.5 km Thailand รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_450 successful operation
type Metadata_450 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`11`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=hLzgSW8c3X4ESN4lzP9HaxMYQWB7YA-grz6CYQKgKO487F-OrM8RkCiZ6BtUIsMS34pR4RIdUTt-Yxy8vp83MXWUutZlLJZtUSQUo2oLxN7Bp21cuMbIqEFAhjZ75gyzVxPqKhT3kH-6QIqkQDF2Tw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_wind_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`WRF-ROMS, Wind Speed and Direction at 1.5km. above Sea Level`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=451
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_451 successful operation
type Metadata_451 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=fREqRMFT8IWWK-o7bZY25-OQz97sirsaHcXXCGIgOdT7MMR7ZcZlwKBqtu_a2rNxg2tNkWWk2S_9XrNBrLF6qBunaLLZDHYt60DDH1JirCtAKm2jDw5B_3YOyd86wtQUoI67BuK2frhgSMo5FBrGp4rS8EiO-4yxVdXMR0yv2SmoO3mUzomayj4cDDX4ODaW`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=452
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_452 successful operation
type Metadata_452 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=t-p2go0dRvoAoCLD5JlKna888pr4TggF2x35hPVT5O-WaI0yrt_S9hkXlSA-AT3sq48rS1qodugRip9j2QaYouHZEm-AZmxmOIK5gNwXFz1Cl8CHoYCXF98-NkrEe9doxCZF0r4Oa8OBNEwzdTYdSa2Dp_RQkRzaU00ZuhGRImgnp6WYAHGYhXZy0wQ68TWS`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=453
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Southeast Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_453 successful operation
type Metadata_453 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=qecmUSYc-tpqjelFtDIAZLJNDtKWN8sYmayKioNVZL406r79HX84I2ZQ7CzY7E0U0WPCosM5U8j1hV6gEPOHQjkNRXe0BX2KdLZsIaCugwrpjzAKMOPCYLSAPNLzuP-c-LylBu00GxSvPt3N5c9z1JkUAjgH7NmMLjEINehMj6Zl24SL0vNtHf1prBYx0mzQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=454
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Southeast Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_454 successful operation
type Metadata_454 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=u7Dnl6JcVEfh1Jnra-2PkMImLaLGHUajarB-LNQMncS7q72VqNCrU_QGPX0xImYUZbPAJHCkR0Q-p7np27LaL8VUGMj9JIO3InqDaKS-c5zLf46U2KIwJ1_PwLAB1KbonCzeUli-HG01k2moeSFV2krTwipHIYGFQV_zeSusbFepuMVSDIJOFIdDZ3cbe3nQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=455
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Thailand รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_455 successful operation
type Metadata_455 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=88MLpCwwC-YFA5DLvmaA2nFceBZ1G4TAXdD4Zg7e7gpPoorFFhlw4WnngsVlu8CPKpJAaF3RKznSkxSWtLAaP1q-8LT7YFHnyEqXmrYUcpwNS438TJdzzWNANqGJL27ySsn4JszMhUS1wo-sC2IRTbyKJjkc_fH384RA1_TBrRmCk-zYns6XoT7fVpKiU29w`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=456
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 0.6 km Thailand รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_456 successful operation
type Metadata_456 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`15`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=dP-iFg6Th6X_QVRZdWZz2cUNt20bjd6NxirejvFvcYsisxOHCd1heBeLDz5NbxYEDMGXdUt7WRCG0he-LUU9W1xSwx-A7IXCKxWka7pv9cqto5bavVFs_5CG6fozVXZ3ay-WIkukPNPo4p6y3Cs565WqNuDx03MYvqb8MljTcyNBSvmoOBSnbu-rYEgegZMu`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_0.6km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 0.6km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=458
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_458 successful operation
type Metadata_458 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=9OSIluswDlmU2Y8IuWaJljJGI1w9acl5OUPxQr444bclewWoa4evXaOGhsoIFGEBfXwXP34jFrvF-IeIhc6T47IfbB161M3yT0yKtJcwhkaRe4DcYamYykueOVTz1-SPalzLvMrTUcQsBr7rN1DYGVedqPxACCJ1LP8LeYpA7GrW6uGRpyESFW6IXva0ecKO`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=459
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_459 successful operation
type Metadata_459 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=ttej0_IMyPVKZKCZJ-jp2JQl31eRCBV5Q_mHPKp0Rj47-9zrHjDSB5jOvT-hfVgt-MA6toHedQmk7rPakALutIy4RhEKkFkTpkOQjfb_iKlpsykOwzkN6yptxKBq35hIaZXoIOK3xoUWBnQqV7q6zjXlMdCmdsuQJmlOqL8FZAeqSCbRyu9mhKsp40qVtxqY`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=460
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Southeast Asia รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_460 successful operation
type Metadata_460 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=3UG-u6_ycMKtkRirimwP1gwwUQID0eLaPeFMruUj5Yl-_N3ev0rzSAL3Tfx4u2LzjZTt8XW90ouddCPy59Zlv-hY-_XLDIblXuYqIdbiru6hdpXGbyY8RyDFNi5JX4VCTJTAJv38P1mOTJdF_kz9JM-mxXP5KxI8-A3l8F_fVrjO6a-KoHLCPpJLYjX5JgYV`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=461
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Southeast Asia รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_461 successful operation
type Metadata_461 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=GUkDRM-UmMIXvbhwLBjcKq4meBlNArR0o2DgDCyXJG5ZGQ8Wcphc6RdSj08PfQ-VIxDYB0fJXc1wL_FaxGFpVT8cr1JCZn4VY9DFwvc1GSN-qCgK4Ei6MJMUe0Z-BscY59Cs_gONvxll5bxHhhrcCMAtg3azQydhvGUMq02ixgvWWCJyhvopv1kFeAxjOIm4`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=462
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Thailand รายวัน เวลา 07:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_462 successful operation
type Metadata_462 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=H99vwmUlAk8B6yrIgLN-V6nuMLeE6M6VJ6jmsb2MVziRNzljOxxgGXL2_DxKhfzVDD9O59UMkcVk9ifU08bugs8cTSChSJY_y81URD-DCnCxqKMajGAunjNjVV5PfqbMOY636ZC_EDSe-Jg_XmP8Yr_9BW8gTzr470FbMxotXyBrwHqFB2l5RCsD5Xy17ta4`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=463
// @Summary 		ข้อมูลภาพคาดการณ์ความกดอากาศ 1.5 km Thailand รายวัน เวลา 19:00
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_463 successful operation
type Metadata_463 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`16`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-22T19:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=rohnkYlLhp3Jk_4OWaI2R-SVpJwVMfYCoBhB2pDlo5acWvI-HXu1_HRPnMtb9bzi9xXq_nlbe9z8BqJA7YSE5ZkzYLvuKLo_wS-_9pfbxrpD9Axb9Dg28TQLYafmPDNFDR-TqCxKR04Ps2RUdKs6In-S36CtOjwHOtGWFB0SKJgrc3ux7AOK-krraw2I0bdl`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`upper_pressure_1.5km.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Pressure Map 1.5km Forecasts by WRF-ROMS Model `  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=466
// @Summary 		คาดการณ์สมดุลน้ำ  นอกเขตชลประทาน สัปดาห์ที่ผ่านมา
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_466 successful operation
type Metadata_466 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`51`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T07:53:47+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=s-643VXdZPYPYucsPW25nTZvpMhQyIhqMA4vXBZ-K7ROLmLfENy3w9y9XhAwt2uVtgtGi9z1jPPllUnOa9chj5nBadZ-lpIHqP2o1a9SEKTb3_i3bZk2tMMr1blaWXXq2QOT4S7sKsytF5Iv0Zibdg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`we_bac.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพย้อนหลังสมดุลน้ำ นอกเขตชลประทาน รายสัปดาห์ ล่าสุด`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=467
// @Summary 		คาดการณ์สมดุลน้ำ นอกเขตชลประทาน พยากรณ์สัปดาห์ถัดไป
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_467 successful operation
type Metadata_467 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`57`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-13T07:55:10+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=HhU9lhz9OzUgpctve2oQmUxDBDqv6je-DrnF8dpmguqc8IYPlYl3Y2N308aQwwRPSpmL0EIiAcyEjGjeRpgPyJfT1J4KrPxx3JyB0dey6rf41_N8Fkjwef6nCc54GiYRptOBoOg2xjFJqGLdgxiFcQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`we_for.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพคาดการณ์สมดุลน้ำ นอกเขตชลประทาน รายสัปดาห์ ล่าสุด`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=468
// @Summary 		คาดการณ์สมดุลน้ำ นอกเขตชลประทาน รายเดือน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_468 successful operation
type Metadata_468 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`52`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-07-05T07:54:59+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=x2WAORfiY_61_Stv5hMG-WJat-g2cSkwzk3Cd1DkZs_V_eRemwc1cjxNSAfmsEDtoZd924FrXjI-i4ud5lD7ldtr1fHsYDY1zotIwoHB8VNAHqNVC8OZ3XVs90Y7kkM8`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`th_6.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`haii-swat-w-forecast`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=469
// @Summary 		ข้อมูลคาดการณ์ฝน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_469 successful operation
type Metadata_469 struct {
	PROVINCE_NAME          *json.RawMessage `json:"province_name"`          // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME            *json.RawMessage `json:"amphoe_name"`            // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME            *json.RawMessage `json:"tumbon_name"`            // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	RAINFORECAST_DATETIME  string           `json:"rainforecast_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่และเวลาที่เก็บข้อมูลคาดการณ์ฝน
	RAINFORECAST_VALUE     float64          `json:"rainforecast_value"`     // example:`2.3596` ข้อมูลคาดการณ์ฝน
	RAINFORECAST_LEVELTEXT string           `json:"rainforecast_leveltext"` // example:`ฝนตกเล็กน้อย` รายละเอียดเกณฑ์ของการคาดการณ์
	RAINFORECAST_LEVEL     string           `json:"rainforecast_level"`     // example:`2` เกณฑ์ของการคาดการณ์
	QC_STATUS              *json.RawMessage `json:"qc_status"`              // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
	AGENCY_ID              int64            `json:"agency_id"`              // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=470
// @Summary 		ข้อมูลภาพถ่ายดาวเทียมสภาพภูมิอากาศ Himawarii WV
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_470 successful operation
type Metadata_470 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`47`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T15:55:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=L9XWFJQnIlxHwMmVew8abbmaS0WvndItJu3tqTblQ9KOsekbP3kJTn-8ocOpL-yHjlWCBqcva5Ni-x79QbBMtU5y8kF_tTMdgkr5ohjvo3KpHKkwYLOY54UIGraVkPqDfOoJna2zS2abU7X5Pxx3Z_TP6uRyWOCS-NmI-ceE3gk`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`WV201708280900.JPG`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd himawari-wv`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=471
// @Summary 		ข้อมูลภาพถ่ายดาวเทียมสภาพภูมิอากาศ Himawarii VIS
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_471 successful operation
type Metadata_471 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`48`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T15:45:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=MYdqAby2KuiGJBKXBa6C4V7txrGnlGq0HgI-d1Wt3FZIIMjeSk7gVL1GH5rVjTWC6Uq7VqWHbHk51_Vq1kXz6FeDaFyeKCiDinh7rAQnS3uXJ48VC-t331iAPLrBCM1GRY8IaoXVGERzdXkj_uuowsBjeG231QLwVxpdiDsHg0E`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`VS201708280850.JPG`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd himawari-vis`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=472
// @Summary 		ข้อมูลภาพถ่ายดาวเทียมสภาพภูมิอากาศ Himawarii I4
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_472 successful operation
type Metadata_472 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`49`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T15:55:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=LQFFDzWQ3BB7NlQUMd482yPsZ1up14kcZ4B005EjfuhsKFj99OacqJkDkdbHp97m1D4QZn1WjFJD3qsrcQDOpq33Trdon8jZPWTVCcbeueKh3MBLBRI2WQxAgriY-fazoFWD07ZVWh_WAAi6oruOaWtVlKipgLOcttX-nLdF430`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`I4201708280900.JPG`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd himawari-i4`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=473
// @Summary 		ข้อมูลภาพถ่ายดาวเทียมสภาพภูมิอากาศ Himawarii S4
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_473 successful operation
type Metadata_473 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`13`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`63`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T15:55:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=0txBYqCwTeIOcwiKSHAHngH26t_QnZqVwWCOQ_BnvKpeRYVMNAaSXITOo-Ta_KkX-hf5F9Ws5u-zbekzApJpxfIVBHALuP3uCdw-Vc16fXFgokwgfmojXNh0tcARf3E42nG0WvUTn1VJnCRFFn15Yl8E54D0-cMd_7jhmaulXfs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`S4201708280900.JPG`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`tmd himawari-s4`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=475
// @Summary 		ฝน จากระบบโทรมาตรเขื่อนอุบลรัตน์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_475 successful operation
type Metadata_475 struct {
	TELE_STATION_ID   int64   `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string  `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL1H        float64 `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=476
// @Summary 		ฝน จากระบบโทรมาตรเขื่อนปากมูล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_476 successful operation
type Metadata_476 struct {
	TELE_STATION_ID   int64   `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string  `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL1H        float64 `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=477
// @Summary 		ฝน จากระบบโทรมาตรเขื่อนสิริกิติ์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_477 successful operation
type Metadata_477 struct {
	TELE_STATION_ID   int64   `json:"tele_station_id"`   // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME string  `json:"rainfall_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL1H        float64 `json:"rainfall1h"`        // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=478
// @Summary 		ระดับน้ำ จากระบบโทรมาตรเขื่อนสิริกิติ์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_478 successful operation
type Metadata_478 struct {
	TELE_STATION_ID     int64   `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string  `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64 `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=479
// @Summary 		ระดับน้ำ จากระบบโทรมาตรเขื่อนอุบลรัตน์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_479 successful operation
type Metadata_479 struct {
	TELE_STATION_ID     int64   `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string  `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64 `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=480
// @Summary 		ระดับน้ำ จากระบบโทรมาตรเขื่อนปากมูล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_480 successful operation
type Metadata_480 struct {
	TELE_STATION_ID     int64   `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_DATETIME string  `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	WATERLEVEL_MSL      float64 `json:"waterlevel_msl"`      // example:`22.549` ระดับน้ำ ม.รทก
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=481
// @Summary 		พื้นฐานโทรมาตรเขื่อนสิริกิติ์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_481 successful operation
type Metadata_481 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=482
// @Summary 		พื้นฐานโทรมาตรเขื่อนอุบลรัตน์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_482 successful operation
type Metadata_482 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=483
// @Summary 		พื้นฐานโทรมาตรเขื่อนปากมูล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_483 successful operation
type Metadata_483 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=484
// @Summary 		ข้อมูลความชื้นสัมพัทธ์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_484 successful operation
type Metadata_484 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	HUMID_DATETIME  string           `json:"humid_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความชื้นสัมพัทธ์ record date
	HUMID_VALUE     float64          `json:"humid_value"`     // example:`95.33` ค่าความชื้นสัมพัทธ์ humid value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=485
// @Summary 		ข้อมูลความกดอากาศ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_485 successful operation
type Metadata_485 struct {
	TELE_STATION_ID   int64            `json:"tele_station_id"`   // example:`899` รหัสสถานีโทรมาตร tele station's serial number
	PRESSURE_DATETIME string           `json:"pressure_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความกดอากาศ record date
	PRESSURE_VALUE    float64          `json:"pressure_value"`    // example:`1012.67` ค่าความกดอากาศ pressure value
	QC_STATUS         *json.RawMessage `json:"qc_status"`         // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=486
// @Summary 		ข้อมูลความเข้มแสง
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_486 successful operation
type Metadata_486 struct {
	TELE_STATION_ID int64            `json:"tele_station_id"` // example:`151` รหัสสถานีโทรมาตร tele station's serial number
	SOLAR_DATETIME  string           `json:"solar_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บค่าความเข้มแสง record date
	SOLAR_VALUE     float64          `json:"solar_value"`     // example:`120.37` ค่าความเข้มแสง solar's value
	QC_STATUS       *json.RawMessage `json:"qc_status"`       // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=487
// @Summary 		ข้อมูลระดับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_487 successful operation
type Metadata_487 struct {
	TELE_STATION_ID    int64            `json:"tele_station_id"`    // example:`2073` รหัสสถานีโทรมาตร tele station  number
	RAINFALL_DATETIME  string           `json:"rainfall_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL_DATE_CALC string           `json:"rainfall_date_calc"` // example:`2006-01-02` วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ เนื่องจากลักษณะการจัดเก็บของปริมาณน้ำฝนจะเริ่มจาก7.00 น.ของเมื่อวาน ถึง 6.59 น.ของวันนี้ Date for calculate rainfall
	RAINFALL10M        float64          `json:"rainfall10m"`        // example:`0` ปริมาณน้ำฝนทุก 10 นาที Rainfall Every 10 minute
	RAINFALL1H         float64          `json:"rainfall1h"`         // example:`1.5` ปริมาณน้ำฝนทุก 1 ชั่วโมง Rainfall Every 1 hour
	RAINFALL24H        float64          `json:"rainfall24h"`        // example:`12.5` ปริมาณน้ำฝนทุก 24 ชั่วโมง Rainfall Every 24  hours
	QC_STATUS          *json.RawMessage `json:"qc_status"`          // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=488
// @Summary 		ข้อมูลปริมาณน้ำฝน รายวัน
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_488 successful operation
type Metadata_488 struct {
	TELE_STATION_ID    int64   `json:"tele_station_id"`    // example:`135` รหัสสถานีโทรมาตร tele station number
	RAINFALL_DATETIME  string  `json:"rainfall_datetime"`  // example:`2006-01-02T15:04:05Z07:00` วันที่เก็บปริมาณน้ำฝน Rainfall date
	RAINFALL_DATE_CALC string  `json:"rainfall_date_calc"` // example:`2006-01-02` วันที่ของปริมาณน้ำฝนสำหรับใช้ในการคำนวณ
	RAINFALL_VALUE     float64 `json:"rainfall_value"`     // example:`3.2` ปริมาณฝนรายวัน เวลา 7:01 - 7:00  Rainfall daily
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=489
// @Summary 		ข้อมูลความหนาแน่นของเมฆทั่วทั้งประเทศ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_489 successful operation
type Metadata_489 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`51`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`145`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=0PIxm6pP-VZ96GWet2NLHWsu6qzLFlzGbeRuohxxwitaiOj2gIuMy4P1obloZekT2bScntQSYnevNqKs4UiOZUoqYxIQidLwKbDv-5sr3BJgnQ6yLXdPYp6BqYLxT8mOBdZkRjwyTmf3YC8B-_5PlIEdgYfTKwnAv6yRJ8BEYsTEeK5RvyJmauPrygLtiX1QDIEwU45cuMAwv_TlJQTRabKA2nHGXpWl_oMY4qUltn0`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ข้อมูลความหนาแน่นของเมฆทั่วทั้งประเทศ.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ข้อมูลความหนาแน่นของเมฆทั่วทั้งประเทศ`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=512
// @Summary 		แผนภาพฝนสะสมรายวัน Gsmaps 10km
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_512 successful operation
type Metadata_512 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`157`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-31T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=XO-UJWqdCWpn7dvphtN5yatXhEj1YLwJX8cNrfc-MSTbzYEucialIDR95cBhG5wxwimKMPmE5_KmpnXSLD7PByOqEKULI5aYNCVXAFkOyjYmQ12nvUTDXq3f206HiY10ghUhnyP59mQ5Leb1yoIpY51u3wvkvfkw7Q69FyAS9cFlUp3E9GerQEqAZAq-Vz4L`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`gsm_010_20180831_bias.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GSMaP (10km x 10km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=513
// @Summary 		แผนภาพฝนสะสมรายวัน Gsmaps 25km
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_513 successful operation
type Metadata_513 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`159`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-31T07:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=pHN7nzf2KiZ3Y2Z1wwnpwAz8YGRzQGCmmmqa3OoWJf2n_aXmN7x2bLnyechK0L7hCK_KOQRsksEh0yNr6gIl7gActq16hkkvK-ai6GQupd0uiYRgUmkL58HxgNpBYxRWy3N4MIB7xRbaXZ3Ug_xGvagkF9gZVdhxHQisL8cqxAAVbGCMK1yc5YkSFmL23qfc`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`gsm_025_20180831_bias.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`GSMaP (25km x 25km)`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=515
// @Summary 		ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_515 successful operation
type Metadata_515 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`50`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`141`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-14T16:06:32+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=cBLJydQjVzVMDRLbrix-IOuq8TDkfleakvUxhGM5DUUxueRnjoVAy3jqMHmwMkG9RelypTQFR5C_93-ZMDK3LbuwIGx0HpgyfyyBDYSHEb9LiVmlSB-O3QmVmQaOofaFP1iEgz7Qwj6LxHM_LgHki40OlB0rhmJpz4muiD5IcSI`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`00Latest.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=516
// @Summary 		แผนภาพค่าเบี่ยงเบนความสูงระดับน้ำทะเล
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_516 successful operation
type Metadata_516 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`57`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`153`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-10T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=E7MjuwNH3pCbGxOWLeHwnqpYxkRgE_tz3VdylOL99dv4mEdP-sou0j5nu98eJMU2yHsHaDMOt6LaJvEzjJ7Da8U3rSlyoxbqmjRRRQbKdq9vFiQ7lBYM_toUPTPFmnLOmBMx9ArNkwpoQHsOvDkxHg8XknKyHH8tZJLW-yy10V02tDv-LLiKisyeGCdS_8p4gecL3L3t8GINZ3Ip1m7dcw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`duacs_global_nrt_msla_merged_h_20180810_glo_msla_n0_t0.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=517
// @Summary 		ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_517 successful operation
type Metadata_517 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`57`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`141`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=SwHw2ZzjdHxiGifBb9P0CRbf0NF9qBXtJtxp9gh-_YQQt18ufOMKLQI_OQpTJwzm4fe-OXOhs1ItKfE3ow8futajFzhYKcI7eAEOaLTlzS3vKW0DbinoTB32aCDtqazS9PPTV4WhGevZUdk32hgejYosCSDsOnUPjT3q48UL4bGd7ZK-ELlkjp8HQuWGzTdaoJSEub3d5CCHCLFgUat22w`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=518
// @Summary 		ค่าดัชนีพืชพรรณจากภาพถ่ายดาวเทียม TERRA
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_518 successful operation
type Metadata_518 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`37`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-03-02T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=aeLzTMWeOILKk3Ptze-FkcYRsuKyjNFrOyenOXanvTQkI-vEyn4CdG150g8O31l6pybSZ818uraPVVEshgb4TCyHIx_Q1Z7t_4fk2REoVLct8QMXO3AmMn8Fa9003pEArPdcjm02VwZXiGqq77GkKQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`next.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=519
// @Summary 		ค่าดัชนีพืชพรรณจากภาพถ่ายดาวเทียม AQUA
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_519 successful operation
type Metadata_519 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`36`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-03-02T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=KSGo41jSQg7xlXaFxH3ZwJHurKpBB27Wqn2zvpzw8eSVrIl8lrMws904u2OlfcgwKuknJqN0ZzETBYib7CSB8z7IrlvM0zOf1sJ4_7p8eBgXZc_7u6qT0sH2fgc5-92v0Rx3nOK2eFsfPVGfjAarmg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`next.gif`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=520
// @Summary 		แผนภาพอุณหภูมิผิวน้ำทะเลบริเวณทั่วโลก
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_520 successful operation
type Metadata_520 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`145`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2016-12-13T21:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=H4iuUbuH2lYZmvnvSpCMc8RCUgjngjveLyu_ND8tq-EycJ3v7hVFq1xklUVg2JPnq6OvuXifDucAX2UCfRooXloKCGSrry1eXNEQ0p3pRlyTyKwNxg1xbiOMKaV1HlhQJUIRGzWhrXRPs9GXWSoWKmqh7fPdiw49W68S979hidNXZwMgEQi-HKcUR-jVkAo_wAyLbqgj9bxRv9iR-BXBrhY9pDyfPwT06Ss-R7PTSzQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`แผนภาพอุณหภูมิผิวน้ำทะเลบริเวณทั่วโลก.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`แผนภาพอุณหภูมิผิวน้ำทะเลบริเวณทั่วโลก`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:` `  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=521
// @Summary 		แผนภาพความสูงคลื่นทะเล Global
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_521 successful operation
type Metadata_521 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`56`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`155`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T23:38:03+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=n9-j9QyLtA8XUktQlBsXHI0dHX8d7R3UTihYnI3M9Lku2Z3BbphScjbKnsnuS_cR6A07z49d2Zs3yTe17dtseDiMia1MTamKHL9OdZN3xnl6AqGwzMyLdsBavLjsDLbmQ6CGrO7jJh8s3xFRzj5Ekw`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`WAVE000.GIF`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=522
// @Summary 		แผนภาพอุณหภูมิผิวน้ำทะเลบริเวณ South China Sea
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_522 successful operation
type Metadata_522 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`56`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`143`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-16T23:38:24+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=79SINCI5NA3YqXvdjfqQovvGpPJAjnWANaRv31Xonj7f7c1KHyox9evZLSitFo-sqgYlh5cc01zZMVBflVDTjCMvk31n9WkMiSpRjXNkYHYUScam9TxZWBwpBVQQnxqk3jiDeTIGrMg_X4crcfyfUQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`SST.GIF`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=523
// @Summary 		ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_523 successful operation
type Metadata_523 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`51`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`141`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-15T04:30:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=W370Zktdl6z2mnoy6tzMv9_ugkf_CJTUZag9PPUVBmk0DIFoHnIfXGhCFRHLVBVutCjvIgndH6tst3kYjJDnYTwaqlbn_GZIzX-kAbiovZT8Kr0_ieZS4G4WjrlRxo-ioJJTlcuZ05iq8Q4sD_pLkYaombzREX4Qbd4opl1L015ZWgOwHZN3HAJS-ItdlMckNWLLKln_t4IUhP8jHFZ2UpD4q5IBRpPvQ6deYTG3ZBs`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`20180815.0430.himawari_8.visir.bckgr.7SEAS_Exp_Overview.DAY.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=524
// @Summary 		ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_524 successful operation
type Metadata_524 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`52`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`141`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-17T08:06:55+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=yAhRpqlgwzSOR5zODef3FmeYX_HqvmByHrv_x2ZMxTukr9O-zteLcBlWAjEcKLao41Ta1Mf9uZ2Q0ZLYCngxasyQK0UIVqHwbCpRUH7-n4cRpWzihhZhzBmTJMTA10T_vv09Maw6-6QHRB-SaGTgzI7Y3V3r1rSqh93MR4HO4Ns`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`latest.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=525
// @Summary 		ค่าความชื้นในดิน Surface
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_525 successful operation
type Metadata_525 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`163`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-12T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=NRdlTl5mxEI73RZLGqwfPFq6lWpDjMVGkqAfLFhiiOZTMo7AJ1d9Inw271l6o4UYYMBB-IQiTRMiUKqbkWDQIAaVF-wXJfpr4H7iBffNuhZmat0F0QTbSTuzVy0illiwJJWsmQoDa7As0sRj0d9HLvhfYX-Tvj3sAPzg6yi2POU`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`seasia0ssm7_201808_12_s.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=526
// @Summary 		ปริมาณและการกระจายตัวของฝน AFWA
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_526 successful operation
type Metadata_526 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`149`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-10T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=_fto3VZFO9hnK9S5CBkWQq4Nmgzi2UFAaa-zFLA3cgiWmcoeeBOOyFRL7Pb4FKZkFG_wEdbCPmMmxFJL54U7l_6HZTxYZDya_fSLQOvfXfbnUn6DwJhyTIex3Z0LReeUomU6nLbKHxD2F9HIliTL9ctPBmyLWRQz0ZF2onut7bQ`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`seasia0l_p201710_10.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=527
// @Summary 		ปริมาณและการกระจายตัวของฝน AFWA Decadal Percent Normal
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_527 successful operation
type Metadata_527 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`150`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-10T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=1wQmJPfYQXqo_XSe9JbOT9VK5-TMtW6qtCc-k6dAIlhynZUqb1ejYkAp2lUE-r6EFjzpwalKN1jEO4f0VCJv9eu3JtA6AsTdCExx4KupFi7MAq-cKTYoyz6F0u9yG6EPA3ciTDcunPiniwqOchhXcX1gKumiHcCsarrbmsJpfS0`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`seasia0l_pn201710_10.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=528
// @Summary 		ปริมาณและการกระจายตัวของฝน WMO
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_528 successful operation
type Metadata_528 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`151`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-10T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=aY5tvV1JtK1uTxb5HiTUyKS6WE1ebZkO09TZBxMwY9ND60lAHOvKXtnDrPNctaQE2aT323J-iX8NiHqQhxAPn1cHEbTxPGRBzTk_g70PP23XcBYdfEEnvWoduNjyEkJSj57TCaHZUWvuQVKvFCHlQORwuryO36OQGnmBn0opPS0`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`seasia0p201710_10_s.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=529
// @Summary 		ปริมาณและการกระจายตัวของฝน WMO Decadal Percent Normal
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_529 successful operation
type Metadata_529 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`55`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`152`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-10T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=b-PzxgD9KR3Dy4lG8ENc0DK6Ah5djBZ52A6QCHW7JW_TmgObv_U4De3W9OkzoSvQnDrbUqnU26xkHZealFOQKAqGq2LUpkjCu2KwfIDpT_OziwSOVdtHveA7bmoRP8bjZRhX9e2sQC4lgzSH5d7cAxuQ7vruIyfy7wayBXxARo0`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`seasia0pn201710_10_s.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=530
// @Summary 		ข้อมูลปริมาณน้ำ ฝนจากภาพถ่ายดาวเทียม COMS ภาพกว้าง Thailand
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_530 successful operation
type Metadata_530 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`4`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-12-31T22:20:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=1CokoANt57SJn6_NjJkuu_H6q4ioPU5uypAvVW3uHr6fHsiHEisUWdDlWbHUF-pEitfLa86F3DibV4cxcbMZW69p5SXkE5fv0Sh0fehIv10I2ROLNRr8LlY2KbLqRQttqmlRpa0BpwulVbR_ZMxKC3rHrzRloYcjBhb0XtuMtXr4BsKV_mqrLXqamOdy46mt`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`coms_ri1d_20171231_thailand.jpg`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=531
// @Summary 		ข้อมูลปริมาณน้ำ ฝนจากภาพถ่ายดาวเทียม COMS ภาพกว้าง South East Asia
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_531 successful operation
type Metadata_531 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`6`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T16:16:32+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=Rp7fxx5wPM0gRuTyyGdgl86Mfo4-ExGcnKJNNmF8lFELxJmLzvWmAv9BeeiBWNY8KLSLwYDrUKLw5z9VFq3h9-gNwX8YX-EVnbfYZzOeJG3aEfhWloADrKgKew0FvW0vfb1zBIGFPt9XyIY4baDR2ZJZPIMRPQCrh32l93bHkWx-R30mAq4FzCLMNK1grrEbihe5oQ1dqahXF8mVAw_YLA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`coms_mi_le2_ri_geo_201708280845.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=532
// @Summary 		ข้อมูลปริมาณน้ำ ฝนจากภาพถ่ายดาวเทียม COMS ภาพกว้าง Asia
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_532 successful operation
type Metadata_532 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`11`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`5`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-08-28T16:12:03+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=7ZGjkS6QuVI9CA0ADGg8g4HrXod0EBMVnr6fP9uPBZWtAHi5l6Bn9P5DQYqTOszuv41qQ_KUgyqM337CA0NCiCGSO5VxyCVOjoPZps0uguUssLtC5lgU4oGqX5vW4lXUAQKbDTDFdSiAc4TAfJIlht-ULq9SUwEa3tsTA2mAmN8Ix4o34Db_neNGcNDWIdtE`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`coms_mi_le2_ri_cn_201708280845.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:``  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=533
// @Summary 		ภาพถ่ายล่าสุดจากดาวเทียม Himawari-8 Animation
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_533 successful operation
type Metadata_533 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`50`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`141`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2018-08-14T15:09:30+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=1kESB3lNFz3NoIf_cMPdRhrtc6oOHLh3_vstfVzoaKvXWtchPaFkBjasbV690eU2ArYonmjBOpVqmC7HTXj6BSkKExuXBJU1dBUSaEv3k5sMzRK_HERe18oO42uzLInPCR-OYjywGUZfrJVQGsmtPA`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`00Movie.mp4`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`ภาพเมฆล่าสุด ที่มาจาก มหาวิทยาลัย kochi`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=534
// @Summary 		แผนภาพการเปลี่ยนแปลงของอณุหภูมิผิวน้ำทะเล แสดงราย 2 สัปดาห์
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_534 successful operation
type Metadata_534 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`9`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`14`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-10-01T00:00:00+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=bX8_HLCQWnGArWgtvwG9rtmKRpd488MEZp_jdM-16KJQafmVwTdQnebLMwqIrCeWx9b3obuewmFRvDrYX3Z4FYMOKv_-llaX-5Ar_VA68CcKVVafzN7SuF9cDgg7HE-b2pMqNHfWYCWbuJvu1iFBdU8n20i1mesr2Ao8G6ZU6jE`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`ssh_diff_20171001_20171008.png`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`Change in weekly Sea Surface Height Anomalies 2017`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=539
// @Summary 		บัญชีข้อมูล (ทดสอบ)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_539 successful operation
type Metadata_539 struct {
	ID                   int64            `json:"id"`                   // รหัสบัญชีข้อมูล metadata serial number
	METADATASERVICE_NAME *json.RawMessage `json:"metadataservice_name"` // ชื่อบัญชีข้อมูลที่ให้บริการในคลังข้อมูล
	SUBCATEGORY_ID       int64            `json:"subcategory_id"`       // รหัสหมวดย่อยของข้อมูล subcategory number
	AGENCY_ID            int64            `json:"agency_id"`            // รหัสหน่วยงานเจ้าของข้อมูล agency id
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=540
// @Summary 		แผนที่แม่น้ำ (ทดสอบ)
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_540 successful operation
type Metadata_540 struct {
	AGENCY_ID      int64  `json:"agency_id"`      // example:`32`  รหัสหน่วยงาน agency's serial number
	MEDIA_TYPE_ID  int64  `json:"media_type_id"`  // example:`54`  รหัสแสดงชนิดข้อมูลสื่อ
	MEDIA_DATETIME string `json:"media_datetime"` // example:`2017-11-02T11:53:26+07:00`  วันที่เก็บข้อมูลสื่อ record date
	MEDIA_PATH     string `json:"media_path"`     // example:`http://api2.thaiwater.net:9200/api/v1/thaiwater30/shared/file?file=M5UQNlFkRlNzGwePuWVTH-BnvyD629JllfbS5FTO_99zinfchvzny6JZ5lz4bzFWrYDNm5G6Cqr-3_5k4SfWl_p7rNEoFJAkB3odksKPrBsZtHNnhKuT8cQYdAJX-kjLM0rtG-dNqp7QLVuXO2jlUg`  ที่อยู่ของไฟล์ข้อมูลสื่อ file path location
	FILENAME       string `json:"filename"`       // example:`river_map22.pdf`  ชื่อไฟล์ของข้อมูลสื่อ file name
	MEDIA_DESC     string `json:"media_desc"`     // example:`nso_rivermap`  รายละเอียดของข้อมูลสื่อ description
	REFER_SOURCE   string `json:"refer_source"`   // example:``  แหล่งข้อมูลสำหรับอ้างอิง reference source
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=544
// @Summary 		ข้อมูลระดับน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_544 successful operation
type Metadata_544 struct {
	WATERLEVEL_DATETIME string           `json:"waterlevel_datetime"` // example:`2006-01-02T15:04:05Z07:00` วันที่ตรวจสอบค่าระดับน้ำ
	TELE_STATION_ID     int64            `json:"tele_station_id"`     // example:`3458` รหัสสถานีโทรมาตร tele station's serial number
	WATERLEVEL_M        float64          `json:"waterlevel_m"`        // ระดับน้ำ เมตร รสม.
	QC_STATUS           *json.RawMessage `json:"qc_status"`           // example:`null` สถานะของการตรวจคุณภาพข้อมูล quality control status
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=545
// @Summary 		พายุ US Naval Research Laboratory
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_545 successful operation
type Metadata_545 struct {
	AGENCY_ID           int64   `json:"agency_id"`           // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	STORM_DATETIME      string  `json:"storm_datetime"`      // วันที่และเวลาที่เก็บข้อมูลเส้นทางพายุ
	STORM_LAT           float64 `json:"storm_lat"`           // ละติจูด
	STORM_DIRECTIONLAT  string  `json:"storm_directionlat"`  // ทิศทาง N หรือ S
	STORM_LONG          float64 `json:"storm_long"`          // ลองติจูด
	STORM_DIRECTIONLONG string  `json:"storm_directionlong"` // ทิศทาง E หรือ W
	STORM_PRESSURE      string  `json:"storm_pressure"`      // ความกดอากาศ mb (pressure is always in mb)
	STORM_NAME          string  `json:"storm_name"`          // ชื่อพายุ
}

// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=546
// @Summary 		พายุ wunderground
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Parameter		start_date	query	string  วันที่เริ่มต้นของข้อมูล
// @Parameter		end_date	query	string  วันที่สิ้นสุดของข้อมูล ช่วงวันที่เลือกต้องไปเกิน 3 วัน
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_546 successful operation
type Metadata_546 struct {
	AGENCY_ID           int64   `json:"agency_id"`           // รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency's serial number
	STORM_DATETIME      string  `json:"storm_datetime"`      // วันที่และเวลาที่เก็บข้อมูลเส้นทางพายุ
	STORM_LAT           float64 `json:"storm_lat"`           // ละติจูด
	STORM_DIRECTIONLAT  string  `json:"storm_directionlat"`  // ทิศทาง N หรือ S
	STORM_LONG          float64 `json:"storm_long"`          // ลองติจูด
	STORM_DIRECTIONLONG string  `json:"storm_directionlong"` // ทิศทาง E หรือ W
	STORM_WIND          string  `json:"storm_wind"`          // Maximum sustained wind in storm (kt)
	STORM_PRESSURE      string  `json:"storm_pressure"`      // ความกดอากาศ mb (pressure is always in mb)
	STORM_NAME          string  `json:"storm_name"`          // ชื่อพายุ
}


// @DocumentName	v1.dataservice
// @Service 		thaiwater30/api_service?mid=550
// @Summary 		ข้อมูลพื้นฐานของสถานีวัดระดับน้ำ โทรมาตรขนาดเล็ก เช่น รหัส,ชื่อ,พิกัด,ระดับตลิ่ง,ระดับท้องน้ำ
// @Parameter		eid	query	string	required:true	รหัสเฉพาะที่ไม่ซ้ำกัน เพื่อใช้ในการเข้าถึงข้อมูล
// @Method			GET
// @Response		404	-	no eid
// @Response		422	-	invalid parameter
// @Response		200 Metadata_550 successful operation
type Metadata_550 struct {
	ID                   int64            `json:"id"`                   // example:`19` รหัสสถานีโทรมาตร tele station's serial number
	TELE_STATION_OLDCODE string           `json:"tele_station_oldcode"` // example:`BKK012` รหัสโทรมาตรเดิมของแต่ละหน่วยงาน old tele station  number
	TELE_STATION_NAME    *json.RawMessage `json:"tele_station_name"`    // example:`{"en":"Krung Thep 12","th":"คลองสำโรง บางเสาธง","jp":"バンコク12"}` ชื่อสถานีโทรมาตร tele station's name
	AGENCY_ID            int64            `json:"agency_id"`            // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ agency  number
	PROVINCE_NAME        *json.RawMessage `json:"province_name"`        // example:`{"th": "กรุงเทพมหานคร"}` ชื่อจังหวัดของประเทศไทย
	AMPHOE_NAME          *json.RawMessage `json:"amphoe_name"`          // example:`{"th": "พระบรมมหาราชวัง"}` ชื่ออำเภอของประเทศไทย
	TUMBON_NAME          *json.RawMessage `json:"tumbon_name"`          // example:`{"th": "พระนคร"}` ชื่อตำบลของประเทศไทย
	TELE_STATION_LAT     float64          `json:"tele_station_lat"`     // example:`13.589267` ละติจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) latitude
	TELE_STATION_LONG    float64          `json:"tele_station_long"`    // example:`100.802235` ลองจิจูดของสถานีโทรมาตร (หน่วย : Decimal Degree) longitude
}
