package stackweight

import (
	"fmt"
)

// Request received by the lambda function.
type Request struct {
	Length float32		`json:"length"`
	Width float32		`json:"width"`
	Height float32		`json:"height"`
	AvgWeight float32	`json:"avg_weight"`
}

// Response returned by the lambda function.
type Response struct {
	ApproxWeight float32	`json:"approx_weight"`
}

// ValidationError tells which field of the Request was incorrect.
// The resaon for the error is also supplied.
type ValidationError struct {
	Field string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Field %s is incorrect. Reason: %s", e.Field, e.Reason)
}

// CalcApproxOutputWeight returns the approximate weight of the stack.
func (r Request) CalcApproxOutputWeight() float32 {
	return r.Length * r.Width * r.Height * r.AvgWeight
}

// Validate the incoming request.
func (r Request) Validate() error {
	if r.Length <= 0 {
		return &ValidationError{"Length", "Length must be greater than zero"}
	}

	if r.Width <= 0 {
		return &ValidationError{"Width", "Width must be greater than zero"}
	}

	if r.Height <= 0 {
		return &ValidationError{"Height", "Height must be greater than zero"}
	}

	if r.AvgWeight <= 0 {
		return &ValidationError{"AvgWeight", "AvgWeight must be greater than zero"}
	}

	return nil
}
