package buy

import (
	f "fmt"
)

func GenerateParticipation(cmd string) string {
	picks := ConverPickNumberToHex(GetPickNumberFromUrl(cmd))
	consecutive := IntToHex(GetConsecDrawFromUrl(cmd))
	picksLen := IntToHex(6)
	result := f.Sprintf("%s%s%s%s", HEX_FILL, picks, picksLen, consecutive)
	return f.Sprintf("0x%s", result[len(result)-32:len(result)])
}

func Buy649(command string) {
	f.Println(GenerateParticipation(command))
}
