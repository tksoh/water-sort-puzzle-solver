package solverlib

import (
	"encoding/json"
	"fmt"
	"time"

	watersortpuzzle "github.com/pkositsyn/water-sort-puzzle-solver"
)

type (
	PuzzleSolver struct {
	}
)

func (p *PuzzleSolver) DoSolvePuzzle(initialStateStr string) bool {
	solver := watersortpuzzle.NewAStarSolver()

	var initialState watersortpuzzle.State
	if err := initialState.FromString(initialStateStr); err != nil {
		fmt.Printf("Invalid puzzle state provided: %s\n", err.Error())
		return false
	}

	t0 := time.Now()
	steps, err := solver.Solve(initialState)
	duration := time.Since(t0)
	if err != nil {
		fmt.Printf("Cannot solve puzzle: %s\n", err.Error())
		return false
	}
	fmt.Printf("Solution took: %v in %d steps\n", duration, len(steps))
	return true
}

func (p *PuzzleSolver) GetPuzzleSolution(initialStateStr string) string {
	solver := watersortpuzzle.NewAStarSolver()

	var initialState watersortpuzzle.State
	if err := initialState.FromString(initialStateStr); err != nil {
		fmt.Printf("Invalid puzzle state provided: %s\n", err.Error())
		return ""
	}

	t0 := time.Now()
	steps, err := solver.Solve(initialState)
	duration := time.Since(t0)
	if err != nil {
		fmt.Printf("Cannot solve puzzle: %s\n", err.Error())
		return ""
	}
	fmt.Printf("Solution took: %v in %d steps\n", duration, len(steps))

	if len(steps) > 0 {
		jsonStr, err := json.Marshal(steps)
		if err != nil {
			fmt.Println("error:", err)
			return ""
		}
		return string(jsonStr)
	} else {
		return "[]"
	}
}
