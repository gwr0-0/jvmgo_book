package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

}

// 读取单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

}

// 创建具体的属性实例
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

}
