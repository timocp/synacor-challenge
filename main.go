package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	vm := new(machine)
	vm.load(file)
	vm.debug = true
	vm.run()
}
