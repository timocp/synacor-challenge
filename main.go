package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	hack := flag.Bool("hack", false, "hack the teleporter")
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	vm := newMachine()
	vm.load(file)
	if *hack {
		vm.callback(0x154b, func() {
			// this instruction is about to check register 7
			// trick the teleporter into going to the alternate location
			vm.reg[7] = 6
			// replace CALL 178b with JMP 167a.  This bypasses the teleporter
			// verification code
			vm.mem[0x1571] = 6
			vm.mem[0x1572] = 0x0157a
		})
	}
	vm.run()
}
