build:
	@ go build -o ./bin/plaudren ./cmd/.
run: build
	@ ./bin/plaudren
test:
	@ go test -v --race ./... 
