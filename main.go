package main

import (
	"encoding/json"
	f "fmt"
	"log"
	"lottery/ultis"
	"net/http"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

var api_path = "https://qtl-performance-kyc-api.quanta.im/lottery/latest-round/info"
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	response := new(ultis.Response)
	bot, err := tgbotapi.NewBotAPI("552167461:AAEKBaMvqKSKjotbaWih_5HzK7Eqcqxc4MM")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			command := update.Message.Command()
			f.Println("update.Message.Command()", command)
			switch update.Message.Command() {
			case "drawinfo":
				var result ultis.RoundInfoResult
				ultis.GetJson(api_path, response)
				result = ultis.FormatDrawInfo(response.Result)
				f.Println("response %v", response)
				f.Println("data%v", result)
				textMsg := f.Sprintf("RAFFLE\nDraw:%d\nLotteryState:%s\nJackpotWinning(USD):%s\nTicketPrice(USD):%s", result.Draw, result.Stage, result.JackpotInUsd, result.TicketPriceInUsd)
				msg.Text = textMsg
			case "buy":
				msg.Text = "Enter your eth"
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Text {
			case "open":
				msg.Text = "Open keyboard"
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.Text = "Close keyboard"
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
			bot.Send(msg)
		}
	}

}
