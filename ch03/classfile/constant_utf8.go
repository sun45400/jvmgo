package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readByte(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	//todo 仅适用于不含null和补充字符，完整版需要补全
	return string(bytes)
}
