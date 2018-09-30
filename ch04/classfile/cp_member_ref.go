package classfile

/**
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

tag
    The tag item of a CONSTANT_Fieldref_info structure has the value CONSTANT_Fieldref (9).
    The tag item of a CONSTANT_Methodref_info structure has the value CONSTANT_Methodref (10).
    The tag item of a CONSTANT_InterfaceMethodref_info structure has the value CONSTANT_InterfaceMethodref (11).

class_index
    The value of the class_index item must be a valid index into the constant_pool table.
	The constant_pool entry at that index must be a CONSTANT_Class_info structure (§4.4.1) representing a class or interface type that has the field or method as a member.
    The class_index item of a CONSTANT_Methodref_info structure must be a class type, not an interface type.
    The class_index item of a CONSTANT_InterfaceMethodref_info structure must be an interface type.
    The class_index item of a CONSTANT_Fieldref_info structure may be either a class type or an interface type.

name_and_type_index
    The value of the name_and_type_index item must be a valid index into the constant_pool table.
	The constant_pool entry at that index must be a CONSTANT_NameAndType_info structure (§4.4.6).
	This constant_pool entry indicates the name and descriptor of the field or method.
    In a CONSTANT_Fieldref_info, the indicated descriptor must be a field descriptor (§4.3.2).
	Otherwise, the indicated descriptor must be a method descriptor (§4.3.3).
    If the name of the method of a CONSTANT_Methodref_info structure begins with a '<' ('\u003c'), then the name must be the special name <init>,
	representing an instance initialization method (§2.9.1). The return type of such a method must be void.
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// 字段符号引用
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

// 普通（非接口）方法符号引用
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

// 接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
