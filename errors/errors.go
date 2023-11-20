package errors

import "fmt"

// Define error codes and descriptions
const (
	ErrInvalidInput = iota + 1
	ErrDatabase     	
	// Add more error codes as needed
)

var errorMessages = map[int]string{
	ErrInvalidInput: "Invalid input",
	ErrDatabase:     "Database error",
	// Add more error messages as needed
}

// CustomError represents a custom error with code and description.
type CustomError struct {
	Code    int
	Message string
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NewError creates a new CustomError.
func NewError(code int) error {
	return &CustomError{
		Code:    code,
		Message: errorMessages[code],
	}
}
