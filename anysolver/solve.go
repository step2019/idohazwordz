// This program runs one of the solvers interactively.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/step2018/idohazwordz/solver"
	"github.com/step2018/idohazwordz/words"
)

var (
	dictFile   = flag.String("dictionary", "/usr/share/dict/words", "File to use as a dictionary")
	size       = flag.Int("size", 16, "How many letters to expect")
	solverName = flag.String("solver", "List", fmt.Sprintf("Solver strategy to use. Must be one of %v", solver.Names(solver.AllSolvers)))
)

func main() {
	flag.Parse()
	s, err := solver.New(*solverName)
	if err != nil {
		log.Fatalf("invalid solver %v: %v", *solverName, err)
	}
	if err := s.Init(words.Load(*dictFile, *size)); err != nil {
		log.Fatalf("error initializing solver: %v", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Input letters: ")
		input, err := reader.ReadString('\n')
		if input == "\n" || err != nil {
			break
		}
		fmt.Printf("Best answer: %s\n", s.Solve(words.Normalize(input)))
	}
}
