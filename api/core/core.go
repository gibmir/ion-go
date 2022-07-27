package core

type JsonRemoteProcedure0[R any] interface {
	Named
	Call() R
	Notify()
}

type JsonRemoteProcedure1[T, R any] interface {
	Named
	PositionalCall(arg T) R
	NamedCall(arg T) R
}

type JsonRemoteProcedure2[T1, T2, R any] interface {
	Named
	PositionalCall(arg1 T1, arg2 T2) R
	NamedCall(arg1 T1, arg2 T2) R

}

type JsonRemoteProcedure3[T1, T2, T3, R any] interface {
	Named
	PositionalCall(arg1 T1, arg2 T2, arg3 T3) R
	NamedCall(arg1 T1, arg2 T2, arg3 T3) R
}

type Named interface {
	GetName() string
}
