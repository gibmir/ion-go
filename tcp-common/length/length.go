package length

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

type LengthFieldService struct {
	// lengthFieldLength must be either 1, 2, 3, 4, or 8
	lengthFieldLength     int
	lengthFieldLengthPool sync.Pool
}

func NewLengthFieldService(lengthFieldLength int) (*LengthFieldService, error) {
	if lengthFieldLength != 1 && lengthFieldLength != 2 &&
		lengthFieldLength != 3 && lengthFieldLength != 4 &&
		lengthFieldLength != 8 {
		return nil, fmt.Errorf("lengthFieldLength must be either 1, 2, 3, 4, or 8: [%d]",
			lengthFieldLength)
	}
	return &LengthFieldService{lengthFieldLength: lengthFieldLength}, nil
}

func (s *LengthFieldService) ReadLengthField(buffer *bytes.Buffer) error {
	length := s.lengthFieldLengthPool.Get().([]byte)
	_, err := buffer.Read(length)
	if err != nil {
		return err
	}

	logrus.Debugf("Response length was read [%s]", length)

	return nil
}

func (s *LengthFieldService) WriteLengthField(length int, buffer *bytes.Buffer) error {
	if buffer == nil {
		return fmt.Errorf("provided buffer is nil")
	}
	if length < 0 {
		return fmt.Errorf("buffer length is less than zero")
	}
	switch s.lengthFieldLength {
	case 1:
		if length >= 256 {
			return fmt.Errorf("length does not fit into a byte: [%d] ", length)
		}
		return buffer.WriteByte(uint8(length))
	case 2:
		if length >= 65536 {
			return fmt.Errorf("length does not fit into a short: [%d] ", length)
		}
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 3:
		if length >= 16777216 {
			return fmt.Errorf("length does not fit into a medium: [%d] ", length)
		}

		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 4:
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	case 8:
		n, err := buffer.Write([]byte(strconv.Itoa(length)))
		if err != nil {
			return err
		}
		if n != 2 {
			return fmt.Errorf("writes more than 2 bytes for length")
		}
		return nil
	default:
		return fmt.Errorf("lengthFieldLength has incorrect value: [%d]", s.lengthFieldLength)
	}
}
