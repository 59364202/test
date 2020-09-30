package order_detail_status

import ()

type Struct_OrderDetailStatus struct {
	Id            int64  `json:"id"`                      // example:`3` รหัสของสถานะการขอข้อมูล
	Detail_Status string `json:"detail_status,omitempty"` // example:`อนุมัติคำขอข้อมูล	` สถานะการขอข้อมูล
}
