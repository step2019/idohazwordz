package solver

import (
	"fmt"
	"reflect"
	"strings"
)

// Solver is an abstraction of one of the many possible ways to solve
// icanhazwordz. A caller first loads a dictionary with Init, then can
// lookup the highest scoring words using Solve.
type Solver interface {
	// Initializes the Solver with the specified dictionary of words.
	Init(dict []string) error
	// Returns one of the highest scoring words that can be made using
	// the given letters.
	Solve(letters string) string
}

var AllSolvers = []Solver{
	&RecursiveSolver{},
	&RecursiveScoredSolver{},
	&MemoSolver{},
	&EnumSolver{},
	&ListSolver{},
	&ListScoredSolver{},
	&RegexSolver{},
	&RegexScoredSolver{},
	&CountListSolver{},
	&BitfieldSolver{},
	&PrimeWordSolver{},
}

// Returns a reasonable name for the given Solver.
func Name(s Solver) string {
	return strings.TrimSuffix(reflect.TypeOf(s).Elem().Name(), "Solver")
}

// Returns the names of a slice of Solvers, as a convenience.
func Names(sl []Solver) []string {
	names := []string{}
	for _, s := range AllSolvers {
		names = append(names, Name(s))
	}
	return names
}

// Returns a new solver by the specified name.
func New(name string) (Solver, error) {
	for _, s := range AllSolvers {
		if strings.ToLower(Name(s)) != strings.ToLower(name) {
			continue
		}
		// Actually make a new one rather than just returning the existing
		// one, just for clarity's sake.
		return reflect.New(reflect.TypeOf(s).Elem()).Interface().(Solver), nil
	}
	return nil, fmt.Errorf("unknown solver: %v", name)
}
