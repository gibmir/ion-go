package tcp

import "github.com/gibmir/ion-go/ion-client/client"
import "github.com/gibmir/ion-go/ion-api/core"

type IonTcpClient struct {
}

type TcpRequest0[R any] struct {
}

func (r *TcpRequest0[R]) Call() <-chan *R {
	return nil
}

func (r *TcpRequest0[R]) Notification() {
}
func NoArg[R any](tcpClient *IonTcpClient, procedure core.JsonRemoteProcedure0[R]) *client.Request0[R] {
        var request client.Request0[R] = &TcpRequest0[R]{}
	return &request
}
