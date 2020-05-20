package tracer

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/uber-go/atomic"
)

var (
	defaultTracer = newTracer()
)

// tracer represents a logger for saving a gRPC connection's lifecycle
//
// A tracer will record all the important steps that a connection
// dialing => re-connecting => closing
type tracer struct {
	enabled bool

	created sync.Map
	writer io.Writer
}

// newTracer create a new tracer
func newTracer() *tracer {
	return &tracer{
		writer: os.Stdout,
	}
}

func (t *tracer) setWriter(w io.Writer) {
	if w == nil {
		return
	}

	t.writer = w
}

// enable turn-on the tracer
func (t *tracer) enable() {
	t.enabled = true
}

// enable turn-off the tracer
func (t *tracer) disable() {
	t.enabled = false
}

// trace will write a tracing message to a writer
// default writer is os.Stdout
func (t *tracer) trace(msg ...interface{}) {
	if !t.enabled {
		return
	}

	t.write(msg...)
}

// newConnection record when dialing success
func (t *tracer) newConnection(addr string) {
	if !t.enabled {
		return
	}

	var c atomic.Uint64
	act, _ := t.created.LoadOrStore(addr, &c)
	n := act.(*atomic.Uint64).Inc()
	t.write(fmt.Sprintf("[grpc] new connection for [%s], and created count is %v", addr, n))
}

func (t *tracer) write(msg ...interface{}) {
	_, _ = fmt.Fprintln(t.writer, msg...)
}
