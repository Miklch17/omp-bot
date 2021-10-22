package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/solution"
	"log"
)

type SolutionCommanderStruct struct {
	bot              *tgbotapi.BotAPI
	SolutionService *solution.DummySolutionService
}

type SolutionCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	NewCommit(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
	EditCommit(inputMsg *tgbotapi.Message)
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CallbackListData struct {
	Start uint64  `json:"start"`
	Offset uint64 `json:"offset"`
}

func NewSolutionCommander(bot *tgbotapi.BotAPI, ) *SolutionCommanderStruct {
	service := solution.NewDummySolutionService()
	return &SolutionCommanderStruct{
		bot:             bot,
		SolutionService: service,
	}
}

func (c *SolutionCommanderStruct) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *SolutionCommanderStruct) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "edit_commit":
		c.EditCommit(msg)
	case "new_commit":
		c.NewCommit(msg)
	default:
		c.Default(msg)
	}
}
