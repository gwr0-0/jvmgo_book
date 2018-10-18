package constants

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Do nothing
Format
	nop
Forms
	nop = 0 (0x0)
Operand Stack
	No change
Description
	Do nothing.
*/

// Do nothing
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
