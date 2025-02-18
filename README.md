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
- Ocaml
- Rust
- C#
- FSharp 
- Zig

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C            | 63.95        | 100          |
| Rust         | 73.25        | 100          |
| Zig          | 160.41       | 100          |
| C#           | 227.94       | 100          |
| F#           | 344.09       | 88           |
| Go           | 464.56       | 65           |
| OCaml        | 1820.10      | 17           |
| Perl         | 10335.59     | 3            |

Note: All the code is idiomatic for the language and not hightly optimized, that is not what is being tested.

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.