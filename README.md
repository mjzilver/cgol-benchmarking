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
- C (99)
- Perl
- Ocaml (Functional)
- Rust
- C# (Mono)

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C            | 93.56        | 100          |
| Rust         | 96.17        | 100          |
| C#           | 257.70       | 100          |
| Go           | 278.32       | 100          |
| OCaml        | 2095.23      | 15           |
| Perl         | 9924.07      | 4            |

Note: All the code is idiomatic for the language and not hightly optimized, that is not what is being tested.

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.