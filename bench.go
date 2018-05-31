package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/step18/ihazwordz/solver"
	"github.com/step18/ihazwordz/words"
)

var (
	dictionaryFile = flag.String("dictionary", "/usr/share/dict/words", "Dictionary file to use")
	letterCount    = flag.Int("letters", 16, "How many letters are provided on board?")
	repitions      = flag.Int("reptitions", 3, "How many times to run each workload.")
	workloadCount  = flag.Int("workloads", 123, "Run this many different workloads.")
	workloadSize   = flag.Int("size", 123, "How many games to run per workload.")
	seed           = flag.Int64("seed", 0, "Initial random seed.")
)

// A workload is collection of game boards that might appear
// (i.e. something like ["abcdefg", "jqszpf"] and so on).
type workload []string

func sampleString() string {
	var pool string
	vals := map[int]string{
		1: "abdeginorstu",
		2: "lcfhmpvwy",
		3: "jkqxz",
	}
	for p, s := range vals {
		for i := 0; i < p; i++ {
			pool += s
		}
	}
	return pool
}

func fakeWorkload(rng *rand.Rand, size int) workload {
	pool := sampleString()
	work := make(workload, size)
	for i := 0; i < size; i++ {
		var board string
		for l := 0; l < *letterCount; l++ {
			board += string(pool[rng.Intn(len(pool))])
		}
		work[i] = board
	}
	return work
}

func runWorkload(s solver.Solver, wl workload) time.Duration {
	start := time.Now()
	for _, w := range wl {
		s.Solve(w)
	}
	return time.Since(start)
}

func name(s solver.Solver) string {
	return strings.TrimPrefix(fmt.Sprintf("%T", s), "*solver.")
}

func main() {
	flag.Parse()
	dict := words.Load(*dictionaryFile, *letterCount)
	// warmup all the solvers.
	solvers := solver.AllSolvers
	for _, s := range solvers {
		s.Init(dict)
	}
	rng := rand.New(rand.NewSource(*seed))
	fmt.Println(strings.Join([]string{"solver", "workload", "rep", "size", "nanos", "nanos_per_solution"}, ","))
	for wi := 0; wi < *workloadCount; wi++ {
		wl := fakeWorkload(rng, *workloadSize)
		for _, s := range solvers {
			for rep := 0; rep < *repitions; rep++ {
				dur := runWorkload(s, wl)
				fmt.Printf("%v,%v,%v,%v,%v,%v\n", name(s), wi, rep, *workloadSize, dur.Nanoseconds(), float64(dur.Nanoseconds())/float64(*workloadSize))
			}
		}
	}
}
