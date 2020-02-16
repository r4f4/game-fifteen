package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
)

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

func printGameReplay(b *board, moves []direction) {
	for _, dir := range moves {
		fmt.Println(dir)
		b.move(dir)
		printBoard(b)
	}
}

func main() {
	generate := flag.Bool("random", false, "Generate a random board")
	replay := flag.Bool("replay", false, "Play the game move by move")
	flag.Parse()

	var b board
	if *generate {
		fmt.Println("Generating a random board:")
		b = generateRandomBoard()
	} else {
		var err error
		reader := bufio.NewReader(os.Stdin)
		tiles := [16]int{}
		for i := 0; i < size; i++ {
			idx := i * size
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("error reading board: %v\n", err)
				return
			}
			values := [4]int{}
			fmt.Sscanf(text, "%d %d %d %d", &values[0], &values[1], &values[2], &values[3])
			for j := 0; j < size; j++ {
				tiles[idx + j] = values[j]
			}
		}
		b, err = makeBoardFrom(tiles[:])
		if err != nil || !b.isValid() {
			fmt.Printf("invalid board supplied: %v\n", err)
			return
		}
	}
	printBoard(&b)
	if !solvable(&b) {
		fmt.Println("Impossible to solve board")
		return
	}
	if sol := aStarSerial(&b); sol != nil {
		if *replay {
			printGameReplay(&b, sol.moves)
		} else {
			fmt.Println(sol.moves)
		}
	} else {
		fmt.Println("Could not solve board")
	}
}
