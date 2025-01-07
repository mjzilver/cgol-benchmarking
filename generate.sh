#!/bin/bash

# defaults
board_size=20
output_file="input.txt"
chance=50

while [[ $# -gt 0 ]]; do
    case "$1" in
        --size)
            board_size="$2"
            shift 2
            ;;
        --file)
            output_file="$2"
            shift 2
            ;;
        --chance)
            chance="$2"
            shift 2
            ;;
        *)
            echo "Unknown option: $1"
            echo "Usage: $0 [--size BOARD_SIZE] [--file OUTPUT_FILE] [--chance CHANCE]"
            echo "  --size BOARD_SIZE: The size of the board to generate (default: 20)"
            echo "  --file OUTPUT_FILE: The file to save the generated board to (default: input.txt)"
            echo "  --chance CHANCE: The percentage chance of a cell being alive must be between 0 and 100 (default: 50)"
            exit 1
            ;;
    esac
done

if (( chance < 0 || chance > 100 )); then
    echo "Chance must be between 0 and 100."
    exit 1
fi

function generate_random() {
    local scaled_random=$(( RANDOM * 100 / 32767 ))
    local chance="$1"
    if (( scaled_random < chance )); then
        echo "1"
    else
        echo "0"
    fi
}

true > "$output_file"

output=""

for ((i = 0; i < board_size; i++)); do
    row=""
    for ((j = 0; j < board_size; j++)); do
        row+="$(generate_random "$chance")" 
    done

    output+="${row}"
    
    if (( i < board_size - 1 )); then
        output+="\n"
    fi
done

echo -e "$output" > "$output_file"

echo "Random ${board_size}x${board_size} board generated and saved to $output_file."