package main

import "testing"

func TestBoardMoveDirection(t *testing.T) {
	board, _ := makeBoardFrom(tilesOrdered)

	var res bool
	var sPos, target int
	// Move all the way to the right
	for i := 0; i < 3; i++ {
		sPos = board.spaceIdx
		target = board.tiles[sPos+1]
		res = board.moveRight()
		if !res || board.spaceIdx != sPos+1 || board.tiles[sPos] != target {
			t.Errorf("should have moved right and swapped values")
		}
	}

	sPos = board.spaceIdx
	res = board.moveRight()
	if res || board.spaceIdx != sPos {
		t.Errorf("invalid move beyound board boundaries: Right")
	}

	// Move all the way down
	for i := 0; i < 3; i++ {
		sPos = board.spaceIdx
		target = board.tiles[sPos+size]
		res = board.moveDown()
		if !res || board.spaceIdx != sPos+size || board.tiles[sPos] != target {
			t.Errorf("should have moved down")
		}
	}

	sPos = board.spaceIdx
	res = board.moveDown()
	if res || board.spaceIdx != sPos {
		t.Errorf("invalid move beyound board boundaries: Down")
	}

	// Move all the way up
	for i := 0; i < 3; i++ {
		sPos = board.spaceIdx
		target = board.tiles[sPos-size]
		res = board.moveUp()
		if !res || board.spaceIdx != sPos-size || board.tiles[sPos] != target {
			t.Errorf("should have moved up")
		}
	}

	sPos = board.spaceIdx
	res = board.moveUp()
	if res || board.spaceIdx != sPos {
		t.Errorf("invalid move beyound board boundaries: Up")
	}

	// Move all the way left
	for i := 0; i < 3; i++ {
		sPos = board.spaceIdx
		target = board.tiles[sPos-1]
		res = board.moveLeft()
		if !res || board.spaceIdx != sPos-1 || board.tiles[sPos] != target {
			t.Errorf("should have moved left")
		}
	}

	sPos = board.spaceIdx
	res = board.moveLeft()
	if res || board.spaceIdx != sPos {
		t.Errorf("invalid move beyound board boundaries: Left")
	}
}

func TestBoardMove(t *testing.T) {
	for _, dir := range [...]direction{left, right, up, down} {
		expected, _ := makeBoardFrom(tilesAlldirs)
		switch dir {
		case left:
			expected.moveLeft()
		case right:
			expected.moveRight()
		case up:
			expected.moveUp()
		case down:
			expected.moveDown()
		}
		current, _ := makeBoardFrom(tilesAlldirs)
		current.move(dir)
		if current.spaceIdx != expected.spaceIdx {
			t.Errorf("move and move%v should do the same", dir)
		}
	}
}

func TestBoardString(t *testing.T) {
	b, err := makeBoardFrom(tilesOrdered)
	if err != nil {
		t.Errorf("should be able to make board")
	}
	const expected = "[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]"
	if b.String() != expected {
		t.Errorf("wrong board string: %v", b)
		t.Errorf("expected: %s", expected)
	}
}

func TestBoardIsValid(t *testing.T) {
	testData := []struct {
		tiles []int
		valid bool
		msg   string
	}{
		{tilesMissing, false, "board missing tiles should not be valid"},
		{tilesWrong, false, "board with invalid tiles should not be valid"},
		{tilesMissing, false, "board with missing tiles should not be valid"},
		{tilesRepeated, false, "board with repeated tiles should not be valid"},
		{tilesExcess, true, "board created from excess tiles should be valid"},
		{tilesSolved, true, "solved board should be valid"},
	}
	for _, td := range testData {
		if b, _ := makeBoardFrom(td.tiles); b.isValid() != td.valid {
			t.Errorf(td.msg)
		}
	}

	if b := generateRandomBoard(); !b.isValid() {
		t.Errorf("generated board should be valid")
	}
}
