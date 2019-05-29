package solver

// EnumSolver enumerates all of the subsets of a given string
// iteratively instead of recursively.
type EnumSolver struct {
	c Common
	Solver
}

func (s *EnumSolver) Init(dict []string) error {
	return s.c.LexInit(dict)
}

func (s *EnumSolver) Solve(letters string) string {
	sorted := SortScore(letters)
	l := uint(len(letters))
	max := (1 << l) - 1
	var best *choices
	for i := max; i > 0; i-- {
		var sub string
		for j := l; j > 0; j-- {
			if (i & (1 << (j - 1))) != 0 {
				sub += string(sorted[l-j])
			}
		}
		cs := s.c.sorted[sub]
		if best.score() < cs.score() {
			best = cs
		}
	}
	return best.first()
}
