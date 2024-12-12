package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7882527472:AAHZVc451PxyAr-JDvPBCxxhujj-XN2Ui5Y")
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

		// Логируем сообщение
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Твой высер: "+update.Message.Text)
		// msg.ReplyToMessageID = update.Message.MessageID

		// Отправляем сообщение
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}
