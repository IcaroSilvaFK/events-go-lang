package main

import "fmt"

func main() {

	sl := []string{
		"Icaro", "Vieira", "da", "silva", "a caca", "Cacascas",
	}

	sl = append(sl[:2], sl[4:]...)

	fmt.Println(len(sl))
	fmt.Println(sl)

}
