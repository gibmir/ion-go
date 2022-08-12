package core

import "fmt"

type Schema struct {
	Namespaces []Namespace
}

func (s *Schema) String() string {
	return fmt.Sprintf("Schema(Namespaces: %v)", s.Namespaces)
}

type Namespace struct {
	*SchemaElement
	Procedures    []Procedure
	Types         map[string]TypeDeclaration
}

func (n *Namespace) String() string {
	return fmt.Sprintf("Namespace(SchemaElement: %v, Procedures: %v, Types: %v)", n.SchemaElement, n.Procedures, n.Types)
}

type Procedure struct {
	*SchemaElement
	ArgumentTypes []PropertyType
	ReturnType    *PropertyType
}

func (p *Procedure) String() string {
	return fmt.Sprintf("Procedure(SchemaElement: %v, ArgumentTypes: %v, ReturnType: %v)", p.SchemaElement, p.ArgumentTypes, p.ReturnType)
}

type PropertyType struct {
	*SchemaElement
	TypeName      string
}

func (p *PropertyType) String() string {
	return fmt.Sprintf("PropertyType(TypeName:%v, SchemaElement:%v)", p.TypeName, p.SchemaElement)
}

type TypeParameter struct {
	SchemaElement *SchemaElement
}

func (t *TypeParameter) String() string {
	return fmt.Sprintf("TypeParameter(SchemaElement: %v)", t.SchemaElement)
}

type TypeDeclaration struct {
	PropertyTypes  []PropertyType
	TypeParameters []TypeParameter
	SchemaElement  *SchemaElement
}

func (t *TypeDeclaration) String() string {
	return fmt.Sprintf("TypeDeclaration(PropertyTypes: %v, TypeParameters: %v, SchemaElement: %v)", t.PropertyTypes, t.TypeParameters, t.SchemaElement)
}

type SchemaElement struct {
	Id          string
	Name        string
	Description string
}

func (s *SchemaElement) String() string {
	return fmt.Sprintf("SchemaElement(Id:%v, Name:%v, Description:%v)", s.Id, s.Name, s.Description)
}
