package solver

import "github.com/step17/ihazwordz/words"

// MemoSolver is a basic recursive solver, but with memoization.
type MemoSolver struct {
	c    Common
	memo map[string]*choices
	Solver
}

func (s *MemoSolver) Init(dict []string) error {
	return s.c.LexInit(dict)
}

func (s *MemoSolver) Solve(letters string) string {
	s.memo = make(map[string]*choices)
	cs := s.resolve("", words.Sort(letters))
	return cs.first()
}

func (s *MemoSolver) key(picked, remain string) string {
	return picked + "," + remain
}

func (s *MemoSolver) resolve(picked, remain string) *choices {
	if remain == "" {
		return s.c.sorted[picked]
	}
	if len(picked)+len(remain) < s.c.minLen {
		return nil
	}

	// Have we already evaluated these arguments?
	key := s.key(picked, remain)
	memoRes, done := s.memo[key]
	if done {
		return memoRes
	}

	// Solve this the normal recursive way.
	next := remain[1:]
	kept := s.resolve(picked+remain[:1], next)
	res := kept
	skip := s.resolve(picked, next)
	if skip.score() > kept.score() {
		res = skip
	}
	// Record this result so we can skip it if it comes up again.
	s.memo[key] = res
	return res
}
