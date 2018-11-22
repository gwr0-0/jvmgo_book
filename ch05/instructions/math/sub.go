package math

/**
Operation
	Subtract int
Format
	isub
Forms
	isub = 100 (0x64)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type int. The values are popped from the operand stack. The int result is value1 - value2. The result is pushed onto the operand stack.
	For int subtraction, a-b produces the same result as a+(-b). For int values, subtraction from zero is the same as negation.
	The result is the 32 low-order bits of the true mathematical result in a sufficiently wide two's-complement format, represented as a value of type int.
	If overflow occurs, then the sign of the result may not be the same as the sign of the mathematical difference of the two values.
	Despite the fact that overflow may occur, execution of an isub instruction never throws a run-time exception.

Operation
	Subtract long
Format
	lsub
Forms
	lsub = 101 (0x65)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type long. The values are popped from the operand stack. The long result is value1 - value2. The result is pushed onto the operand stack.
	For long subtraction, a-b produces the same result as a+(-b). For long values, subtraction from zero is the same as negation.
	The result is the 64 low-order bits of the true mathematical result in a sufficiently wide two's-complement format, represented as a value of type long.
	If overflow occurs, then the sign of the result may not be the same as the sign of the mathematical difference of the two values.
	Despite the fact that overflow may occur, execution of an lsub instruction never throws a run-time exception.

Operation
	Subtract float
Format
	fsub
Forms
	fsub = 102 (0x66)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type float. The values are popped from the operand stack and undergo value set conversion (§2.8.3), resulting in value1' and value2'.
	The float result is value1' - value2'. The result is pushed onto the operand stack.
	For float subtraction, it is always the case that a-b produces the same result as a+(-b).
	However, for the fsub instruction, subtraction from zero is not the same as negation, because if x is +0.0, then 0.0-x equals +0.0, but -x equals -0.0.
	The Java Virtual Machine requires support of gradual underflow as defined by IEEE 754.
	Despite the fact that overflow, underflow, or loss of precision may occur, execution of an fsub instruction never throws a run-time exception.

Operation
	Subtract double
Format
	dsub
Forms
	dsub = 103 (0x67)
Operand Stack
	..., value1, value2 →
	..., result
Description
	Both value1 and value2 must be of type double. The values are popped from the operand stack and undergo value set conversion (§2.8.3), resulting in value1' and value2'.
	The double result is value1' - value2'. The result is pushed onto the operand stack.
	For double subtraction, it is always the case that a-b produces the same result as a+(-b).
	However, for the dsub instruction, subtraction from zero is not the same as negation, because if x is +0.0, then 0.0-x equals +0.0, but -x equals -0.0.
	The Java Virtual Machine requires support of gradual underflow as defined by IEEE 754.
	Despite the fact that overflow, underflow, or loss of precision may occur, execution of a dsub instruction never throws a run-time exception.

*/
