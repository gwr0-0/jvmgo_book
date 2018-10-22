package stack

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Swap the top two operand stack values
Format
	swap
Forms
	swap = 95 (0x5f)
Operand Stack
	..., value2, value1 →
	..., value1, value2
Description
	Swap the top two values on the operand stack.
	The swap instruction must not be used unless value1 and value2 are both values of a category 1 computational type (§2.11.1).
Notes
	The Java Virtual Machine does not provide an instruction implementing a swap on operands of category 2 computational types.

*/
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
