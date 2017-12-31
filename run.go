package main

import "fmt"

func (vm *machine) run() {
	for !vm.halted {
		vm.exec()
	}
}

func (vm *machine) exec() {
	if vm.debug {
		vm.traceOp()
	}
	op := vm.read()
	switch op {
	case 0: // halt
		vm.halted = true
	case 1: // set a b
		a := vm.read()
		b := vm.readValue()
		vm.setValue(a, b)
	case 2: // push a
		vm.push(vm.readValue())
	case 3: // pop a
		v, ok := vm.pop()
		if !ok {
			panic(fmt.Errorf("pop from empty stack"))
		}
		vm.setValue(vm.read(), v)
	case 4: // eq a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		if b == c {
			vm.setValue(a, 1)
		} else {
			vm.setValue(a, 0)
		}
	case 5: // gt a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		if b > c {
			vm.setValue(a, 1)
		} else {
			vm.setValue(a, 0)
		}
	case 6: // jmp a
		vm.pc = vm.mem[vm.pc]
	case 7: // jt a b
		a := vm.readValue()
		b := vm.readValue()
		if a != 0 {
			vm.pc = b
		}
	case 8: // jf a b
		a := vm.readValue()
		b := vm.readValue()
		if a == 0 {
			vm.pc = b
		}
	case 9: // add a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		vm.setValue(a, (b+c)%32768)
	case 10: // mult a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		vm.setValue(a, (b*c)%32768)
	case 11: // mod 1 b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		vm.setValue(a, (b%c)%32768)
	case 12: // and a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		vm.setValue(a, b&c)
	case 13: // or a b c
		a := vm.read()
		b := vm.readValue()
		c := vm.readValue()
		vm.setValue(a, b|c)
	case 14: // not a b
		a := vm.read()
		b := vm.readValue()
		vm.setValue(a, (^b)%32768)
	case 15: // rmem a b
		a := vm.read()
		b := vm.readAt(vm.readValue())
		vm.setValue(a, b)
	case 17: // call a
		a := vm.readValue()
		vm.push(vm.pc)
		vm.pc = a
	case 19: // out a
		fmt.Print(string(rune(vm.read())))
	case 21: // noop
	default:
		panic(fmt.Errorf("unimplemented opcode: %d at %d", op, vm.pc-1))
	}
}

// information about each op
var ops = []struct {
	name string
	args uint16
}{
	{"HALT", 0},
	{"SET", 2},
	{"PUSH", 1},
	{"POP", 1},
	{"EQ", 3},
	{"GT", 3},
	{"JMP", 1},
	{"JT", 2},
	{"JF", 2},
	{"ADD", 3},
	{"MULT", 3},
	{"MOD", 3},
	{"AND", 3},
	{"OR", 3},
	{"NOT", 2},
	{"RMEM", 2},
	{"WMEM", 2},
	{"CALL", 1},
	{"RET", 0},
	{"OUT", 1},
	{"IN", 1},
	{"NOOP", 0},
}

func (vm *machine) traceOp() {
	op := vm.mem[vm.pc]
	if op == 19 {
		return // ignore OUT
	}
	fmt.Printf("%04x %-4s", vm.pc, ops[op].name)
	for i := uint16(0); i <= 3; i++ {
		if i < ops[op].args {
			m := vm.mem[vm.pc+i+1]
			if m <= 32767 {
				// memory address or literal
				fmt.Printf(" %4x", m)
			} else if m <= 32775 {
				// register reference
				fmt.Printf("   r%d", m-32768)
			} else {
				panic(fmt.Errorf("invalid memory: %d at %d", m, vm.pc+i+1))
			}
		} else {
			fmt.Printf("     ")
		}
	}
	fmt.Printf(" r=[")
	for i, r := range vm.reg {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%4x", r)
	}
	fmt.Printf("] s=%s\n", vm.stack)
}
