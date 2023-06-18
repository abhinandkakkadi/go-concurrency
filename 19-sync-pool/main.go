package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, debug string) {

	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString(" : ")
	b.WriteString(debug)
	w.Write(b.Bytes())
	bufPool.Put(b)
}
