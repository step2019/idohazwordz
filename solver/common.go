package solver

import (
	"math"

	"github.com/step17/ihazwordz/words"
)

// Some common solver components used by most solvers.
type Common struct {
	sorter func(string) string // how to "sort" words
	sorted map[string]*choices // map of some "sorted" words to the corresponding original choices.
	minLen int                 // minimum length of words in sorted
}

func (c *Common) LexInit(dict []string) error {
	c.sorter = words.Sort
	return c.init(dict)
}

func (c *Common) ScoredInit(dict []string) error {
	c.sorter = SortScore
	return c.init(dict)
}

func (c *Common) init(dict []string) error {
	c.sorted = make(map[string]*choices)
	c.minLen = math.MaxInt32
	for _, word := range dict {
		if len(word) < c.minLen {
			c.minLen = len(word)
		}
		norm := words.Normalize(word)
		sNorm := c.sorter(norm)
		v := c.sorted[sNorm]
		if v == nil {
			v = &choices{points: Score(norm)}
			c.sorted[sNorm] = v
		}
		v.words = append(v.words, norm)
	}
	return nil
}

// Choices is a type for keeping track of lists of word choices with
// equivalent scores.
type choices struct {
	words  []string
	points int
}

// Score returns the score of this choice.
func (cs *choices) score() int {
	if cs == nil {
		return 0
	}
	return cs.points
}

// First arbitrarily picks the first out of a list of potential word
// choices, or an empty string if there are no choices.
func (cs *choices) first() string {
	if cs == nil {
		return ""
	}
	return cs.words[0]
}
