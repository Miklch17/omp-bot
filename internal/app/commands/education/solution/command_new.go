package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	education2 "github.com/ozonmp/omp-bot/internal/service/education"
	"log"
)

func (c *SolutionCommander) New(inputMsg *tgbotapi.Message){
	log.Println("Пытаемся запусть создание новой записи")
	education2.AddOperationDataInEditedChat(inputMsg.Chat.ID, 0, education2.NewOperationData)
	TextMsg := "Добавление новой записи.\n" + education.DescriptionNewOrEditCommand
	c.SendMessage(inputMsg, TextMsg)
	log.Println("Запуск прошел успешно")
}

