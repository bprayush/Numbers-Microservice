package main

import "errors"

// NumberService interface
type NumberService interface {
	Add(int, int) (int, error)
	Substract(int, int) (int, error)
	Multiply(int, int) (int, error)
	Divide(int, int) (int, error)
}

type numberService struct{}

func (numberService) Add(a int, b int) (int, error) {
	sum := a + b
	return sum, nil
}

func (numberService) Substract(a int, b int) (int, error) {
	diff := a - b
	return diff, nil
}

func (numberService) Multiply(a int, b int) (int, error) {
	product := a * b
	return product, nil
}

func (numberService) Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Invalid division by zero")
	}
	return a / b, nil
}
