package solution

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
)

func (s *DummySolutionService) Describe(SolutionId uint64) (*education.Solution, error) {
	log.Println("Пытаемся отдать элемент")
	if _, ok := education.Data[SolutionId]; !ok {
		return nil, fmt.Errorf("item not found")
	}
	r := education.Data[SolutionId]
	log.Println("Удачно отдали элемент")
	return &r, nil
}

