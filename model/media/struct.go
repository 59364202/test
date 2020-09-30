package media

import (
	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_media_type "haii.or.th/api/thaiwater30/model/media_type"
)

type Struct_Media struct {
	Agency      *model_agency.Struct_Agency        `json"agency,omitempty"`     // หน่วยงาน
	MediaType   *model_media_type.Struct_MediaType `json"media_type,omitempty"` // ชนิดของสื้อ
	ID          int64                              `json:"id"`                  // example:`8264994` รหัสแสดงข้อมูลสื่อ
	Datetime    string                             `json:"media_datetime"`      // example:`2006-01-02` วันที่เก็บข้อมูลสื่อ
	Path        string                             `json:"media_path"`          // example:`product/precipitation/asia/latest/haii` ที่อยู่ของไฟล์ข้อมูลสื่อ
	Description string                             `json:"media_desc"`          // example:`WRF-ROMS Model, Asia (27km x 27km)` รายละเอียดของข้อมูลสื่อ
	Filename    string                             `json:"filename"`            // example:`d01_day01.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	ReferSource string                             `json:"refer_source"`        // example:`https://api.haii.or.th/tiservice/v1/ws/MBtOTp6IUXbjaCxhQoFQNrFgZUCzNgbo/model/wrfroms/latest` แหล่งข้อมูลสำหรับอ้างอิง
	FileStatus  bool                               `json:"file_status"`         // example:`true` มีไฟล์
	ImageID     string                             `json:"image_id"`            // example:`AAECAwQFBgcICQoLDA0ODz2sq-vcpj7l00s__U-WBd0SNWAl6wOJvLeSEe_SSJcmyw9mYoy7s0HaiNgksfhyvZ2kdDs=` ที่อยู่ของไฟล์แบบเข้ารหัส
}

type Struct_Media_InputParam struct {
	AgencyID    string `json:"agency_id"`     // example: 19 รหัสหน่วยงาน
	MediaTypeID string `json:"media_type_id"` // example: 18  รหัสชนิดข้อมูลสื้อ
	StartDate   string `json:"start_date"`    // example: 2006-01-02 วันที่เริ่มต้น
	EndDate     string `json:"end_date"`      // example: 2006-01-02 วันที่สิ้นสุด
}

type MediaFileOutput struct {
	MediaTypeID   		int64      	`json:"media_type_id"`     		// example: 18  รหัสชนิดข้อมูลสื้อ
	MediaType     		string      `json:"media_type"`       		// example:`Precipitation` ชนิดของสื่อ
	MediaSubType  		string      `json:"media_subtype"`    		// example:`Thailand` ชนิดย่อยข้อมูลสื่อ
	MediaCategory 		string      `json:"media_category"`   		// example:`image` ประเภทของ media เช่น image, animation, excel etc
	Description   		string      `json:"description"`      		// example:`description` รายละเอียดของข้อมูลสื่อ
	DateTime      		string      `json:"datetime"`         		// example:`2006-01-02 15:04` วันที่เก็บข้อมูลสื่อ
	URL           		string      `json:"media_path"`       		// example:`QWE1QTH35BDS51ASD` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	URLThumb      		interface{} `json:"media_path_thumb"` 		// example:`QWE1QTH3AD@LKH1238D` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	Filename      		string      `json:"filename"`         		// example:`thailand.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	Filepath      		string      `json:"filepath"`         		// example:`/product/precipitation/thailand/2016/06/10` ที่อยู่ของไฟล์ข้อมูลสื่อ
	FilenameThumb 		interface{} `json:"filename_thumb"`   		// example:`thumb-thailand.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	FilepathThumb 		interface{} `json:"filepath_thumb"`   		// example:`/product/precipitation/thailand/2016/06/10` ที่อยู่ของไฟล์ thumb ข้อมูลสื่อ
	FilePathAnimation  	interface{} `json:"file_path_animation"`  	// example:`cri240.gif`ที่อยู่ของไฟล์ animation ข้อมูลสื่อ
	FilenameAnimation  	interface{} `json:"filename_animation"`   	// example:`cri240.jpg`ชื่อไฟล์ animation ของข้อมูลสื่อ
	MediaPathAnimation  interface{} `json:"media_path_animation"`	// example:`QWE1QTH3AD@LKH1238D` ลิ้งค์ของ animation ข้อมูลสื่อ
	// DebugSql			string		`json:"sql"`					// example:`SELECT * FROM ....` sql ที่ใช้ดึงข้อมูล
	// DebugParam			string		`json:"param"`					// example:`....` Parameter ที่ใช้ดึงข้อมูล
}

type PdfFileOutput struct {
	AgencyID    		int64 		`json:"agency_id"`				// example: 74 รหัสหน่วยงานเจ้าของข้อมูล
	MediaTypeID   		int64      	`json:"media_type_id"`     		// example: 18  รหัสชนิดข้อมูลสื้อ
	Description   		string      `json:"description"`      		// example:`description` รายละเอียดของข้อมูลสื่อ
	DateTime      		string      `json:"datetime"`         		// example:`2006-01-02 15:04` วันที่เก็บข้อมูลสื่อ
	URL           		string      `json:"media_path"`       		// example:`QWE1QTH35BDS51ASD` ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Filename      		string      `json:"filename"`         		// example:`thailand.jpg` ชื่อไฟล์ของข้อมูลสื่อ
	Filepath      		string      `json:"filepath"`         		// example:`/product/precipitation/thailand/2016/06/10` ที่อยู่ของไฟล์ข้อมูลสื่อ
	// DebugSql			string		`json:"sql"`					// example:`SELECT * FROM ....` sql ที่ใช้ดึงข้อมูล
	// DebugParam			string		`json:"param"`					// example:`....` Parameter ที่ใช้ดึงข้อมูล
}

type Struct_Radar struct {
	RadarType          string      `json:"radar_type"`           // example:`cri240`ชนิดเรดาร์
	RadarName          string      `json:"radar_name"`           // example:`เรดาห์ที่เชียงราย รัศมี 240 กม.`ชื่อเรดาร์
	MediaDatetime      string      `json:"media_datetime"`       // example:`2006-01-02 15:04`วันที่เก็บข้อมูลสื่อ
	FilePath           string      `json:"file_path"`            // example:`product/radar/cri240/2006/01/02`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename           string      `json:"filename"`             // example:`cri240.jpg`ชื่อไฟล์ของข้อมูลสื่อ
	MediaPath          string      `json:"media_path"`           // example:`QW123SW3158FT89HDA6C7`ลิ้งค์ของไฟล์ข้อมูลสื่อ
	FilenameThumb      interface{} `json:"filename_thumb"`       // example:`thumb-cri240.jpg`ชื่อไฟล์ thumb ของข้อมูลสื่อ
	MediaPathThumb     interface{} `json:"media_path_thumb"`     // example:`QW123SW3158FT89HDA6C7`ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
	FilePathAnimation  interface{} `json:"file_path_animation"`  // example:`cri240.gif`ที่อยู่ของไฟล์ animation ข้อมูลสื่อ
	FilenameAnimation  interface{} `json:"filename_animation"`   // example:`cri240.jpg`ชื่อไฟล์ animation ของข้อมูลสื่อ
	MediaPathAnimation interface{} `json:"media_path_animation"` // example:`QW123SW3158FT89HDA6C7` ลิ้งค์ของไฟล์ animation ข้อมูลสื่อ
}

type Struct_RadarHistory struct {
	MediaDatetime  string      `json:"media_datetime"`   // example:`2006-01-02 15:04`วันที่เก็บข้อมูลสื่อ
	FilePath       interface{} `json:"file_path"`        // example:`product/radar/cri240/2006/01/02`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename       interface{} `json:"filename"`         // example:`cri240.jpg`ชื่อไฟล์ของข้อมูลสื่อ
	MediaPath      interface{} `json:"media_path"`       // example:`QW123SW3158FT89HDA6C7`ลิ้งค์ของไฟล์ข้อมูลสื่อ
	FilenameThumb  interface{} `json:"filename_thumb"`   // example:`thumb-cri240.jpg`ชื่อไฟล์ thumb ของข้อมูลสื่อ
	MediaPathThumb interface{} `json:"media_path_thumb"` // example:`QW123SW3158FT89HDA6C7` ลิ้งค์ของไฟล์ thumb ข้อมูลสื่อ
}

type Result_Struct_RadarHistory struct {
	DateTime 	string                 	`json:"datetime"` 	// example:`2006-01-02`วันเวลา
	RadarName	string      			`json:"radar_name"`	// example:`เรดาห์ที่เชียงราย รัศมี 240 กม.`ชื่อเรดาร์
	Timezone	string      			`json:"timezone"`   // example:`GMT` zone เวลา
	Agency	    string      			`json:"agency"`		// example:`tmd` ชื่อหน่วยงาน
	Data     	[]*Struct_RadarHistory 	`json:"data"`     	// ข้อมูล
	// DebugSql	string					`json:"sql"`		// example:`SELECT * FROM ....` sql ที่ใช้ดึงข้อมูล
}

type Param_Media struct {
	AgencyId    int64 `json:"agency_id"`
	MediaTypeId int64 `json:"media_type_id"`
}

type Struct_ReportHistory struct {
	DateTime 	string                 	`json:"datetime"` 		// example:`2006-01-02`วันเวลา
	MediaDesc	string      			`json:"desc"`			// example:`เรดาห์ที่เชียงราย รัศมี 240 กม.`ชื่อเรดาร์
	FilePath    string      			`json:"file_path"`  	// example:`product/radar/cri240/2006/01/02`ที่อยู่ของไฟล์ข้อมูลสื่อ
	Filename    string      			`json:"filename"`   	// example:`cri240.jpg`ชื่อไฟล์ของข้อมูลสื่อ
	MediaPath   interface{} 			`json:"media_path"` 	// example:`QW123SW3158FT89HDA6C7`ลิ้งค์ของไฟล์ข้อมูลสื่อ
	Agency	    string      			`json:"agency"`			// example:`tmd` ชื่อหน่วยงานStruct_ReportHistory
	MediaTypeId int64 					`json:"media_type_id"`	// example: `รหัสหมวด media เช่น 74`
	// DebugSql	string					`json:"sql"`			// example:`SELECT * FROM ....` sql ที่ใช้ดึงข้อมูล
}
