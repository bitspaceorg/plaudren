build:
	@ go build -o ./bin/plaudren ./cmd/plaudren
run: build
	@ ./bin/plaudren
test:
	@ go test -v --race ./... 
