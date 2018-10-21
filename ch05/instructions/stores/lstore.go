package stores

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Store long into local variable
Format
	lstore
	index
Forms
	lstore = 55 (0x37)
Operand Stack
	..., value →
	...
Description
	The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current frame (§2.6).
	The value on the top of the operand stack must be of type long. It is popped from the operand stack, and the local variables at index and index+1 are set to value.
Notes
	The lstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.


Operation
	Store long into local variable
Format
	lstore_<n>
Forms
	lstore_0 = 63 (0x3f)
	lstore_1 = 64 (0x40)
	lstore_2 = 65 (0x41)
	lstore_3 = 66 (0x42)
Operand Stack
	..., value →
	...
Description
	Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6).
	The value on the top of the operand stack must be of type long. It is popped from the operand stack, and the local variables at <n> and <n>+1 are set to value.
Notes
	Each of the lstore_<n> instructions is the same as lstore with an index of <n>, except that the operand <n> is implicit.
*/

type LSTORE struct {
	base.Index8Instruction
}
type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}
func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}
func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}
func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
