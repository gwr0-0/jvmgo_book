package stores

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Store int into local variable
Format
	istore
	index
Forms
	istore = 54 (0x36)
Operand Stack
	..., value →
	...
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6). The value on the top of the operand stack must be of type int. It is popped from the operand stack, and the value of the local variable at index is set to value.
Notes
	The istore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Store int into local variable
Format
	istore_<n>
Forms
	istore_0 = 59 (0x3b)
	istore_1 = 60 (0x3c)
	istore_2 = 61 (0x3d)
	istore_3 = 62 (0x3e)
Operand Stack
	..., value →
	...
Description
	The <n> must be an index into the local variable array of the current frame (§2.6). The value on the top of the operand stack must be of type int.
	It is popped from the operand stack, and the value of the local variable at <n> is set to value.
Notes
	Each of the istore_<n> instructions is the same as istore with an index of <n>, except that the operand <n> is implicit.
*/

type ISTORE struct {
	base.Index8Instruction
}
type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, uint(self.Index))
}
func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}
func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}
func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}
func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
