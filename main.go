package main

import (
	"log"
	"trip-forwarder-bot/config"

	TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var conf = config.New()

func main() {
	bot, err := TelegramBotAPI.NewBotAPI(conf.TelegramBot.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = conf.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	newUpdate := TelegramBotAPI.NewUpdate(0)
	newUpdate.Timeout = 60

	updates := bot.GetUpdatesChan(newUpdate)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/start" {
				message := TelegramBotAPI.NewMessage(update.Message.Chat.ID, "Буду рад вам помочь :)")
				_, err := bot.Send(message)
				if err != nil {
					log.Print("Error")
				}
			} else {
				message := TelegramBotAPI.NewMessage(update.Message.Chat.ID, update.Message.Text)
				message.ReplyToMessageID = update.Message.MessageID

				_, err := bot.Send(message)
				if err != nil {
					log.Print("Error")
				}
			}
		}
	}
}
