package schema

import (
	"fmt"
	"strings"
)

const (
	BooleanType = "boolean"
	StringType  = "string"
	NumberType  = "number"
	IntType     = "int"
	CustomType  = "custom"
	ListType    = "list"
	MapType     = "map"

	BoolGolangTypeName    = "bool"
	StringGolangTypeName  = "string"
	IntGolangTypeName     = "int"
	Float64GolangTypeName = "float64"
)

type TypeNode struct {
	Children []TypeNode
	Parent   *TypeNode
	Value    *strings.Builder
}

func From(typeName string) *TypeNode {

	return &TypeNode{}
}

func (node *TypeNode) String() string {
	return fmt.Sprintf("TypeNode(Children: %v, Parent: %v, Value: %s",
		node.Children, node.Parent, node.Value)
}
