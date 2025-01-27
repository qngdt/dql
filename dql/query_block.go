package dql

import (
	"fmt"
	"strings"
)

// QueryBlock represents a block of a DQL query.
//
// A QueryBlock typically includes a name, criteria, directives, and attributes.
type QueryBlock struct {
	// Name is the name of the query block.
	Name string

	// Criteria defines the function or condition used in the query block.
	Criteria []string

	// Directives is a list of directives applied to the query block.
	Directives []string

	// Attributes is a list of attributes included in the query block.
	Attributes []*Attribute
}

// NewQueryBlock creates a new QueryBlock.
//
// Parameters:
//   - name: The name of the query block.
//   - criteria: The root criteria of the query block.
//
// Returns:
//   - A pointer to a QueryBlock object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)")
//	fmt.Println(queryBlock.String()) // Output: getUser(func: has(user)) { }
//
// See: https://dgraph.io/docs/dql/dql-syntax/dql-query/#query-block
func NewQueryBlock(name string, criteria string) *QueryBlock {
	return &QueryBlock{
		Name:      name,
		Criteria: []string{criteria},
	}
}

// WithCriteria adds one or more criteria to the query block.
//
// Parameters:
//   - criteria: One or more criteria to add to the query block.
//
// Returns:
//   - The updated QueryBlock object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)").
//	    WithCriteria("orderasc: name@en")
//	fmt.Println(queryBlock.String()) // Output: getUser(func: has(user), orderasc: name@en) { }
func (qb *QueryBlock) WithCriteria(criteria ...string) *QueryBlock {
	for _, c := range criteria {
		qb.Criteria = append(qb.Criteria, c)
	}
	return qb
}


// WithDirectives adds one or more directives to the query block.
//
// Parameters:
//   - directives: One or more directives to add to the query block.
//
// Returns:
//   - The updated QueryBlock object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)").
//	    WithDirectives("@filter(eq(name, \"John\"))")
//	fmt.Println(queryBlock.String()) // Output: getUser(func: has(user)) @filter(eq(name, "John")) { }
func (qb *QueryBlock) WithDirectives(directives ...string) *QueryBlock {
	for _, d := range directives {
		qb.Directives = append(qb.Directives, d)
	}
	return qb
}

// WithAttributes adds one or more attributes to the query block.
//
// Parameters:
//   - attrs: One or more Attribute objects to add to the query block.
//
// Returns:
//   - The updated QueryBlock object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)").
//	    WithAttributes(NewAttribute("name"), NewAttribute("age"))
//	fmt.Println(queryBlock.String()) // Output: getUser(func: has(user)) { name age }
func (qb *QueryBlock) WithAttributes(attrs ...*Attribute) *QueryBlock {
	for _, a := range attrs {
		qb.Attributes = append(qb.Attributes, a)
	}
	return qb
}

// String generates a string representation of the query block.
//
// The string includes the name, criteria, directives, and attributes of the query block.
//
// Returns:
//   - A string representation of the query block.
func (qb *QueryBlock) String() string {
	components := []string{qb.Name, fmt.Sprintf("(func: %s)", strings.Join(qb.Criteria, ", "))}
	for _, f := range qb.Directives {
		components = append(components, f)
	}
	components = append(components, "{")
	for _, attr := range qb.Attributes {
		components = append(components, attr.String())
	}
	components = append(components, "}")

	return strings.Join(components, " ")
}
