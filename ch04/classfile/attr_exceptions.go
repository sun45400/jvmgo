package classfile

type ExceptionsAttribute struct {
	exceptionsIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionsIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptinIndexTable() []uint16 {
	return self.exceptionsIndexTable
}
