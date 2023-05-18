package handle

import "fmt"

type MethodHandle interface {
	Call(args []interface{}) (interface{}, error)
}

type IncorrectArgumentErr struct {
	message string
}

func NewIncorrectArgsCountError(argsCount int) *IncorrectArgumentErr {
	return &IncorrectArgumentErr{message: fmt.Sprintf("args array has incorrect length [%d]", argsCount)}
}

func NewIncorrectArgType(arg any) *IncorrectArgumentErr {
	return &IncorrectArgumentErr{message: fmt.Sprintf("arg has incorrect type [%T]", arg)}
}
func (err *IncorrectArgumentErr) Error() string {
	return err.message
}

type MethodHandle0[R any] struct {
	CallFn func() (R, error)
}

func (m MethodHandle0[R]) Call(args []interface{}) (interface{}, error) {
	r, err := m.CallFn()
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle1[T, R any] struct {
	CallFn func(t T) (R, error)
}

func (m MethodHandle1[T, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, NewIncorrectArgsCountError(len(args))
	}

	arg, ok := args[0].(T)
	if !ok {
		return nil, NewIncorrectArgType(args[0])
	}

	r, err := m.CallFn(arg)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle2[T1, T2, R any] struct {
	CallFn func(t1 T1, t2 T2) (R, error)
}

func (m MethodHandle2[T1, T2, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, NewIncorrectArgsCountError(len(args))
	}

	arg1, ok := args[0].(T1)
	if !ok {
		return nil, NewIncorrectArgType(args[0])
	}

	arg2, ok := args[1].(T2)
	if !ok {
		return nil, NewIncorrectArgType(args[1])
	}
	r, err := m.CallFn(arg1, arg2)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle3[T1, T2, T3, R any] struct {
	CallFn func(t1 T1, t2 T2, t3 T3) (R, error)
}

func (m MethodHandle3[T1, T2, T3, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 3 {
		return nil, NewIncorrectArgsCountError(len(args))
	}

	arg1, ok := args[0].(T1)
	if !ok {
		return nil, NewIncorrectArgType(args[0])
	}

	arg2, ok := args[1].(T2)
	if !ok {
		return nil, NewIncorrectArgType(args[1])
	}

	arg3, ok := args[2].(T3)
	if !ok {
		return nil, NewIncorrectArgType(args[2])
	}

	r, err := m.CallFn(arg1, arg2, arg3)
	if err != nil {
		return nil, err
	}
	return r, nil
}
