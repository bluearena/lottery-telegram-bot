package drawinfo

import (
	f "fmt"
)

func DrawInfoToString(data GameDrawInfo, game string) string {
	return f.Sprintf("%s\nDraw:%d\nLotteryState:%s\nJackpotWinning(USD):%s\nTicketPrice(USD):%s", game, data.Draw, data.Stage, data.JackpotInUsd, data.TicketPriceInUsd)
}
