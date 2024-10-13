
# THA-Backend-Tezos

This is a Go-based service that collects new delegations from the Tezos blockchain using the tzKT API. The service continuously polls delegations and stores them in a database, which can be accessed through a public API endpoint.

## Features

- **Continuous Delegation Polling**: Fetches new Tezos delegations every 15 seconds and stores them.
- **Public API**:
  - `GET /xtz/delegations`: Retrieves stored delegations in JSON format, sorted by the most recent first.
- **Configurable Server Port**: Define the server's port via command-line arguments or environment variables.
  
## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/FayxChance/tha-backend-tezos
    cd tha-backend-tezos
    ```

2. Install the dependencies:

    Make sure you have Go installed. Then, run:

    ```bash
    go get -u github.com/gin-gonic/gin
    ```

## Running the Application

You can run the service with a custom port via command-line argument or environment variable.

### 1. Running on the default port (`8080`):

```bash
go run main.go
```

### 2. Running with a custom port using a command-line argument:

```bash
go run main.go -port=9090
```

### 3. Running with a custom port using the `PORT` environment variable:

```bash
export PORT=9090
go run main.go
```

The server will be available at `http://localhost:<PORT>/xtz/delegations`.

## Endpoints

### `GET /xtz/delegations`

Retrieves the stored delegations in JSON format.

Example response:
```json
{
  "data": [
    {
      "timestamp": "2023-10-10T12:34:56Z",
      "amount": "1000000",
      "delegator": "tz1...",
      "level": "123456"
    }
  ]
}
```

## Dependencies

- **Gin**: Web framework for Go.
- **SQLite**: Lightweight database to persist delegation data.
- **tzKT API**: Used for fetching new delegations from Tezos.

Install all dependencies by running:

```bash
go get -u github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
```

## License

This project is licensed under the MIT License.
