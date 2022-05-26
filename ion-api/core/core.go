package core

type JsonRemoteProcedure0[R any] interface {
	Call() R
}

type JsonRemoteProcedure1[T,R any] interface {
	Call(arg T) R
}

type JsonRemoteProcedure2[T1,T2,R any] interface {
	Call(arg1 T1, arg2 T2) R
}

type JsonRemoteProcedure3[T1,T2, T3,R any] interface {
	Call(arg1 T1, arg2 T2, arg3 T3) R
}
