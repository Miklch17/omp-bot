package servicedata

import (
	"fmt"
	"log"
	"strconv"
)

const (
	NewOperationData = iota
	EditOperationData
)
type operationData struct{
	ProductID uint64
	OperationType int
}
var editedChat = make(map[int64]operationData, 100)

func DeleteEditedChatElement(Key int64) {
	delete(editedChat, Key)
}

func IsHaveEditedChatElement(Key int64) bool {
	_, ok := editedChat[Key]
	return ok
}

func GetEditedChatElement(Key int64) (*operationData, error) {
	log.Printf("Попытка получить элемент %d", Key)
	if element, ok := editedChat[Key]; !ok {
		log.Printf("Ошибка элемент не найден")
		return nil, fmt.Errorf("Такого элемента %d не существует", Key)
	} else {
		log.Println("Элемент получен ", Key, element)
		return &element, nil
	}
}

func AddOperationDataInEditedChat(Key int64, productID uint64, operationType int) {
	log.Println("инициализация структуры operationData")
	editedChat[Key] = operationData{
		ProductID: productID,
		OperationType: operationType,
	}
}

func (s *operationData) String() string {
	log.Println("Структура operationData ввиде строки")
	OperationName := ""
	switch s.OperationType {
	case NewOperationData:
		OperationName = "New"
	case EditOperationData:
		OperationName = "Edit"
	default:
		OperationName = "Unknown:" + strconv.Itoa(s.OperationType)
	}
	return fmt.Sprintf("Тип операции: %s, для ID = %d", OperationName, s.ProductID)
}