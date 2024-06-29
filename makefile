APP = penguin-app
BIN_FOLDER = ./bin
CMD_SRC = cmd/penguin-project/main.go

build:
	go build -o $(BIN_FOLDER)/$(APP) $(CMD_SRC)

run:
	./bin/$(APP) --addr 127.0.0.1 --port 8082

clean:
	rm -rf $(BIN_FOLDER)