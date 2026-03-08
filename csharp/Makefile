.PHONY: build
dotnet:
	dotnet publish --configuration Release --self-contained true -r linux-x64 -p:PublishSingleFile=true -o ./bin

.PHONY: run
run:
	dotnet run -- --file ../input.txt --amount 10

