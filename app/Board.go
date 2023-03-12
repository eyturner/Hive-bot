package main

type Hex struct {
	q int
	r int
	z int
}

type Board struct {
	grid map[Hex]Piece
}

// May rename since FEN doesn't apply here.
func (b Board) importFromFEN(fen string) Board {
	// Code to setup board from FEN ()
}

func (b Board) printASCII() string {
	// Code to print an ascii version of the current grid
}

func (b *Board) move(p Piece, newHex Hex) Board {
	// update board to reflect move
}
