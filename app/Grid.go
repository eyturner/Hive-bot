package main

type Hex struct {
	q int
	r int
	z int
}

type Grid map[Hex]Piece

type Move struct {
	piece Piece
	hex   Hex
}

// IGN: Inline Grid Notation
func importFromIGN(fen string) Grid {
	// Pick an orientation of the board

	// If this piece has a flat top in your chosen orientation, find the column that is the left-most and the piece at the top of that column
	// you will be reading up --> down and then left --> right

	// If this piece has a pointed top in your chosen orientation, find the row that is top-most and the piece at the left of that row
	// you will be reading left --> right and then up --> down

	// Flat orientation:

	newGrid := Grid()
	return newGrid
	// Code to setup board from FEN ()
}

func (grid Grid) rotate90() {
	// Code to rotate the grid 90 degrees
}

func (grid Grid) printASCII() string {
	// Code to print an ascii version of the current grid
	return "wow!"
}

func (grid *Grid) move(piece Piece, prevHex Hex, newHex Hex) {
	// update board to reflect move
}

func (grid *Grid) addToGrid(newPiece Piece, hex Hex) {

}
