package check

import (
	"strconv"
)

//引数がすべて数字かのチェック
func OnlyNumber(args []string) bool {
	isNum := true
	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			isNum = false
			break
		}
	}
	return isNum
}
