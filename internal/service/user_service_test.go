package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	s := NewUserService()

	// fixed DOB
	dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	age := s.CalculateAge(dob)

	if age <= 0 {
		t.Fatalf("expected age > 0, got %d", age)
	}
}
