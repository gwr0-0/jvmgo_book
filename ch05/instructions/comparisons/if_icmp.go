package comparisons

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
if_icmp<cond>
Operation
	Branch if int comparison succeeds
Format
	if_icmp<cond>
	branchbyte1
	branchbyte2
Forms
	if_icmpeq = 159 (0x9f)
	if_icmpne = 160 (0xa0)
	if_icmplt = 161 (0xa1)
	if_icmpge = 162 (0xa2)
	if_icmpgt = 163 (0xa3)
	if_icmple = 164 (0xa4)
Operand Stack
	..., value1, value2 →
	...
Description
	Both value1 and value2 must be of type int. They are both popped from the operand stack and compared. All comparisons are signed.
		The results of the comparison are as follows:
		if_icmpeq succeeds if and only if value1 = value2
		if_icmpne succeeds if and only if value1 ≠ value2
		if_icmplt succeeds if and only if value1 < value2
		if_icmple succeeds if and only if value1 ≤ value2
		if_icmpgt succeeds if and only if value1 > value2
		if_icmpge succeeds if and only if value1 ≥ value2
	If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset, where the offset is calculated to be (branchbyte1 << 8) | branchbyte2.
	Execution then proceeds at that offset from the address of the opcode of this if_icmp<cond> instruction.
	The target address must be that of an opcode of an instruction within the method that contains this if_icmp<cond> instruction.
	Otherwise, execution proceeds at the address of the instruction following this if_icmp<cond> instruction.

*/

type IF_ICMPEQ struct {
	base.BranchInstruction
}
type IF_ICMPNE struct {
	base.BranchInstruction
}
type IF_ICMPLT struct {
	base.BranchInstruction
}
type IF_ICMPGE struct {
	base.BranchInstruction
}
type IF_ICMPGT struct {
	base.BranchInstruction
}
type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}
