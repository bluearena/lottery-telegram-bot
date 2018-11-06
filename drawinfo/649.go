package drawinfo

import (
	f "fmt"
	"lottery/ultis"
	"math"
	"strconv"
)

type Game649Result struct {
	ColumnPrice        string
	TotalTicketSold    int
	Draw               int
	Stage              string
	JackpotAmount      string
	JackpotAmountInUSD string
	FiatRate           string
}

type Game649Response struct {
	Status  int
	Message string
	Result  Game649Result
}

var PERCENT_RATE float64 = 100000000
var G649DrawInfoPath string = "https://qtl-performance-kyc-api.quanta.im/lotto/api?game=lotto-6-49&module=draw&action=info&draw=latest"

func Format649DrawInfo(rawData Game649Result) GameDrawInfo {
	var result GameDrawInfo
	jackpot, err := strconv.ParseFloat(rawData.JackpotAmount, 64)
	if err != nil {
		f.Printf("Parse jackpot err %s", err)
	}
	ticketPrice, errTickets := strconv.ParseFloat(rawData.ColumnPrice, 64)
	if err != nil {
		f.Printf("Parse ticket price err %s", errTickets)
	}
	fiatRate, errFiat := strconv.ParseFloat(rawData.FiatRate, 64)
	if err != nil {
		f.Printf("Parse fiatRate to float err %s", errFiat)
	}
	f.Printf("get g649Draw %+v", rawData)
	ethToUsd := fiatRate / PERCENT_RATE
	var weight float64 = math.Pow10(18)
	jackpotInEth := jackpot / weight
	ticketPriceInEth := ticketPrice / weight
	result.JackpotInEth = ultis.FormatFloat(jackpotInEth)
	result.JackpotInUsd = ultis.FormatFloat(jackpotInEth * ethToUsd)
	result.TicketPriceInEth = ultis.FormatFloat(ticketPriceInEth)
	result.TicketPriceInUsd = ultis.FormatFloat(ticketPriceInEth * ethToUsd)
	result.SellingTicket = rawData.TotalTicketSold
	result.Draw = rawData.Draw
	result.Stage = rawData.Stage
	return result
}
func Get649DrawInfo() GameDrawInfo {
	response := new(Game649Response)
	ultis.GetJson(G649DrawInfoPath, response)
	return Format649DrawInfo(response.Result)
}
func GetAndFormat649DrawInfo() string {
	var result GameDrawInfo = Get649DrawInfo()
	return DrawInfoToString(result, "649")
}
