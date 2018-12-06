package conversions

import (
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	"github.com/gwr0-0/jvmgo/ch05/rtda"
)

/**
i2l = 133 (0x85)
i2f = 134 (0x86)
i2d = 135 (0x87)
i2b = 145 (0x91)
i2c = 146 (0x92)
i2s = 147 (0x93)
*/

type I2L struct {
	base.NoOperandsInstruction
}
type I2F struct {
	base.NoOperandsInstruction
}
type I2D struct {
	base.NoOperandsInstruction
}
type I2B struct {
	base.NoOperandsInstruction
}
type I2C struct {
	base.NoOperandsInstruction
}
type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}
func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}
func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int8(i)
	stack.PushInt(int32(b))
}
func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := uint16(i)
	stack.PushInt(int32(c))
}
func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int16(i)
	stack.PushInt(int32(s))
}
