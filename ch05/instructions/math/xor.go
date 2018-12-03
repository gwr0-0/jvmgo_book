package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
ixor
Operation
	Boolean XOR int
Format
	ixor
Forms
	ixor = 130 (0x82)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. They are popped from the operand stack.
	An int result is calculated by taking the bitwise exclusive OR of value1 and value2.
	The result is pushed onto the operand stack.

lxor
Operation
	Boolean XOR long
Format
	lxor
Forms
	lxor = 131 (0x83)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. They are popped from the operand stack.
	A long result is calculated by taking the bitwise exclusive OR of value1 and value2.
	The result is pushed onto the operand stack.

*/

type IXOR struct {
	base.NoOperandsInstruction
}
type LXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}
func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
