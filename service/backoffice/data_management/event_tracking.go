package data_management

import (
	model_event_tracking "haii.or.th/api/thaiwater30/model/event_tracking"
	result "haii.or.th/api/thaiwater30/util/result"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/service"
)

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_option_list
// @Method			GET
// @Summary			ตัวเลือกสำหรับแสดงผลหน้าติดตามเหตุการณ์
// @Description		data event tracking option
// @Produces		json
// @Response		200		EventTrackingSelectOptionList successful operation
// @Response		404			-		the request service name was not found
type ResultNull struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`null`
}
type EventTrackingSelectOptionList struct {
	Result string                                             `json:"result"` // example:`OK`
	Data   model_event_tracking.EventTrackingSelectOptionList `json:"data"`
}

func (srv *HttpService) getSelectOptionEventTracking(ctx service.RequestContext) error {
	page := ctx.GetServiceParams("id")
	dataResult, err := model_event_tracking.GetEventTrackingSelectOption(page)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_option_list_invalid_data
// @Method			GET
// @Summary			ตัวเลือกสำหรับหน้าติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Description		ตัวเลือกสำหรับหน้าติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Produces		json
// @Response		200		EventTrackingSelectOptionInvalidDataSwagger successful operation

type EventTrackingSelectOptionInvalidDataSwagger struct {
	Result string                                                      `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventTrackingSelectOptionInvalidData `json:"data"`
}

func (srv *HttpService) getSelectOptionEventTrackingInvalidData(ctx service.RequestContext) error {

	dataResult, err := model_event_tracking.GetEventTrackingDownloadInvalidDataSelectOption()
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking
// @Method			GET
// @Summary			ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์
// @Description		ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์
// @Parameter		-	query	model_event_tracking.EventTrackingOption
// @Produces		json
// @Response		200		EVTrackingSwagger successful operation

type EVTrackingSwagger struct {
	Result string                                   `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventTrackingData `json:"data"`
}

func (srv *HttpService) getEventTracking(ctx service.RequestContext) error {

	p := &model_event_tracking.EventTrackingOption{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetEventTrackingData(p.DateStart, p.DateEnd, p.Agency, p.EventType, p.EventSubType, p.SolveEvent)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_solve
// @Method			GET
// @Summary			ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าแก้ไขปัญหา
// @Description		ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าแก้ไขปัญหา
// @Parameter		-	query	EventTrackingOptionSwaggerTrackingSolve
// @Produces		json
// @Response		200		EventTrackingSolveSwagger successful operation

type EventTrackingOptionSwaggerTrackingSolve struct {
	DateStart    string  `json:"date_start"`     // example:`2017-08-24 00:00` วันที่เริ่มต้น เช่น  2017-08-24 00:00
	DateEnd      string  `json:"date_end"`       // example:`2017-08-24 23:59` วันที่สิ้นสุด เช่น  2017-08-24 23:59
	Agency       []int64 `json:"agency"`         // example:`[9]` รหัสหน่วยงาน เช่น [9]
//	EventType    []int64 `json:"event_type"`     // example:`[4]` รหัสเหตุการณ์ เช่น [4]
//	EventSubType []int64 `json:"event_sub_type"` // example:`[42]` รหัสเหตุการณ์ย่อย เช่น [42]
	SolveEvent   bool    `json:"solve_event"`    // example:`false` แก้ปัญหาเหตุการณ์
}

type EventTrackingSolveSwagger struct {
	Result string                                   `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventTrackingData `json:"data"`
}

func (srv *HttpService) getSolveEventTracking(ctx service.RequestContext) error {

	p := &model_event_tracking.EventTrackingOption{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetEventTrackingSolveData(p.DateStart, p.DateEnd, p.Agency, p.EventType, p.EventSubType)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_download_invalid_data
// @Method			GET
// @Summary			ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าดาวน์โหลดข้อมูลผิดพลาด
// @Description		ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าดาวน์โหลดข้อมูลผิดพลาด
// @Parameter		-	query	model_event_tracking.EventInvalidDataOption
// @Produces		json
// @Response		200		DLEventTrackingInvalidDataSwagger successful operation

type DLEventTrackingInvalidDataSwagger struct {
	Result string                                  `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventInvalidData `json:"data"`
}

func (srv *HttpService) getDownloadEventInvalidData(ctx service.RequestContext) error {

	p := &model_event_tracking.EventInvalidDataOption{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetDownloadEventInvalidData(p.Agency, p.Date)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_send_invalid_data
// @Method			GET
// @Summary			ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าส่งข้อมูลผิดพลาด
// @Description		ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าส่งข้อมูลผิดพลาด
// @Parameter		-	query	model_event_tracking.EventSendInvalidDataOption
// @Produces		json
// @Response		200		SendInvalidDataSwagger successful operation
type SendInvalidDataSwagger struct {
	Result string                                      `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventSendInvalidData `json:"data"`
}

func (srv *HttpService) getSendEventInvalidData(ctx service.RequestContext) error {

	p := &model_event_tracking.EventSendInvalidDataOption{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetSendEventInvalidData(p.Agency, p.DateStart, p.DateEnd)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_invalid_data
// @Method			GET
// @Summary			ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Description		ข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Parameter		-	query	model_event_tracking.EventSendInvalidDataOption
// @Produces		json
// @Response		200		EVTrackingInvalidDataSwagger successful operation

type EVTrackingInvalidDataSwagger struct {
	Result string                                      `json:"result"` //example:`OK`
	Data   []model_event_tracking.EventSendInvalidData `json:"data"`
}

func (srv *HttpService) getTrackingEventInvalidData(ctx service.RequestContext) error {

	p := &model_event_tracking.EventSendInvalidDataOption{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetTrackingEventInvalidData(p.Agency, p.DateStart, p.DateEnd)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_solve
// @Method			PUT
// @Summary			อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าแก้ไขข้อมูลผิดพลาด
// @Description		อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าแก้ไขข้อมูลผิดพลาด
// @Consumes		json
// @Parameter		-	form	model_event_tracking.EventTrackingUpdate
// @Produces		json
// @Response		200		ResultPutEventTracking successful operation
type ResultPutEventTracking struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`Update EventLog Solve Successful`
}

func (srv *HttpService) putEventTracking(ctx service.RequestContext) error {

	p := &model_event_tracking.EventTrackingUpdate{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.UpdateEventTracking(p.EventLogID, p.Message, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_send_invalid_data
// @Method			PUT
// @Summary			อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าส่งข้อมูลผิดพลาด
// @Description		อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์หน้าส่งข้อมูลผิดพลาด
// @Consumes		json
// @Parameter		-	body	model_event_tracking.EventTrackingUpdate
// @Produces		json
// @Response		200		ResultSendEventInvalidData successful operation
type ResultSendEventInvalidData struct {
	Result string `json:"result"` // example:`OK`
	Data   string `json:"data"`   // example:`Update EventLog Send Error At Successful`
}

func (srv *HttpService) putSendEventInvalidData(ctx service.RequestContext) error {

	p := &model_event_tracking.EventSendTrackingUpdate{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.UpdateSendEventInvalidData(p.EventLogID, p.Date, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_tracking_invalid_data
// @Method			PUT
// @Consumes		json
// @Summary			อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Description		อัพเดทข้อมูลสำหรับแสดงในตารางติดตามเหตุการณ์ข้อมูลผิดพลาด
// @Parameter		-	body	model_event_tracking.EventTrackingUpdate
// @Produces		json
// @Response		200		ResultTrackingEventInvalidData successful operation

type ResultTrackingEventInvalidData struct {
	Result string      `json:"result"` // example:`OK`
	Data   interface{} `json:"data"`   // example:`Update EventLog Solve Successful`
}

func (srv *HttpService) putTrackingEventInvalidData(ctx service.RequestContext) error {

	p := &model_event_tracking.EventTrackingUpdate{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.UpdateEventTrackingInvalidData(p.EventLogID, p.Message, ctx.GetUserID())
	if err != nil {
		ctx.ReplyError(err)
	} else {
		ctx.ReplyJSON(result.Result1(dataResult))
	}

	return nil
}

// @DocumentName	v1.webservice
// @Service			thaiwater30/backoffice/data_management/event_file_csv
// @Method			GET
// @Summary			ดาวน์โหลด ไฟล์ csv ของข้อมูลที่ผิดพลาด
// @Description		ดาวน์โหลด ไฟล์ csv ของข้อมูลที่ผิดพลาด
// @Parameter		event_log_id	query	int	รหัส eventlog เช่น 672296
// @Produces		Binary
// @Response		200		file successful operation

func (srv *HttpService) getFileInvalidData2(ctx service.RequestContext) error {

	p := &model_event_tracking.EventInvalidData{}
	if err := ctx.GetRequestParams(p); err != nil {
		return errors.Repack(err)
	}
	dataResult, err := model_event_tracking.GetFileInvalidData2(ctx.GetUserID(), p.EventLogId)
	if err != nil {
		ctx.ReplyError(err)
	} else {
		if dataResult != "" {
			ctx.ReplyFile(dataResult)
		} else {
			ctx.ReplyJSON(result.Result0(dataResult))
		}
	}

	return nil
}
