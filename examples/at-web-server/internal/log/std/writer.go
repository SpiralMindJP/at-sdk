package std

import (
	"bytes"
	"fmt"
	syslog "log"
	"os"
	"sync"

	"github.com/SpiralMindJP/at-sdk/examples/at-web-server/internal/log"
)

type writer struct {
	d *syslog.Logger
	i *syslog.Logger
	w *syslog.Logger
	e *syslog.Logger

	bufLen  int
	bufPool sync.Pool
}

var DefaultBufferLength = 4096

func newWriter() *writer {
	w := &writer{
		d:      syslog.New(os.Stderr, "[D] ", syslog.LstdFlags|syslog.Lmsgprefix),
		i:      syslog.New(os.Stderr, "[I] ", syslog.LstdFlags|syslog.Lmsgprefix),
		w:      syslog.New(os.Stderr, "[W] ", syslog.LstdFlags|syslog.Lmsgprefix),
		e:      syslog.New(os.Stderr, "[E] ", syslog.LstdFlags|syslog.Lmsgprefix),
		bufLen: DefaultBufferLength,
	}

	w.bufPool.New = func() interface{} {
		return w.newBuffer()
	}

	return w
}

func (w *writer) newBuffer() []byte {
	return make([]byte, w.bufLen)
}

func (w *writer) logger(level log.Level) *syslog.Logger {
	switch level {
	case log.Debug:
		return w.d
	case log.Info:
		return w.i
	case log.Warn:
		return w.w
	case log.Error:
		return w.e
	default:
		panic("invalid log level")
	}
}

func (w *writer) logw(level log.Level, msg string, kvs ...interface{}) {
	buf := bytes.NewBuffer(w.bufPool.Get().([]byte))
	defer w.bufPool.Put(buf.Bytes())

	buf.WriteString(msg)

	for i := 0; i+1 < len(kvs); i += 2 {
		k := kvs[i]
		v := kvs[i+1]

		buf.WriteString(", ")

		key, ok := k.(string)
		if !ok {
			continue
		}
		buf.WriteString(key)
		buf.WriteString(": ")

		fmt.Fprint(buf, v)
	}

	w.logger(level).Print(buf.String())
}
