.PHONY: bench gen bench-gen compile light

bench:
	./benchmark.sh

gen:
	./generate.sh

bench-gen: gen bench

fresh: compile gen bench

# Runs a lightweight version of the benchmark
light:
	./benchmark.sh --iterations 1

# This compiles all (where applicable) solutuons
compile:
	cd go && make