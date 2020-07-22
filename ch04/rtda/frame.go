package rtda

type Frame struct {
	lower *Frame
	//localVars		LocalVars		// 局部变量表指针
	//operandStack	*OperandStack   // 操作数栈？todo
}

//func NewFrame(maxLocals, maxStack uint) *Frame{
//	return &Frame{
//		localVars: maxLocals,
//		operandStack: newOperandStack(maxStack),
//	}
//}
