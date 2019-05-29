package solver

// Basic recursive solver. Tries all 2^N possible subsets of letters.
type RecursiveClassSolver struct {
	c Common
	Solver
}

func (s *RecursiveClassSolver) Init(dict []string) error {
	return s.c.LexInit(dict)
}

func (s *RecursiveClassSolver) Solve(letters string) string {
	cs := s.resolve("", ToIsochars(letters))
	return cs.first()
}

func (s *RecursiveClassSolver) resolve(picked string, remain Isochars) *choices {
	if remain == nil {
		return s.c.sorted[picked]
	}
	if len(picked)+remain.Len() < s.c.minLen {
		return nil
	}
	next := remain.Next()
	kept := s.resolve(picked+remain.First(), next)
	if kept != nil {
		return kept
	}
	return s.resolve(picked, next)
}
