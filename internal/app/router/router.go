package router

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/demo"
	"github.com/ozonmp/omp-bot/internal/app/commands/education"
	education2 "github.com/ozonmp/omp-bot/internal/service/education"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type EducationCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot *tgbotapi.BotAPI

	// demoCommander
	 demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	// work
	// service
	// exchange
	// estate
	// rating
	// security
	// cinema
	// logistic
	// product
	// education
	EducationCommander EducationCommander
}

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
	return &Router{
		// bot
		bot: bot,
		// demoCommander
		 demoCommander: demo.NewDemoCommander(bot),
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		// work
		// service
		// exchange
		// estate
		// rating
		// security
		// cinema
		// logistic
		// product
		// education
		EducationCommander: education.NewEducationCommander(bot),
	}
}

func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		c.demoCommander.HandleCallback(callback, callbackPath)
	case "user":
	case "access":
	case "buy":
	case "delivery":
	case "recommendation":
	case "travel":
	case "loyalty":
	case "bank":
	case "subscription":
	case "license":
	case "insurance":
	case "payment":
	case "storage":
	case "streaming":
	case "business":
	case "work":
	case "service":
	case "exchange":
	case "estate":
	case "rating":
	case "security":
	case "cinema":
	case "logistic":
	case "product":
	case education2.Education:
		c.EducationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	var commandPath path.CommandPath
	//Вот тут тоже не уверен что правильно так реализовывать
	if item, err := education2.GetEditedChatElement(msg.Chat.ID); err == nil {
		commandPath = path.CommandPath{
			CommandName: item.OperationType,
			Domain:      education2.Education,
			Subdomain:   education2.Solution,
		}
	} else {
		if !msg.IsCommand() {
			c.showCommandFormat(msg)

			return
		}
		commandPath, err = path.ParseCommand(msg.Command())
		if err != nil {
			log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
			return
		}
	}

	switch commandPath.Domain {
	case "demo":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "user":
	case "access":
	case "buy":
	case "delivery":
	case "recommendation":
	case "travel":
	case "loyalty":
	case "bank":
	case "subscription":
	case "license":
	case "insurance":
	case "payment":
	case "storage":
	case "streaming":
	case "business":
	case "work":
	case "service":
	case "exchange":
	case "estate":
	case "rating":
	case "security":
	case "cinema":
	case "logistic":
	case "product":
	case education2.Education:
		c.EducationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	c.bot.Send(outputMsg)
}
