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

### Wishlist
- C++
- Zig
- Python
- Bash

## Results  
Board size: 50x50 & 1000 generations

| Language | Time (ms) |
|----------|-----------|
| C        | 109       |
| Rust     | 115       |
| Go       | 292       |
| C#       | 359       |
| OCaml    | 2047      |
| Perl     | 12134     |
