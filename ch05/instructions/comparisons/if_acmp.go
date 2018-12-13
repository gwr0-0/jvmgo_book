package comparisons

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
if_acmp<cond>
Operation
	Branch if reference comparison succeeds
Format
	if_acmp<cond>
	branchbyte1
	branchbyte2
Forms
	if_acmpeq = 165 (0xa5)
	if_acmpne = 166 (0xa6)
Operand Stack
	..., value1, value2 →
	...
Description
	Both value1 and value2 must be of type reference. They are both popped from the operand stack and compared. The results of the comparison are as follows:
		if_acmpeq succeeds if and only if value1 = value2
		if_acmpne succeeds if and only if value1 ≠ value2
	If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset, where the offset is calculated to be (branchbyte1 << 8) | branchbyte2.
	Execution then proceeds at that offset from the address of the opcode of this if_acmp<cond> instruction.
	The target address must be that of an opcode of an instruction within the method that contains this if_acmp<cond> instruction.
	Otherwise, if the comparison fails, execution proceeds at the address of the instruction following this if_acmp<cond> instruction.

*/

type IF_ACMPEQ struct {
	base.BranchInstruction
}
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}
func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopRef()
	val1 := stack.PopRef()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}
