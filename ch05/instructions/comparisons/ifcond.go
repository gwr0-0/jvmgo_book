package comparisons

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
if<cond>
Operation
	Branch if int comparison with zero succeeds
Format
	if<cond>
	branchbyte1
	branchbyte2
Forms
	ifeq = 153 (0x99)
	ifne = 154 (0x9a)
	iflt = 155 (0x9b)
	ifge = 156 (0x9c)
	ifgt = 157 (0x9d)
	ifle = 158 (0x9e)
Operand Stack
	..., value →
	...
Description
	The value must be of type int. It is popped from the operand stack and compared against zero.
	All comparisons are signed. The results of the comparisons are as follows:
		ifeq succeeds if and only if value = 0
		ifne succeeds if and only if value ≠ 0
		iflt succeeds if and only if value < 0
		ifle succeeds if and only if value ≤ 0
		ifgt succeeds if and only if value > 0
		ifge succeeds if and only if value ≥ 0
	If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset,
	where the offset is calculated to be (branchbyte1 << 8) | branchbyte2.
	Execution then proceeds at that offset from the address of the opcode of this if<cond> instruction.
	The target address must be that of an opcode of an instruction within the method that contains this if<cond> instruction.
	Otherwise, execution proceeds at the address of the instruction following this if<cond> instruction.

*/

type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
