package main

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
}

func newMachine() *machine {
	vm := new(machine)
	vm.stack = newStack()
	return vm
}

func (vm *machine) push(v uint16) {
	vm.stack.push(v)
}

func (vm *machine) pop() (uint16, bool) {
	return vm.stack.pop()
}
