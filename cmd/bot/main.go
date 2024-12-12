package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// const token = "7882527472:AAHZVc451PxyAr-JDvPBCxxhujj-XN2Ui5Y"

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал обновлений
	updates := bot.GetUpdatesChan(u)

	// Обрабатываем обновления
	for update := range updates {
		// Проверяем, что сообщение существует
		if update.Message == nil {
			continue
		}

		
		if update.Message.Command() == "help" {
			helpCommand(bot, update.Message)
			continue
		}

		defaultBehavior(bot, update.Message)

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help  - help")
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Твой высер: "+inputMessage.Text)

	bot.Send(msg)
}
