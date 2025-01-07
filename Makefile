.PHONY: bench gen bench-gen compile

bench:
	./benchmark.sh

gen:
	./generate.sh

bench-gen: gen bench

compile:
	cd go && make