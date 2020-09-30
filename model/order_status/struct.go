package order_status

import ()

type Struct_OrderStatus struct {
	Id           int64  `json:"id"`           // example:`1` รหัสของสถานะการขอข้อมูล
	Order_Status string `json:"order_status"` // example:`รับคำขอข้อมูล` สถานะการขอข้อมูล
}
