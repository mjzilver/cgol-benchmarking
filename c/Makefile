CC = gcc
CFLAGS = -Wall -std=c99 -O3
BIN_DIR = bin
EXE_NAME = cgol

# Default target to build
all: build

build: main.c
	@mkdir -p $(BIN_DIR)
	$(CC) $(CFLAGS) -o $(BIN_DIR)/$(EXE_NAME) main.c

clean:
	rm -rf $(BIN_DIR)

run: build
	./$(BIN_DIR)/$(EXE_NAME) --file ../input.txt --amount 10
