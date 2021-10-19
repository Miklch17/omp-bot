package education

import "fmt"

type Solution struct {
	ID     uint64
	TaskID uint64 `json:"task_id"`
	StudentID   uint64 `json:"student_id"`
	Description string `json:"description"`
}

func (c Solution) String() string{
	return fmt.Sprintf("ID: %d TaskID: %d StudentID: %d Description: %s", c.ID, c.TaskID, c.StudentID, c.Description)
}

var data map[uint64]Solution

func GetData() *map[uint64]Solution {
	return &data
}

func init() {
	data = make(map[uint64]Solution, 100)
	data[1] = Solution{1, 100, 111, "Cool Description 1"}
	data[2] = Solution{2, 200, 222, "Cool Description 2"}
	data[3] = Solution{3, 300, 333, "Cool Description 3"}
	data[4] = Solution{4, 400, 444, "Cool Description 4"}
	data[5] = Solution{5, 500, 555, "Cool Description 5"}
	data[6] = Solution{6, 600, 666, "Cool Description 6"}
	data[7] = Solution{7, 700, 777, "Cool Description 7"}
}

const
	DescriptionNewOrEditCommand = `Данные можно заполнить в 3х видах: 
		1. 3 строки без каких либо меток:
		   TaskID
		   StudentID
		   Description.
		Все поля должны быть в одном сообщении каждое поле в отдельной строке.
		2. 3 строки с метками, информация по какому полю предоставляется:
		   task_id: 10
		   student_id: 20
		   description: Тест для описания
		Все поля должны быть в одном сообщении каждое поле в отдельной строке.
		3. Правильный json заполненный необходимыми данными:
		   {"task_id": 123, "student_id": 34235, "description": "test Description"} `