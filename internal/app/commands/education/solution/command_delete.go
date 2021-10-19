package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SolutionCommander) Delete(inputMsg *tgbotapi.Message){
	log.Println("Попытка удалить документ")
	textMsg := ""
	defer func() {
		c.SendMessage(inputMsg, textMsg)
	}()
	idx, textMsg := GetArgument(inputMsg)
	if textMsg != "" { return}

	_, err := c.SolutionService.Remove(idx)
	if err != nil {
		textMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(textMsg)
		return
	}
	textMsg = "Запись удалена"
	log.Println("Удаление прошло успешно")
}

