# DQL Package

The `dql` package provides utilities for programmatically constructing and managing DQL (Dgraph Query Language) queries. It includes types and methods for building queries, managing parameters, variable blocks, query blocks, fragments, and attributes.

## Features

- Create and manage DQL queries with ease.
- Add parameters, variable blocks, query blocks, and fragments to queries.
- Support for reusable fragments and nested attributes.
- Generate compact or human-readable query strings.

## Installation

To use the `dql` package in your Go project, you can import it directly:

```go
import "path/to/dql"
```

## Usage

### Creating a Query

You can create a query using the `NewQuery` function and add various components like parameters, variable blocks, query blocks, and fragments.

```go
package main

import (
	"fmt"
	"dql"
)

func main() {
	// Create a query block
	queryBlock := dql.NewQueryBlock("getUser", "has(user)").
		WithAttributes(
			dql.NewAttribute("name"),
			dql.NewAttribute("age"),
		)

	// Create a parameter
	param := dql.NewParam("id", "string").WithDefault("123")

	// Create a query
	query := dql.NewQuery("GetUserQuery", queryBlock).
		WithParam(param)

	// Print the query
	fmt.Println(query.String())
    // Output:
    // query GetUserQuery($id: string = 123) { getUser(func: has(user)) { name age } }
}
```

### Adding Variable Blocks

Variable blocks can be added to a query using the `WithVarBlocks` method.

```go
varBlock := dql.NewVarBlock("has(user)").WithName("userVar")

query := dql.NewQuery("GetUserQuery", dql.NewQueryBlock("getUser", "has(user)")).
	WithVarBlocks(varBlock)

fmt.Println(query.String())
// Output:
// query GetUserQuery { userVar AS var(func: has(user)) { } getUser(func: has(user)) { } }
```

### Using Fragments

Fragments allow you to reuse parts of a query.

```go
fragment := dql.NewFragment("userFragment").
	WithAttributes(
		dql.NewAttribute("name"),
		dql.NewAttribute("age"),
	)

query := dql.NewQuery("GetUserQuery", dql.NewQueryBlock("getUser", "has(user)")).
	WithFragments(fragment)

fmt.Println(query.String())
// Output:
// query GetUserQuery { getUser(func: has(user)) { ...userFragment } fragment userFragment { name age } }
```


### Pretty Printing Queries

You can generate a human-readable version of the query using the `PrettyPrint` method.

```go
fmt.Println(query.PrettyPrint())
// Output:
// query GetUserQuery($id: string = 123) {
//   getUser(func: has(user)) {
//     name
//     age
//   }
// }
```

## API Reference

### Query

- `NewQuery(name string, queryBlock *QueryBlock) *Query`: Creates a new query.
- `WithParam(params ...*Param) *Query`: Adds parameters to the query.
- `WithVarBlocks(vbs ...*VarBlock) *Query`: Adds variable blocks to the query.
- `WithQueryBlocks(qbs ...*QueryBlock) *Query`: Adds query blocks to the query.
- `WithFragments(fragments ...*Fragment) *Query`: Adds fragments to the query.
- `String() string`: Generates a single-line string representation of the query.
- `PrettyPrint() string`: Generates a human-readable version of the query.

### QueryBlock

- `NewQueryBlock(name string, criteria string) *QueryBlock`: Creates a new query block.
- `WithCriteria(criteria ...string) *QueryBlock`: Adds one or more criteria to the query block.
- `WithDirectives(directives ...string) *QueryBlock`: Adds directives to the query block.
- `WithAttributes(attrs ...*Attribute) *QueryBlock`: Adds attributes to the query block.
- `String() string`: Generates a string representation of the query block.

### VarBlock

- `NewVarBlock(criteria string) *VarBlock`: Creates a new variable block.
- `WithName(name string) *VarBlock`: Sets the name of the variable block.
- `WithCriteria(criteria ...string) *VarBlock`: Adds one or more criteria to the variable block.
- `WithDirectives(directives ...string) *VarBlock`: Adds directives to the variable block.
- `WithAttributes(attrs ...*Attribute) *VarBlock`: Adds attributes to the variable block.
- `String() string`: Generates a string representation of the variable block.

### Fragment

- `NewFragment(name string) *Fragment`: Creates a new fragment.
- `WithAttributes(attrs ...*Attribute) *Fragment`: Adds attributes to the fragment.
- `String() string`: Generates a string representation of the fragment.

### Attribute

- `NewAttribute(name string) *Attribute`: Creates a new attribute.
- `WithAlias(alias string) *Attribute`: Sets an alias for the attribute.
- `WithDirectives(directives ...string) *Attribute`: Adds directives to the attribute.
- `WithAttributes(attributes ...*Attribute) *Attribute`: Adds nested attributes to the attribute.
- `String() string`: Generates a string representation of the attribute.

### Param

- `NewParam(name string, type string) *Param`: Creates a new parameter.
- `WithDefault(val string) *Param`: Sets a default value for the parameter.
- `String() string`: Generates a string representation of the parameter.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
