package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education"
	"log"
)

func (c *SolutionCommanderStruct) EditCommit(inputMessage *tgbotapi.Message) {
	log.Println("Пытаемся завершить редактирование")
	textMsg := ""
	defer func() {
		c.SendMessage(inputMessage, textMsg)
	}()
	if solution, Text := c.SolutionService.DecodeMessage(inputMessage); Text == ""{
		c.SolutionService.Update(solution.ID, solution)
		education.DeleteEditedChatElement(inputMessage.Chat.ID)
		sol, _ := c.SolutionService.Describe(solution.ID)
		textMsg = "Запись заменена: \n " + sol.String()
		log.Println("Запись удачно изменена")
	} else {
		textMsg = Text
		log.Println("Что-то пошло не по плану: ", textMsg)
	}
}

