package describer

import (
	"encoding/json"
	goerrors "errors"
	"fmt"

	"github.com/gibmir/ion-go/api/pkg/dto"
	"github.com/gibmir/ion-go/api/pkg/errors"
)

var (
	emptyArr        = make([]interface{}, 0)
	unmarshableJson = `{"jsonrpc":"2.0","error":{"code": %d,"message": "%s"},"id":"%s"}`
)

type ProcedureDescription struct {
	ProcedureName string
	ArgNames      []string
}

type Marshaller interface {
	Unmarshal(bytes []byte) ([]interface{}, error)
	Marshal(id string, result interface{}, err error) []byte
}

// Describer0 no args procedure describer
type Describer0[R any] struct {
	*Describer
}

func (d *Describer0[R]) Marshal(id string, result any, err error) []byte {
	return marshalResult[R](id, result, err)
}

func (d *Describer0[R]) Unmarshal(bytes []byte) ([]interface{}, error) {
	return emptyArr, nil
}

// Describer1 one arg procedure describer
type Describer1[T, R any] struct {
	*Describer
}

func (d *Describer1[T1, R]) Marshal(id string, result any, err error) []byte {
	return marshalResult[R](id, result, err)
}

func (d *Describer1[T, R]) Unmarshal(bytes []byte) ([]interface{}, error) {
	t, err := d.unmarshal(bytes)
	if err != nil {
		return emptyArr, err
	}

	return []interface{}{t}, nil
}

func (d *Describer1[T, R]) unmarshal(bytes []byte) (T, error) {
	var (
		t T
	)
	if len(bytes) == 0 {
		return t, nil
	}

	firstRune := rune(bytes[0])
	if firstRune == '{' {
		//named request
		var argsMap map[string]*json.RawMessage
		err := json.Unmarshal(bytes, &argsMap)
		if err != nil {
			return t, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsMap[d.Description.ArgNames[0]], t)
		if err != nil {
			return t, errors.NewInvalidRequestError(err.Error())
		}

		return t, nil
	} else if firstRune == '[' {
		//positional request
		var argsArray []*json.RawMessage

		err := json.Unmarshal(bytes, &argsArray)
		if err != nil {
			return t, errors.NewInvalidRequestError(err.Error())
		}

		if len(argsArray) != 1 {
			return t, errors.NewInvalidParamsError(fmt.Sprintf("params array has incorrect size [%d]", len(argsArray)))
		}
		err = json.Unmarshal(*argsArray[0], &t)
		if err != nil {
			return t, errors.NewInvalidRequestError(err.Error())
		}

		return t, nil
	} else {
		return t, errors.NewInvalidRequestError("")
	}
}

// Describer2 two arg procedure describer
type Describer2[T1, T2, R any] struct {
	*Describer
}

func (d *Describer2[T1, T2, R]) Marshal(id string, result any, err error) []byte {
	return marshalResult[R](id, result, err)
}

func (d *Describer2[T1, T2, R]) Unmarshal(bytes []byte) ([]interface{}, error) {
	t1, t2, err := d.unmarshal(bytes)
	if err != nil {
		return emptyArr, err
	}

	return []interface{}{t1, t2}, nil
}

func (d *Describer2[T1, T2, R]) unmarshal(bytes []byte) (T1, T2, error) {
	var (
		t1 T1
		t2 T2
	)
	if len(bytes) == 0 {
		return t1, t2, nil
	}

	firstRune := rune(bytes[0])
	if firstRune == '{' {
		//named request
		var argsMap map[string]*json.RawMessage
		err := json.Unmarshal(bytes, &argsMap)
		if err != nil {
			return t1, t2, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsMap[d.Description.ArgNames[0]], t1)
		if err != nil {
			return t1, t2, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsMap[d.Description.ArgNames[1]], t2)
		if err != nil {
			return t1, t2, errors.NewInvalidRequestError(err.Error())
		}

		return t1, t2, nil
	} else if firstRune == '[' {
		//positional request

		return t1, t2, nil
	} else {
		return t1, t2, errors.NewInvalidRequestError("")
	}
}

// Describer3 three arg procedure describer
type Describer3[T1, T2, T3, R any] struct {
	*Describer
}

func (d *Describer3[T1, T2, T3, R]) Marshal(id string, result any, err error) []byte {
	return marshalResult[R](id, result, err)
}

func marshalResult[R any](id string, result any, err error) []byte {
	if err != nil {
		var jsonRpcError *errors.JsonRpcError
		isJsonRpcError := goerrors.As(err, &jsonRpcError)
		if isJsonRpcError {
			return marshal(id, dto.Response[R]{
				Protocol: dto.DefaultJsonRpcProtocolVersion,
				Id:       id,
				Error:    jsonRpcError,
			})
		}
		return marshal(id, dto.Response[R]{
			Protocol: dto.DefaultJsonRpcProtocolVersion,
			Id:       id,
			Error:    errors.NewInternalError(err.Error()),
		})
	}

	r, ok := result.(R)
	if !ok {
		return marshal(id, dto.Response[R]{
			Protocol: dto.DefaultJsonRpcProtocolVersion,
			Id:       id,
			Error:    errors.NewInternalError(fmt.Sprintf("result has incorrect type [%T]", result)),
		})
	}

	return marshal(id, dto.Response[R]{Id: id, Result: r})
}

func MarshalError(err error) []byte {
	var jsonRpcError *errors.JsonRpcError
	isJsonRpcError := goerrors.As(err, &jsonRpcError)
	if isJsonRpcError {
		return []byte(fmt.Sprintf(`{"jsonrpc":"2.0","error":{"code": %d,"message": "%s"},"id":null}`,
			jsonRpcError.Code, jsonRpcError.Message))
	}
	return []byte(fmt.Sprintf(`{"jsonrpc":"2.0","error":{"code": %d,"message": "%s"},"id":null}`,
		errors.InvalidRequestErrorCode, err.Error()))
}

func marshal(id string, v any) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		return []byte(fmt.Sprintf(`{"jsonrpc":"2.0","error":{"code": %d,"message": "%s"},"id":"%s"}`,
			errors.InternalErrorCode, err.Error(), id))
	}
	return bytes
}

func (d *Describer3[T1, T2, T3, R]) Unmarshal(bytes []byte) ([]interface{}, error) {
	t1, t2, t3, err := d.unmarshal(bytes)
	if err != nil {
		return emptyArr, err
	}

	return []interface{}{t1, t2, t3}, nil
}

func (d *Describer3[T1, T2, T3, R]) unmarshal(bytes []byte) (T1, T2, T3, error) {
	var (
		t1 T1
		t2 T2
		t3 T3
	)

	if len(bytes) == 0 {
		return t1, t2, t3, nil
	}

	firstRune := rune(bytes[0])
	if firstRune == '{' {
		//named request
		var argsMap map[string]*json.RawMessage
		err := json.Unmarshal(bytes, &argsMap)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		firstRawArg, ok := argsMap[d.Description.ArgNames[0]]
		if !ok {
			return t1, t2, t3, errors.NewInvalidParamsError(fmt.Sprintf("[%s] not found", d.Description.ArgNames[0]))
		}
		err = json.Unmarshal(*firstRawArg, &t1)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsMap[d.Description.ArgNames[1]], &t2)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsMap[d.Description.ArgNames[2]], &t3)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		return t1, t2, t3, nil
	} else if firstRune == '[' {
		//positional request
		var argsArray []*json.RawMessage

		err := json.Unmarshal(bytes, &argsArray)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsArray[0], &t1)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsArray[1], &t2)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		err = json.Unmarshal(*argsArray[2], &t3)
		if err != nil {
			return t1, t2, t3, errors.NewInvalidRequestError(err.Error())
		}

		return t1, t2, t3, nil
	} else {
		return t1, t2, t3, errors.NewInvalidRequestError("incorrect request params type")
	}
}

type Describer struct {
	Description *ProcedureDescription
}

func (d *Describer) Describe() *ProcedureDescription {
	return d.Description
}
