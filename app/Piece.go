package main

type PieceType int64

const (
	QueenBee PieceType = iota
	Ant
	Beetle
	Grasshopper
	Ladybug
	Spider
	Mosquito
	Pillbug
)

type Piece struct {
	color     string
	pieceType PieceType
	name      string
}

func (p Piece) generateMoves(grid Grid) string {
	// Code for finding moves for a specific piece
	return "wow!"
}
