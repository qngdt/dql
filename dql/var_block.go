package dql

import (
	"fmt"
	"strings"
)

// VarBlock represents a variable block in a DQL query.
//
// A VarBlock is used to define variables in a query, including their criteria, attributes, and directives.
type VarBlock struct {
	// Name is the name of the variable block.
	Name string

	// Criteria defines the function or condition used in the variable block.
	Criteria string

	// Attributes is a list of attributes included in the variable block.
	Attributes []*Attribute

	// Directives is a list of directives applied to the variable block.
	Directives []string
}

// NewVarBlock creates a new VarBlock with the specified criteria.
//
// Parameters:
//   - criteria: The function or condition used in the variable block.
//
// Returns:
//   - A pointer to a VarBlock object.
//
// Example:
//   varBlock := NewVarBlock("has(user)")
//   fmt.Println(varBlock.String()) // Output: var(func: has(user)) { }
//
// See: https://dgraph.io/docs/dql/dql-syntax/dql-query/#variable-var-block
func NewVarBlock(criteria string) *VarBlock {
	return &VarBlock{
		Criteria: criteria,
	}
}

// WithName sets the name of the variable block.
//
// Parameters:
//   - name: The name to set for the variable block.
//
// Returns:
//   - The updated VarBlock object.
//
// Example:
//   varBlock := NewVarBlock("has(user)").WithName("userVar")
//   fmt.Println(varBlock.String()) // Output: userVar AS var(func: has(user)) { }
func (vb *VarBlock) WithName(name string) *VarBlock {
	vb.Name = name
	return vb
}

// WithDirectives adds one or more directives to the variable block.
//
// Parameters:
//   - directives: One or more directives to add to the variable block.
//
// Returns:
//   - The updated VarBlock object.
//
// Example:
//   varBlock := NewVarBlock("has(user)").
//       WithDirectives("@filter(eq(name, \"John\"))")
//   fmt.Println(varBlock.String()) // Output: var(func: has(user)) @filter(eq(name, "John")) { }
func (vb *VarBlock) WithDirectives(directives ...string) *VarBlock {
	for _, d := range directives {
		vb.Directives = append(vb.Directives, d)
	}
	return vb
}

// WithAttributes adds one or more attributes to the variable block.
//
// Parameters:
//   - attrs: One or more Attribute objects to add to the variable block.
//
// Returns:
//   - The updated VarBlock object.
//
// Example:
//   varBlock := NewVarBlock("has(user)").
//       WithAttributes(NewAttribute("name"), NewAttribute("age"))
//   fmt.Println(varBlock.String()) // Output: var(func: has(user)) { name age }
func (vb *VarBlock) WithAttributes(attrs ...*Attribute) *VarBlock {
	for _, a := range attrs {
		vb.Attributes = append(vb.Attributes, a)
	}
	return vb
}

// String generates a string representation of the variable block.
//
// The string includes the name (if set), criteria, directives, and attributes of the variable block.
//
// Returns:
//   - A string representation of the variable block.
func (vb *VarBlock) String() string {
	components := []string{}
	if vb.Name != "" {
		components = append(components, vb.Name, "AS")
	}
	components = append(components, "var", fmt.Sprintf("(func: %s)", vb.Criteria))
	for _, f := range vb.Directives {
		components = append(components, f)
	}
	components = append(components, "{")
	for _, attr := range vb.Attributes {
		components = append(components, attr.String())
	}
	components = append(components, "}")
	return strings.Join(components, " ")
}
