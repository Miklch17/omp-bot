package solution

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service_consts"
	"log"
	"strings"
)

func (c *SolutionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Println("Пытаемся отразить следующую страницу")
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	var msg = tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		strings.Join(c.SolutionService.List(parsedData.Start, parsedData.Offset), "\n"))

	if c.SolutionService.Len() > parsedData.Start + parsedData.Offset {
		parsedData.Start += parsedData.Offset
		data, _ := json.Marshal(parsedData)
		callbackData := path.CallbackPath{
			Domain:       service_consts.Education,
			Subdomain:    service_consts.Solution,
			CallbackName: "list",
			CallbackData: string(data),
		}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next", callbackData.String(),
				)))
	}
	c.bot.Send(msg)
	log.Println("Следующую страницу отобразили")
}

