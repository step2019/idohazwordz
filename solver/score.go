package solver

import (
	"sort"

	"github.com/step2019/idohazwordz/words"
)

type PointMap map[rune]int

var LetterPoints PointMap

func init() {
	LetterPoints = make(PointMap)
	// Assign everything 1 for now:
	for c := 'A'; c <= 'Z'; c++ {
		LetterPoints[c] = 1
	}
	// Assign some letters more.
	LetterPoints.assignPoints("LCFHMPVWY", 2)
	LetterPoints.assignPoints("JKQXZ", 3)
}

func (m PointMap) assignPoints(letters string, points int) {
	for _, l := range letters {
		m[l] = points
	}
}

func SumScore(word string) int {
	if word == "" {
		return 0
	}
	score := 1
	for _, l := range words.Normalize(word) {
		score += LetterPoints[l]
	}
	return score
}

func Score(word string) int {
	score := SumScore(word)
	return score * score
}

type scoredRuneSlice []rune

func (s scoredRuneSlice) Len() int      { return len(s) }
func (s scoredRuneSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s scoredRuneSlice) Less(i, j int) bool {
	pi, pj := LetterPoints[s[i]], LetterPoints[s[j]]
	switch {
	case pi != pj:
		return pi > pj
	default:
		return s[i] < s[j]
	}
}

// Sort a string's by descending letter score (i.e. high scoring
// letters first).
func SortScore(letters string) string {
	rs := scoredRuneSlice(letters)
	sort.Sort(rs)
	return string(rs)
}
