package data_service

import (
	"haii.or.th/api/thaiwater30/util/file"
	"haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/service"
	"path/filepath"
	"strconv"
	"time"
	//	"log"
	model "haii.or.th/api/server/model"
	model_order_detail "haii.or.th/api/thaiwater30/model/order_detail"
)

type UploadResult struct {
	LetterPath string `json:"path"`
}

type Struct_getUploadResult struct {
	Result string                                   `json:"result"` // example:`OK`
	Data   []*model_order_detail.Struct_OrderDetail `json:"data"`   // คำขอที่ยังไม่ได้รับผลคำขอจากหน่วยงานเฉพาะที่ส่งคำขอไปยังหน่วยงานเจ้าของข้อมูลแล้ว
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/upload_result
// @Summary			คำขอที่ยังไม่ได้รับผลคำขอจากหน่วยงาน
// @Description		คำขอที่ยังไม่ได้รับผลคำขอจากหน่วยงานเฉพาะที่ส่งคำขอไปยังหน่วยงานเจ้าของข้อมูลแล้ว
// @Method			GET
// @Produces		json
// @Response		200	Struct_getUploadResult successful operation

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/upload_result?
// @Summary			ไฟล์ผลอนุมัติคำขอ
// @Parameter		path	query	string	example:`OetIbuhGAomS26N6RB-bojBW-OmFGXX55ehGlxvuV-1vPzQC3CeGC-Raub6gWTfXddd7HERANFw2LYWt4rQ9SUMyTt3Y2HSCXO-_fWeCkAqp5ue5dK1U-J8wD7iUPVTj` encryptที่อยู่ไฟล์
// @Method			GET
// @Produces		Binary
// @Response		200 file successful operation
func (srv *HttpService) getUploadResult(ctx service.RequestContext) error {
	p := &UploadResult{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}

	if p.LetterPath != "" {
		fname, err := model.GetCipher().DecryptText(p.LetterPath)
		if err != nil {
			return rest.NewError(422, "invalid letterpath", err)
		}
		ctx.ReplyFile(fname)
	} else {
		rs, err := model_order_detail.GetOrderDetailGroupByAgency()
		if err != nil {
			ctx.ReplyError(err)
			return nil
		}
		ctx.ReplyJSON(result.Result1(rs))
	}
	return nil
}

type Struct_putUploadResult struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`lirfcr9UlRpMjt7FXqv55_NQ4sKKUjE--RUWVKuFlXZxv80hJIJqvX9N_z66oveGnfiqJ21N8MmfFgZ5hKEQqmyibZUMRofrdvOvCrjKs9Wk4z56uEnIfarRLy7ZkqPDirmuPsYWQaV4axJmU12sqA` encryptext path
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_service/upload_result
// @Summary			บันทึกเอกสารคำขอหน่วยงาน
// @Consumes		multipart/form-data
// @Method			PUT
// @Parameter		order_header_id	form	int64	example:`35` รหัสคำขอ
// @Parameter		agency_id	form	int64	example:`1` รหัสหน่วยงาน
// @Parameter		file	form	file ไฟล์ผลอนุมัติคำขอ(pdf)
// @Produces		json
// @Response		200	Struct_putUploadResult	successful operation
func (srv *HttpService) putUploadResult(ctx service.RequestContext) error {
	p := &model_order_detail.Param_OrderLetterPath{}
	if err := ctx.GetRequestParams(p); err != nil {
		return err
	}
	ctx.LogRequestParams(p)

	fn, err := ctx.GetUploadFile("file")
	if err != nil {
		return err
	}
	f, err := fn.GetFile()
	if err != nil {
		return err
	}

	// dataservice/agency_id/order_header_id
	var LetterPath = filepath.Join(file.DataserviceUploadLetterPath, strconv.FormatInt(p.Agency_Id, 10), strconv.FormatInt(p.Order_Header_Id, 10))
	var LetterFileName string = time.Now().Format("20060102") + ".pdf"
	err = file.SaveFile(f, LetterPath, LetterFileName)
	if err != nil {
		return err
	}
	p.Detail_Letterpath = filepath.Join(LetterPath, LetterFileName)

	err = model_order_detail.UpdaeLetterPath(p, ctx.GetUserID())
	if err != nil {
		ctx.ReplyJSON(result.Result0(err.Error()))
	} else {
		_detail_Letterpath, _ := model.GetCipher().EncryptText(filepath.Join(file.UploadPath, LetterPath, LetterFileName))
		ctx.ReplyJSON(result.Result1(_detail_Letterpath))
	}

	return nil
}
