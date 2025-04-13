package main

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"testing"
)

// ================== Test Case Fiber  ==================

func TestUserRoute(t *testing.T) {
	app := setUp()

	// Define test cases
	tests := []struct {
		description  string
		requestBody  User
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  User{"jane.doe@example.com", "Jane Doe", 30},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Valid input",
			requestBody:  User{"wrong-email", "Jane Doe", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Valid input",
			requestBody:  User{"jane.doe@example.com", "1234", 30},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Valid input",
			requestBody:  User{"jane.doe@example.com", "Jane Doe", -2},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tc.expectStatus, resp.StatusCode)
		})
	}
}

// ================== Test Add and Test Factorial  ==================

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add positive numbers", 2, 3, 5},
		{"Add negative numbers", -1, -2, -3},
		{"Add zero", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			expectedResult := tc.expected
			if result != expectedResult {
				t.Errorf("Add(%d, %d) = %d is wrong, corret is %d", tc.a, tc.b, result, expectedResult)
			}
		})
	}

	//result := Add(2, 3)
	//expectResult := 5
	//if result != expectResult {
	//	t.Errorf("Add(2,3) = %d is wrong, correct is %d", result, expectResult)
	//}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		num      int
		expected int
	}{
		{"Case 2", 2, 2},
		{"Case 5", 5, 120},
		{"Case -1", -1, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Factorial(tc.num)
			expectedResult := tc.expected
			if result != expectedResult {
				t.Errorf("Factorial(%d) = %d is wrong, corret is %d", tc.num, result, expectedResult)
			}
		})
	}
}
