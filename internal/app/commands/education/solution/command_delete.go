package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SolutionCommander) Delete(inputMsg *tgbotapi.Message){
	log.Println("Попытка удалить документ")
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	_, err := c.SolutionService.Remove(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	TextMsg = "Запись удалена"
	log.Println("Удаление прошло успешно")
}

