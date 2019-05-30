package solver

import (
	"log"
	"regexp"
	"sort"

	"github.com/step17/ihazwordz/words"
)

// rePair stores an anagram word pair. It contains the original word,
// and a regular expression that matches sorted strings that can spell
// this word.
type rePair struct {
	re   *regexp.Regexp
	word string
}

// rePairs sortable by score.
type rePairSlice []rePair

func (p rePairSlice) Len() int      { return len(p) }
func (p rePairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p rePairSlice) Less(i, j int) bool {
	return Score(p[j].word) < Score(p[i].word)
}

// RegexSolver goes through a sorted dictionary, checking if each word
// can be made out of the given characters using regexs.
type RegexSolver struct {
	c    Common
	dict []rePair // Highest scoring words first.
	Solver
}

func (s *RegexSolver) Init(dict []string) error {
	s.c.LexInit(dict)
	for sorted, cs := range s.c.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, rePair{re: asRE(sorted), word: cs.first()})
	}
	sort.Sort(rePairSlice(s.dict))
	return nil
}

func (s RegexSolver) Solve(letters string) string {
	sorted := words.Sort(letters)
	for _, p := range s.dict {
		log.Printf("check %v vs %v", sorted, p.re)
		if p.re.MatchString(sorted) {
			return p.word
		}
	}
	return ""
}

func chrTween(prev, c rune) string {
	if prev+1 == c {
		return ""
	}
	s := "["
	for i := prev + 1; i < c; i++ {
		s += string(i)
	}
	s += "]*"
	return s
}

func asRE(sorted string) *regexp.Regexp {
	if sorted == "" {
		return regexp.MustCompile("")
	}
	prev := 'A' - 1
	re := ""
	for _, c := range sorted {
		if prev != c {
			if prev >= 'A' {
				re += "+"
			}
			re += chrTween(prev, c)
			prev = c
		}
		re += string(c)
	}
	re += ".*"
	log.Printf("%v -> %v", sorted, re)
	return regexp.MustCompile(re)
}
