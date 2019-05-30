// This program benchmarks all of the solvers by running a series of
// randomly generated workloads using all of the solvers.  It writes a
// CSV table of execution times to stdout.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/step2019/idohazwordz/solver"
	"github.com/step2019/idohazwordz/words"
)

var (
	dictionaryFile   = flag.String("dictionary", "/usr/share/dict/words", "Dictionary file to use")
	dictProbableFrac = flag.Float64("sample", 1, "probabilistically sample this fraction of the dictionary")
	dictExactFrac    = flag.Float64("exact", 1, "sample exactly this fraction of the dictionary (after --sample)")
	letterCount      = flag.Int("letters", 16, "How many letters are provided on board?")
	repitions        = flag.Int("reps", 3, "How many times to run each workload.")
	workloadCount    = flag.Int("workloads", 123, "Run this many different workloads.")
	workloadSize     = flag.Int("size", 123, "How many games to run per workload.")
	seed             = flag.Int64("seed", 0, "Initial random seed. Uses time.Now() if 0.")
	ignoreSolverStr  = flag.String("ignore", "RecursiveScored,ListScored", "CSV of solvers to skip")
	onlySolverStr    = flag.String("only", "", "if specified, benchmarks only these solvers.")
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

func runif(r *rand.Rand, p float64) func(string) bool {
	return func(string) bool {
		if p >= 1 {
			return true
		}
		return r.Float64() < p
	}
}

func resample(r *rand.Rand, words []string, f float64) []string {
	if f >= 1 {
		return words
	}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	n := int(float64(len(words)) * f)
	log.Printf("reducing to %v words", n)
	return words[:n]
}

func parseSolvers(s string) map[string]bool {
	m := map[string]bool{}
	for _, f := range strings.Split(s, ",") {
		m[f] = true
	}
	return m
}

func main() {
	flag.Parse()
	ignoreSolvers := parseSolvers(*ignoreSolverStr)
	onlySolvers := parseSolvers(*onlySolverStr)
	if *seed == 0 {
		*seed = time.Now().UnixNano()
	}
	rng := rand.New(rand.NewSource(*seed))
	dict := words.SamplingLoad(*dictionaryFile, *letterCount, runif(rng, *dictProbableFrac))
	dict = resample(rng, dict, *dictExactFrac)
	// warmup all the solvers.
	solvers := []solver.Solver{}
	for _, s := range solver.AllSolvers {
		name := solver.Name(s)
		if len(onlySolvers) == 0 {
			if ignoreSolvers[name] {
				continue
			}
		} else if !onlySolvers[name] {
			continue
		}
		solvers = append(solvers, s)
		s.Init(dict)
	}
	fmt.Println(strings.Join([]string{
		"solver", "workload", "rep", "size",
		"dict", "letters",
		"nanos", "log10_nanos_per_solution"}, ","))
	for wi := 0; wi < *workloadCount; wi++ {
		wl := fakeWorkload(rng, *workloadSize)
		for _, s := range solvers {
			for rep := 0; rep < *repitions; rep++ {
				dur := runWorkload(s, wl)
				fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v\n",
					solver.Name(s), wi, rep, *workloadSize,
					len(dict), *letterCount,
					dur.Nanoseconds(), math.Log10(float64(dur.Nanoseconds())/float64(*workloadSize)))
			}
		}
	}
}
