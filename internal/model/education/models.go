package education

import "fmt"

type Solution struct {
	Id uint64
	TaskID      uint64
	StudentID   uint64
	Description string
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

