package upload

import (
	"mime/multipart"
	"time"

	"haii.or.th/api/thaiwater30/model/datafile"
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/pqx"
)

func UploadToTemp(filename string, file multipart.File) (int64, string, error) {

	defer file.Close()
	path := "temp/" + time.Now().Format("20060102150405") + "/"
	fpname := path + filename
	datafile.CopyFile(path, filename, file)

	// q := `INSERT INTO tempfile2 (path) VALUES ($1) Returning id`

	tx, err := pqx.Open()
	if err != nil {
		return 0, "", errors.Repack(err)
	}
	// defer tx.Rollback()

	var fileID int64
	const q = `INSERT INTO tempfile2 (path) VALUES ($1) Returning id`
	err = tx.QueryRow(q, filename).Scan(&fileID)

	if err != nil {
		return 0, "", errors.Repack(err)
	}

	stmt, err := tx.Prepare(q)
	if err != nil {
		return 0, "", errors.Repack(err)
	}

	p := []interface{}{fpname}

	_, err = stmt.Exec(p...)
	if err != nil {
		return 0, "", errors.Repack(err)
	}

	// var id int64
	// id, err = rs.LastInsertId()

	// tx.Commit()
	if err != nil {
		return 0, "", errors.Repack(err)
	}
	return fileID, fpname, nil
}
