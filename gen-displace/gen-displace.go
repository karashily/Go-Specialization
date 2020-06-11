package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64
	fmt.Printf("accleration = ")
	fmt.Scan(&a)

	fmt.Printf("intitial velocity = ")
	fmt.Scan(&v0)

	fmt.Printf("intitial displacement = ")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("intitial displacement = ")
	fmt.Scan(&t)

	fmt.Println(fn(t))
}
