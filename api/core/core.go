package core

type JsonRemoteProcedure0[R any] interface {
	Describer
	Call() R
	Notify()
}

type JsonRemoteProcedure1[T, R any] interface {
	Describer
	PositionalCall(arg T) R
	NamedCall(arg T) R
}

type JsonRemoteProcedure2[T1, T2, R any] interface {
	Describer
	PositionalCall(arg1 T1, arg2 T2) R
	NamedCall(arg1 T1, arg2 T2) R
}

type JsonRemoteProcedure3[T1, T2, T3, R any] interface {
	Describer
	PositionalCall(arg1 T1, arg2 T2, arg3 T3) R
	NamedCall(arg1 T1, arg2 T2, arg3 T3) R
}

type ProcedureDescription struct {
	ProcedureName string
	ArgNames      []string
}

type Describer interface {
	Describe() *ProcedureDescription
}
