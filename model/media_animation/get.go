// Copyright 2016 Hydro and Agro Informatics Institute <www.haii.or.th>.
// All rights reserved. Use of this source code is governed by HAII license.
//     Author: CIM Systems (Thailand) <cim@cim.co.th>
//

// Package media_animation is a model for public.media_animation table. This table store media_animation information.
package media_animation

import (
	"database/sql"

	//	"log"
	"os"
	"path/filepath"

	"haii.or.th/api/server/model/setting"

	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/thumbnail"

	"haii.or.th/api/thaiwater30/util/b64"
	"haii.or.th/api/thaiwater30/util/sqltime"
)

//	scan query
//	Parameters:
//		strSQL
//			sql query
//	Return:
//		Array Struct_Media
func scanSql(strSQL string) ([]*Struct_Media, error) {
	//	log.Println(strSQL)
	db, err := pqx.Open()
	if err != nil {
		return nil, errors.Repack(err)
	}

	var (
		obj *Struct_Media

		_filepath   sql.NullString
		_filename   sql.NullString
		_media_time sqltime.NullTime

		_filename_thumb string
	)
	data := make([]*Struct_Media, 0)

	row, err := db.Query(strSQL)
	if err != nil {
		return nil, pqx.GetRESTError(err)
	}

	for row.Next() {
		err = row.Scan(&_filepath, &_filename, &_media_time)
		if err != nil {
			return nil, err
		}

		obj = &Struct_Media{}
		//		obj.Path, _ = model.GetCipher().EncryptText(filepath.Join(_filepath.String, _filename.String))
		obj.Path, _ = b64.EncryptText(filepath.Join(_filepath.String, _filename.String))
		obj.Filename = _filename.String
		obj.FilePath = _filepath.String
		obj.Datetime = _media_time.Time.Format(setting.GetSystemSetting("setting.Default.DatetimeFormat"))

		_filename_thumb = thumbnail.GetThumbName(_filename.String, "", "")
		if _, err := os.Stat(filepath.Join(setting.GetSystemSetting("server.service.dataimport.DataPathPrefix"), _filepath.String, _filename_thumb)); err == nil {
			//			obj.PathThumb, _ = model.GetCipher().EncryptText(filepath.Join(_filepath.String, _filename_thumb))
			obj.PathThumb, _ = b64.EncryptText(filepath.Join(_filepath.String, _filename_thumb))
			obj.FilenameThumb = _filename_thumb[1:]
		}

		obj.Dt = _media_time.Time

		data = append(data, obj)
	}
	return data, nil
}

//  animation ของ คาดการ์ณฝน
//	Return:
//		Array Struct_Media
func GetPreRainAnimation() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreRainAnimation)
}

//  animation ของ คาดการ์ณลม
//	Return:
//		Array Struct_Media
func GetPreWindAnimation() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreWindAnimation)
}

//  animation ของ คาดการ์ณคลื่น
//	Return:
//		Array Struct_Media
func GetPreWaveAnimation() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreWaveAnimation)
}

//  animation ของ คาดการ์ณคลื่น
//	Return:
//		Array Struct_Media
func GetPreWaveAnimationMp4() ([]*Struct_Media, error) {
	return scanSql(SQL_SelectPreWaveAnimationMP4)
}
