BIN_DIR = bin

.PHONY: build clean run compare

all: build

build: 
	@mkdir -p $(BIN_DIR)
	cargo build --release
	cp target/release/cgol $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)

run: build
	./$(BIN_DIR)/cgol --file ../input.txt --amount 10

compare:
	@diff -u ./output.txt ../c/output.txt