package main

import (
	"fmt"
	"io"
)

/*
== binary format ==
- each number is stored as a 16-bit little-endian pair (low byte, high byte)
- numbers 0..32767 mean a literal value
- numbers 32768..32775 instead mean registers 0..7
- numbers 32776..65535 are invalid
- programs are loaded into memory starting at address 0
- address 0 is the first 16-bit value, address 1 is the second 16-bit value, etc
*/

func (vm *machine) load(input io.Reader) error {
	fmt.Printf("Loading ... ")
	buf := make([]byte, 2)
	for i := 0; ; i++ {
		n, err := input.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%d pairs read\n", i)
				return nil
			}
			return err
		}
		if n != 2 {
			return fmt.Errorf("non-even input")
		}
		vm.mem[i] = uint16(buf[0]) + (uint16(buf[1]) << 8)
	}
}
