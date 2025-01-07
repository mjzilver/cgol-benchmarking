# cgol-benchmarking

This repository contains the code for benchmarking the performance of the Game of Life implementation in different languages.

How to benchmark
```bash
    ./benchmark.sh --file <filename> --depth <depth>
```
Generate new input files
```bash
    ./generate_board.sh --file <filename> --size <size> 
```

## Languages
- [ ] Go
- [ ] C
- [ ] C++
- [ ] Rust
- [ ] Javascript (Node.js)
- [ ] Python
- [ ] C#
- [ ] OCaml
- [ ] Bash
- [ ] ASM (x86_64)

### General API
The general API for all the implementations is as follows:

--silent flag: If this flag is set, the program will not print the final state of the board. This is useful for benchmarking purposes.

--amount flag: This flag is used to specify the number of iterations the game should run for.

--size flag: This flag is used to specify the size of the board. Needs to be a square board.

--file flag: This flag is used to specify the file from which the initial state of the board should be read. The file should contain the initial state of the board in the form of a 2D array of 0s and 1s. File extensions don't matter.


### Rules
- No AI.
- No input validation.