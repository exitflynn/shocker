package main

import "os"
import "fmt"

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("Weird command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}