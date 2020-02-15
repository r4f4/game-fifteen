package main

import (
	"fmt"
	"testing"
)

func TestSolvable(t *testing.T) {
	testData := []struct {
		tiles    []int
		solvable bool
		msg      string
	}{
		{tilesSolved, true, "solved board should be solvable"},
		{tilesUnsolved, true, "unsolved board should be solvable"},
		{tilesImpossible, false, "impossible board should not be solvable"},
	}
	for _, td := range testData {
		if b, _ := makeBoardFrom(td.tiles); solvable(&b) != td.solvable {
			t.Errorf(td.msg)
		}
	}
}

func TestCheckSolved(t *testing.T) {
	testData := []struct {
		tiles  []int
		solved bool
		msg    string
	}{
		{tilesUnsolved, false, "unsolved board should not be solved"},
		{tilesSolved, true, "solved board should be solved"},
		{tilesAlmost, false, "almost board should not be solved"},
	}
	for _, td := range testData {
		if b, _ := makeBoardFrom(td.tiles); solved(&b) != td.solved {
			t.Errorf(td.msg)
		}
	}
}

func TestAStarSerial(t *testing.T) {
	b, _ := makeBoardFrom(tilesSolved)
	sol := aStarSerial(&b)
	if sol == nil {
		t.Errorf("aStar should not return nil to solved board")
	}
	if !solved(sol.state) {
		t.Errorf("final state should be solved")
	}
	if len(sol.moves) != 0 {
		t.Errorf("no moves should be needed")
	}
	if sol.cost != 0 {
		t.Errorf("cost of solved board should be 0")
	}

	b, _ = makeBoardFrom(tilesUnsolved)
	sol = aStarSerial(&b)
	if sol == nil {
		t.Errorf("aStar should not return nil to unsolved board")
	}
	fmt.Println(sol.state.String())
	if !solved(sol.state) || !sol.state.isValid() {
		t.Errorf("final state should be solved")
		t.Errorf(sol.state.String())
	}
	if len(sol.moves) == 0 {
		t.Errorf("cannot solve unsolved with 0 moves")
	}
	// apply the moves and check final state
	for _, dir := range sol.moves {
		b.move(dir)
	}
	if !solved(&b) {
		t.Errorf("movement list does not result in solution")
	}
	fmt.Println(sol.cost, sol.moves)
}
