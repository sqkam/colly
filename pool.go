package colly

import (
	"bytes"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64*1024))
	},
}
