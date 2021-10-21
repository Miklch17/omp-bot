package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	education2 "github.com/ozonmp/omp-bot/internal/service/education"
	"log"
)

func (c *SolutionCommander) Edit(inputMsg *tgbotapi.Message){
	log.Println("Попытка начать редактирование записи")
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
	education2.AddOperationDataInEditedChat(inputMsg.Chat.ID, idx, education2.EditOperationData)
	textMsg = product.String() + education.DescriptionNewOrEditCommand
	log.Println("Редактирование начали")
}

