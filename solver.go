package main

import (
	"fmt"
	"math"
	"sort"
	"container/heap"
)

type solution struct {
	state *board
	moves []direction
	cost  int
}

func (s *solution) String() string {
	return fmt.Sprintf("%v\n%v\n%v", s.state, s.moves, s.cost)
}

// A MinHeap implements heap.Interface and holds Solution items
type minHeap []*solution

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*solution)) }

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*h = old[0 : n-1]
	return item
}

// Calculate the Manhattan distance of a tile with value `val` at index `idx`
func manhattanDistance(val, idx int) int {
	// Linear distance from where tile should be
	diff := int(math.Abs(float64(val - 1 - idx)))
	return (diff / size) + (diff % size) // # of rows + cols to move
}

// Calculates the cost of (possibly) solving the current Board's state
// The cost is defined as `cost = moves + estimate`, where
// `moves` is the number of moves made so far
// `estimate` is the Manhattan Distance of moving the remaining out-of-place tiles
func calcManhattanDistance(b *board) int {
	cost := 0
	for i, tile := range b.tiles {
		// Tile is in the right position
		if tile == 0 || tile == i+1 {
			continue
		}
		cost += manhattanDistance(tile, i)
	}
	return cost
}

func aStarSerial(b *board) *solution {
	start := &solution{state: b, moves: []direction{}, cost: calcManhattanDistance(b)}
	h := &minHeap{start}
	heap.Init(h)

	explored := make(map[string]int)
	explored[b.String()] = start.cost
	nIgnored := 0

	for h.Len() > 0 {
		sol := heap.Pop(h).(*solution)

		if solved(sol.state) {
			fmt.Printf("Explored: %d Ignored: %d\n", len(explored), nIgnored)
			return sol
		}

		nMoves := len(sol.moves)
		oldIdx := sol.state.spaceIdx
		// Try moving in all possible directions
		for _, dir := range []direction{left, right, up, down} {
			// Do not undo the last move
			if nMoves > 0 && opposites[dir] == sol.moves[nMoves-1] {
				continue
			}
			// Temporarily modify the board so we do not copy if the move is impossible
			if !sol.state.move(dir) {
				continue
			}
			str := sol.state.String()
			tile := sol.state.tiles[oldIdx]
			diff := manhattanDistance(tile, oldIdx) - manhattanDistance(tile, sol.state.spaceIdx)
			copyCost := sol.cost + diff + 1 // one more move
			if cost, ok := explored[str]; ok && copyCost > cost {
				nIgnored++
				continue
			}
			explored[str] = copyCost
			copyBoard := sol.state.copy()
			// Restore board to previous state
			sol.state.move(opposites[dir])
			copyMoves := append(append([]direction{}, sol.moves...), dir)
			newSol := solution{state: &copyBoard, moves: copyMoves, cost: copyCost}
			heap.Push(h, &newSol)
		}
	}
	// No solution found
	return nil
}

// https://www.cs.bham.ac.uk/~mdr/teaching/modules04/java2/TilesSolvability.html
// If empty tile is in an even row and the number of inversions is odd ---> then it is solvable
// If empty tile is in an odd row and the number of inversions is even ---> then it is solvable
// It is not solvable otherwise
func solvable(b *board) bool {
	nInversions := 0
	for i, tile := range b.tiles {
		if tile == 0 {
			continue
		}
		for _, next := range b.tiles[i+1:] {
			if next == 0 {
				continue
			}
			if tile > next {
				nInversions++
			}
		}
	}
	isEvenRow := (b.spaceIdx/size)%2 == 0
	isEvenInversions := nInversions%2 == 0
	return (isEvenRow && !isEvenInversions) || (!isEvenRow && isEvenInversions)
}

// Check if a board is in a solved state
func solved(b *board) bool {
	lastIdx := len(b.tiles) - 1
	return b.spaceIdx == lastIdx && sort.IntsAreSorted(b.tiles[:lastIdx])
}
