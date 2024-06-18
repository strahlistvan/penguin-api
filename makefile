APP = penguin-app
BIN_FOLDER = ./bin
CMD_SRC = cmd/penguin-project/main.go

build:
	go build -o $(BIN_FOLDER)/$(APP) $(CMD_SRC)

run:
	./bin/$(APP)

clean:
	rm -rf $(BIN_FOLDER)