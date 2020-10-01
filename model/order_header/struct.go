// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package order_header is a model for dataservice.order_header table. This table store order_header.
package order_header

import (
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
	model_order_status "haii.or.th/api/thaiwater30/model/order_status"

	"encoding/json"
)

type Field_Id struct {
	Id int64 `json:"id"` // example:`47` รหัส order header
}
type Field_Order_Datetime struct {
	Order_Datetime string `json:"order_datetime"` // example:`2006-01-02 15:04` วันที่ขอข้อมูล
}
type Field_Order_Quality struct {
	Order_Quality int64 `json:"order_quality"` // example:`4` จำนวนรายการของคำขอ
}
type Field_Order_Forexternal struct {
	Order_Forexternal bool `json:"order_forexternal"` // example:`false` การให้บริการข้อมูลสำหรับบุคคลภายนอก
}

type Struct_OH struct {
	Field_Id
	Field_Order_Datetime
	Field_Order_Quality
	Field_Order_Forexternal

	OrderStatus *model_order_status.Struct_OrderStatus `json:"order_status,omitempty"` // สถานะคำขอ
}

type Struct_OrderHeader struct {
	Id                int64           `json:"id"`                         // example:`47` รหัส order header
	User_Id           int64           `json:"user_id"`                    // example:`68` รหัสผู้ขอใช้บริการ
	User_Fullname     string          `json:"user_fullname"`              // example:`Aditep Aditep(HAII User)` ชื่อผู้ขอใช้บริการ
	User_Agency_Id    int64           `json:"user_agency_id"`             // example:`9` รหัสหน่วยงานของผู้ขอใช้บริการ
	User_Agency_Name  json.RawMessage `json:"user_agency_name,omitempty"` // example:`{"th": "สถาบันสารสนเทศทรัพยากรน้ำและการเกษตร (องค์การมหาชน)", "en": "Hydro and Agro Informatics Institute"}` ชื่อหน่วยงานผู้ขอใช้บริการ
	Order_Status_id   int64           `json:"order_status_id"`            // example:`4` รหัสของสถานะการขอข้อมูล
	Order_Datetime    string          `json:"order_datetime"`             // example:`2006-01-02 15:04` วันที่ขอข้อมูล
	Order_Quality     int64           `json:"order_quality"`              // example:`4` จำนวนรายการของคำขอ
	Order_Purpose     string          `json:"order_purpose"`              // example:`กรมเจ้าท่าทดสอบ` วัตถุประสงค์ของการขอข้อมูล
	Order_Forexternal bool            `json:"order_forexternal"`          // example:`false` การให้บริการข้อมูลสำหรับบุคคลภายนอก

	OrderStatus *model_order_status.Struct_OrderStatus   `json:"order_status,omitempty"` // สถานะคำขอ
	OrderDetail []*model_order_detail.Struct_OrderDetail `json:"order_detail,omitempty"` // รายการที่ขอ
}

type Param_OrderHeader struct {
	Id                int64  `json:"-"`
	User_Id           int64  `json:"user_id"` // example:`0` รหัสของผู้ใช้ที่ต้องการทำการแทน ไม่มีใส่ 0
	Order_Quality     int64  `json:"-"`
	Order_Purpose     string `json:"order_purpose"`     // example:`เพื่อทดสอบ` วัตถุประสงค์การนำข้อมูลไปใช้งาน
	Order_Forexternal bool   `json:"order_forexternal"` // example:`false` สำหรับบุคคลภายนอก

	OrderDetail []*model_order_detail.Param_OrderDetail `json:"order_detail,omitempty"` // รายละเอียดคำขอแต่ละรายการ
}

type Param struct {
	Date_Start      string `json:"datestart"`       // example:`2006-01-02` วันที่เริ่มต้น
	Date_End        string `json:"dateend"`         // example:`2006-01-02` วันที่สิ้นสุด
	Agency_Id       int64  `json:"agency_id"`       // example:`9` รหัสหน่วยงาน
	Statud_Id       int64  `json:"status_id"`       // example:`4` รหัสของสถานะการขอข้อมูล
	Order_Header_Id int64  `json:"order_header_id"` // example:`47` รหัส order header
	Date            string `json:"date"`            // example:`2006-01-02` วันที่
	User_Id         string `json:"user_id"`         // example:`68` รหัสผู้ขอใช้บริการ
}

type Struct_Order_Purpose struct {
	Id            int64  `json:"id"`            // example:`47` รหัส order header
	Order_Purpose string `json:"order_purpose"` // example:`เพื่อทดสอบ` วัตถุประสงค์การนำข้อมูลไปใช้งาน
}
