#!/bin/bash

# defaults
input_file="input.txt"
depth=30

while [[ $# -gt 0 ]]; do
    case "$1" in
        --file)
            input_file="$2"
            shift 2
            ;;
        --depth)
            depth="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [--file INPUT_FILE] [--depth DEPTH]"
            echo "  --file INPUT_FILE: The file to read the initial board state from (default: input.txt)"
            echo "  --depth DEPTH: The number of generations to simulate (default: 30)"
            exit 1
            ;;
    esac
done

executables=("go/bin/cgol")

log_file="benchmark_results.log"
echo "Benchmark Results" > "$log_file"
echo "------------------" >> "$log_file"

for exe in "${executables[@]}"; do
    echo "Benchmarking $exe..." >> "$log_file"
    total_time=0

    start_time=$(date +%s%3N)
    ./"$exe" --silent --depth "$depth" "$input_file" > /dev/null
    end_time=$(date +%s%3N)

    elapsed_time=$((end_time - start_time))
    total_time=$((total_time + elapsed_time))

    echo "Executable $exe: ${elapsed_time} ms" >> "$log_file"
    echo "------------------" >> "$log_file"
done

echo "Benchmarking complete! Results saved to $log_file."
