package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthFieldPrependerWithEightByte_Success(t *testing.T) {
	a := assert.New(t)

	prepender, err := NewLengthFieldService(8)
	a.Nil(err)
	a.NotNil(prepender)
}

func TestLengthFieldPrependerWithFourByte_Success(t *testing.T) {
	a := assert.New(t)

	prepender, err := NewLengthFieldService(4)
	a.Nil(err)
	a.NotNil(prepender)
}

func TestLengthFieldPrependerWithThreeByte_Success(t *testing.T) {
	a := assert.New(t)

	prepender, err := NewLengthFieldService(3)
	a.Nil(err)
	a.NotNil(prepender)
}

func TestLengthFieldPrependerWithTwoByte_Success(t *testing.T) {
	a := assert.New(t)

	prepender, err := NewLengthFieldService(2)
	a.Nil(err)
	a.NotNil(prepender)
}

func TestLengthFieldPrependerWithByte_Success(t *testing.T) {
	a := assert.New(t)

	prepender, err := NewLengthFieldService(1)
	a.Nil(err)
	a.NotNil(prepender)
}
