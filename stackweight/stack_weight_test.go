package stackweight

import "testing"

func TestCalcApproxOutputWeight(t *testing.T) {
	req := Request{5, 12, 4, 3}
	res := req.CalcApproxOutputWeight()
	if res != 720 {
		t.Error("5 x 12 x 4 x 3 was not 720")
	}
}

func TestValidateWithNegativeLength(t * testing.T) {
	req := Request{-5, 12, 4, 3}
	res := req.Validate()
	if res == nil {
		t.Fatal("Validation error expected")
	}
	message := res.Error()
	if message != "Field Length is incorrect. Reason: Length must be greater than zero" {
		t.Error("Incorrect error message")
	}
}

func TestValidateWithNegativeWidth(t * testing.T) {
	req := Request{5, -12, 4, 3}
	res := req.Validate()
	if res == nil {
		t.Fatal("Validation error expected")
	}
	message := res.Error()
	if message != "Field Width is incorrect. Reason: Width must be greater than zero" {
		t.Error("Incorrect error message")
	}
}

func TestValidateWithNegativeHeight(t * testing.T) {
	req := Request{5, 12, -4, 3}
	res := req.Validate()
	if res == nil {
		t.Fatal("Validation error expected")
	}
	message := res.Error()
	if message != "Field Height is incorrect. Reason: Height must be greater than zero" {
		t.Error("Incorrect error message")
	}
}

func TestValidateWithNegativeAvgWeight(t * testing.T) {
	req := Request{5, 12, 4, -3}
	res := req.Validate()
	if res == nil {
		t.Fatal("Validation error expected")
	}
	message := res.Error()
	if message != "Field AvgWeight is incorrect. Reason: AvgWeight must be greater than zero" {
		t.Error("Incorrect error message")
	}
}

func TestValidateWithValidRequest(t * testing.T) {
	req := Request{5, 12, 4, 3}
	res := req.Validate()
	if res != nil {
		t.Fatal("Unexpected validation error")
	}
}
