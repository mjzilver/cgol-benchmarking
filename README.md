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

### Wishlist
- C
- C++
- Rust
- Javascript (Node.js)
- Python
- C#
- OCaml
- Bash
- ASM (x86_64)

### General API
The general API for all the implementations is as follows:

--silent flag: If this flag is set, the program will not print the final state of the board. This is useful for debugging.

--amount flag: This flag is used to specify the number of iterations the game should run for.

--file flag: This flag is used to specify the file from which the initial state of the board should be read. 

You don't have to validate anything just assume the input it always correct.

### Rules
- No AI code generation. (Not fun)
- No dependencies.
- Cells must wrap around the edges. (Less chance of total cell death)
- Try to be fast.