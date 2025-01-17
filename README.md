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

## Current implemented Languages
- Go
- C 
- Perl
- Ocaml
- Rust
- C#

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C            | 105.15       | 100          |
| Rust         | 107.03       | 100          |
| C#           | 265.60       | 100          |
| Go           | 285.57       | 100          |
| OCaml        | 1934.83      | 16           |
| Perl         | 11015.35     | 3            |