package words

import (
	"sort"
)

type runeSlice []rune

func (s runeSlice) Len() int      { return len(s) }
func (s runeSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s runeSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func Sort(letters string) string {
	rs := runeSlice(letters)
	sort.Sort(rs)
	return string(rs)
}

func SortRunes(rs []rune) {
	sort.Sort(rs)
}
