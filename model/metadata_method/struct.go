package metadata_method

import ()

type MetadataMethodParams struct {
	ID         int64  `json:"id"`                   // example:`1` รหัสวิธีการได้มาของข้อมูล
	MethodID   string `json:"metadata_method_id"`   // example:`tci3CPw5YTdhuV9PuN8oh3e-I9GySj9SZmpYRIbct4iUiSSjyOqsVPewcVLgVHgDDtmzcX35dU5bkjt4qtzD_g` รหัสวิธีการได้มาของข้อมูล
	MethodName string `json:"metadata_method_name"` // example:`Web Service` ชื่อวิธีการได้มาซึ่งข้อมูล
}
