package servicedata

import (
	"fmt"
	"strconv"
)

const (
	NewoperationData  = iota
	EditoperationData
)
type operationData struct{
	ProductID uint64
	OperationType int
}
var EditedChat = make(map[int64]operationData, 100)

func GetOperationData(productID uint64, operationType int) *operationData{
	return &operationData{
		ProductID: productID,
		OperationType: operationType,
	}
}

func (s *operationData) String() string {
	OperationName := ""
	switch s.OperationType {
	case NewoperationData:
		OperationName = "New"
	case EditoperationData:
		OperationName = "Edit"
	default:
		OperationName = "Unknown:" + strconv.Itoa(s.OperationType)
	}
	return fmt.Sprintf("Тип операции: %s, для ID = %d", OperationName, s.ProductID)
}

const
	DescriptionNewOrEditCommand = "Данные можно заполнить в 3х видах: \n" +
		"1. 3 строки без каких либо меток: \n"+
		"   TaskID \n"+
		"   StudentID \n"+
		"   Description. \n"+
		"Все поля должны быть в одном сообщении каждое поле в отдельной строке.\n"+
		"2. 3 строки с метками, информация по какому полю предоставляется:\n" +
		"   TaskID: 10 \n"+
		"   StudentID: 20 \n"+
		"   Description: Тест для описания\n"+
		"Все поля должны быть в одном сообщении каждое поле в отдельной строке.\n"+
		"3. Правильный json заполненный необходимыми данными:\n" +
		"   {\"task_id\": 123, \"StudentID\": 34235, \"Description\": \"test Description\"} "