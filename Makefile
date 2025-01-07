.PHONY: bench gen bench-gen compile light

bench:
	./benchmark.sh

gen:
	./generate.sh

bench-gen: gen bench

light:
	./benchmark.sh --iterations 1

compile:
	cd go && make