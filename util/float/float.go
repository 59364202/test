package float

import (
	"haii.or.th/api/util/datatype"

	"github.com/dustin/go-humanize"

	"math"
	"strings"
)

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}

func TwoDigit(num float64) float64 {
	return ToFixed(num, 2)
}
func OneDigit(num float64) float64 {
	return ToFixed(num, 1)
}
func NoDigit(num float64) float64 {
	return ToFixed(num, 0)
}

//	remove trailing 0 from string
func RemoveTrialZero(str string) string {
	if str == "0" {
		return str
	}
	if !strings.Contains(str, ".") {
		// ไม่มี ทศนิยม ไม่ต้องทำอะไร
		return str
	}
	return strings.TrimRight(strings.TrimRight(str, "0"), ".")
}

//	number to string with max decimals length
func String(v interface{}, maxDigit int) string {
	str := StringTrailingZero(v, maxDigit)
	//	return str
	return RemoveTrialZero(str)
}

//	number to string with trailing 0
func StringTrailingZero(v interface{}, maxDigit int) string {
	f := ToFixed(datatype.MakeFloat(v), maxDigit)
	return datatype.MakeString(f)
}

//	number to string with comma
func Comma(v interface{}, maxDigit int) string {
	str := CommaTrailingZero(v, maxDigit)
	return RemoveTrialZero(str)
}

//	number to string with comma trailing 0
func CommaTrailingZero(v interface{}, digit int) string {
	if digit < 0 {
		digit = 0
	}

	//	f := ToFixed(datatype.MakeFloat(v), digit)
	f := datatype.MakeFloat(v)
	switch digit {
	case 0:
		f = NoDigit(f)
	case 1:
		f = OneDigit(f)
	case 2:
		f = TwoDigit(f)
	}

	return humanize.Commaf(f)
}
