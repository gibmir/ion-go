package pool

import (
	"bytes"
)

type BufferPool struct {
	poolSize     int
	bufferLength int
	pool         chan *bytes.Buffer
}

func NewBufferPool(poolSize, bufferLength int) *BufferPool {
	return &BufferPool{
		poolSize:     poolSize,
		bufferLength: bufferLength,
		pool:         make(chan *bytes.Buffer, poolSize),
	}
}

func (bufferPool *BufferPool) Get() *bytes.Buffer {
	select {
	case buffer := <-bufferPool.pool:
		return buffer
	default:
		return bytes.NewBuffer(make([]byte, bufferPool.bufferLength))
	}
}

func (bufferPool *BufferPool) Put(buffer *bytes.Buffer) {
	if buffer == nil {
		return
	}
	buffer.Reset()
	bufferPool.pool <- buffer
}
