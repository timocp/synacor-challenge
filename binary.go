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
	buf := make([]byte, 2)
	for i := 0; ; i++ {
		n, err := input.Read(buf)
		if err != nil {
			if err == io.EOF {
				if vm.debug {
					fmt.Printf("%d pairs read\n", i)
				}
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

// read returns the value pointed to by the program counter and increments
// the pc
func (vm *machine) read() uint16 {
	r := vm.mem[vm.pc]
	vm.pc++
	return r
}

// readValue reads the next address and interprets it as a literal or register
// read.
func (vm *machine) readValue() uint16 {
	v := vm.read()
	if v <= 32767 {
		return v
	}
	if v <= 32775 {
		return vm.reg[v-32768]
	}
	panic(fmt.Errorf("invalid memory: %d at %x", v, vm.pc-1))
}

// setValue sets a register or a memory address a to a specific value n
func (vm *machine) setValue(a, n uint16) {
	if n > 32775 {
		panic(fmt.Errorf("invalid value to store in memory: %d", n))
	}
	if a <= 32767 {
		if vm.debug {
			fmt.Printf(" setValue: memory location %x being set to %d\n", a, n)
		}
		vm.mem[a] = n
	} else if a <= 32775 {
		if vm.debug {
			fmt.Printf(" setValue: register %d being set to %d\n", a-32768, n)
		}
		vm.reg[a-32768] = n
	} else {
		panic(fmt.Errorf("invalid memory: %d", a))
	}
}
