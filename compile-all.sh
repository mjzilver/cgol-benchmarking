#!/bin/bash

cd go && make &
cd c && make &
cd ocaml && make &
cd rust && make &
cd csharp && make &
cd fsharp && make &

wait

echo "compiled all solutions"
