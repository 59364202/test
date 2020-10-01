// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package agency is a model for public.agency table. This table store agency.
package agency

import (
	"encoding/json"
)

// โครงสร้างใหญ่ที่มีเพิ่มมาจากโครงสร้างเล็ก
type Struct_Agency struct {
	Struct_A
	Agency_shortname json.RawMessage `json:"agency_shortname,omitempty"` // example:`{"en": "HAII"}` ชื่อย่อของหน่วยงาน (ภาษาอังกฤษ)

	Department_id   int64           `json:"department_id,omitempty"`   // example:`156` รหัสกรม
	Department_name json.RawMessage `json:"department_name,omitempty"` // example:`{"th":"สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)"}` ชื่อกรม
	Ministry_id     int64           `json:"ministry_id,omitempty"`     // example:`17` ลำดับข้อมูลกระทรวง
	Ministry_name   json.RawMessage `json:"ministry_name,omitempty"`   // example:`{"th":"กระทรวงวิทยาศาสตร์และเทคโนโลยี","en":"Ministry of Science and Technology"}` ชื่อกระทรวง
	Logo            string          `json:"logo,omitempty"`            // example:`iVBORw0KGgoAAAANSU...` โลโก้หน่วยงาน
	Aspects         json.RawMessage `json:"aspects,omitempty"`         // example:`{apects:[5,6],leader:null}` ข้อมูลด้านน้ำและภูมิอากาศ
}

// โครงสร้างเล็กๆ ที่มีแค่ id, agency_name
type Struct_A struct {
	Id          int64           `json:"id,omitempty"`          // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ
	Agency_name json.RawMessage `json:"agency_name,omitempty"` // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
}

// โครงสร้างสำหรับบอกรายละเอียดว่า หน่วยงานมีบัญชีข้อมูลแต่ละสถานะเป็นจำนวนเท่าไหร่
type Struct_AgencyShopping struct {
	AgencyId                   int64           `json:"agency_id"`                    // example:`9` รหัสหน่วยงานที่เชื่อมโยงกับคลังฯ
	AgencyName                 json.RawMessage `json:"agency_name,omitempty"`        // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงานที่เชื่อมโยงกับคลังฯ
	Metadata                   int64           `json:"metadata"`                     // example:`9` รหัสบัญชีข้อมูล metadata serial number
	MetadataStatus_Connect     int64           `json:"metadata_status_connect"`      // example:`6` จำนวนบัญชีข้อมูลที่สถานะเป็น connect
	MetadataStatus_WaitUpdate  int64           `json:"metadata_status_wait_update"`  // example:`2` จำนวนบัญชีข้อมูลที่สถานะเป็น wait for update
	MetadataStatus_WaitConnect int64           `json:"metadata_status_wait_connect"` // example:`1` จำนวนบัญชีข้อมูลที่สถานะเป็น wait for connect
	DataService                int64           `json:"dataservice"`                  // example:`10` จำนวนครั้งที่มีคนขอใช้บริการ
}

// โครงสร้างที่ใช้รับค่าจาก user ที่ใช้สำหรับ insert, update, delete
type Param_PostAgency struct {
	AgencyName      json.RawMessage `json:"agency_name"`      // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงาน
	AgencyShortName json.RawMessage `json:"agency_shortname"` // example:`{"en": "HAII"}` ชื่อย่อหน่วยงาน
	DepartmentId    string          `json:"department_id"`    // example:`156` รหัสกรม
	MinistryId      string          `json:"ministry_id"`      // example:`17` รหัสกระทรวง
	Logo            string          `json:"logo"`             //example:`iVBORw0KGgoAAAANSU...` โลโก้หน่วยงาน
}

// โครงสร้างหลักที่ใช้รับค่าจาก user ที่ใช้สำหรับ insert, update, delete
type Param_Agency struct {
	Param_PostAgency
	Id     string `json:"id"` // example:`9` รหัสหน่วยงาน
	UserId int64  `json:"-"`  // รหัสผู้ใช้
}

type Logo_Agency struct {
	Id   int64  `json:"id"`   //example:`9` รหัสหน่วยงาน
	Logo string `json:"logo"` //example:`iVBORw0KGgoAAAANSU...` โลโก้หน่วยงาน
}
