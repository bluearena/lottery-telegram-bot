package ultis

import (
	"strconv"
	"strings"
)

type DrawInfoParmas struct {
	Game string
}

func FormatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', 6, 64)
}
func FilterGameDrawInfo(command string) DrawInfoParmas {
	var params DrawInfoParmas
	commandArr := strings.Split(command, " ")
	if v := len(commandArr); v == 1 {
		params.Game = "raffle"
	} else {
		params.Game = commandArr[1]
	}
	return params
}
