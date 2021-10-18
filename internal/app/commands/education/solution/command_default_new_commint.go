package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/servicedata"
)

func (c *SolutionCommander) newCommit(inputMessage *tgbotapi.Message) {
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMessage, TextMsg)
	}()
	if solution, Text := c.SolutionService.DecodeMessage(inputMessage); Text == ""{
		solution.Id = c.SolutionService.CreateNewID()
		c.SolutionService.Create(solution.Id, solution)
		delete(servicedata.EditedChat, inputMessage.Chat.ID)
		sol, _ := c.SolutionService.Describe(solution.Id)
		TextMsg = "Запись добавлена: \n " + sol.String()
	} else { TextMsg = Text}
}

