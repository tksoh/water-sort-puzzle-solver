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

func (p *PuzzleSolver) GetPuzzleSolutionMap(initialStateStr string) string {
	solver := watersortpuzzle.NewAStarSolver()

	var initialState watersortpuzzle.State
	if err := initialState.FromString(initialStateStr); err != nil {
		fmt.Printf("Invalid puzzle state provided: %s\n", err.Error())
		jsonStr := makeSolutionJson("invalid", 0, nil, err.Error())
		return jsonStr
	}

	t0 := time.Now()
	steps, err := solver.Solve(initialState)
	duration := time.Since(t0)
	if err != nil {
		fmt.Printf("Cannot solve puzzle: %s\n", err.Error())
		jsonStr := makeSolutionJson("unsolved", duration, nil, err.Error())
		return jsonStr
	}

	fmt.Printf("Solution took: %v in %d steps\n", duration, len(steps))

	if verifySolution(initialState, steps) {
		fmt.Printf("Solution verification PASS!\n")
	} else {
		fmt.Printf("Solution verification FAIL!\n")
	}

	jsonStr := makeSolutionJson("solved", duration, steps, "")
	return jsonStr
}

func makeSolutionJson(status string, duration time.Duration,
	steps []watersortpuzzle.Step, error string) string {
	dataMap := map[string]interface{}{
		"status":   status,
		"duration": duration,
		"steps":    steps,
		"message":  error,
	}

	jsonStr, err := json.Marshal(dataMap)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return string(jsonStr)
}

func verifySolution(initialState watersortpuzzle.State, steps []watersortpuzzle.Step) bool {
	var err error
	state := initialState
	count := 0
	for _, step := range steps {
		if state, err = state.Step(step); err != nil {
			fmt.Printf("Verification FAIL at Step #%d: [%d] -> [%d]\n", count, step.From+1, step.To+1)
			return false
		}
		count++
	}

	return state.IsTerminal()
}
