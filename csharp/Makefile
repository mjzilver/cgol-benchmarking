CSC = mcs 
TARGET = cgol
SRC = Program.cs
BIN = bin

build: $(BIN)/$(TARGET)

$(BIN):
	mkdir -p $(BIN)

all: $(BIN) $(TARGET)

$(BIN)/$(TARGET): $(SRC) $(BIN)
	$(CSC) -out:$(BIN)/$(TARGET) $(SRC)

run: $(BIN)/$(TARGET)
	mono $(BIN)/$(TARGET) --file ../input.txt --amount 10

clean:
	rm -f $(BIN)/$(TARGET)

compare:
	@diff -u ./output.txt ../c/output.txt