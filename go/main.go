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
	inputFile  string
	size       = -1
	amount     int
	silent     bool
	serverMode bool
)

func init() {
	flag.StringVar(&inputFile, "file", "<flag not provided>", "Input file containing initial board state")
	flag.IntVar(&amount, "amount", -1, "Number of iterations")
	flag.BoolVar(&silent, "silent", false, "If true, do not print")
	flag.BoolVar(&serverMode, "server", false, "Run in server mode")
}

func main() {
	flag.Parse()

	if serverMode {
		runServer()
		return
	}

	board := parseBoard(inputFile)
	buffer := initBoard(size)

	for i := 0; i < amount; i++ {
		board, buffer = nextState(board, buffer)
	}

	if !silent {
		printBoard(board)
	}
}

func initBoard(s int) [][]int {
	res := make([][]int, s)
	for i := range res {
		res[i] = make([]int, s)
	}
	return res
}

func parseBoard(file string) [][]int {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %s: %v", file, err))
	}

	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")

	size = len(lines)
	board := initBoard(size)

	for row, line := range lines {
		for col, ch := range line {
			board[row][col] = int(ch - '0')
		}
	}

	return board
}

func nextState(board, buffer [][]int) ([][]int, [][]int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if shouldCellLive(board, y, x) {
				buffer[y][x] = 1
			} else {
				buffer[y][x] = 0
			}
		}
	}
	return buffer, board
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
	for y := range board {
		for x := range board {
			strb.WriteString(strconv.Itoa(board[y][x]))
		}
		strb.WriteRune('\n')
	}
	file.WriteString(strb.String())
}

func runServer() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("READY")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "SHUTDOWN" {
			break
		}
		if strings.HasPrefix(line, "RUN ") {
			parts := strings.Fields(line[4:])
			if len(parts) >= 1 {
				f := parts[0]
				n := 1
				if len(parts) >= 2 {
					if v, err := strconv.Atoi(parts[1]); err == nil {
						n = v
					}
				}
				board := parseBoard(f)
				buffer := initBoard(size)
				for i := 0; i < n; i++ {
					board, buffer = nextState(board, buffer)
				}
				if !silent {
					printBoard(board)
				}
				fmt.Println("DONE")
			} else {
				fmt.Println("ERROR bad request")
			}
		} else {
			fmt.Println("ERROR unknown command")
		}
	}
}
