package solver

import (
	"sort"
)

// ListScoredSolver goes through a score-sorted dictionary, checking if each word
// can be made out of the given characters.
type ListScoredSolver struct {
	c    Common
	dict []anaPair // Highest scoring words first.
	Solver
}

func (s *ListScoredSolver) Init(dict []string) error {
	s.c.ScoredInit(dict)
	for sorted, cs := range s.c.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, anaPair{sorted: sorted, word: cs.first()})
	}
	sort.Sort(anaPairSlice(s.dict))
	return nil
}

func (s ListScoredSolver) Solve(letters string) string {
	sorted := SortScore(letters)
	for _, p := range s.dict {
		if s.canSpell(p.sorted, sorted) {
			return p.word
		}
	}
	return ""
}

func (s ListScoredSolver) canSpell(sub, sup string) bool {
	// precondition: both sub and sup are sorted.
	bi, pi := 0, 0
	// Check each character of the substring (bi) to see if it's a
	// subsequence of the superstring (pi).
	for bi < len(sub) && pi < len(sup) {
		switch {
		case sub[bi] == sup[pi]:
			bi++
			pi++
		case LetterPoints[rune(sub[bi])] > LetterPoints[rune(sup[pi])]:
			// sub contains a character that isn't in sup
			return false
		default:
			// sup contains some letters that aren't in sub
			pi++
		}
	}
	// Did we successfully make it to the end of sub? If so, it's
	// entirely a subsequence of sup.
	return bi == len(sub)
}
