.PHONY: bench gen bench-gen

bench:
	go run ./bench.go --chance 50 --amount 1000 --iterations 100 --silent

bench-gen: 
	go run ./bench.go --generate --size 50 --chance 50 --amount 1000 --iterations 100 --silent

fresh: compile light

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

.PHONY: fresh-comp 
fresh-comp: fresh compare

.PHONY: compare
compare:
	@diff ./c/output.txt ./go/output.txt
	@diff ./csharp/output.txt ./go/output.txt
	@diff ./ocaml/output.txt ./go/output.txt
	@diff ./rust/output.txt ./go/output.txt
	@diff ./perl/output.txt ./go/output.txt