package drawinfo

import (
	"fmt"
	"lottery/ultis"
	"math"
	"strconv"
)

type GameRaffleResult struct {
	TicketPrice      string
	NumberTicketSold int
	RoundNumber      int
	LotteryState     string
	JackpotWinning   string
	Eth_to_usd       float64
}
type GameDrawInfo struct {
	TicketPriceInUsd string
	TicketPriceInEth string
	JackpotInUsd     string
	JackpotInEth     string
	Draw             int
	Stage            string
	SellingTicket    int
}
type GameRaffleResponse struct {
	Status  int
	Message string
	Result  GameRaffleResult
}

var RaffleDrawInfoPath string = "https://qtl-performance-kyc-api.quanta.im/lottery/latest-round/info"

func FormatRaffleDrawInfo(rawData GameRaffleResult) GameDrawInfo {
	fmt.Printf("value %+v", rawData)
	var result GameDrawInfo
	jackpot, err := strconv.ParseFloat(rawData.JackpotWinning, 64)
	if err != nil {
		fmt.Printf("Parse jackpot err %s", err)
	}
	ticketPrice, errTickets := strconv.ParseFloat(rawData.TicketPrice, 64)
	if err != nil {
		fmt.Printf("Parse ticket price err %s", errTickets)
	}
	var weight float64 = math.Pow10(18)
	jackpotInEth := jackpot / weight
	ticketPriceInEth := ticketPrice / weight
	result.JackpotInEth = ultis.FormatFloat(jackpotInEth)
	result.JackpotInUsd = ultis.FormatFloat(jackpotInEth * rawData.Eth_to_usd)
	result.TicketPriceInEth = ultis.FormatFloat(ticketPriceInEth)
	result.TicketPriceInUsd = ultis.FormatFloat(ticketPriceInEth * rawData.Eth_to_usd)
	result.Draw = rawData.RoundNumber
	result.Stage = rawData.LotteryState
	return result
}
func GetRaffleDrawInfo() GameDrawInfo {
	response := new(GameRaffleResponse)
	ultis.GetJson(RaffleDrawInfoPath, response)
	return FormatRaffleDrawInfo(response.Result)
}
func GetAndFormatRaffleDrawInfo() string {
	var result GameDrawInfo = GetRaffleDrawInfo()
	return DrawInfoToString(result, "Raffle")
}
