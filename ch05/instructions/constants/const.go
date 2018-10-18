package constants

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
这一系列指令把隐含在操作码中的常量值推入操作数栈顶
*/

/**
Operation
	Push null
Format
	aconst_null
Forms
	aconst_null = 1 (0x1)
Operand Stack
	... →
	..., null
Description
	Push the null object reference onto the operand stack.
Notes
	The Java Virtual Machine does not mandate a concrete value for null.
*/
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

/**
Operation
	Push double
Format
	dconst_<d>
Forms
	dconst_0 = 14 (0xe)
	dconst_1 = 15 (0xf)
Operand Stack
	... →
	..., <d>
Description
	Push the double constant <d> (0.0 or 1.0) onto the operand stack.
*/
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

/**
Operation
	Push float
Format
	fconst_<f>
Forms
	fconst_0 = 11 (0xb)
	fconst_1 = 12 (0xc)
	fconst_2 = 13 (0xd)
Operand Stack
	... →
	..., <f>
Description
	Push the float constant <f> (0.0, 1.0, or 2.0) onto the operand stack.
*/
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

/**
Operation
	Push int constant
Format
	iconst_<i>
Forms
	iconst_m1 = 2 (0x2)
	iconst_0 = 3 (0x3)
	iconst_1 = 4 (0x4)
	iconst_2 = 5 (0x5)
	iconst_3 = 6 (0x6)
	iconst_4 = 7 (0x7)
	iconst_5 = 8 (0x8)
Operand Stack
	... →
	..., <i>
Description
	Push the int constant <i> (-1, 0, 1, 2, 3, 4 or 5) onto the operand stack.
Notes
	Each of this family of instructions is equivalent to bipush <i> for the respective value of <i>, except that the operand <i> is implicit.
*/
type ICONST_M1 struct {
	base.NoOperandsInstruction
}
type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}
func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}
func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}
func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}
func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

/**
Operation
	Push long constant
Format
	lconst_<l>
Forms
	lconst_0 = 9 (0x9)
	lconst_1 = 10 (0xa)
Operand Stack
	... →
	..., <l>
Description
	Push the long constant <l> (0 or 1) onto the operand stack.
*/
type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}
func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
