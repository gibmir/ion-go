package server

import (
	api "github.com/gibmir/ion-go/api/core"
)

type RpcDescriptor struct {
	MethodHandle MethodHandle
	Marshaller api.Marshaller
}

type Registry struct {
	Registry map[string]*RpcDescriptor
}

func (r *Registry) Add(procedureName string, descriptor *RpcDescriptor) {
	r.Registry[procedureName] = descriptor
}

func (r *Registry) Descriptor(procedureName string) (*RpcDescriptor, bool) {
	d, found := r.Registry[procedureName]
	return d, found
}
