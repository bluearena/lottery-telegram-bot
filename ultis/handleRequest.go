package ultis

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

type RoundInfoResponse struct {
	TicketPrice      string
	NumberTicketSold int
	RoundNumber      int
	LotteryState     string
	JackpotWinning   string
	Eth_to_usd       float64
}
type RoundInfoResult struct {
	TicketPriceInUsd string
	TicketPriceInEth string
	JackpotInUsd     string
	JackpotInEth     string
	Draw             int
	Stage            string
	SellingTicket    int
}
type Response struct {
	Status  int
	Message string
	Result  RoundInfoResponse
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func FormatDrawInfo(rawData RoundInfoResponse) RoundInfoResult {
	fmt.Printf("value %+v", rawData)
	var result RoundInfoResult
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
	result.JackpotInEth = strconv.FormatFloat(jackpotInEth, 'f', 6, 64)
	result.JackpotInUsd = strconv.FormatFloat(jackpotInEth*rawData.Eth_to_usd, 'f', 6, 64)
	result.TicketPriceInEth = strconv.FormatFloat(ticketPriceInEth, 'f', 6, 64)
	result.TicketPriceInUsd = strconv.FormatFloat(ticketPriceInEth*rawData.Eth_to_usd, 'f', 2, 64)
	result.Draw = rawData.RoundNumber
	result.Stage = rawData.LotteryState
	return result
}
func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
