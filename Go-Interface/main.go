package main

import "fmt"

// Creating an Interface
type bot interface {
	getGreetings() string
}

// English Bot Struct
type englishBot struct{}

// Spanish Bot Struct
type spanishBot struct{}

// Note📝: Interfaces are usefull, when we have ti use repetative logic for diff functions.
// 		It allows us to combine them together in creating one generic function and use it as many times.

func main() {
	fmt.Println("Hello from Go Interface DS!!")

	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)

	// printGreetingEn(eb)
	// printGreetingSp(sb)

}

// Creating the geGreetings() func, which takes an interface of type bot (here)
func printGreetings(b bot) {
	fmt.Println("Interface Called:", b.getGreetings())
}

func (eb englishBot) getGreetings() string {
	fmt.Println("Greetings from the English Bot ..", eb)
	return "Hey there!"
}

func (sb spanishBot) getGreetings() string {
	fmt.Println("Hello from Spanish Bot ..", sb)
	return "Que pasa?"
}

// func printGreetingEn(eb englishBot) {
// 	fmt.Printf("%+v", eb)
// 	fmt.Println(eb.getGreetings())
// }

// func printGreetingSp(sb spanishBot) {
// 	fmt.Printf("%+v", sb)
// 	fmt.Println(sb.getGreetings())
// }
