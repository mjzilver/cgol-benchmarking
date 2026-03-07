#!/usr/bin/env python3

import sys
import argparse

OUTPUT_FILE = 'output.txt'

board = []
size = -1
silent = False

neighbors = [
    (-1, -1), (-1, 0), (-1, 1),
    (0, -1), (0, 1),
    (1, -1), (1, 0), (1, 1)
]

def parse_board(file_path):
    global board, size
    
    with open(file_path, 'r') as f:
        lines = f.read().strip().split('\n')
    
    size = len(lines)
    board = []
    
    for y in range(size):
        row = []
        for x in range(size):
            row.append(1 if lines[y][x] == '1' else 0)
        board.append(row)

def next_state():
    global board
    
    new_board = [[0] * size for _ in range(size)]
    
    for y in range(size):
        for x in range(size):
            count = 0
            
            for dy, dx in neighbors:
                ny = (y + dy + size) % size
                nx = (x + dx + size) % size
                count += board[ny][nx]
            
            if board[y][x] == 1:
                if count < 2 or count > 3:
                    new_board[y][x] = 0
                else:
                    new_board[y][x] = 1
            else:
                if count == 3:
                    new_board[y][x] = 1
                else:
                    new_board[y][x] = 0
    
    board = new_board

def print_board():
    with open(OUTPUT_FILE, 'w') as f:
        for y in range(size):
            for x in range(size):
                f.write(str(board[y][x]))
            f.write('\n')

def run_simulation(file_path, amount):
    parse_board(file_path)
    
    for _ in range(amount):
        next_state()
    
    if not silent:
        print_board()

def server_loop():
    print('READY', flush=True)
    
    try:
        for line in sys.stdin:
            line = line.strip()
            
            if not line:
                continue
            
            if line == 'SHUTDOWN':
                break
            
            if line.startswith('RUN '):
                parts = line[4:].split()
                if len(parts) >= 1:
                    file_path = parts[0]
                    req_amount = int(parts[1]) if len(parts) >= 2 else 1
                    
                    parse_board(file_path)
                    for _ in range(req_amount):
                        next_state()
                    
                    if not silent:
                        print_board()
                    
                    print('DONE', flush=True)
                else:
                    print('ERROR bad request', flush=True)
            else:
                print('ERROR unknown command', flush=True)
    except EOFError:
        pass

def main():
    global silent
    
    parser = argparse.ArgumentParser(description='Conway\'s Game of Life')
    parser.add_argument('--file', type=str, help='Input file')
    parser.add_argument('--amount', type=int, default=10, help='Number of generations')
    parser.add_argument('--silent', action='store_true', help='Silent mode')
    parser.add_argument('--server', action='store_true', help='Server mode')
    
    args = parser.parse_args()
    silent = args.silent
    
    if args.server:
        server_loop()
    else:
        if not args.file:
            print('Error: --file is required', file=sys.stderr)
            sys.exit(1)
        run_simulation(args.file, args.amount)

if __name__ == '__main__':
    main()
