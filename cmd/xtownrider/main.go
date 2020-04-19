package main

import (
	"fmt"

	"github.com/protheory8/xtownrider/internal/app/xtownrider"
)

func main() {
	fmt.Println("Xtownrider Copyright © 2020 The Xtownrider Contributors")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY.")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions. Read file 'COPYING' for details.")
	fmt.Println()

	xtownrider.MainGameLoop()
}
