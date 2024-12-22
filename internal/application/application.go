package application

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	calc "github.com/dmitriikuleshov/http_calculation_service/pkg/calculation"
)

// CalculationRequest represents the structure of incoming calculation requests
type CalculationRequest struct {
	Expression string `json:"expression" `
}

// CalculationResponse represents the structure of outgoing calculation responses
type CalculationResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// sendJSONResponse sends a JSON response with a specific status code
func sendJSONResponse(w http.ResponseWriter, response CalculationResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// CalculateHandler processes incoming HTTP requests to evaluate mathematical expressions
func CalculateHandler(w http.ResponseWriter, r *http.Request, calculate func(string) (float64, error)) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var request CalculationRequest
	if err = json.Unmarshal(body, &request); err != nil {
		log.Printf("JSON Unmarshal error: %v", err)
		response := CalculationResponse{Error: "Invalid JSON format"}
		sendJSONResponse(w, response, http.StatusBadRequest)
		return
	}

	result, err := calculate(request.Expression)
	if err != nil {
		log.Printf("Calculation error: %v", err)
		response := CalculationResponse{}
		status := http.StatusInternalServerError

		switch {
		case errors.Is(err, calc.ErrInvalidExpression):
			response.Error = "Invalid expression"
			status = http.StatusUnprocessableEntity
		case errors.Is(err, calc.ErrDivisionByZero):
			response.Error = "Division by zero"
			status = http.StatusUnprocessableEntity
		case errors.Is(err, calc.ErrMismatchedParentheses):
			response.Error = "Mismatched parentheses"
			status = http.StatusUnprocessableEntity
		case errors.Is(err, calc.ErrInvalidChacacter):
			response.Error = "Invalid character in expression"
			status = http.StatusUnprocessableEntity
		case errors.Is(err, calc.ErrInvalidNumber):
			response.Error = "Invalid number format"
			status = http.StatusUnprocessableEntity
		default:
			response.Error = "Unknown calculation error"
		}
		sendJSONResponse(w, response, status)
		return
	}

	sendJSONResponse(w, CalculationResponse{Result: result}, http.StatusOK)
}