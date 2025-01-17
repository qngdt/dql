package dql

import "fmt"

// Param represents a parameter for a DQL query.
type Param struct {
	// Name is the name of the parameter.
	Name string

	// Type is the type of the parameter.
	Type string

	// Default is the default value of the parameter (optional).
	Default string
}

// NewParam creates a new parameter for a DQL query.
//
// Parameters:
//   - n: The name of the parameter.
//   - t: The type of the parameter.
//
// Returns:
//   - A pointer to a Param object.
//
// Example:
//   param := NewParam("id", "string")
//   fmt.Println(param.String()) // Output: id: string
//
// See: https://dgraph.io/docs/dql/dql-syntax/dql-query/#query-parameterization
func NewParam(n string, t string) *Param {
	return &Param{
		Name: n,
		Type: t,
	}
}

// WithDefault sets the default value for the parameter.
//
// Parameters:
//   - val: The default value to set.
//
// Returns:
//   - The updated Param object.
//
// Example:
//   param := NewParam("id", "string").WithDefault("123")
//   fmt.Println(param.String()) // Output: id: string = 123
func (p *Param) WithDefault(val string) *Param {
	p.Default = val
	return p
}

// String generates a string representation of the parameter.
//
// The string includes the parameter's name, type, and default value (if set).
//
// Returns:
//   - A string representation of the parameter.
func (p *Param) String() string {
	res := fmt.Sprintf("%s: %s", p.Name, p.Type)
	if p.Default != "" {
		res += fmt.Sprintf(" = %s", p.Default)
	}
	return res
}
