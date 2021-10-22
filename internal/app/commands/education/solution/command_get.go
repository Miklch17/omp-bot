package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SolutionCommanderStruct) Get(inputMsg *tgbotapi.Message){
	log.Println("Попытка показать запись")
	textMsg := ""
	defer func() {
		c.SendMessage(inputMsg, textMsg)
	}()
	idx, textMsg := GetArgument(inputMsg)
	if textMsg != "" { return}

	product, err := c.SolutionService.Describe(idx)
	if err != nil {
		textMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(textMsg)
		return
	}
	textMsg = product.String()
	log.Println("Попытка успешна")
}

