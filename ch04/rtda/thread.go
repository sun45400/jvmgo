package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		//todo 可以通过命令行参数来接收栈大小
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int      { return self.pc } // getter
func (self *Thread) SetPC(pc int) { self.pc = pc }   // setter
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
