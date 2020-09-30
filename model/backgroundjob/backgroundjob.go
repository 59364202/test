package backgroundjob

import (
	"haii.or.th/api/util/errors"
	"haii.or.th/api/util/eventcode"
	"haii.or.th/api/util/pqx"

	"database/sql"
)

// 	คิวรี่ดูว่ามี cmd รันอยู่รึป่าว
// 	โดยดูจากตัวล่าสุด
// 	ถ้า finish_at แสดงว่า มี process รันอยู่
//	Parameters:
//		cmd
//			command
//	Return:
//		true if cmd is running
func IsBgJobRunning(cmd string) (bool, error) {
	db, err := pqx.Open()
	if err != nil {
		return false, errors.NewEvent(eventcode.EventNetworkCriticalUnableConDB, err)
	}

	q := `SELECT finish_at FROM api.backgroundjob_log WHERE command = $1 ORDER BY id DESC LIMIT 1`
	p := []interface{}{cmd}
	var _finish_at sql.NullString
	var isBgjobRunning bool = false

	row, err := db.Query(q, p...)
	if err != nil {
		return false, errors.New(err.Error())
	}
	for row.Next() {
		err = row.Scan(&_finish_at)
		if err != nil {
			return false, errors.New(err.Error())
		}
		if !_finish_at.Valid {
			isBgjobRunning = true
		}
	}

	return isBgjobRunning, nil
}
