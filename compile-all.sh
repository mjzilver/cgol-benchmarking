#!/bin/bash

cd go && make &
cd c && make &
cd cpp && make &
cd ocaml && make &
cd rust && make &
cd csharp && make &
cd fsharp && make &
cd java && make &
cd python && make &

wait

echo "compiled all solutions"
