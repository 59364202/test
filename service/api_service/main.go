package api_service

import (
	"haii.or.th/api/server/model"
	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"

	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	// "haii.or.th/api/server/model/cronjob"
	model_metadata "haii.or.th/api/thaiwater30/model/metadata"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

const (
	DataServiceName = "thaiwater30/api_service"
	ServiceVersion  = service.APIVersion1
)

type HttpService struct {
}
type ApiService struct {
	Eid        string `json:"eid"`
	Id         int64  `json:"id"`
	MetadataId int64  `json:"mid"`

	FormDate string `json:"start_date"`
	ToDate   string `json:"end_date"`
}

func RegisterService(dpt service.Dispatcher) {
	srv := &HttpService{}
	dpt.Register(ServiceVersion, service.MethodGET, DataServiceName, srv.handleGetData)

	// cronjob
	// setting.SetSystemDefault("thaiwater30.service.api_service.cronRevmoeExpFile", "0 3 * * *")
	// cronjob.NewClusterFunc("thaiwater30.service.api_service.cronRevmoeExpFile", removeExpFile)
}

// Send post to laravel
func removeExpFile() error {
	url := "http://nhc.thaiwater.net/thaiwater30/data_service/check"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	var fw io.Writer
	var err error

	// add field
	if fw, err = w.CreateFormField("time"); err != nil {
		return err
	}
	time := int(setting.GetSystemSettingInt("service.api-service.download.ExpireDate") * 60 * 60)
	if _, err = fw.Write([]byte(strconv.Itoa(time))); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	// Check the response
	if res.StatusCode != http.StatusOK {
		return rest.NewError(res.StatusCode, res.Status, nil)
	}
	return nil
}

// handler get data
func (srv *HttpService) handleGetData(ctx service.RequestContext) error {
	media_url := ctx.BuildURL(0, "thaiwater30/shared/file", true)
	p := &ApiService{}
	if err := ctx.GetRequestParams(p); err != nil { // has eid ?
		return rest.NewError(404, "no eid", nil)
	}
	if p.MetadataId <= 0 { // has metadata id ?
		return rest.NewError(422, "invalid mid", nil)
	}
	_s_id, err := model.GetCipher().DecryptText(p.Eid)
	if err != nil { // decrypt eid
		return rest.NewError(422, "invalid eid", nil)
	}
	_, err = strconv.ParseInt(_s_id, 10, 64)
	if err != nil { // check _s_id is int
		return rest.NewError(422, "invalid eid", nil)
	}

	sd, err := model_order_detail.GetMetadataByEId(p.Eid)
	if err != nil {
		return err
	}
	p.Id = sd.OrderDetail_id

	if p.MetadataId != sd.Metadata_id { // check metadata_id
		return rest.NewError(422, "invalid mid", nil)
	}

	if !sd.IsEnabled.Bool { // service inactive
		return rest.NewError(422, "service inactive", nil)
	}

	ctx.LogRequestParams(p)

	if !sd.Agency_id.Valid || !sd.Connection_format.Valid || !sd.Table_name.Valid { // ยังไม่ได้ผูก dataset
		return rest.NewError(403, "this metadata no dataset_id ", nil)
	}

	if sd.Service_id == 2 { // เป็น download ให้คืน url ไป laravel
		// เช็ค ไฟล์ อายุเกิน .. วัน ให้บอกว่า file หมดอายุ
		expDate := sd.CreateAt.AddDate(0, 0, int(setting.GetSystemSettingInt("service.api-service.download.ExpireDate")))
		if time.Now().After(expDate) {
			return rest.NewError(410, "file expaired", nil)
		}
		// redirect to laravel
		_path := fmt.Sprintf("%v", sd.OrderDetail_id)
		ctx.ReplyRedirect("http://web.thaiwater.net/thaiwater30/api_service/"+base64.RawURLEncoding.EncodeToString([]byte(_path)), 0)
		return nil
	}
	if sd.Service_id == 1 { // เป็นข้อมูลย้อนหลัง เช็ค form_date, to_date ถ้าไม่มีให้ default ย้อนหลัง 3 วัน
		fd := time.Now().AddDate(0, 0, -2)
		//		sd.Detail_fromdate.String = fd.Format("2006-01-02")

		td := time.Now()
		//		sd.Detail_todate.String = td.Format("2006-01-02")

		if p.FormDate != "" {
			_fd, err := time.Parse("2006-01-02", p.FormDate)
			if err == nil {
				fd = _fd
				td = _fd
				//				sd.Detail_fromdate.String = fd.Format("2006-01-02")
			}
		}
		if p.ToDate != "" {
			_td, err := time.Parse("2006-01-02", p.ToDate)
			if err == nil {
				td = _td
				//				sd.Detail_todate.String = td.Format("2006-01-02")
			}
		}

		if int(td.Sub(fd).Hours())/24 > 2 {
			fd = td.AddDate(0, 0, -2)
			ctx.ReplyError(rest.NewError(422, "ช่วงวันที่เลือกต้องไม่เกิน 3 วัน", nil))
			return nil
		}
		sd.Detail_fromdate.String = fd.Format("2006-01-02")
		sd.Detail_todate.String = td.Format("2006-01-02")
	}

	row, err := model_order_detail.GetMetadataQueryResult(sd) // get query result
	if err != nil {
		return err
	}

	var data_media []*model_metadata.Struct_Data_Media
	var data *model_metadata.Struct_MetadataImportByAgency_Table

	if model_order_detail.IsMedia(sd.Table_name.String) {
		// scan media query result
		data_media, err = model_metadata.ScanData_Media(row, media_url)
		if err != nil {
			return rest.NewError(500, " Scan Media Data_ "+sd.Table_name.String+" err ", err)
		}
		ctx.ReplyJSON(data_media)
	} else {
		// scan data query result
		data, err = model_order_detail.ScanData(row)
		if err != nil {
			return rest.NewError(500, " ScanData_ "+sd.Table_name.String+" err ", err)
		}
		ctx.ReplyJSON(data.Data)
	}

	return nil
}
