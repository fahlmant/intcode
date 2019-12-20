package intcode

import (
	"fmt"
)

type Computer struct {
	PC, Offset, Output  int
	Instructions, Input []int
}

func (c *Computer) RunProgram() {

	for {
		opcode := c.Instructions[c.PC]
		arg1, arg2, arg3 := c.parseOpcode(opcode)
		switch opcode {
		case 1:
			c.Instructions[arg3] = c.Instructions[arg1] + c.Instructions[arg2]
			c.PC += 4
		case 2:
			c.Instructions[arg3] = c.Instructions[arg1] * c.Instructions[arg2]
			c.PC += 4
		case 3:
			if len(c.Input) >= 1{
				input := c.Input[0]
				c.Input = c.Input[1:]
				c.Instructions[arg1] = input
			}
			c.PC += 2
		case 4:
			c.Output = c.Instructions[arg1]
			c.PC +=2
		case 99:
			return
		default:
			fmt.Printf("Invalid Opcode: %d\n", opcode)
			return
		}
	}
}

func (c *Computer) parseOpcode(opcode int) (int, int, int) {

	return c.Instructions[c.PC+1], c.Instructions[c.PC+2], c.Instructions[c.PC+3]
}
