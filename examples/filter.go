package examples

import (
	"fmt"

	"dql/dql"
)

func Filter() {
	directorFilmBlock := dql.NewAttribute("director.film").
		WithDirectives("@filter(allofterms(name@en, \"jones indiana\"))").
		WithAttributes(
			dql.NewAttribute("name@en"),
		)

	queryBlock := dql.NewQueryBlock("me", `eq(name@en, "Steven Spielberg")`).
		WithDirectives("@filter(has(director.film))").
		WithAttributes(
			dql.NewAttribute("name@en"),
			directorFilmBlock,
		)

	query := dql.NewQuery("", queryBlock)
	fmt.Println(query.PrettyPrint())
}
