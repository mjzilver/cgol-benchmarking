#!/bin/bash

# defaults
input_file="input.txt"
amount=1000
iterations=100

while [[ $# -gt 0 ]]; do
    case "$1" in
        --file)
            input_file="$2"
            shift 2
            ;;
        --amount)
            amount="$2"
            shift 2
            ;;
        --iterations)
            iterations="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [--file INPUT_FILE] [--amount AMOUNT] [--iterations ITERATIONS]"
            echo "  --file INPUT_FILE: The file to read the initial board state from (default: input.txt)"
            echo "  --amount AMOUNT: The number of generations to simulate (default: 1000)"
            echo "  --iterations ITERATIONS: The number of times to run each solution to average the runtime (default 100)"
            exit 1
            ;;
    esac
done

executables=("./go/bin/cgol" "./c/bin/cgol" "perl ./perl/cgol.pl")
names=("go" "c" "perl") 
timestamp=$(date +%s)

mkdir -p logs

# Function to run the benchmark for each executable
benchmark_executable() {
    exe=$1
    name=$2

    log_file_name="logs/benchmark_${name}_${timestamp}.log"
    total_time=0

    echo "Benchmarking $exe..." > "$log_file_name"
    echo "------------------" >> "$log_file_name"

    for ((i=1; i<=iterations; i++)); do
        echo "  Running iteration $i..." >> "$log_file_name"
        
        start_time=$(date +%s%3N)
        $exe --silent --amount "$amount" --file "$input_file" > /dev/null
        end_time=$(date +%s%3N)

        elapsed_time=$((end_time - start_time))
        total_time=$((total_time + elapsed_time))

        echo "    Iteration $i: ${elapsed_time} ms" >> "$log_file_name"
    done

    # Calculate average time
    avg_time=$((total_time / iterations))
    echo "Executable $exe ($name): Average time over $iterations iterations: ${avg_time} ms" >> "$log_file_name"
    echo "------------------" >> "$log_file_name"
}

# Run each executable in parallel
for i in "${!executables[@]}"; do
    benchmark_executable "${executables[$i]}" "${names[$i]}" &
done

# Wait for all background processes to finish
wait

echo "Benchmarking complete. Results are in the logs directory."
