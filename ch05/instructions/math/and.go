package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
iand
Operation
	Boolean AND int
Format
	iand
Forms
	iand = 126 (0x7e)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. They are popped from the operand stack.
	An int result is calculated by taking the bitwise AND (conjunction) of value1 and value2. The result is pushed onto the operand stack.

land
Operation
	Boolean AND long
Format
	land
Forms
	land = 127 (0x7f)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. They are popped from the operand stack.
	A long result is calculated by taking the bitwise AND of value1 and value2. The result is pushed onto the operand stack.

*/

type IAND struct {
	base.NoOperandsInstruction
}
type LAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}
func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
