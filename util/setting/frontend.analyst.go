package setting

import (
	//	"encoding/json"
	"os"
	"path/filepath"

	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/thaiwater30/util/b64"
	
)

type Struct_RadarTypeOrder struct {
	Radar_type      string `json:"radar_type"`      // example:`cri240` ประเภทเรดาร์
	Radar_name      string `json:"radar_name"`      // example:`เรดาห์เชียงราย รัศมี 240 กม.` ชื่อเรดาร์
	Radar_frequency int    `json:"radar_frequency"` // example:`60`
	Agency          string `json:"agency"`          // example:`tmd` หน่วยงาน
	Timezone	    string `json:"timezone"` 		// example:`GMT` zone เวลา
	Band			string `json:"band"`			// example:`sband`
	
	FileName  string // error file name
	FilePath  string // error file path
	MediaPath string // error file path (encrypt)
}

//	สร้างตัว RadarTypeOrder พร้อม setting ค่าจาก db
func New_Struct_RadarTypeOrder() ([]*Struct_RadarTypeOrder, error) {
	s := make([]*Struct_RadarTypeOrder, 0)
	err := setting.GetSystemSettingPtr("Frontend.analyst.Radar.RadarTypeOrder", &s)

	if err != nil {
		return nil, err
	}
	return s, nil
}

//	get radar error file
//	Return:
//		radar error file if exist
func (s *Struct_RadarTypeOrder) DefaultErrorFile() {
	dataPathPrefix := setting.GetSystemSetting("server.service.dataimport.DataPathPrefix")

	if s.Agency == "drraa" {
		folderBand := s.Band
	
		errorPath := filepath.Join("product/radar/history", s.Agency, folderBand, s.Radar_type, s.Radar_type+"_error.jpg")
		
		if _, err := os.Stat(filepath.Join(dataPathPrefix, errorPath)); err == nil {
			// exist
			s.FilePath = errorPath
			s.MediaPath, _ = b64.EncryptText(errorPath)
			s.FileName = s.Radar_type + "_error.jpg"
		}
			
	} else {
		// เอาแค่ชื่อเรดาห์ 3ตัวแรก ไม่เอารัศมี (ubn240)
		folderName := s.Radar_type[0:3]
		// บางเรดาห์ชื่อ 2ตัวแรก (pn240)
		if folderName[2:] == "1" || folderName[2:] == "2" {
			folderName = folderName[0:2]
		}
		errorPath := filepath.Join("product/radar/history", s.Agency, folderName, s.Radar_type+"_error.jpg")
		// find /data/thaiwater/thaiwaterdata/test1/product/radar/history/tmd/cri240_error.jpg
		if _, err := os.Stat(filepath.Join(dataPathPrefix, errorPath)); err == nil {
			// exist
			s.FilePath = errorPath
			s.MediaPath, _ = b64.EncryptText(errorPath)
			s.FileName = s.Radar_type + "_error.jpg"
		}
	}
	
	
}

