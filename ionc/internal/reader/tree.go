package reader

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gibmir/ion-go/ionc/internal/schema"
)

type Tree[T any] struct {
	root *Node[T]
}

func NewTypeTree(typeString string) *Tree[*strings.Builder] {
	stack := Stack[*Node[*strings.Builder]]{
		elements: make([]*Node[*strings.Builder], 0),
	}

	root := NewTypeNode(nil)
	stack.Push(root)

	for _, symbol := range typeString {
		if symbol == '<' {
			current, _ := stack.Peek()
			child := NewTypeNode(current)
			current.children = append(current.children, child)
			stack.Push(child)
		} else if symbol == ',' {
			current, _ := stack.Pop()
			parentChild := NewTypeNode(current.parent)
			current.parent.children = append(current.parent.children, parentChild)
			stack.Push(parentChild)
		} else if symbol == '>' {
			current, _ := stack.Peek()
			parentValue := current.parent.value.String()
			for !(current.value.String() == parentValue) {
				current, _ = stack.Pop()
			}
			stack.Push(current)
		} else {
			node, _ := stack.Peek()
			if node == nil {
				node = &Node[*strings.Builder]{value: &strings.Builder{}}
			}
			node.value.WriteRune(symbol)
			stack.Push(node)
		}
	}
	return &Tree[*strings.Builder]{
		root: root,
	}
}

func (t *Tree[T]) String() string {
	return fmt.Sprintf(
		`tree{
			root: %v
		}`, t.root)
}

type Node[T any] struct {
	value    T
	children []*Node[T]
	parent   *Node[T]
}

func NewTypeNode(parent *Node[*strings.Builder]) *Node[*strings.Builder] {
	return &Node[*strings.Builder]{
		value:    &strings.Builder{},
		children: make([]*Node[*strings.Builder], 0),
		parent:   parent,
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf(
		`node{
			value: %v,
			parent: %v,
			children: %v,
		}`, n.value, n.parent, n.children)
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	var result T
	if s.IsEmpty() {
		return result, false
	} else {
		index := len(s.elements) - 1
		result = s.elements[index]
		s.elements = s.elements[:index]
		return result, true
	}
}

func (s *Stack[T]) Peek() (T, bool) {
	var result T
	if s.IsEmpty() {
		return result, false
	} else {
		index := len(s.elements) - 1
		result = s.elements[index]
		return result, true
	}
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf(
		`stack{
			elements: %v,
		}`, s.elements)
}

// ToTypeName prepares type name from type tree
func ToTypeName(t *Tree[*strings.Builder]) string {
	traverse := Stack[*Node[*strings.Builder]]{}
	resultStack := Stack[*Node[*strings.Builder]]{}

	traverse.Push(t.root)
	resultStack.Push(t.root)

	for !traverse.IsEmpty() {
		current, _ := traverse.Pop()
		if len(current.children) > 0 {
			for _, child := range current.children {
				traverse.Push(child)
				resultStack.Push(child)
			}
		}
	}

	nodeTypeName := make(map[*Node[*strings.Builder]]string)
	for !resultStack.IsEmpty() {
		current, _ := resultStack.Pop()
		typeName := current.value.String()
		parametrization := make([]string, len(current.children))
		for i, child := range current.children {

			childTypeName, ok := nodeTypeName[child]
			if ok {
				parametrization[i] = childTypeName
			} else {
				parametrization[i] = child.value.String()

			}
		}
		if isCustomType(typeName) || isParametrizedType(typeName) {
			nodeTypeName[current] = buildParametrizedTypeName(typeName, parametrization)
		} else {
			nodeTypeName[current] = typeName
		}
	}

	return nodeTypeName[t.root]
}

func isCustomType(typeName string) bool {
	switch typeName {
	case schema.BooleanType:
		return false
	case schema.StringType:
		return false
	case schema.NumberType:
		return false
	case schema.IntType:
		return false
	case schema.ListType:
		return false
	case schema.MapType:
		return false
	default:
		return true
	}
}

func isParametrizedType(typeName string) bool {
	switch typeName {
	case schema.BooleanType:
		return false
	case schema.StringType:
		return false
	case schema.NumberType:
		return false
	case schema.IntType:
		return false
	default:
		return true
	}
}

func buildParametrizedTypeName(typeName string, typeParametrization []string) string {
	var result strings.Builder
	result.WriteString(prepareTypeName(typeName))
	if len(typeParametrization) == 0 {
		return result.String()
	}
	result.WriteRune('[')

	for i, parametrizationTypeName := range typeParametrization {
		preparedTypeName := prepareGolangType(parametrizationTypeName)
		// custom type should starts with title letter
		if isCustomType(parametrizationTypeName) {
			preparedTypeName = prepareTypeName(parametrizationTypeName)
		}
		result.WriteString(preparedTypeName)
		if i != len(typeParametrization)-1 {
			result.WriteRune(',')
		}
	}

	result.WriteRune(']')
	return result.String()
}

func prepareTypeName(typeName string) string {
	typeNameRunes := []rune(typeName)
	typeNameRunes[0] = unicode.ToUpper(typeNameRunes[0])
	return string(typeNameRunes)
}

func prepareGolangType(typeName string) string {
	switch typeName {
	case schema.BooleanType:
		return schema.BoolGolangTypeName
	case schema.StringType:
		return schema.StringGolangTypeName
	case schema.NumberType:
		return schema.Float64GolangTypeName
	case schema.IntType:
		return schema.IntGolangTypeName
	default:
		// return type as is
		return typeName
	}
}
