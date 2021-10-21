package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SolutionCommanderStruct) Default(inputMessage *tgbotapi.Message) {
	log.Println("Реакция на текст без комманд")
	log.Printf("Error enter: [%s] %s", inputMessage.From.UserName, inputMessage.Text)
	c.SendMessage(inputMessage, "You wrote: "+inputMessage.Text)
}
