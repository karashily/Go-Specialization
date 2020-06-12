package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func Do(req []string) {

}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var animal, act string

	for {
		fmt.Print("> ")
		fmt.Scan(&animal)
		fmt.Scan(&act)
		if animal == "cow" {
			if act == "eat" {
				cow.Eat()
			} else if act == "move" {
				cow.Move()
			} else if act == "speak" {
				cow.Speak()
			} else {
				println("Error! act not defined!")
			}
		} else if animal == "bird" {
			if act == "eat" {
				bird.Eat()
			} else if act == "move" {
				bird.Move()
			} else if act == "speak" {
				bird.Speak()
			} else {
				println("Error! act not defined!")
			}
		} else if animal == "snake" {
			if act == "eat" {
				snake.Eat()
			} else if act == "move" {
				snake.Move()
			} else if act == "speak" {
				snake.Speak()
			} else {
				println("Error! act not defined!")
			}
		} else {
			println("Error! animal not defined")
		}
	}
}
