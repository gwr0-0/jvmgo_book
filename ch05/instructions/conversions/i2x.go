package conversions

import "github.com/gwr0-0/jvmgo/ch05/instructions/base"

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
