package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
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
	PeakMemory int64
}

const (
	boardFile = "./input.txt"
)

func humanBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

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
		{Name: "C++", Args: []string{"./bin/cgol"}, Dir: "./cpp"},
		{Name: "Perl", Args: []string{"perl", "./cgol.pl"}, Dir: "./perl"},
		{Name: "Python", Args: []string{"python3", "./cgol.py"}, Dir: "./python"},
		{Name: "OCaml", Args: []string{"./bin/cgol"}, Dir: "./ocaml"},
		{Name: "Rust", Args: []string{"./bin/cgol"}, Dir: "./rust"},
		{Name: "C#", Args: []string{"./bin/dotnet/cgol"}, Dir: "./csharp"},
		{Name: "F#", Args: []string{"./bin/cgol"}, Dir: "./fsharp"},
		{Name: "Java", Args: []string{"java", "-jar", "./bin/cgol.jar"}, Dir: "./java"},
		{Name: "Node.js", Args: []string{"node", "./cgol.js"}, Dir: "./nodejs"},
	}

	if generate {
		genBoard(size)
	}

	stdArgs := []string{
		"--file", fmt.Sprintf(".%s", boardFile), // Make ./ into ../
		"--amount", fmt.Sprintf("%d", amount),
	}
	if silent {
		stdArgs = append(stdArgs, "--silent")
	}

	for i := range exes {
		benchmarkExe(&exes[i], stdArgs)
		if len(exes[i].Benchmarks) > 0 {
			exes[i].Average = float32(exes[i].TotalTime) / float32(len(exes[i].Benchmarks))
		}
	}

	sort.Slice(exes, func(i, j int) bool {
		return exes[i].Average < exes[j].Average
	})

	tableOutput := fmt.Sprintf("Benchmark for a %dx%d board with %d generations with a %d sec timeout\n", size, size, amount, timeoutSec)
	tableOutput += fmt.Sprintf("| %-12s | %-12s | %-12s | %-12s |\n", "Language", "Avg Time ms", "Peak Mem", "Iterations")
	tableOutput += fmt.Sprintf("|%s|%s|%s|%s|\n", strings.Repeat("-", 14), strings.Repeat("-", 14), strings.Repeat("-", 14), strings.Repeat("-", 14))

	for _, exe := range exes {
		avgMilli := exe.Average / 1000
		memBytes := uint64(exe.PeakMemory) * 1024
		tableOutput += fmt.Sprintf("| %-12s | %-12.2f | %-12s | %-12d |\n", exe.Name, avgMilli, humanBytes(memBytes), len(exe.Benchmarks))
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
	fmt.Printf("Starting benchmark: %s\n", exe.Name)

	filePath, reqAmount := parseStdArgs(stdArgs)

	args := append([]string{}, exe.Args...)
	args = append(args, "--server")
	if silent {
		args = append(args, "--silent")
	}

	cmd, stdin, reader, err := startServer(args, exe.Dir)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}

	defer func() {
		cleanupServer(cmd, stdin)
	}()

	if err := waitReady(reader); err != nil {
		fmt.Println("Error waiting for READY:", err)
		cmd.Process.Kill()
		return
	}

	runIterations(reader, stdin, exe, filePath, reqAmount, cmd)
}

func parseStdArgs(stdArgs []string) (string, int) {
	filePath := ""
	reqAmount := amount
	for i := 0; i < len(stdArgs); i++ {
		if stdArgs[i] == "--file" && i+1 < len(stdArgs) {
			filePath = stdArgs[i+1]
			i++
		} else if stdArgs[i] == "--amount" && i+1 < len(stdArgs) {
			if v, err := strconv.Atoi(stdArgs[i+1]); err == nil {
				reqAmount = v
			}
			i++
		}
	}
	return filePath, reqAmount
}

func startServer(args []string, dir string) (*exec.Cmd, io.WriteCloser, *bufio.Reader, error) {
	cmd := exec.Command(args[0], args[1:]...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	if err := cmd.Start(); err != nil {
		return nil, nil, nil, err
	}

	reader := bufio.NewReader(stdout)
	return cmd, stdin, reader, nil
}

func cleanupServer(cmd *exec.Cmd, stdin io.WriteCloser) {
	if stdin != nil {
		stdin.Write([]byte("SHUTDOWN\n"))
		stdin.Close()
	}
	if cmd.Process != nil {
		cmd.Process.Kill()
		cmd.Wait()
	}
}

func waitReady(reader *bufio.Reader) error {
	readyCh := make(chan string, 1)
	errCh := make(chan error, 1)
	go func() {
		s, err := reader.ReadString('\n')
		if err != nil {
			errCh <- err
			return
		}
		readyCh <- s
	}()

	select {
	case s := <-readyCh:
		if strings.TrimSpace(s) != "READY" {
			return fmt.Errorf("server did not signal READY, got: %s", s)
		}
		return nil
	case e := <-errCh:
		return e
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout waiting for server READY")
	}
}

func runIterations(
	reader *bufio.Reader,
	stdin io.WriteCloser,
	exe *Exe,
	filePath string,
	reqAmount int,
	cmd *exec.Cmd,
) {
	for i := 0; i < iterations; i++ {
		fmt.Printf("\r  Progress: %d/%d ", i+1, iterations)
		os.Stdout.Sync()

		req := fmt.Sprintf("RUN %s %d\n", filePath, reqAmount)
		startTime := time.Now().UnixMicro()
		if _, err := stdin.Write([]byte(req)); err != nil {
			fmt.Println("Error writing to server stdin:", err)
			return
		}

		respCh := make(chan string, 1)
		respErr := make(chan error, 1)
		go func() {
			s, err := reader.ReadString('\n')
			if err != nil {
				respErr <- err
				return
			}
			respCh <- s
		}()

		select {
		case <-respCh:
			endTime := time.Now().UnixMicro()
			dur := endTime - startTime
			exe.Benchmarks = append(exe.Benchmarks, dur)
			exe.TotalTime += dur
			if mem := getPeakMemory(cmd.Process.Pid); mem > exe.PeakMemory {
				exe.PeakMemory = mem
			}
		case e := <-respErr:
			fmt.Println("\nError reading response:", e)
			cleanupServer(cmd, stdin)
			return
		case <-time.After(time.Duration(timeoutSec) * time.Second):
			fmt.Println("\nIteration timeout")
			cleanupServer(cmd, stdin)
			return
		}

		if exe.TotalTime > int64(timeoutSec*1000*1000) {
			fmt.Println()
			return
		}
	}
	fmt.Println()
}

func getPeakMemory(pid int) int64 {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return 0
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "VmHWM:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				if val, err := strconv.ParseInt(fields[1], 10, 64); err == nil {
					return val
				}
			}
		}
	}
	return 0
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
