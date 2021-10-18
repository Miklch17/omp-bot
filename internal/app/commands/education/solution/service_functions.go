package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func OneCommandName(Begin string) string{
	return fmt.Sprintf("/%s__education__solution", Begin)
}

func (c *SolutionCommander) SendMessage(inputMsg *tgbotapi.Message, msgtext string){
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, msgtext)
	c.bot.Send(msg)
}

func GetArgument(inputMsg *tgbotapi.Message) (uint64, string){
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		TextMsg := fmt.Sprintf("За командой должно ID элемента с которым работаем, а полчили \"%s\"", args)
		log.Println(TextMsg)
		return 0, TextMsg
	}
	return idx, ""
}

