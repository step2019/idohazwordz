package solver

import (
	"sort"
)

// RegexScoredSolver goes through a sorted dictionary, checking if each word
// can be made out of the given characters using regexs.
type RegexScoredSolver struct {
	c    Common
	dict []rePair // Highest scoring words first.
	Solver
}

func (s *RegexScoredSolver) Init(dict []string) error {
	s.c.ScoredInit(dict)
	for sorted, cs := range s.c.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, rePair{re: asRE(sorted), word: cs.first()})
	}
	sort.Sort(rePairSlice(s.dict))
	return nil
}

func (s RegexScoredSolver) Solve(letters string) string {
	sorted := SortScore(letters)
	for _, p := range s.dict {
		if p.re.MatchString(sorted) {
			return p.word
		}
	}
	return ""
}
