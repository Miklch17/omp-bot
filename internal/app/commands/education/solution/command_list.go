package solution

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education"
	"log"
	"strings"
)

const (
	StartShowElementNumber = 0
	OffsetShowElementNumber = 2
)
func (c *SolutionCommander) List(inputMsg *tgbotapi.Message){
	log.Println("Пытаемся отразить 1 страницу списка и кнопку продолжения")
	cb := CallbackListData{ StartShowElementNumber, OffsetShowElementNumber}
	data, _ := json.Marshal(cb)

	callbackPath := path.CallbackPath{
		Domain:       education.Education,
		Subdomain:    education.Solution,
		CallbackName: "list",
		CallbackData: string(data),
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strings.Join(c.SolutionService.List(cb.Start, cb.Offset), "\n"))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String(),
			)))
	c.bot.Send(msg)
	log.Println("1 страница и кнпока удачно")
}

