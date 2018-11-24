package instructions

import (
	"fmt"
	"github.com/gwr0-0/jvmgo/ch05/instructions/base"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/constants"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/loads"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/math"
	. "github.com/gwr0-0/jvmgo/ch05/instructions/stack"
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
	lconst_0    = &LCONST_0{}
	lconst_1    = &LCONST_1{}
	fconst_0    = &FCONST_0{}
	fconst_1    = &FCONST_1{}
	fconst_2    = &FCONST_2{}
	dconst_0    = &DCONST_0{}
	dconst_1    = &DCONST_1{}
	bipush      = &BIPUSH{}
	sipush      = &SIPUSH{}

	iload   = &ILOAD{}
	lload   = &LLOAD{}
	fload   = &FLOAD{}
	dload   = &DLOAD{}
	aload   = &ALOAD{}
	iload_0 = &ILOAD_0{}
	iload_1 = &ILOAD_1{}
	iload_2 = &ILOAD_2{}
	iload_3 = &ILOAD_3{}
	lload_0 = &LLOAD_0{}
	lload_1 = &LLOAD_1{}
	lload_2 = &LLOAD_2{}
	lload_3 = &LLOAD_3{}
	fload_0 = &FLOAD_0{}
	fload_1 = &FLOAD_1{}
	fload_2 = &FLOAD_2{}
	fload_3 = &FLOAD_3{}
	dload_0 = &DLOAD_0{}
	dload_1 = &DLOAD_1{}
	dload_2 = &DLOAD_2{}
	dload_3 = &DLOAD_3{}
	aload_0 = &ALOAD_0{}
	aload_1 = &ALOAD_1{}
	aload_2 = &ALOAD_2{}
	aload_3 = &ALOAD_3{}

	istore   = &ISTORE{}
	lstore   = &LSTORE{}
	fstore   = &FSTORE{}
	dstore   = &DSTORE{}
	astore   = &ASTORE{}
	istore_0 = &ISTORE_0{}
	istore_1 = &ISTORE_1{}
	istore_2 = &ISTORE_2{}
	istore_3 = &ISTORE_3{}
	lstore_0 = &LSTORE_0{}
	lstore_1 = &LSTORE_1{}
	lstore_2 = &LSTORE_2{}
	lstore_3 = &LSTORE_3{}
	fstore_0 = &FSTORE_0{}
	fstore_1 = &FSTORE_1{}
	fstore_2 = &FSTORE_2{}
	fstore_3 = &FSTORE_3{}
	dstore_0 = &DSTORE_0{}
	dstore_1 = &DSTORE_1{}
	dstore_2 = &DSTORE_2{}
	dstore_3 = &DSTORE_3{}
	astore_0 = &ASTORE_0{}
	astore_1 = &ASTORE_1{}
	astore_2 = &ASTORE_2{}
	astore_3 = &ASTORE_3{}

	pop     = &POP{}
	pop2    = &POP2{}
	dup     = &DUP{}
	dup_x1  = &DUP_X1{}
	dup_x2  = &DUP_X2{}
	dup2    = &DUP2{}
	dup2_x1 = &DUP2_X1{}
	dup2_x2 = &DUP2_X2{}
	swap    = &SWAP{}

	iadd = &IADD{}
	ladd = &LADD{}
	fadd = &FADD{}
	dadd = &DADD{}
	isub = &ISUB{}
	lsub = &LSUB{}
	fsub = &FSUB{}
	dsub = &DSUB{}
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
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return bipush
	case 0x11:
		return sipush

		// TODO 0x12 ~ 0x14

	case 0x15:
		return iload
	case 0x16:
		return lload
	case 0x17:
		return fload
	case 0x18:
		return dload
	case 0x19:
		return aload
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3

		// TODO 0x2e ~ 0x35

	case 0x36:
		return istore
	case 0x37:
		return lstore
	case 0x38:
		return fstore
	case 0x39:
		return dstore
	case 0x3a:
		return astore
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3

		// TODO 0x4f ~ 0x56

	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap

	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub

	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
