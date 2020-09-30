package setting

import (
	"haii.or.th/api/util/datatype"
	"haii.or.th/api/util/errors"

	//	"log"
	"strings"
)

//  เปรียบเทียบค่า ของ value ตาม opeartion, term
//	Parameters:
//		value
//			ค่าที่ต้องการเปรี่ยบเทียบ
//		opeartion
//			เงื่อนไขในการเปรียบเทียบ ex. >, >=, <
//		term
//			ค่าที่ใช้เปรียบเทียบ
//	Return:
//		true, false
func Compare(value interface{}, opearion string, term ...interface{}) (bool, error) {
	// แปลง value เป็น float64
	v, b := datatype.ToFloat(value)
	if b == false {
		return false, errors.New("cannot convert to float")
	}
	opearion = strings.ToLower(opearion)
	//	log.Println(v, opearion, term)

	termFloat := make([]float64, 0)
	for _, t := range term {
		f := datatype.MakeFloat(t)
		termFloat = append(termFloat, f)
	}
	rs := false
	// compare
	switch opearion {
	case ">":
		rs = v > termFloat[0]
	case ">=":
		rs = v >= termFloat[0]
	case "<":
		rs = v < termFloat[0]
	case "<=":
		rs = v <= termFloat[0]
	case "!=":
		rs = v != termFloat[0]
	case "=":
		rs = v == termFloat[0]
	case "between":
		rs = v >= termFloat[0] && v <= termFloat[1]
	}
	return rs, nil
}
