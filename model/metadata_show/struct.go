// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: Thitiporn Meeprasert <thitiporn@haii.or.th>

package metadata_show

type Struct_MetadataShow_Param struct {
	ID                     int64  `json:"id"`                        // example:`11111` รหัส
	MetadataID             int64  `json:"metadata_id"`               // example:`1` รหัสบัญชีข้อมูล
	MetadataName           string `json:"metadata_name"`             // example:`ข้อมูลฝนจากระบบโทรมาตร` บัญชีข้อมูล
	AgencyID               int64  `json:"agency_id"`                 // example:`1` รหัสหน่วยงาน
	AgencyName             string `json:"agency_name"`               // example:`สถาบันสารสนเทศทรัพยากรน้ำและการเกษตรฯ` หน่วยงาน
	MetadataShowSystemID   int64  `json:"metadata_show_system_id"`   // example:`1` รหัสระบบที่นำข้อมูลไปแสดง
	MetadataShowSystemName string `json:"metadata_show_system_name"` // example:`คลังข้อมูลน้ำและภูมิอากาศแห่งชาติ` ชื่อระบบที่นำข้อมูลไปแสดง
	SubCategoryID          int64  `json:"subcategory_id"`            // example:`1` รหัสส่วนข้อมูลที่นำไปแสดง
	SubCategoryName        string `json:"subcategory_name"`          // example:`ฝน` ส่วนข้อมูลที่นำไปแสดง
	ConnectionFormat       string `json:"connection_format"`         // example:`ฝน` ส่วนข้อมูลที่นำไปแสดง
	MetadataMethod         string `json:"metadata_method"`           // example:`ฝน` ส่วนข้อมูลที่นำไปแสดง

}
