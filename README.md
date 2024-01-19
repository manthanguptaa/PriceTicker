## PriceTicker Microservice

**A Go microservice for fetching prices, designed with flexibility and extensibility in mind.**

## Key Features

* **Price Fetching:** Retrieves prices for specified ticker symbols (currently mocked for development).
* **Decorator Pattern:** Enhances core functionality with logging and metrics through decorators.
* **Mocked Metrics and Logging:** Placeholder implementations for easy integration with real systems later.
* **Client Library:** Provides a convenient way to interact with the microservice from other Go applications.
* **Dockerization:** Includes a Dockerfile for easy containerization.

## Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/manthanguptaa/PriceTicker.git
```

2. **(Choose one - Local or Docker):**

**Local:**

   a. Install dependencies:

   ```bash
   cd PriceTicker
   go mod download
   ```

   b. Run the microservice:

   ```bash
   go run main.go
   ```

**Docker:**

   a. Build the Docker image:

   ```bash
   cd PriceTicker
   docker build -t price-ticker .
   ```

   b. Run the Docker container:

   ```bash
   docker run -p 8080:8080 price-ticker
   ```

3. **Use the client library:**

```go
import "github.com/manthanguptaa/PriceTicker/client"

price, err := client.FetchPrice("ETH")
```
