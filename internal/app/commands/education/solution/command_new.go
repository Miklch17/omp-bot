package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"log"
)

func (c *SolutionCommander) New(inputMsg *tgbotapi.Message){
	log.Println("Пытаемся запусть создание новой записи")
	servicedata.EditedChat[inputMsg.Chat.ID] = *(servicedata.GetOperationData(0, servicedata.NewoperationData))
	TextMsg := "Добавление новой записи.\n" + education.DescriptionNewOrEditCommand
	c.SendMessage(inputMsg, TextMsg)
	log.Println("Запуск прошел успешно")
}

