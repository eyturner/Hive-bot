package main

type PlayerType int64

const (
	Human PlayerType = iota
	Computer
)

type Player struct {
	playerType PlayerType
	nest       []Piece
}
