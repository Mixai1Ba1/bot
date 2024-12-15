package commands

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Твой высер: "+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *Commander) HadlerUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			// "купить: "+update.CallbackQuery.Data,
			// fmt.Sprintf("command %s\n", args[0])+
			// fmt.Sprintf("offset  %s\n", args[1]),
			fmt.Sprintf("Parsed %+v\n", parsedData),
		)
		c.bot.Send(msg)
		return
	}
	// defer func() {
	// 	fmt.Println("first")
	// }()
	// первый вошел последний вышел LIFO
	// аргументы все передаются по значению и сохраняются в моменте когда вывзывается дефер
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
