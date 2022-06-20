package client

import "github.com/gibmir/ion-go/ion-api/dto"

type Request0[R any] interface {
	Call(id string) (chan *R, chan *dto.ErrorResponse)

	Notification()
}

type Request1[T, R any] interface {
	PositionalCall(id string, argument *T) <-chan *R
	NamedCall(id string, argument *T) <-chan *R

	PositionalNotification(argument *T)
	NamedNotification(argument *T)
}

type Request2[T1, T2, R any] interface {
	PositionalCall(id string, argument1 *T1, argument2 *T2) <-chan *R
	NamedCall(id string, argument1 *T1, argument2 *T2) <-chan *R

	PositionalNotification(argument1 *T1, argument2 *T2)
	NamedNotification(argument1 *T1, argument2 *T2)
}
type Request3[T1, T2, T3, R any] interface {
	PositionalCall(id string, argument1 *T1, argument2 *T2, argument3 *T3) <-chan *R
	NamedCall(id string, argument1 *T1, argument2 *T2, argument3 *T3) <-chan *R

	PositionalNotification(argument1 *T1, argument2 *T2, argument3 *T3)
	NamedNotification(argument1 *T1, argument2 *T2, argument3 *T3)
}
