package instructions

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/constants"
)

var (
	nop = &NOP{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
