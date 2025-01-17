package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const outputFile = "output.txt"

var (
	inputFile string
	size      = -1
	amount    int
	silent    bool
)

func init() {
	flag.StringVar(&inputFile, "file", "<flag not provided>", "Input file containing initial board state")
	flag.IntVar(&amount, "amount", -1, "Number of iterations")
	flag.BoolVar(&silent, "silent", false, "If true, do not print")
}

func main() {
	flag.Parse()

	board := parseBoard()

	for range amount {
		board = nextState(board)
	}

	if !silent {
		printBoard(board)
	}
}

func initBoard() [][]int {
	res := make([][]int, size)
	for i := range res {
		res[i] = make([]int, size)
	}
	return res
}

func parseBoard() [][]int {
	var res [][]int

	file, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Sprintf("File %s not found!", inputFile))
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	row := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		if size == -1 {
			size = len(str)
			res = initBoard()
		}

		for i, ch := range str {
			res[row][i] = int(ch - '0')
		}
		row = row + 1
	}

	return res
}

func nextState(board [][]int) [][]int {
	res := initBoard()

	for y := range size {
		for x := range size {
			if shouldCellLive(board, y, x) {
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

func shouldCellLive(board [][]int, y, x int) bool {
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
	file, _ := os.Create(outputFile)
	defer file.Close()

	var strb strings.Builder
	for y := range len(board) {
		for x := range len(board) {
			strb.WriteString(strconv.Itoa(board[y][x]))
		}
		strb.WriteRune('\n')
	}
	file.WriteString(strb.String())
}
