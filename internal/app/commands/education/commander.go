package education

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/solution"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education"
	"log"
)

type EducationCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type EducationCommanderStruct struct {
	bot                *tgbotapi.BotAPI
	SolutionCommander solution.SolutionCommander
}

func NewEducationCommander(
	bot *tgbotapi.BotAPI,
) *EducationCommanderStruct {
	return &EducationCommanderStruct{
		bot: bot,
		// SolutionCommanderStruct
		SolutionCommander: solution.NewSolutionCommander(bot),
	}
}

func (c *EducationCommanderStruct) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case education.Solution:
		c.SolutionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("EducationCommanderStruct.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *EducationCommanderStruct) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case education.Solution:
		c.SolutionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("EducationCommanderStruct.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
