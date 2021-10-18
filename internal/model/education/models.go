package education

import "fmt"

type Solution struct {
	Id uint64
	TaskID      uint64 `json:"task_id"`
	StudentID   uint64 `json:"student_id"`
	Description string `json:"description"`
}

func (c Solution) String() string{
	return fmt.Sprintf("ID: %d TaskID: %d StudentID: %d Description: %s", c.Id, c.TaskID, c.StudentID, c.Description)
}

var Data map[uint64]Solution

func init() {
	Data = make(map[uint64]Solution, 100)
	Data[1] = Solution{1, 100, 111, "Cool Description 1"}
	Data[2] = Solution{2, 200, 222, "Cool Description 2"}
	Data[3] = Solution{3, 300, 333, "Cool Description 3"}
	Data[4] = Solution{4, 400, 444, "Cool Description 4"}
	Data[5] = Solution{5, 500, 555, "Cool Description 5"}
	Data[6] = Solution{6, 600, 666, "Cool Description 6"}
	Data[7] = Solution{7, 700, 777, "Cool Description 7"}
}

const
	DescriptionNewOrEditCommand = "Данные можно заполнить в 3х видах: \n" +
		"1. 3 строки без каких либо меток: \n"+
		"   TaskID \n"+
		"   StudentID \n"+
		"   Description. \n"+
		"Все поля должны быть в одном сообщении каждое поле в отдельной строке.\n"+
		"2. 3 строки с метками, информация по какому полю предоставляется:\n" +
		"   task_id: 10 \n"+
		"   student_id: 20 \n"+
		"   description: Тест для описания\n"+
		"Все поля должны быть в одном сообщении каждое поле в отдельной строке.\n"+
		"3. Правильный json заполненный необходимыми данными:\n" +
		"   {\"task_id\": 123, \"student_id\": 34235, \"description\": \"test Description\"} "