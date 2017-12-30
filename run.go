package main

import "fmt"

func (vm *machine) run() {
	for !vm.halted {
		vm.exec()
	}
}

func (vm *machine) exec() {
	op := vm.read()
	switch op {
	case 0: // halt
		vm.halted = true
	case 6: // jmp
		vm.pc = vm.mem[vm.pc]
	case 19: // out
		fmt.Print(string(rune(vm.read())))
	case 21: // noop
	default:
		panic(fmt.Errorf("unimplemented opcode: %d at %x", op, vm.pc))
	}
}
