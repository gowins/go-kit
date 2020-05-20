package tracing

import (
	"io"

	"github.com/gowins/go-kit/grpc/client/pool/internal/tracer"
)

func init()  {
	tracer.Enable()
}

func WithWriter(w io.Writer) {
	tracer.SetWriter(w)
}
