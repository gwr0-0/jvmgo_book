package stack

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
Operation
	Duplicate the top operand stack value
Format
	dup
Forms
	dup = 89 (0x59)
Operand Stack
	..., value →
	..., value, value
Description
	Duplicate the top value on the operand stack and push the duplicated value onto the operand stack.
	The dup instruction must not be used unless value is a value of a category 1 computational type (§2.11.1).

Operation
	Duplicate the top operand stack value and insert two values down
Format
	dup_x1
Forms
	dup_x1 = 90 (0x5a)
Operand Stack
	..., value2, value1 →
	..., value1, value2, value1
Description
	Duplicate the top value on the operand stack and insert the duplicated value two values down in the operand stack.
	The dup_x1 instruction must not be used unless both value1 and value2 are values of a category 1 computational type (§2.11.1).

Operation
	Duplicate the top operand stack value and insert two or three values down
Format
	dup_x2
Forms
	dup_x2 = 91 (0x5b)
Operand Stack
	Form 1:
	..., value3, value2, value1 →
	..., value1, value3, value2, value1
	where value1, value2, and value3 are all values of a category 1 computational type (§2.11.1).
	Form 2:
	..., value2, value1 →
	..., value1, value2, value1
	where value1 is a value of a category 1 computational type and value2 is a value of a category 2 computational type (§2.11.1).
Description
	Duplicate the top value on the operand stack and insert the duplicated value two or three values down in the operand stack.

*/

type DUP struct {
	base.NoOperandsInstruction
}
type DUP_X1 struct {
	base.NoOperandsInstruction
}
type DUP_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot3)
	stack.PushSlot(slot1)
}

/**
Operation
	Duplicate the top operand stack value and insert two values down
Format
	dup_x1
Forms
	dup_x1 = 90 (0x5a)
Operand Stack
	..., value2, value1 →
	..., value1, value2, value1
Description
	Duplicate the top value on the operand stack and insert the duplicated value two values down in the operand stack.
	The dup_x1 instruction must not be used unless both value1 and value2 are values of a category 1 computational type (§2.11.1).

Operation
	Duplicate the top one or two operand stack values and insert two or three values down
Format
	dup2_x1
Forms
	dup2_x1 = 93 (0x5d)
Operand Stack
	Form 1:
	..., value3, value2, value1 →
	..., value2, value1, value3, value2, value1
	where value1, value2, and value3 are all values of a category 1 computational type (§2.11.1).
	Form 2:
	..., value2, value1 →
	..., value1, value2, value1
	where value1 is a value of a category 2 computational type and value2 is a value of a category 1 computational type (§2.11.1).
Description
	Duplicate the top one or two values on the operand stack and insert the duplicated values, in the original order, one value beneath the original value or values in the operand stack.

Operation
	Duplicate the top one or two operand stack values and insert two, three, or four values down
Format
	dup2_x2
Forms
	dup2_x2 = 94 (0x5e)
Operand Stack
	Form 1:
	..., value4, value3, value2, value1 →
	..., value2, value1, value4, value3, value2, value1
	where value1, value2, value3, and value4 are all values of a category 1 computational type (§2.11.1).
	Form 2:
	..., value3, value2, value1 →
	..., value1, value3, value2, value1
	where value1 is a value of a category 2 computational type and value2 and value3 are both values of a category 1 computational type (§2.11.1).
	Form 3:
	..., value3, value2, value1 →
	..., value2, value1, value3, value2, value1
	where value1 and value2 are both values of a category 1 computational type and value3 is a value of a category 2 computational type (§2.11.1).
	Form 4:
	..., value2, value1 →
	..., value1, value2, value1
	where value1 and value2 are both values of a category 2 computational type (§2.11.1).
Description
	Duplicate the top one or two values on the operand stack and insert the duplicated values, in the original order, into the operand stack.

*/
type DUP2 struct {
	base.NoOperandsInstruction
}
type DUP2_X1 struct {
	base.NoOperandsInstruction
}
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
