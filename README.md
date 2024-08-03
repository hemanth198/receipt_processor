# Receipt Processor

This is a Go application that processes receipts and calculates points based on a set of rules. The application exposes a REST API with two endpoints: one for processing receipts and another for retrieving points.

## API Endpoints

### Process Receipt

- **Endpoint**: `/receipts/process`
- **Method**: `POST`
- **Payload**: JSON receipt
- **Response**: JSON containing an ID for the receipt

### Get Points

- **Endpoint**: `/receipts/{id}/points`
- **Method**: `GET`
- **Response**: JSON object containing the number of points awarded

## Running the Application

### Using Docker

1. **Build the Docker image**:

    ```sh
    docker build -t receipt-processor -f docker .
    ```

2. **Run the Docker container**:

    ```sh
    docker run -p 8080:8080 receipt-processor
    ```

### Using `go run`

1. **Install dependencies**:

    ```sh
    go mod tidy
    ```

2. **Run the application**:

    ```sh
    go run main.go controller.go calculation.go models.go validation.go
    ```
