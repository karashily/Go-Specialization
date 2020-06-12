package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animalsMap := make(map[string]Animal)
	for {
		fmt.Print("> ")
		var command string
		fmt.Scan(&command)
		if command == "newanimal" {
			var animalName, animalType string
			fmt.Scan(&animalName, &animalType)
			switch animalType {
			case "cow":
				animalsMap[animalName] = Cow{animalName}
			case "bird":
				animalsMap[animalName] = Bird{animalName}
			case "snake":
				animalsMap[animalName] = Snake{animalName}
			}
		} else if command == "query" {
			var animalName, reqType string
			fmt.Scan(&animalName, &reqType)
			switch reqType {
			case "eat":
				if val, ok := animalsMap[animalName]; ok {
					val.Eat()
				} else {
					fmt.Println("Bad Query.. Name Doesn't exist")
				}
			case "move":
				if val, ok := animalsMap[animalName]; ok {
					val.Move()
				} else {
					fmt.Println("Bad Query.. Name Doesn't exist")
				}
			case "speak":
				if val, ok := animalsMap[animalName]; ok {
					val.Speak()
				} else {
					fmt.Println("Bad Query.. Name Doesn't exist")
				}
			default:
				fmt.Println("Bad Query... Action not defined")
			}
		} else {
			fmt.Println("Wrong Command")
		}
	}
}
