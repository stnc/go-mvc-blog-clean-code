package stnccollection

import (
	"strconv"
)

//FloatToString float 2 string
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 10, 64)
}

//StringToFloat  to convert a float number to a string
func StringToFloat(str string) (returnData float64, err2 error) {
	/* bu float32 yapar
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s) // 3.1415927410125732
	}
	*/
	//  var returnData float64
	// var err2 error

	returnData, err2 = strconv.ParseFloat(str, 64)
	return returnData, err2

}

//FloatToString10 float 2 string
func FloatToString10(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 10, 64)
}

//FloatToString6 float 2 string
func FloatToString6(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

/*
https://selmantunc.com.tr/post/635793192618442752/golang-numeric-conversions

Atoi (string to int)
i, err := strconv.Atoi(“-42”)

——————————-

Itoa (int to string).
s := strconv.Itoa(-42)

  ———————

int64 to string
str:= strconv.FormatInt(int64(165), 10)

——————————-

uint64 to string
lastID := strconv.FormatUint(uint64(5656556666), 10)

——————————–

string to  uint64
catID, _ := strconv.ParseUint(“string”, 10, 64)

interface return to string
session.Get(key).(string)
*/

//Uint64toString uint64 2 string
func Uint64toString(inputNum uint64) string {
	return strconv.FormatUint(uint64(inputNum), 10)
}

//StringtoUint64 string 2 uint64
func StringtoUint64(inputStr string) (uintInt uint64) {
	uintInt, _ = strconv.ParseUint(inputStr, 10, 64)
	return uintInt
}
