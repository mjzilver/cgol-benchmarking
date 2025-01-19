.PHONY: build
build: 
	dotnet publish --configuration Release --self-contained true -r linux-x64 -p:PublishSingleFile=true -o ./bin/

.PHONY: run
run: build
	./bin/cgol --file ../input.txt --amount 10
