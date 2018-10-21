package stack

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Pop the top operand stack value
Format
	pop
Forms
	pop = 87 (0x57)
Operand Stack
	..., value →
	...
Description
	Pop the top value from the operand stack.
	The pop instruction must not be used unless value is a value of a category 1 computational type (§2.11.1).

Operation
	Pop the top one or two operand stack values
Format
	pop2
Forms
	pop2 = 88 (0x58)
Operand Stack
	Form 1:
	..., value2, value1 →
	...
	where each of value1 and value2 is a value of a category 1 computational type (§2.11.1).
	Form 2:
	..., value →
	...
	where value is a value of a category 2 computational type (§2.11.1).
Description
	Pop the top one or two values from the operand stack.
*/

type POP struct {
	base.NoOperandsInstruction
}
type POP2 struct {
	base.NoOperandsInstruction
}

// int、float 等占一个操作数栈位置
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// double、long 占两个操作数栈位置
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
