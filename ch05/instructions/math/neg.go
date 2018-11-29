package math

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
ineg = 116 (0x74)
lneg = 117 (0x75)
fneg = 118 (0x76)
dneg = 119 (0x77)
*/

type INEG struct {
	base.NoOperandsInstruction
}
type LNEG struct {
	base.NoOperandsInstruction
}
type FNEG struct {
	base.NoOperandsInstruction
}
type DNEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := -val
	stack.PushInt(result)
}
func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := -val
	stack.PushLong(result)
}
func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := -val
	stack.PushFloat(result)
}
func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := -val
	stack.PushDouble(result)
}
