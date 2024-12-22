# Arithmetic Expression Evaluation Web Service

This project implements a lightweight and efficient web service for evaluating arithmetic expressions submitted via HTTP requests.

## 📂 Project Structure

-   **`cmd/`**: Contains the entry point of the application (`main.go`).
-   **`config/`**: Stores the server configuration files.
-   **`internal/application/`**: Implements the web server logic.
-   **`pkg/calculator/`**: Includes the core calculator logic for evaluating expressions.

## 🚀 Quick Start

### Prerequisites

Ensure [Go](https://go.dev/dl/) is installed on your system.

### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/dmitriikuleshov/http_calculation_service
    cd http_calculation_service
    ```

2. Set the environment variable CALC_CONFIG_PATH to the path of the configuration file. By default, the configuration file (config.json) is located in the configs folder in the project directory.

3. Run the server:
    ```bash
    go run ./cmd/main.go
    ```

The server will start and listen for HTTP requests on `http://localhost:8080`.

## 🛠️ Usage

Send a `POST` request to the server with the expression in JSON format.

### Example Request

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"52+52\"}" http://localhost:8080
```

**Response:**

```json
{
	"result": 104
}
```

**HTTP Status Code:** `200 OK`

---

## ❗ Possible Errors and Status Codes

Below is a comprehensive list of all possible errors, their descriptions, and associated HTTP status codes:

| **Error**                         | **Description**                                | **HTTP Status Code**        |
| --------------------------------- | ---------------------------------------------- | --------------------------- |
| `Invalid JSON format`             | The request body contains malformed JSON.      | `400 Bad Request`           |
| `Invalid request method`          | The HTTP method used is not `POST`.            | `405 Method Not Allowed`    |
| `Invalid expression`              | The arithmetic expression is not valid.        | `422 Unprocessable Entity`  |
| `Division by zero`                | Attempted division by zero in the expression.  | `422 Unprocessable Entity`  |
| `Mismatched parentheses`          | Parentheses in the expression are unbalanced.  | `422 Unprocessable Entity`  |
| `Invalid character in expression` | The expression contains unsupported symbols.   | `422 Unprocessable Entity`  |
| `Invalid number format`           | Numbers in the expression have invalid format. | `422 Unprocessable Entity`  |
| `Failed to read request body`     | Error reading the HTTP request body.           | `500 Internal Server Error` |
| `Internal server error`           | A generic error occurred on the server.        | `500 Internal Server Error` |

---

### Example Error Scenarios

1. **Malformed JSON Request**

    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"10+\"}" http://localhost:8080
    ```

    **Response:**

    ```json
    {
    	"error": "Invalid JSON format"
    }
    ```

    **Status Code:** `400`

2. **Division by Zero**

    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"10/0\"}" http://localhost:8080
    ```

    **Response:**

    ```json
    {
    	"error": "Division by zero"
    }
    ```

    **Status Code:** `422`

3. **Unbalanced Parentheses**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"(10+5\"}" http://localhost:8080
    ```
    **Response:**
    ```json
    {
    	"error": "Mismatched parentheses"
    }
    ```
    **Status Code:** `422`

---
