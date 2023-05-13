package registry

import (
	api "github.com/gibmir/ion-go/api/core"

	"github.com/gibmir/ion-go/server/internal/handle"
)

type RpcDescriptor struct {
	MethodHandle handle.MethodHandle
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
