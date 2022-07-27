package configuration

type Configuration struct {
	url                     string
	bufferPoolConfiguration *BufferPoolConfiguration
}

type BufferPoolConfiguration struct {
	poolSize     int
	bufferLength int
}

func (c *Configuration) GetUrl() string {
	return c.url
}

func (c *Configuration) GetPoolSize() int {
	return c.bufferPoolConfiguration.poolSize
}

func (c *Configuration) GetBufferLength() int {
	return c.bufferPoolConfiguration.bufferLength
}
