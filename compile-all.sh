#!/bin/bash

cd go && make &
cd c && make &
cd ocaml && make &
cd rust && make &
cd csharp && make &
cd csharp && make dotnet &
cd fsharp && make &

wait

echo "compiled all solutions"
