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

func (c *SolutionCommanderStruct) SendMessage(inputMsg *tgbotapi.Message, msgtext string){
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, msgtext)
	c.bot.Send(msg)
}

func GetArgument(inputMsg *tgbotapi.Message) (uint64, string){
	log.Println("Пытаемся получить 1ый аргумент после команда, должно быть число")
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		textMsg := fmt.Sprintf("За командой должно ID элемента с которым работаем, а получили \"%s\"", args)
		log.Println(textMsg)
		return 0, textMsg
	}
	log.Println("Число найдено")
	return idx, ""
}

