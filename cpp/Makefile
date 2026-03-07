CXX = g++
CXXFLAGS = -Wall -std=c++17 -O3
BIN_DIR = bin
EXE_NAME = cgol

all: build

build: main.cpp
	@mkdir -p $(BIN_DIR)
	$(CXX) $(CXXFLAGS) -o $(BIN_DIR)/$(EXE_NAME) main.cpp

clean:
	rm -rf $(BIN_DIR)

run: build
	./$(BIN_DIR)/$(EXE_NAME) --file ../input.txt --amount 10
