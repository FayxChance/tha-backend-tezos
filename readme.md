# THA-BACKEND-TEZOS

This is a basic API built using the Gin web framework in Go. It provides an endpoint to retrieve delegation data and can be run on a custom port specified via command line or environment variables.

## Features

- **GET `/xtz/delegations`**: Returns a JSON response with a welcome message.
- Command-line flag support to specify the server's port (defaults to `8080`).
- Option to use the `PORT` environment variable to set the server's port.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/FayxChance/tha-backend-tezos
   cd gin-delegations-api
   ```

2. **Install dependencies:**

   Make sure you have Go installed. Then, install Gin using:

   ```bash
   go get -u github.com/gin-gonic/gin
   ```

## Running the Application

You can run the application with a custom port via command-line argument or environment variable.

### 1. Running with a default port (`8080`):

   ```bash
   go run main.go
   ```

### 2. Running with a custom port using the command-line argument:

   ```bash
   go run main.go -port=9090
   ```

### 3. Running with a custom port using the `PORT` environment variable:

   ```bash
   export PORT=9090
   go run main.go
   ```

The server will start on the specified port. You can access the API at `http://localhost:<PORT>/xtz/delegations`.

## Endpoints

### `GET /xtz/delegations`

Returns a simple JSON message.

**Response:**

```json
{
  "message": "Hello, world!"
}
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): Web framework used for building the API.

To install the dependencies, run:

```bash
go get -u github.com/gin-gonic/gin
```
