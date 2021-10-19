package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
)

func (s *DummySolutionService) Remove(SolutionId uint64) (bool, error) {
	log.Println("Пытаемся удалить запись")
	if _, ok := (*education.GetData())[SolutionId]; !ok {
		return false, fmt.Errorf("item not found")
	}
	delete((*education.GetData()), SolutionId)
	log.Println("Запись успешно удалена")
	return true, nil
}

