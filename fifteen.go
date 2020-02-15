package main

import "fmt"

// Print a board, one row per line; each row as an array
func printBoard(b *board) {
	if len(b.tiles) == 0 {
		fmt.Println("[]")
		return
	}

	for i := 0; i < size; i++ {
		fmt.Println(b.tiles[i*size : i*size+size])
	}
}

func main() {
	fmt.Println("Generating a random board:")
	board := generateRandomBoard()
	printBoard(&board)
	if !solvable(&board) {
		fmt.Println("Impossible to solve board")
		return
	}
	if sol := aStarSerial(&board); sol != nil {
		fmt.Println(sol.moves)
	} else {
		fmt.Println("Could not solve board")
	}
}
