package solver

// Basic recursive solver. Tries all 2^N possible subsets of letters.
type ScoredMemoSolver struct {
	c    Common
	todo map[int][]func() *choices
	Solver
}

func (s *ScoredMemoSolver) Init(dict []string) error {
	return s.c.ScoredInit(dict)
}

func reverse(str string) string {
	b := []byte(str)
	for i, j := 0, len(b)-1; i < j; {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func (s *ScoredMemoSolver) Solve(letters string) string {
	maxScore := SumScore(letters)
	s.todo = map[int][]func() *choices{
		0: []func() *choices{func() *choices { return s.resolve("", SortScore(letters), 0) }},
	}
	for skipping := 0; skipping < maxScore; skipping++ {
		for _, f := range s.todo[skipping] {
			if cs := f(); cs != nil {
				return cs.first()
			}
		}
	}
	return ""
}

func (s *ScoredMemoSolver) resolve(picked, remain string, skipped int) *choices {
	if remain == "" {
		return s.c.sorted[picked]
	}
	if len(picked)+len(remain) < s.c.minLen {
		return nil
	}
	next := remain[1:]
	kept := s.resolve(picked+remain[:1], next, skipped)
	if kept != nil {
		return kept
	}
	p := LetterPoints[rune(remain[0])]
	s.todo[skipped+p] = append(s.todo[skipped+p], func() *choices { return s.resolve(picked, next, skipped+p) })
	return nil
}
