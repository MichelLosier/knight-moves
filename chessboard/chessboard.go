package chessboard

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BoardSquare struct {
	Xcoord   int
	Ycoord   int
	Board    *Board
	Previous *BoardSquare
}

func (b BoardSquare) Equals(boardSquare BoardSquare) bool {
	if (b.Xcoord == boardSquare.Xcoord) && (b.Ycoord == boardSquare.Ycoord) {
		return true
	}
	return false
}

func (b BoardSquare) GetStringCoords() string {
	return fmt.Sprintf("%s%s", IntCoordToString(b.Xcoord), strconv.Itoa(b.Ycoord))
}

func (b *BoardSquare) TracePath() []*BoardSquare {
	squares := []*BoardSquare{b}
	previous := b.Previous
	if previous != nil {
		previousSquares := previous.TracePath()
		squares = append(previousSquares, squares...)
	}
	return squares
}

func NewBoardSquare(coordinates []int, board *Board) (BoardSquare, error) {

	isInBounds := board.IsInBounds(coordinates)

	if !isInBounds {
		return BoardSquare{}, errors.New("Board Square coordinates out of bounds of board size")
	}

	newBoardSquare := BoardSquare{
		Xcoord: coordinates[0],
		Ycoord: coordinates[1],
		Board:  board,
	}

	return newBoardSquare, nil
}

func NewBoardSquareFromString(coordinates string, board *Board) (BoardSquare, error) {
	parsedCoordinates := strings.Split(coordinates, "")

	xCoordInt := StringCoordToInt(parsedCoordinates[0])
	yCoordInt, _ := strconv.Atoi(parsedCoordinates[1])

	intCoordinates := []int{xCoordInt, yCoordInt}
	return NewBoardSquare(intCoordinates, board)
}

func (b BoardSquare) FindNextKnightMoves() []BoardSquare {
	nextMoveSquares := []BoardSquare{}
	moveSets := [][]int{ //inelegant, i know
		{2, 1},
		{2, -1},
		{-2, 1},
		{-2, -1},
		{1, 2},
		{-1, 2},
		{1, -2},
		{-1, -2},
	}

	for _, moveSet := range moveSets {
		newCoord := []int{
			(b.Xcoord + moveSet[0]),
			(b.Ycoord + moveSet[1]),
		}
		if b.Board.IsInBounds(newCoord) {
			boardSquare, _ := NewBoardSquare(newCoord, b.Board)

			if !b.Board.IsRestrictedSquare(boardSquare) {
				nextMoveSquares = append(nextMoveSquares, boardSquare)
			}
		}
	}

	return nextMoveSquares

}

type BoardSquareQueue struct {
	items []BoardSquare
}

func (q *BoardSquareQueue) Enqueue(boardSquare BoardSquare) {
	newStart := []BoardSquare{boardSquare}
	q.items = append(newStart, q.items[:q.Size()]...)
}

func (q *BoardSquareQueue) Dequeue() (BoardSquare, error) {
	size := q.Size()
	if size == 0 {
		return BoardSquare{}, errors.New("Empty Queue")
	}
	nextInt := q.items[size-1]
	q.items = q.items[:size-1]
	return nextInt, nil
}

func (q *BoardSquareQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *BoardSquareQueue) Size() int {
	return len(q.items)
}

func FindShortestKnightPath(startingSquare BoardSquare, targetSquare BoardSquare) []*BoardSquare {
	var searchQueue BoardSquareQueue
	searchQueue.Enqueue(startingSquare)
	path := []*BoardSquare{}

	for !searchQueue.IsEmpty() {
		currentSquare, _ := searchQueue.Dequeue()
		if currentSquare.Equals(targetSquare) {
			return currentSquare.TracePath()
		}
		nextSquares := currentSquare.FindNextKnightMoves()
		for _, square := range nextSquares {
			square.Previous = &currentSquare
			searchQueue.Enqueue(square)
		}
	}
	return path
}

type Board struct {
	Domain            []int //[0,8]
	Range             []int //[0,8]
	RestrictedSquares []BoardSquare
}

func (b Board) IsInBounds(coordinates []int) bool {
	if (coordinates[0] >= b.Domain[0]) &&
		(coordinates[0] <= b.Domain[1]) &&
		(coordinates[1] >= b.Range[0]) &&
		(coordinates[1] <= b.Range[1]) {
		return true
	}
	return false
}

func (b Board) IsRestrictedSquare(boardSquare BoardSquare) bool {
	for _, rSquare := range b.RestrictedSquares {
		if rSquare.Equals(boardSquare) {
			return true
		}
	}
	return false
}

func (b *Board) SetRestrictedSquares(coordinates []string) {
	var restrictedSquares []BoardSquare
	for _, coord := range coordinates {
		boardSquare, _ := NewBoardSquareFromString(coord, b)
		restrictedSquares = append(restrictedSquares, boardSquare)
	}

	b.RestrictedSquares = restrictedSquares
}

func NewBoard(bDomain []int, bRange []int) Board {
	newBoard := Board{
		Domain: bDomain,
		Range:  bRange,
	}
	return newBoard
}

const asciiOffset = 97

func StringCoordToInt(coord string) int {
	lowerString := strings.ToLower(coord)
	runeCode := []rune(lowerString)[0]
	runeInt := int(runeCode)
	return (runeInt - asciiOffset)
}

func IntCoordToString(coord int) string {
	runeCode := rune(coord + asciiOffset)
	coordString := string(runeCode)

	return coordString
}
