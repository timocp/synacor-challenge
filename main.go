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
	vm := newMachine()
	vm.load(file)
	vm.breakpoint = 0x154b // JF which checks r7
	vm.run()
}
