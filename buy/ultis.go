package buy

import (
	f "fmt"
	"strconv"
	"strings"
)

var HEX_FILL = "00000000000000000000000000000000"

func GetPickNumberFromUrl(url string) [6]int {
	cmdArr := strings.Split(url, " ")
	var result [6]int
	if v := len(cmdArr); v > 1 {
		picks := strings.Split(cmdArr[1], "-")
		for i := 0; i < len(picks); i++ {
			numb, err := strconv.Atoi(picks[i])
			if err != nil {
				numb = 0
			}
			result[i] = numb
		}
	}
	return result
}
func GetConsecDrawFromUrl(url string) int {
	cmdArr := strings.Split(url, " ")
	strNum := ""
	if l := len(cmdArr); l > 1 {
		strNum = cmdArr[2]
	} else {
		strNum = "1"
	}
	num, err := strconv.Atoi(strNum)
	if err != nil {
		f.Println("GetConsecDrawFromUrl parse strNum err")
	}
	return num
}
func ConverPickNumberToHex(picks [6]int) string {
	result := ""
	for i := 0; i < len(picks); i++ {
		result = f.Sprintf("%s%s", result, IntToHex(picks[i]))
	}
	return result
}
func IntToHex(num int) string {
	conv := f.Sprintf("0%x", num)
	return conv[len(conv)-2 : len(conv)]
}
