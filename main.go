package main

import (
	"encoding/json"
	f "fmt"
	"log"
	"lottery/buy"
	"lottery/drawinfo"
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
				var formatParams ultis.DrawInfoParmas = ultis.FilterGameDrawInfo(update.Message.Text)
				var textMsg string
				if formatParams.Game == "649" {
					textMsg = drawinfo.GetAndFormat649DrawInfo()
				} else if formatParams.Game == "raffle" {
					textMsg = drawinfo.GetAndFormatRaffleDrawInfo()
				} else {
					textMsg = "We support 2 games Raffle and 649 now!"
				}
				msg.Text = textMsg
			case "buy649":
				buy.Buy649(update.Message.Text)
				msg.Text = "Enter your eth"
			case "status":
				msg.Text = "I'm ok."
			case "link":
				msg.Text = "http://zing.vn"
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
