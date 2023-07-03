package errors

import "runtime"

type Source struct {
	Function string
	File     string
	Line     int
}

func CaptureSource(skip int) *Source {
	var pcs [1]uintptr

	runtime.Callers(
		skip+2,
		pcs[:],
	)

	fs := runtime.CallersFrames(pcs[:])
	f, _ := fs.Next()
	return &Source{
		Function: f.Function,
		File:     f.File,
		Line:     f.Line,
	}
}
