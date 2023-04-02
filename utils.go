package main

import "math"

// функция возведения в степень. Первый аргумент - возводимое число, второй - желаемая степень. Она может быть любым целым числом.
func Pow(a, b int) float64 {
	var op int

	if b == 0 {
		return 1.0
	}
	
	if a < 0 {
		op = -1
	} else {
		op = 1
	}

	if b < 0 {
		for i := b; i < 0; i++ {
			op *= a
		}

		res := 1.0 / float64(op)
		return res
	}

	for i := 0; i < b; i++ {
		op *= a
	}

	return float64(op)
}

// a * (a - 1) * (a - 2) ... * (a - n) 
// возвращает это произведение, называемое факториалом. Максимальный возвращаемый факториал - 65
func Factorial(a uint) uint {
	res := uint(1)

	for i := a; i > 0; i-- {
		res *= i
	}

	return res
}

func LinearEquation(a, b, c float64) (*float64, *float64) {
	d :=  b*b - 4 * a * c

	if d < 0 {
		return nil, nil
	}

	if d == 0 {
		res := -b / (2 * a)
		return &res, nil
	}

	resA := -b + math.Sqrt(d) / (2 * a)
	resB := -b - math.Sqrt(d) / (2 * a)
	
	return &resA, &resB
}