package main

import (
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"

	"log"
	//	"strconv"
	//	"haii.or.th/api/util/pqx"
)

func TestApiService() error {
	var m_id int64 = 105
	p, err := model_order_detail.GetMetadataByMetadata(m_id)
	if err != nil {
		log.Println(err, " from FindExampleValueMedia ", m_id)
		return err
	}
	log.Println(p)
	p.Service_id = 4

	row, err := model_order_detail.GetMetadataQueryResult(p) // get query result
	if err != nil {
		log.Println(err, " from FindMediaTypeId ", m_id)
		return err
	}

	if err != nil {
		return err
	}

	var data_media []*model_metadata.Struct_Data_Media
	var data *model_metadata.Struct_MetadataImportByAgency_Table

	if model_order_detail.IsMedia(p.Table_name.String) {
		// scan media query result
		data_media, err = model_metadata.ScanData_Media(row, media_url)
		if err != nil {
			return err
		}
		for _, v := range data_media {
			log.Println(v)
		}
	} else {
		// scan data query result
		data, err = model_order_detail.ScanData(row)
		if err != nil {
			return err
		}
		for _, v := range data.Data {
			log.Println(v)
		}
	}

	return nil
}
