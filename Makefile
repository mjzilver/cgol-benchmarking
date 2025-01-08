.PHONY: bench gen bench-gen

bench:
	./benchmark.sh

gen:
	./generate.sh

bench-gen: gen bench

fresh: compile gen bench

.PHONY: light
light:
	./benchmark.sh --iterations 1

.PHONY: compile
compile:
	cd go && make
	cd c && make

.PHONY: web
web:
	cd web && python3 server.py