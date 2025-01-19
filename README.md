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
- C# (Mono & .NET)
- FSharp 

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C            | 117.48       | 100          |
| Rust         | 126.77       | 100          |
| C# (mono)    | 270.27       | 100          |
| Go           | 315.10       | 96           |
| C# (.NET)    | 1437.91      | 21           |
| F#           | 1698.31      | 18           |
| OCaml        | 2385.61      | 13           |
| Perl         | 11434.61     | 3            |

Note: All the code is idiomatic for the language and not hightly optimized, that is not what is being tested.

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.