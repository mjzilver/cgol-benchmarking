package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	size       int
	amount     int
	iterations int
	chance     int
	timeoutSec int
	silent     bool
	generate   bool
)

type Exe struct {
	Name string
	Args []string
	Dir  string

	Average    float32
	Benchmarks []int64
	TotalTime  int64
}

const (
	boardFile = "../input.txt"
)

func init() {
	flag.IntVar(&amount, "amount", 10, "Amount of generations to benchmark")
	flag.IntVar(&iterations, "iterations", 10, "Amount of times to run the binary to get an average")
	flag.IntVar(&size, "size", 20, "Size of the board (always square)")
	flag.IntVar(&chance, "chance", 50, "Chance for a random cell to be alive")
	flag.BoolVar(&generate, "generate", false, "If true, generate a new board")
	flag.BoolVar(&silent, "silent", false, "If true, do not print an output file")
	flag.IntVar(&timeoutSec, "timeout", 30, "Timeout to cancel iteration (in seconds)")
}

func main() {
	flag.Parse()

	exes := []Exe{
		{Name: "Go", Args: []string{"./bin/cgol"}, Dir: "./go"},
		{Name: "C", Args: []string{"./bin/cgol"}, Dir: "./c"},
		{Name: "Perl", Args: []string{"perl", "./cgol.pl"}, Dir: "./perl"},
		{Name: "OCaml", Args: []string{"./bin/cgol"}, Dir: "./ocaml"},
		{Name: "Rust", Args: []string{"./bin/cgol"}, Dir: "./rust"},
		{Name: "C#", Args: []string{"./bin/cgol"}, Dir: "./csharp"},
	}

	if generate {
		genBoard(size)
	}

	stdArgs := []string{
		"--file", boardFile,
		"--amount", fmt.Sprintf("%d", amount),
	}
	if silent {
		stdArgs = append(stdArgs, "--silent")
	}

	var wg sync.WaitGroup
	for i := range exes {
		wg.Add(1)
		go func(exe *Exe) {
			defer wg.Done()
			benchmarkExe(exe, stdArgs)
			exe.Average = float32(exe.TotalTime) / float32(len(exe.Benchmarks))
		}(&exes[i])
	}
	wg.Wait()

	sort.Slice(exes, func(i, j int) bool {
		return exes[i].Average < exes[j].Average
	})

	tableOutput := fmt.Sprintf("Benchmark for a %dx%d board with %d generations with a %d sec timeout\n", size, size, amount, timeoutSec)
	tableOutput += fmt.Sprintf("| %-12s | %-12s | %-12s |\n", "Language", "Avg Time ms", "Iterations")
	tableOutput += fmt.Sprintf("|%s|%s|%s|\n", strings.Repeat("-", 14), strings.Repeat("-", 14), strings.Repeat("-", 14))

	for _, exe := range exes {
		avgMilli := exe.Average / 1000
		tableOutput += fmt.Sprintf("| %-12s | %-12.2f | %-12d |\n", exe.Name, avgMilli, len(exe.Benchmarks))
	}

	fmt.Println(tableOutput)

	os.Mkdir("./logs", 0755)
	logFileName := fmt.Sprintf("./logs/benchmark_%d.log", time.Now().Unix())
	logFile, err := os.Create(logFileName)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	_, err = logFile.WriteString(tableOutput)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func benchmarkExe(exe *Exe, stdArgs []string) {
	exe.Args = append(exe.Args, stdArgs...)

	for range iterations {
		cmd := exec.Command(exe.Args[0], exe.Args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = exe.Dir

		startTime := time.Now().UnixMicro()
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error executing %s: %s\n", exe.Args[0], err)
			panic(err)
		}
		endTime := time.Now().UnixMicro()

		exe.Benchmarks = append(exe.Benchmarks, endTime-startTime)
		exe.TotalTime = exe.TotalTime + endTime - startTime
		if exe.TotalTime > int64(timeoutSec*1000*1000) {
			return
		}
	}
}

func genBoard(size int) {
	file, _ := os.Create(boardFile)
	defer file.Close()

	var strb strings.Builder
	for range size {
		for range size {
			if chance > rand.Intn(100) {
				strb.WriteRune('1')
			} else {
				strb.WriteRune('0')
			}
		}
		strb.WriteRune('\n')
	}
	file.WriteString(strb.String())
}
