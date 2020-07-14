package classfile
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool{
	cpCount := int(reader.readUint16())
	constantPool := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i ++ {  //i = 1 开始
		constantPool[i] = readConstantInfo(reader,constantPool)
		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return constantPool
}

func readConstantInfo(reader *ClassReader, constantPool ConstantPool) ConstantInfo{
	tag := reader.readUint8()
	c := newConstantInfo(tag, constantPool)
	c.readInfo(reader)
	return c
}


func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if constantInfo := self[index]; constantInfo != nil{
		return constantInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string){
	info := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(info.nameIndex)
	_type := self.getUtf8(info.descriptorIndex)
	return name, _type
}
func (self ConstantPool) getClassName(index uint16) string{
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
func (self ConstantPool) getUtf8(index uint16) string{
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
