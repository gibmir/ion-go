package server

import "fmt"

type MethodHandle interface {
	Call(args []interface{}) (interface{}, error)
}

type IncorrectArgumentErr struct {
	message string
}

func (err *IncorrectArgumentErr) Error() string {
	return err.message
}

type MethodHandle0[R any] struct {
	call func() (R, error)
}

func (m MethodHandle0[R]) Call(args []interface{}) (interface{}, error) {
	r, err := m.call()
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle1[T, R any] struct {
	call func(t T) (R, error)
}

func (m MethodHandle1[T, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, &IncorrectArgumentErr{fmt.Sprintf("args array has incorrect length [%d]", len(args))}
	}

	arg, ok := args[0].(T)
	if !ok {
		return nil, &IncorrectArgumentErr{fmt.Sprintf("arg has incorrect type [%T].",args[0])}
	}

	r, err := m.call(arg)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle2[T1, T2, R any] struct {
	call func(t1 T1, t2 T2) (R, error)
}

func (m MethodHandle2[T1, T2, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, &IncorrectArgumentErr{}
	}

	arg1, ok := args[0].(T1)
	if !ok {
		return nil, &IncorrectArgumentErr{}
	}

	arg2, ok := args[0].(T2)
	if !ok {
		return nil, &IncorrectArgumentErr{}
	}
	r, err := m.call(arg1, arg2)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type MethodHandle3[T1, T2, T3, R any] struct {
	call func(t1 T1, t2 T2, t3 T3) (R, error)
}

func (m MethodHandle3[T1, T2, T3, R]) Call(args []interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, &IncorrectArgumentErr{}
	}

	arg1, ok := args[0].(T1)
	if !ok {
		return nil, &IncorrectArgumentErr{}
	}

	arg2, ok := args[0].(T2)
	if !ok {
		return nil, &IncorrectArgumentErr{}
	}

	arg3, ok := args[0].(T3)
	if !ok {
		return nil, &IncorrectArgumentErr{}
	}

	r, err := m.call(arg1, arg2, arg3)
	if err != nil {
		return nil, err
	}
	return r, nil
}
