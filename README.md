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
It ignores the startup time by setting up a server that receives the board path via stdin.

This benchmark tests
- Handling process flags
- Reading a file
- Parsing text into a 2d array
- Computations on a 2d array
- Writing to file

## Results  
Benchmark for a 50x50 board with 1000 generations with a 30 sec timeout
| Language     | Avg Time ms  | Iterations   |
|--------------|--------------|--------------|
| C++          | 56.92        | 100          |
| Rust         | 73.81        | 100          |
| C            | 74.36        | 100          |
| Java         | 164.89       | 100          |
| C#           | 209.98       | 100          |
| F#           | 236.01       | 100          |
| Node.js      | 389.49       | 78           |
| Go           | 524.00       | 58           |
| OCaml        | 624.70       | 49           |
| Python       | 3829.92      | 8            |
| Perl         | 9217.58      | 4            |

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.

## Benchmark configuration

Board size of 50:
```
50 × 50 = 2500 cells
```

1000 Generations
```
Total cell updates:
2500 × 1000 = 2,500,000
```

Each cell checks **8 neighbors**:

```
2,500,000 × 8 = 20,000,000 neighbor checks
```

So the benchmark roughly performs:

* **2.5 million cell updates**
* **20 million neighbor checks**
