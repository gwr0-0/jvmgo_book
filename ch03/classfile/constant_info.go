package classfile

/**
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

// tag常量值定义，详见 https://docs.oracle.com/javase/specs/jvms/se10/html/jvms-4.html#jvms-4.4-140
const (
	CONSTANT_Class              = 7  // 1.0.2
	CONSTANT_Fieldref           = 9  // 1.0.2
	CONSTANT_Methodref          = 10 // 1.0.2
	CONSTANT_InterfaceMethodref = 11 // 1.0.2
	CONSTANT_String             = 8  // 1.0.2
	CONSTANT_Integer            = 3  // 1.0.2
	CONSTANT_Float              = 4  // 1.0.2
	CONSTANT_Long               = 5  // 1.0.2
	CONSTANT_Double             = 6  // 1.0.2
	CONSTANT_NameAndType        = 12 // 1.0.2
	CONSTANT_Utf8               = 1  // 1.0.2
	CONSTANT_MethodHandle       = 15 // 7
	CONSTANT_MethodType         = 16 // 7
	CONSTANT_InvokeDynamic      = 18 // 7
	CONSTANT_Module             = 19 // 9
	CONSTANT_Package            = 20 // 9
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}

	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
