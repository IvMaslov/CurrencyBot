package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	goenv "github.com/joho/godotenv"
)

var currencyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("USD"),
		tgbotapi.NewKeyboardButton("EUR"),
		tgbotapi.NewKeyboardButton("GBP"),
		tgbotapi.NewKeyboardButton("Crypto"),
	),
)

var cryptoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("BTC"),
		tgbotapi.NewKeyboardButton("ETH"),
	),
)

func main() {
	err := goenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if command := update.Message.Command(); command == "start" {
			msg.Text = "Привет, этот бот будет сообщает курс основных валют и криптовалют"
			msg.ReplyMarkup = currencyKeyboard
			bot.Send(msg)
			continue
		}
		if command := update.Message.Command(); command == "open" {
			msg.Text = "Открываю..."
			msg.ReplyMarkup = currencyKeyboard
			bot.Send(msg)
			continue
		}
		if text := update.Message.Text; text == "BTC" || text == "ETH" {
			data, err := GetCrypto(text)
			if err != nil {
				msg.Text = "Что-то пошло не так, попробуйте позже"
				msg.ReplyMarkup = currencyKeyboard
				bot.Send(msg)
				continue
			}
			msg.Text = "Курс " + text + ": " + data.Price
			msg.ReplyMarkup = currencyKeyboard
			bot.Send(msg)
			continue
		}

		switch update.Message.Text {
		case "USD":
			data, err := GetCurrency("usd")
			if err != nil {
				msg.Text = "Что-то пошло не так, попробуйте позже"
				bot.Send(msg)
				continue
			}
			msg.Text = "Доллар к рублю: " + fmt.Sprintf("%f", data.Price) + " | Обновлено: " + data.Date
			bot.Send(msg)
		case "EUR":
			data, err := GetCurrency("eur")
			if err != nil {
				msg.Text = "Что-то пошло не так, попробуйте позже"
				bot.Send(msg)
				continue
			}
			msg.Text = "Евро к рублю: " + fmt.Sprintf("%f", data.Price) + " | Обновлено: " + data.Date
			bot.Send(msg)
		case "GBP":
			data, err := GetCurrency("gbp")
			if err != nil {
				msg.Text = "Что-то пошло не так, попробуйте позже"
				bot.Send(msg)
				continue
			}
			msg.Text = "Фунт к рублю: " + fmt.Sprintf("%f", data.Price) + " | Обновлено: " + data.Date
			bot.Send(msg)
		case "Crypto":
			msg.ReplyMarkup = cryptoKeyboard
			msg.Text = "Выберите криптовалюту"
			bot.Send(msg)

		default:
			msg.Text = "Неправильный формат сообщения"
			bot.Send((msg))
		}
	}
}

// https://t.me/CurrentEquivalentBot
