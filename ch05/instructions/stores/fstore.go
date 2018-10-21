package stores

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Store float into local variable
Format
	fstore
	index
Forms
	fstore = 56 (0x38)
Operand Stack
	..., value →
	...
Description
	The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
	The value on the top of the operand stack must be of type float.
	It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'. The value of the local variable at index is set to value'.
Notes
	The fstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.

Operation
Store float into local variable
Format
	fstore_<n>
Forms
	fstore_0 = 67 (0x43)
	fstore_1 = 68 (0x44)
	fstore_2 = 69 (0x45)
	fstore_3 = 70 (0x46)
Operand Stack
	..., value →
	...
Description
	The <n> must be an index into the local variable array of the current frame (§2.6). The value on the top of the operand stack must be of type float.
	It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'. The value of the local variable at <n> is set to value'.
Notes
	Each of the fstore_<n> instructions is the same as fstore with an index of <n>, except that the operand <n> is implicit.
*/

type FSTORE struct {
	base.Index8Instruction
}
type FSTORE_0 struct {
	base.NoOperandsInstruction
}
type FSTORE_1 struct {
	base.NoOperandsInstruction
}
type FSTORE_2 struct {
	base.NoOperandsInstruction
}
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.Index))
}
func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}
func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}
func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}
func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
