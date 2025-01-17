package dql

import (
	"strings"
)

// Query represents a DQL query.
type Query struct {
	// Name is the name of the query.
	Name string

	// Params is a list of parameters for the query.
	Params []*Param

	// QueryBlocks is a list of query blocks that define the main body of the query.
	QueryBlocks []*QueryBlock

	// VarBlocks is a list of variable blocks used in the query.
	VarBlocks []*VarBlock

	// Fragments is a list of reusable fragments included in the query.
	Fragments []*Fragment
}

// NewQuery creates a new DQL query.
//
// Parameters:
//   - name: The name of the query.
//   - queryBlock: The default query block.
//
// Returns:
//   - A pointer to a Query object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)")
//	query := NewQuery("GetUserQuery", queryBlock)
//	fmt.Println(query.String()) // Output: query GetUserQuery { getUser(func: has(user)) { } }
//
// See: https://dgraph.io/docs/dql/dql-syntax/dql-query
func NewQuery(name string, queryBlock *QueryBlock) *Query {
	return &Query{
		Name:        name,
		QueryBlocks: []*QueryBlock{queryBlock},
	}
}

func (q *Query) concatenate() []string {
	components := []string{}
	if q.Name != "" {
		components = append(components, "query", q.Name)
	}
	if len(q.Params) != 0 {
		components = append(components, "(")
		paramComps := make([]string, len(q.Params))
		for i, param := range q.Params {
			paramComps[i] = param.String()
		}
		components = append(components, strings.Join(paramComps, ", "))
		components = append(components, ")")
	}
	components = append(components, "{")
	for _, vBlock := range q.VarBlocks {
		components = append(components, vBlock.String())
	}
	for _, qBlock := range q.QueryBlocks {
		components = append(components, qBlock.String())
	}
	components = append(components, "}")
	for _, f := range q.Fragments {
		components = append(components, f.String())
	}
	return components
}

// String generates the full query as a single-line string.
//
// Returns:
//   - A string representation of the query.
func (q Query) String() string {
	components := q.concatenate()
	return strings.Join(components, " ")
}

// PrettyPrint generates a formatted, human-readable version of the query with proper indentation.
//
// Returns:
//   - A formatted string representation of the query.
func (q Query) PrettyPrint() string {
	raw := q.String()
	var result strings.Builder
	indent := 0
	step := "  "
	for i := 0; i < len(raw); i++ {
		char := raw[i]
		switch char {
		case '{':
			result.WriteByte(char)
			result.WriteByte('\n')
			indent++
			result.WriteString(strings.Repeat(step, indent))
			i += 1 // Skip the " "
		case '}':
			result.WriteByte('\n')
			indent--
			result.WriteString(strings.Repeat(step, indent))
			result.WriteByte(char)
			if i < len(raw)-1 {
				peak := raw[i+2]
				if peak != '}' {
					result.WriteByte('\n')
				}
			}
			result.WriteString(strings.Repeat(step, indent))
			i += 1 // Skip the " "
		default:
			result.WriteByte(char)
		}
	}

	return result.String()
}

// WithParam adds one or more parameters to the query.
//
// Parameters:
//   - params: One or more Param objects to add to the query.
//
// Returns:
//   - The updated Query object.
//
// Example:
//
//	param := NewParam("id", "string").WithDefault("123")
//	query := NewQuery("GetUserQuery", NewQueryBlock("getUser", "has(user)")).
//	    WithParam(param)
//	fmt.Println(query.String()) // Output: query GetUserQuery($id: string = 123) { getUser(func: has(user)) { } }
func (q *Query) WithParam(params ...*Param) *Query {
	for _, p := range params {
		q.Params = append(q.Params, p)
	}
	return q
}

// WithVarBlocks adds one or more variable blocks to the query.
//
// Parameters:
//   - vbs: One or more VarBlock objects to add to the query.
//
// Returns:
//   - The updated Query object.
//
// Example:
//
//	varBlock := NewVarBlock("has(user)").WithName("userVar")
//	query := NewQuery("GetUserQuery", NewQueryBlock("getUser", "has(user)")).
//	    WithVarBlocks(varBlock)
//	fmt.Println(query.String()) // Output: query GetUserQuery { userVar AS var(func: has(user)) { } getUser(func: has(user)) { } }
func (q *Query) WithVarBlocks(vbs ...*VarBlock) *Query {
	for _, vb := range vbs {
		q.VarBlocks = append(q.VarBlocks, vb)
	}
	return q
}

// WithQueryBlocks adds one or more query blocks to the query.
//
// Parameters:
//   - qbs: One or more QueryBlock objects to add to the query.
//
// Returns:
//   - The updated Query object.
//
// Example:
//
//	queryBlock := NewQueryBlock("getUser", "has(user)")
//	query := NewQuery("GetUserQuery", queryBlock).
//	    WithQueryBlocks(NewQueryBlock("getPosts", "has(post)"))
//	fmt.Println(query.String()) // Output: query GetUserQuery { getUser(func: has(user)) { } getPosts(func: has(post)) { } }
func (q *Query) WithQueryBlocks(qbs ...*QueryBlock) *Query {
	for _, qb := range qbs {
		q.QueryBlocks = append(q.QueryBlocks, qb)
	}
	return q
}

// WithFragments adds one or more fragments to the query.
//
// Parameters:
//   - fragments: One or more Fragment objects to add to the query.
//
// Returns:
//   - The updated Query object.
//
// Example:
//
//	fragment := NewFragment("userFragment").WithAttributes(NewAttribute("name"), NewAttribute("age"))
//	query := NewQuery("GetUserQuery", NewQueryBlock("getUser", "has(user)")).
//	    WithFragments(fragment)
//	fmt.Println(query.String()) // Output: query GetUserQuery { getUser(func: has(user)) { ...userFragment } fragment userFragment { name age } }
func (q *Query) WithFragments(fragments ...*Fragment) *Query {
	for _, f := range fragments {
		q.Fragments = append(q.Fragments, f)
	}
	return q
}
