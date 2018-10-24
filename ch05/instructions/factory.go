package instructions

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/constants"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/loads"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/math"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/stores"
)

var (
	nop         = &NOP{}
	aconst_null = &ACONST_NULL{}
	iconst_m1   = &ICONST_M1{}
	iconst_0    = &ICONST_0{}
	iconst_1    = &ICONST_1{}
	iconst_2    = &ICONST_2{}
	iconst_3    = &ICONST_3{}
	iconst_4    = &ICONST_4{}
	iconst_5    = &ICONST_5{}

	iload_1 = &ILOAD_1{}
	iload_2 = &ILOAD_2{}

	istore_1 = &ISTORE_1{}
	istore_2 = &ISTORE_2{}
	istore_3 = &ISTORE_3{}

	iadd = &IADD{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5

	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2

	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3

	case 0x60:
		return iadd
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
