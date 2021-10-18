package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"log"
)

func (c *SolutionCommander) editCommit(inputMessage *tgbotapi.Message) {
	log.Println("Пытаемся завершить редактирование")
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMessage, TextMsg)
	}()
	if solution, Text := c.SolutionService.DecodeMessage(inputMessage); Text == ""{
		c.SolutionService.Update(solution.Id, solution)
		delete(servicedata.EditedChat, inputMessage.Chat.ID)
		sol, _ := c.SolutionService.Describe(solution.Id)
		TextMsg = "Запись заменена: \n " + sol.String()
		log.Println("Запись удачно изменена")
	} else {
		TextMsg = Text
		log.Println("Что-то пошло не по плану: ", TextMsg)
	}
}

