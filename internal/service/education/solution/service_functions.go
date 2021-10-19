package solution

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/service/education/servicedata"
	"log"
	"strconv"
	"strings"
)

func (s *DummySolutionService) CreateNewID() uint64 {
	log.Println("Получить следующий ID для записи, скорее всего потокоНЕбезопастна, но потоки еще не проходили")
	max := uint64(0)
	for i := range *education.GetData() {
		if max < i {max = i}
	}
	return max + 1
}

func (s *DummySolutionService) Len() uint64 {
	log.Println("Получаем длинну данных")
	return uint64(len(*education.GetData()))
}

const errorMessageFormat = "Внимательно ознакомьтесь с подсказкой, переданные данные не соотвествуют ни одному из " +
		"3х возможных форматов, повторите ввод заново."
const (
	taskIDStringNumber = iota
	studentIDStringNumber
	descriptionStringNumber
)
func SimpleInputData(inputMessage *tgbotapi.Message, data []string)  (education.Solution, string){
	log.Println("Похоже пришло просто 3 строки")
	textMsg := ""
	if len(data) != 3 {
		textMsg = errorMessageFormat
		return education.Solution{}, textMsg
	}
	taskID, err := strconv.ParseUint(data[taskIDStringNumber], 0, 64)
	if err != nil {
		textMsg = errorMessageFormat
		return education.Solution{}, textMsg
	}
	studentID, err := strconv.ParseUint(data[studentIDStringNumber], 0, 64)
	if err != nil {
		textMsg = errorMessageFormat
		return education.Solution{}, textMsg
	}
	idx, _ := servicedata.GetEditedChatElement(inputMessage.Chat.ID)
	log.Println("Похоже пришло просто 3 строки - успешно преобразовали")
	return education.Solution{
		ID:          idx.ProductID,
		TaskID:      taskID,
		StudentID:   studentID,
		Description: data[descriptionStringNumber],
	}, textMsg

}

func LabelInputData(inputMessage *tgbotapi.Message, data []string)  (education.Solution, string){
	log.Println("Похоже пришло 3 строки с метками")
	textMsg := ""
	task := ""
	student := ""
	description := ""
	for i := 0; i < 3; i++ {
		if strings.Contains(data[i], "task_id") {
			idx := strings.Index(data[i], "task_id:") + len("task_id:")
			task = strings.TrimSpace(data[i][idx:])
		}
		if strings.Contains(data[i], "student_id:") {
			idx := strings.Index(data[i], "student_id:") + len("student_id:")
			student = strings.TrimSpace(data[i][idx:])
		}
		if strings.Contains(data[i], "description:") {
			idx := strings.Index(data[i], "description:") + len("description:")
			description = strings.TrimSpace(data[i][idx:])
		}
	}

	taskID, err := strconv.ParseUint(task, 0, 64)
	if err != nil {
		textMsg = errorMessageFormat
		return education.Solution{}, textMsg
	}
	studentID, err := strconv.ParseUint(student, 0, 64)
	if err != nil {
		textMsg = errorMessageFormat
		return education.Solution{}, textMsg
	}
	idx, _ := servicedata.GetEditedChatElement(inputMessage.Chat.ID)
	log.Println("Похоже пришло 3 строки с метками - успешно преобразовали")
	return education.Solution{
		ID:          idx.ProductID,
		TaskID:      taskID,
		StudentID:   studentID,
		Description: description,
	}, textMsg
}

func (s *DummySolutionService) DecodeMessage(inputMessage *tgbotapi.Message) (education.Solution, string) {
	log.Println("Пробует разобрать что же нам пришло")
	textMsg := ""
	data := strings.Split(inputMessage.Text, "\n")
	if strings.Contains(inputMessage.Text, "task_id:") || strings.Contains(inputMessage.Text, "student_id:") ||
		strings.Contains(inputMessage.Text, "description:") && len(data) == 3{
		//Тут вариант данных с метками полей
		return LabelInputData(inputMessage, data)
	}
	parsedData := education.Solution{}
	err := json.Unmarshal([]byte(inputMessage.Text), &parsedData)
	if err == nil {
		//Тут вариант с json
		idx, _ := servicedata.GetEditedChatElement(inputMessage.Chat.ID)
		parsedData.ID = idx.ProductID
		log.Println("пришел json")
		return parsedData, textMsg
	}
	//Последний возможный вариант с 3 обычными строками
	return SimpleInputData(inputMessage, data)
}
