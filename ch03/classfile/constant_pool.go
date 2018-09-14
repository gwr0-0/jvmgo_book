package classfile

/*
	常量池是一张表。注意：
	1. 表头给出的常量池大小比实际大1.假设表头给出的值是n，那么常量池的实际大小是n-1
	2. 有效的常量池索引是1 - n-1。0是无效索引，表示不指向任何常量
	3. CONSTANT_Long_info 和 CONSTANT_Double_info各占两个位置。
	   故如果常量池中存在这两种常量，实际的常量数量比n-1还有少，且范围内某些数会变成无效索引
*/

type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
	}
}

// 按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {

}

// 从常量池查找字段或方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {

}

// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {

}

// 从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info
}
