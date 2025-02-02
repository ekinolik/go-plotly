package validation

import (
	"fmt"
	"reflect"
)

// ValidationError represents an error during validation
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error for %s: %s", e.Field, e.Message)
}

// Validator interface for objects that can validate themselves
type Validator interface {
	Validate() error
}

// ValidateType checks if a value matches expected type
func ValidateType(field string, value interface{}, expectedType reflect.Kind) error {
	if value == nil {
		return nil
	}

	actualType := reflect.TypeOf(value).Kind()
	if actualType != expectedType {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("expected %v, got %v", expectedType, actualType),
		}
	}
	return nil
}
