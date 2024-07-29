# Cache Manager API

Cache Manager API is a Go application that provides a simple in-memory cache with a GET, SET, and DELETE API. This API allows you to manage cache entries efficiently.

## Features

- Set cache with a key, value, and duration.
- Retrieve cache values by key.
- Delete cache entries by key.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (version 1.15 or later) installed on your machine.

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/cache-manager-api.git
    cd cache-manager-api
    ```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```

## Usage

1. **Run the application:**

    ```bash
    go run main.go
    ```

2. **Access the API:**

    The API will be available at `http://localhost:8080`.

## API Endpoints

- **GET /cache/{key}**: Retrieve a cache entry by key.

    **Example Request:**
    ```bash
    curl -X GET http://localhost:8080/cache/myKey
    ```

    **Example Response:**
    ```json
    {
      "key": "myKey",
      "value": "myValue"
    }
    ```

- **POST /cache**: Set a new cache entry with a key, value, and duration.

    **Example Request:**
    ```bash
    curl -X POST http://localhost:8080/cache \
    -H "Content-Type: application/json" \
    -d '{"key": "myKey", "value": "myValue", "duration": 60}'
    ```

    **Example Response:**
    ```json
    {
      "message": "Cache set successfully"
    }
    ```

- **DELETE /cache/{key}**: Remove a cache entry by key.

    **Example Request:**
    ```bash
    curl -X DELETE http://localhost:8080/cache/myKey
    ```

    **Example Response:**
    ```json
    {
      "message": "Cache deleted successfully"
    }
    ```

## Project Structure

- **main.go**: The main entry point of the application.
- **handlers.go**: Contains the HTTP handlers for the API endpoints.
- **cache.go**: Implements the in-memory cache logic.

