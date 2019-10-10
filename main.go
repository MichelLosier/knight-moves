package main

import (
	"bufio"
	"fmt"
	chessboard "knight-moves/chessboard"
	"os"
	"strings"
)

func main() {
	var parsedArgs []string
	reader := bufio.NewReader(os.Stdin)

	rawArgs, readErr := reader.ReadString('\n')

	if readErr != nil {
		fmt.Printf("Error reading input: %s", readErr.Error())
		return
	}

	parsedArgs = parseArguments(rawArgs)

	bDomain := []int{0, 8}
	bRange := []int{0, 8}

	board := chessboard.NewBoard(bDomain, bRange)
	board.SetRestrictedSquares(parsedArgs[2:])

	startingSquare, _ := chessboard.NewBoardSquareFromString(parsedArgs[0], &board)
	targetSquare, _ := chessboard.NewBoardSquareFromString(parsedArgs[1], &board)

	path := chessboard.FindShortestKnightPath(startingSquare, targetSquare)
	printPath(path)

}

//TODO or return error if validation fails
func parseArguments(rawArgs string) []string {
	argsSlice := strings.Fields(rawArgs)
	//validation
	return argsSlice
}

func printPath(path []*chessboard.BoardSquare) {
	var pathString string

	for _, node := range path {
		pathString = fmt.Sprintf("%s %s", pathString, node.GetStringCoords())
	}

	fmt.Println(pathString)
}
