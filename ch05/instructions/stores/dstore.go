package stores

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Store double into local variable
Format
	dstore
	index
Forms
	dstore = 57 (0x39)
Operand Stack
	..., value →
	...
Description
	The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current frame (§2.6).
	The value on the top of the operand stack must be of type double.
	It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'. The local variables at index and index+1 are set to value'.
Notes
	The dstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
	Store double into local variable
Format
	dstore_<n>
Forms
	dstore_0 = 71 (0x47)
	dstore_1 = 72 (0x48)
	dstore_2 = 73 (0x49)
	dstore_3 = 74 (0x4a)
Operand Stack
	..., value →
	...
Description
	Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6).
	The value on the top of the operand stack must be of type double. It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'.
	The local variables at <n> and <n>+1 are set to value'.
Notes
	Each of the dstore_<n> instructions is the same as dstore with an index of <n>, except that the operand <n> is implicit.
*/

type DSTORE struct {
	base.Index8Instruction
}
type DSTORE_0 struct {
	base.NoOperandsInstruction
}
type DSTORE_1 struct {
	base.NoOperandsInstruction
}
type DSTORE_2 struct {
	base.NoOperandsInstruction
}
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(self.Index))
}
func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}
func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}
func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}
func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
