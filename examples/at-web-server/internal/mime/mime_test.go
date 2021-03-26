package mime

import (
	"mime"
	"testing"
)

func Test_init(t *testing.T) {
	for ext, typ := range mimeTypes {
		got := mime.TypeByExtension(ext)
		if got != typ {
			t.Errorf("got: %v, want %v", got, typ)
		}
	}
}

func Test_TypeByExtension(t *testing.T) {
	for ext, typ := range mimeTypes {
		got := TypeByExtension(ext)
		if got != typ {
			t.Errorf("got: %v, want %v", got, typ)
		}
	}

	const mimeOctetStream = "application/octet-stream"
	got := TypeByExtension(".invalidext")
	if got != mimeOctetStream {
		t.Errorf("got: %v, want: %v", got, mimeOctetStream)
	}
}
