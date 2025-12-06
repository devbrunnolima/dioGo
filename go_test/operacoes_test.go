package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	testValue := Sum(3, 2, 1)
	expected := 6

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	testValue := Sum(3, 2, 1)
	expected := 7

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	testValue := Multiply(10, 10)
	expected := 100

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldMultiplyIncorrect(t *testing.T) {
	testValue := Multiply(10, 10)
	expected := 2560

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	testValue := Subtract(10, 5)
	expected := 5

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldSubtractIncorrect(t *testing.T) {
	testValue := Subtract(10, 10)
	expected := 5

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	testValue, err := Divide(10)
	expected := 10

	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	testValue, err := Divide(10)
	expected := 5

	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if testValue != expected {
		t.Error("Expected:", expected, "Got:", testValue)
	}
}
