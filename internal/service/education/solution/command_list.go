package solution

import (
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
	"sort"
)

func (s *DummySolutionService) List(cursor uint64, limit uint64) []string {
	log.Println("Формируем первую страницу списка")
	if uint64(len(*education.GetSolution())) < cursor {
		return []string{}
	}
	if uint64(len(*education.GetSolution())) < cursor + limit {
		limit = uint64(len(*education.GetSolution())) - cursor
	}
	//Наверное есть более правильный метод, но я не смог придумать как из мапы вернуть элементы, а если через массив
	//делать то сильно усложняются другие методы
	rs := make([]string, 0, len(*education.GetSolution()))
	for _, v := range *education.GetSolution() {
		rs = append(rs, v.String())
	}
	sort.Strings(rs)
	res := make([]string, 0, limit)
	for i:= cursor; i < cursor + limit; i++ {
		res = append(res, rs[i])
	}
	log.Println("Первая страница сформирована")
	return res
}

