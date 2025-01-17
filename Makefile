.PHONY: bench gen bench-gen

bench:
	go run ./bench.go --chance 50 --amount 1000 --iterations 100

bench-gen: 
	go run ./bench.go --generate --size 50 --chance 50 --amount 1000 --iterations 100

fresh: compile gen bench

.PHONY: light
light:
	go run ./bench.go --iterations 10 --amount 30

.PHONY: compile
compile:
	cd go && make
	cd c && make
	cd ocaml && make
	cd rust && make
	cd csharp && make

.PHONY: web
web:
	cd web && python3 server.py