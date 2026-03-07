# cgol-benchmarking

This repository contains the code for benchmarking the performance of the Game of Life implementation in different languages.

How to benchmark
```bash
    make bench
```
Generate new input files
```bash
    make gen
```

## Idea

I wanted a benchmark that actually tests something and not just a hello world or useless loops. 

This benchmark tests
- Handling process flags
- Reading a file
- Parsing text into a 2d array
- Computations on a 2d array
- Writing to file

## Current implemented Languages
- Go
- C (99)
- Perl
- Ocaml
- Rust
- C#
- FSharp 

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C            | 77.67        | 100          |
| Rust         | 80.90        | 100          |
| C#           | 304.17       | 99           |
| F#           | 455.20       | 66           |
| Go           | 515.32       | 59           |
| OCaml        | 699.29       | 43           |
| Perl         | 12957.62     | 3            |

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.