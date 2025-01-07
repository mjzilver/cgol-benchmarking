package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	inputFile string
	size      int
	amount    int
	silent    bool
)

func init() {
	flag.StringVar(&inputFile, "file", "input.txt", "Input file containing initial board state")
	flag.IntVar(&size, "size", 20, "Size of the board (square)")
	flag.IntVar(&amount, "amount", 30, "Number of iterations")
	flag.BoolVar(&silent, "silent", false, "If true, do not print")
}

func main() {
	flag.Parse()

	board := parseBoard(inputFile, size)

	for _ = range amount {
		board = nextState(board, size)
	}

	if !silent {
		printBoard(board)
	}
}

func initBoard(size int) [][]int {
	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}
	return res
}

func parseBoard(filePath string, size int) [][]int {
	res := initBoard(size)

	file, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Sprintf("File %s not found!", filePath))
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	row := 0

	for fileScanner.Scan() {
		for i, ch := range fileScanner.Text() {
			res[row][i] = int(ch - '0')
		}
		row = row + 1
	}

	return res
}

func nextState(board [][]int, size int) [][]int {
	res := initBoard(size)

	for y := range len(board) {
		for x := range len(board[y]) {
			if shouldCellLive(board, y, x, size) {
				res[y][x] = 1
			}
		}
	}

	return res
}

var (
	neighbors = [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
)

func shouldCellLive(board [][]int, y, x, size int) bool {
	count := 0

	for _, n := range neighbors {
		count = count + board[(y+n[0]+size)%size][(x+n[1]+size)%size]
	}

	if board[y][x] == 1 {
		if count < 2 || count > 3 {
			return false
		}
	} else {
		if count == 3 {
			return true
		}
	}

	return board[y][x] == 1
}

func printBoard(board [][]int) {
	fmt.Println(strings.Repeat("-", len(board)*2))

	for y := range len(board) {
		fmt.Printf("%v\n", board[y])
	}
	fmt.Println(strings.Repeat("-", len(board)*2))
}
