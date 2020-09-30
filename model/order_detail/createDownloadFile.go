// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package metadata is a model for dataservice.order_detail table. This table store order_detail.
package order_detail

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	data "haii.or.th/api/server/model/eventlog"
	"haii.or.th/api/server/model/setting"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/log"
	"haii.or.th/api/util/rest"
)

//OnCreateDataService ...
//  หลังจาก post data_service
//	ให้เช็คว่ามี service ที่เป็น download หรือ cd/dvd
// 	ให้สร้างไฟล์ แล้วงส่งเมล์ อีกรอบ
func OnCreateDataService(mData *MailData, media_url string) {
	var hasDownload bool = false
	var hasCd bool = false
	var hasErr bool = false

	for _, d := range mData.Data {
		if d.Service_Id != 2 && d.Service_Id != 3 { // ไม่ใช่ download, cd/dvd ไม่ต้องทำอะไร
			continue
		}

		sd, err := GetMetadataByOd(d.Id)
		if !sd.Agency_id.Valid || !sd.Connection_format.Valid || !sd.Table_name.Valid || err != nil {
			d.IsSuccess = false
			continue // this metadata no dataset_id
		}

		row, err := GetMetadataQueryResult(sd)
		if err != nil {
			log.Locationf("order_id : %d, GetMetadataQueryResult err : %s", d.Id, err)
			d.IsSuccess = false
			d.ErrMsg = errors.Repack(err).Error()
			hasErr = true
			continue
		}

		var data_media []*model_metadata.Struct_Data_Media
		var data *model_metadata.Struct_MetadataImportByAgency_Table

		if IsMedia(sd.Table_name.String) {
			data_media, err = model_metadata.ScanData_Media(row, media_url)
		} else {
			data, err = ScanData(row)
		}

		if err != nil {
			d.IsSuccess = false
			d.ErrMsg = errors.Repack(err).Error()
			hasErr = true
			continue
		}
		strID := datatype.MakeString(d.Id)
		folderPath, err := GenerateFileFromData(strID, data, data_media) // สร้างไฟล์จาก ข้อมูล
		if err != nil {
			log.Locationf("order_id : %d, GenerateFileFromData err : %s", d.Id, err)
			d.ErrMsg = errors.Repack(err).Error()
			hasErr = true
			continue
		}
		log.Locationf("order_id : %d, folderPath : %s", d.Id, folderPath)
		if folderPath == "" {
			continue
		}
		d.FolderPath = folderPath
		if d.Service_Id == 2 {
			// นำไฟล์ ไปไว้บนเครื่องเว็ป
			err = UploadFileToWeb(folderPath, strID)
			if err != nil {
				// อัพไฟล์ไม่สำเร็จ อาจจะเกิดจาก limit upload size
				d.IsSuccess = false
				//  เก็บลง array ไว้ สำหรับแจ้ง admin ว่ามีรายการไหนที่อัพโหลดไม่สำเร็จ
				d.ErrMsg = errors.Repack(err).Error()
				hasErr = true
			} else {
				hasDownload = true // mark ไว้ว่ามี download อยู่ ต้องส่งเมลหลังทำทุกอย่างเสร็จ
				d.IsSuccess = true

			}
		} else if d.Service_Id == 3 {
			// นำไฟล์ไปเตรียมไว้สำหรับทำ cd/dvd
			// moveFileCdDvd(folderPath)
			hasCd = true
		}
	}
	// ลง eventlog เพื่อส่งเมล ให้ผู้ขอใช้บริการ
	if hasDownload { // มี download ส่งเมล
		data.LogSystemEvent(mData.ServiceId, mData.AgentUserId, mData.UserId, eventcode.EventDataServiceSendMail, "POST : data_service shopping after upload to laravel", mData)
	}
	if hasCd { // มี cd/dvd ต้อง email แจ้งผู้ดูแลเรื่อง shopping
		data.LogSystemEvent(mData.ServiceId, mData.AgentUserId, 0, eventcode.EventDataServiceSendMailToAdmin, "POST : data_service shopping after upload to laravel send email to admin", mData)
	}
	if hasErr { // มี err ต้องแจ้งผู้ดูแล
		data.LogSystemEvent(mData.ServiceId, mData.AgentUserId, 0, eventcode.EventDataServiceSendMailUploadErr, "POST : data_service upload to laravel error send email to admin", mData)
	}
}

//UploadFileToWeb ...
// send POST to laravel
func UploadFileToWeb(folderPath, folderName string) error {
	// upload to laravel
	url := "http://web.thaiwater.net/thaiwater30/data_service/upload" // url laravel

	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	var fw io.Writer
	var err error

	// add folder field
	if fw, err = w.CreateFormField("folder"); err != nil {
		return err
	}
	log.Locationf("add folder to body %s", folderName)
	if _, err = fw.Write([]byte(folderName)); err != nil {
		return err
	}

	// add all file in folder
	filepath.Walk(folderPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			// add file
			log.Locationf("add file to body %s", path)
			f, err := os.Open(path)
			if err != nil {
				log.Locationf("err add file to body %s", err)
				return err
			}
			defer f.Close()
			log.Locationf("CreateFormFile %s", path)
			fw, err = w.CreateFormFile("zip[]", path)
			if err != nil {
				log.Locationf("err CreateFormFile %s", err)
				return err
			}
			if _, err = io.Copy(fw, f); err != nil {
				log.Log(err)
				return err
			}
		}

		return nil
	})
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}
	//	log.Locationf("%s", vs)
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	log.Locationf("post upload file %s %s", folderPath, folderName)
	res, err := client.Do(req)
	if err != nil {
		log.Locationf(err.Error())
		return err
	}

	// Check the response
	log.Locationf("post upload file status code %s", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		responseData, _ := ioutil.ReadAll(res.Body)
		log.Locationf("resbody : %s", responseData)
		return rest.NewError(res.StatusCode, res.Status, nil)
	}
	// remove file and folder after post complete
	os.RemoveAll(folderPath)
	log.Log("post upload file remove all zip")
	return nil
}

// ย้ายที่เก็บสำหรับ service cd/dvd ไปไว้ตามที่ setting ไว้
func moveFileCdDvd(folderPath string) {
	newFolderPath := setting.GetSystemSetting("bof.order_detail.PrepareFileForCD")
	// move folder
	exec.Command("mv", folderPath, newFolderPath) // mv tempZipFile/200 /data/thaiwater/thaiwaterdata/test/tempCD
}
