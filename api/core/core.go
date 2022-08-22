package core

type ProcedureDescription struct {
	ProcedureName string
	ArgNames      []string
}

type Type[T any] struct {
}

type Describer0[R any] struct {
	*Describer
	ReturnType *Type[R]
}

type Describer1[T, R any] struct {
	FirstArgument *Type[T]
	*Describer0[R]
}

type Describer2[T1, T2, R any] struct {
	FirstArgument  *Type[T1]
	SecondArgument *Type[T2]
	*Describer0[R]
}

type Describer3[T1, T2, T3, R any] struct {
	FirstArgument  *Type[T1]
	SecondArgument *Type[T2]
	ThirdArgument  *Type[T3]
	*Describer0[R]
}

type Describer struct {
	Description *ProcedureDescription
}

func (d *Describer) Describe() *ProcedureDescription {
	return d.Description
}
