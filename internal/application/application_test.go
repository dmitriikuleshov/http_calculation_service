package application

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	calc "github.com/dmitriikuleshov/http_calculation_service/pkg/calculation"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		body         string
		expectedCode int
		expectedBody string
		mockCalc     func(string) (float64, error)
	}{
		{
			name:         "Valid Expression",
			method:       http.MethodPost,
			body:         `{"expression": "3 + 5"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"result":8}`,
			mockCalc: func(s string) (float64, error) {
				return 8, nil
			},
		},
		{
			name:         "Invalid Expression",
			method:       http.MethodPost,
			body:         `{"expression": "2 + +"}`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"error":"Invalid expression"}`,
			mockCalc: func(s string) (float64, error) {
				return 0, calc.ErrInvalidExpression
			},
		},
		{
			name:         "Division by Zero",
			method:       http.MethodPost,
			body:         `{"expression": "10 / 0"}`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"error":"Division by zero"}`,
			mockCalc: func(s string) (float64, error) {
				return 0, calc.ErrDivisionByZero
			},
		},
		{
			name:         "Mismatched Parentheses",
			method:       http.MethodPost,
			body:         `{"expression": "(2 + 3"}`,
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: `{"error":"Mismatched parentheses"}`,
			mockCalc: func(s string) (float64, error) {
				return 0, calc.ErrMismatchedParentheses
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			CalculateHandler(rec, req, tt.mockCalc)

			res := rec.Result()
			defer res.Body.Close()

			body, _ := io.ReadAll(res.Body)

			if res.StatusCode != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, res.StatusCode)
			}

			if strings.TrimSpace(string(body)) != tt.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tt.expectedBody, body)
			}
		})
	}
}
