package event_log_sink_target

import (
	"encoding/json"
	"haii.or.th/api/thaiwater30/util/selectoption"
)

type TargetDetails struct {
	ID        int64  `json:"id"`             // example:`1` รหัสของ target
	Condition string `json:"condition_name"` // example:`condition` condition name
	Method    string `json:"method_name"`    // example:`method` method name
	Lang      string `json:"lang"`           // example:`th` ภาษาที่ใช้ในการส่ง
	Group     string `json:"group_name"`     // example:`System` ชื่อกลุ่ม
}

type TargetEdit struct {
	ID        int64  `json:"id"`           // example:1 รหัสของ target
	Condition int64  `json:"condition_id"` // example:3 condition id
	Method    int64  `json:"method_id"`    // example:5 method id
	Lang      string `json:"lang"`         // example:`th` ภาษาที่ใช้ในการส่ง
	Group     int64  `json:"group_id"`     // example:3 ชื่อกลุ่ม
}

type TargetInputAddSwagger struct {
	Condition int64   `json:"condition_id"` // example:5 condition id เช่น 1
	Method    int64   `json:"method_id"`    // example:4 method id เช่น  3
	Lang      string  `json:"lang"`         // example:`th` ภาษาที่ใช้ในการส่ง เช่น th
	Group     []int64 `json:"group_id"`     // example:[1,2,3,4] รหัสกลุ่มที่ใช้ในการส่ง เช่น [1,2,3,4]
	Color     string  `json:"color"`        // example:`AAAAAA` สีของกลุ่ม เช่น FFFFFF
}

type TargetInputAdd struct {
	ID        int64   `json:"id"`           // example:1 รหัสของ target เช่น 1
	Condition int64   `json:"condition_id"` // example:4 condition id เช่น 4
	Method    int64   `json:"method_id"`    // example:5 method id เช่น  5
	Lang      string  `json:"lang"`         // example:`th` ภาษาที่ใช้ในการส่ง เช่น  th
	Group     []int64 `json:"group_id"`     // example:[1,2,3,4] รหัสกลุ่มที่ใช้ในการส่ง เช่น [1,2,3,4]
	Color     string  `json:"color"`        // example:`AAAAAA` สีของกลุ่ม เช่น AAAAAA
}

type TargetInput struct {
	ID        int64  `json:"id"`           // example:1 รหัสของ target เช่น 1
	Condition int64  `json:"condition_id"` // example:5 รหัส condition เช่น 5
	Method    int64  `json:"method_id"`    // example:4 method name เช่น 4
	Lang      string `json:"lang"`         // example:`th` ภาษาที่ใช้ในการส่ง เช่น th
	Group     int64  `json:"group_id"`     // example:1 รหัสกลุ่มที่ใช้ในการส่ง เช่น 1
	Color     string `json:"color"`        // example:`AAAAAA` สีของกลุ่ม เช่น AAAAAA
}

type TargetSelectOptionSwagger struct {
	Condition []*ConList          `json:"condition_list"`        // รายชื่อเงื่อนไขการส่ง
	Method    []*MethodList       `json:"method_list"`           // รายชื่อ method
	Lang      json.RawMessage     `json:"lang,omitempty"`        // ภาษา
	Group     []*GroupTargetColor `json:"permission_group_list"` // รายชื่อกลุ่มที่มีสิทธิ
}

type ConList struct {
	Text  string `json:"text"`  // example:`เหตุการณ์ส่งรายงาน`
	Value int64  `json:"value"` // example:`1`
}

type MethodList struct {
	Text  string `json:"text"`  // example:`SMTP send 1 group`
	Value int64  `json:"value"` // example:`1`
}

type TargetSelectOption struct {
	Condition []*selectoption.Option `json:"condition_list"`        // รายชื่อเงื่อนไขการส่ง
	Method    []*selectoption.Option `json:"method_list"`           // รายชื่อ method
	Lang      json.RawMessage        `json:"lang,omitempty"`        // ภาษา
	Group     []*GroupTargetColor    `json:"permission_group_list"` // รายชื่อกลุ่มที่มีสิทธิ
}

type ConditionSwagger struct {
	Text  string `json:"text"`  // example:`Match All condition`
	Value int64  `json:"value"` // example:`1`
}

type MethodSwagger struct {
	Text  string `json:"text"`  // example:`SMTP send 2 user`
	Value int64  `json:"value"` // example:`3`
}

type GroupTargetColorSW struct {
	Value int64  `json:"value"` //example:`4`
	Text  string `json:"text"`  //example:`กลุ่มผู้ใช้งานทั่วไป`
	Color string `json:"color"` //example:`FFFFFF`
}

type GroupTargetColor struct {
	Value interface{} `json:"value"` //example:`220`
	Text  interface{} `json:"text"`  //example:`สำหรับผู้ใช้งานระบบให้บริการข้อมูล`
	Color string      `json:"color"` //example:`FFFFFF`
}
