package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/servicedata"
	"log"
)

func (c *SolutionCommander) Default(inputMessage *tgbotapi.Message) {
	log.Println("Реакция на текст без комманд, если не идет редактирование или создание новой, то это ошибка")
	if idx, ok := servicedata.GetEditedChatElement(inputMessage.Chat.ID); ok == nil {
		switch idx.OperationType {
		case servicedata.NewOperationData:
			c.newCommit(inputMessage)
		case servicedata.EditOperationData:
			c.editCommit(inputMessage)
		default:
			c.SendMessage(inputMessage, "Неизвстная комманда " + idx.String())
		}
	} else {
		log.Println(ok)
		log.Printf("Error enter: [%s] %s", inputMessage.From.UserName, inputMessage.Text)
		c.SendMessage(inputMessage, "You wrote: "+inputMessage.Text)
	}
}
