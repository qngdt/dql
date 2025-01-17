package dql

import "strings"

// Attribute represents an attribute in a DQL query.
//
// An Attribute can have an alias, directives, and nested attributes.
type Attribute struct {
	// Alias is an optional alias for the attribute.
	Alias string

	// Name is the name of the attribute.
	Name string

	// Directives is a list of directives applied to the attribute.
	Directives []string

	// Attributes is a list of nested attributes under this attribute.
	Attributes []*Attribute
}

// NewAttribute creates a new Attribute with the specified name.
//
// Parameters:
//   - name: The name of the attribute.
//
// Returns:
//   - A pointer to an Attribute object.
//
// Example:
//
//	attr := NewAttribute("name")
//	fmt.Println(attr.String()) // Output: name
func NewAttribute(name string) *Attribute {
	return &Attribute{
		Name: name,
	}
}

// WithDirectives adds one or more directives to the attribute.
//
// Parameters:
//   - directives: One or more directives to add to the attribute.
//
// Returns:
//   - The updated Attribute object.
//
// Example:
//
//	attr := NewAttribute("name").WithDirectives("@filter(eq(name, \"John\"))")
//	fmt.Println(attr.String()) // Output: name @filter(eq(name, "John"))
func (a *Attribute) WithDirectives(directives ...string) *Attribute {
	for _, d := range directives {
		a.Directives = append(a.Directives, d)
	}
	return a
}

// WithAttributes adds one or more nested attributes to the attribute.
//
// Parameters:
//   - attributes: One or more Attribute objects to add as nested attributes.
//
// Returns:
//   - The updated Attribute object.
//
// Example:
//
//	attr := NewAttribute("person").
//	    WithAttributes(NewAttribute("name"), NewAttribute("age"))
//	fmt.Println(attr.String()) // Output: person { name age }
func (a *Attribute) WithAttributes(attributes ...*Attribute) *Attribute {
	for _, attr := range attributes {
		a.Attributes = append(a.Attributes, attr)
	}
	return a
}

// String generates a string representation of the attribute.
//
// The string includes the alias (if set), name, directives, and any nested attributes.
//
// Returns:
//   - A string representation of the attribute.
func (a *Attribute) String() string {
	components := []string{}
	if a.Alias != "" {
		components = append(components, a.Alias, ":")
	}
	components = append(components, a.Name)
	for _, f := range a.Directives {
		components = append(components, f)
	}
	if len(a.Attributes) != 0 {
		components = append(components, "{")
		for _, attr := range a.Attributes {
			components = append(components, attr.String())
		}
		components = append(components, "}")
	}
	return strings.Join(components, " ")
}
