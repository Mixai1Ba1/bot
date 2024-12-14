package commands

import (
<<<<<<< HEAD
	"fmt"
	"log"
	"strconv"

=======
>>>>>>> origin/master
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
<<<<<<< HEAD
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID, fmt.Sprintf("success: %v", arg),
	)
	c.bot.Send(msg)
}
=======

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "ЕИВ")
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
>>>>>>> origin/master
