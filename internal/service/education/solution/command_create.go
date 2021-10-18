package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
)

func (s *DummySolutionService) Create(SolutionId uint64, Solution education.Solution) (uint64, error) {
	log.Println("Пытаемся добавить элемент в набор данных")
	if _, ok := education.Data[SolutionId]; ok {
		return 0, fmt.Errorf("item already exists")
	}
	education.Data[SolutionId] = Solution
	log.Println("Элемент успешно добавлен")
	return SolutionId, nil
}

