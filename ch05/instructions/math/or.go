package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
ior
Operation
	Boolean OR int
Format
	ior
Forms
	ior = 128 (0x80)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. They are popped from the operand stack.
	An int result is calculated by taking the bitwise inclusive OR of value1 and value2.
	The result is pushed onto the operand stack.

lor
Operation
	Boolean OR long
Format
	lor
Forms
	lor = 129 (0x81)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. They are popped from the operand stack.
	A long result is calculated by taking the bitwise inclusive OR of value1 and value2.
	The result is pushed onto the operand stack.

*/

type IOR struct {
	base.NoOperandsInstruction
}
type LOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}
func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
