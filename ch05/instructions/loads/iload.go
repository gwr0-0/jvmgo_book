package loads

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
从局部变量表获取变量让后推入操作数栈顶
*/

/**
Operation
	Load int from local variable
Format
	iload
	index
Forms
	iload = 21 (0x15)
Operand Stack
	... →
	..., value
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
	The local variable at index must contain an int. The value of the local variable at index is pushed onto the operand stack.
Notes
	The iload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Load int from local variable
Format
	iload_<n>
Forms
	iload_0 = 26 (0x1a)
	iload_1 = 27 (0x1b)
	iload_2 = 28 (0x1c)
	iload_3 = 29 (0x1d)
Operand Stack
	... →
	..., value
Description
	The <n> must be an index into the local variable array of the current frame (§2.6).
	The local variable at <n> must contain an int. The value of the local variable at <n> is pushed onto the operand stack.
Notes
	Each of the iload_<n> instructions is the same as iload with an index of <n>, except that the operand <n> is implicit.

*/

type ILOAD struct {
	base.Index8Instruction
}
type ILOAD_0 struct {
	base.NoOperandsInstruction
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, self.Index)
}
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
