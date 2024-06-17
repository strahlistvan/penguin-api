build:
	go build -o ./bin/penguin-app cmd/penguin-project/main.go

run:
	./bin/penguin-app
clean:
	rm -rf ./bin