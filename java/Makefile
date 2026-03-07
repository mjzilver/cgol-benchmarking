.PHONY: build
build:
	javac -O -g:none Main.java
	mkdir -p bin
	jar cfe bin/cgol.jar Main *.class

.PHONY: run
run:
	java Main --file ../input.txt --amount 10

.PHONY: clean
clean:
	rm -f *.class
	rm -rf bin
