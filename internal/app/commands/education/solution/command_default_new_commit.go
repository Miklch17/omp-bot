package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education"
	"log"
)

func (c *SolutionCommander) newCommit(inputMessage *tgbotapi.Message) {
	log.Println("Пытаемся завершить создание записи")
	textMsg := ""
	defer func() {
		c.SendMessage(inputMessage, textMsg)
	}()
	if solution, Text := c.SolutionService.DecodeMessage(inputMessage); Text == ""{
		solution.ID = c.SolutionService.CreateNewID()
		if _, ok := c.SolutionService.Create(solution.ID, solution); ok!=nil {
			textMsg = "Ошибка добавления записи"
			log.Println("Ошибка при добавлении записи")
			return
		}
		education.DeleteEditedChatElement(inputMessage.Chat.ID)
		sol, ok := c.SolutionService.Describe(solution.ID)
		if ok != nil{
			textMsg = "Не смогли получить добавленную запись"
			log.Println("Не смогли получить добавленную запись")
			return
		}
		textMsg = "Запись добавлена: \n " + sol.String()
		log.Println("Запись удачно добавлена")
	} else {
		textMsg = Text
		log.Println("Что-то пошло не по плану: ", textMsg)
	}
}

