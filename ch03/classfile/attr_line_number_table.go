package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntity
}

type LineNumberTableEntity struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntity, length)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntity{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
