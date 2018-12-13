package comparisons

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
lcmp
Operation
	Compare long
Format
	lcmp
Forms
	lcmp = 148 (0x94)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. They are both popped from the operand stack, and a signed integer comparison is performed.
	If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
	If value1 is equal to value2, the int value 0 is pushed onto the operand stack.
	If value1 is less than value2, the int value -1 is pushed onto the operand stack.

*/

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
