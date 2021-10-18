package servicedata

import (
	"fmt"
	"log"
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
	log.Println("инициализация структуры operationData")
	return &operationData{
		ProductID: productID,
		OperationType: operationType,
	}
}

func (s *operationData) String() string {
	log.Println("Структура operationData ввиде строки")
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