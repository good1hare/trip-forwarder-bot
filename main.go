package main

import (
	TelegramBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"trip-forwarder-bot/config"
	"trip-forwarder-bot/models"
)

var conf = config.New()

func main() {
	db := models.ConnectDB()

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
				user := models.User{Name: update.Message.From.UserName}
				db.Create(&user)
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
