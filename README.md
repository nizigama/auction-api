# Ethereum auction contract API

#### Description
This service provides restful APIs for managing and querying data related to an auction
contract. Authentication is required to access the APIs.

#### Setup
To run the project locally, you need to have golang installed. Preferably the latest version(1.22),
the codebase in this project has been only tested on version 1.21 and 1.22 of Golang.
You may download and install golang locally to run the project or use docker, we'll explore both options

1. Install golang, the instructions can be found on the [golang website](https://go.dev/doc/install)
2. Or use docker, the instructions to install it can also be found on the [docker website](https://docs.docker.com/engine/install/)
3. Copy `.env.example` to `.env` and set the values of the empty variables to your respective ones
    ```shell
      # Linux & Unix
      cp .env.example .env
      # Windows, just use the file explorer😉
    ```
4. Create a `database.sqlite` that will be used to store users and invalidated tokens
    ```shell
      # Linux & Unix
      touch database.sqlite
      # Windows, just use the file explorer😉
    ```

#### Usage

##### Using docker
- Start server using docker
    ```shell
    docker compose up
    ```
- Run tests
    ```shell
    # Unix & Linux
    docker run --rm -v $(pwd):/opt/auction-api -w /opt/auction-api golang:1.21.8-alpine go test tests/*.go -v
    ```

##### Using installed go
- Download dependencies
    ```shell
    go get .
    ```
- Start server using installed golang
    ```shell
    go run *.go
    ```
-  Install documentation generator
    ```shell
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
- Generate documentation
    ```shell
    swag init
    ```
- Run tests
    ```shell
    go test tests/*.go -v
    ```