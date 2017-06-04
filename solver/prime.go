package solver

import (
	"math/big"
	"sort"

	"github.com/step17/ihazwordz/words"
)

var (
	primes    []int64
	bigPrimes []*big.Int
)

func init() {
	for p := int64(2); len(primes) < 26; p++ {
		if isPrime(p) {
			primes = append(primes, p)
			bigPrimes = append(bigPrimes, big.NewInt(p))
		}
	}
}

func isPrime(p int64) bool {
	switch {
	case p < 2:
		return false
	case p == 2:
		return true
	}

	for i := int64(2); i <= p/2; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func asPrime(letter byte) *big.Int {
	return bigPrimes[letter-'A']
}

// primeWord expresses a sequence of letters as a product of primes
// where each letter in the alphabet corresponds to a unique prime
// number.
type primeWord struct {
	i *big.Int
}

var bigZero = big.NewInt(0)

// Contains returns true iff needle is a subset of haystack.  If
// haystack does contain needle, haystack must be an integer multiple
// of needle.
func (haystack primeWord) Contains(needle primeWord) bool {
	var z, m big.Int
	z.DivMod(haystack.i, needle.i, &m)
	return z.Cmp(bigZero) > 0 && m.Cmp(bigZero) == 0
}

// pwPair stores an anagram word pair (the original word, and a count
// of the required letters expressed as a primeWord).
type pwPair struct {
	pw   primeWord
	word string
}

// pwPairs sortable by score.
type pwPairSlice []pwPair

func (p pwPairSlice) Len() int      { return len(p) }
func (p pwPairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pwPairSlice) Less(i, j int) bool {
	return Score(p[j].word) < Score(p[i].word)
}

// PrimeWordSolver goes through a sorted dictionary like ListSolver,
// but uses primeWords instead of sorted strings to represent anagram
// clusters.
type PrimeWordSolver struct {
	rs   RecursiveSolver
	dict []pwPair // Highest scoring words first.
	Solver
}

func (s *PrimeWordSolver) Init(dict []string) error {
	s.rs.Init(dict)
	for sorted, cs := range s.rs.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, pwPair{pw: s.PrimeWord(sorted), word: cs.first()})
	}
	sort.Sort(pwPairSlice(s.dict))
	return nil
}

// Expresses the given word as a primeWord.
func (s PrimeWordSolver) PrimeWord(word string) primeWord {
	word = words.Normalize(word)
	i := big.NewInt(1)
	for li := 0; li < len(word); li++ {
		i.Mul(i, asPrime(word[li]))
	}
	return primeWord{i}
}

func (s PrimeWordSolver) Solve(letters string) string {
	pw := s.PrimeWord(letters)
	for _, p := range s.dict {
		if pw.Contains(p.pw) {
			return p.word
		}
	}
	return ""
}
