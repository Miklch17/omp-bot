package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"strconv"
	"strings"
)

func (c *SolutionCommander) editCommit(inputMessage *tgbotapi.Message) {
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMessage, TextMsg)
	}()
	data := strings.Split(inputMessage.Text, "\n")
	if len(data) != 3 {
		TextMsg = "В сообщение должно быть 3 строки, повторите ввод, пожалуйста"
		return
	}
	taskID, err := strconv.ParseUint(data[0], 0, 64)
	if err != nil {
		TextMsg = "Первая строка не содержит число, повторите ввод, пожалуйста"
		return
	}
	studentID, err := strconv.ParseUint(data[1], 0, 64)
	if err != nil {
		TextMsg = "Вторая строка не содержит число, повторите ввод, пожалуйста"
		return
	}
	idx, _ := servicedata.EditedChat[inputMessage.Chat.ID]
	solution := education.Solution{	}
	solution.Id = idx.ProductID
	solution.TaskID = taskID
	solution.StudentID = studentID
	solution.Description = data[2]
	c.SolutionService.Update(idx.ProductID, solution)
	delete(servicedata.EditedChat, inputMessage.Chat.ID)
	sol, _ := c.SolutionService.Describe(idx.ProductID)
	TextMsg = "Запись заменена: \n " + sol.String()
}

