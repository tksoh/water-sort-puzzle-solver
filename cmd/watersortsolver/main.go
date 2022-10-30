package main

import (
	"flag"
	"fmt"
	"time"

	watersortpuzzle "github.com/pkositsyn/water-sort-puzzle-solver"
)

var algorithmType = flag.String("algorithm", "astar",
	`Algorithm to solve with. Choices: [astar, idastar, dijkstra]`)

func main() {
	flag.Parse()
	fmt.Println("Input initial puzzle state")

	var initialStateStr string = "JHGFF;CAHEI;FHBBG;DDHEC;IFJGD;ECAAE;JJIFD;BJAHB;DBCII;GECAG;;"

	if initialStateStr == "" {
		n, err := fmt.Scanln(&initialStateStr)
		if err != nil {
			fmt.Printf("Error getting input: %s\n", err.Error())
			return
		}
		if n != 1 {
			fmt.Printf("Scanned %d values but needed one position\n", n)
			return
		}
	}

	var solver watersortpuzzle.Solver
	switch *algorithmType {
	case "astar":
		solver = watersortpuzzle.NewAStarSolver()
	case "idastar":
		solver = watersortpuzzle.NewIDAStarSolver()
	case "dijkstra":
		solver = watersortpuzzle.NewDijkstraSolver()
	}

	var initialState watersortpuzzle.State
	if err := initialState.FromString(initialStateStr); err != nil {
		fmt.Printf("Invalid puzzle state provided: %s\n", err.Error())
		return
	}

	t0 := time.Now()
	steps, err := solver.Solve(initialState)
	duration := time.Since(t0)
	if err != nil {
		fmt.Printf("Cannot solve puzzle: %s\n", err.Error())
		return
	}

	suffix := ""
	if statsSolver, ok := solver.(watersortpuzzle.SolverWithStats); ok {
		suffix = fmt.Sprintf(" Algorithm took %d iterations to find solution.", statsSolver.Stats().Steps)
	}

	fmt.Printf("Puzzle solved in %d steps!%s\n", len(steps), suffix)
	fmt.Printf("Solution took: %v\n", duration)
	for _, step := range steps {
		fmt.Println(step.From+1, step.To+1)
	}
}
