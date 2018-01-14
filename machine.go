package main

import "fmt"

/*
== architecture ==
- three storage regions
  - memory with 15-bit address space storing 16-bit values
  - eight registers
  - an unbounded stack which holds individual 16-bit values
- all numbers are unsigned integers 0..32767 (15-bit)
- all math is modulo 32768; 32758 + 15 => 5
*/

type machine struct {
	mem   [32768]uint16
	reg   [8]uint16
	stack *stack

	// program counter
	pc uint16

	// flags
	halted bool
	debug  bool

	// funcs the vm should call when an address is reached
	callbacks map[uint16]func()
}

func newMachine() *machine {
	vm := new(machine)
	vm.stack = newStack()
	vm.callbacks = make(map[uint16]func())
	return vm
}

func (vm *machine) push(v uint16) {
	vm.stack.push(v)
}

func (vm *machine) pop() (uint16, bool) {
	return vm.stack.pop()
}

func (vm *machine) callback(addr uint16, f func()) {
	vm.callbacks[addr] = f
}

func (vm *machine) dumpMem() {
	for i := 0; i < 32768; i++ {
		v := vm.mem[i]
		if v <= 32767 {
			fmt.Printf("%04x %4x\n", i, v)
		} else if v <= 32775 {
			fmt.Printf("%04x r%d\n", i, v-32768)
		} else {
			panic(fmt.Errorf("invalid value in memory: %x at %04x", v, i))
		}
	}
}
