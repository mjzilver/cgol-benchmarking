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
| Language     | Avg Time ms  | Peak Mem     | Iterations   |
|--------------|--------------|--------------|--------------|
| C++          | 58.23        | 3.38 MB      | 100          |
| C            | 76.78        | 1.52 MB      | 100          |
| Rust         | 79.71        | 1.89 MB      | 100          |
| Java         | 173.24       | 121.38 MB    | 100          |
| C#           | 234.51       | 35.53 MB     | 100          |
| F#           | 250.68       | 39.46 MB     | 100          |
| Node.js      | 405.63       | 68.27 MB     | 74           |
| Go           | 535.03       | 5.23 MB      | 57           |
| OCaml        | 664.99       | 5.52 MB      | 46           |
| Python       | 4140.69      | 14.66 MB     | 8            |
| Perl         | 10685.67     | 7.62 MB      | 3            |

## Time Complexity
O(T * N^2)

Where:
- `T` is the number of generations.
- `N` is the size.

## For example

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
