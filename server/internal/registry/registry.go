package registry

import (
	"github.com/gibmir/ion-go/api/pkg/describer"

	"github.com/gibmir/ion-go/server/internal/handle"
)

type RpcDescriptor struct {
	MethodHandle handle.MethodHandle
	Marshaller describer.Marshaller
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
