package solver

import (
	"strings"

	"github.com/step2019/idohazwordz/words"
)

// A set of characters all having the same score. Within an Isochar
// the runes are sorted lexicographically.
type Isochars []string

func ToIsochars(s string) (res Isochars) {
	x := map[int]string{}
	for _, c := range s {
		x[LetterPoints[c]] = x[LetterPoints[c]] + string(c)
	}
	for _, p := range []int{3, 2, 1} {
		s := x[p]
		if s == "" {
			continue
		}
		res = append(res, words.Sort(s))
	}
	return res
}

func (i Isochars) String() string {
	return strings.Join([]string(i), ",")
}

func (i Isochars) Len() int {
	tot := 0
	for _, x := range i {
		tot += len(x)
	}
	return tot
}

func (i Isochars) First() string {
	if i == nil {
		return ""
	}
	return string(i[0][0])
}

func (i Isochars) Next() Isochars {
	if i == nil {
		return nil
	}
	f := i[0]
	if len(f) == 1 {
		return i[1:]
	}
	i[0] = i[0][1:]
	return i
}
