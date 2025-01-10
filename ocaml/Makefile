OCAMLC = ocamlc
TARGET = cgol
SRC = cgol.ml
BIN = bin

build: $(BIN)/$(TARGET)

$(BIN):
	mkdir -p $(BIN)

all: $(BIN) $(TARGET)

$(BIN)/$(TARGET): $(SRC) $(BIN)
	$(OCAMLC) -o $(BIN)/$(TARGET) $(SRC)

run: $(BIN)/$(TARGET)
	./$(BIN)/$(TARGET) --file ../input.txt --amount 10

clean:
	rm -f $(BIN)/$(TARGET)
	rm -f *.cmi
	rm -f *.cmo
