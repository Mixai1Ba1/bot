package main

import (
	"log"
	"os"

	"github.com/Mixai1Ba1/bot/internal/app/commands"
	"github.com/Mixai1Ba1/bot/internal/service/product"
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

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	// Обрабатываем обновления
	for update := range updates {
		commander.HadlerUpdate(update)
	}
}
