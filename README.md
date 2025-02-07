# Receipt Processor

The Receipt Processor is a Go-based application designed to process receipts efficiently and allow retrieving the calculated points for each receipt based on its id.
This project serves as a submission for the first-round of interviews for fetch-rewards, showcasing clean, structured, and maintainable Go code.
There is also an accompanying flake.nix file if that is desired for a dev environment. As that's what I use for my dev environments.

## Technologies Used
The dependencies used are as follows
* github.com/spf13/viper
* github.com/sirupsen/logrus
* github.com/cucumber/godog
* github.com/google/uuid
* github.com/gorilla/mux

## Nix Development Environment Instructions

To set up a nix develop environment, run the following command:
```
nix develop
```

## Build Instructions

To build the application, run the following command:
```
go build ./cmd/receipt-processor
```

## Run Instructions

To run the application, execute:
```
go run ./cmd/receipt-processor
```

## Testing Instructions

To run the integration tests, use:

```
go test -v ./tests/integration
```

## Project Structure
```
receipt-processor/
│── cmd/                # Entry point of the application
│── examples/           # Example json provided by fetch rewards
│── internal/           # Business logic and core processing
│── pkg/                # General configuration and setup
│── tests/features/     # Feature files for the godog tests tests
│── tests/integration/  # Integrated godog tests
│── go.mod              # Go module file
│── flake.nix           # Nix development environment file
│── .env                # Environment file for configurations
```

## Notes

This application is built using Go.

Tests follow a bdd approach and have been implemented by creating a feature file in gherkin and setting up the step based tests.

The project is structured for maintainability and scalability.

For any questions regarding this submission, feel free to reach out.

**Author**: Joseph Redmond Jr

**Submission For**: Fetch Rewards