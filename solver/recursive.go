package solver

import (
	"github.com/step17/ihazwordz/words"
)

// Basic recursive solver. Tries all 2^N possible subsets of letters.
type RecursiveSolver struct {
	c Common
	Solver
}

func (s *RecursiveSolver) Init(dict []string) error {
	return s.c.LexInit(dict)
}

func (s *RecursiveSolver) Solve(letters string) string {
	cs := s.resolve("", words.Sort(letters))
	return cs.first()
}

func (s *RecursiveSolver) resolve(picked, remain string) *choices {
	if remain == "" {
		return s.c.sorted[picked]
	}
	if len(picked)+len(remain) < s.c.minLen {
		return nil
	}
	next := remain[1:]
	kept := s.resolve(picked+remain[:1], next)
	skip := s.resolve(picked, next)
	if skip.score() > kept.score() {
		return skip
	}
	return kept
}
