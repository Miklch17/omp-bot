package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
)

func (s *DummySolutionService) Update(SolutionId uint64, Solution education.Solution) error {
	log.Println("Пытаемся обновить запись")
	if _, ok := (*education.GetSolution())[SolutionId]; !ok {
		return fmt.Errorf("item not found")
	}
	(*education.GetSolution())[SolutionId] = Solution
	log.Println("Успешно обновили")
	return nil
}

