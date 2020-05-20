package tracer

import "io"

func Enable() {
	defaultTracer.enable()
}

func Disable() {
	defaultTracer.disable()
}

func SetWriter(w io.Writer) {
	defaultTracer.setWriter(w)
}

func AddTrace(msg ...interface{})  {
	defaultTracer.trace(msg...)
}

func NewConnection(addr string) {
	defaultTracer.newConnection(addr)
}