package core

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	rpcServer *rpc.Server
}

func (s *Server) server(){
	
	s.rpcServer.ServeHTTP()
}
