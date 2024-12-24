package models

// SuccessResponse for successful API responses.
type SuccessResponse struct {
    Message string      `json:"message"`     
    Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse for API error responses.
type ErrorResponse struct {
    Message string `json:"message"`
}

// ValidationError for a single validation error.
type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// ValidationErrorResponse for validation errors.
type ValidationErrorResponse struct {
    Errors []ValidationError `json:"errors"`
}
