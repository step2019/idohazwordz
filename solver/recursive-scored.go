package solver

// Basic recursive solver. Tries all 2^N possible subsets of letters.
type RecursiveScoredSolver struct {
	c Common
	Solver
}

func (s *RecursiveScoredSolver) Init(dict []string) error {
	return s.c.ScoredInit(dict)
}

func (s *RecursiveScoredSolver) Solve(letters string) string {
	maxScore := SumScore(letters)
	for skipping := 0; skipping < maxScore; skipping++ {
		if cs := s.resolve("", SortScore(letters), skipping); cs != nil {
			return cs.first()
		}
	}
	return ""
}

func (s *RecursiveScoredSolver) resolve(picked, remain string, skip int) *choices {
	if remain == "" {
		return s.c.sorted[picked]
	}
	if len(picked)+len(remain) < s.c.minLen {
		return nil
	}
	next := remain[1:]
	kept := s.resolve(picked+remain[:1], next, skip)
	if kept != nil {
		return kept
	}
	p := LetterPoints[rune(remain[0])]
	if p > skip {
		return nil
	}
	return s.resolve(picked, next, skip-p)
}
