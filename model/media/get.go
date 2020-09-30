// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package media is a model for public.media table. This table store media information.
package media

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"haii.or.th/api/server/model/setting"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/rest"
	"haii.or.th/api/util/thumbnail"

	"haii.or.th/api/thaiwater30/util/b64"
	udt "haii.or.th/api/thaiwater30/util/datetime"
	"haii.or.th/api/thaiwater30/util/validdata"

	model_setting "haii.or.th/api/server/model/setting"

	model_agency "haii.or.th/api/thaiwater30/model/agency"
	model_media_type "haii.or.th/api/thaiwater30/model/media_type"
)

// 	get media
//	Parameters:
//		param
//			Struct_Media_InputParam
//	Return:
//		Array Struct_Media
func GetMedia(param *Struct_Media_InputParam) ([]*Struct_Media, error) {

	err := checkInputParam(param)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Find datetime default format
	strDatetimeFormat := model_setting.GetSystemSetting("bof.Default.DatetimeFormat")
	if strDatetimeFormat == "" {
		strDatetimeFormat = model_setting.GetSystemSetting("setting.Default.DatetimeFormat")
	}

	//Convert AgencyID type from string to int64
	intAgencyId, err := strconv.ParseInt(param.AgencyID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Convert MediaTypeID type from string to int64
	intMediaTypeId, err := strconv.ParseInt(param.MediaTypeID, 10, 64)
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Open Database
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	//Variables
	var (
		resultData []*Struct_Media
		objMedia   *Struct_Media

		_id           sql.NullInt64
		_datetime     time.Time
		_path         sql.NullString
		_desc         sql.NullString
		_filename     sql.NullString
		_refer_source sql.NullString

		_media_type_id      sql.NullInt64
		_media_type_name    sql.NullString
		_media_subtype_name sql.NullString

		_agency_id        sql.NullInt64
		_agency_shortname sql.NullString
		_agency_name      sql.NullString

		_result *sql.Rows
	)

	//Query
	//log.Printf(sqlGetMedia + sqlGetMediaOrderBy, intAgencyId, intMediaTypeId, param.Start_date, param.End_date)
	_result, err = db.Query(sqlGetMedia+sqlGetMediaOrderBy, intAgencyId, intMediaTypeId, param.StartDate, param.EndDate)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer _result.Close()

	//Loop data result
	resultData = make([]*Struct_Media, 0)

	for _result.Next() {
		//Scan to execute query with variables
		err := _result.Scan(&_id, &_datetime, &_path, &_desc, &_filename, &_refer_source, &_media_type_id, &_media_type_name, &_media_subtype_name, &_agency_id, &_agency_shortname, &_agency_name)
		if err != nil {
			return nil, err
		}

		//strImagePath := model_setting.GetSystemSetting("setting.Default.ImagePath")

		if _agency_name.String == "" {
			_agency_name.String = "{}"
		}

		if _agency_shortname.String == "" {
			_agency_shortname.String = "{}"
		}

		//Generate MediaType object
		objMedia = &Struct_Media{}
		objMedia.ID = _id.Int64
		objMedia.Datetime = _datetime.Format(strDatetimeFormat)
		objMedia.Path = _path.String
		objMedia.Filename = _filename.String
		objMedia.ImageID, _ = b64.EncryptText(objMedia.Path + "," + objMedia.Filename)
		objMedia.Description = _desc.String
		objMedia.ReferSource = _refer_source.String

		_Struct_MediaType := &model_media_type.Struct_MediaType{}
		_Struct_MediaType.Id = _media_type_id.Int64
		_Struct_MediaType.Type_name = _media_type_name.String
		_Struct_MediaType.Subtype_name = _media_subtype_name.String
		_Struct_MediaType.Type_subtype_name = _media_type_name.String + " - " + _media_subtype_name.String
		objMedia.MediaType = _Struct_MediaType

		_Struct_Agency := &model_agency.Struct_Agency{}
		_Struct_Agency.Id = _agency_id.Int64
		_Struct_Agency.Agency_shortname = json.RawMessage(_agency_shortname.String)
		_Struct_Agency.Agency_name = json.RawMessage(_agency_name.String)
		objMedia.Agency = _Struct_Agency

		if _, err := os.Stat(path.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), objMedia.Path, objMedia.Filename)); os.IsNotExist(err) {
			objMedia.FileStatus = false
		} else {
			objMedia.FileStatus = true
		}

		resultData = append(resultData, objMedia)
	}

	return resultData, nil
}

//	validate input param
//	Parameters:
//		param
//			Struct_Media_InputParam
//	Return:
//		nil, error
func checkInputParam(param *Struct_Media_InputParam) error {

	//Check Parameters
	if param.AgencyID == "" {
		return errors.New("'agency_id' is not null.")
	}
	if param.MediaTypeID == "" {
		return errors.New("'media_type_id' is not null.")
	}
	if param.StartDate == "" {
		return errors.New("'start_date' is not null.")
	}
	if param.EndDate == "" {
		return errors.New("'end_date' is not null.")
	}

	return nil
}

// get upper wind lastest
//  Parameters:
//		None
//	Return:
//		[]MediaFileOutput
func GetUpperWindLatest() ([]*MediaFileOutput, error) {

	// define mediatype
	mediaType := []int64{10, 11, 12, 15, 16} //70, 71, 72

	data := make([]*MediaFileOutput, 0)

	// loop get media type by id
	for _, v := range mediaType {
		media, err := GetMediaLatestByMediaTypeID(v, 1)
		if err != nil {
			return nil, err
		}
		data = append(data, media...)
	}
	// reutnr media type data
	return data, nil
}

// get swat latest
//  Parameters:
//		None
//	Return:
//		[]MediaFileOutput
func GetSwatLatest() ([]*MediaFileOutput, error) {

	// define media type id
	mediaType := []int64{51, 57}

	data := make([]*MediaFileOutput, 0)
	// get media latest by media type id
	for _, v := range mediaType {
		media, err := GetMediaLatestByMediaTypeID(v, 1)
		if err != nil {
			return nil, err
		}
		data = append(data, media...)
	}
	media := make([]*MediaFileOutput, 0)
	// get first day and last day of month
	t := time.Now()
	d := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-1"
	dts, _ := time.Parse("2006-1-2", d)

	// get image latest 6 month
	dte := dts.AddDate(0, 1, -1)
	md, err := mediaSelectMonth(52, 1, dts, dte)
	if err != nil {
		return nil, err
	}
	media = append(media, md)
	for i := 0; i < 5; i++ {
		dts = dts.AddDate(0, 1, 0)
		dte = dts.AddDate(0, 1, -1)
		md, err := mediaSelectMonth(52, 1, dts, dte)
		if err != nil {
			return nil, err
		}
		media = append(media, md)
	}
	//	media, err := GetMediaLatestByMediaTypeID(52, 6)

	data = append(data, media...)

	return data, nil
}

// Get WaterSituation
//  Parameters:
//		None
//  Return:
//		[]MediaFileOutput
func GetWaterSituation() ([]*MediaFileOutput, error) {
	mediaType := []int64{74, 75, 76, 77}
	// get image by media type id
	return getMediaByMediaTypeIDAllDate(mediaType)
}

// 	get data by media type id all date
//	Parameters:
//		mediaTypeID
//			[]รหัสประเภทข้อมูลสื่อ
//	Return:
//		[]MediaFileOutput
func getMediaByMediaTypeIDAllDate(mediaTypeID []int64) ([]*MediaFileOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := sqlMediaOther
	p := []interface{}{}
	// loop make condition media type id
	if len(mediaTypeID) > 0 {
		var condition string
		for i, v := range mediaTypeID {
			if i != 0 {
				condition += " OR media_type_id=$" + strconv.Itoa(i+1)
			} else {
				condition = "media_type_id=$" + strconv.Itoa(i+1)
			}
			p = append(p, v)
		}
		q += " AND (" + condition + ")"
	}
	q += " ORDER BY m.media_datetime DESC"

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*MediaFileOutput, 0)

	for rows.Next() {
		var (
			mType     sql.NullString
			mSubType  sql.NullString
			mDatetime time.Time
			mPath     sql.NullString
			mDesc     sql.NullString
			mName     sql.NullString
		)
		dataRow := &MediaFileOutput{}
		// define image data output by row
		rows.Scan(&mType, &mSubType, &mDatetime, &mPath, &mDesc, &mName)
		dataRow.MediaType = mType.String
		dataRow.MediaSubType = mSubType.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")
		dataRow.Description = mDesc.String
		image, _ := b64.EncryptText(mPath.String + "," + mName.String)
		dataRow.URL = image
		dataRow.Filename = mName.String
		dataRow.Filepath = mPath.String
		filenameThumb := thumbnail.GetThumbName(mName.String, "", "")
		// check thumbnail from image
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), mName.String, filenameThumb)); err == nil {
			dataRow.URLThumb, _ = b64.EncryptText(filepath.Join(mPath.String, filenameThumb))
			dataRow.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
			dataRow.FilepathThumb = mPath.String
		}
		data = append(data, dataRow)
	}

	return data, nil

}

// get media latest by month
//  Parameters:
//		media_type_id
//				media_type_id
//		limit
//				limit sql
//		startDate
//				startDate get data
//		endDate
//				endDate get data
//  Return:
//		Array MediaFileOutput
func mediaSelectMonth(media_type_id int64, limit int64, startDate, endDate time.Time) (*MediaFileOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := sqlMediaLatestMonth + strconv.FormatInt(limit, 10)
	p := []interface{}{media_type_id, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")}

	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	dataRow := &MediaFileOutput{}

	for rows.Next() {
		var (
			mType     sql.NullString
			mSubType  sql.NullString
			mDatetime time.Time
			mPath     sql.NullString
			mDesc     sql.NullString
			mName     sql.NullString
		)

		rows.Scan(&mType, &mSubType, &mDatetime, &mPath, &mDesc, &mName)
		dataRow.MediaType = mType.String
		dataRow.MediaSubType = mSubType.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")
		dataRow.Description = mDesc.String
		// encrypt real path for get data
		image, _ := b64.EncryptText(mPath.String + "," + mName.String)
		dataRow.URL = image
		dataRow.Filename = mName.String
		dataRow.Filepath = mPath.String
		filenameThumb := thumbnail.GetThumbName(mName.String, "", "")
		// check thumbnail from filename
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), mName.String, filenameThumb)); err == nil {
			dataRow.URLThumb, _ = b64.EncryptText(filepath.Join(mPath.String, filenameThumb))
			dataRow.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
			dataRow.FilepathThumb = mPath.String
		}
	}

	return dataRow, nil
}

//  Parameters:
//		mediaTypeID
//				array media type id
//		limit
//			limit sql
//  Return:
//		Array MediaFileOutput
func GetMediaLatestByMediaTypeID(media_type_id int64, limit int64) ([]*MediaFileOutput, error) {

	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := sqlMediaLatest + strconv.FormatInt(limit, 10)

	p := []interface{}{media_type_id}

	//	fmt.Println(p)
	fmt.Println(q)
	rows, err := db.Query(q, p...)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	// media output
	data := make([]*MediaFileOutput, 0)

	for rows.Next() {
		var (
			mId           sql.NullInt64
			mType         sql.NullString
			mSubType      sql.NullString
			mCategory     sql.NullString
			mDatetime     time.Time
			mPath         sql.NullString
			mDesc         sql.NullString
			mName         sql.NullString
			mAnimPath     string
			mAnimFilename string
		)
		dataRow := &MediaFileOutput{}
		// media file output one row
		rows.Scan(&mId, &mType, &mSubType, &mCategory, &mDatetime, &mPath, &mDesc, &mName)
		dataRow.MediaTypeID = mId.Int64
		dataRow.MediaType = mType.String
		dataRow.MediaSubType = mSubType.String
		dataRow.MediaCategory = mCategory.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")
		dataRow.Description = mDesc.String
		// encrype real path
		image, _ := b64.EncryptText(mPath.String + "," + mName.String)
		dataRow.URL = image
		dataRow.Filename = mName.String
		dataRow.Filepath = mPath.String
		filenameThumb := thumbnail.GetThumbName(mName.String, "", "")
		full_path := filepath.Join(mPath.String, filenameThumb)

		// check thumbnail from filename
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
			dataRow.URLThumb, _ = b64.EncryptText(filepath.Join(mPath.String, filenameThumb))
			dataRow.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
			dataRow.FilepathThumb = mPath.String
		}

		// animation
		if mId.Int64 == 10 || mId.Int64 == 15 { // upper wind 0.6km, ressure 0.6km
			mAnimPath = "product/animation/upper_wind_600m/asia/haii/media"
			mAnimFilename = "upper_0.6km_large.mp4"
		} else if mId.Int64 == 11 || mId.Int64 == 16 { // upper wind 1.5km, pressure 1.5km
			mAnimPath = "product/animation/upper_wind_600m/asia/haii/media"
			mAnimFilename = "upper_1.5km_large.mp4"
		} else if mId.Int64 == 12 { // upper wind 5km
			mAnimPath = "product/animation/upper_wind_5000m/asia/haii/media"
			mAnimFilename = "upper_5.0km_large.mp4"
		}

		full_path_animation := filepath.Join(mAnimPath, mAnimFilename) // join path and filename
		dataRow.FilenameAnimation = mAnimPath
		dataRow.FilePathAnimation = mAnimFilename
		dataRow.MediaPathAnimation, _ = b64.EncryptText(full_path_animation) // encrypt path

		data = append(data, dataRow)
	}

	return data, nil
}

// precipitation rain history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//		datatype
//				datatype for select media type id
//  Return:
//		Array MediaFileOutput
func GetPrecipitationRainHistory(year, month, datatype string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time
	datatype = strings.ToLower(datatype)
	// select media type id
	switch datatype {
	case "animation thailand":
		return GetMediaLatestByMediaTypeID(80, 1)
	case "animation southeast asia":
		return GetMediaLatestByMediaTypeID(81, 1)
	case "animation asia":
		return GetMediaLatestByMediaTypeID(82, 1)
	case "asia":
		mediaTypeID = []int64{18}
	case "southeast asia":
		mediaTypeID = []int64{19}
	case "thailand":
		mediaTypeID = []int64{17}
	case "thailand basin":
		mediaTypeID = []int64{20}
	default:
		return nil, errors.New("No datatype")
	}

	// get first day and last day
	var err error
	if month != "" {
		firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	} else {
		firstDay, lastDay, err = getFirstDayAndLastDay(year, "")
	}
	if err != nil {
		return nil, err
	}
	agencyID := []int64{}
	// get media history and return
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 07:00", lastDay.Format("2006-01-02")+" 07:00")
}

// wave history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//		day
//				day for generate datetime
//		datatype
//				datatype for select media type id
//  Return:
//		Array MediaFileOutput
func GetWaveHistory(year, month, day, datatype string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time
	datatype = strings.ToLower(datatype)
	// select datatype and get mediatype id
	if datatype == "animation" {
		mediaTypeID = []int64{83}
	} else if datatype == "image" {
		mediaTypeID = []int64{13}
	} else {
		return nil, errors.New("No datatype")
	}
	agencyID := []int64{9}
	// get first day and last day
	var err error
	if day != "" {
		firstDay, err = time.Parse("2006-01-02", year+"-"+month+"-"+day)
		lastDay = firstDay
	} else {
		firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	}
	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59")
}

// vertical wind history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//  Return:
//		Array MediaFileOutput
func GetWind10mHistory(year, month, tinit_time string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time

	// select mediatype id
	mediaTypeID = []int64{182}
	agencyID := []int64{9}

	// get first day and last day
	var err error
	firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getMediaWind10mHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59", tinit_time)
}

//  rain accumulat history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//  Return:
//		Array MediaFileOutput
func GetRainAccumulatHistory(year, month string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time

	// select mediatype id
	mediaTypeID = []int64{184}
	agencyID := []int64{51}

	// get first day and last day
	var err error
	firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59")
}

// vertical wind history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//  Return:
//		Array MediaFileOutput
func GetVerticalWindHistory(year, month string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time

	// select mediatype id
	mediaTypeID = []int64{12}
	agencyID := []int64{9}

	// get first day and last day
	var err error
	firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59")
}

// upper wind and pressure 0.6 or 1.5 km history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//  Return:
//		Array MediaFileOutput
func GetUpperWindHistory(height, year, month string) ([]*MediaFileOutput, error) {
	var firstDay time.Time
	var lastDay time.Time
	var mediaTypeID []int64

	// select mediatype id
	if height == "1dot5km" {
		mediaTypeID = []int64{11, 16} // 11=upper wind 1.5km, 16=pressure 1.5km
	} else {
		mediaTypeID = []int64{10, 15} // 10=upper wind 0.6km asia, 15=pressure 0.6km asia
	}

	agencyID := []int64{9} // 9=haii

	// get first day and last day
	var err error
	firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59")
}

// swat history
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//		datatype
//				datatype for select media type id
//  Return:
//		Array MediaFileOutput
func GetSwatHistory(year, month, datatype string) ([]*MediaFileOutput, error) {
	var mediaTypeID []int64
	var firstDay time.Time
	var lastDay time.Time
	var err error
	// select mediatype id by datatype
	if datatype == "swat-w-forecast" {
		mediaTypeID = []int64{57}
		firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	} else if datatype == "swat-w-back" {
		mediaTypeID = []int64{51}
		firstDay, lastDay, err = getFirstDayAndLastDay(year, month)
	} else if datatype == "swat-m" {
		mediaTypeID = []int64{52}
		firstDay, lastDay, err = getFirstDayAndLastDay(year, "")
	} else {
		return nil, errors.New("No datatype")
	}
	if err != nil {
		return nil, err
	}
	agencyID := []int64{9}
	// get media history and return data
	return getMediaHistory(mediaTypeID, agencyID, firstDay.Format("2006-01-02")+" 00:00", lastDay.Format("2006-01-02")+" 23:59:59")
}

// func for get first day and last day in month
//  Parameters:
//		year
//				year for generate datetime
//		month
//				month for generate datetime
//  Return:
//		firstday of month and lastday of month
func getFirstDayAndLastDay(year, month string) (time.Time, time.Time, error) {
	var firstDay time.Time
	var lastDay time.Time
	var err error
	if month != "" {
		firstDay, err = time.Parse("2006-01-02", year+"-"+month+"-01")
		if err != nil {
			return firstDay, lastDay, err
		}
		lastDay = firstDay.AddDate(0, 1, -1)
	} else {
		firstDay, err = time.Parse("2006-01-02", year+"-01-01")
		if err != nil {
			return firstDay, lastDay, err
		}
		lastDay, err = time.Parse("2006-01-02", year+"-12-31")
		if err != nil {
			return firstDay, lastDay, err
		}
	}
	return firstDay, lastDay, nil
}

// func get media wind10m history
//  Parameters:
//		mediaTypeID
//				array media_type_id
//		agencyID
//				array agency id
//		dateStart
//				datestart for get data
//		dateEnd
//				dateend for get data
//  Return:
//		Array MediaFileOutput
func getMediaWind10mHistory(mediaTypeID []int64, agencyID []int64, dateStart, dateEnd string, tinit_time string) ([]*MediaFileOutput, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := sqlGetMediaHistory

	condition := " WHERE (media_datetime>=$1 AND media_datetime<=$2) AND (mt.deleted_at=to_timestamp(0) AND m.deleted_at=to_timestamp(0)) "

	p := []interface{}{dateStart, dateEnd}
	// var pText 	string
	count := 2

	// condition media type id
	if len(mediaTypeID) > 0 {
		var conditionMedia string
		for i, v := range mediaTypeID {
			count += 1
			if i != 0 {
				conditionMedia += " OR media_type_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " media_type_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
			// pText += strconv.Itoa(count) + "=" + strconv.FormatInt(v, 10) + ", "
		}
		condition += " AND (" + conditionMedia + ")"
	}

	// condition agency id
	if len(agencyID) > 0 {
		var conditionMedia string
		for i, v := range agencyID {
			count += 1
			if i != 0 {
				conditionMedia += " OR agency_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " agency_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
		}
		condition += " AND (" + conditionMedia + ")"
	}

	// swan criteria, check init date < datadate one day
	// For example, data date = 2019-01-18, created_at should be 2019-01-17
	if mediaTypeID[0] == 13 {
		condition += " AND DATE_PART('day', date(media_datetime)::timestamp - date(m.created_at)::timestamp) = 0"
	}

	if tinit_time != "" {
		condition += " AND (SUBSTRING(m.media_path FROM 47 FOR 2) = '" + tinit_time + "')"
	}

	c9 := true
	// sf := ""
	for _, v := range mediaTypeID {
		if v == 17 || v == 18 || v == 19 || v == 20 {
			c9 = false
			// sf = "07/"
		}
	}

	if !c9 {
		q = sqlGetMediaHistory2 + " WHERE (gs.date>=$1 AND gs.date<=$2) AND m.filename LIKE '%' || to_char(m.media_datetime, 'YYYYMMDD') || '%'"
	} else {
		q += condition + " ORDER BY media_datetime"
	}

	fmt.Println(p)
	fmt.Println(q)

	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*MediaFileOutput, 0)

	for rows.Next() {
		var (
			mTypeId   int64
			mType     sql.NullString
			mSubType  sql.NullString
			mDatetime time.Time
			mPath     sql.NullString
			mDesc     sql.NullString
			mName     sql.NullString
		)
		dataRow := &MediaFileOutput{}
		// media output
		rows.Scan(&mTypeId, &mType, &mSubType, &mDatetime, &mPath, &mDesc, &mName)
		dataRow.MediaTypeID = mTypeId
		dataRow.MediaType = mType.String
		dataRow.MediaSubType = mSubType.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")
		dataRow.Description = mDesc.String
		// encrypt text real path
		if mPath.Valid && mName.Valid {
			dataRow.URL, _ = b64.EncryptText(mPath.String + "/" + mName.String)
		}
		dataRow.Filename = mName.String
		dataRow.Filepath = mPath.String
		filenameThumb := thumbnail.GetThumbName(mName.String, "", "")
		full_path := filepath.Join(mPath.String, filenameThumb)

		// check file thumbnail
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
			dataRow.URLThumb, _ = b64.EncryptText(filepath.Join(mPath.String, filenameThumb))
			dataRow.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
			dataRow.FilepathThumb = mPath.String
		}

		// debug
		data = append(data, dataRow)
	}

	// return data
	return data, nil
}

// func get media history
//  Parameters:
//		mediaTypeID
//				array media_type_id
//		agencyID
//				array agency id
//		dateStart
//				datestart for get data
//		dateEnd
//				dateend for get data
//  Return:
//		Array MediaFileOutput
func getMediaHistory(mediaTypeID []int64, agencyID []int64, dateStart, dateEnd string) ([]*MediaFileOutput, error) {
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}
	q := sqlGetMediaHistory

	condition := " WHERE (media_datetime>=$1 AND media_datetime<=$2) AND (mt.deleted_at=to_timestamp(0) AND m.deleted_at=to_timestamp(0)) "

	p := []interface{}{dateStart, dateEnd}
	// var pText 	string
	count := 2

	// condition media type id
	if len(mediaTypeID) > 0 {
		var conditionMedia string
		for i, v := range mediaTypeID {
			count += 1
			if i != 0 {
				conditionMedia += " OR media_type_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " media_type_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
			// pText += strconv.Itoa(count) + "=" + strconv.FormatInt(v, 10) + ", "
		}
		condition += " AND (" + conditionMedia + ")"
	}

	// condition agency id
	if len(agencyID) > 0 {
		var conditionMedia string
		for i, v := range agencyID {
			count += 1
			if i != 0 {
				conditionMedia += " OR agency_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " agency_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
		}
		condition += " AND (" + conditionMedia + ")"
	}

	// swan criteria, check init date < datadate one day
	// For example, data date = 2019-01-18, created_at should be 2019-01-17
	if mediaTypeID[0] == 13 {
		condition += " AND DATE_PART('day', date(media_datetime)::timestamp - date(m.created_at)::timestamp) = 0"
	}

	c9 := true
	// sf := ""
	for _, v := range mediaTypeID {
		if v == 17 || v == 18 || v == 19 || v == 20 {
			c9 = false
			// sf = "07/"
		}
	}

	if !c9 {
		q = sqlGetMediaHistory2 + " WHERE (gs.date>=$1 AND gs.date<=$2) AND m.filename LIKE '%' || to_char(m.media_datetime, 'YYYYMMDD') || '%'"
	} else {
		q += condition + " ORDER BY media_datetime"
	}

//		fmt.Println(p)
//		fmt.Println(q)

	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}
	defer rows.Close()

	data := make([]*MediaFileOutput, 0)

	for rows.Next() {
		var (
			mTypeId   int64
			mType     sql.NullString
			mSubType  sql.NullString
			mDatetime time.Time
			mPath     sql.NullString
			mDesc     sql.NullString
			mName     sql.NullString
		)
		dataRow := &MediaFileOutput{}
		// media output
		rows.Scan(&mTypeId, &mType, &mSubType, &mDatetime, &mPath, &mDesc, &mName)
		dataRow.MediaTypeID = mTypeId
		dataRow.MediaType = mType.String
		dataRow.MediaSubType = mSubType.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")
		dataRow.Description = mDesc.String
		// encrypt text real path
		if mPath.Valid && mName.Valid {
			dataRow.URL, _ = b64.EncryptText(mPath.String + "/" + mName.String)
			// dataRow.URL, _ = b64.EncryptText(mPath.String + "/" + sf + mName.String)
		}
		dataRow.Filename = mName.String
		dataRow.Filepath = mPath.String
		filenameThumb := thumbnail.GetThumbName(mName.String, "", "")
		full_path := filepath.Join(mPath.String, filenameThumb)

		// check file thumbnail
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
			dataRow.URLThumb, _ = b64.EncryptText(filepath.Join(mPath.String, filenameThumb))
			dataRow.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
			dataRow.FilepathThumb = mPath.String
		}

		// debug
		// dataRow.DebugSql = q
		// dataRow.DebugParam = pText
		data = append(data, dataRow)
	}

	// return data
	return data, nil
}

//  Parameters:
//		radar_type
//				radar type
//		date
//				date for get data
//  Return:
//		Array Result_Struct_RadarHistory
func GetRadarHistory(radar_type, date string) ([]*Result_Struct_RadarHistory, error) {
	if radar_type == "" {
		return nil, rest.NewError(422, "no radar_type", nil)
	}
	if date == "" {
		return nil, rest.NewError(422, "no date", nil)
	}
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	var _setting = make([]map[string]interface{}, 0)
	var _radar_setting map[string]interface{} = nil
	json.Unmarshal(model_setting.GetSystemSettingJSON("Frontend.analyst.Radar.RadarTypeOrder"), &_setting)
	// หา setting ของ radar_type
	for _, v := range _setting {
		if v["radar_type"].(string) == radar_type {
			_radar_setting = v
		}
	}
	if _radar_setting == nil {
		// ไม่มี setting ของ radar_type นี้
		return nil, rest.NewError(422, "invalid radar_type ", nil)
	}
	// สร้าง sql ตาม radar_type, date, ความถี่ของ radar_type
	var q, p = Gen_SQL_RadarHistory(radar_type, date, fmt.Sprintf("%v", _radar_setting["radar_frequency"]))
	row, err := db.Query(q, p...)
	if err != nil {
		return nil, err
	}
	var (
		obj  *Result_Struct_RadarHistory
		data []*Result_Struct_RadarHistory = make([]*Result_Struct_RadarHistory, 0)

		data_radar []*Struct_RadarHistory
		o          *Struct_RadarHistory

		_tempCurrentHour int = -1
	)
	for row.Next() {
		var (
			_media_datetime time.Time
			_media_path     sql.NullString
			_filename       sql.NullString
		)
		err = row.Scan(&_media_datetime, &_media_path, &_filename)
		if err != nil {
			return nil, err
		}
		// ชั่วโมงเปลี่ยน ขึ้น array ใหม่
		if _tempCurrentHour != _media_datetime.Hour() {
			if _tempCurrentHour != -1 {
				obj.Data = data_radar // เก็บ frequency ของแต่ละ ชม.
			}
			obj = &Result_Struct_RadarHistory{}
			obj.DateTime = _media_datetime.Format(setting.GetSystemSetting("setting.Default.DateFormat")+" 15") + ":00" // 2006-01-02 15:00 fixed นาที ที่ 00
			obj.RadarName = _radar_setting["radar_name"].(string)
			obj.Timezone = _radar_setting["timezone"].(string)
			obj.Agency = _radar_setting["agency"].(string)
			// obj.DebugSql = q

			data_radar = make([]*Struct_RadarHistory, 0)

			data = append(data, obj)
			_tempCurrentHour = _media_datetime.Hour()
		}

		o = &Struct_RadarHistory{}
		o.MediaDatetime = _media_datetime.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))
		o.FilePath = validdata.ValidData(_media_path.Valid, _media_path.String)
		o.Filename = validdata.ValidData(_filename.Valid, _filename.String)
		if _media_path.Valid && _filename.Valid {
			o.MediaPath, _ = b64.EncryptText(filepath.Join(_media_path.String, _filename.String))
		}

		filenameThumb := thumbnail.GetThumbName(_filename.String, "", "")
		// filenameThumb := thumbnail.GetThumbName(_filename, "", "")
		full_path := filepath.Join(_media_path.String, filenameThumb)

		// thumb
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), full_path)); err == nil {
			o.MediaPathThumb, _ = b64.EncryptText(filepath.Join(_media_path.String, filenameThumb))
			o.FilenameThumb = strings.Replace(filenameThumb, "/", "", -1)
		}

		data_radar = append(data_radar, o)
	}
	obj.Data = data_radar
	return data, nil
}

type Struct_Media_sql struct {
	Media_datetime string
	Location       string
	Media_path     string
	Filename       string
}

type Struct_MediaLatest struct {
	Media_date string `json:"media_date"` //example:`2017-07-20` วันที่.
	Media_time string `json:"media_time"` //example:`04:03:00` เวลา
	Location   string `json:"location"`   //example:`cri`
	Media_file string `json:"media_file"` //example:`http://www.nhc.in.th/product/report/radar/phs/2017/07/20/phs240_201707200430.jpg` url ภาพ
}

func GetMediaLatest(media_url string) ([]*Struct_MediaLatest, error) {
	//Connect DB
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	//Query
	var q string = `
SELECT
    media_datetime
    , substring(filename from 1 for 3) as location
    , media_path
    , filename
FROM
    cache.latest_media
WHERE
    agency_id = 13
    AND media_type_id = 30
ORDER BY
    substring(filename from 1 for 3) ASC
	`

	//Query result
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	var rs []*Struct_MediaLatest = make([]*Struct_MediaLatest, 0)

	for rows.Next() {
		var (
			_media_datetime time.Time
			_location       sql.NullString
			_media_path     sql.NullString
			_filename       sql.NullString
		)

		err = rows.Scan(&_media_datetime, &_location, &_media_path, &_filename)
		if err != nil {
			return nil, err
		}
		encrypt_url, _ := b64.EncryptText(filepath.Join(_media_path.String, _filename.String))
		s := &Struct_MediaLatest{
			Media_date: _media_datetime.Format("2006-01-02"),
			Media_time: _media_datetime.Format("15:04:05"),
			Location:   _location.String,
			Media_file: media_url + "?image=" + encrypt_url,
		}

		rs = append(rs, s)

	}

	return rs, err
}

//  Parameters:
//		limit
//				record limit
//		agency_id
//				agency id
//  Return:
//		Array Struct_MediaReportHistory
func GetReportHistory(limit, agency_id string) ([]*PdfFileOutput, error) {
	var mediaTypeID []int64
	var agencyID []int64

	// convert string param to int64
	dataLimit, err := strconv.Atoi(limit)

	// set media type by agency
	if agency_id == "9" { // haii
		mediaTypeID = []int64{74, 75, 76, 77} // 74=daily, 75=weekly, 76=monthly, 77=yearly
		agencyID = []int64{9}
	} else {
		return nil, errors.New("No datatype")
	}

	if err != nil {
		return nil, err
	}

	// get media history and return data
	return getPdfHistory(mediaTypeID, agencyID, dataLimit)
}

// func get pdf history
//  Parameters:
//		mediaTypeID
//				array media_type_id
//		agencyID
//				array agency id
//		limit
//				data limit
//  Return:
//		Array MediaFileOutput
func getPdfHistory(mediaTypeID []int64, agencyID []int64, limit int) ([]*PdfFileOutput, error) {
	var condition string

	// connection
	db, err := pqx.Open()
	if err != nil {
		return nil, err
	}

	// setup sql
	q := sqlGetPdfHistory
	p := []interface{}{}
	count := 0

	// condition media type id
	if len(mediaTypeID) > 0 {
		var conditionMedia string
		for i, v := range mediaTypeID {
			count += 1
			if i != 0 {
				conditionMedia += " OR media_type_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " media_type_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
		}
		condition += " WHERE (" + conditionMedia + ")"
	}

	// condition agency id
	if len(agencyID) > 0 {
		var conditionMedia string
		for i, v := range agencyID {
			count += 1
			if i != 0 {
				conditionMedia += " OR agency_id=$" + strconv.Itoa(count)
			} else {
				conditionMedia = " agency_id=$" + strconv.Itoa(count)
			}
			p = append(p, v)
		}
		condition += " AND (" + conditionMedia + ")"
	}

	// replace $where by condition
	q = strings.Replace(q, "$where", condition, 1)
	q += " WHERE RANK <= " + strconv.Itoa(limit)

	// query
	rows, err := db.Query(q, p...)

	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	// close connection
	defer rows.Close()

	data := make([]*PdfFileOutput, 0)

	for rows.Next() {
		var (
			mAgencyId int64
			mTypeId   int64
			mPath     sql.NullString
			mFileName sql.NullString
			mDesc     sql.NullString
			mDatetime time.Time
		)
		dataRow := &PdfFileOutput{}

		// media output
		rows.Scan(&mAgencyId, &mTypeId, &mPath, &mFileName, &mDesc, &mDatetime)

		dataRow.AgencyID = mAgencyId
		dataRow.MediaTypeID = mTypeId
		dataRow.Filepath = mPath.String
		dataRow.Filename = mFileName.String
		dataRow.Description = mDesc.String
		dataRow.DateTime = udt.DatetimeFormat(mDatetime, "datetime")

		// encrypt text real path
		if mPath.Valid && mFileName.Valid {
			dataRow.URL, _ = b64.EncryptText(mPath.String + "/" + mFileName.String)
		}

		// dataRow.DebugSql = q
		data = append(data, dataRow)
	}

	// return data
	return data, nil
}
