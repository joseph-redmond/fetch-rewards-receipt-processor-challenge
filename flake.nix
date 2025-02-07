{
  description = "Go development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.git
          ];

          shellHook = ''
            echo "Welcome to the Go development environment!"

            echo "To build the application run the following command"
            echo "go build ./cmd/receipt-processor"

            echo "To run the application run the following command"
            echo "go run ./cmd/receipt-processor

            echo "To test the integration tests run the following command"
            echo "go test -v ./tests/integration"
          '';
        };
      });
}
