package classfile

/**
field_info {
	u2 				access_flags;
	u2 				name_index;
	u2 				descriptor_index;
	u2				attributes_count;
	attribute_info  attributes[attributes_count];
}
*/
type MemberInfo struct {
	cp              ConstantPool // 常量池指针
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {

}

// 从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
