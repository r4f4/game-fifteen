package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const empty = 0
const size = 4

type board struct {
	tiles    [16]int
	spaceIdx int
}

func (b board) String() string {
	return fmt.Sprint(b.tiles)
}

func makeBoardFrom(tiles []int) (board, error) {
	b := board{}
	nTiles := size * size
	if c := copy(b.tiles[:], tiles); c < nTiles {
		return b, errors.New(fmt.Sprintf("expected %d tiles, got %d instead", nTiles, c))
	}
	b.spaceIdx = b.emptyTileIndex()
	return b, nil
}

func generateRandomBoard() board {
	b, _ := makeBoardFrom(tilesSolved)
	rand.Shuffle(len(b.tiles), func(i, j int) {
		// Swap contents
		b.tiles[i], b.tiles[j] = b.tiles[j], b.tiles[i]
	})
	b.spaceIdx = b.emptyTileIndex()
	return b
}

// Make a deep-copy of a Board
func (b *board) copy() board {
	return board{b.tiles, b.spaceIdx}
}

// Finds the index of the empty tile in the board
func (b *board) emptyTileIndex() int {
	for i, tile := range b.tiles {
		if tile == 0 {
			return i
		}
	}
	// XXX:Should never happen
	return -1
}

func (b *board) isValid() bool {
	max := size*size - 1
	// Make sure it contains all tiles from [0, size * size - 1]
	set := make(map[int]bool, max+1)
	for _, tile := range b.tiles {
		if tile > max {
			return false
		}
		if _, ok := set[tile]; ok {
			return false
		}
		set[tile] = true
	}
	return true
}

// Moves the empty tile's position at the given direction by 1 and swaps the contents
func (b *board) makeMove(dir int) bool {
	index := b.spaceIdx + dir
	if index < 0 || index >= len(b.tiles) {
		return false
	}
	// Swap values
	b.tiles[b.spaceIdx], b.tiles[index] = b.tiles[index], empty
	// Update empty tile's position
	b.spaceIdx = index
	return true
}

type direction int

const (
	left direction = -1
	right = 1
	up = -size
	down = size
)

var opposites = map[direction]direction{left: right, right: left, up: down, down: up}

func (d direction) String() string {
	switch d {
	case left:
		return "left"
	case right:
		return "right"
	case up:
		return "up"
	case down:
		return "down"
	default:
		return "invalid"
	}
}

// Convenience method to be used in loops
func (b *board) move(d direction) bool {
	switch d {
	case left, right:
		if (b.spaceIdx / size) != ((b.spaceIdx + int(d)) / size) {
			return false
		}
		fallthrough
	case up, down:
		return b.makeMove(int(d))
	default:
		return false
	}
}
