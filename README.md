# BitWars Player Go

## How it works

### Run it locally
Run the player directly:
```bash
go run main.go
```
or build an executable:
```bash
go build
./player-Go
```

### Functionality
This application template provides a HTTP server with a single POST endpoint (/) on port `3000`.
The bit-dealer sends a POST request containing the current game state as a JSON object.
Your task is to implement a function which returns a `PlayerAction` as a response.
You'll find a predefined function `Decide()` in this file: `/logic/strategy.go`.
In `/logic/strategy_test.go` you can add unit-tests to quickly debug your application locally.
Run all unit-tests with the following command (executed in the root path of this project):
```bash
go test ./...
```
