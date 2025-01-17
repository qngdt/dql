package dql

import "strings"

// Fragment represents a reusable fragment in a DQL query.
//
// A Fragment is a named reusable block that contains a list of attributes.
type Fragment struct {
	// Name is the name of the fragment.
	Name string

	// Attributes is a list of attributes included in the fragment.
	Attributes []*Attribute
}

// NewFragment creates a new Fragment with the specified name.
//
// Parameters:
//   - name: The name of the fragment.
//
// Returns:
//   - A pointer to a Fragment object.
//
// Example:
//   fragment := NewFragment("userFragment")
//   fmt.Println(fragment.String()) // Output: fragment userFragment { }
//
// See: https://dgraph.io/docs/query-language/fragments/
func NewFragment(name string) *Fragment {
	return &Fragment{
		Name: name,
	}
}

// WithAttributes adds one or more attributes to the fragment.
//
// Parameters:
//   - attrs: One or more Attribute objects to add to the fragment.
//
// Returns:
//   - The updated Fragment object.
//
// Example:
//   fragment := NewFragment("userFragment").
//       WithAttributes(NewAttribute("name"), NewAttribute("age"))
//   fmt.Println(fragment.String()) // Output: fragment userFragment { name age }
func (f *Fragment) WithAttributes(attrs ...*Attribute) *Fragment {
	for _, a := range attrs {
		f.Attributes = append(f.Attributes, a)
	}
	return f
}

// String generates a string representation of the fragment.
//
// The string includes the fragment's name and its attributes.
//
// Returns:
//   - A string representation of the fragment.
func (f *Fragment) String() string {
	components := []string{"fragment", f.Name}
	components = append(components, "{")
	for _, attr := range f.Attributes {
		components = append(components, attr.String())
	}
	components = append(components, "}")
	return strings.Join(components, " ")
}
