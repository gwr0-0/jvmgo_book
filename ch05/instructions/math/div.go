package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
idiv = 108 (0x6c)
ldiv = 109 (0x6d)
fdiv = 110 (0x6e)
ddiv = 111 (0x6f)

Operand Stack
..., value1, value2 →
..., result
	value1 / value2
*/

type IDIV struct {
	base.NoOperandsInstruction
}
type LDIV struct {
	base.NoOperandsInstruction
}
type FDIV struct {
	base.NoOperandsInstruction
}
type DDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 / v2
	stack.PushInt(result)
}
func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 / v2
	stack.PushLong(result)
}
func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}
func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}
