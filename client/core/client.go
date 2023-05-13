package core

type Request0[R any] interface {
	Call(id string, responseChannel chan<- R, errorChannel chan<- error)
	Notify()
}

type Request1[T, R any] interface {
	PositionalCall(id string, argument T, responseChannel chan<- R, errorChannel chan<- error)
	PositionalNotification(argument T)
}

type Request2[T1, T2, R any] interface {
	PositionalCall(id string, firstArgument T1, secondArgument T2, responseChannel chan<- R, errorChannel chan<- error)
	PositionalNotification(firstArgument T1, secondArgument T2)
}

type Request3[T1, T2, T3, R any] interface {
	PositionalCall(id string, firstArgument T1, secondArgument T2, thirdArgument T3, responseChannel chan<- R, errorChannel chan<- error)
	PositionalNotification(firstArgument T1, secondArgument T2, thirdArgument T3)
}
