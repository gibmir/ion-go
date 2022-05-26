package schema

import "fmt"

type Schema struct {
	Procedures []Procedure
	Types      map[string]TypeDeclaration
}

func (s *Schema) String() string {
	return fmt.Sprintf("Schema(Procedures: %v, Types: %v)", s.Procedures, s.Types)
}

type Procedure struct {
	SchemaElement *SchemaElement
	ArgumentTypes []PropertyType
	ReturnType    *PropertyType
}

func (p *Procedure) String() string {
	return fmt.Sprintf("Procedure(SchemaElement: %v, ArgumentTypes: %v, ReturnType: %v)", p.SchemaElement, p.ArgumentTypes, p.ReturnType)
}

type PropertyType struct {
	TypeName      string
	SchemaElement *SchemaElement
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
